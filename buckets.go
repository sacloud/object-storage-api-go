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

package objectstorage

import (
	"context"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// BucketAPI バケット操作関連API
type BucketAPI interface {
	// Create バケットの作成
	Create(ctx context.Context, siteId, bucketName string) (*v1.Bucket, error)
	// Delete バケットの削除
	Delete(ctx context.Context, siteId, bucketName string) error
}

var _ BucketAPI = (*bucketOp)(nil)

type bucketOp struct {
	client *Client
}

// NewBucketOp バケット操作関連API
func NewBucketOp(client *Client) BucketAPI {
	return &bucketOp{client: client}
}

func (op *bucketOp) Create(ctx context.Context, siteId, bucketName string) (*v1.Bucket, error) {
	params := v1.CreateBucketJSONRequestBody{
		ClusterId: siteId,
	}

	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.CreateBucketWithResponse(ctx, v1.BucketName(bucketName), params)
	if err != nil {
		return nil, err
	}
	bucket, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &bucket.Data, nil
}

func (op *bucketOp) Delete(ctx context.Context, siteId, bucketName string) error {
	params := v1.DeleteBucketJSONRequestBody{
		ClusterId: siteId,
	}

	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeleteBucketWithResponse(ctx, v1.BucketName(bucketName), params)
	if err != nil {
		return err
	}
	return resp.Result()
}
