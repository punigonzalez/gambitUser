package bd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/punigonzalez/gambitUser/models"

	"fmt"

	"github.com/punigonzalez/gambitUser/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO users (User_Uuid, User_Email, User_DateAdd) VALUES ('" + sign.UserUUID + "','" + sign.UserEmail + "','" + tools.FechaMySQL() + "')"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecucion Exitosa")
	return nil
}
