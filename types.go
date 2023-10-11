// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bprocessor

type bEvent struct {
	Type           string         `json:"type"`
	TraceID        string         `json:"traceId"`
	EntityMetadata entityMetadata `json:"entityMetadata"`
	Payload        interface{}
}

type entityMetadata struct {
	Type entityMetadataType `json:"type"`
}

type entityMetadataType struct {
	Name      string    `json:"name"`
	Namespace namespace `json:"namespace"`
}

type namespace struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}
