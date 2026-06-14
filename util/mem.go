package util

func Ptr[T any](v T) *T {
	return &v

}
func Deref[T any](p *[]T) []T {
	if p == nil {
		return nil
	}
	return *p
}
