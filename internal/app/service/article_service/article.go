package article_service

import (
	"fast-go/internal/app/models"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	if err := models.Add(a); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	return models.EditByModel(a, a.ID)
}

//func (a *Article) Get() (article *models.Article, e error) {
//	var cacheArticle *models.Article
//
//	cache := cache_service.Article{ID: a.ID}
//	key := cache.GetArticleKey()
//	if gredis.Exists(key) {
//		data, err := gredis.Get(key)
//		if err != nil {
//			logging.Info(err)
//		} else {
//			json.Unmarshal(data, &cacheArticle)
//			return cacheArticle, nil
//		}
//	}
//
//	err := models.Get(article, a.ID)
//	if err != nil {
//		return nil, err
//	}
//	gredis.Set(key, a, 3600)
//	return article, nil
//}
//
//func (a *Article) GetAll() ([]*models.Article, error) {
//	var (
//		articles, cacheArticles []*models.Article
//	)
//
//	cache := cache_service.Article{
//		TagID: a.TagID,
//		State: a.State,
//
//		PageNum:  a.PageNum,
//		PageSize: a.PageSize,
//	}
//	key := cache.GetArticlesKey()
//	if gredis.Exists(key) {
//		data, err := gredis.Get(key)
//		if err != nil {
//			//logging.Info(err)
//		} else {
//			json.Unmarshal(data, &cacheArticles)
//			return cacheArticles, nil
//		}
//	}
//
//	err := models.Find(&articles, a.PageNum, a.PageSize, a.getMaps())
//	if err != nil {
//		return nil, err
//	}
//
//	gredis.Set(key, articles, 3600)
//	return articles, nil
//}

func (a *Article) Delete() error {
	var article models.Article
	return models.Delete(&article, a.ID)
}

func (a *Article) ExistByID() (bool, error) {
	var article models.Article
	return models.ExistByID(&article, a.ID)
}

func (a *Article) Count() (int, error) {
	var article models.Article
	return models.GetTotal(&article, a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
