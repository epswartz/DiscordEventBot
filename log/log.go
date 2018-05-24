// This package serves as a wrapper around go-logging, so that I can use it from anywhere without making more loggers.

package log

import (
	"DiscordEventBot/config"
	"github.com/op/go-logging"
	"os"
)

var logger = logging.MustGetLogger("DiscordEventBot") // Create logger
var format = logging.MustStringFormatter(          // Format string for the logger.
	// `%{color}%{time:15:04:05.000} %{shortfunc} %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	`%{color}[%{time:2006-01-02T15:04:05}] %{level:.4s} %{id:03x} > %{color:reset} %{message} `,
)

var Debug = logger.Debug
var Info = logger.Info
var Notice = logger.Notice
var Warning = logger.Warning
var Error = logger.Error
var Critical = logger.Critical

func init(){
	// Initialize the logger's backend so it has somewhere to send messages.
	loggingBck := logging.NewLogBackend(os.Stdout, "", 0)
	loggingBackendFormatter := logging.NewBackendFormatter(loggingBck, format) // Format using the string we made earlier.
	logging.SetBackend(loggingBackendFormatter) // Set the backends to be used by the logger.

	// If verbose logging is off we set the debug function to something lame
	if(!config.Cfg.VerboseLogging){
		Debug = func(a ...interface{}){
			// Do nothing, this shuts debug up
		}
	}

}