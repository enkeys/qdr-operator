/*
Copyright 2019 The Interconnectedcloud Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/enkeys/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InterconnectLister helps list Interconnects.
type InterconnectLister interface {
	// List lists all Interconnects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Interconnect, err error)
	// Interconnects returns an object that can list and get Interconnects.
	Interconnects(namespace string) InterconnectNamespaceLister
	InterconnectListerExpansion
}

// interconnectLister implements the InterconnectLister interface.
type interconnectLister struct {
	indexer cache.Indexer
}

// NewInterconnectLister returns a new InterconnectLister.
func NewInterconnectLister(indexer cache.Indexer) InterconnectLister {
	return &interconnectLister{indexer: indexer}
}

// List lists all Interconnects in the indexer.
func (s *interconnectLister) List(selector labels.Selector) (ret []*v1alpha1.Interconnect, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Interconnect))
	})
	return ret, err
}

// Interconnects returns an object that can list and get Interconnects.
func (s *interconnectLister) Interconnects(namespace string) InterconnectNamespaceLister {
	return interconnectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InterconnectNamespaceLister helps list and get Interconnects.
type InterconnectNamespaceLister interface {
	// List lists all Interconnects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Interconnect, err error)
	// Get retrieves the Interconnect from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Interconnect, error)
	InterconnectNamespaceListerExpansion
}

// interconnectNamespaceLister implements the InterconnectNamespaceLister
// interface.
type interconnectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Interconnects in the indexer for a given namespace.
func (s interconnectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Interconnect, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Interconnect))
	})
	return ret, err
}

// Get retrieves the Interconnect from the indexer for a given namespace and name.
func (s interconnectNamespaceLister) Get(name string) (*v1alpha1.Interconnect, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("interconnect"), name)
	}
	return obj.(*v1alpha1.Interconnect), nil
}
