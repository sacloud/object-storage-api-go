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

// ListAccountAccessKeys サイトアカウントのアクセスキーの取得
// (GET /{site_name}/v2/account/keys)
func (s *Server) GetAccountKeys(w http.ResponseWriter, r *http.Request, siteId string) {
	keys, err := s.Engine.ListAccountAccessKeys(siteId)
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
	_ = json.NewEncoder(w).Encode(&v1.AccountKeysResponseBody{Data: keys})
}

// CreateAccountAccessKey サイトアカウントのアクセスキーの発行
// (POST /{site_name}/v2/account/keys)
func (s *Server) CreateAccountKey(w http.ResponseWriter, r *http.Request, siteId string) {
	key, err := s.Engine.CreateAccountAccessKey(siteId)
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
	_ = json.NewEncoder(w).Encode(&v1.AccountKeyResponseBody{Data: *key})
}

// DeleteAccountAccessKey サイトアカウントのアクセスキーの削除
// (DELETE /{site_name}/v2/account/keys/{id})
func (s *Server) DeleteAccountKey(w http.ResponseWriter, r *http.Request, siteId string, accountKeyId v1.AccessKeyID) {
	if err := s.Engine.DeleteAccountAccessKey(siteId, accountKeyId.String()); err != nil {
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

// ReadAccountAccessKey サイトアカウントのアクセスキーの取得
// (GET /{site_name}/v2/account/keys/{id})
func (s *Server) GetAccountKey(w http.ResponseWriter, r *http.Request, siteId string, accountKeyId v1.AccessKeyID) {
	key, err := s.Engine.ReadAccountAccessKey(siteId, accountKeyId.String())
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
	_ = json.NewEncoder(w).Encode(&v1.AccountKeyResponseBody{Data: *key})
}
