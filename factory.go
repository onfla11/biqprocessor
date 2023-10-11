// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bprocessor

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

func NewFactory() processor.Factory {

	return processor.NewFactory(
		"bprocessor",
		createDefaultConfig,
		processor.WithTraces(createBProcessor, component.StabilityLevelDevelopment),
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
