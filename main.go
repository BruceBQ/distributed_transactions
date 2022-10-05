package main

import (
	"fmt"
	"net/http"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/lithammer/shortuuid"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultHTTPServer default url for http server. used by test and examples
	DefaultHTTPServer = "http://localhost:36789/api/dtmsvr"
	// DefaultJrpcServer default url for http json-rpc server. used by test and examples
	DefaultJrpcServer = "http://localhost:36789/api/json-rpc"
	// DefaultGrpcServer default url for grpc server. used by test and examples
	DefaultGrpcServer = "localhost:36790"
)

type Request struct {
	Amount uint `json:"amount"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/customer/verify", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Verify",
		})
	})

	r.Run()
	server := "localhost:5555"
	gid := shortuuid.New()
	saga := dtmcli.NewSaga(DefaultHTTPServer, gid)
	saga.Add(fmt.Sprintf("http://%s/order/create", server), fmt.Sprintf("http://%s/order/cancel", server), &Request{Amount: 10}).
		Add(fmt.Sprintf("http://%s/customer/verify", "localhost:5556"), "", nil)

	if err := saga.Submit(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(saga.Gid)
}
