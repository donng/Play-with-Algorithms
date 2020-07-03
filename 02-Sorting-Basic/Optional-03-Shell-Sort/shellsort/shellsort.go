package shellsort

func Sort(arr []int, length int) {
	h := 1
	for h < length/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < length; i++ {
			e := arr[i]
			j := i
			for ; j >= h && e < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}
