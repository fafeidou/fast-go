package test

import (
	"fast-go/internal/app/models"
	"fast-go/pkg/gredis"
	"fast-go/pkg/logging"
	"fast-go/setting"
	"fmt"
	"testing"
)
func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
}
func TestMain(m *testing.M) {
	//setting.Setup()
	//models.Setup()
	//logging.Setup()
	//gredis.Setup()
	var whereOrder []models.PageWhereOrder
	var list []models.Article
	var total uint64
	var arr []interface{}
	arr = append(arr, "%batman%")
	whereOrder = append(whereOrder, models.PageWhereOrder{Where: "title like ?", Value: arr})
	order := "ID DESC , title ASC"
	whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})
	field := "id,tag_id,title"
	whereOrder = append(whereOrder,models.PageWhereOrder{Select:field})
	preload := "Tag"
	whereOrder = append(whereOrder,models.PageWhereOrder{PreLoad:preload})
	models.GetPage(&models.Article{},&models.Article{},&list,1,10,&total,whereOrder...)
	fmt.Println(list)
	//m.Log("第一个测试通过了") //记录一些你期望记录的信息
}