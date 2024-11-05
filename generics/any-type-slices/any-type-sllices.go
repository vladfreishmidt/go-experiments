package anytypeslices

// Slice constrains - the function parameter can be a slice of any data type.
// All three generic functions are equivalent
// ~[]E specifies that underlying data type should be a slice if it is a type by a different name.
func f1[S interface{ ~[]E }, E interface{}](x S) int {
	return len(x)
}

func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}

func f3[S ~[]E, E any](x S) int {
	return len(x)
}
