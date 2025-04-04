package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/punigonzalez/gambitUser/models"
	"github.com/punigonzalez/gambitUser/secretM"
)

var SecretModel models.SecretRDSjson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretM.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion exitosa a la BD")
	return nil
}

func ConnStr(keys models.SecretRDSjson) string {
	var dbUser, authoken, dbEndpoint, dbname string

	dbUser = keys.Username
	authoken = keys.Password
	dbEndpoint = keys.Host
	dbname = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authoken, dbEndpoint, dbname)

	fmt.Println(dsn) // para mostrar en cloudWatch
	return dsn
}
