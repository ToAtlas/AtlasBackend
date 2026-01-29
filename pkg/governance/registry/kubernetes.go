package registry

import (
	"fmt"

	conf "github.com/ToAtlas/AtlasBackend/api/gen/go/conf/v1"
	"github.com/ToAtlas/AtlasBackend/pkg/k8s"
	kuberegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	"github.com/go-kratos/kratos/v2/registry"
)

func NewKubernetesRegistry(c *conf.KubernetesConfig) registry.Registrar {
	if c == nil || !c.Enable {
		return nil
	}

	clientset, err := k8s.NewClientset()
	if err != nil {
		panic(fmt.Sprintf("failed to create kubernetes clientset: %v", err))
	}

	return kuberegistry.NewRegistry(clientset, k8s.GetCurrentNamespace())
}

func NewKubernetesDiscovery(c *conf.KubernetesConfig) registry.Discovery {
	if c == nil || !c.Enable {
		return nil
	}

	clientset, err := k8s.NewClientset()
	if err != nil {
		panic(fmt.Sprintf("failed to create kubernetes clientset: %v", err))
	}

	return kuberegistry.NewRegistry(clientset, k8s.GetCurrentNamespace())
}
