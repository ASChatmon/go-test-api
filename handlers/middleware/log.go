package middleware

import (
	"net/http"
	//"encoding/base64"
	"github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"go-test-api/config"
	"go-test-api/types"
	//"strings"
)

type gojiMiddleware struct {
	h http.Handler
	a *config.Config
}

func (g gojiMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := g.a.Log.Log
	transactionId := r.Header.Get("X-Transaction-Id")

	if transactionId == "" {
		transactionId = uuid.NewV4().String()
		g.a.Log.LogInfo("ServerHttp", "create transaction", "No transaction header sent, using "+transactionId)
	}

	/*
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
	*/
	g.a.Log = &types.Logger{Log: logger.WithFields(
		logrus.Fields{"transaction_id": transactionId,
			"uri":    r.RequestURI,
			"method": r.Method,
			//"remote": r.RemoteAddr,
			//"user":   pair[0],
		})}
	w.Header().Add("X-Transaction-Id", transactionId)
	g.h.ServeHTTP(w, r)
}

func LogWithTransactionId(a *config.Config) func(http.Handler) http.Handler {
	fn := func(h http.Handler) http.Handler {
		return gojiMiddleware{h, a}
	}
	return fn
}
