package controller

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"wechat-robot-admin-backend/pkg/appx"
)

type PprofProxy struct{}

var hrefRegexp = regexp.MustCompile(`href=(["'])([^"']+)(["'])`)

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
	proxy.ModifyResponse = func(res *http.Response) error {
		if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
			return nil
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		_ = res.Body.Close()
		body = rewritePprofHTMLLinks(body, c.Request.URL.Path, robot.ID)
		res.Body = io.NopCloser(bytes.NewReader(body))
		res.ContentLength = int64(len(body))
		res.Header.Set("Content-Length", strconv.Itoa(len(body)))
		return nil
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("pprof proxy error: " + err.Error()))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func rewritePprofHTMLLinks(body []byte, requestPath string, robotID int64) []byte {
	basePath := "/api/v1/pprof/debug/pprof/"
	if prefix, _, found := strings.Cut(requestPath, "/debug/pprof"); found {
		basePath = strings.TrimRight(prefix, "/") + "/debug/pprof/"
	}
	robotIDText := strconv.FormatInt(robotID, 10)

	return hrefRegexp.ReplaceAllFunc(body, func(match []byte) []byte {
		parts := hrefRegexp.FindSubmatch(match)
		if len(parts) != 4 {
			return match
		}
		rewritten, ok := rewritePprofHref(string(parts[2]), basePath, robotIDText)
		if !ok {
			return match
		}
		return []byte("href=" + string(parts[1]) + rewritten + string(parts[3]))
	})
}

func rewritePprofHref(rawHref string, basePath string, robotID string) (string, bool) {
	parsed, err := url.Parse(rawHref)
	if err != nil || parsed.IsAbs() || parsed.Host != "" {
		return "", false
	}
	profilePath := parsed.Path
	if _, suffix, found := strings.Cut(profilePath, "/debug/pprof"); found {
		profilePath = suffix
	} else if strings.HasPrefix(profilePath, "/") {
		return "", false
	}
	profilePath = strings.TrimLeft(profilePath, "/")

	query := parsed.Query()
	query.Set("id", robotID)
	rewritten := basePath + profilePath
	if encodedQuery := query.Encode(); encodedQuery != "" {
		rewritten += "?" + encodedQuery
	}
	if parsed.Fragment != "" {
		rewritten += "#" + parsed.Fragment
	}
	return rewritten, true
}
