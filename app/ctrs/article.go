package ctrs

import (
	"context"
	"github.com/astaxie/beego/validation"
	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/app/service/impl"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

	"github.com/davveo/lemonShop/pkg/app"
	"github.com/davveo/lemonShop/pkg/util"
)

type ArticleController struct {
	articleService impl.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		articleService: impl.NewArticleService(),
	}
}

func (ac *ArticleController) GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}

	req := impl.ArticleReq{ID: id}
	exists, err := ac.articleService.ExistByID(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, consts.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := ac.articleService.Get(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, article)
}

func (ac *ArticleController) GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	tagId := -1
	if arg := c.PostForm("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}

	req := impl.ArticleReq{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: 20,
	}

	total, err := ac.articleService.Count(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := ac.articleService.GetAll(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, consts.SUCCESS, data)
}

func (ac *ArticleController) AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form entity.AddArticleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != consts.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//tagService := tag_service.TagReq{ID: form.TagID}
	//exists, err := tagService.ExistByID()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
	//	return
	//}
	//
	//if !exists {
	//	appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
	//	return
	//}

	articleReq := impl.ArticleReq{
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
		CreatedBy:     form.CreatedBy,
	}
	if err := ac.articleService.Add(context.Background(), articleReq); err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, nil)
}

func (ac *ArticleController) EditArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = entity.EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != consts.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	req := impl.ArticleReq{
		ID:            form.ID,
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		ModifiedBy:    form.ModifiedBy,
		State:         form.State,
	}
	exists, err := ac.articleService.ExistByID(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, consts.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	//tagService := tag_service.TagReq{ID: form.TagID}
	//exists, err = tagService.ExistByID()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
	//	return
	//}
	//
	//if !exists {
	//	appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
	//	return
	//}
	//
	//err = a.articleService.Edit(context.Background(), req)
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
	//	return
	//}

	appG.Response(http.StatusOK, consts.SUCCESS, nil)
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		appG.Response(http.StatusOK, consts.INVALID_PARAMS, nil)
		return
	}

	req := impl.ArticleReq{ID: id}
	exists, err := ac.articleService.ExistByID(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, consts.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = ac.articleService.Delete(context.Background(), req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, nil)
}
