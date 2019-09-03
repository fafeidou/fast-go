package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

// NewTrans 创建事务管理实例
func NewTrans(db *gorm.DB) *Trans {
	return &Trans{db}
}

// Trans 事务管理
type Trans struct {
	db *gorm.DB
}

func (a *Trans) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.Trans.%s", name)
}

// Begin 开启事务
func (a *Trans) Begin(ctx context.Context) (interface{}, error) {

	result := a.db.Begin()
	if err := result.Error; err != nil {
		return nil, errors.New("开启事务发生错误")
	}
	return result, nil
}

// Commit 提交事务
func (a *Trans) Commit(ctx context.Context, trans interface{}) error {

	db, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("未知的事务类型")
	}

	result := db.Commit()
	if err := result.Error; err != nil {
		return errors.New("提交事务发生错误")
	}
	return nil
}

// Rollback 回滚事务
func (a *Trans) Rollback(ctx context.Context, trans interface{}) error {

	db, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("未知的事务类型")
	}

	result := db.Rollback()
	if err := result.Error; err != nil {
		return errors.New("回滚事务发生错误")
	}
	return nil
}
