package controllers

import (
	"check-domain-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	oneCount int = 0
	minCount int = 0
	maxCount int = 0
)

type Domain interface {
	Save(models.Domain) error
	GetByName(name string) (models.Domain, error)
	GetMin() (models.Domain, error)
	GetMax() (models.Domain, error)
}

type Handler struct {
	domainService Domain
}

func NewHandler(domain Domain) *Handler {
	return &Handler{
		domainService: domain,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/site", h.getByName)
	r.GET("/min", h.getMin)
	r.GET("/max", h.getMax)
	r.GET("/stat", h.getStat)

	return r

}

func (h *Handler) getByName(c *gin.Context) {
	oneCount++
	name := c.Query("name")
	domain, err := h.domainService.GetByName(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, domain)

}

func (h *Handler) getMin(c *gin.Context) {
	minCount++
	domain, err := h.domainService.GetMin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, domain)

}
func (h *Handler) getMax(c *gin.Context) {
	maxCount++
	domain, err := h.domainService.GetMax()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, domain)

}

func (h *Handler) getStat(c *gin.Context) {
	c.String(http.StatusOK, "Кол-во запросов к конкретному сайту: %d\n", oneCount)
	c.String(http.StatusOK, "Кол-во запросов к сайтам с минимальной задержкой: %d\n", minCount)
	c.String(http.StatusOK, "Кол-во запросов к сайтам с максимальной задержкой: %d\n", maxCount)
	// 	c.JSON(200, gin.H{"Кол-во запросов к конкретному сайту": oneCount})
	// 	c.JSON(200, gin.H{"Кол-во запросов к сайтам с минимальной задержкой": minCount})
	// 	c.JSON(200, gin.H{"Кол-во запросов к сайтам с максимальной задержкой": maxCount})
}
