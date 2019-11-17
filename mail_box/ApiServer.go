package mail_box

import (
	"aaa/mail_box/configs"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	//单例对象
	G_apiServer *ApiServer
)

// 初始化服务
func InitApiServer() (err error) {
	var (
		mux           *http.ServeMux
		listener      net.Listener
		httpServer    *http.Server
		staticDir     http.Dir     // 静态文件根目录
		staticHandler http.Handler // 静态文件的HTTP回调
	)

	// 配置路由
	mux = http.NewServeMux()

	//mail_box
	mux.HandleFunc("/job/mail_box1", HandleAcceptLetterMail)
	mux.HandleFunc("/job/mail_box2", HandleNoLetter)
	mux.HandleFunc("/job/mail_box3", SearchLetter)

	//  /index.html

	// 静态文件目录
	staticDir = http.Dir(configs.G_config.WebRoot)
	staticHandler = http.FileServer(staticDir)
	mux.Handle("/", http.StripPrefix("/", staticHandler)) //   ./webroot/index.html

	// 启动TCP监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(configs.G_config.ApiPort)); err != nil {
		return
	}

	// 创建一个HTTP服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(configs.G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(configs.G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	// 赋值单例
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	// 启动了服务端
	go httpServer.Serve(listener)

	return
}
