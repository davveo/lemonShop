package dao

import (
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"
	"gorm.io/gorm"
)

type ArticleDao struct {
	repo db.Repo
}

func NewArticleDao(db db.Repo) ArticleDao {
	return ArticleDao{
		repo: db,
	}
}

// ExistArticleByID checks if an article exists based on ID
func (articleDao ArticleDao) ExistArticleByID(id int) (bool, error) {
	var article models.Article

	readDB := articleDao.repo.GetDbR()

	err := readDB.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func (articleDao ArticleDao) GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	readDB := articleDao.repo.GetDbR()

	if err := readDB.Model(&models.Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func (articleDao ArticleDao) GetArticles(pageNum int, pageSize int, maps interface{}) ([]*models.Article, error) {
	var articles []*models.Article
	readDB := articleDao.repo.GetDbR()

	err := readDB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func (articleDao ArticleDao) GetArticle(id int) (*models.Article, error) {
	var article models.Article
	readDB := articleDao.repo.GetDbR()
	err := readDB.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

// EditArticle modify a single article
func (articleDao ArticleDao) EditArticle(id int, data interface{}) error {
	writeDB := articleDao.repo.GetDbW()
	if err := writeDB.Model(&models.Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddArticle add a single article
func (articleDao ArticleDao) AddArticle(data map[string]interface{}) error {
	article := models.Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	writeDB := articleDao.repo.GetDbW()
	if err := writeDB.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// DeleteArticle delete a single article
func (articleDao ArticleDao) DeleteArticle(id int) error {
	writeDB := articleDao.repo.GetDbW()
	if err := writeDB.Where("id = ?", id).Delete(models.Article{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllArticle clear all article
func (articleDao ArticleDao) CleanAllArticle() error {
	writeDB := articleDao.repo.GetDbW()
	if err := writeDB.Unscoped().Where("deleted_on != ? ", 0).Delete(&models.Article{}).Error; err != nil {
		return err
	}

	return nil
}
