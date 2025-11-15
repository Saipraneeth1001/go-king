package custommarshall


import (
	"goking/models"
	"encoding/json"
	"fmt"

)

func CustomMarshal() {

	u := models.User{Username: "Praneeth", Password: "King"}
	fmt.Println("Printing without marshal", u)

	output, _ := json.Marshal(u)
	fmt.Printf("structs: %s", output)
}