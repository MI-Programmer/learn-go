package main

// var (
// 	ValidationError = errors.New("Validation error")
// 	NotFoundError   = errors.New("Not found error")
// )

// func GetById(id string) error {
// 	if id == "" {
// 		return ValidationError
// 	}
// 	if id != "ilham" {
// 		return NotFoundError
// 	}
// 	return nil
// }

// func main() {
// 	err := GetById("David")

// 	if errors.Is(err, ValidationError) {
// 		fmt.Println("Validation error")
// 	} else if errors.Is(err, NotFoundError) {
// 		fmt.Println("Not found error")
// 	} else {
// 		fmt.Println("Unknown error")
// 	}
// }

// var (
// 	ClientError     = errors.New("Client error")
// 	BadRequestError = errors.New("Bad request error")
// )

// func Request(path string) error {
// 	if path == "" {
// 		return ClientError
// 	}
// 	if path != "/home" {
// 		return BadRequestError
// 	}
// 	return nil
// }

// func main() {
// 	err := Request("")

// 	if errors.Is(err, ClientError) {
// 		fmt.Println(401)
// 	} else if errors.Is(err, BadRequestError) {
// 		fmt.Println(400)
// 	} else {
// 		fmt.Println(500)
// 	}
// }
