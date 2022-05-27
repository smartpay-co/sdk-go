package smartpay

import "encoding/json"

// Ptr Inline pointer helper
func Ptr[T any](v T) *T {
	return &v
}

func ConvertToStruct[T any](from interface{}) (to T, err error) {
	jsonData, err := json.Marshal(from)
	if err != nil {
		return to, err
	}
	err = json.Unmarshal(jsonData, &to)
	if err != nil {
		return to, err
	}
	return
}
