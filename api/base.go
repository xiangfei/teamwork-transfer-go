package api

import (
        "github.com/kataras/iris/v12"
)

type AwifiBaseApi struct {
}

func NewAwifiBaseApi() *AwifiBaseApi {

        return &AwifiBaseApi{}

}


func (zk *AwifiBaseApi) Index(ctx iris.Context) {
        ctx.Writef("Hello from the server")

}

func (zk *AwifiBaseApi) IndexJson(ctx iris.Context) {
        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    "",
        })

}



