package api

import (
	"net/http"

	"github.com/Fonzeca/Chatline/src/db/model"
	"github.com/Fonzeca/Chatline/src/server/manager"
	"github.com/Fonzeca/Chatline/src/services"
	"github.com/labstack/echo/v4"
)

type ApiComment struct {
	manager manager.CommentManager
}

func NewApiComment() ApiComment {
	m := manager.NewCommentManager()
	return ApiComment{manager: m}
}

func (api *ApiComment) CreateComment(c echo.Context) error {
	data := model.Comentario{}
	c.Bind(&data)

	err := api.manager.CreateComment(data)

	if err != nil {
		return err
	}

	go services.ProcessData(model.ComentarioMQ{
		UsuarioID: data.UsuarioID,
		Fecha:     data.Fecha,
		Tema:      data.Tema,
		TemaId:    data.TemaID,
		Mensaje:   data.Mensaje,
	})

	return c.NoContent(http.StatusOK)
}

func (api *ApiComment) GetAllComments(c echo.Context) error {
	comments, err := api.manager.GetAllComments()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Se produjo un error al consultar la bdd")
	}

	return c.JSON(http.StatusOK, comments)
}

func (api *ApiComment) GetCommentsByUserIds(c echo.Context) error {
	data := model.UserIds{}
	c.Bind(&data)

	comments, err := api.manager.GetCommentsByUserIds(data.Ids)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Se produjo un error al consultar la bdd")
	}

	return c.JSON(http.StatusOK, comments)
}
