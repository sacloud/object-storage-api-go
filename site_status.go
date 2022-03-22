// Copyright 2022 The sacloud/object-storage-api-go authors
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

// SiteStatusAPI サイトステータスAPI
type SiteStatusAPI interface {
	// Read サイトステータスの参照
	Read(ctx context.Context, siteId string) (*v1.Status, error)
}

var _ SiteStatusAPI = (*siteStatusOp)(nil)

type siteStatusOp struct {
	client *Client
}

// NewSiteStatusOp サイトステータスAPI
func NewSiteStatusOp(client *Client) SiteStatusAPI {
	return &siteStatusOp{client: client}
}

func (op *siteStatusOp) Read(ctx context.Context, siteId string) (*v1.Status, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetStatusWithResponse(ctx, siteId)
	if err != nil {
		return nil, err
	}
	status, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &status.Data, nil
}
