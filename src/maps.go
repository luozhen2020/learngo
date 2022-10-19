package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "not bad",
	}

	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map...")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values from map...")
	qualityValue, ok := m["quality"]
	fmt.Println(qualityValue, ok)
	if wrongQuality, ok := m["quall"]; ok {
		fmt.Println(wrongQuality)
	} else {
		fmt.Println("Quality does not exist!")
	}

	fmt.Println("Deleting values from map...")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)
}
