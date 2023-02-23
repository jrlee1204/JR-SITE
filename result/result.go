package result

import (
	"github.com/gin-gonic/gin"
	"jr_site/constant"
	"jr_site/model"
	"net/http"
)

func Success(c *gin.Context) {
	SuccessWithData(c, nil)
}

func SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, model.Result{
		Code: constant.SUCCESS,
		Data: data,
	})
}

func Fail(c *gin.Context, errorCode int, errorMsg string) {
	c.JSON(http.StatusOK, model.Result{
		Code:      constant.FAIL,
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
	})
}
