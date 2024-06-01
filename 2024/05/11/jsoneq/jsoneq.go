package jsoneq

import (
	jsonpatch "github.com/evanphx/json-patch"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// generateJSONPatch 生成两个对象之间的 JSON Patch
func generateJSONPatch(oldObj, newObj *unstructured.Unstructured) ([]byte, error) {
	oldData, err := oldObj.MarshalJSON()
	if err != nil {
		return nil, err
	}
	newData, err := newObj.MarshalJSON()
	if err != nil {
		return nil, err
	}
	patch, err := jsonpatch.CreateMergePatch(oldData, newData)
	if err != nil {
		return nil, err
	}
	return patch, nil
}

// isPatchEmpty 检查 JSON Patch 是否为空
func isPatchEmpty(patch []byte) bool {
	return len(patch) == 0 || string(patch) == "{}"
}
