package x

// If es una función genérica que devuelve el valor x si condition es true, de lo contrario devuelve y.
func If[T any](condition bool, True, False T) T {
	if condition {
		return True
	}
	return False
}

// SwitchMap es una función genérica que devuelve el valor asociado a una clave si la encuentra, de lo contrario devuelve el valor por defecto.
func SwitchMap[K comparable, V any](key K, cases map[K]V, defaultValue V) V {
	if value, found := cases[key]; found {
		return value
	}
	return defaultValue
}

// If es una función genérica que devuelve el valor x si condition es true, de lo contrario devuelve y. Pero los valores son funciones
func IfFn[T any](condition bool, TrueFn, FalseFn func() T) T {
	if condition {
		return TrueFn()
	}
	return FalseFn()
}
