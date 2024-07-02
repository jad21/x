package x

import (
	"fmt"
)

func Int64ToString(v int64) string {
	return fmt.Sprintf("%d", v)
}

func UInt64ToString(v uint64) string {
	return fmt.Sprintf("%d", v)
}

func IntToString(v int) string {
	return fmt.Sprintf("%d", v)
}

func UIntToString(v uint) string {
	return fmt.Sprintf("%d", v)
}

func ToListInterfaces[T any](array []T) []interface{} {
	return MapIndex(len(array), func(i int) interface{} {
		return array[i]
	})
}
