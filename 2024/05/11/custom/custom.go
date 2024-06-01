package custom

import (
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// comparePointers 比较指针指向的值
func comparePointers(ptr1, ptr2 interface{}) bool {
	if ptr1 == nil && ptr2 == nil {
		return true
	}
	if ptr1 == nil || ptr2 == nil {
		return false
	}
	return reflect.DeepEqual(ptr1, ptr2)
}

// compareCustomObject 使用自定义比较函数比较对象
func compareCustomObject(obj1, obj2 *unstructured.Unstructured) bool {
	for key, value1 := range obj1.Object {
		value2, exists := obj2.Object[key]
		if !exists || !comparePointers(value1, value2) {
			return false
		}
	}
	return true
}
