package gin_api

import (
	"github.com/gin-gonic/gin"
)

func NewGinApiSdk(r *gin.Engine) *GinApiSdk {
	sdk := &GinApiSdk{
		R:           r,
		middlewares: []MiddlewareFunc{},
	}

	return sdk
}

type MiddlewareFunc func(api Api)

type GinApiSdk struct {
	R *gin.Engine

	middlewares []MiddlewareFunc
}

func (sdk *GinApiSdk) GetGinEngine() *gin.Engine {
	return sdk.R
}

func (sdk *GinApiSdk) Group(relativePath string) *GinSdkGroup {
	newGroup := GinSdkGroup{
		sdk:      sdk,
		G:        sdk.R.Group(relativePath),
		Basepath: relativePath,
	}

	return &newGroup
}

type RegisterFunc func(api Api, handlers ...gin.HandlerFunc) gin.IRoutes

func (sdk *GinApiSdk) RegisterGroup(relativePath string, groupHandler func(group *gin.RouterGroup, register RegisterFunc)) {
	r := sdk.R.Group(relativePath)

	var registfn RegisterFunc = func(api Api, handlers ...gin.HandlerFunc) gin.IRoutes {
		api.SetGroupPath(relativePath)

		for _, middlewareFunc := range sdk.middlewares {
			middlewareFunc(api)
		}

		return r.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
	}

	groupHandler(r, registfn)
}

func (sdk *GinApiSdk) Register(api Api, handlers ...gin.HandlerFunc) gin.IRoutes {

	for _, middlewareFunc := range sdk.middlewares {
		middlewareFunc(api)
	}

	return sdk.R.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
}

func (sdk *GinApiSdk) Use(handler func(Api)) *GinApiSdk {
	sdk.middlewares = append(sdk.middlewares, handler)
	return sdk
}
