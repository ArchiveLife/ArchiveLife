package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/ArchiveLife/ui"
	"github.com/ArchiveLife/ui/lib"
	"github.com/webview/webview"
)

func main() {

	http.HandleFunc("/services", lib.APIGetServices)
	http.Handle("/", http.FileServer(http.FS(ui.GetUIContent())))

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port

	fmt.Println(port)

	go http.Serve(listener, nil)
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Archive Life")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(fmt.Sprintf("http://127.0.0.1:%v", port))
	w.Run()

}
