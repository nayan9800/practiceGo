package main

import "fmt"

type person struct {
	Name      string
	ContactNo int
}

/*methods are defined on  struct*/
/*here ToString method is defined on person struct*/
func (p person) Tostring() string { return fmt.Sprintf("%s has contact no %d", p.Name, p.ContactNo) }

/*methods can be defined as value receiver or pointer receiver*/
/*here ChangeContactNo method reciveing pointer and changing ContactNo field*/
func (p *person) ChangeContactNo(newNo int) { p.ContactNo = newNo }

/*methods can also be declared on non-struct types*/
/*myFloat is float64 type*/
type myFloat float64

/*abs method is defined on myFolat*/
func (f myFloat) abs() myFloat {
	if f < 0 {
		return -f
	}
	return f
}

/*interfaces in go*/
//An interface type is defined as a set of method signatures.
type college interface {
	showDetails()
}

type student struct {
	Name   string
	rollNo int
	class  int
}
type teacher struct {
	Name string
	ID   int
}

//here student and teacher types inpleamenting
//interface college by impleamenting showDeatails()
//method
func (s student) showDetails() {
	fmt.Printf("%s is student,roll no:-%d in class %d\n", s.Name, s.rollNo, s.class)
}

func (t teacher) showDetails() {
	fmt.Printf("%s is Teacher in colleage having ID:-%d\n", t.Name, t.ID)
}

//display function takes multiple college
//types as argument
/*college interface is used as value in here*/
func display(people ...college) {
	for _, p := range people {
		p.showDetails()
	}
}
func MethodsAndInterfaces() {
	fmt.Println("methods and Interfaces in go")

	foo := person{Name: "foo", ContactNo: 8555647551}
	fmt.Println(foo.Tostring())
	foo.ChangeContactNo(455785351)
	fmt.Println(foo.Tostring())

	myf := myFloat(-14545.5454)
	fmt.Println(myf.abs())

	alice := student{Name: "alice", rollNo: 4, class: 5}
	bob := teacher{Name: "bob", ID: 5}
	display(alice, bob)
}
