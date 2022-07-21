package models

import (
	"context"
	"fmt"
	"github.com/fajarardiyanto/boiler-monolith-rest-api/interfaces"
	"github.com/fajarardiyanto/boiler-monolith-rest-api/internal/database"
	"github.com/fajarardiyanto/boiler-monolith-rest-api/internal/errors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ModelPost struct {
	db interfaces.SQL
}

func NewModelPost() PostRepository {
	return &ModelPost{
		db: database.DBConn,
	}
}

func (m *ModelPost) GetPost() ([]interfaces.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result []interfaces.Post

	tx := m.db.Orm().WithContext(ctx).Debug().
		Raw("select * from post").
		Debug().Model(&interfaces.Post{}).Scan(&result)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			errors.ErrorMessage(http.StatusNotFound, tx.Error.Error())
			return nil, tx.Error
		}
		errors.ErrorMessage(http.StatusInternalServerError, tx.Error.Error())
		return nil, tx.Error
	}

	if len(result) == 0 {
		errors.ErrorMessage(http.StatusNotFound, "data tidak ditemukan")
		return nil, fmt.Errorf("data tidak di temukan")
	}

	return result, nil
}
