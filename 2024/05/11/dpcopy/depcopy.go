package dpcopy

import (
	"reflect"

	"github.com/mohae/deepcopy"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// deepCopyObject 深度复制对象
func deepCopyObject(obj *unstructured.Unstructured) *unstructured.Unstructured {
	copy := deepcopy.Copy(obj.Object).(map[string]interface{})
	return &unstructured.Unstructured{Object: copy}
}

// compareObjectsUsingDeepCopy 使用深度复制比较对象
func compareObjectsUsingDeepCopy(obj1, obj2 *unstructured.Unstructured) bool {
	copy1 := deepCopyObject(obj1)
	copy2 := deepCopyObject(obj2)
	return reflect.DeepEqual(copy1.Object, copy2.Object)
}
