package main

import (
	"log"
	"os"

	"github.com/op/go-logging"
)

var logYao = logging.MustGetLogger("yao")
var logLu = logging.MustGetLogger("lu")
var yaoFormat = logging.MustStringFormatter(
	`%{color} => %{level:.4s} %{id:04x}%{color} %{message}`,
)
var luFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {

	// init
	// f, _ := os.OpenFile("/data/logs/yao.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// defer f.Close()
	// yaoBackend := logging.NewLogBackend(f, "[Yao]", log.Ldate|log.Ltime)
	yaoBackend := logging.NewLogBackend(os.Stderr, "[Yao]", log.Ldate|log.Ltime)
	luBackend := logging.NewLogBackend(os.Stdout, "[Lu]", 0)

	luBackendLevel := logging.AddModuleLevel(luBackend)
	luBackendLevel.SetLevel(logging.ERROR, "lu")
	// luBackendFmt := logging.NewBackendFormatter(luBackend, luFormat)

	yaoBackendFmt := logging.NewBackendFormatter(yaoBackend, yaoFormat)
	yaoBackendLevel := logging.AddModuleLevel(yaoBackendFmt)
	yaoBackendLevel.SetLevel(logging.ERROR, "yao")
	logging.SetBackend(yaoBackendLevel)

	// log
	logYao.Debugf("yao-debug %s", logging.Redact("yao-secret"))
	logYao.Info("yao-info")
	logYao.Notice("yao-notice")
	logYao.Warning("yao-warning")
	logYao.Error("yao-err")
	logYao.Critical("yao-crit")

	logLu.Debugf("lu-debug %s", logging.Redact("lu-secret"))
	logLu.Info("lu-info")
	logLu.Notice("lu-notice")
	logLu.Warning("lu-warning")
	logLu.Error("lu-err")
	logLu.Critical("lu-crit")
}
