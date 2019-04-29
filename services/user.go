package services

import (
	"encoding/json"
	"workshop-meli/db"
	"workshop-meli/models"
	"workshop-meli/utils"
)

var databaseUsers = db.DBUsers{}

//Servicio Alta Usuario
//Recibe un User y devuelve el usuario cargado o error (User,error)
func AltaUsuario(usuario models.User) (models.User, error) {
	return usuario, databaseUsers.Guardar(int(usuario.ID), usuario)
}

func MostrarTodosUsuarios() []models.User {
	result := []models.User{}
	for _, usuario := range databaseUsers.MostrarTodos() {
		result = append(result, usuario)
	}
	return result
}

func MostrarUsuarioPorID(key int) (interface{}, error) {
	if usuario, error := databaseUsers.MostrarPorID(key); error != nil {
		return nil, error
	} else {
		return usuario, nil
	}
}

func ActualizarUsuario(key int, usuario models.User) (interface{}, error) {
	usuarioGuardado, err := databaseUsers.MostrarPorID(key)
	if err != nil {
		return nil, err
	}
	usuarioActual := models.User{}
	if err := json.Unmarshal(utils.InterfaceToBytes(usuarioGuardado), &usuarioActual); err != nil {
		return nil, err
	}
	if usuario.DNI != usuarioActual.DNI && usuario.DNI > 0 {
		usuarioActual.DNI = usuario.DNI
	}
	if usuario.LastName != usuarioActual.LastName && usuario.LastName != "" {
		usuarioActual.LastName = usuario.LastName
	}
	if usuario.Name != usuarioActual.Name && usuario.Name != "" {
		usuarioActual.Name = usuario.Name
	}

	return databaseUsers.Actualizar(key, usuarioActual)
}

func EliminarUsuario(key int) string {
	return databaseUsers.Eliminar(key)
}
