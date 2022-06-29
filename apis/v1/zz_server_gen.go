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

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// バケットの削除
	// (DELETE /fed/v1/buckets/{bucket_name})
	DeleteBucket(c *gin.Context, bucketName BucketName)
	// バケットの作成
	// (PUT /fed/v1/buckets/{bucket_name})
	CreateBucket(c *gin.Context, bucketName BucketName)
	// サイト一覧の取得
	// (GET /fed/v1/clusters)
	GetClusters(c *gin.Context)
	// サイトの取得
	// (GET /fed/v1/clusters/{site_id})
	GetCluster(c *gin.Context, siteId string)
	// サイトアカウントの削除
	// (DELETE /{site_id}/v2/account)
	DeleteAccount(c *gin.Context, siteId string)
	// サイトアカウントの取得
	// (GET /{site_id}/v2/account)
	GetAccount(c *gin.Context, siteId string)
	// サイトアカウントの作成
	// (POST /{site_id}/v2/account)
	CreateAccount(c *gin.Context, siteId string)
	// サイトアカウントのアクセスキーの取得
	// (GET /{site_id}/v2/account/keys)
	GetAccountKeys(c *gin.Context, siteId string)
	// サイトアカウントのアクセスキーの発行
	// (POST /{site_id}/v2/account/keys)
	CreateAccountKey(c *gin.Context, siteId string)
	// サイトアカウントのアクセスキーの削除
	// (DELETE /{site_id}/v2/account/keys/{account_key_id})
	DeleteAccountKey(c *gin.Context, siteId string, accountKeyId AccessKeyID)
	// サイトアカウントのアクセスキーの取得
	// (GET /{site_id}/v2/account/keys/{account_key_id})
	GetAccountKey(c *gin.Context, siteId string, accountKeyId AccessKeyID)
	// パーミッション一覧の取得
	// (GET /{site_id}/v2/permissions)
	GetPermissions(c *gin.Context, siteId string)
	// パーミッションの作成
	// (POST /{site_id}/v2/permissions)
	CreatePermission(c *gin.Context, siteId string)
	// パーミッションの削除
	// (DELETE /{site_id}/v2/permissions/{permission_id})
	DeletePermission(c *gin.Context, siteId string, permissionId PermissionID)
	// パーミッションの取得
	// (GET /{site_id}/v2/permissions/{permission_id})
	GetPermission(c *gin.Context, siteId string, permissionId PermissionID)
	// パーミッションの更新
	// (PUT /{site_id}/v2/permissions/{permission_id})
	UpdatePermission(c *gin.Context, siteId string, permissionId PermissionID)
	// パーミッションが保有するアクセスキー一覧の取得
	// (GET /{site_id}/v2/permissions/{permission_id}/keys)
	GetPermissionKeys(c *gin.Context, siteId string, permissionId PermissionID)
	// パーミッションのアクセスキーの発行
	// (POST /{site_id}/v2/permissions/{permission_id}/keys)
	CreatePermissionKey(c *gin.Context, siteId string, permissionId PermissionID)
	// パーミッションが保有するアクセスキーの削除
	// (DELETE /{site_id}/v2/permissions/{permission_id}/keys/{permission_key_id})
	DeletePermissionKey(c *gin.Context, siteId string, permissionId PermissionID, permissionKeyId AccessKeyID)
	// パーミッションが保有するアクセスキーの取得
	// (GET /{site_id}/v2/permissions/{permission_id}/keys/{permission_key_id})
	GetPermissionKey(c *gin.Context, siteId string, permissionId PermissionID, permissionKeyId AccessKeyID)
	// サイトのステータスの取得
	// (GET /{site_id}/v2/status)
	GetStatus(c *gin.Context, siteId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// DeleteBucket operation middleware
func (siw *ServerInterfaceWrapper) DeleteBucket(c *gin.Context) {

	var err error

	// ------------- Path parameter "bucket_name" -------------
	var bucketName BucketName

	err = runtime.BindStyledParameter("simple", false, "bucket_name", c.Param("bucket_name"), &bucketName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter bucket_name: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteBucket(c, bucketName)
}

// CreateBucket operation middleware
func (siw *ServerInterfaceWrapper) CreateBucket(c *gin.Context) {

	var err error

	// ------------- Path parameter "bucket_name" -------------
	var bucketName BucketName

	err = runtime.BindStyledParameter("simple", false, "bucket_name", c.Param("bucket_name"), &bucketName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter bucket_name: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateBucket(c, bucketName)
}

// GetClusters operation middleware
func (siw *ServerInterfaceWrapper) GetClusters(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetClusters(c)
}

// GetCluster operation middleware
func (siw *ServerInterfaceWrapper) GetCluster(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetCluster(c, siteId)
}

// DeleteAccount operation middleware
func (siw *ServerInterfaceWrapper) DeleteAccount(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteAccount(c, siteId)
}

// GetAccount operation middleware
func (siw *ServerInterfaceWrapper) GetAccount(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAccount(c, siteId)
}

// CreateAccount operation middleware
func (siw *ServerInterfaceWrapper) CreateAccount(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateAccount(c, siteId)
}

// GetAccountKeys operation middleware
func (siw *ServerInterfaceWrapper) GetAccountKeys(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAccountKeys(c, siteId)
}

// CreateAccountKey operation middleware
func (siw *ServerInterfaceWrapper) CreateAccountKey(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateAccountKey(c, siteId)
}

// DeleteAccountKey operation middleware
func (siw *ServerInterfaceWrapper) DeleteAccountKey(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "account_key_id" -------------
	var accountKeyId AccessKeyID

	err = runtime.BindStyledParameter("simple", false, "account_key_id", c.Param("account_key_id"), &accountKeyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter account_key_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteAccountKey(c, siteId, accountKeyId)
}

// GetAccountKey operation middleware
func (siw *ServerInterfaceWrapper) GetAccountKey(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "account_key_id" -------------
	var accountKeyId AccessKeyID

	err = runtime.BindStyledParameter("simple", false, "account_key_id", c.Param("account_key_id"), &accountKeyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter account_key_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAccountKey(c, siteId, accountKeyId)
}

// GetPermissions operation middleware
func (siw *ServerInterfaceWrapper) GetPermissions(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissions(c, siteId)
}

// CreatePermission operation middleware
func (siw *ServerInterfaceWrapper) CreatePermission(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreatePermission(c, siteId)
}

// DeletePermission operation middleware
func (siw *ServerInterfaceWrapper) DeletePermission(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeletePermission(c, siteId, permissionId)
}

// GetPermission operation middleware
func (siw *ServerInterfaceWrapper) GetPermission(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermission(c, siteId, permissionId)
}

// UpdatePermission operation middleware
func (siw *ServerInterfaceWrapper) UpdatePermission(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdatePermission(c, siteId, permissionId)
}

// GetPermissionKeys operation middleware
func (siw *ServerInterfaceWrapper) GetPermissionKeys(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissionKeys(c, siteId, permissionId)
}

// CreatePermissionKey operation middleware
func (siw *ServerInterfaceWrapper) CreatePermissionKey(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreatePermissionKey(c, siteId, permissionId)
}

// DeletePermissionKey operation middleware
func (siw *ServerInterfaceWrapper) DeletePermissionKey(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_key_id" -------------
	var permissionKeyId AccessKeyID

	err = runtime.BindStyledParameter("simple", false, "permission_key_id", c.Param("permission_key_id"), &permissionKeyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_key_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeletePermissionKey(c, siteId, permissionId, permissionKeyId)
}

// GetPermissionKey operation middleware
func (siw *ServerInterfaceWrapper) GetPermissionKey(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_id" -------------
	var permissionId PermissionID

	err = runtime.BindStyledParameter("simple", false, "permission_id", c.Param("permission_id"), &permissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_id: %s", err)})
		return
	}

	// ------------- Path parameter "permission_key_id" -------------
	var permissionKeyId AccessKeyID

	err = runtime.BindStyledParameter("simple", false, "permission_key_id", c.Param("permission_key_id"), &permissionKeyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter permission_key_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissionKey(c, siteId, permissionId, permissionKeyId)
}

// GetStatus operation middleware
func (siw *ServerInterfaceWrapper) GetStatus(c *gin.Context) {

	var err error

	// ------------- Path parameter "site_id" -------------
	var siteId string

	err = runtime.BindStyledParameter("simple", false, "site_id", c.Param("site_id"), &siteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter site_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetStatus(c, siteId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.DELETE(options.BaseURL+"/fed/v1/buckets/:bucket_name", wrapper.DeleteBucket)

	router.PUT(options.BaseURL+"/fed/v1/buckets/:bucket_name", wrapper.CreateBucket)

	router.GET(options.BaseURL+"/fed/v1/clusters", wrapper.GetClusters)

	router.GET(options.BaseURL+"/fed/v1/clusters/:site_id", wrapper.GetCluster)

	router.DELETE(options.BaseURL+"/:site_id/v2/account", wrapper.DeleteAccount)

	router.GET(options.BaseURL+"/:site_id/v2/account", wrapper.GetAccount)

	router.POST(options.BaseURL+"/:site_id/v2/account", wrapper.CreateAccount)

	router.GET(options.BaseURL+"/:site_id/v2/account/keys", wrapper.GetAccountKeys)

	router.POST(options.BaseURL+"/:site_id/v2/account/keys", wrapper.CreateAccountKey)

	router.DELETE(options.BaseURL+"/:site_id/v2/account/keys/:account_key_id", wrapper.DeleteAccountKey)

	router.GET(options.BaseURL+"/:site_id/v2/account/keys/:account_key_id", wrapper.GetAccountKey)

	router.GET(options.BaseURL+"/:site_id/v2/permissions", wrapper.GetPermissions)

	router.POST(options.BaseURL+"/:site_id/v2/permissions", wrapper.CreatePermission)

	router.DELETE(options.BaseURL+"/:site_id/v2/permissions/:permission_id", wrapper.DeletePermission)

	router.GET(options.BaseURL+"/:site_id/v2/permissions/:permission_id", wrapper.GetPermission)

	router.PUT(options.BaseURL+"/:site_id/v2/permissions/:permission_id", wrapper.UpdatePermission)

	router.GET(options.BaseURL+"/:site_id/v2/permissions/:permission_id/keys", wrapper.GetPermissionKeys)

	router.POST(options.BaseURL+"/:site_id/v2/permissions/:permission_id/keys", wrapper.CreatePermissionKey)

	router.DELETE(options.BaseURL+"/:site_id/v2/permissions/:permission_id/keys/:permission_key_id", wrapper.DeletePermissionKey)

	router.GET(options.BaseURL+"/:site_id/v2/permissions/:permission_id/keys/:permission_key_id", wrapper.GetPermissionKey)

	router.GET(options.BaseURL+"/:site_id/v2/status", wrapper.GetStatus)

	return router
}
