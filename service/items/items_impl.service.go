package items

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"go-invoice-system/repository/mysql/items"
	"go-invoice-system/repository/mysql/types"
	"log"
	"math"
	"net/http"
	"time"
)

type Item struct {
	itemMysql items.ItemsMysql
	typeMysql types.TypesMysql
}

func NewItemService(cm items.ItemsMysql, tm types.TypesMysql) ItemsService {
	return &Item{
		itemMysql: cm,
		typeMysql: tm,
	}
}

func (s *Item) GetAllItems(c *gin.Context, pageParam, limitParam, orderByParam, itemName string, startTime time.Time) ([]model.GetItem, helper.Paginate, error) {
	var (
		dataItems []model.GetItem
		paginate  helper.Paginate
		totalData int64
		err       error
	)

	page, limit, offset, err := helper.Pagination(pageParam, limitParam)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return dataItems, paginate, err
	}

	dataItems, totalData, err = s.itemMysql.GetAll(limit, offset, orderByParam, itemName)
	if err != nil {
		log.Printf("error item service GetAll: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return dataItems, paginate, err
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	paginate = helper.Paginate{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  totalData,
	}

	return dataItems, paginate, nil
}

func (s *Item) FindItemById(c *gin.Context, itemId string, startTime time.Time) (model.GetItem, error) {
	var (
		item model.GetItem
		err  error
	)

	item, err = s.itemMysql.FindById(itemId)
	if err != nil {
		err = fmt.Errorf("item_id '%s' not found", itemId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return item, err
	}

	return item, nil
}

func (s *Item) CreateItem(c *gin.Context, request model.RequestItem, startTime time.Time) (string, error) {
	var err error

	findByName, _ := s.itemMysql.FindByName(request.ItemName)
	if findByName.ItemName != "" {
		err = fmt.Errorf("item_name '%s' already exists", request.ItemName)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	_, err = s.typeMysql.FindById(request.TypeID)
	if err != nil {
		err = fmt.Errorf("type_id '%s' not found", request.TypeID)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	typeIdUuid, err := uuid.Parse(request.TypeID)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	item := domain.Items{
		IDItem:       uuid.New(),
		ItemName:     request.ItemName,
		ItemQuantity: request.ItemQuantity,
		ItemPrice:    request.ItemPrice,
		TypeID:       typeIdUuid,
		CreatedAt:    time.Now(),
	}

	err = s.itemMysql.Create(&item)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessCreatedData, nil
}

func (s *Item) UpdateItem(c *gin.Context, request model.RequestItem, itemId string, startTime time.Time) (string, error) {
	var err error

	_, err = s.itemMysql.FindById(itemId)
	if err != nil {
		err = fmt.Errorf("item_id '%s' not found", itemId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	_, err = s.typeMysql.FindById(request.TypeID)
	if err != nil {
		err = fmt.Errorf("type_id '%s' not found", request.TypeID)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	itemIdUuid, err := uuid.Parse(itemId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	timeUpdate := time.Now()
	item := domain.Items{
		IDItem:       itemIdUuid,
		ItemName:     request.ItemName,
		ItemQuantity: request.ItemQuantity,
		ItemPrice:    request.ItemPrice,
		TypeID:       itemIdUuid,
		UpdatedAt:    &timeUpdate,
	}

	exists, _ := s.itemMysql.CheckUpdateExists(item)

	if exists == true {
		err = fmt.Errorf("item_name '%s' already exists", request.ItemName)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	err = s.itemMysql.Update(&item)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessUpdatedData, nil
}

func (s *Item) DeleteItem(c *gin.Context, itemId string, startTime time.Time) (string, error) {
	var err error

	_, err = s.itemMysql.FindById(itemId)
	if err != nil {
		err = fmt.Errorf("item_id '%s' not found", itemId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	itemIdUuid, err := uuid.Parse(itemId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	timeDelete := time.Now()
	item := domain.Items{
		IDItem:    itemIdUuid,
		DeletedAt: &timeDelete,
	}

	err = s.itemMysql.Delete(&item)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessDeletedData, nil
}
