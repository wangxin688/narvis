package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HasParams(g *gin.Context, key string) bool {
	if _, ok := g.Params.Get(key); ok {
		return true
	}
	if _, ok := g.GetQuery(key); ok {
		return true
	}
	if _, ok := g.GetPostForm(key); ok {
		return true
	}
	return false
}

func ValidateUuidString(Id string) error {
	_, err := uuid.Parse(Id)
	return err
}
