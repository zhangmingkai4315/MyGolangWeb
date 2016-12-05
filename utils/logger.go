package utils

import (
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TraceLog   *log.Logger
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func createLoggers(traceHandler, infoHandler, warningHandler, errorHandler io.Writer) {

	TraceLog = log.New(traceHandler, "Trace：", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog = log.New(infoHandler, "Info：", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(warningHandler, "Warning：", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(errorHandler, "Error：", log.Ldate|log.Ltime|log.Lshortfile)

}

func InitLogger() {
	if viper.GetString("env") == "production" {
		//	create log files for production
	} else {
		createLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}
}
