package smartpay

import "encoding/json"

// Ptr Inline pointer helper
func Ptr[T any](v T) *T {
	return &v
}

func ConvertToStruct[T any](from interface{}) (to T) {
	jsonData, err := json.Marshal(from)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonData, &to)
	if err != nil {
		panic(err)
	}
	return
}
