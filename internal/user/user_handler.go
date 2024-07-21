package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s service) Handler {
	return Handler{
		service: s,
	}
}

func (h Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, CreateUserRes{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	})
}

func (h Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", res.Token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, LoginRes{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	})
}

func (h Handler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
