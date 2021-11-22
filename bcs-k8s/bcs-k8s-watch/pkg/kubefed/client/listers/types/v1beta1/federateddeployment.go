/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedDeploymentLister helps list FederatedDeployments.
type FederatedDeploymentLister interface {
	// List lists all FederatedDeployments in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedDeployment, err error)
	// FederatedDeployments returns an object that can list and get FederatedDeployments.
	FederatedDeployments(namespace string) FederatedDeploymentNamespaceLister
	FederatedDeploymentListerExpansion
}

// federatedDeploymentLister implements the FederatedDeploymentLister interface.
type federatedDeploymentLister struct {
	indexer cache.Indexer
}

// NewFederatedDeploymentLister returns a new FederatedDeploymentLister.
func NewFederatedDeploymentLister(indexer cache.Indexer) FederatedDeploymentLister {
	return &federatedDeploymentLister{indexer: indexer}
}

// List lists all FederatedDeployments in the indexer.
func (s *federatedDeploymentLister) List(selector labels.Selector) (ret []*v1beta1.FederatedDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedDeployment))
	})
	return ret, err
}

// FederatedDeployments returns an object that can list and get FederatedDeployments.
func (s *federatedDeploymentLister) FederatedDeployments(namespace string) FederatedDeploymentNamespaceLister {
	return federatedDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedDeploymentNamespaceLister helps list and get FederatedDeployments.
type FederatedDeploymentNamespaceLister interface {
	// List lists all FederatedDeployments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedDeployment, err error)
	// Get retrieves the FederatedDeployment from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedDeployment, error)
	FederatedDeploymentNamespaceListerExpansion
}

// federatedDeploymentNamespaceLister implements the FederatedDeploymentNamespaceLister
// interface.
type federatedDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedDeployments in the indexer for a given namespace.
func (s federatedDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedDeployment))
	})
	return ret, err
}

// Get retrieves the FederatedDeployment from the indexer for a given namespace and name.
func (s federatedDeploymentNamespaceLister) Get(name string) (*v1beta1.FederatedDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federateddeployment"), name)
	}
	return obj.(*v1beta1.FederatedDeployment), nil
}