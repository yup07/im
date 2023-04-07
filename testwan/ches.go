package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	t, _ := json.Marshal("123456")
	//t := time.Now()
	fmt.Printf("t:%T,值：%s", t, t)
}
