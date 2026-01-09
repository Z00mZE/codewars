package kata

func MoveZeros(arr []int) []int {
	out := make([]int, len(arr))
	var i int
	for _, tmp := range arr {
		if tmp != 0 {
			out[i] = tmp
			i++
		}
	}
	return out
}
