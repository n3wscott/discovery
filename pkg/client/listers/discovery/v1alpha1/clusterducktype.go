/*
Copyright 2021 The Knative Authors

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "knative.dev/discovery/pkg/apis/discovery/v1alpha1"
)

// ClusterDuckTypeLister helps list ClusterDuckTypes.
// All objects returned here must be treated as read-only.
type ClusterDuckTypeLister interface {
	// List lists all ClusterDuckTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterDuckType, err error)
	// Get retrieves the ClusterDuckType from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterDuckType, error)
	ClusterDuckTypeListerExpansion
}

// clusterDuckTypeLister implements the ClusterDuckTypeLister interface.
type clusterDuckTypeLister struct {
	indexer cache.Indexer
}

// NewClusterDuckTypeLister returns a new ClusterDuckTypeLister.
func NewClusterDuckTypeLister(indexer cache.Indexer) ClusterDuckTypeLister {
	return &clusterDuckTypeLister{indexer: indexer}
}

// List lists all ClusterDuckTypes in the indexer.
func (s *clusterDuckTypeLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterDuckType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterDuckType))
	})
	return ret, err
}

// Get retrieves the ClusterDuckType from the index for a given name.
func (s *clusterDuckTypeLister) Get(name string) (*v1alpha1.ClusterDuckType, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterducktype"), name)
	}
	return obj.(*v1alpha1.ClusterDuckType), nil
}
