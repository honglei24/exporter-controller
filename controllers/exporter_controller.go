/*
Copyright 2020 honglei.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"reflect"

	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appv1 "exporter-controller/api/v1"
)

// ExporterReconciler reconciles a Exporter object
type ExporterReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=app.ci.com,resources=exporters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app.ci.com,resources=exporters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps;extensions,resources=deployments,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups="",resources=services;events;configmaps;persistentvolumeclaims,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;list;watch;create;update;delete
func (r *ExporterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("exporter", req.NamespacedName)

	// your logic here
	// 声明 finalizer 字段，类型为字符串
	myFinalizerName := "storage.finalizers.exporters.app.ci.com"

	instance := &appv1.Exporter{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if instance.DeletionTimestamp != nil {
		log.Info("Get deleted App, clean up subResources.")

		// 如果不为 0 ，则对象处于删除中
		if containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			// 如果存在 finalizer 且与上述声明的 finalizer 匹配，那么执行对应 hook 逻辑
			if err := r.deleteExternalResources(instance); err != nil {
				// 如果删除失败，则直接返回对应 err，controller 会自动执行重试逻辑
				return ctrl.Result{}, err
			}

			// 如果对应 hook 执行成功，那么清空 finalizers， k8s 删除对应资源
			instance.ObjectMeta.Finalizers = removeString(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Update(context.Background(), instance); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, nil
	}

	if !containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, myFinalizerName)
		if err := r.Update(context.Background(), instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	deploy := getDeployFromExporter(instance)

	found := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace}, found)

	if err != nil && errors.IsNotFound(err) {
		log.Info("Old Deployment NotFound and Creating new one", "namespace", deploy.Namespace, "name", deploy.Name)
		if err = r.Create(ctx, deploy); err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		log.Error(err, "Get Deployment info Error", "namespace", deploy.Namespace, "name", deploy.Name)
		return ctrl.Result{}, err
	} else if !reflect.DeepEqual(deploy.Spec, found.Spec) {
		// Update the found object and write the result back if there are any changes
		found.Spec = deploy.Spec
		log.Info("Old deployment changed and Updating Deployment to reconcile", "namespace", deploy.Namespace, "name", deploy.Name)
		err = r.Update(ctx, found)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *ExporterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.Exporter{}).
		Complete(r)
}

func (r *ExporterReconciler) deleteExternalResources(instance *appv1.Exporter) error {
	//
	// 删除 Exporter关联的外部资源逻辑
	//
	// 需要确保实现是幂等的
	ctx := context.Background()
	deploy := getDeployFromExporter(instance)
	r.Log.Info("Delete exporter", "exporter", instance)
	err := r.Delete(ctx, deploy)
	return err
}

func getDeployFromExporter(instance *appv1.Exporter) *appsv1.Deployment {
	labels := make(map[string]string)
	labels["app"] = instance.Name

	deploySpec := appsv1.DeploymentSpec{}
	deploySpec = instance.Spec.Deploy
	deploySpec.Selector = &metav1.LabelSelector{MatchLabels: labels}

	sidecar := corev1.Container{
		Name:            "nginx",
		Image:           "nginx:latest",
		ImagePullPolicy: corev1.PullIfNotPresent,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "nginx-default-conf",
				MountPath: "/etc/nginx/conf.d",
			},
		},
	}
	volume := corev1.Volume{
		Name: "nginx-default-conf",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: "nginx-conf"},
			},
		},
	}
	if deploySpec.Template.ObjectMeta.Annotations == nil {
		deploySpec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	// 设置注解
	setDefaultValue(deploySpec.Template.ObjectMeta.Annotations, "prometheus.io/scrape", "true")
	setDefaultValue(deploySpec.Template.ObjectMeta.Annotations, "prometheus.io/path", "/metrics")
	setDefaultValue(deploySpec.Template.ObjectMeta.Annotations, "prometheus.io/port", "8081")

	deploySpec.Template.Spec.Containers = append(deploySpec.Template.Spec.Containers, sidecar)
	deploySpec.Template.Spec.Volumes = append(deploySpec.Template.Spec.Volumes, volume)

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name + "-deploy",
			Namespace: instance.Namespace,
			Labels:    labels,
		},
		Spec: deploySpec,
	}

	return deploy
}
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func setDefaultValue(m map[string]string, key string, value string) {
	if _, ok := m[key]; !ok {
		m[key] = value
	}
}
