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

//go:build acctest
// +build acctest

package objectstorage_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/packages-go/envvar"
	"github.com/sacloud/packages-go/testutil"
	"github.com/stretchr/testify/require"
)

const (
	siteId = "isk01" // Note: 本来はサイトAPI経由で取得すべきだが現状ではサイトが1つしかないため定数として保持/利用する
)

func s3Client(t *testing.T, token, secret string) *minio.Client {
	t.Helper()

	endpoint := "s3.isk01.sakurastorage.jp"
	s3Client, err := minio.New(endpoint, &minio.Options{
		Creds:        credentials.NewStaticV4(token, secret, ""),
		Region:       "jp-north-1",
		Secure:       true,
		BucketLookup: minio.BucketLookupPath,
	})
	if err != nil {
		t.Fatal(err)
	}
	return s3Client
}

func s3ClientFromEnv(t *testing.T) *minio.Client {
	return s3Client(t, os.Getenv("SACLOUD_OJS_ACCESS_KEY_ID"), os.Getenv("SACLOUD_OJS_SECRET_ACCESS_KEY"))
}

// TestAccSiteAndStatusAPI サイト/ステータス関連APIの疎通確認
func TestAccSiteAndStatusAPI(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx := context.Background()
	siteOp := objectstorage.NewSiteOp(accTestClient)

	// List
	sites, err := siteOp.List(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, sites)

	// Read
	site, err := siteOp.Read(ctx, sites[0].Id)
	require.NoError(t, err)
	require.NotEmpty(t, site)

	// Site Status
	statusOp := objectstorage.NewSiteStatusOp(accTestClient)
	status, err := statusOp.Read(ctx, site.Id)

	require.NoError(t, err)
	require.NotEmpty(t, status)
}

// TestAccBucketHandling バケット周りの一連の操作のテスト
//
// 一連の操作は以下のステップで行う
//   - Step1: バケットを作成
//   - Step2: アクセスキーを用いてAWS SDK経由でバケットが作成されたかを確認
//   - Step3: 各リソースのクリーンアップ
//
// Note: バケット一覧の参照のためにサイトアカウントのアクセスキーが必要
func TestAccBucketHandling(t *testing.T) {
	skipIfNoAPIKey(t)
	skipIfNoEnv(t, "SACLOUD_OJS_ACCESS_KEY_ID", "SACLOUD_OJS_SECRET_ACCESS_KEY")

	ctx := context.Background()
	bucketName := testutil.Random(28, testutil.CharSetAlpha)

	// Step1: バケット作成
	bucketOp := objectstorage.NewBucketOp(accTestClient)
	{
		created, err := bucketOp.Create(ctx, siteId, bucketName)
		require.NoError(t, err)
		require.NotEmpty(t, created)
	}

	// Step2: バケットにアクセス
	{
		s3Client := s3ClientFromEnv(t)

		output, err := s3Client.ListBuckets(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, output)

		exist := false
		for _, b := range output {
			if b.Name == bucketName {
				exist = true
				break
			}
		}
		require.True(t, exist, "bucket %q is not exist", bucketName)
	}

	// Step3: クリーンアップ
	require.NoError(t, bucketOp.Delete(ctx, siteId, bucketName))
}

// TestAccAccessToBucketWithPermissionKey パーミッションキーによるオブジェクトへのアクセス
func TestAccAccessToBucketObjectWithPermissionKey(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx := context.Background()
	bucketName := testutil.Random(28, testutil.CharSetAlpha)

	// Step1: バケット作成
	bucketOp := objectstorage.NewBucketOp(accTestClient)
	{
		created, err := bucketOp.Create(ctx, siteId, bucketName)
		require.NoError(t, err)
		require.NotEmpty(t, created)
	}

	// Step2: バケットにアクセスできるパーミッション/アクセスキーの作成
	permissionOp := objectstorage.NewPermissionOp(accTestClient)
	var permission *v1.Permission
	var key *v1.PermissionKey
	{
		created, err := permissionOp.Create(ctx, siteId, &v1.CreatePermissionParams{
			BucketControls: v1.BucketControls{
				{
					BucketName: v1.BucketName(bucketName),
					CanRead:    true,
					CanWrite:   true,
				},
			},
			DisplayName: v1.DisplayName(bucketName),
		})
		require.NoError(t, err)
		require.NotEmpty(t, created)
		permission = created

		createdKey, err := permissionOp.CreateAccessKey(ctx, siteId, permission.Id.Int64())
		require.NoError(t, err)
		require.NotEmpty(t, createdKey)
		key = createdKey
	}

	// Step3: 作成したアクセスキーでバケットにアクセス
	{
		s3Client := s3Client(t, key.Id.String(), key.Secret.String())

		objectKey := "foobar"
		objectBodyText := "body of s3://[bucket_name]/foobar"
		objectBody := bytes.NewBufferString(objectBodyText)

		// オブジェクトの作成
		created, err := s3Client.PutObject(ctx, bucketName, objectKey, objectBody, int64(objectBody.Len()), minio.PutObjectOptions{})
		require.NoError(t, err)
		require.NotEmpty(t, created)

		// オブジェクトの読み込み
		read, err := s3Client.GetObject(ctx, bucketName, objectKey, minio.GetObjectOptions{})
		require.NoError(t, err)
		require.NotEmpty(t, read)

		readText, err := io.ReadAll(read)
		require.NoError(t, err)
		require.Equal(t, objectBodyText, string(readText))

		// オブジェクトの削除
		err = s3Client.RemoveObject(ctx, bucketName, objectKey, minio.RemoveObjectOptions{})
		require.NoError(t, err)
	}

	// Step4: クリーンアップ
	require.NoError(t, permissionOp.DeleteAccessKey(ctx, siteId, permission.Id.Int64(), key.Id.String()))
	require.NoError(t, permissionOp.Delete(ctx, siteId, permission.Id.Int64()))
	require.NoError(t, bucketOp.Delete(ctx, siteId, bucketName))
}

var accTestClient = &objectstorage.Client{
	APIRootURL: envvar.StringFromEnv("SAKURACLOUD_OJS_ROOT_URL", defaultServerURL),
}

// skipIfNoEnv 指定の環境変数のいずれかが空の場合はt.SkipNow()する
func skipIfNoEnv(t *testing.T, envs ...string) {
	var emptyEnvs []string
	for _, env := range envs {
		if os.Getenv(env) == "" {
			emptyEnvs = append(emptyEnvs, env)
		}
	}
	if len(emptyEnvs) > 0 {
		for _, env := range emptyEnvs {
			t.Logf("environment variable %q is not set", env)
		}
		t.SkipNow()
	}
}

func skipIfNoAPIKey(t *testing.T) {
	skipIfNoEnv(t, "SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")
}
