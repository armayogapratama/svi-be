package handler

import (
	"net/http"
	"strconv"
	"svi-be/internal/model"
	"svi-be/internal/service"
	"svi-be/internal/validation"

	"github.com/gin-gonic/gin"
)

type ArtikelHandler struct {
	Service *service.ArticleService
}

func NewArticleHandler(svc *service.ArticleService) *ArtikelHandler {
	return &ArtikelHandler{
		Service: svc,
	}
}

func (h *ArtikelHandler) CreateArticle(c *gin.Context) {
	var p model.Posts
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: "Request body tidak sesuai",
			Data:    err.Error(),
		})
		return
	}

	isValid, validationMessage := validation.ValidatePost(&p)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": validationMessage,
			"data":    "",
		})
		return
	}

	result, err := h.Service.CreateArticle(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ArtikelHandler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "10")
	perPage, _ := strconv.Atoi(perPageStr)

	status := c.DefaultQuery("status", "")

	var search model.Search

	search.Status = status

	result, err := h.Service.GetAll(page, perPage, &search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ArtikelHandler) GetDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	result, err := h.Service.GetDetail(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ArtikelHandler) DeleteArtikel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: "Invalid ID format",
			Data:    err.Error(),
		})
		return
	}

	detail, err := h.Service.GetDetail(uint(id))
	if err != nil || detail.Data == nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  "Error",
			Message: "Post not found",
			Data:    "",
		})
		return
	}

	dataMap, ok := detail.Data.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  "Error",
			Message: "Failed to parse data",
			Data:    "",
		})
		return
	}

	if dataMap["id"] == nil || dataMap["id"].(uint) == 0 {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  "Error",
			Message: "Post not found",
			Data:    "",
		})
		return
	}

	currentStatus := dataMap["status"].(string)
	if currentStatus != "publish" && currentStatus != "draft" {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: "Only posts with 'Publish' or 'Draft' status can be marked as 'Thrash'",
			Data:    "",
		})
		return
	}

	var p model.Posts
	p.ID = uint(id)
	p.Status = "thrash"
	p.Title = dataMap["title"].(string)
	p.Content = dataMap["content"].(string)
	p.Category = dataMap["category"].(string)

	result, err := h.Service.UpdateArtikel(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ArtikelHandler) UpdateArtikel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: "Invalid ID format",
			Data:    err.Error(),
		})
		return
	}

	var p model.Posts
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: "Request body tidak sesuai",
			Data:    err.Error(),
		})
		return
	}

	p.ID = uint(id)

	detail, err := h.Service.GetDetail(uint(id))
	if err != nil || detail.Data == nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  "Error",
			Message: "Post not found",
			Data:    "",
		})
		return
	}

	dataMap, ok := detail.Data.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  "Error",
			Message: "Failed to parse data",
			Data:    "",
		})
		return
	}

	if dataMap["id"] == nil || dataMap["id"].(uint) == 0 {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  "Error",
			Message: "Post not found",
			Data:    "",
		})
		return
	}

	isValid, validationMessage := validation.ValidatePost(&p)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": validationMessage,
			"data":    "",
		})
		return
	}

	result, err := h.Service.UpdateArtikel(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
