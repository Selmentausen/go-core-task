package main

func MapIntersection(slice1, slice2 []int) (bool, []int) {
	elementsMap := make(map[int]struct{})
	for _, item := range slice1 {
		elementsMap[item] = struct{}{}
	}

	var intersection []int
	for _, item := range slice2 {
		if _, exists := elementsMap[item]; exists {
			intersection = append(intersection, item)
		}
	}
	if len(intersection) > 0 {
		return true, intersection
	}
	return false, []int{}
}
