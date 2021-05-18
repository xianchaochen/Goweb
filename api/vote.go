package api

import (
	"bluebell/common"
	"bluebell/entity"
	"bluebell/pkg/errno"
	"bluebell/pkg/translator"
	"bluebell/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func PostVoteHandler(c *gin.Context)  {
	p := new(entity.ParamVoteData)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, errno.ErrParam)
			return
		}
		//// validator.ValidationErrors类型错误则进行翻译
		msg, _ := json.Marshal(common.FormatTranslateMsg(errs.Translate(translator.Trans)))
		s := string(msg)
		c.JSON(http.StatusOK, errno.ErrParam.WithMsg(s))
		return
	}

	value, _ := c.Get(ContextUserIDKey)
	user_id, _ := value.(int)
	err := service.PostVote(strconv.Itoa(user_id), p.PostID, float64(p.Direction))
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrUserVoteFAILED)
		return
	}

	c.JSON(http.StatusOK, errno.OK)
}
