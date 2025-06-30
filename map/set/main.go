package main

import "fmt"

type Set map[int]bool

func (s Set) ToSlice() []int {
	slice := make([]int, 0, len(s))
	for element := range s {
		slice = append(slice, element)
	}
	return slice
}

func NewSet(elements ...int) Set {
	set := make(Set)
	for _, element := range elements {
		set[element] = true
	}
	return set
}

func (s1 Set) Union(s2 Set) []int {
	result := NewSet()
	for elem := range s1 {
		result[elem] = true
	}

	for elem := range s2 {
		result[elem] = true
	}
	return result.ToSlice()
}

func (s Set) Difference(s2 Set) []int {
	result := NewSet()
	for elem := range s {
		if !s2[elem] {
			result[elem] = true
		}
	}
	return result.ToSlice()
}

func (s Set) Intersection(s2 Set) []int {
	result := NewSet()
	for elem := range s {
		if s2[elem] {
			result[elem] = true
		}
	}
	return result.ToSlice()
}

func main() {

	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)

	// Объединение
	union := set1.Union(set2)
	fmt.Println("Union:", union)

	// Различие
	difference := set1.Difference(set2)
	fmt.Println("difference:", difference)

	// Пересечние
	intersection := set1.Intersection(set2)
	fmt.Println("intersection:", intersection)

}
