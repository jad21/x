package hgo

// If es una función genérica que devuelve el valor x si condition es true, de lo contrario devuelve y.
func If[T any](condition bool, x, y T) T {
	if condition {
		return x
	}
	return y
}
