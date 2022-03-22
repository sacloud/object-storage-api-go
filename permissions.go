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

// PermissionAPI パーミッション関連API
type PermissionAPI interface {
	List(ctx context.Context, siteId string) ([]*v1.Permission, error)
	// Create パーミッションの作成
	Create(ctx context.Context, siteId string, params *v1.CreatePermissionParams) (*v1.Permission, error)
	// Read パーミッションの参照
	Read(ctx context.Context, siteId string, permissionId int64) (*v1.Permission, error)
	// Update パーミッションの更新
	Update(ctx context.Context, siteId string, permissionId int64, params *v1.UpdatePermissionParams) (*v1.Permission, error)
	// Delete パーミッションの削除
	Delete(ctx context.Context, siteId string, permissionId int64) error

	// ListAccessKeys アクセスキー一覧
	ListAccessKeys(ctx context.Context, siteId string, permissionId int64) ([]*v1.PermissionKey, error)
	// CreateAccessKey アクセスキーの作成
	//
	// Secretはこの戻り値でのみ参照可能
	CreateAccessKey(ctx context.Context, siteId string, permissionId int64) (*v1.PermissionKey, error)
	// ReadAccessKey アクセスキーの参照
	//
	// Secretは常に空文字になっている
	ReadAccessKey(ctx context.Context, siteId string, permissionId int64, accessKeyId string) (*v1.PermissionKey, error)
	// DeleteAccessKey アクセスキーの削除
	DeleteAccessKey(ctx context.Context, siteId string, permissionId int64, accessKeyId string) error
}

var _ PermissionAPI = (*permissionOp)(nil)

type permissionOp struct {
	client *Client
}

// NewPermissionOp パーミッション関連API
func NewPermissionOp(client *Client) PermissionAPI {
	return &permissionOp{client: client}
}

func (op *permissionOp) List(ctx context.Context, siteId string) ([]*v1.Permission, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetPermissionsWithResponse(ctx, siteId)
	if err != nil {
		return nil, err
	}
	permissions, err := resp.Result()
	if err != nil {
		return nil, err
	}
	var results []*v1.Permission
	for _, p := range permissions.Data {
		results = append(results, &p)
	}
	return results, nil
}

func (op *permissionOp) Create(ctx context.Context, siteId string, params *v1.CreatePermissionParams) (*v1.Permission, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.CreatePermissionWithResponse(ctx, siteId, v1.CreatePermissionJSONRequestBody(*params))
	if err != nil {
		return nil, err
	}
	permission, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &permission.Data, nil
}

func (op *permissionOp) Read(ctx context.Context, siteId string, permissionId int64) (*v1.Permission, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetPermissionWithResponse(ctx, siteId, v1.PermissionID(permissionId))
	if err != nil {
		return nil, err
	}
	permission, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &permission.Data, nil
}

func (op *permissionOp) Update(ctx context.Context, siteId string, permissionId int64, params *v1.UpdatePermissionParams) (*v1.Permission, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.UpdatePermissionWithResponse(ctx, siteId,
		v1.PermissionID(permissionId), v1.UpdatePermissionJSONRequestBody(*params))
	if err != nil {
		return nil, err
	}
	permission, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &permission.Data, nil
}

func (op *permissionOp) Delete(ctx context.Context, siteId string, permissionId int64) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeletePermissionWithResponse(ctx, siteId, v1.PermissionID(permissionId))
	if err != nil {
		return err
	}
	return resp.Result()
}

func (op *permissionOp) ListAccessKeys(ctx context.Context, siteId string, permissionId int64) ([]*v1.PermissionKey, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetPermissionKeysWithResponse(ctx, siteId, v1.PermissionID(permissionId))
	if err != nil {
		return nil, err
	}
	permissions, err := resp.Result()
	if err != nil {
		return nil, err
	}
	var results []*v1.PermissionKey
	for _, p := range permissions.Data {
		results = append(results, &p)
	}
	return results, nil
}

func (op *permissionOp) CreateAccessKey(ctx context.Context, siteId string, permissionId int64) (*v1.PermissionKey, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.CreatePermissionKeyWithResponse(ctx, siteId, v1.PermissionID(permissionId))
	if err != nil {
		return nil, err
	}
	key, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &key.Data, nil
}

func (op *permissionOp) ReadAccessKey(ctx context.Context, siteId string, permissionId int64, accessKeyId string) (*v1.PermissionKey, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetPermissionKeyWithResponse(ctx, siteId,
		v1.PermissionID(permissionId), v1.AccessKeyID(accessKeyId))
	if err != nil {
		return nil, err
	}
	key, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &key.Data, nil
}

func (op *permissionOp) DeleteAccessKey(ctx context.Context, siteId string, permissionId int64, accessKeyId string) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeletePermissionKeyWithResponse(ctx, siteId,
		v1.PermissionID(permissionId), v1.AccessKeyID(accessKeyId))
	if err != nil {
		return err
	}
	return resp.Result()
}
