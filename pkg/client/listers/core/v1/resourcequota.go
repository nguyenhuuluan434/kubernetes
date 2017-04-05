/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
)

// ResourceQuotaLister helps list ResourceQuotas.
type ResourceQuotaLister interface {
	// List lists all ResourceQuotas in the indexer.
	List(selector labels.Selector) (ret []*v1.ResourceQuota, err error)
	// ResourceQuotas returns an object that can list and get ResourceQuotas.
	ResourceQuotas(namespace string) ResourceQuotaNamespaceLister
	ResourceQuotaListerExpansion
}

// resourceQuotaLister implements the ResourceQuotaLister interface.
type resourceQuotaLister struct {
	indexer cache.Indexer
}

// NewResourceQuotaLister returns a new ResourceQuotaLister.
func NewResourceQuotaLister(indexer cache.Indexer) ResourceQuotaLister {
	return &resourceQuotaLister{indexer: indexer}
}

// List lists all ResourceQuotas in the indexer.
func (s *resourceQuotaLister) List(selector labels.Selector) (ret []*v1.ResourceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResourceQuota))
	})
	return ret, err
}

// ResourceQuotas returns an object that can list and get ResourceQuotas.
func (s *resourceQuotaLister) ResourceQuotas(namespace string) ResourceQuotaNamespaceLister {
	return resourceQuotaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceQuotaNamespaceLister helps list and get ResourceQuotas.
type ResourceQuotaNamespaceLister interface {
	// List lists all ResourceQuotas in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ResourceQuota, err error)
	// Get retrieves the ResourceQuota from the indexer for a given namespace and name.
	Get(name string) (*v1.ResourceQuota, error)
	ResourceQuotaNamespaceListerExpansion
}

// resourceQuotaNamespaceLister implements the ResourceQuotaNamespaceLister
// interface.
type resourceQuotaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceQuotas in the indexer for a given namespace.
func (s resourceQuotaNamespaceLister) List(selector labels.Selector) (ret []*v1.ResourceQuota, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResourceQuota))
	})
	return ret, err
}

// Get retrieves the ResourceQuota from the indexer for a given namespace and name.
func (s resourceQuotaNamespaceLister) Get(name string) (*v1.ResourceQuota, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("resourcequota"), name)
	}
	return obj.(*v1.ResourceQuota), nil
}
