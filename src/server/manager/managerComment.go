package manager

import (
	"github.com/Fonzeca/Chatline/src/db"
	"github.com/Fonzeca/Chatline/src/db/model"
)

type CommentManager struct {
}

func NewCommentManager() CommentManager {
	return CommentManager{}
}

func (ma *CommentManager) CreateComment(commentRequest model.Comentario) error {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return err
	}

	tx := db.Create(&commentRequest)

	return tx.Error
}

func (ma *CommentManager) GetAllComments() ([]model.Comentario, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.Comentario{}, err
	}

	comments := []model.Comentario{}
	tx := db.Find(&comments)

	return comments, tx.Error
}

func (ma *CommentManager) GetCommentsByUserIds(userIdsRequest []string) ([]model.Comentario, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.Comentario{}, err
	}

	comments := []model.Comentario{}
	tx := db.Where("usuario_id IN ?", userIdsRequest).Find(&comments)

	return comments, tx.Error
}
