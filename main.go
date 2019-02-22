package main


func main() {
	left := make([]int, 6)
	left[0] = 2
	left[1] = 5
	left[2] = 7
	left[3] = 5
	left[4] = 5
	left[5] = 6
	inplace(0, len(left) / 2, left)
	for _, val := range left {
		println(val)
	}
}

func inplace(left, right int, vals []int) {
	length := len(vals)
	start := right
	for left < length {
		if vals[left] > vals[right] {
			swap(left, right, vals)
			left++
		} else {
			right++
		}
		if right == length {
			right = start
			if right < left {
				right = left
			}
			left++
		}
	}
}

func swap(left, right int, vals []int) {
	temp := vals[left]
	vals[left] = vals[right]
	vals[right] = temp
}