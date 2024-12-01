package compare

func If[T any](condition bool, doIfTrue T, doIfFalse T) T {
	if condition {
		return doIfTrue
	}
	return doIfFalse
}
