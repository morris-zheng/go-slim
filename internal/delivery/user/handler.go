package user

import (
	"strconv"

	"github.com/morris-zheng/go-slim/internal/common/response"
	"github.com/morris-zheng/go-slim/internal/domain"
	userDomain "github.com/morris-zheng/go-slim/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	us *userDomain.Service
}

func NewHandler(svc *domain.ServiceContext) *Handler {
	return &Handler{
		us: userDomain.NewService(svc),
	}
}

func (h *Handler) Query(c *gin.Context) {
	var qp userDomain.QueryParams
	if err := c.ShouldBindQuery(&qp); err != nil {
		response.Fail(c, err.Error(), 404)
		return
	}
	if qp.Page <= 0 {
		qp.Page = 1
	}
	ul, total, _ := h.us.Query(qp)
	result := struct {
		List  []userDomain.User `json:"list"`
		Total int64             `json:"total"`
	}{
		List:  ul,
		Total: total,
	}
	response.Success(c, result)
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := h.us.Get(id)
	if err != nil {
		response.Fail(c, err.Error(), 404)
		return
	}
	response.Success(c, u)
}

func (h *Handler) Create(c *gin.Context) {
	err := h.us.Create(userDomain.User{
		Name: "test",
	})
	if err != nil {
		response.Fail(c, err.Error(), 404)
		return
	}
	response.Success(c, "success")
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := h.us.Get(id)
	if err != nil {
		response.Fail(c, err.Error(), 404)
		return
	}
	u.Name = "lala"
	err = h.us.Update(u)
	if err != nil {
		response.Fail(c, err.Error(), 404)
		return
	}
	response.Success(c, "success")
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.us.Delete(id)
	if err != nil {
		response.Fail(c, err.Error(), 404)
		return
	}
	response.Success(c, "success")
}
