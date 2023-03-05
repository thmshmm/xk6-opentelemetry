package generator

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
)

const unsupportedTypeValue = "unsupported_type"

func ToAnyValue(attr interface{}) *commonpb.AnyValue {
	var value *commonpb.AnyValue

	switch attrValue := attr.(type) {
	case string:
		value = &commonpb.AnyValue{
			Value: &commonpb.AnyValue_StringValue{
				StringValue: attrValue,
			},
		}
	case int64:
		value = &commonpb.AnyValue{
			Value: &commonpb.AnyValue_IntValue{
				IntValue: attrValue,
			},
		}
	case int:
		value = &commonpb.AnyValue{
			Value: &commonpb.AnyValue_IntValue{
				IntValue: int64(attrValue),
			},
		}
	case float64:
		value = &commonpb.AnyValue{
			Value: &commonpb.AnyValue_DoubleValue{
				DoubleValue: attrValue,
			},
		}
	case bool:
		value = &commonpb.AnyValue{
			Value: &commonpb.AnyValue_BoolValue{
				BoolValue: attrValue,
			},
		}
	default:
		value = &commonpb.AnyValue{
			Value: &commonpb.AnyValue_StringValue{
				StringValue: unsupportedTypeValue,
			},
		}
	}

	return value
}
