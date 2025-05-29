package main

import "fmt"

func main() {
	//	myMap := NewHashMap()
	//
	//	myMap.Insert("key1", "value1")
	//
	//	value, exists := myMap.Get("key1")
	//	if exists {
	//		println("key1:", value.(string))
	//	} else {
	//		println("key1 does not exist")
	//	}
	//
	//	result := hashFunction("my-key")
	//	fmt.Println(result)

	tinyMap := NewHashMapWithCapacity(1)
	tinyMap.Insert("1", "value1")
	tinyMap.Insert("2", 1337)
	tinyMap.Insert("3", "value3")
	tinyMap.Insert("4", "value4")
	tinyMap.Insert("5", "value5")
	tinyMap.Insert("6", "value6")

	allValues := tinyMap.GetAll()
	for _, value := range allValues {
		fmt.Printf("%s: %v\n", value.key, value.value)
	}
}
