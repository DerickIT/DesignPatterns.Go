package httpgo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var MsgFlag = map[int]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	409: "Conflict",
	500: "Internal Server Error",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return "Unknown"
}

func ReturnBadRequest(ctx *gin.Context, err error) {

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
			"message": GetMsg(http.StatusBadRequest),
		})
		log.Println(err.Error())

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  GetMsg(http.StatusBadRequest),
		})
		log.Println(GetMsg(http.StatusBadRequest))
	}
}

func ReturnJson(response *http.Response, ctx *gin.Context) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicln("Error reading response body:", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response body:", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  GetMsg(http.StatusOK),
		"data": data,
	})
}
