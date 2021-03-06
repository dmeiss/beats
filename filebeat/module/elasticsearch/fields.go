// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
