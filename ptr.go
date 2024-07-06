package ptr

import (
	"time"
)

// Ptr 是一个通用的指针生成器
func Ptr[T any](v T) *T {
	return &v
}

// String 返回字符串的指针
func String(v string) *string {
	return Ptr(v)
}

// Int 返回整数的指针
func Int(v int) *int {
	return Ptr(v)
}

// Bool 返回布尔值的指针
func Bool(v bool) *bool {
	return Ptr(v)
}

// Float64 返回float64的指针
func Float64(v float64) *float64 {
	return Ptr(v)
}

// Time 返回time.Time的指针
func Time(v time.Time) *time.Time {
	return Ptr(v)
}

// SlicePtr 返回切片的指针
func SlicePtr[T any](v []T) *[]T {
	return Ptr(v)
}

// MapPtr 返回映射的指针
func MapPtr[K comparable, V any](v map[K]V) *map[K]V {
	return Ptr(v)
}

// IsNil 检查指针是否为nil
func IsNil[T any](v *T) bool {
	return v == nil
}

// GetValueOrDefault 返回指针的值，如果指针为nil则返回默认值
func GetValueOrDefault[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}
	return defaultValue
}

// GetNonZeroValueOrDefault 返回非零值，或者默认值
func GetNonZeroValueOrDefault[T comparable](value *T, defaultValue T) T {
	if value == nil || *value == *new(T) {
		return defaultValue
	}
	return *value
}

// NewPtrIf 根据条件创建指针
func NewPtrIf[T any](condition bool, v T) *T {
	if condition {
		return Ptr(v)
	}
	return nil
}

// Equal 比较两个指针指向的值是否相等
func Equal[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// Map 对切片中的每个元素应用函数f
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Filter 过滤切片中满足条件的元素
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
