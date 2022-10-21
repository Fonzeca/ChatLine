package api

import (
	"net/http"

	"github.com/Fonzeca/Chatline/src/db/model"
	"github.com/Fonzeca/Chatline/src/server/manager"
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

func (api *ApiComment) GetCommentsByTopicAndTopicId(c echo.Context) error {
	val, paramErr := c.FormParams()
	tema := val.Get("tema")
	temaId := val.Get("tema_id")

	if paramErr != nil {
		return c.JSON(http.StatusBadRequest, paramErr.Error())
	}

	if tema == "" || temaId == "" {
		return c.JSON(http.StatusBadRequest, "Ningún parametro puede estar vacío")
	}

	comments, err := api.manager.GetCommentsByTopicAndTopicId(tema, temaId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Se produjo un error al consultar la bdd")
	}

	return c.JSON(http.StatusOK, comments)
}
