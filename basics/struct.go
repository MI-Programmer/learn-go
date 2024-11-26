package main

// type Customer struct {
// 	Name, Address string
// 	Age           int
// }

// func (customer Customer) sayHello(name string) {
// 	fmt.Println("Hello", name, ", My name is", customer.Name)
// }

// func main() {
// 	var ilham Customer
// 	fmt.Println(ilham)

// 	ilham.Name = "Ilham"
// 	ilham.Address = "Indonesian"
// 	ilham.Age = 30

// 	fmt.Println(ilham)
// 	fmt.Println(ilham.Name)
// 	fmt.Println(ilham.Address)
// 	fmt.Println(ilham.Age)

// 	mipro := Customer{
// 		Name:    "MI Programmer",
// 		Address: "Indonesian",
// 		Age:     20,
// 	}

// 	fmt.Println(mipro)

// 	ferry := Customer{"Ferry Irwandi", "Jawa timur", 30}

// 	fmt.Println(ferry)

// 	ilham.sayHello("Peter")
// 	mipro.sayHello("Ferry")
// 	ferry.sayHello("Deddy")
// }

// type Student struct {
// 	Name, Address, Hobby string
// 	Age                  int
// 	IsPass               bool
// }

// func (s Student) introduce(name string) (result string) {
// 	result = fmt.Sprintf("Hello %s, my name is %s and I am %d years old.", name, s.Name, s.Age)
// 	return result
// }

// func main() {
// 	var peter = Student{
// 		Name:    "Peter",
// 		Address: "American",
// 		Hobby:   "Game",
// 		Age:     20,
// 	}

// 	fmt.Println(peter)

// 	ilham := Student{
// 		"Ilham",
// 		"Indonesian",
// 		"Sing",
// 		20,
// 		true,
// 	}

// 	fmt.Println(ilham)
// 	fmt.Println(ilham.IsPass)
// 	fmt.Println(peter.introduce("David"))
// 	fmt.Println(ilham.introduce("Coki"))
// }

// type Worker struct {
// 	FirstName, LastName string
// 	Salary              int
// }

// func (w Worker) getFullName() string {
// 	return fmt.Sprintf("%s %s", w.FirstName, w.LastName)
// }

// func main() {
// 	david := Worker{
// 		FirstName: "David",
// 		LastName:  "Divad",
// 		Salary:    200,
// 	}

// 	fmt.Println(david)
// 	fmt.Println(david.getFullName())
// }
