/*
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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	serverlessreleasev1alpha1 "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease/v1alpha1"
	versioned "github.com/serverless-operator/serverless-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/serverless-operator/serverless-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/serverless-operator/serverless-operator/pkg/client/listers/serverlessrelease/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServerlessReleaseInformer provides access to a shared informer and lister for
// ServerlessReleases.
type ServerlessReleaseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServerlessReleaseLister
}

type serverlessReleaseInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServerlessReleaseInformer constructs a new informer for ServerlessRelease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServerlessReleaseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServerlessReleaseInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServerlessReleaseInformer constructs a new informer for ServerlessRelease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServerlessReleaseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReleaseV1alpha1().ServerlessReleases(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReleaseV1alpha1().ServerlessReleases(namespace).Watch(options)
			},
		},
		&serverlessreleasev1alpha1.ServerlessRelease{},
		resyncPeriod,
		indexers,
	)
}

func (f *serverlessReleaseInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServerlessReleaseInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serverlessReleaseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&serverlessreleasev1alpha1.ServerlessRelease{}, f.defaultInformer)
}

func (f *serverlessReleaseInformer) Lister() v1alpha1.ServerlessReleaseLister {
	return v1alpha1.NewServerlessReleaseLister(f.Informer().GetIndexer())
}
