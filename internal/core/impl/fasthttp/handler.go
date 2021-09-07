package impl

import (
	"net/http"

	"github.com/RedAFD/mega/internal/core/context"
	"github.com/RedAFD/mega/internal/utils/i18n"
	"github.com/RedAFD/mega/internal/utils/logger"
	"github.com/RedAFD/mega/third_party/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func init() {
	swagger.Handler = func(ctx context.Context) {

		fctx := ctx.(*_coreContextFasthttp).RequestCtx

		var r http.Request
		if err := fasthttpadaptor.ConvertRequest(fctx, &r, true); err != nil {
			logger.Error("cannot parse requestURI %q: %v", r.RequestURI, err)
			ctx.SetRespCode(fasthttp.StatusInternalServerError, i18n.Sprintf("服务器错误，请重新尝试"))
			return
		}

		var w netHTTPResponseWriter
		httpSwagger.WrapHandler.ServeHTTP(&w, r.WithContext(fctx))

		ctx.SetRespCode(w.StatusCode(), w.body)

		haveContentType := false
		for k, vv := range w.Header() {
			if k == fasthttp.HeaderContentType {
				haveContentType = true
			}

			for _, v := range vv {
				ctx.SetRespHeader(k, v)
			}
		}
		if !haveContentType {
			// From net/http.ResponseWriter.Write:
			// If the Header does not contain a Content-Type line, Write adds a Content-Type set
			// to the result of passing the initial 512 bytes of written data to DetectContentType.
			l := 512
			if len(w.body) < 512 {
				l = len(w.body)
			}

			ctx.SetRespContentType(http.DetectContentType(w.body[:l]))
		}
	}
}

type netHTTPResponseWriter struct {
	statusCode int
	h          http.Header
	body       []byte
}

func (w *netHTTPResponseWriter) StatusCode() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}

func (w *netHTTPResponseWriter) Header() http.Header {
	if w.h == nil {
		w.h = make(http.Header)
	}
	return w.h
}

func (w *netHTTPResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *netHTTPResponseWriter) Write(p []byte) (int, error) {
	w.body = append(w.body, p...)
	return len(p), nil
}
