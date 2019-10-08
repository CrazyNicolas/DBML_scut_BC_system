package PBFT_module

import "strconv"

func ToString(args interface{}) string {
	switch args.(type) {
	case int:
		return strconv.FormatInt(int64(args.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(args.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(args.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(args.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(args.(int64)), 10)
	}
	return ""
}
