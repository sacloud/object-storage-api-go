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
	"time"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestEngine_Permissions(t *testing.T) {
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
		Buckets: []*v1.Bucket{
			{
				ClusterId: siteId,
				Name:      "bucket1",
			},
			{
				ClusterId: siteId,
				Name:      "bucket2",
			},
		},
	}

	var permissionId v1.PermissionID
	t.Run("create permission", func(t *testing.T) {
		permissions, err := engine.ListPermissions(siteId)
		require.NoError(t, err)
		require.Len(t, permissions, 0)

		bucketControl := v1.BucketControl{
			BucketName: "bucket1",
			CanRead:    true,
			CanWrite:   false,
		}
		permission, err := engine.CreatePermission(siteId, &v1.PermissionRequestBody{
			BucketControls: v1.BucketControls{bucketControl},
			DisplayName:    "foobar",
		})
		require.NoError(t, err)
		require.NotNil(t, permission)

		require.Len(t, permission.BucketControls, 1)
		require.Equal(t, bucketControl, v1.BucketControl{
			BucketName: permission.BucketControls[0].BucketName,
			CanRead:    permission.BucketControls[0].CanRead,
			CanWrite:   permission.BucketControls[0].CanWrite,
		})
		permissionId = permission.Id

		permissions, err = engine.ListPermissions(siteId)
		require.NoError(t, err)
		require.Len(t, permissions, 1)
	})

	t.Run("read permission", func(t *testing.T) {
		key, err := engine.ReadPermission(siteId, permissionId.Int64())
		require.NoError(t, err)
		require.NotNil(t, key)
	})

	t.Run("update permission", func(t *testing.T) {
		bucketControl := v1.BucketControl{
			BucketName: "bucket2",
			CanRead:    false,
			CanWrite:   true,
		}
		permission, err := engine.UpdatePermission(siteId, permissionId.Int64(), &v1.PermissionRequestBody{
			BucketControls: v1.BucketControls{bucketControl},
			DisplayName:    "foobar",
		})
		require.NoError(t, err)
		require.NotNil(t, permission)

		require.Equal(t, bucketControl, v1.BucketControl{
			BucketName: permission.BucketControls[0].BucketName,
			CanRead:    permission.BucketControls[0].CanRead,
			CanWrite:   permission.BucketControls[0].CanWrite,
		})
	})

	t.Run("delete permission", func(t *testing.T) {
		err := engine.DeletePermission(siteId, permissionId.Int64())
		require.NoError(t, err)
		require.Len(t, engine.Permissions, 0)
	})
}
