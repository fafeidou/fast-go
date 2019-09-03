package models

import (
	"context"
	"fast-go/conf"
	icontext "fast-go/internal/app/context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedOn  int    `json:"createdOn"`
	ModifiedOn int    `json:"modifiedOn"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

// 分页条件
type PageWhereOrder struct {
	Order   string
	Where   string
	Value   []interface{}
	Select  string
	PreLoad string
}

type Element interface{}
type List [] Element

/**
  根据条件查询记录数
 */
func GetTotal(result interface{}, maps interface{}) (int, error) {
	var count int
	if err := db.Model(result).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

/**
  根据id查询是否存在
 */
func ExistByID(result interface{}, id int) (bool, error) {
	var count int
	err := db.Select("id").Where("id = ?", id).First(result).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func Add(result interface{}) error {
	if err := db.Create(result).Error; err != nil {
		return err
	}
	return nil
}

func Delete(result interface{}, id int) error {
	if err := db.Where("id = ?", id).Delete(result).Error; err != nil {
		return err
	}
	return nil
}

func Edit(result interface{}, id int, data interface{}) error {
	if err := db.Model(result).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditByModel(result interface{}, id int) error {
	if err := db.Model(result).Where("id = ?", id).Updates(result).Error; err != nil {
		return err
	}
	return nil
}

func Get(result interface{}, id int) error {
	if err := db.Where("id = ?", id).First(result).Error; err != nil {
		return err
	}
	return nil
}

func Find(result interface{}, pageNum int, pageSize int, maps interface{}) error {
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(result).Error; err != nil {
		return err
	}
	return nil
}

// GetPage
func GetPage(model, where interface{}, out interface{}, pageIndex, pageSize uint64, totalCount *uint64, whereOrder ...PageWhereOrder) error {
	db := db.Model(model).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Select != "" {
				db = db.Select(wo.Select)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
			if wo.PreLoad != "" {
				db = db.Preload(wo.PreLoad)
			}

		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// Setup initializes the database instance
func Setup() {
	var err error

	db, err = gorm.Open(conf.App.Database.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.App.Database.User,
		conf.App.Database.Password,
		conf.App.Database.Host,
		conf.App.Database.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.App.Database.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, db *gorm.DB, fn func(context.Context) error) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	transModel := NewTrans(db)
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(icontext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

func CloseDB() {
	defer db.Close()
}

func GetDB() (data *gorm.DB) {
	return db
}
