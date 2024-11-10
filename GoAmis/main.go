package main

import (
	"github.com/gin-gonic/gin"
)

// 使用Node.js如何做？
// https://github.com/ZhongFuCheng3y/austin-admin/blob/master/server.js

func main() {
	r := gin.Default()

	// 托管静态文件夹 `/public`, `/pages`, `/sdk`, `/history`
	r.Static("/public", "./public")
	r.Static("/pages", "./pages")
	r.Static("/sdk", "./sdk")
	r.Static("/history", "./history")

	// 设置根路径返回 `index.html`
	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	r.Run(":3000") // 默认监听并在 http://localhost:3000 提供服务
}

//func main() {
//    // 设置静态文件夹，托管 `/public`, `/pages`, `/sdk` 等静态资源
//    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
//    http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages"))))
//    http.Handle("/sdk/", http.StripPrefix("/sdk/", http.FileServer(http.Dir("./sdk"))))
//    http.Handle("/history/", http.StripPrefix("/history/", http.FileServer(http.Dir("./history"))))
//
//    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//        http.ServeFile(w, r, "./index.html")
//    })
//
//    // 设置端口
//    port := ":3000"
//    log.Printf("Server running at http://localhost%s\n", port)
//
//    // 启动服务器
//    err := http.ListenAndServe(port, nil)
//    if err != nil {
//        log.Fatalf("Failed to start server: %v", err)
//    }
//}
