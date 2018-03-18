package framework

import (
	"github.com/appscode/go/crypto/rand"
	cs "github.com/kubedb/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"
	"k8s.io/client-go/kubernetes"
)

type Framework struct {
	kubeClient   kubernetes.Interface
	extClient    cs.KubedbV1alpha1Interface
	namespace    string
	name         string
	StorageClass string
}

func New(kubeClient kubernetes.Interface, extClient cs.KubedbV1alpha1Interface, storageClass string) *Framework {
	return &Framework{
		kubeClient:   kubeClient,
		extClient:    extClient,
		name:         "redis-operator",
		namespace:    rand.WithUniqSuffix("redis"),
		StorageClass: storageClass,
	}
}

func (f *Framework) Invoke() *Invocation {
	return &Invocation{
		Framework: f,
		app:       rand.WithUniqSuffix("redis-e2e"),
	}
}

type Invocation struct {
	*Framework
	app string
}