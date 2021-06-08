package controller

import (
	"Assignemnts/APIs/models"
	"Assignemnts/APIs/repo"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	//"github.com/lucky-786/gorm-api/models"
	//"github.com/lucky-786/gorm-api/repo"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var validate *validator.Validate
var DB *gorm.DB

func IsvalidID() int64 {
	var u []models.User //array of user structure
	b := repo.DB.Find(&u).RowsAffected
	fmt.Println("value of b", b)
	return b
}

func FieldValidation(user models.User) error {
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}

func PassHash(u string) string {
	h := sha1.New()
	h.Write([]byte(u))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	fmt.Println("Create user called")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.CreatedAt = time.Now().Unix()
	user.Password = PassHash(user.Password)

	err = FieldValidation(user)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		fmt.Fprintf(w, "Json Body is not correct")
		return
	}
	repo.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("single user")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("id is", id)
	var user models.User
	repo.DB.Find(&user, id)
	x := IsvalidID()
	fmt.Println("value of user.Id", user.Id)
	if user.Id == 0 || user.Id > int(x) {
		fmt.Fprintf(w, "No such user exist")
		return
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func GetMultiUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ids models.Result //only one field -> Ids []int `json:ids`
	var user []models.User
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		//v := structs.Values(ids) //convert type struct to interface
		//fmt.Println(v, reflect.TypeOf(v))
		fmt.Println(ids.Ids)
		repo.DB.Where(ids.Ids).Find(&user)
		json.NewEncoder(w).Encode(user)
	}
	/*var users []int
	DB.Table("users").Select("id").Scan(&users)
	if (len(users)) > 0 {
		id := make(map[string][]int)
		id["ids"] = users
		json.NewEncoder(w).Encode(id)
	} else {
		fmt.Fprintf(w, "No data is present in database")
	}*/
}
