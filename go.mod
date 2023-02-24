module exporter-controller

go 1.13

require (
	github.com/GehirnInc/crypt v0.0.0-20190301055215-6c0105aabd46
	github.com/go-logr/logr v0.2.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.7.0
	github.com/prometheus/client_golang v1.2.0 // indirect
	k8s.io/api v0.20.0-alpha.2
	k8s.io/apimachinery v0.20.0-alpha.2
	k8s.io/client-go v0.20.0-alpha.2
	sigs.k8s.io/controller-runtime v0.4.0
)

replace github.com/prometheus/prometheus => ./vendor/github.com/prometheus/prometheus_v2.15.2
