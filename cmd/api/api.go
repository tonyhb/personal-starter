package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"gitlab.com/tonyhb/keepupdated/pkg/api"
)

func main() {
	api := api.New(api.Opts{
		Log: logrus.StandardLogger(),
	})
	http.ListenAndServe(":80", api.Handler())
}
