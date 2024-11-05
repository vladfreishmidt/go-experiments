package newdatatypes

import "errors"

type TreeLast[T any] []T

// replaceLast is a method that operates on TreeLast variables
func (t TreeLast[T]) replaceLast(element T) (TreeLast[T], error) {
	if len(t) == 0 {
		return t, errors.New("This is empty!")
	}

	t[len(t)-1] = element
	return t, nil
}

// Usage:
// tempStr := TreeLast[string]{"aa", "bb"}
// tempStr.replaceLast("cc")
