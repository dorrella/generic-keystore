// Terrible database using generic keystore
package main

import (
	"fmt"

	keystore "github.com/dorrella/generic-keystore"
)

// person table
type Person struct {
	Name string
	Age  uint32
}

// class table
type Class struct {
	Name     string
	Teacher  uint32
	Students []uint32
}

func main() {
	// using uint32s as the primary key, create tables for
	// person and class objects
	store_people := keystore.NewKeyStore[uint32, *Person]()
	store_class := keystore.NewKeyStore[uint32, *Class]()

	// helper to fill table
	people := map[uint32]*Person{
		0: &Person{"bob baker", 80},
		1: &Person{"drew carey", 45},
		2: &Person{"alex trebeck", 64},
	}

	// actually fill people table
	for k, v := range people {
		store_people.Put(k, v)
	}

	// only have 1 class, so make it directly
	// using uids from people map
	c := &Class{
		Name:     "Public Speech",
		Teacher:  0,
		Students: []uint32{1, 2},
	}

	//store class object
	store_class.Put(0, c)

	//....

	//get class object
	new_class, ok := store_class.Get(0)
	if !ok {
		panic("some error")
	}

	//print class name
	fmt.Printf("%s:\n", new_class.Name)

	//print teacher
	teacher, ok := store_people.Get(new_class.Teacher)
	if !ok {
		panic("failed to get teacher")
	}

	fmt.Printf("  Teacher: %s\n", teacher.Name)

	//get each student
	fmt.Println("Students:")
	for _, student_id := range new_class.Students {
		student, ok := store_people.Get(student_id)
		if !ok {
			panic("failed to get student")
		}

		fmt.Printf("  %s, Age: %d\n", student.Name, student.Age)
	}

}
