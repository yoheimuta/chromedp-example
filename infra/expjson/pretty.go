package expjson

import (
	"encoding/json"
	"fmt"
)

// PrettyFormat はオブジェクトの JSON を返す。デバッグ目的で print するときに使う想定。
func PrettyFormat(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("orig=%v, err=%v", v, err)
	}
	return string(b)
}
