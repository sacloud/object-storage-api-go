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

package server

import (
	"encoding/json"
	"net/http"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
)

// ListPermissionAccessKeys パーミッションが保有するアクセスキー一覧の取得
// (GET /{site_name}/v2/permissions/{id}/keys)
func (s *Server) GetPermissionKeys(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID) {
	keys, err := s.Engine.ListPermissionAccessKeys(siteId, permissionId.Int64())
	if err != nil {
		switch e := err.(type) {
		case *fake.Error:
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&v1.PermissionKeysResponseBody{Data: keys})
}

// CreatePermissionAccessKey パーミッションのアクセスキーの発行
// (POST /{site_name}/v2/permissions/{id}/keys)
func (s *Server) CreatePermissionKey(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID) {
	key, err := s.Engine.CreatePermissionAccessKey(siteId, permissionId.Int64())
	if err != nil {
		switch e := err.(type) {
		case *fake.Error:
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&v1.PermissionKeyResponseBody{Data: *key})
}

// DeletePermissionAccessKey パーミッションが保有するアクセスキーの削除
// (DELETE /{site_name}/v2/permissions/{id}/keys/{key_id})
func (s *Server) DeletePermissionKey(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID, permissionKeyId v1.AccessKeyID) {
	if err := s.Engine.DeletePermissionAccessKey(siteId, permissionId.Int64(), permissionKeyId.String()); err != nil {
		switch e := err.(type) {
		case *fake.Error:
			if e.Type == fake.ErrorTypeNotFound {
				http.Error(w, e.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// ReadPermissionAccessKey パーミッションが保有するアクセスキーの取得
// (GET /{site_name}/v2/permissions/{id}/keys/{key_id})
func (s *Server) GetPermissionKey(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID, permissionKeyId v1.AccessKeyID) {
	key, err := s.Engine.ReadPermissionAccessKey(siteId, permissionId.Int64(), permissionKeyId.String())
	if err != nil {
		switch e := err.(type) {
		case *fake.Error:
			if e.Type == fake.ErrorTypeNotFound {
				http.Error(w, e.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&v1.PermissionKeyResponseBody{Data: *key})
}
