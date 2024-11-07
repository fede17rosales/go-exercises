package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"twittergo/database"
	"twittergo/models"
)

func Register(ctx context.Context) models.RespAPI {
	var t models.User
	var r models.RespAPI
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el Email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, find, _ := database.CheckUser(t.Email)
	if find {
		r.Message = "Ya existe un usuario registrado con ese mail"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := database.InsertRegister(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println(r.Message)
	return r
}
