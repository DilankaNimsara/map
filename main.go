package main

import "fmt"

func main() {

	////method one
	//color := map[string]string{
	//	"red":   "ff0000",
	//	"white": "000000",
	//	"black": "ffffff",
	//}
	//
	//fmt.Println(color)

	// method two
	var colors = map[string]string{}
	fmt.Println(colors)

	//method three
	colorsMap := make(map[string]string)
	colorsMap["red"] = "ff0000"
	colorsMap["white"] = "000000"

	fmt.Println(colorsMap["white"])
	delete(colorsMap, "red")

	fmt.Println(colorsMap)
	mapTest()
}

func mapTest() {
	colors := map[int]string{
		1: "red",
		2: "black",
		3: "green",
		4: "white",
	}
	printMap(colors)
}

func printMap(m map[int]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}
