package repository

import (
	"svi-be/internal/model"

	"gorm.io/gorm"
)

type ArtikelRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArtikelRepository {
	return &ArtikelRepository{DB: db}
}

func (r *ArtikelRepository) SaveArticle(artikel *model.Posts) error {
	return r.DB.Create(artikel).Error
}

func (r *ArtikelRepository) GetAll(page int, perPage int, search *model.Search) ([]model.Posts, error) {
	var artikel []model.Posts
	// offset := (page - 1) * perPage
	if search.Status != "" {
		err := r.DB.Table("posts as a").
			Select("a.*").
			Where("a.status ILIKE ?", "%"+search.Status+"%").
			Where("a.deleted_at IS NULL").
			Find(&artikel).Error
		return artikel, err
	}
	err := r.DB.Where("title ILIKE ?", "%"+search.Status+"%").Order("created_at DESC").Find(&artikel).Error
	return artikel, err
}

func (r *ArtikelRepository) GetDetail(id uint) (model.Posts, error) {
	var artikel model.Posts
	err := r.DB.Where("id = ?", id).First(&artikel).Error
	return artikel, err
}

func (r *ArtikelRepository) GetPrev(id uint) (model.Posts, error) {
	var artikel model.Posts
	err := r.DB.Where("id = ?", (id - 1)).First(&artikel).Error
	return artikel, err
}

func (r *ArtikelRepository) GetNext(id uint) (model.Posts, error) {
	var artikel model.Posts
	err := r.DB.Where("id = ?", (id + 1)).First(&artikel).Error
	return artikel, err
}

func (r *ArtikelRepository) UpdateArticle(p *model.Posts) error {
	return r.DB.Save(&p).Error
}

func (r *ArtikelRepository) DeleteArticle(id uint) error {
	return r.DB.Delete(&model.Posts{}, id).Error
}
