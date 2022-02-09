// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded zlib format compressed contents of module/gcp.
func AssetGcp() string {
	return "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"
}