package types

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
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

func (s *Type) GetAllTypes(c *gin.Context, pageParam, limitParam, orderByParam, typeName string, startTime time.Time) ([]model.Types, helper.Paginate, error) {
	var (
		dataTypes []model.Types
		paginate  helper.Paginate
		totalData int64
		err       error
	)

	page, limit, offset, err := helper.Pagination(pageParam, limitParam)
	if err != nil {
		helper.ResponseAPI(c, http.StatusBadRequest, err.Error(), startTime)
		return dataTypes, paginate, err
	}

	dataTypes, totalData, err = s.typeMysql.GetAll(limit, offset, orderByParam, typeName)
	if err != nil {
		log.Printf("error type service GetAll: %s", err)
		helper.ResponseAPI(c, http.StatusInternalServerError, err.Error(), startTime)
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
