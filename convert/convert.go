package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func ToInt(value any, defaultValue int) int {
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		res, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return defaultValue
		}
		return res
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return defaultValue
	}
}

func ToFloat64(value any, defaultValue float64) float64 {
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case string:
		res, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return defaultValue
		}
		return res
	case bool:
		if v {
			return 1.0
		}
		return 0.0
	default:
		return defaultValue
	}
}

func ToString(value any) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case bool:
		return strconv.FormatBool(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%g", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func ToBool(value any, defaultValue bool) bool {
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case bool:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return ToInt(v, 0) != 0
	case float32, float64:
		return ToFloat64(v, 0.0) != 0.0
	case string:
		res, err := strconv.ParseBool(strings.ToLower(strings.TrimSpace(v)))
		if err != nil {
			return defaultValue
		}
		return res
	default:
		return defaultValue
	}
}

