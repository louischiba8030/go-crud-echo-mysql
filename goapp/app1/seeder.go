package main

import (
	"os"
//	"go-crud-echo-mysql/controller"
	"go-crud-echo-mysql/model"

//	"strconv"
	"encoding/json"
	"log"
	"fmt"
)

// Error handling
//type FStreamReadError struct {
//	Msg string
//	Code int
//}
// Implement Error func on MyErr struct
//func (err *FStreamReadError) Error() string {
//	return fmt.Sprintf("err %s [code=%d]", err.Msg, err.Code)
//}

func main () {
	DoSeed()
}

func DoSeed () error {
	// DROP TABLE members
	m := model.Member{}
	m.DropTable()
	// Create TABLE members
	m.CreateTable()

	// Get current directory
	var stcData []*model.Member
	p, _ := os.Getwd()
	fmt.Println(p)

	// Load local json file
	f, err := os.ReadFile("goapp/app1/n46_members.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read stream from json data
	if err := json.Unmarshal([]byte(f), &stcData); err != nil {
		fmt.Println(err)
		return err
	}

	// Put each row and store into the SQL table
	for _, memberRow := range stcData {
//		p, _ := strconv.Atoi(memberRow.Age)
//		age := uint(p)
	
		member := model.Member {
			Name: memberRow.Name,
			Age: memberRow.Age,
			Bloodtype: memberRow.Bloodtype,
			Origin: memberRow.Origin,
		}

		// Insert into SQL
		member.Create()
		fmt.Printf("Written data: %s %v\n", memberRow.Name, memberRow.Age)
	}

	return nil
}
