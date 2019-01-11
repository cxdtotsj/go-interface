package tools

import (
	"strconv"
	"time"
)

// RandInt every timestamp for sleep 1s
func RandInt() string {

	currentTime := time.Now().Unix()
	ts := strconv.FormatInt(currentTime, 10)
	time.Sleep(1 * time.Second)
	return ts
}

// func main() {
// 	for i := 0; i <= 20; i++ {
// 		fmt.Println(RandInt())
// 	}
// }
