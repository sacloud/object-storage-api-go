// Copyright 2022-2023 The sacloud/object-storage-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import (
	"fmt"

	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// ListClusters サイト一覧の取得
// (GET /fed/v1/clusters)
func (engine *Engine) ListClusters() ([]v1.Cluster, error) {
	defer engine.rLock()()

	return engine.clusters(), nil
}

// ReadCluster サイトの取得
// (GET /fed/v1/clusters/{id})
func (engine *Engine) ReadCluster(id string) (*v1.Cluster, error) {
	defer engine.rLock()()

	c := engine.getClusterById(id)
	if c != nil {
		return engine.copyCluster(c)
	}
	return nil, newError(ErrorTypeNotFound, "cluster", id)
}

// clusters engine.Clustersを非ポインタ型にして返す
func (engine *Engine) clusters() []v1.Cluster {
	var clusters []v1.Cluster
	for _, c := range engine.Clusters {
		clusters = append(clusters, *c)
	}
	return clusters
}

func (engine *Engine) getClusterById(id string) *v1.Cluster {
	if id == "" {
		return nil
	}
	for _, c := range engine.Clusters {
		if c.Id == id {
			return c
		}
	}
	return nil
}

func (engine *Engine) copyCluster(source *v1.Cluster) (*v1.Cluster, error) {
	if source == nil {
		return nil, fmt.Errorf("source is nil")
	}
	var cluster v1.Cluster
	if err := deepcopy.Copy(&cluster, source); err != nil {
		return nil, err
	}
	return &cluster, nil
}

func (engine *Engine) siteExist(siteName string) error {
	if cluster := engine.getClusterById(siteName); cluster == nil {
		return newError(ErrorTypeNotFound, "cluster", "",
			"指定のサイトは存在しません。site_name: %s", siteName)
	}
	return nil
}
