package main

import "fmt"

// assign2dMap 二维Map赋值
// k1 一维key
// k2 二维key
// v 二维值
func assign2dMap(k1 string, k2 string, v interface{}, ret map[string]map[string]interface{}) {

	item, ok := ret[k1]
	if !ok {
		item = make(map[string]interface{})
		ret[k1] = item
	}
	item[k2] = v
}

func main() {


	//判断key是否存在
	//a1 := map[string]string{
	//	"name": "张三",
	//	"address": "西安市",
	//}
	//
	//fmt.Printf("%#v", a1)
	//
	//k, v := a1["age"]
	//fmt.Printf("%#v", k) //
	//fmt.Printf("%#v", v)

	//
	ret := make(map[string]map[string]interface{})
	assign2dMap("key1", "k_a", "a_value", ret)
	assign2dMap("key1", "k_b", "b_value", ret)

	fmt.Println(" main() 执行结果 ")
	fmt.Printf("%#v", ret)


}
