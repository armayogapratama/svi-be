package service

import (
	"svi-be/internal/model"
	"svi-be/internal/repository"
)

type ArticleService struct {
	Repo *repository.ArtikelRepository
}

func NewArticleService(repo *repository.ArtikelRepository) *ArticleService {
	return &ArticleService{Repo: repo}
}

func (s *ArticleService) GetAll(page int, perPage int, search *model.Search) (*model.Response, error) {
	result, err := s.Repo.GetAll(page, perPage, search)
	if err != nil {
		return &model.Response{
			Status:  "Error",
			Message: "Gagal mengambil data",
			Data:    err.Error(),
		}, err
	}

	var data []map[string]interface{}

	for i := 0; i <= len(result)-1; i++ {
		data = append(data, map[string]interface{}{
			"id":       result[i].ID,
			"title":    result[i].Title,
			"content":  result[i].Content,
			"category": result[i].Category,
			"status":   result[i].Status,
		})
	}

	return &model.Response{
		Status:  "Success",
		Message: "Berhasil ambil data",
		Data:    data,
	}, nil
}

func (s *ArticleService) CreateArticle(p *model.Posts) (*model.Response, error) {
	err := s.Repo.SaveArticle(p)
	if err != nil {
		return &model.Response{
			Status:  "Error",
			Message: "Gagal membuat artikel",
			Data:    err.Error(),
		}, err
	}

	return &model.Response{
		Status:  "Success",
		Message: "Berhasil membuat artikel",
		Data: map[string]interface{}{
			"id":       p.ID,
			"title":    p.Title,
			"content":  p.Content,
			"category": p.Category,
			"status":   p.Status,
		},
	}, nil
}

func (s *ArticleService) GetDetail(id uint) (*model.Response, error) {
	result, err := s.Repo.GetDetail(id)
	if err != nil {
		return &model.Response{
			Status:  "Error",
			Message: "Gagal mengambil detail",
			Data:    err.Error(),
		}, err
	}

	return &model.Response{
		Status:  "Success",
		Message: "Berhasil mengambil detail",
		Data: map[string]interface{}{
			"id":       result.ID,
			"title":    result.Title,
			"content":  result.Content,
			"category": result.Category,
			"status":   result.Status,
		},
	}, nil
}

func (s *ArticleService) UpdateArtikel(p *model.Posts) (*model.Response, error) {
	err := s.Repo.UpdateArticle(p)
	if err != nil {
		return &model.Response{
			Status:  "Error",
			Message: "Gagal mengubah data",
			Data:    err.Error(),
		}, err
	}

	return &model.Response{
		Status:  "Success",
		Message: "Berhasil mengubah data",
		Data:    p,
	}, nil
}

func (s *ArticleService) DeleteArtikel(id uint) (*model.Response, error) {
	err := s.Repo.DeleteArticle(id)
	if err != nil {
		return &model.Response{
			Status:  "Error",
			Message: "Gagal menghapus data",
			Data:    err.Error(),
		}, err
	}

	return &model.Response{
		Status:  "Success",
		Message: "Berhasil menghapus data",
		Data:    id,
	}, nil
}
