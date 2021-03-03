package newrelic

import (
	"net/http"
	"net/url"

	"github.com/gogearbox/gearbox"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/valyala/fasthttp"
)

type handler struct {
	repanic bool
	app     *newrelic.Application
}

// Options struct holds newrelic middleware settings
type Options struct {
	// Repanic configures whether newrelic should repanic after recovery
	Repanic bool
}

// New returns middleware handler
func New(app *newrelic.Application, options ...Options) func(ctx gearbox.Context) {
	var op Options
	if len(options) > 0 {
		op = options[0]
	}

	handler := &handler{
		repanic: op.Repanic,
		app:     app,
	}

	return handler.handle
}

func (h *handler) handle(ctx gearbox.Context) {
	name := string(ctx.Context().Method()) + " " + string(ctx.Context().Path())
	txn := h.app.StartTransaction(name)

	txn.SetWebRequestHTTP(convert(ctx.Context()))

	defer h.recoverWithNewrelic(txn)

	ctx.Next()
}

func (h *handler) recoverWithNewrelic(txn *newrelic.Transaction) {
	err := recover()
	if err != nil {
		switch err := err.(type) {
		case error:
			txn.NoticeError(err)
		default:
			txn.NoticeError(errWrapper{err})
		}
	}

	txn.End()

	if h.repanic && err != nil {
		panic(err)
	}
}

// convert converts fasthttp request to net/http request
func convert(ctx *fasthttp.RequestCtx) *http.Request {
	r := new(http.Request)

	r.Method = string(ctx.Method())

	uri := ctx.URI()
	r.URL = &url.URL{
		Scheme:   string(uri.Scheme()),
		Path:     string(uri.Path()),
		Host:     string(uri.Host()),
		RawQuery: string(uri.QueryString()),
	}

	// Headers
	r.Header = make(http.Header)
	r.Header.Add("Host", string(ctx.Host()))
	ctx.Request.Header.VisitAll(func(key, value []byte) {
		r.Header.Add(string(key), string(value))
	})

	return r
}
