// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/bprocessor"

import (
	"context"

	"github.com/onfla11/bprocessor/internal/metadata"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

func NewFactory() processor.Factory {

	return processor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processor.WithTraces(createBProcessor, metadata.MetricsStability),
	)
}

// createsBiqProcessor creates a trace processor based on this config.
func createBProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	next consumer.Traces) (processor.Traces, error) {

	// the only supported storage for now
	generator := &eventGenerator{}
	oCfg := cfg.(*Config)
	b := newBProcessor(ctx, oCfg, generator)

	return processorhelper.NewTracesProcessor(ctx, set, cfg, next, b.processTraces)
}
