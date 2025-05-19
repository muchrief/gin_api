package gin_api

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

type GinSdkGroup struct {
	sdk      *GinApiSdk
	G        *gin.RouterGroup
	Basepath string
}

func (grp *GinSdkGroup) GetGinEngine() *gin.Engine {
	return grp.sdk.R
}

func (grp *GinSdkGroup) Group(path string) *GinSdkGroup {
	base, _ := url.JoinPath(grp.Basepath, path)
	newGroup := GinSdkGroup{
		sdk:      grp.sdk,
		G:        grp.G.Group(path),
		Basepath: base,
	}

	return &newGroup
}

func (grp *GinSdkGroup) Register(api Api, handlers ...gin.HandlerFunc) gin.IRoutes {
	api.SetGroupPath(grp.Basepath)

	for _, middlewareFunc := range grp.sdk.middlewares {
		middlewareFunc(api)
	}

	return grp.G.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
}
