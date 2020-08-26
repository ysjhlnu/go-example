package service

import (
	"encoding/json"
	"go-gin-example/models"
	"go-gin-example/pkg/gredis"
	"go-gin-example/pkg/logging"
	"go-gin-example/service/cache_service"
)

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}
	article, err := models.GetArticles(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 3600)
	return article, nil
}