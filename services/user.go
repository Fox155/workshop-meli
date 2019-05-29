package services

import (
	"encoding/json"
	"workshop-meli/models"
	"workshop-meli/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbmysql, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/prueba_go")

if err != nil {
	panic(err.Error())
}

defer dbmysql.Close()

//Servicio Alta Usuario
//Recibe un User y devuelve el usuario cargado o error (User,error)
func AltaUsuario(usuario models.User) (models.User, error) {
	insert, err := dbmysql.Query("INSERT Users VALUES(?,?,?,?);",usuario.ID,usuario.DNI,usuario.LastName,usuario.Name)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	defer insert.Close()
	return usuario, nil
}

func MostrarTodosUsuarios() []models.User {
	result := []models.User{}

	results_q, err := dbmysql.Query("SELECT * FROM Users;")

	if err != nil {
		panic(err.Error())
		return nil
	}

	for results_q.Next(){
		var usuario User

		err = results_q.Scan(&usuario.ID,&usuario.DNI,&usuario.LastName,&usuario.Name);

		if err != nil{
			panic(err.Error())
			return nil
		}

		result = append(result, usuario)
	}
	return result
}

func MostrarUsuarioPorID(key int) (models.User, error) {
	result := User{}

	result_q, err := dbmysql.QueryRow("SELECT * FROM Users WHERE user_id=?;",key)

	if err != nil {
		panic(err.Error())
		return nil, err
	}

	err = result_q.Scan(&result.ID,&result.DNI,&result.LastName,&result.Name);

	if err != nil{
		panic(err.Error())
		return nil, err
	}

	return result
}

func ActualizarUsuario(key int, usuario models.User) (models.User, error) {
	usuarioGuardado := User{}

	result_q, err := dbmysql.QueryRow("SELECT * FROM Users WHERE user_id=?;",key)

	if err != nil {
		panic(err.Error())
		return nil, err
	}

	err = result_q.Scan(&usuarioGuardado.ID,&usuarioGuardado.DNI,&usuarioGuardado.LastName,&usuarioGuardado.Name);

	if err != nil{
		panic(err.Error())
		return nil, err
	}
	
	result_q, err := dbmysql.Query("UPDATE Users SET dni=?,last_name=?,name=? WHERE user_id=?;",usuario.DNI,usuario.LastName,usuario.Name,key)

	if err != nil{
		panic(err.Error())
		return nil, err
	}

	return usuario, nil
}

func EliminarUsuario(key int) string {
	result_q, err := dbmysql.QueryRow("DELETE FROM Users WHERE user_id=?;",key)

	if err != nil {
		panic(err.Error())
		return err.Error()
	}

	return nil
}
