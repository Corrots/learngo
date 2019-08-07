package main

import "fmt"

func getCountries() (c []string) {
	countries := []string{"United states", "United kingdom", "Austrilia", "India",
		"China", "Russia", "France", "Germany", "Spain"} // can be much more

	//return countries[:3]
	/**
		为了避免写出上面糟糕的代码 (函数返回的切片引用了一个比较大的数组，占用了较多的内存)，
		我们必须创建一个新的匿名数组切片，使其长度固定。下面的程序是个不错的方案。
	 */
	 c = make([]string, 2)
	 copy(c, countries[:2])
	 return
}

func main() {
	c := getCountries()

	fmt.Println(c)
	fmt.Println(cap(c))
}
