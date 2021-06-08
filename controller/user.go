package controller

import (
	//"Assignemnts/APIs/models"
	//"Assignemnts/APIs/repo"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"net/http"

	"github.com/lucky-786/gorm-api/models"
	"github.com/lucky-786/gorm-api/repo"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var validate *validator.Validate
var DB *gorm.DB

func IsvalidID() int64 {
	var u []models.User //array of user structure
	b := repo.DB.Find(&u).RowsAffected
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
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	repo.DB.Find(&user, id)
	x := IsvalidID()
	if user.Id == 0 || user.Id > int(x) {
		fmt.Fprintf(w, "User with id "+id+" does not exist")
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
		for _, i := range ids.Ids {
			fmt.Print(i)
			repo.DB.Find(&user, i)
			x := IsvalidID()
			if i == 0 || i > int(x) {
				fmt.Fprintln(w, "User with id "+strconv.Itoa(i)+" does not exist")
			} else {
				json.NewEncoder(w).Encode(user)
			}
		}
	}
}
