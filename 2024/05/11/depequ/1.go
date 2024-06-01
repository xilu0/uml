package depequ

import (
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// compareObjects 使用 DeepEqual 比较两个对象
func compareObjects(obj1, obj2 *unstructured.Unstructured) bool {
	return reflect.DeepEqual(obj1.Object, obj2.Object)
}
