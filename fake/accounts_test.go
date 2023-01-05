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
	"testing"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestEngine_Accounts(t *testing.T) {
	siteId := "isk01"
	engine := &Engine{
		Clusters: []*v1.Cluster{
			{
				Id: siteId,

				ControlPanelUrl: "https://secure.sakura.ad.jp/objectstorage/",
				DisplayNameEnUs: "Ishikari Site #1",
				DisplayNameJa:   "石狩第1サイト",
				DisplayName:     "石狩第1サイト",
				DisplayOrder:    1,
				EndpointBase:    "isk01.sakurastorage.jp",
			},
		},
		Account: nil,
	}

	t.Run("create account", func(t *testing.T) {
		account, err := engine.CreateSiteAccount(siteId)
		require.NoError(t, err)
		require.NotNil(t, account)
		require.NotNil(t, engine.Account)
		require.Equal(t, v1.Code("member@account@"+siteId), account.Code)
	})

	t.Run("read account", func(t *testing.T) {
		account, err := engine.ReadSiteAccount(siteId)
		require.NoError(t, err)
		require.NotNil(t, account)
		require.Equal(t, v1.Code("member@account@"+siteId), account.Code)
	})

	t.Run("delete account with invalid siteId", func(t *testing.T) {
		err := engine.DeleteSiteAccount("invalid")
		require.Error(t, err)
	})

	t.Run("delete account", func(t *testing.T) {
		err := engine.DeleteSiteAccount(siteId)
		require.NoError(t, err)
		require.Nil(t, engine.Account)
	})
}
