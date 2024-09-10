package main

import "fmt"

func main() {

	var smsSendingLimit int = 10
	var costPerSMS float32 = 5
	var hasPermission bool = true
	// var username string = "Om"
	username := "Om"
	name := username // name is also is a string
	const premiumPlanName = "Premium Plan"
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	fmt.Printf("%q", name)
	fmt.Println("plan:", premiumPlanName)
}
