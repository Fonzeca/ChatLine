package manager

import (
	"strconv"
	"time"

	"github.com/Fonzeca/Chatline/src/db"
	"github.com/Fonzeca/Chatline/src/db/model"
	"github.com/Fonzeca/Chatline/src/services"
	"gorm.io/gorm"
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

	transactionErr := db.Transaction(func(tx *gorm.DB) error {

		newComment := model.ComentarioView{
			UsuarioID: commentRequest.UsuarioID,
			Fecha:     time.Now(),
			Tema:      commentRequest.Tema,
			TemaId:    commentRequest.TemaID,
			Mensaje:   commentRequest.Mensaje,
		}

		if err := tx.Create(&commentRequest).Error; err != nil {
			return err
		}

		rabbitErr := services.ProcessData(newComment)

		return rabbitErr
	})

	return transactionErr
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

func (ma *CommentManager) GetCommentsByTopicAndTopicId(topic string, topicId string) ([]model.Comentario, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.Comentario{}, err
	}

	topicIdInt, convErr := strconv.Atoi(topicId)

	if convErr != nil {
		return []model.Comentario{}, convErr
	}

	comments := []model.Comentario{}
	tx := db.Where("tema = ? AND tema_id = ?", topic, topicIdInt).Find(&comments)

	return comments, tx.Error
}
