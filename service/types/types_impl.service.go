package types

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"go-invoice-system/repository/mysql/types"
	"log"
	"math"
	"net/http"
	"time"
)

type Type struct {
	typeMysql types.TypesMysql
}

func NewTypeService(tm types.TypesMysql) TypesService {
	return &Type{
		typeMysql: tm,
	}
}

func (s *Type) GetAllTypes(c *gin.Context, pageParam, limitParam, orderByParam, typeName string, startTime time.Time) ([]model.GetType, helper.Paginate, error) {
	var (
		dataTypes []model.GetType
		paginate  helper.Paginate
		totalData int64
		err       error
	)

	page, limit, offset, err := helper.Pagination(pageParam, limitParam)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return dataTypes, paginate, err
	}

	dataTypes, totalData, err = s.typeMysql.GetAll(limit, offset, orderByParam, typeName)
	if err != nil {
		log.Printf("error type service GetAll: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return dataTypes, paginate, err
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	paginate = helper.Paginate{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  totalData,
	}

	return dataTypes, paginate, nil
}

func (s *Type) FindTypeById(c *gin.Context, typeId string, startTime time.Time) (model.GetType, error) {
	var (
		typ model.GetType
		err error
	)

	typ, err = s.typeMysql.FindById(typeId)
	if err != nil {
		err = fmt.Errorf("type_id '%s' not found", typeId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return typ, err
	}

	return typ, nil
}

func (s *Type) CreateType(c *gin.Context, request model.RequestType, startTime time.Time) (string, error) {
	var err error

	findByName, _ := s.typeMysql.FindByName(request.TypeName)

	if findByName.TypeName != "" {
		err = fmt.Errorf("type_name '%s' already exists", request.TypeName)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	types := domain.Types{
		IDType:    uuid.New(),
		TypeName:  request.TypeName,
		CreatedAt: time.Now(),
	}

	err = s.typeMysql.Create(&types)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessCreatedData, nil
}

func (s *Type) UpdateType(c *gin.Context, request model.RequestType, typeId string, startTime time.Time) (string, error) {
	var err error

	_, err = s.typeMysql.FindById(typeId)
	if err != nil {
		err = fmt.Errorf("type_id '%s' not found", typeId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	typeIdUuid, err := uuid.Parse(typeId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	timeUpdate := time.Now()
	types := domain.Types{
		IDType:    typeIdUuid,
		TypeName:  request.TypeName,
		UpdatedAt: &timeUpdate,
	}

	exists, _ := s.typeMysql.CheckUpdateExists(types)
	if exists == true {
		err = fmt.Errorf("type_name '%s' already exists", request.TypeName)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	err = s.typeMysql.Update(&types)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessUpdatedData, nil
}

func (s *Type) DeleteType(c *gin.Context, typeId string, startTime time.Time) (string, error) {
	var err error

	_, err = s.typeMysql.FindById(typeId)
	if err != nil {
		err = fmt.Errorf("type_id '%s' not found", typeId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	typeIdUuid, err := uuid.Parse(typeId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	timeDelete := time.Now()
	types := domain.Types{
		IDType:    typeIdUuid,
		DeletedAt: &timeDelete,
	}

	err = s.typeMysql.Delete(&types)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessDeletedData, nil
}
