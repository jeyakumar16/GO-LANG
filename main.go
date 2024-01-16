package main

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with the capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with the capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)

	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("\nThe length is %v with the capacity %v\n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"JK": 20, "Bharath": 21}
	fmt.Println(myMap2["Bharath"])

	var age, ok = myMap2["Hariesh"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name, key doesn't exist")
	}

	for name, age := range myMap2 {
		fmt.Printf("The name is %v and the their age is %v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("The index is %v and the corresponding value is %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i int = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}
