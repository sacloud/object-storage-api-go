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

// AccountAPI アカウント操作関連API
type AccountAPI interface {
	// Create アカウントの作成
	Create(ctx context.Context, siteId string) (*v1.Account, error)
	// Read アカウントの参照
	Read(ctx context.Context, siteId string) (*v1.Account, error)
	// Delete アカウントの削除
	Delete(ctx context.Context, siteId string) error

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
	resp, err := op.client.apiClient().CreateSiteAccountWithResponse(ctx, siteId)
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
	resp, err := op.client.apiClient().ReadSiteAccountWithResponse(ctx, siteId)
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
	resp, err := op.client.apiClient().DeleteSiteAccountWithResponse(ctx, siteId)
	if err != nil {
		return err
	}
	return resp.Result()
}

func (op *accountOp) CreateAccessKey(ctx context.Context, siteId string) (*v1.AccountKey, error) {
	resp, err := op.client.apiClient().CreateAccountAccessKeyWithResponse(ctx, siteId)
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
	resp, err := op.client.apiClient().ReadAccountAccessKeyWithResponse(ctx, siteId, v1.AccessKeyID(accessKeyId))
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
	resp, err := op.client.apiClient().DeleteAccountAccessKeyWithResponse(ctx, siteId, v1.AccessKeyID(accessKeyId))
	if err != nil {
		return err
	}
	return resp.Result()
}
