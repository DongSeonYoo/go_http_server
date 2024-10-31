package network

import "github.com/gin-gonic/gin"

/*
	configuration of router about API
*/

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	n := &Network{
		engin: gin.New(),
	}

	newUserRouter(n)

	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.POST(path, handler...)
}

func (n *Network) registerPUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.DELETE(path, handler...)
}
