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
func (engine *Engine) DeleteSiteAccount(siteName string) error {
	defer engine.lock()()

	// Note: API定義上は定義されていないがサイトがないケースでは404が返される
	if cluster := engine.getClusterById(siteName); cluster == nil {
		return NewError(ErrorTypeNotFound, "account", "",
			"指定のサイトは存在しません。site_name: %s", siteName)
	}

	if len(engine.Buckets) > 0 {
		return NewError(ErrorTypeConflict, "account", "",
			"アカウントに紐づくバケットが存在します")
	}

	engine.Account = nil
	return nil
}

// ReadSiteAccount サイトアカウントの取得
// (GET /{site_name}/v2/account)
func (engine *Engine) ReadSiteAccount(siteName string) (*v1.Account, error) {
	defer engine.rLock()()

	if err := engine.siteAndAccountExist(siteName); err != nil {
		return nil, err
	}
	return engine.copyAccount(engine.Account)
}

// CreateSiteAccount サイトアカウントの作成
// (POST /{site_name}/v2/account)
func (engine *Engine) CreateSiteAccount(siteName string) (*v1.Account, error) {
	defer engine.lock()()

	// Note: API定義上は定義されていないがサイトがないケースでは404が返される
	cluster := engine.getClusterById(siteName)
	if cluster == nil {
		return nil, NewError(ErrorTypeNotFound, "account", "",
			"指定のサイトは存在しません。site_name: %s", siteName)
	}

	if engine.Account != nil {
		return nil, NewError(ErrorTypeConflict, "account", "",
			"すでにサイトにアカウントが存在します。site_name: %s", siteName)
	}

	engine.Account = &v1.Account{
		Code:       v1.Code("member@account@" + siteName),
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

func (engine *Engine) siteAndAccountExist(siteName string) error {
	// Note: API定義上は定義されていないがサイトがないケースでは404が返される
	if cluster := engine.getClusterById(siteName); cluster == nil {
		return NewError(ErrorTypeNotFound, "account", "",
			"指定のサイトは存在しません。site_name: %s", siteName)
	}

	if engine.Account == nil {
		return NewError(ErrorTypeNotFound, "account", "",
			"サイトにアカウントが存在しません。site_name: %s", siteName)
	}

	return nil
}
