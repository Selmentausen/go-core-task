package main

func MapDifference(slice1, slice2 []string) []string {
	exclusionMap := make(map[string]struct{})
	for _, item := range slice2 {
		exclusionMap[item] = struct{}{}
	}

	result := make([]string, 0, len(slice1))
	for _, item := range slice1 {
		if _, exists := exclusionMap[item]; !exists {
			result = append(result, item)
		}
	}
	return result
}
