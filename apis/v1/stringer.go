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

package v1

import "fmt"

// String .
func (v *AccessKeyID) String() string {
	return string(*v)
}

// String .
func (v *BucketName) String() string {
	return string(*v)
}

// String .
func (v *Code) String() string {
	return string(*v)
}

// String .
func (v *DisplayName) String() string {
	return string(*v)
}

// String .
func (v *ErrorMessage) String() string {
	return string(*v)
}

// String .
func (v *ErrorTraceId) String() string {
	return string(*v)
}

// String .
func (v *ErrorsDomain) String() string {
	return string(*v)
}

// String .
func (v *ErrorsLocation) String() string {
	return string(*v)
}

// String .
func (v *ErrorsLocationType) String() string {
	return string(*v)
}

// String .
func (v *ErrorsMessage) String() string {
	return string(*v)
}

// String .
func (v *ErrorsReason) String() string {
	return string(*v)
}

// String .
func (v *PermissionSecret) String() string {
	return string(*v)
}

// String .
func (v *ResourceID) String() string {
	return string(*v)
}

// String .
func (v *SecretAccessKey) String() string {
	return string(*v)
}

// String .
func (v PermissionID) String() string {
	return fmt.Sprintf("%d", v)
}
