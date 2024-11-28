package main

// import (
// 	"encoding/csv"
// 	"os"
// )

// func main() {
// 	csvString := "mi,programmer,ganteng\n" + "eko,kurni,khan\n" + "david,peter,parker"

// 	reader := csv.NewReader(strings.NewReader(csvString))

// 	for {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}

// 		fmt.Println(record)
// 	}
// }

// func main() {
// 	writer := csv.NewWriter(os.Stdout)

// 	_ = writer.Write([]string{"mi", "programmer", "ganteng"})
// 	_ = writer.Write([]string{"mi", "programmer", "ganteng"})
// 	_ = writer.Write([]string{"mi", "programmer", "ganteng"})

// 	writer.Flush()
// }
