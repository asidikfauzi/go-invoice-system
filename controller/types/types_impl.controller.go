package types

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"log"
	"net/http"
	"time"
)

func (m *MasterTypes) GetAllTypes(c *gin.Context) {
	startTime := time.Now()

	var request model.Types
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, http.StatusInternalServerError, err.Error(), startTime)
		return
	}

	pageParam := c.Query("page")
	limitParam := c.Query("limit")
	orderByParam := c.Query("orderBy")

	dataType, paginate, err := m.TypeService.GetAllTypes(c, pageParam, limitParam, orderByParam, request.TypeName, startTime)
	if err != nil {
		log.Printf("error type controller GetAllTypes :%s", err)
		return
	}

	helper.ResponseDataAPIWithPagination(c, http.StatusOK, helper.SuccessGetData, dataType, paginate, startTime)
	return
}

func (m *MasterTypes) FindTypeById(c *gin.Context) {
	startTime := time.Now()

	typeId := c.Param("typeId")
	dataType, err := m.TypeService.FindTypeById(c, typeId, startTime)
	if err != nil {
		log.Printf("error type controller FindTypeById :%s", err)
		return
	}
	helper.ResponseDataAPI(c, http.StatusOK, helper.SuccessGetData, dataType, startTime)
	return

}
