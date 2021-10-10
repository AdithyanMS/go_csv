package main

import (
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// type tableData struct {
// 	Sno      int
// 	Categ    string
// 	ProdName string
// 	Descp    string
// 	Mrp      int
// }

func main() {

	csvFile, err := os.Open("table.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	// for _, line := range csvLines {
	//     emp := empData{
	//         Name: line[0],
	//         Age: line[1],
	// 		City: line[2],
	//     }
	//     fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
	// }
	fmt.Println(csvLines[1][3])
}
