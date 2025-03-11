// Basic GORM CRUD operations for database entities that ease the development of handlers.
package repository

import (
	"github.com/cogniia/core-api-template/src/database"
	database_model "github.com/cogniia/core-api-template/src/database/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPaginated[T any](
	entity T,
	pagination database_model.Paginable,
	order database_model.Orderable,
	whereable database_model.Whereable,
	prefix string,
	db *gorm.DB,
) ([]T, error) {
	if db == nil {
		db = database.Connection().Model(&entity)
	}

	if pagination != nil {
		pagination.PaginateQuery(&db)
	}
	if order != nil {
		order.OrderQuery(&db, prefix)
	}
	if whereable != nil {
		whereable.Where(&db, prefix)
	}

	var entities []T

	if err := db.Where(&entity).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func First[T any](
	entity T,
	offset int,
	order database_model.Orderable,
	whereable database_model.Whereable,
	prefix string,
	db *gorm.DB,
) (T, error) {
	if db == nil {
		db = database.Connection().Model(&entity)
	}

	if order != nil {
		order.OrderQuery(&db, prefix)
	}
	if whereable != nil {
		whereable.Where(&db, prefix)
	}

	var result T

	if err := db.Where(&entity).Offset(offset).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func Create[T any](entity T, db *gorm.DB) (T, error) {
	if db == nil {
		db = database.Connection()
	}

	if err := db.Create(&entity).Error; err != nil {
		return entity, err
	}

	return entity, nil
}

func Updates[T any](updateData interface{}, where *T, db *gorm.DB) (T, error) {
	var entity T
	if db == nil {
		db = database.Connection()
	}
	db = db.Model(&entity)
	if where != nil {
		db = db.Where(where)
	}

	if err := db.Updates(updateData).Error; err != nil {
		return entity, err
	}

	return entity, nil
}

func DeleteById[T any](entityId uuid.UUID, db *gorm.DB) error {
	if db == nil {
		db = database.Connection()
	}

	var entity T

	if err := db.Delete(&entity, entityId).Error; err != nil {
		return err
	}

	return nil
}

func Count[T any](
	entity T,
	order database_model.Orderable,
	whereable database_model.Whereable,
	prefix string,
	db *gorm.DB,
) (int64, error) {
	if db == nil {
		db = database.Connection().Model(&entity)
	}

	if whereable != nil {
		whereable.Where(&db, prefix)
	}
	if order != nil {
		order.OrderQuery(&db, prefix)
	}

	var entityCount int64
	if err := db.Where(&entity).Count(&entityCount).Error; err != nil {
		return entityCount, err
	}

	return entityCount, nil
}
