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

	fmt.Println("Traversing Map...")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values from map...")
	courseName, ok := m["course"]
	fmt.Println("CourseName = ", courseName, ok)
	if causeName, ok := m["cause"]; ok {
		fmt.Println("CauseName = ", causeName)
	} else {
		fmt.Println("CauseName = key does not exist!")
	}

	fmt.Println("Deleting values from map...")
	delete(m, "name")
	fmt.Println(m)
}
