package main

// type HasName interface {
// 	GetName() string
// }

// type Person struct {
// 	Name string
// }

// func (p Person) GetName() string {
// 	return p.Name
// }

// type Animal struct {
// 	Name string
// }

// func (a Animal) GetName() string {
// 	return a.Name
// }

// func sayHello(value HasName) {
// 	fmt.Println("Hello", value.GetName())
// }

// func main() {
// 	ilham := Person{"MI Programmer"}
// 	sayHello(ilham)

// 	kucing := Animal{Name: "Kucing indo"}
// 	sayHello(kucing)
// }

// type Person interface {
// 	GetFullName() string
// 	Greeting(string) string
// }

// type Student struct {
// 	FirstName, LastName, School string
// 	Age                         int
// }

// func (s Student) GetFullName() string {
// 	return fmt.Sprintf("%s %s", s.FirstName, s.LastName)
// }

// func (s Student) Greeting(name string) string {
// 	return fmt.Sprintf("Hello %s, where you from? - %s", name, s.GetFullName())
// }

// type Worker struct {
// 	FirstName, LastName, Firm string
// 	Age                       int
// }

// func (w Worker) GetFullName() string {
// 	return fmt.Sprintf("%s %s", w.FirstName, w.LastName)
// }

// func (w Worker) Greeting(name string) string {
// 	return fmt.Sprintf("Hello %s, where you from? - %s", name, w.GetFullName())
// }

// func fullname(value Person) {
// 	fmt.Println("Hi my fullname is ", value.GetFullName())
// }

// func main() {
// 	ferry := Student{
// 		FirstName: "Ferry",
// 		LastName:  "Irwandi",
// 		School:    "STAN",
// 		Age:       33,
// 	}

// 	ilham := Worker{
// 		FirstName: "MI",
// 		LastName:  "Programmer",
// 		Firm:      "Rakamin",
// 		Age:       20,
// 	}

// 	fullname(ferry)
// 	fullname(ilham)
// }

// type Car interface {
// 	IsOld() bool
// 	GetCurrentPrice(int) int
// }

// type Tesla struct {
// 	Name    string
// 	Release int
// 	Price   int
// }

// func (t Tesla) IsOld() bool {
// 	return t.Release < 2010
// }

// func (t Tesla) GetCurrentPrice(year int) int {
// 	return t.Price - year/5
// }

// func isRecommend(value Car) {
// 	if value.IsOld() {
// 		fmt.Println("Is not recommended")
// 	} else {
// 		fmt.Println("Is recommended")
// 	}
// }

// func main() {
// 	x := Tesla{Name: "X", Price: 200, Release: 2020}

// 	isRecommend(x)
// }
