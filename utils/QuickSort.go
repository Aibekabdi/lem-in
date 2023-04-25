package lemin

func partition(arr [][]string, low, high int) ([][]string, int) {
	pivot := len(arr[high])
	i := low
	for j := low; j < high; j++ {
		if len(arr[j]) < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort(arr [][]string, low, high int) [][]string {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}
