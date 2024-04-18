package main

// SOLUTION
func TwoSum(list []int, total int) []int {
	m := make(map[int]int)

	for i, v := range list {
		rest := total - v
		_, exists := m[rest]
		if !exists {
			m[v] = i
		} else {
			return []int{m[rest], i}
		}
	}

	return []int{0, 0}
}
