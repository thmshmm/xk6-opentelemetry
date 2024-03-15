package generator

import (
	"encoding/hex"
	"log/slog"
	"strings"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

type TraceConfig struct {
	Data []SpanData `js:"data"`
}

type SpanData struct {
	TraceID           string                 `js:"traceId"`
	SpanID            string                 `js:"spanId"`
	ParentSpanID      string                 `js:"parentSpanId"`
	Name              string                 `js:"name"`
	Attributes        map[string]interface{} `js:"attributes"`
	Kind              string                 `js:"kind"`
	StartTimeUnixNano uint64                 `js:"startTimeUnixNano"`
	EndTimeUnixNano   uint64                 `js:"endTimeUnixNano"`
}

func ResourceSpans(resourceAttrs []*commonpb.KeyValue, config TraceConfig) *tracepb.ResourceSpans {
	return &tracepb.ResourceSpans{
		Resource: resource(resourceAttrs),
		ScopeSpans: []*tracepb.ScopeSpans{
			{
				Spans: spans(config.Data),
			},
		},
	}
}

func spans(data []SpanData) []*tracepb.Span {
	spans := []*tracepb.Span{}

	for _, spanData := range data {
		spans = append(spans, span(spanData))
	}

	return spans
}

func span(data SpanData) *tracepb.Span {
	traceID, err := hex.DecodeString(data.TraceID)
	if err != nil {
		slog.Error("invalid trace id")
	}

	spanID, err := hex.DecodeString(data.SpanID)
	if err != nil {
		slog.Error("invalid span id")
	}

	parentSpanID, err := hex.DecodeString(data.ParentSpanID)
	if err != nil {
		slog.Error("invalid parent span id")
	}

	return &tracepb.Span{
		TraceId:           traceID,
		SpanId:            spanID,
		ParentSpanId:      parentSpanID,
		Name:              data.Name,
		Attributes:        ToAttributes(data.Attributes),
		Kind:              getSpanKind(data.Kind),
		StartTimeUnixNano: data.StartTimeUnixNano,
		EndTimeUnixNano:   data.EndTimeUnixNano,
	}
}

func getSpanKind(kind string) tracepb.Span_SpanKind {
	switch strings.ToLower(kind) {
	case "client":
		return tracepb.Span_SPAN_KIND_CLIENT
	case "server":
		return tracepb.Span_SPAN_KIND_SERVER
	case "consumer":
		return tracepb.Span_SPAN_KIND_CONSUMER
	case "producer":
		return tracepb.Span_SPAN_KIND_PRODUCER
	}

	return tracepb.Span_SPAN_KIND_INTERNAL
}
