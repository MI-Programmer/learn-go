package main

// type validationError struct {
// 	Message string
// }

// func (v *validationError) Error() string {
// 	return v.Message
// }

// type notFoundError struct {
// 	Message string
// }

// func (n *notFoundError) Error() string {
// 	return n.Message
// }

// func SaveData(value string) error {
// 	if value == "" {
// 		return &validationError{"Validation error"}
// 	}
// 	if value != "ilham" {
// 		return &notFoundError{"Not found data"}
// 	}
// 	return nil
// }

// func main() {
// 	err := SaveData("agus")

// 	if err == nil {
// 		fmt.Println("Success")
// 	} else {
// 		switch finalError := err.(type) {
// 		case *validationError:
// 			fmt.Println("Validation error:", finalError.Error())
// 		case *notFoundError:
// 			fmt.Println("Not Found Error:", finalError.Error())
// 		default:
// 			fmt.Println("Unknown Error:", err.Error())
// 		}

// 		if validationError, ok := err.(*validationError); ok {
// 			fmt.Println("Validation error:", validationError.Error())
// 		} else if notFoundError, ok := err.(*notFoundError); ok {
// 			fmt.Println("Not Found Error:", notFoundError.Error())
// 		} else {
// 			fmt.Println("Unknown Error:", err.Error())
// 		}
// 	}
// }

// type serverError struct {
// 	Message string
// }

// func (s *serverError) Error() string {
// 	return s.Message
// }

// type clientError struct {
// 	Message string
// }

// func (c *clientError) Error() string {
// 	return c.Message
// }

// func Request(path string) error {
// 	if path == "" {
// 		return &serverError{"Server error"}
// 	}
// 	if path != "/home" {
// 		return &clientError{"Client error"}
// 	}
// 	return nil
// }

// func main() {
// 	err := Request("")

// 	if err == nil {
// 		fmt.Println("Request success")
// 	} else {
// 		if serverError, ok := err.(*serverError); ok {
// 			fmt.Println("(500):", serverError.Error())
// 		} else if clientError, ok := err.(*clientError); ok {
// 			fmt.Println("(400):", clientError.Error())
// 		} else {
// 			fmt.Println(err.Error())
// 		}
// 	}
// }
