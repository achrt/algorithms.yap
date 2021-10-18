package main

// import "fmt"

// func main() {
// 	test()
// }

func brokenSearch(arr []int, k int) int {
	d := &Data{
		digits: arr,
		lenght: len(arr),
	}
	return d.Search(0, d.lenght-1, k)
}

// func test() {
// 	arr := []int{5,1}
// 	t := brokenSearch(arr, 8)

// 	fmt.Println(t)
// 	if t != 6 {
// 		fmt.Println("WA")
// 	}
// }

type Data struct {
	digits []int
	lenght int
}

func (d *Data) Len() int {
	return d.lenght
}

func (d *Data) Search(left, right, k int) int {
	middle := (right + left) / 2

	// проверка правых краев разбитого массива
	if d.digits[right] == k {
		return right
	}
	if d.digits[middle] == k {
		return middle
	}

	// если с края не нашлось, проверяем в какой половине искать

	if d.digits[middle] > k {
		return d.Search(0, middle, k)
	}
	if d.digits[right] > k {
		return d.Search(middle, right, k)
	}
	return -1
}
