package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	//"net/http/pprof"
	"flag"
	"github.com/golang/glog"
	"io"
)

func rootHandle(w http.ResponseWriter, r *http.Request)  {
	glog.V(2).Info(fmt.Sprintf("请求路径:[%s]  请求主机:[%s]",r.URL,r.Host))
	responseStr := "请求路径如下:\n/healthz: 健康检查\n/test:测试\n"
	io.WriteString(w,responseStr)
}

func healthzHandler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"200\n")
}
func testHandler(w http.ResponseWriter , r *http.Request)  {
	//获取主机名并返回
	hosntame ,err := os.Hostname()
	if err == nil {
		w.Header().Set("hostname",hosntame)
	}

	rawEnv := os.Environ()
	for _,curRawEnv := range rawEnv {

		k1 := strings.Split(curRawEnv,"=")[0]
		w.Header().Set(fmt.Sprintf("%s",k1),fmt.Sprintf("%s",os.Getenv(k1)))
		glog.V(2).Info(fmt.Sprintf("k:%s  v:%s\n",k1,os.Getenv(k1)))
	}


	for k,v := range r.Header {
		glog.V(2).Info(fmt.Sprintf("k:%s  v:%s\n",k,v))
		w.Header().Set(k,fmt.Sprintf("%s",v))
	}
	io.WriteString(w,"测试接口\n")

}

func main()  {
	flag.Set("v","4")
	flag.Set("log_dir","logs")
	port := flag.String("port","8000"," default port ")

	flag.Parse()
	defer glog.Flush()

	glog.V(2).Info("http  server  starting ......")

	http.HandleFunc("/",rootHandle)
	http.HandleFunc("/healthz",healthzHandler)
	http.HandleFunc("/test",testHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%s",*port),nil)
	if err != nil {
		log.Fatal(err)
	}

}
