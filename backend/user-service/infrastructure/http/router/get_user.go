package router

import (
	"monolith/application"
	errApp "monolith/application/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getByIdParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type getUserByIdRouter struct {
	getUserApplication application.GetUserApplication
}

func newGetUserByIdRouter(router *gin.Engine, getUserApplication application.GetUserApplication) (*gin.Engine, error) {
	if getUserApplication == nil {
		return nil, errApp.NewBadArgumentError("get user application is required")
	}
	getRouter := &getUserByIdRouter{
		getUserApplication: getUserApplication,
	}
	router.GET("/api/users/:id", getRouter.GetUserById)
	return router, nil

}

func (gur *getUserByIdRouter) GetUserById(c *gin.Context) {
	var params getByIdParams
	if err := c.ShouldBindUri(&params); err != nil {
		panic(errApp.NewBadArgumentErrorWithCause("Error to parse body.", err))
	}
	resp, err := gur.getUserApplication.GetUser(params.ID)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}
