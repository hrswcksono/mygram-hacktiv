package photo_repository

type PhotoRepository interface {
	CreatePhoto()
	GetAllPhoto()
	UpdatePhoto(photoId int)
	DeletePhoto(photoId int)
}
