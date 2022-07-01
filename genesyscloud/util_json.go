package genesyscloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func suppressEquivalentJsonDiffs(k, old, new string, d *schema.ResourceData) bool {
	if old == new {
		return true
	}

	ob := bytes.NewBufferString("")
	if err := json.Compact(ob, []byte(old)); err != nil {
		return false
	}

	nb := bytes.NewBufferString("")
	if err := json.Compact(nb, []byte(new)); err != nil {
		return false
	}

	return jsonBytesEqual(ob.Bytes(), nb.Bytes())
}

func jsonBytesEqual(b1, b2 []byte) bool {
	var o1 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}

	var o2 interface{}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}

	return reflect.DeepEqual(o1, o2)
}

func interfaceToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func interfaceToJson(val interface{}) string {
	j, _ := json.Marshal(val)
	return string(j)
}

func jsonStringToInterface(jsonStr string) (interface{}, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal %s: %v", jsonStr, err)
	}
	return obj, nil
}
