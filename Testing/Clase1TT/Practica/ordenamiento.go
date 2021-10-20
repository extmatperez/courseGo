package ordenamientos

import "sort"

func ordenar(num []int) {
	sort.Slice(num, func(i, j int) bool { return num[i] < num[j] })
}
