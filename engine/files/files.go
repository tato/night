package files

import (
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"supok.es/night/basepath"
)

type FileItem struct {
	Path string `json:"path"`
}

type GetFileTreeResponse struct {
	Files []FileItem `json:"files"`
}

func GetFileTree(c *gin.Context) {
	resp := GetFileTreeResponse{Files: make([]FileItem, 0)}

	if len(basepath.BasePath) == 0 {
		c.JSON(http.StatusNoContent, resp)
		return
	}

	err := filepath.Walk(basepath.BasePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				rel, err := filepath.Rel(basepath.BasePath, path)
				if err != nil {
					return err
				}

				item := FileItem{
					Path: rel,
				}
				resp.Files = append(resp.Files, item)
			}

			return nil
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetFile(c *gin.Context) {

	if len(basepath.BasePath) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	filePath := c.Param("path")

	c.File(path.Join(basepath.BasePath, filePath))
}
