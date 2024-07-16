package monolithic

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/grafana/tempo-operator/apis/tempo/v1alpha1"
	"github.com/grafana/tempo-operator/internal/manifests/gateway"
	"github.com/grafana/tempo-operator/internal/manifests/manifestutils"
	"github.com/grafana/tempo-operator/internal/manifests/naming"
)

const (
	componentName = "serviceaccount"
)

// BuildServiceAccount creates a Kubernetes service account for Tempo.
func BuildServiceAccount(opts Options) *corev1.ServiceAccount {
	tempo := opts.Tempo
	labels := ComponentLabels(componentName, tempo.Name)
	var annotations map[string]string
	if tempo.Spec.Multitenancy.IsGatewayEnabled() && tempo.Spec.Multitenancy.Mode == v1alpha1.ModeOpenShift {
		annotations = gateway.BuildServiceAccountAnnotations(tempo.Spec.Multitenancy.TenantsSpec, naming.Name("jaegerui", tempo.Name))
	}

	if opts.StorageParams.S3 != nil && opts.StorageParams.S3.ShortLived != nil {
		awsAnnotations := manifestutils.S3AWSSTSAnnotations(*opts.StorageParams.S3.ShortLived)
		if annotations == nil {
			annotations = map[string]string{}
		}
		for k, v := range awsAnnotations {
			annotations[k] = v
		}
	}
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:        naming.DefaultServiceAccountName(tempo.Name),
			Namespace:   tempo.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
	}
}
