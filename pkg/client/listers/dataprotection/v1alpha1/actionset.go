/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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
	v1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ActionSetLister helps list ActionSets.
// All objects returned here must be treated as read-only.
type ActionSetLister interface {
	// List lists all ActionSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActionSet, err error)
	// Get retrieves the ActionSet from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ActionSet, error)
	ActionSetListerExpansion
}

// actionSetLister implements the ActionSetLister interface.
type actionSetLister struct {
	indexer cache.Indexer
}

// NewActionSetLister returns a new ActionSetLister.
func NewActionSetLister(indexer cache.Indexer) ActionSetLister {
	return &actionSetLister{indexer: indexer}
}

// List lists all ActionSets in the indexer.
func (s *actionSetLister) List(selector labels.Selector) (ret []*v1alpha1.ActionSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionSet))
	})
	return ret, err
}

// Get retrieves the ActionSet from the index for a given name.
func (s *actionSetLister) Get(name string) (*v1alpha1.ActionSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("actionset"), name)
	}
	return obj.(*v1alpha1.ActionSet), nil
}
