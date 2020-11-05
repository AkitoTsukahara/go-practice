package main

import(
	"fmt"
	"./animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.ElephpantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}