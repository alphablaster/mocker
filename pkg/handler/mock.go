package handler

import (
	"github.com/alphablaster/mocker/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createMock(c *gin.Context)  {
	var input entity.Mock
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Mock.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllMocks(c *gin.Context)  {

}

func (h *Handler) getMockById(c *gin.Context)  {

}

func (h *Handler) updateMock(c *gin.Context)  {

}

func (h *Handler) deleteMock(c *gin.Context)  {

}