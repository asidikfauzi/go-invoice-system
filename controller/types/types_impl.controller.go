package types

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/common/validator"
	"go-invoice-system/model"
	"log"
	"net/http"
	"time"
)

func (m *MasterTypes) GetAllTypes(c *gin.Context) {
	startTime := time.Now()

	var request model.Types
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
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

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataType, paginate, startTime)
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

	helper.ResponseDataAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataType, startTime)
	return
}

func (m *MasterTypes) CreateType(c *gin.Context) {
	startTime := time.Now()

	var request model.RequestType
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(request)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, validate, startTime)
		return
	}

	msg, err := m.TypeService.CreateType(c, request, startTime)
	if err != nil {
		log.Printf("error type controller CreateType :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusCreated, helper.SuccessCreated, []string{msg}, startTime)
	return

}

func (m *MasterTypes) UpdateType(c *gin.Context) {
	startTime := time.Now()

	typeId := c.Param("typeId")
	var request model.RequestType
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(request)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, validate, startTime)
		return
	}

	msg, err := m.TypeService.UpdateType(c, request, typeId, startTime)
	if err != nil {
		log.Printf("error type controller CreateType :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, helper.Success, []string{msg}, startTime)
	return

}
