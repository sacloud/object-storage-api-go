// Copyright 2022-2025 The sacloud/object-storage-api-go authors
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

package objectstorage

import (
	"context"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// SiteAPI サイト(クラスター)関連API
type SiteAPI interface {
	// List サイト一覧
	List(ctx context.Context) ([]*v1.Cluster, error)
	// Read サイト詳細
	Read(ctx context.Context, siteId string) (*v1.Cluster, error)
}

var _ SiteAPI = (*siteOp)(nil)

type siteOp struct {
	client *Client
}

// NewSiteOp .
func NewSiteOp(client *Client) SiteAPI {
	return &siteOp{client: client}
}

func (op *siteOp) List(ctx context.Context) ([]*v1.Cluster, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetClustersWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	clusters, err := resp.Result()
	if err != nil {
		return nil, err
	}

	var results []*v1.Cluster
	for i := range clusters.Data {
		results = append(results, &clusters.Data[i])
	}
	return results, nil
}

func (op *siteOp) Read(ctx context.Context, id string) (*v1.Cluster, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetClusterWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	cluster, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return cluster.Data, nil
}
