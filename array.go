package hgo

// Find: encontrar el primero que cumpla el criterio
func Find[T any](array []T, cb func(item T) bool) *T {
	for _, v := range array {
		if cb(v) {
			return &v
		}
	}
	return nil
}
