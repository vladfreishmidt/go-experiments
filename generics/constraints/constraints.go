package constraints

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

type AnotherInt int
type AllInts interface {
	~int
}

// AddElements of the slice int type
func AddElements[T AllInts](s []T) T {
	sum := T(0)

	for _, v := range s {
		sum = sum + v
	}

	return sum
}
