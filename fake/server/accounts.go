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

// DeleteSiteAccount サイトアカウントの削除
// (DELETE /{site_name}/v2/account)
func (s *Server) DeleteAccount(w http.ResponseWriter, r *http.Request, siteId string) {
	if err := s.Engine.DeleteSiteAccount(siteId); err != nil {
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

// ReadSiteAccount サイトアカウントの取得
// (GET /{site_name}/v2/account)
func (s *Server) GetAccount(w http.ResponseWriter, r *http.Request, siteId string) {
	account, err := s.Engine.ReadSiteAccount(siteId)
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
	_ = json.NewEncoder(w).Encode(&v1.AccountResponseBody{Data: *account})
}

// CreateSiteAccount サイトアカウントの作成
// (POST /{site_name}/v2/account)
func (s *Server) CreateAccount(w http.ResponseWriter, r *http.Request, siteId string) {
	account, err := s.Engine.CreateSiteAccount(siteId)
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
	_ = json.NewEncoder(w).Encode(&v1.AccountResponseBody{Data: *account})
}
