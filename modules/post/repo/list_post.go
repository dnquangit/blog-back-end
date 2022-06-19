package postrepo

import (
	"go-module/component"
	models "go-module/modules/post/model"
)

func (repo *postRepo) FindPostsWithCondition(f interface{}) (*[]models.Post, error) {
	switch f.(type) {
	case *models.FilterTitle:
		return repo.FindPostsByTitle(models.NewFilterTitle(f))
	case *models.FilterPublished:
		return repo.FindPostsByPublished(models.NewFilterPublished(f))
	case *models.FilterAll:
		return repo.FindAllPosts(models.NewFilterAll(f))
	}
	return nil, nil
}

func (repo *postRepo) FindPostsByPublished(filter *models.FilterPublished) (*[]models.Post, error) {

	if filter != nil {
		var posts []models.Post
		db := repo.db
		if err := db.Table(models.Post{}.TableName()).Where("published = ? AND deleted = ?", filter.Published, false).Find(&posts).Error; err != nil {
			return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
		}
		return &posts, nil
	}

	return nil, nil
}

func (repo *postRepo) FindPostsByTitle(filter *models.FilterTitle) (*[]models.Post, error) {

	if filter != nil {
		var posts []models.Post
		db := repo.db
		if err := db.Table(models.Post{}.TableName()).Where("title = ? AND id != ? AND deleted = ?", filter.Title, filter.Id, false).Find(&posts).Error; err != nil {
			return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
		}
		return &posts, nil
	}

	return nil, nil
}

func (repo *postRepo) FindAllPosts(filter *models.FilterAll) (*[]models.Post, error) {

	if filter != nil {
		var posts []models.Post
		db := repo.db
		if err := db.Table(models.Post{}.TableName()).Where("deleted = ?", false).Find(&posts).Error; err != nil {
			return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
		}
		return &posts, nil
	}

	return nil, nil
}
