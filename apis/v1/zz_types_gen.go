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

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package v1

import (
	"time"
)

// Access key ID
type AccessKeyID string

// Account info
type Account struct {
	// Code
	Code Code `json:"code"`

	// Created at
	CreatedAt CreatedAt `json:"created_at"`

	// Resource ID
	ResourceId ResourceID `json:"resource_id"`
}

// Root user access key
type AccountKey struct {
	// Created at
	CreatedAt CreatedAt `json:"created_at"`

	// Access key ID
	Id AccessKeyID `json:"id"`

	// Secret Access key
	Secret SecretAccessKey `json:"secret"`
}

// Root user access key
type AccountKeyResponseBody struct {
	// Root user access key
	Data AccountKey `json:"data"`
}

// Root user access keys
type AccountKeysResponseBody struct {
	// data type
	Data []AccountKey `json:"data"`
}

// Account info
type AccountResponseBody struct {
	// Account info
	Data Account `json:"data"`
}

// Bucket defines model for Bucket.
type Bucket struct {
	ClusterId string `json:"cluster_id"`
	Name      string `json:"name"`
}

// Bucket control
type BucketControl struct {
	// Bucket name
	BucketName BucketName `json:"bucket_name"`

	// The flag to read bucket contents
	CanRead CanRead `json:"can_read"`

	// The flag to write bucket contents
	CanWrite CanWrite `json:"can_write"`

	// Created at
	CreatedAt CreatedAt `json:"created_at"`
}

// Bucket controls
type BucketControls []BucketControl

// Bucket name
type BucketName string

// The flag to read bucket contents
type CanRead bool

// The flag to write bucket contents
type CanWrite bool

// Cluster defines model for Cluster.
type Cluster struct {
	// API Servers Zones
	ApiZone []string `json:"api_zone"`

	// URL of Control Panel
	ControlPanelUrl string `json:"control_panel_url"`

	// Display Name (Depending on Accept-Language)
	DisplayName string `json:"display_name"`

	// Display Name (en-us)
	DisplayNameEnUs string `json:"display_name_en_us"`

	// Display Name (ja)
	DisplayNameJa string `json:"display_name_ja"`

	// Display Order (Can be ignored)
	DisplayOrder int `json:"display_order"`

	// Endpoint Base of Cluster
	EndpointBase string `json:"endpoint_base"`

	// URL of IAM-compat API
	IamEndpoint string `json:"iam_endpoint"`

	// URL of IAM-compat API (w/ resigning)
	IamEndpointForControlPanel string `json:"iam_endpoint_for_control_panel"`
	Id                         string `json:"id"`

	// URL of S3-compat API
	S3Endpoint string `json:"s3_endpoint"`

	// URL of S3-compat API (w/ resigning)
	S3EndpointForControlPanel string `json:"s3_endpoint_for_control_panel"`

	// Storage Servers Zones
	StorageZone []string `json:"storage_zone"`
}

// Code
type Code string

// CreateBucketRequestBody defines model for CreateBucketRequestBody.
type CreateBucketRequestBody struct {
	ClusterId string `json:"cluster_id"`
}

// CreateBucketResponseBody defines model for CreateBucketResponseBody.
type CreateBucketResponseBody struct {
	Data Bucket `json:"data"`
}

// Created at
type CreatedAt time.Time

// Display name
type DisplayName string

// Error defines model for Error.
type Error struct {
	// どのサービスで発生したエラーかを判別する。
	// マイクロサービス名に加えてクラスター名を含む文字列が入ることを想定している。
	Domain ErrorsDomain `json:"domain"`

	// エラー発生箇所。
	// どのリソースなのか（どのリソースを操作した時に発生したものなのか）、
	// どのパラメータなのかといった情報。
	Location ErrorsLocation `json:"location"`

	// エラーの発生箇所の種類。
	// HTTPヘッダなのかHTTPパラメータなのか、
	// S3バケットなのかといったlocationの種別情報。
	LocationType ErrorsLocationType `json:"location_type"`

	// エラー発生時のメッセージ内容。
	// このメッセージはエラーを発生させたアプリケーションのメッセージをそのまま含む場合がある。
	Message ErrorsMessage `json:"message"`

	// なぜそのエラーが発生したかがわかる情報。
	// エラーメッセージの原因やエラー解決のためのヒントも含む場合がある。
	Reason ErrorsReason `json:"reason"`
}

// Error400 defines model for Error400.
type Error400 struct {
	// error
	Detail ErrorDetail `json:"error"`
}

// Error401 defines model for Error401.
type Error401 struct {
	// error
	Detail ErrorDetail `json:"error"`
}

// Error403 defines model for Error403.
type Error403 struct {
	// error
	Detail ErrorDetail `json:"error"`
}

// Error404 defines model for Error404.
type Error404 struct {
	// error
	Detail ErrorDetail `json:"error"`
}

// Error409 defines model for Error409.
type Error409 struct {
	// error
	Detail ErrorDetail `json:"error"`
}

// エラーコード。
type ErrorCode int32

// ErrorDefault defines model for ErrorDefault.
type ErrorDefault struct {
	// error
	Detail ErrorDetail `json:"error"`
}

// error
type ErrorDetail struct {
	// エラーコード。
	Code ErrorCode `json:"code"`

	// 認証に関するエラーについて詳細なエラー内容を表示する。
	Errors Errors `json:"errors"`

	// エラー発生時のメッセージ内容。
	// このメッセージはエラーを発生させたアプリケーションのメッセージをそのまま含む場合がある。
	Message ErrorMessage `json:"message"`

	// X-Sakura-Internal-Serial-ID
	TraceId ErrorTraceId `json:"trace_id"`
}

// エラー発生時のメッセージ内容。
// このメッセージはエラーを発生させたアプリケーションのメッセージをそのまま含む場合がある。
type ErrorMessage string

// X-Sakura-Internal-Serial-ID
type ErrorTraceId string

// 認証に関するエラーについて詳細なエラー内容を表示する。
type Errors []Error

// どのサービスで発生したエラーかを判別する。
// マイクロサービス名に加えてクラスター名を含む文字列が入ることを想定している。
type ErrorsDomain string

// エラー発生箇所。
// どのリソースなのか（どのリソースを操作した時に発生したものなのか）、
// どのパラメータなのかといった情報。
type ErrorsLocation string

// エラーの発生箇所の種類。
// HTTPヘッダなのかHTTPパラメータなのか、
// S3バケットなのかといったlocationの種別情報。
type ErrorsLocationType string

// エラー発生時のメッセージ内容。
// このメッセージはエラーを発生させたアプリケーションのメッセージをそのまま含む場合がある。
type ErrorsMessage string

// なぜそのエラーが発生したかがわかる情報。
// エラーメッセージの原因やエラー解決のためのヒントも含む場合がある。
type ErrorsReason string

// ListClustersResponseBody defines model for ListClustersResponseBody.
type ListClustersResponseBody struct {
	// If use a pointer type, braek output
	Data []Cluster `json:"data"`
}

// Permission defines model for Permission.
type Permission struct {
	// Bucket controls
	BucketControls BucketControls `json:"bucket_controls"`

	// Created at
	CreatedAt CreatedAt `json:"created_at"`

	// Display name
	DisplayName DisplayName `json:"display_name"`

	// Permission ID
	Id PermissionID `json:"id"`
}

// Permission ID
type PermissionID int64

// Permission Key
type PermissionKey struct {
	// Created at
	CreatedAt CreatedAt `json:"created_at"`

	// Access key ID
	Id AccessKeyID `json:"id"`

	// Permission secret key
	Secret PermissionSecret `json:"secret"`
}

// data type
type PermissionKeyResponseBody struct {
	// Permission Key
	Data PermissionKey `json:"data"`
}

// Permission Keys
type PermissionKeys []PermissionKey

// data type
type PermissionKeysResponseBody struct {
	// Permission Keys
	Data PermissionKeys `json:"data"`
}

// Request body for bucket controls for Permission
type PermissionRequestBody struct {
	// Bucket controls
	BucketControls BucketControls `json:"bucket_controls"`

	// Display name
	DisplayName DisplayName `json:"display_name"`
}

// PermissionResponseBody defines model for PermissionResponseBody.
type PermissionResponseBody struct {
	Data Permission `json:"data"`
}

// Permission secret key
type PermissionSecret string

// Permissions
type Permissions []Permission

// PermissionsResponseBody defines model for PermissionsResponseBody.
type PermissionsResponseBody struct {
	// Permissions
	Data Permissions `json:"data"`
}

// ReadClusterResponseBody defines model for ReadClusterResponseBody.
type ReadClusterResponseBody struct {
	Data *Cluster `json:"data,omitempty"`
}

// Resource ID
type ResourceID string

// Secret Access key
type SecretAccessKey string

// data type
type Status struct {
	AcceptNew  bool       `json:"accept_new"`
	Message    string     `json:"message"`
	StartedAt  time.Time  `json:"started_at"`
	StatusCode StatusCode `json:"status_code"`
}

// StatusCode defines model for StatusCode.
type StatusCode struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

// Status
type StatusResponseBody struct {
	// data type
	Data Status `json:"data"`
}

// DeleteBucketJSONBody defines parameters for DeleteBucket.
type DeleteBucketJSONBody CreateBucketRequestBody

// CreateBucketJSONBody defines parameters for CreateBucket.
type CreateBucketJSONBody CreateBucketRequestBody

// CreatePermissionJSONBody defines parameters for CreatePermission.
type CreatePermissionJSONBody PermissionRequestBody

// UpdatePermissionJSONBody defines parameters for UpdatePermission.
type UpdatePermissionJSONBody PermissionRequestBody

// DeleteBucketJSONRequestBody defines body for DeleteBucket for application/json ContentType.
type DeleteBucketJSONRequestBody DeleteBucketJSONBody

// CreateBucketJSONRequestBody defines body for CreateBucket for application/json ContentType.
type CreateBucketJSONRequestBody CreateBucketJSONBody

// CreatePermissionJSONRequestBody defines body for CreatePermission for application/json ContentType.
type CreatePermissionJSONRequestBody CreatePermissionJSONBody

// UpdatePermissionJSONRequestBody defines body for UpdatePermission for application/json ContentType.
type UpdatePermissionJSONRequestBody UpdatePermissionJSONBody
