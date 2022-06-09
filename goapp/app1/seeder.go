package main

import (
//	"os"
//	"go-crud-echo-mysql/controller"
	"go-crud-echo-mysql/model"
/*
	"encoding/json"
	"log"
	"fmt"*/
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
/*
	var stcData []*model.Member
	p, _ := os.Getwd()
	fmt.Println(p)

	f, err := os.ReadFile("goapp/app1/n46_members.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(f), &stcData); err != nil {
		fmt.Println(err)
		return err
	}

	for _, memberRow := range stcData {
		p, _ := strconv.Atoi(memberRow.Age)
		age := uint(p)
	
		member := model.Member {
			Name: memberRow.Name,
			Age: age,
			Bloodtype: memberRow.Bloodtype,
			Origin: memberRow.Origin,
		}

		// Insert into SQL
		member.Create()
		fmt.Printf("%s (@%s)\n", memberRow.Name)
	}
*/
	return nil
}
