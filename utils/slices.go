package utils

// Contains checks if an element is present in a slice.
//
// Parameters:
// - s: The slice of elements to search.
// - e: The element to check for.
//
// Returns:
// - bool: True if the element is found, false otherwise.
func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Remove removes the first occurrence of element e from slice s.
//
// Parameters:
// - s: The input slice of type []T, where T is a comparable type.
// - e: The element to be removed from the slice.
//
// Returns:
// - []T: The modified slice with the element removed.
func Remove[T comparable](s []T, e T) []T {
	for i, a := range s {
		if a == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// RemoveAll removes all occurrences of element e from slice s.
// Like Remove, but all.
func RemoveAll[T comparable](s []T, e T) []T {
	var result []T
	for _, a := range s {
		if a != e {
			result = append(result, a)
		}
	}
	return result
}

// Unique returns a new slice containing the unique elements of the input slice s.
//
// The parameter s is a slice of elements of a comparable type.
// The return type is a slice of the same type as the elements of s.
func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
