package heap

func getParent(i int) int {
	return (i - 1) / 2
}

func getLeft(i int) int {
	return (2*i + 1)
}

func getRight(i int) int {
	return (2*i + 2)
}

func MaxHeapify(nums []int, i int) {
	n := len(nums)
	l := getLeft(i)
	r := getRight(i)
	var largest int
	if l < n && nums[l] > nums[i] {
		largest = l
	} else {
		largest = i
	}
	if r < n && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		MaxHeapify(nums, largest)
	}
}
