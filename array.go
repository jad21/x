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

// Reverse invierte el orden de los elementos en un array
func Reverse[T any](array []T) []T {
	length := len(array)
	reversed := make([]T, length)
	for i, v := range array {
		reversed[length-1-i] = v
	}
	return reversed
}

// Contains verifica si un array contiene un valor específico
func Contains[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
