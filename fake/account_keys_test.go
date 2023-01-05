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
	"time"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestEngine_AccountKeys(t *testing.T) {
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
		Account: &v1.Account{
			Code:       v1.Code("member@account@" + siteId),
			CreatedAt:  v1.CreatedAt(time.Now()),
			ResourceId: "100000000001",
		},
	}

	var keyId v1.AccessKeyID
	t.Run("create key", func(t *testing.T) {
		keys, err := engine.ListAccountAccessKeys(siteId)
		require.NoError(t, err)
		require.Len(t, keys, 0)

		key, err := engine.CreateAccountAccessKey(siteId)
		require.NoError(t, err)
		require.NotNil(t, key)
		require.Equal(t, v1.SecretAccessKey("secret"), key.Secret)
		keyId = key.Id

		keys, err = engine.ListAccountAccessKeys(siteId)
		require.NoError(t, err)
		require.Len(t, keys, 1)
	})

	t.Run("read account", func(t *testing.T) {
		key, err := engine.ReadAccountAccessKey(siteId, keyId.String())
		require.NoError(t, err)
		require.NotNil(t, key)
		require.Empty(t, key.Secret) // 参照時は空のはず
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
