package basepath

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// OS path where the night library is reading files from
var BasePath string

type GetBasePathResponse struct {
	BasePath string `json:"basepath"`
}

type PutBasePathRequest struct {
	BasePath string `json:"basepath"`
}

func GetBasePath(c *gin.Context) {

	resp := GetBasePathResponse{
		BasePath: BasePath,
	}
	c.JSON(http.StatusOK, resp)
}

func PutBasePath(c *gin.Context) {

	var req PutBasePathRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	BasePath = filepath.Clean(req.BasePath)

	resp := GetBasePathResponse{
		BasePath: BasePath,
	}
	c.JSON(http.StatusOK, resp)
}
