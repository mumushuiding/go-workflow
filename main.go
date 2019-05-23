package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mumushuiding/go-workflow/workflow-engine/service"

	config "github.com/mumushuiding/go-workflow/workflow-config"
	controller "github.com/mumushuiding/go-workflow/workflow-controller"
	model "github.com/mumushuiding/go-workflow/workflow-engine/model"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/workflow/", controller.Index)
	//-------------------------流程定义----------------------
	mux.HandleFunc("/workflow/procdef/save", controller.SaveProcdef)
	mux.HandleFunc("/workflow/procdef/findAll", controller.FindAllProcdefPage)
	mux.HandleFunc("/workflow/procdef/delById", controller.DelProcdefByID)
	// -----------------------流程实例-----------------------
	mux.HandleFunc("/workflow/process/start", controller.StartProcessInstance)
	mux.HandleFunc("/workflow/process/findTask", controller.FindMyProcInstPageAsJSON)
	// mux.HandleFunc("/workflow/process/moveToHistory", controller.MoveFinishedProcInstToHistory)
	// -----------------------任务--------------------------
	mux.HandleFunc("/workflow/task/complete", controller.CompleteTask)
	mux.HandleFunc("/workflow/task/withdraw", controller.WithDrawTask)
	// 配置
	var config = *config.Config
	// 启动数据库连接
	model.Setup()
	// 启动服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("the application start up at port%s", server.Addr)
	// 启动定时任务
	service.CronJobs()
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
