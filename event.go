// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bprocessor

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type eventGenerator struct{}

// consume takes a single trace and generate b event
func (e *eventGenerator) generate(td ptrace.Traces) (*bEvent, error) {
	traceID, err := getTraceID(td)
	if err != nil {
		return nil, err
	}

	eType := entityMetadataType{Name: "business_transaction", Namespace: namespace{Name: "apm", Version: 1}}

	event := &bEvent{
		Type:           "bmetric:b_event",
		TraceID:        traceID.String(),
		EntityMetadata: entityMetadata{eType},
	}
	return event, nil
}

func getTraceID(td ptrace.Traces) (pcommon.TraceID, error) {
	rss := td.ResourceSpans()
	if rss.Len() == 0 {
		return pcommon.NewTraceIDEmpty(), errors.New("no resource spans are present")
	}

	ilss := rss.At(0).ScopeSpans()
	if ilss.Len() == 0 {
		return pcommon.NewTraceIDEmpty(), errors.New("no scope spans are present")
	}

	spans := ilss.At(0).Spans()
	if spans.Len() == 0 {
		return pcommon.NewTraceIDEmpty(), errors.New("no trace id is present")
	}

	return spans.At(0).TraceID(), nil
}
