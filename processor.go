// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bprocessor
import (
	"context"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

type bProcessor struct {
	eventGenerator *eventGenerator
	// b processor configuration
	config *Config
}

func newBProcessor(_ context.Context, config *Config, generator *eventGenerator) *bProcessor {
	bq := &bProcessor{
		eventGenerator: generator,
		config:         config,
	}
	return bq
}

func (bq *bProcessor) processTraces(_ context.Context, batch ptrace.Traces) (ptrace.Traces, error) {
	for _, singleTrace := range batchpersignal.SplitTraces(batch) {
		event, err := bq.eventGenerator.generate(singleTrace)
		if err != nil {
			return batch, err
		}
		if event != nil {
			fmt.Println("Event Generated: {}", event)
		}
		fmt.Println("processTraces: ", singleTrace.ResourceSpans().Len())
		singleTrace.ResourceSpans().Len()
	}
	return batch, nil
}
