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
	"testing"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestEngine_Clusters(t *testing.T) {
	engine := &Engine{
		Clusters: []*v1.Cluster{
			{
				Id: "isk01",

				ControlPanelUrl: "https://secure.sakura.ad.jp/objectstorage/",
				DisplayNameEnUs: "Ishikari Site #1",
				DisplayNameJa:   "石狩第1サイト",
				DisplayName:     "石狩第1サイト",
				DisplayOrder:    1,
				EndpointBase:    "isk01.sakurastorage.jp",
			},
		},
	}

	t.Run("select all", func(t *testing.T) {
		clusters, err := engine.ListClusters()
		require.NoError(t, err)
		require.Len(t, clusters, len(engine.Clusters))
	})

	t.Run("select by id", func(t *testing.T) {
		cluster, err := engine.ReadCluster("isk01")
		require.NoError(t, err)

		require.EqualValues(t, engine.Clusters[0], cluster)
		// 参照は異なるはず
		require.True(t, engine.Clusters[0] != cluster)
	})
}
