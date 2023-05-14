package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/morris-zheng/go-slim/internal/common/response"
	"github.com/morris-zheng/go-slim/internal/domain"
	userDomain "github.com/morris-zheng/go-slim/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *domain.ServiceContext
	uc  *userDomain.UseCase
}

func NewHandler(svc *domain.ServiceContext) *Handler {
	return &Handler{
		svc: svc,
		uc:  userDomain.NewUseCase(svc),
	}
}

type QueryResp struct {
	List  []userDomain.User `json:"list"`
	Total int64             `json:"total"`
}

func (h *Handler) Query(c *gin.Context) {
	var qp userDomain.QueryParams
	if err := c.ShouldBindQuery(&qp); err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	if qp.Page <= 0 {
		qp.Page = 1
	}

	ul, total, _ := h.uc.Query(qp)
	response.Success(c, response.Response{
		Data: QueryResp{
			List:  ul,
			Total: total,
		},
	})
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Logger.Info(c, fmt.Sprintf("get user:  %d", id))

	u, err := h.uc.Get(id)
	if err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	response.Success(c, response.Response{
		Data: u,
	})
}

func (h *Handler) Create(c *gin.Context) {
	err := h.uc.Create(userDomain.User{
		Name: "test",
	})
	if err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	response.Success(c, response.Response{
		Data:     "success",
		HttpCode: http.StatusCreated,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := h.uc.Get(id)
	if err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	u.Name = "lala"
	err = h.uc.Update(u)
	if err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	response.Success(c, response.Response{
		Data: "success",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.uc.Get(id)
	if err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	err = h.uc.Delete(id)
	if err != nil {
		response.Fail(c, response.Response{
			Msg:  err.Error(),
			Code: 10404,
		})
		return
	}

	response.Success(c, response.Response{
		Data: "success",
	})
}
