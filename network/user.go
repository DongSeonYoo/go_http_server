package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}
	})

	userRouterInstance.router.registerGET("/", userRouterInstance.get)
	userRouterInstance.router.registerPOST("/", userRouterInstance.create)
	userRouterInstance.router.registerPUT("/", userRouterInstance.update)
	userRouterInstance.router.registerDELETE("/", userRouterInstance.delete)

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다")
}
