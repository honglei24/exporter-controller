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
	"context"
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	prometheus_conf "github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	//"reflect"
	//"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appv1 "exporter-controller/api/v1"
)

// PrometheusReconciler reconciles a Prometheus object
type PrometheusReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=app.ci.com,resources=prometheus,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app.ci.com,resources=prometheus/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps;extensions,resources=deployments,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups="",resources=services;events;configmaps;persistentvolumeclaims,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;list;watch;create;update;delete

func (r *PrometheusReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("prometheus", req.NamespacedName)

	// your logic here
	// 声明 finalizer 字段，类型为字符串
	//myFinalizerName := "storage.finalizers.prometheus.app.ci.com"
	//

	//instance := &appv1.Prometheus{}
	//err := r.Get(ctx, req.NamespacedName, instance)

	configmap := &corev1.ConfigMap{}
	err := r.Get(ctx, types.NamespacedName{Name: "prometheus-for-app", Namespace: "default"}, configmap)
	//
	//log.Info("=================deploy", "deploy", "###1")
	//log.Info("=================deploy", "deploy namespace", instance.Namespace)
	if err != nil {
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("config detail:", "detail", configmap.Data)
	prom_conf := &prometheus_conf.Config{}
	_ = prom_conf.GlobalConfig.ScrapeInterval.Set("15s")
	_ = prom_conf.GlobalConfig.EvaluationInterval.Set("15s")
	prom_conf.RuleFiles = append(prom_conf.RuleFiles, "prometheus.rules.yml")
	scrapeConfig := &prometheus_conf.ScrapeConfig{}
	scrapeConfig.JobName = "test"
	scrapeConfig.HTTPClientConfig = config_util.HTTPClientConfig{
		BasicAuth: &config_util.BasicAuth{
			Username: "admin",
			Password: "123456",
		},
	}
	scrapeConfig.ServiceDiscoveryConfig.StaticConfigs = []*targetgroup.Group{
		{
			Targets: []model.LabelSet{
				{model.AddressLabel: "192.168.2.8:32000"},
			},
			Labels: model.LabelSet{
				"group": "production",
			},
		},
	}
	prom_conf.ScrapeConfigs = append(prom_conf.ScrapeConfigs, scrapeConfig)
	//log.Info("prometheus.yml detail:", "prometheus_conf", prom_conf)
	//err = yaml.Unmarshal([]byte(configmap.Data["prometheus.yml"]), prom_conf)
	//if err != nil {
	//	return ctrl.Result{}, err
	//}
	//log.Info("prometheus.yml detail:", "prom_conf.GlobalConfig", prom_conf.GlobalConfig)
	//log.Info("prometheus.yml detail:", "prom_conf.ScrapeConfigs", prom_conf.ScrapeConfigs)
	//log.Info("prometheus.yml detail:", "prom_conf.RuleFiles", prom_conf.RuleFiles)
	//log.Info("prometheus.yml detail:", "prom_conf", prom_conf.String())

	config_test2 := &corev1.ConfigMap{}
	config_test2.Namespace = "default"
	config_test2.Name = "test"
	config_test2.Data = make(map[string]string)
	config_test2.Data["prometheus.yml"] = prom_conf.String()
	config_test2.Data["alerting_rules.yml"] = "{}"
	config_test2.Data["alerts"] = "{}"
	config_test2.Data["recording_rules.yml"] = "{}"
	config_test2.Data["rules"] = "{}"
	err = r.Create(ctx, config_test2)

	//
	//if instance.ObjectMeta.DeletionTimestamp.IsZero() {
	//	if !containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
	//		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, myFinalizerName)
	//		if err := r.Update(ctx, instance); err != nil {
	//			return ctrl.Result{}, err
	//		}
	//	}
	//} else {
	//	log.Info("Get deleted App, clean up subResources.")
	//
	//	if containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
	//		if err := r.deleteExternalResources(instance); err != nil {
	//			return ctrl.Result{}, err
	//		}
	//
	//		instance.ObjectMeta.Finalizers = removeString(instance.ObjectMeta.Finalizers, myFinalizerName)
	//		if err := r.Update(ctx, instance); err != nil {
	//			return ctrl.Result{}, err
	//		}
	//	}
	//}

	//deploy := getDeployFromPrometheus(instance)
	//
	//found := &appsv1.Deployment{}
	//err = r.Get(ctx, req.NamespacedName, found)
	//if err != nil {
	//	// Object not found, return.
	//	// Created objects are automatically garbage collected.
	//	// For additional cleanup logic use finalizers.
	//	if errors.IsNotFound(err) {
	//		if err = r.Create(ctx, deploy); err != nil {
	//			return ctrl.Result{}, err
	//		}
	//	}
	//	return reconcile.Result{}, client.IgnoreNotFound(err)
	//}
	//
	//if !reflect.DeepEqual(deploy.Spec, found.Spec) {
	//	// Update the found object and write the result back if there are any changes
	//	found.Spec = deploy.Spec
	//	log.Info("Old deployment changed and Updating Deployment to reconcile", "namespace", deploy.Namespace, "name", deploy.Name)
	//	err = r.Update(ctx, found)
	//	if err != nil {
	//		return ctrl.Result{}, err
	//	}
	//}

	return ctrl.Result{}, nil
}

func (r *PrometheusReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.Prometheus{}).
		Complete(r)
}

//func (r *PrometheusReconciler) deleteExternalResources(instance *appv1.Prometheus) error {
//	//
//	// 删除 Exporter关联的外部资源逻辑
//	//
//	// 需要确保实现是幂等的
//	ctx := context.Background()
//	deploy := getDeployFromPrometheus(instance)
//	r.Log.Info("Delete prometheus", "prometheus", instance)
//	err := r.Delete(ctx, deploy)
//	return err
//}

//func getDeployFromPrometheus(instance *appv1.Prometheus) *appsv1.Deployment {
//	labels := make(map[string]string)
//	labels["app"] = instance.Name
//
//	deploySpec := appsv1.DeploymentSpec{}
//	deploySpec = instance.Spec.Deploy
//	deploySpec.Selector = &metav1.LabelSelector{MatchLabels: labels}
//
//	prometheus := corev1.Container{
//		Name:            "prometheus",
//		Image:           "prom/prometheus:latest",
//		ImagePullPolicy: corev1.PullIfNotPresent,
//		VolumeMounts: []corev1.VolumeMount{
//			{
//				Name:      "config-volume",
//				MountPath: "/etc/config",
//			},
//			{
//				Name:      "storage-volume",
//				MountPath: "/data",
//			},
//		},
//		Args: []string{
//			"--storage.tsdb.retention.time=5d",
//			"--config.file=/etc/config/prometheus.yml",
//			"--storage.tsdb.path=/data",
//			"--web.console.libraries=/etc/prometheus/console_libraries",
//			"--web.console.templates=/etc/prometheus/consoles",
//			"--web.enable-lifecycle",
//		},
//	}
//
//	sidecar := corev1.Container{
//		Name:            "prometheus-server-configmap-reload",
//		Image:           "jimmidyson/configmap-reload:v0.2.2",
//		ImagePullPolicy: corev1.PullIfNotPresent,
//		VolumeMounts: []corev1.VolumeMount{
//			{
//				Name:      "config-volume",
//				MountPath: "/etc/config",
//				//ReadOnly: true,
//			},
//		},
//		Args: []string{"--volume-dir=/etc/config", "--webhook-url=http://127.0.0.1:9090/-/reload"},
//	}
//
//	volume := corev1.Volume{
//		Name: "config-volume",
//		VolumeSource: corev1.VolumeSource{
//			ConfigMap: &corev1.ConfigMapVolumeSource{
//				LocalObjectReference: corev1.LocalObjectReference{Name: "prometheus-for-app"},
//			},
//		},
//	}
//	emptyDir := corev1.Volume{
//		Name: "storage-volume",
//		VolumeSource: corev1.VolumeSource{
//			EmptyDir: &corev1.EmptyDirVolumeSource{
//				Medium: corev1.StorageMediumDefault,
//			}},
//	}
//	if deploySpec.Template.ObjectMeta.Annotations == nil {
//		deploySpec.Template.ObjectMeta.Annotations = make(map[string]string)
//	}
//	// 设置注解
//	setDefaultValue(deploySpec.Template.ObjectMeta.Annotations, "prometheus.app.ci.com", "true")
//
//	deploySpec.Template.Spec.Containers = append(deploySpec.Template.Spec.Containers, prometheus)
//	deploySpec.Template.Spec.Containers = append(deploySpec.Template.Spec.Containers, sidecar)
//	deploySpec.Template.Spec.Volumes = append(deploySpec.Template.Spec.Volumes, volume)
//	deploySpec.Template.Spec.Volumes = append(deploySpec.Template.Spec.Volumes, emptyDir)
//
//	deploy := &appsv1.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "deploy-" + instance.Name,
//			Namespace: instance.Namespace,
//			Labels:    labels,
//		},
//		Spec: deploySpec,
//	}
//
//	return deploy
//}
