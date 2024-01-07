package items

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/common/validator"
	"go-invoice-system/model"
	"log"
	"net/http"
	"time"
)

func (m *MasterItems) GetAllItems(c *gin.Context) {
	startTime := time.Now()

	var request model.Items
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	pageParam := c.Query("page")
	limitParam := c.Query("limit")
	orderByParam := c.Query("orderBy")

	dataItem, paginate, err := m.ItemService.GetAllItems(c, pageParam, limitParam, orderByParam, request.ItemName, startTime)
	if err != nil {
		log.Printf("error item controller GetAllItems :%s", err)
		return
	}

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataItem, paginate, startTime)
	return
}

func (m *MasterItems) FindItemById(c *gin.Context) {
	startTime := time.Now()

	itemId := c.Param("itemId")
	dataItem, err := m.ItemService.FindItemById(c, itemId, startTime)
	if err != nil {
		log.Printf("error item controller FindItemById :%s", err)
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataItem, startTime)
	return
}

func (m *MasterItems) CreateItem(c *gin.Context) {
	startTime := time.Now()

	var request model.RequestItem
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(request)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, validate, startTime)
		return
	}

	msg, err := m.ItemService.CreateItem(c, request, startTime)
	if err != nil {
		log.Printf("error item controller CreateItem :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusCreated, helper.SuccessCreated, []string{msg}, startTime)
	return

}

func (m *MasterItems) UpdateItem(c *gin.Context) {
	startTime := time.Now()

	itemId := c.Param("itemId")
	var request model.RequestItem
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(request)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, validate, startTime)
		return
	}

	msg, err := m.ItemService.UpdateItem(c, request, itemId, startTime)
	if err != nil {
		log.Printf("error item controller UpdateItem :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, helper.Success, []string{msg}, startTime)
	return

}

func (m *MasterItems) DeleteItem(c *gin.Context) {
	startTime := time.Now()

	itemId := c.Param("itemId")
	msg, err := m.ItemService.DeleteItem(c, itemId, startTime)
	if err != nil {
		log.Printf("error item controller DeleteItem :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, helper.Success, []string{msg}, startTime)
	return

}
