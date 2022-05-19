package smartpay

// Ptr Inline pointer helper
func Ptr[T any](v T) *T {
	return &v
}
