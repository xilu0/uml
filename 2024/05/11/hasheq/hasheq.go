package hasheq

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// hashObject 计算对象的哈希值
func hashObject(obj *unstructured.Unstructured) (string, error) {
	data, err := json.Marshal(obj.Object)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}

// compareHashes 比较两个对象的哈希值
func compareHashes(obj1, obj2 *unstructured.Unstructured) (bool, error) {
	hash1, err := hashObject(obj1)
	if err != nil {
		return false, err
	}
	hash2, err := hashObject(obj2)
	if err != nil {
		return false, err
	}
	return hash1 == hash2, nil
}
