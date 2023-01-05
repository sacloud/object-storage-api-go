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

func TestEngine_Buckets(t *testing.T) {
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
		Buckets: []*v1.Bucket{
			{
				ClusterId: "isk01",
				Name:      "bucket1",
			},
			{
				ClusterId: "isk01",
				Name:      "bucket2",
			},
		},
	}

	t.Run("create bucket", func(t *testing.T) {
		bucket, err := engine.CreateBucket("isk01", "foobar")
		require.NoError(t, err)
		require.Equal(t, &v1.Bucket{ClusterId: "isk01", Name: "foobar"}, bucket)
		require.Len(t, engine.Buckets, 3)
	})

	t.Run("delete bucket", func(t *testing.T) {
		err := engine.DeleteBucket("isk01", "foobar")
		require.NoError(t, err)
		require.Len(t, engine.Buckets, 2)
		require.Equal(t, engine.Buckets[0].Name, "bucket1")
		require.Equal(t, engine.Buckets[1].Name, "bucket2")
	})
}
