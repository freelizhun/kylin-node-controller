/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/kylin/kylin-node-controller/pkg/apis/crd/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KylinNodeLister helps list KylinNodes.
type KylinNodeLister interface {
	// List lists all KylinNodes in the indexer.
	List(selector labels.Selector) (ret []*v1.KylinNode, err error)
	// KylinNodes returns an object that can list and get KylinNodes.
	KylinNodes(namespace string) KylinNodeNamespaceLister
	KylinNodeListerExpansion
}

// kylinNodeLister implements the KylinNodeLister interface.
type kylinNodeLister struct {
	indexer cache.Indexer
}

// NewKylinNodeLister returns a new KylinNodeLister.
func NewKylinNodeLister(indexer cache.Indexer) KylinNodeLister {
	return &kylinNodeLister{indexer: indexer}
}

// List lists all KylinNodes in the indexer.
func (s *kylinNodeLister) List(selector labels.Selector) (ret []*v1.KylinNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KylinNode))
	})
	return ret, err
}

// KylinNodes returns an object that can list and get KylinNodes.
func (s *kylinNodeLister) KylinNodes(namespace string) KylinNodeNamespaceLister {
	return kylinNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KylinNodeNamespaceLister helps list and get KylinNodes.
type KylinNodeNamespaceLister interface {
	// List lists all KylinNodes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.KylinNode, err error)
	// Get retrieves the KylinNode from the indexer for a given namespace and name.
	Get(name string) (*v1.KylinNode, error)
	KylinNodeNamespaceListerExpansion
}

// kylinNodeNamespaceLister implements the KylinNodeNamespaceLister
// interface.
type kylinNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KylinNodes in the indexer for a given namespace.
func (s kylinNodeNamespaceLister) List(selector labels.Selector) (ret []*v1.KylinNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KylinNode))
	})
	return ret, err
}

// Get retrieves the KylinNode from the indexer for a given namespace and name.
func (s kylinNodeNamespaceLister) Get(name string) (*v1.KylinNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kylinnode"), name)
	}
	return obj.(*v1.KylinNode), nil
}
