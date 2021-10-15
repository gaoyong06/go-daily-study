package main

import "fmt"

//Unique 切片去重
func Unique(s []int) []int {

	ret := make([]int, 0)
	temp := make(map[int] bool)

	for _, v := range s {
		if _, ok := temp[v]; !ok {
			ret = append(ret, v)
			temp[v] = true
		}
	}
	return  ret
}

func main() {

	a1 := []int {1, 2, 3, 4, 5, 1}
	fmt.Println(a1)

	a2 := Unique(a1)
	fmt.Println(a2)


}
