package models

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

type Model struct {
	Name  string
	Schema map[string]reflect.Kind
}

func CreateModelRoutes(model Model, router *gin.RouterGroup) {



}
