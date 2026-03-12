// You can edit this code!
// Click here and start typing.
package main

import "fmt"

/*
var A int;
A = 2

A := 2

{
  k: v,
  2: 3
}
*/

type model int

/*
{ 
  0: "toyota",
  1: "audi"
*/

const (
   toyota model = iota
   audi
   bmw
   mercedes
)

type Ticket string

type Human struct {
	Name        string
	SSN         string
	Age         int
	Temperature []int
	Dependents  []Human
	Cars	[]Car
}

// func get_fine(t string)
// func get_fine(tkt Ticket) 



type Car struct {
	ModelStr string
	Vin	 int
	tickets  []Ticket
}

type EthicPerson struct {
 	Property Human
	Additional map[string]string
}

type EthicPerson2 struct {
 	Property Human
	Additional map[string]string
}

func main() {

	//list1 := make([]int, 10)

	map1 := make(map[int]int) // {}
	fmt.Println(map1)
	map1[1] = 2

	map2 := map[int]int{}
	fmt.Println(map2)

	map3 := map[int]int{
		1: 12,
		2: 23,
	}
	fmt.Println(map3)

	var map4 map[int]int // nil
	fmt.Println(map4)
	map4 = map[int]int{}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		fmt.Printf("%d\n", i)
	}

	medicalDir := make(map[string]Human)
	medicalDir["abc"] = Human{
		Name:        "abc",
		SSN:         "123-123-123",
		Age:         59,
		Temperature: []int{89, 98, 98},
		Dependents: []Human{
			{
				Name:        "def",
				SSN:         "",
				Age:         1,
				Temperature: []int{98, 98, 98},
			},
		},
	}
	fmt.Println(medicalDir)
}
