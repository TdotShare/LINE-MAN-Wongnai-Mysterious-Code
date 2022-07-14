package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = action_reverse(string(sd))
	fmt.Println(whatIsIt)
}

func action_reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
		//fmt.Println(result)
	}
	return result
}
