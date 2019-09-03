package main

import (
	"fast-go/internal/app/models"
	"fmt"
	"testing"
)
func TestMain(m *testing.M) {
	//setting.Setup()
	//models.Setup()
	//logging.Setup()
	//gredis.Setup()
	m.Run()
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

//router := routers.InitRouter()

//s := &http.Server{
//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
//	//Handler:        router,
//	ReadTimeout:    setting.ReadTimeout,
//	WriteTimeout:   setting.WriteTimeout,
//	MaxHeaderBytes: 1 << 20,
//}
//
//s.ListenAndServe()
//根据id查询
//article := models.GetArticle(1)
//fmt.Println(article)

//根绝条件查询
//maps := make(map[string]interface{})
//maps["state"] = 1
//count := models.GetArticleTotal(maps)
//fmt.Println(count)

//单个添加
//addMaps := make(map[string]interface{})
//addMaps["tag_id"] = 1
//addMaps["title"] = "testTitle"
//addMaps["desc"] = "testDesc"
//addMaps["content"] = "testContent"
//addMaps["created_by"] = "batman2"
//addMaps["state"] = 1
//models.AddArticle(addMaps)

//批量添加
//batchModels := make(map[int]models.Article, 3)
//model1 :=  models.Article{Title:"batman1"}
//model2 :=  models.Article{Title:"batman2"}
//model3 :=  models.Article{Title:"batman3"}
//batchModels[0] = model1
//batchModels[1] = model2
//batchModels[2] = model3
//models.AddArticles(batchModels)
//batchModels

//封装基础增删改查测试
//maps := make(map[string]interface{})
//maps["state"] = 1
//aa := new(models.Article)
//count2 := models.GetTotal(aa, maps)
//fmt.Println(count2)
//
//var article models.Article
//exist := models.ExistByID(&article, 2)
//fmt.Println(exist)

//list := make(models.List, 10)
//list[0] = *aa
//models.FindPage(list,0,10,maps)
//fmt.Println(list)
//var article models.Article
//article.Title = "123123gfdssfdcs"
//models.Add(&article)
//models.Delete(&article,17)
//data := make(map[string]interface {})
//data["title"] = "123123123"
//models.Edit(&article,10,data)
//models.EditByModel(&article,10)

//models.Get(*aa,2)
//fmt.Println(*aa)
//var articles []models.Article
//
////分页测试
//paginator := pagination.Paging(&pagination.Param{
//	DB:      models.GetDB(),
//	Page:    1,
//	Limit:   10,
//	OrderBy: []string{"id desc"},
//	ShowSQL: true,
//}, &articles)
//fmt.Println(*paginator)

//var article models.Article
//
//models.Get(&article,2)
//fmt.Println(article

//var articles []models.Article
//maps := make(map[string]interface{})
//maps["state"] = 1
//models.Find(&articles,1,10,maps)
//fmt.Println(articles)
//自己封装分页
//var whereOrder []models.PageWhereOrder
//var list []models.Article
//var total uint64
//var arr []interface{}
//arr = append(arr, "%batman%")
//whereOrder = append(whereOrder, models.PageWhereOrder{Where: "title like ?", Value: arr})
//order := "ID DESC , title ASC"
//whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})
//field := "id,tag_id,title"
//whereOrder = append(whereOrder,models.PageWhereOrder{Select:field})
//preload := "Tag"
//whereOrder = append(whereOrder,models.PageWhereOrder{PreLoad:preload})
//models.GetPage(&models.Article{},&models.Article{},&list,1,10,&total,whereOrder...)
//fmt.Println(list)
//aa := models.GetArticles(1,10,make(map[string]interface{}))
//aa := models.GetArticle(2)
//
//cache := cache_service.Article{ID: aa.ID}
//key := cache.GetArticleKey()
//redis 测试
//gredis.Set(key, aa, 3600)
//result := gredis.Exists("123")
//result,_ := gredis.Get(key)
//var article models.Article
//
//err:=json.Unmarshal(result,&article)
//if err!=nil{
//	fmt.Println(err)
//}
//fmt.Printf("article => %v" , article)
//logging.Debug("[info] start http server listening")