package grpcx

func ConvertInt(v any) int {
	if v == nil {
		return 0
	}

	switch x := v.(type) {
	case int:
		return x
	case int32:
		return int(x)
	case int64:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	case *int:
		if x == nil {
			return 0
		}

		return *x
	case *int32:
		if x == nil {
			return 0
		}

		return int(*x)
	case *int64:
		if x == nil {
			return 0
		}

		return int(*x)
	default:
		return 0
	}
}

func ConvertFloat(v any) float64 {
	if v == nil {
		return 0.0
	}

	switch x := v.(type) {
	case float32:
		return float64(x)
	case float64:
		return x
	case int:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	case *float32:
		if x == nil {
			return 0.0
		}

		return float64(*x)
	case *float64:
		if x == nil {
			return 0.0
		}

		return *x
	default:
		return 0.0
	}
}
