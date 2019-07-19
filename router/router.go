package router

import (
	"net/http"

	"github.com/mumushuiding/workdiary/config"
	"github.com/mumushuiding/workdiary/controller"
)

// Mux 路由
var Mux = http.NewServeMux()
var conf = *config.Config

func init() {
	setMux()
}
func crossOrigin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", conf.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", conf.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", conf.AccessControlAllowHeaders)
		h(w, r)
	}
}
func setMux() {
	Mux.HandleFunc("/api/v1/workdiary/index", controller.Index)
	// -------------------------- 项目 ----------------
	Mux.HandleFunc("/api/v1/workdiary/project/save", crossOrigin(controller.SaveProject))
	Mux.HandleFunc("/api/v1/workdiary/project/findall", crossOrigin(controller.FindAllProject))
}
