package main

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/go-version"
	//"strconv"
)

func main() {

	//版本号比较
	v1 := "1.0.9"
	v2 := "1.0.10"

	//iv1, _ := strconv.Atoi(v1)
	//iv2, _ := strconv.Atoi(v2)
	//fmt.Println("iv1 = ", iv1)
	//fmt.Println("iv2 = ", iv2)
	//
	//if v1 < v2 {
	//	fmt.Printf("%s < %s", v1,v2)
	//}else {
	//	fmt.Printf("%s > %s", v1,v2)
	//}

	v11, _ := version.NewVersion(v1)
	v22, _ := version.NewVersion(v2)

	if v11.LessThan(v22) {
		fmt.Printf("%s < %s\n", v1, v2)
	} else {
		fmt.Printf("%s > %s\n", v1, v2)
	}

	//数字转字符串
	versionCode := 36
	versionCodeStr := strconv.Itoa(versionCode)
	fmt.Printf("%#v\n", versionCodeStr)

	//字符串转数字
	versionCodeStr1 := "36"
	versionCode1, _ := strconv.Atoi(versionCodeStr1)
	fmt.Printf("%#v\n", versionCode1)

	//TODO: 字符串截取
	//https://www.cnblogs.com/xuweiqiang/p/12558912.html

	traceId := "abcdefg"
	prefixLetter := traceId[len(traceId)-1:]
	fmt.Println("prefixLetter: ", prefixLetter)

}
