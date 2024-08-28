// Copyright (c) 2024 Alibaba Group Holding Ltd.
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

package db

type DbClientCommonAttrsGetter[REQUEST any] interface {
	GetSystem(REQUEST) string
	GetUser(REQUEST) string
	GetName(REQUEST) string
	GetConnectionString(REQUEST) string
}

type DbClientAttrsGetter[REQUEST any] interface {
	DbClientCommonAttrsGetter[REQUEST]
	GetStatement(REQUEST) string
	GetOperation(REQUEST) string
}

type SqlClientAttributesGetter[REQUEST any] interface {
	DbClientCommonAttrsGetter[REQUEST]
	GetRawStatement(REQUEST) string
}
