package comment_pg

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/repository/comment_repository"
	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

func (c *commentPG) CreateComment(commentPayload *entity.Comment) (*entity.Comment, error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&commentPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return commentPayload, tx.Commit().Error
}

func (c *commentPG) GetAllComment() ([]entity.Comment, error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var comment = []entity.Comment{}

	if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.User{}) }).Preload("Photo", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.Photo{}) }).Find(&comment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comment, tx.Commit().Error
}

func (c *commentPG) UpdateComment(commentPayload *entity.Comment) (*entity.Comment, error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Updates(commentPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return commentPayload, tx.Commit().Error
}

func (c *commentPG) DeleteComment(commentId int) error {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var comment = &entity.Comment{}

	if err := tx.Delete(&comment, commentId).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (c *commentPG) GetCommentByID(commentId int) (*entity.Comment, error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var comment = &entity.Comment{}

	if err := tx.Delete(&comment, commentId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comment, tx.Commit().Error
}
