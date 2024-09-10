package main

import "fmt"

func main() {

	var smsSendingLimit int = 10
	var costPerSMS float32 = 5
	var hasPermission bool = true
	var username string = "Om"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
