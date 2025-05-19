# gin_api

```
r := gin.Default()
sdk := gin_api.NewGinApiSdk(r)

sdk.Register(
    &gin_api.ApiData{
        Method: http.MethodGet,
        RelativePath: "health",
    },
    func(ctx *gin.Context){
        return gin.H{
            "status": "OK",
        }
    },
)
```