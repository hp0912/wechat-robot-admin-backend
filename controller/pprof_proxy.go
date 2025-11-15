package controller

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"

	"wechat-robot-admin-backend/pkg/appx"
)

type PprofProxy struct{}

func NewPprofProxyController() *PprofProxy {
	return &PprofProxy{}
}

func (p *PprofProxy) ProxyPprof(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	pprofPath := c.Param("pprof_path")
	target, err := url.Parse(robot.GetBaseURL())
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = target.Path + "/pprof" + pprofPath
		// 移除id参数
		query := req.URL.Query()
		query.Del("id")
		req.URL.RawQuery = query.Encode()
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("pprof proxy error: " + err.Error()))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
