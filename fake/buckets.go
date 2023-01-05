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
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// DeleteBucket バケットの削除
// (DELETE /fed/v1/buckets/{name})
func (engine *Engine) DeleteBucket(siteId, name string) error {
	defer engine.lock()()

	if err := engine.siteExist(siteId); err != nil {
		return err
	}

	bucket := engine.getBucketByName(name)
	if bucket == nil {
		// Note: 実際は存在しないバケット名を指定した場合204 NoContentが返っているが、API定義には記載がない。
		// このためここではUnknownErrorとしておく
		return newError(
			ErrorTypeUnknown, "bucket", name,
			"バケットが存在しません: cluster: %s, bucket: %s", siteId, name)
	}

	engine.deleteBucketByName(name)
	return nil
}

// CreateBucket バケットの作成
// (PUT /fed/v1/buckets/{name})
func (engine *Engine) CreateBucket(siteId, name string) (*v1.Bucket, error) {
	defer engine.lock()()

	if err := engine.siteExist(siteId); err != nil {
		return nil, err
	}

	bucket := engine.getBucketByName(name)
	if bucket != nil {
		return nil, newError(
			ErrorTypeConflict, "bucket", name,
			"同名バケットがすでに存在します: cluster: %s, bucket: %s", siteId, name)
	}

	// Note: バケットのclusterIdはどこからくるのか不明。
	// このためこの実装ではサイト(cluster)の先頭を用いることにする。
	// もしサイトが1つもなければエラーとする。
	if len(engine.Clusters) == 0 {
		return nil, newError(
			ErrorTypeUnknown, "bucket", "",
			"バケットの属するサイトが存在しません")
	}

	bucket = &v1.Bucket{
		ClusterId: engine.Clusters[0].Id,
		Name:      name,
	}
	engine.Buckets = append(engine.Buckets, bucket)
	return bucket, nil
}

func (engine *Engine) getBucketByName(name string) *v1.Bucket {
	if name == "" {
		return nil
	}
	for _, b := range engine.Buckets {
		if b.Name == name {
			return b
		}
	}
	return nil
}

func (engine *Engine) deleteBucketByName(name string) {
	var deleted []*v1.Bucket
	for _, bucket := range engine.Buckets {
		if bucket.Name != name {
			deleted = append(deleted, bucket)
		}
	}
	engine.Buckets = deleted
}
