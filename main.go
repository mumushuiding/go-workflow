package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	model "github.com/mumushuiding/go-workflow/workflow-engine/model"
	"github.com/mumushuiding/Activiti-go/config"
	"github.com/mumushuiding/Activiti-go/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Index)
	//-------------------------流程定义----------------------
	mux.HandleFunc("/procdef/save", controller.SaveProcdef)
	mux.HandleFunc("/procdef/findAll", controller.FindAllProcdefPage)
	mux.HandleFunc("/procdef/delById", controller.DelProcdefByID)
	// -----------------------流程实例-----------------------
	mux.HandleFunc("/process/start", controller.StartProcessInstance)
	mux.HandleFunc("/process/findTask", controller.FindMyProcInstPageAsJSON)
	// -----------------------任务--------------------------
	mux.HandleFunc("/task/complete", controller.CompleteTask)
	mux.HandleFunc("/task/withdraw", controller.WithDrawTask)
	// 配置
	var config = *config.Config
	// 启动数据库连接
	model.Setup()
	// 启动服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("the application start up at port%s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
