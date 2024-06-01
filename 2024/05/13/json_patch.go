package main

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

func main() {
	// 原始 JSON 文档
	original := []byte(`{
        "name": "John",
        "age": 30,
        "city": "New York"
    }`)

	// JSON Patch 文档
	patch := []byte(`[
        { "op": "replace", "path": "/name", "value": "Jane" },
        { "op": "remove", "path": "/age" },
        { "op": "add", "path": "/country", "value": "USA" }
    ]`)

	// 创建补丁对象
	patchObj, err := jsonpatch.DecodePatch(patch)
	if err != nil {
		panic(err)
	}

	// 应用补丁
	patched, err := patchObj.Apply(original)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Patched document: %s\n", patched)
}
