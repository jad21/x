package x

// If es una función genérica que devuelve el valor x si condition es true, de lo contrario devuelve y.
func If[T any](condition bool, True, False T) T {
	if condition {
		return True
	}
	return False
}
