package subscription

import (
	operatorsv1beta1 "github.com/open-cluster-management/multicloudhub-operator/pkg/apis/operators/v1beta1"
	"github.com/open-cluster-management/multicloudhub-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// GRC overrides the grc chart
func GRC(m *operatorsv1beta1.MultiClusterHub, cache utils.CacheSpec) *unstructured.Unstructured {
	sub := &Subscription{
		Name:      "grc",
		Namespace: m.Namespace,
		Overrides: map[string]interface{}{
			"pullSecret": m.Spec.ImagePullSecret,
			"grcuiapi": map[string]interface{}{
				"image": map[string]interface{}{
					"pullPolicy": m.Spec.ImagePullPolicy,
				},
			},
			"grcui": map[string]interface{}{
				"image": map[string]interface{}{
					"pullPolicy": m.Spec.ImagePullPolicy,
				},
			},
			"compliance": map[string]interface{}{
				"image": map[string]interface{}{
					"pullPolicy": m.Spec.ImagePullPolicy,
				},
			},
			"hubconfig": map[string]interface{}{
				"replicaCount": m.Spec.ReplicaCount,
				"nodeSelector": m.Spec.NodeSelector,
			},
			"global": map[string]interface{}{
				"imageOverrides": cache.ImageOverrides,
			},
		},
	}

	return newSubscription(m, sub)
}
