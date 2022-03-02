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

package fake

import (
	"fmt"
	"time"

	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// DeleteSiteAccount サイトアカウントの削除
// (DELETE /{site_name}/v2/account)
func (engine *Engine) DeleteSiteAccount(siteId string) error {
	defer engine.lock()()

	// Note: API定義上は定義されていないがサイトがないケースでは404が返される
	if cluster := engine.getClusterById(siteId); cluster == nil {
		return NewError(ErrorTypeNotFound, "account", "",
			"指定のサイトは存在しません。cluster: %s", siteId)
	}

	if len(engine.Buckets) > 0 {
		return NewError(ErrorTypeConflict, "account", "",
			"アカウントに紐づくバケットが存在します。cluster: %s", siteId)
	}

	engine.Account = nil
	return nil
}

// ReadSiteAccount サイトアカウントの取得
// (GET /{site_name}/v2/account)
func (engine *Engine) ReadSiteAccount(siteId string) (*v1.Account, error) {
	defer engine.rLock()()

	if err := engine.siteAndAccountExist(siteId); err != nil {
		return nil, err
	}
	return engine.copyAccount(engine.Account)
}

// CreateSiteAccount サイトアカウントの作成
// (POST /{site_name}/v2/account)
func (engine *Engine) CreateSiteAccount(siteId string) (*v1.Account, error) {
	defer engine.lock()()

	// Note: API定義上は定義されていないがサイトがないケースでは404が返される
	cluster := engine.getClusterById(siteId)
	if cluster == nil {
		return nil, NewError(ErrorTypeNotFound, "account", "",
			"指定のサイトは存在しません。cluster: %s", siteId)
	}

	if engine.Account != nil {
		return nil, NewError(ErrorTypeConflict, "account", "",
			"すでにサイトにアカウントが存在します。cluster: %s", siteId)
	}

	engine.Account = &v1.Account{
		Code:       v1.Code("member@account@" + siteId),
		CreatedAt:  v1.CreatedAt(time.Now()),
		ResourceId: v1.ResourceID(fmt.Sprintf("%d", engine.nextId())),
	}

	return engine.copyAccount(engine.Account)
}

func (engine *Engine) copyAccount(source *v1.Account) (*v1.Account, error) {
	if source == nil {
		return nil, fmt.Errorf("source is nil")
	}
	var account v1.Account
	if err := deepcopy.Copy(&account, source); err != nil {
		return nil, err
	}
	return &account, nil
}

func (engine *Engine) siteAndAccountExist(siteId string) error {
	if err := engine.siteExist(siteId); err != nil {
		return err
	}

	if engine.Account == nil {
		return NewError(ErrorTypeNotFound, "account", "",
			"サイトにアカウントが存在しません。cluster: %s", siteId)
	}

	return nil
}
