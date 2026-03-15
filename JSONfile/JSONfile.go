package main

import (
	"encoding/json"
	"fmt"
)

type Jstruct struct {
	Id    string
	Name  string
	Total float64
}

func startJSON() {
	fmt.Println("welcome to json covertion")
	jstruct := Jstruct{}
	var JData []byte
	AssinData(&jstruct)
	ConvertToJSON(&jstruct, &JData)
	ConvertToStruct(&JData)

}

func ConvertToStruct(jdata *[]byte) {
	Struct2 := Jstruct{}
	fmt.Println("-------------Converting JSON to Struct -------------------------")
	fmt.Println("Your entered data in the JSON format:")
	fmt.Printf("%s\n", *jdata)
	err := json.Unmarshal(*jdata, &Struct2)
	if err != nil {
		fmt.Println("error in the data or value entered:", err)
	}
	fmt.Println("Your entered data in the Struct format:")
	fmt.Println(Struct2.Name, Struct2.Id, Struct2.Total)
}
func ConvertToJSON(jsn *Jstruct, jdata *[]byte) {
	fmt.Println("-------------COnverting Struct to JSON -------------------------")
	fmt.Println("Your entered data in the struct format:")
	fmt.Println(jsn.Name, jsn.Id, jsn.Total)
	fmt.Println("Your entered data in the JSON format:")
	jsonData, err := json.Marshal(jsn)
	if err != nil {
		fmt.Println("error in the data or value entered:", err)
	}
	*jdata = append(*jdata, jsonData...)
	fmt.Printf("%s\n", jsonData)

}

func AssinData(jsn *Jstruct) {
	var name, id string
	var total float64
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("enter your Student id: ")
	fmt.Scanln(&id)
	fmt.Print("Enter student total marks: ")
	fmt.Scanln(&total)
	jsn.Name = name
	jsn.Id = id
	jsn.Total = total
}

func main() {
	startJSON()
}

//REal developer style of writing
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	ID    string  `json:"id"`
// 	Name  string  `json:"name"`
// 	Total float64 `json:"total"`
// }

// func assignData() Student {
// 	var s Student

// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&s.Name)

// 	fmt.Print("Enter your student id: ")
// 	fmt.Scanln(&s.ID)

// 	fmt.Print("Enter total marks: ")
// 	fmt.Scanln(&s.Total)

// 	return s
// }

// func convertToJSON(s Student) ([]byte, error) {
// 	fmt.Println("\nStruct Data:")
// 	fmt.Println(s.Name, s.ID, s.Total)

// 	jsonData, err := json.Marshal(s)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("\nJSON Format:")
// 	fmt.Println(string(jsonData))

// 	return jsonData, nil
// }

// func convertToStruct(data []byte) (Student, error) {
// 	var s Student

// 	err := json.Unmarshal(data, &s)
// 	if err != nil {
// 		return s, err
// 	}

// 	return s, nil
// }

// func main() {

// 	fmt.Println("Welcome to JSON Conversion Program")

// 	student := assignData()

// 	jsonData, err := convertToJSON(student)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	newStudent, err := convertToStruct(jsonData)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("\nConverted Back To Struct:")
// 	fmt.Println(newStudent.Name, newStudent.ID, newStudent.Total)
// }
