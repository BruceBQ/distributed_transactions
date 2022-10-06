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
		fmt.Println("customer verify")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Customer Verify",
		})
	})
	r.POST("/customer/cancel", func(ctx *gin.Context) {
		fmt.Println("Customer cancel")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Customer cancel",
		})
	})
	r.POST("/do", func(ctx *gin.Context) {
		server := "localhost:8080"
		gid := shortuuid.New()
		saga := dtmcli.NewSaga(DefaultHTTPServer, gid)
		fmt.Println("saga options", saga.TransOptions)
		saga.BuildCustomOptions()
		saga.Add(fmt.Sprintf("http://%s/order/create", server), fmt.Sprintf("http://%s/order/cancel", server), &Request{Amount: 10}).
			Add(fmt.Sprintf("http://%s/customer/verify", server), fmt.Sprintf("http://%s/customer/cancel", server), nil)

		fmt.Println("DB", dtmcli.GetCurrentDBType())
		saga.TimeoutToFail = 20
		fmt.Println("Timeout to Faile ", saga.TimeoutToFail)
		fmt.Println("Retry Interval ", saga.RequestTimeout)
		if err := saga.Submit(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(saga.Gid)
		// wfName := "wf"
		// err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
		// 	return nil
		// })
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "order created",
		})

	})

	r.POST("/order/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "order created",
		})

	})

	r.POST("/order/cancel", func(ctx *gin.Context) {
		fmt.Println("Cancel order")
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "order canceled",
		})
	})

	r.Run()

}
