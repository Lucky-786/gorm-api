/*package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Users struct {
	ID        string `validae:"omitempty,uuid"`
	FirstName string `validate:"required"`
	LastName  string `validate:"alpha,required"`
}

func main() {
	user := Users{
		FirstName: "Lucy",
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(user)
	}
}*/

/*"first_name":"15asf",
  "last_name":"15vsf",
  "city":"Hydrabad",
  "phone":8681921912,
  "height":6,
  "gender":"M",
  "password":"12dsdswe",
  "married":"false"*/

/*var users []int    //multi user
DB.Table("users").Select("id").Scan(&users)
if (len(users)) > 0 {
	id := make(map[string][]int)
	id["ids"] = users
	json.NewEncoder(w).Encode(id)
} else {
	fmt.Fprintf(w, "No data is present in database")
}*/