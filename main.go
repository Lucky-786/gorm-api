package main

import (
	"Assignemnts/APIs/cmd"
	"Assignemnts/APIs/repo"
	//"github.com/lucky-786/gorm-api/cmd"
	//"github.com/lucky-786/gorm-api/repo"
)

func init() {
	_, err := repo.InitialMigration()
	if err != nil {
		panic("Could not connect Database")
	}

}

func main() {
	cmd.Execute()
}
