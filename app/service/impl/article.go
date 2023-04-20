package impl

import (
	"context"
	"encoding/json"
	"github.com/davveo/lemonShop/app/cache_service"
	"github.com/davveo/lemonShop/dao"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/cache"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
	"github.com/davveo/lemonShop/pkg/log"
)

type ArticleReq struct {
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

type ArticleService struct {
	articleRepo dao.ArticleDao
}

func NewArticleService() ArticleService {
	return ArticleService{
		articleRepo: dao.NewArticleDao(dbLocal.GRpo),
	}
}

func (a *ArticleService) Add(ctx context.Context, req ArticleReq) error {
	article := map[string]interface{}{
		"tag_id":          req.TagID,
		"title":           req.Title,
		"desc":            req.Desc,
		"content":         req.Content,
		"created_by":      req.CreatedBy,
		"cover_image_url": req.CoverImageUrl,
		"state":           req.State,
	}

	if err := a.articleRepo.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) Edit(ctx context.Context, req ArticleReq) error {
	return a.articleRepo.EditArticle(req.ID, map[string]interface{}{
		"tag_id":          req.TagID,
		"title":           req.Title,
		"desc":            req.Desc,
		"content":         req.Content,
		"cover_image_url": req.CoverImageUrl,
		"state":           req.State,
		"modified_by":     req.ModifiedBy,
	})
}

func (a *ArticleService) Get(ctx context.Context, req ArticleReq) (*models.Article, error) {
	var cacheArticle *models.Article
	articleCache := cache_service.Article{ID: req.ID}
	key := articleCache.GetArticleKey()
	if cache.Cache.Exists(key) {
		data, err := cache.Cache.Get(key)
		if err != nil {
			log.Info(err)
		} else {
			json.Unmarshal([]byte(data), &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := a.articleRepo.GetArticle(req.ID)
	if err != nil {
		return nil, err
	}

	//cache.Cache.Set(key, article, 3600)
	return article, nil
}

func (a *ArticleService) GetAll(ctx context.Context, req ArticleReq) ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	articleCache := cache_service.Article{
		TagID: req.TagID,
		State: req.State,

		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}
	key := articleCache.GetArticlesKey()
	if cache.Cache.Exists(key) {
		data, err := cache.Cache.Get(key)
		if err != nil {
			log.Info(err)
		} else {
			json.Unmarshal([]byte(data), &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := a.articleRepo.GetArticles(req.PageNum, req.PageSize, a.getMaps(ctx, req))
	if err != nil {
		return nil, err
	}

	//cache.Set(key, articles, 3600)
	return articles, nil
}

func (a *ArticleService) Delete(ctx context.Context, req ArticleReq) error {
	return a.articleRepo.DeleteArticle(req.ID)
}

func (a *ArticleService) ExistByID(ctx context.Context, req ArticleReq) (bool, error) {
	return a.articleRepo.ExistArticleByID(req.ID)
}

func (a *ArticleService) Count(ctx context.Context, req ArticleReq) (int64, error) {
	return a.articleRepo.GetArticleTotal(a.getMaps(ctx, req))
}

func (a *ArticleService) getMaps(ctx context.Context, req ArticleReq) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if req.State != -1 {
		maps["state"] = req.State
	}
	if req.TagID != -1 {
		maps["tag_id"] = req.TagID
	}

	return maps
}
