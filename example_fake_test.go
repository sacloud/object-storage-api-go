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

package objectstorage_test

import (
	"net/http/httptest"
	"time"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
)

func init() {
	initFakeServer()
}

func initFakeServer() {
	siteId := "isk01"
	fakeServer := &server.Server{
		Engine: &fake.Engine{
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
				Code:       "member@account@isk01",
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
		},
	}
	sv := httptest.NewServer(fakeServer.Handler())
	serverURL = sv.URL
}
