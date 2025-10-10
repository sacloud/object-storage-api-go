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

// ListPermissions パーミッション一覧の取得
// (GET /{site_name}/v2/permissions)
func (s *Server) GetPermissions(w http.ResponseWriter, r *http.Request, siteId string) {
	permissions, err := s.Engine.ListPermissions(siteId)
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
	_ = json.NewEncoder(w).Encode(&v1.PermissionsResponseBody{Data: permissions})
}

// CreatePermission パーミッションの作成
// (POST /{site_name}/v2/permissions)
func (s *Server) CreatePermission(w http.ResponseWriter, r *http.Request, siteId string) {
	var paramJSON v1.PermissionRequestBody
	if err := json.NewDecoder(r.Body).Decode(&paramJSON); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	permission, err := s.Engine.CreatePermission(siteId, &paramJSON)
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
	_ = json.NewEncoder(w).Encode(&v1.PermissionResponseBody{Data: *permission})
}

// DeletePermission パーミッションの削除
// (DELETE /{site_name}/v2/permissions/{id})
func (s *Server) DeletePermission(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID) {
	if err := s.Engine.DeletePermission(siteId, permissionId.Int64()); err != nil {
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

// GetPermission パーミッションの取得
// (GET /{site_name}/v2/permissions/{id})
func (s *Server) GetPermission(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID) {
	permission, err := s.Engine.ReadPermission(siteId, permissionId.Int64())
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
	_ = json.NewEncoder(w).Encode(&v1.PermissionResponseBody{Data: *permission})
}

// UpdatePermission パーミッションの更新
// (PUT /{site_name}/v2/permissions/{id})
func (s *Server) UpdatePermission(w http.ResponseWriter, r *http.Request, siteId string, permissionId v1.PermissionID) {
	var paramJSON v1.PermissionRequestBody
	if err := json.NewDecoder(r.Body).Decode(&paramJSON); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	permission, err := s.Engine.UpdatePermission(siteId, permissionId.Int64(), &paramJSON)
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
	_ = json.NewEncoder(w).Encode(&v1.PermissionResponseBody{Data: *permission})
}
