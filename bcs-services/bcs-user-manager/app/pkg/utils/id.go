/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package utils xxx
package utils

import "context"

// StringInSlice returns true if given string in slice
func StringInSlice(s string, l []string) bool {
	for _, objStr := range l {
		if s == objStr {
			return true
		}
	}
	return false
}

// ContextValueKey define context value key
type ContextValueKey string

const (
	// ContextValueKeyRequestID request_id in context value
	ContextValueKeyRequestID ContextValueKey = "X-Request-Id"
)

// GetRequestIDFromContext get request id from context
func GetRequestIDFromContext(ctx context.Context) string {
	requestID, _ := ctx.Value(ContextValueKeyRequestID).(string)
	return requestID
}
