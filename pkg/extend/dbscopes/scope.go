package dbscopes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func Paginate(req *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(req.Query("page"))
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(req.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		fmt.Println("page",offset,pageSize)
		return db.Offset(offset).Limit(pageSize)
	}
}

func UpdateAll(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fields = append(fields, "deleted_at", "created_at", "id")
		return db.Select("*").Omit(fields...)
	}
}