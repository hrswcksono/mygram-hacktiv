package comment_pg

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
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

func (c *commentPG) CreateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Create(&commentPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return commentPayload, nil
}

func (c *commentPG) GetAllComment() ([]entity.Comment, errs.MessageErr) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var comment = []entity.Comment{}

	if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.User{}) }).Preload("Photo", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.Photo{}) }).Find(&comment).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return comment, nil
}

func (c *commentPG) UpdateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Updates(commentPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var comment = &entity.Comment{}

	if err := tx.Preload("Photo", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.Photo{}) }).Find(&comment, commentPayload.ID).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return comment, nil
}

func (c *commentPG) DeleteComment(commentId int) errs.MessageErr {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	var comment = &entity.Comment{}

	if err := tx.Delete(&comment, commentId).Error; err != nil {
		tx.Rollback()
		return errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	return nil
}

func (c *commentPG) GetCommentByID(commentId int) (*entity.Comment, errs.MessageErr) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var comment = &entity.Comment{}

	if err := tx.Find(&comment, commentId).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return comment, nil
}
