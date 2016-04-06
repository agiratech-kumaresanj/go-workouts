package main

import ("fmt")

func arun(a map[int]string) {
	fmt.Println(a)
}

func main() {
	mapVaribale := make(map[int]string)
	for i := 0 ; i<10; i++ {
		mapVaribale[i] = "test"
	}
	fmt.Println(mapVaribale)
	fmt.Printf("the value of the 3rd element in the Map : %v\n",mapVaribale[2])
	mapVaribale[10] = "arun"
	fmt.Println(mapVaribale)
	delete(mapVaribale,1)
	// fmt.Println(mapVaribale)
	arun(mapVaribale)
}