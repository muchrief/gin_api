package gin_api

import "github.com/gin-gonic/gin"

type Api interface {
	GetFullUriPath() string
	GetTags() []string
	GetSummary() string
	GetDescription() string
	GetKeyName() string
	GetQuery() any
	GetPayload() any
	GetResponse() any
	GetMethod() string
	GetGroupPath() string
	GetRelativePath() string

	SetGroupPath(path string)
}

type ApiSdk interface {
	GetGinEngine() *gin.Engine
	Group(relativePath string) *GinSdkGroup
	RegisterGroup(relativePath string, groupHandler func(group *gin.RouterGroup, register RegisterFunc))
	Register(api Api, handlers ...gin.HandlerFunc) gin.IRoutes
	Use(handler func(Api)) *GinApiSdk
}

type SdkGroup interface {
	GetGinEngine() *gin.Engine
	Group(path string) *GinSdkGroup
	Register(api Api, handlers ...gin.HandlerFunc) gin.IRoutes
}
