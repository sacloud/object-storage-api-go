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

// AccountAPI アカウント操作関連API
type AccountAPI interface {
	// Create アカウントの作成
	Create(ctx context.Context, siteId string) (*v1.Account, error)
	// Read アカウントの参照
	Read(ctx context.Context, siteId string) (*v1.Account, error)
	// Delete アカウントの削除
	Delete(ctx context.Context, siteId string) error

	// ListAccessKeys アクセスキーの参照
	//
	// Secretは常に空文字になっている
	ListAccessKeys(ctx context.Context, siteId string) ([]*v1.AccountKey, error)

	// CreateAccessKey アクセスキーの作成
	//
	// Secretはこの戻り値でのみ参照可能
	CreateAccessKey(ctx context.Context, siteId string) (*v1.AccountKey, error)
	// ReadAccessKey アクセスキーの参照
	//
	// Secretは常に空文字になっている
	ReadAccessKey(ctx context.Context, siteId, accessKeyId string) (*v1.AccountKey, error)
	// DeleteAccessKey アクセスキーの削除
	DeleteAccessKey(ctx context.Context, siteId, accessKeyId string) error
}

var _ AccountAPI = (*accountOp)(nil)

type accountOp struct {
	client *Client
}

// NewAccountOp アカウント操作関連API
func NewAccountOp(client *Client) AccountAPI {
	return &accountOp{client: client}
}

func (op *accountOp) Create(ctx context.Context, siteId string) (*v1.Account, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.CreateAccountWithResponse(ctx, siteId)
	if err != nil {
		return nil, err
	}
	account, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &account.Data, nil
}

func (op *accountOp) Read(ctx context.Context, siteId string) (*v1.Account, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetAccountWithResponse(ctx, siteId)
	if err != nil {
		return nil, err
	}
	account, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &account.Data, nil
}

func (op *accountOp) Delete(ctx context.Context, siteId string) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeleteAccountWithResponse(ctx, siteId)
	if err != nil {
		return err
	}
	return resp.Result()
}

func (op *accountOp) ListAccessKeys(ctx context.Context, siteId string) ([]*v1.AccountKey, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetAccountKeysWithResponse(ctx, siteId)
	if err != nil {
		return nil, err
	}
	keys, err := resp.Result()
	if err != nil {
		return nil, err
	}
	var results []*v1.AccountKey
	for i := range keys.Data {
		results = append(results, &keys.Data[i])
	}
	return results, nil
}

func (op *accountOp) CreateAccessKey(ctx context.Context, siteId string) (*v1.AccountKey, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.CreateAccountKeyWithResponse(ctx, siteId)
	if err != nil {
		return nil, err
	}
	account, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &account.Data, nil
}

func (op *accountOp) ReadAccessKey(ctx context.Context, siteId, accessKeyId string) (*v1.AccountKey, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetAccountKeyWithResponse(ctx, siteId, v1.AccessKeyID(accessKeyId))
	if err != nil {
		return nil, err
	}
	account, err := resp.Result()
	if err != nil {
		return nil, err
	}
	return &account.Data, nil
}

func (op *accountOp) DeleteAccessKey(ctx context.Context, siteId, accessKeyId string) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeleteAccountKeyWithResponse(ctx, siteId, v1.AccessKeyID(accessKeyId))
	if err != nil {
		return err
	}
	return resp.Result()
}
