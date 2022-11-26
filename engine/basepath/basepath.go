package basepath

import (
	"fmt"
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

	fmt.Println("Set base path to: ", req.BasePath)
	BasePath = filepath.Clean(req.BasePath)
	c.Status(http.StatusOK)
}
