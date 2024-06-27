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

// Map aplica una función a cada elemento de un array y devuelve un nuevo array con los resultados
func Map[T any, U any](array []T, f func(T) U) []U {
	mapped := make([]U, len(array))
	for i, v := range array {
		mapped[i] = f(v)
	}
	return mapped
}

// Repeat repite el valor el numero de veces que se desea
func Repeat[T any](v T, count int) (results []T) {
	for i := 0; i < count; i++ {
		results = append(results, v)
	}
	return
}
