package main

import (
	"errors"
)

// func BenchmarkInterface(b *testing.B) {
// 	var data []interface{}
// 	for i := 0; i < 10000; i++ {
// 		data = append(data, rand.Intn(10000))
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		sortData(data)
// 		calculateAverage(data)
// 	}
// }

// func BenchmarkGenerics(b *testing.B) {
// 	data := make([]int, 10000)
// 	for i := 0; i < 10000; i++ {
// 		data[i] = rand.Intn(10000)
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		sortDataGeneric(data)
// 		calculateAverageGeneric(data)
// 	}
// }

// // compareValues compares two reflect.Value objects and returns an integer representing the result of the comparison.
// // If the first value is less than the second, the result is negative.
// // If the first value is equal to the second, the result is zero.
// // If the first value is greater than the second, the result is positive.
// func compareValues(a, b reflect.Value) int {
// 	switch {
// 	case a.IsValid() && b.IsValid():
// 		// Both values are valid, compare them
// 		switch {
// 		case a.Type() == b.Type():
// 			// If the types are the same, compare directly
// 			switch a.Kind() {
// 			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 				return int(a.Int()) - int(b.Int())
// 			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 				return int(a.Uint()) - int(b.Uint())
// 			case reflect.Float32, reflect.Float64:
// 				if a.Float() < b.Float() {
// 					return -1
// 				} else if a.Float() > b.Float() {
// 					return 1
// 				}
// 				return 0
// 			case reflect.String:
// 				return compareString(a.String(), b.String())
// 			default:
// 				// For other types, use the default comparison
// 				return compareInterface(a.Interface(), b.Interface())
// 			}
// 		default:
// 			// Types are different, compare as interfaces
// 			return compareInterface(a.Interface(), b.Interface())
// 		}
// 	case a.IsValid():
// 		// Only the first value is valid, it's less than the invalid value
// 		return -1
// 	case b.IsValid():
// 		// Only the second value is valid, it's greater than the invalid value
// 		return 1
// 	default:
// 		// Both values are invalid, they are equal
// 		return 0
// 	}
// }

// // compareString compares two strings and returns an integer representing the result of the comparison.
// // If the first string is less than the second, the result is negative.
// // If the first string is equal to the second, the result is zero.
// // If the first string is greater than the second, the result is positive.
// func compareString(a, b string) int {
// 	if a < b {
// 		return -1
// 	} else if a > b {
// 		return 1
// 	}
// 	return 0
// }

// // compareInterface compares two interface{} values using the default comparison rules.
// // This falls back to the behavior of the "==" operator for interface{} values.
// func compareInterface(a, b interface{}) int {
// 	if a == b {
// 		return 0
// 	} else if a != b {
// 		// Assuming non-equality for interface{} means "a is greater than b"
// 		return 1
// 	}
// 	// Both values are nil
// 	return 0
// }

// func sortData(data []interface{}) {
// 	// Create a slice of reflect.Value to sort
// 	values := make([]reflect.Value, len(data))
// 	for i, item := range data {
// 		values[i] = reflect.ValueOf(item)
// 	}

// 	// Define a custom less function using reflection
// 	less := func(i, j int) bool {
// 		// Compare the values using reflection
// 		return compareValues(values[i], values[j]) < 0
// 	}

// 	// Perform a stable sort using the custom less function
// 	sort.SliceStable(values, less)

// 	// Copy the sorted values back to the original slice
// 	for i, value := range values {
// 		data[i] = value.Interface()
// 	}
// }

// func calculateAverage(data []interface{}) (float64, error) {
// 	// 实现计算平均值
// 	if len(data) == 0 {
// 		return 0, errors.New("empty slice")
// 	}

// 	var sum float64
// 	var count int

// 	for _, item := range data {
// 		switch v := item.(type) {
// 		case int:
// 			sum += float64(v)
// 			count++
// 		case int8:
// 			sum += float64(v)
// 			count++
// 		case int16:
// 			sum += float64(v)
// 			count++
// 		case int32:
// 			sum += float64(v)
// 			count++
// 		case int64:
// 			sum += float64(v)
// 			count++
// 		case float32:
// 			sum += float64(v)
// 			count++
// 		case float64:
// 			sum += v
// 			count++
// 		default:
// 			return 0, fmt.Errorf("unsupported type: %T", v)
// 		}
// 	}

// 	if count == 0 {
// 		return 0, errors.New("no numeric values found")
// 	}
// 	return sum / float64(count), nil
// }

// Comparer 接口定义了一个比较方法，任何实现了这个接口的类型都可以使用这个方法。

func CompareByInterface(a, b interface{}) (bool, error) {
	// Check if a and b are of type int
	aInt, okA := a.(int)
	bInt, okB := b.(int)

	if !okA || !okB {
		return false, errors.New("not int")
	}
	return aInt < bInt, nil
}

func CompareByGeneric[T int | float64](a, b T) bool {
	return a < b // 直接比较，需要类型支持操作符 <
}
