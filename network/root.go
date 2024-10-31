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

	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
