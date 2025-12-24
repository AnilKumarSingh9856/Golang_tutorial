package main

import "fmt"

// A map is an unordered and changeable collection that does not allow duplicates.

// The default value of a map is nil.

// Maps hold references to an underlying hash table.

// Go has multiple ways for creating maps.


func main() {
	// 1st way Create Maps Using var and :=
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
    // b := map[KeyType]ValueType{key1:value1, key2:value2,...}

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)


	// 2nd way Create Maps Using the make() Function
	// var a = make(map[KeyType]ValueType)
	// b := make(map[KeyType]ValueType)

	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"
									// c is no longer empty
	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("a\t%v\n", c)
	fmt.Printf("b\t%v\n", d)

	// Create an Empty Map
	// 	There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.

	// Syntax
	// var a map[KeyType]ValueType
	// Note: The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.

	// 	Removing elements is done using the delete() function.

	// Syntax
	// delete(map_name, key)
	delete(d, "Stavanger")

	// You can check if a certain key exist in a map using
	// Syntax
	// val , ok := map_name[key]

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value
    

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// Note: 
	// Maps Are References
	// Maps are references to hash tables.

	// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

	// Iterate Over Maps by using range
	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}


