package main

func main() {

	slice := make([]int, 0, 10)
	slice = append(slice, 6, 4, 1, 3, 8, 2, 9, 0, 7, 5)

	BubbleSort(slice)

}

func BubbleSort(slice []int) {

	for i := 0; i < len(slice); i++ {

		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(&slice[j], &slice[j+1])

			}
		}
	}
	for k := 0; k < len(slice); k++ {
		print(slice[k])
		print("\n")
	}
}

// we dont use temp
func Swap(first, second *int) {

	*first, *second = *second, *first

}
