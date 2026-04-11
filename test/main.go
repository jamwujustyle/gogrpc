package main

import "fmt"

func test(strs ...string) []string {
	arr:= []string{}

	for _, str:= range strs {
		arr = append(arr, str)
	}

	return arr
}
func debug(args ...any) {
	fmt.Println(args...)
}
func main() {

	str1:= "hello"
	str2:= "bitch"
	str3:= "my"
	str4:= "name"
	str5:= "is"
	str6:= "bitch"

	fmt.Println(test(str1, str2, str3, str4, str5, str6))

	vals:= []any{1, "hello", true}

	debug(vals)
}