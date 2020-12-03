//
//  logging.go
//  go-sdk
//
//  Copyright © 2019 Halim. All rights reserved.
//

package logging

import (
	"bytes"
	"fmt"
	"os"

	opLogging "github.com/op/go-logging"
)

var log = opLogging.MustGetLogger("go-sdk")

const INTERNAL = "INTERNAL"

func MustGetLogger(name string) *GoSDK {
	host, err := os.Hostname()
	if err != nil {
		log.Error(INTERNAL, err.Error())
		host = "unknown"
	}

	format := opLogging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{color}%{message}%{color:reset}`)
	backend := opLogging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := opLogging.NewBackendFormatter(backend, format)
	opLogging.SetBackend(backendFormatter)

	ppl := &GoSDK{opLogging.MustGetLogger(name), host}
	ppl.ExtraCalldepth = 1
	return ppl
}

type GoSDK struct {
	*opLogging.Logger
	Hostname string
}

/*func (ppl *GoSDK) Debug(logID string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Debug("debug %s", args...)
}*/

func (ppl *GoSDK) Debugf(logID string, stringFormat string, args ...interface{}) {
	ppl.Logger.Debugf("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Info(logID string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Info("info", args)
}

func (ppl *GoSDK) Infof(logID string, stringFormat string, args ...interface{}) {
	ppl.Logger.Infof("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Error(logID string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			"", "ERROR", args))
	}
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Error("err", args)
}

func (ppl *GoSDK) Errorf(logID string, stringFormat string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			stringFormat, "ERROR", args))
	}
	ppl.Logger.Errorf("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Critical(logID string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			"", "CRITICAL", args))
	}
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Critical("crit", args)
}

func (ppl *GoSDK) Criticalf(logID string, stringFormat string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			stringFormat, "CRITICAL", args))
	}
	ppl.Logger.Criticalf("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Fatal(logID string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			"", "FATAL", args))
	}
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Fatal(args)
}

func (ppl *GoSDK) Fatalf(logID string, stringFormat string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			stringFormat, "FATAL", args))
	}
	ppl.Logger.Fatalf("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Panic(logID string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			"", "PANIC", args))
	}
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Panic(args)
}

func (ppl *GoSDK) Panicf(logID string, stringFormat string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			stringFormat, "PANIC", args))
	}
	ppl.Logger.Panicf("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Warning(logID string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			"", "WARNING", args))
	}
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Warning("warning", args)
}

func (ppl *GoSDK) Warningf(logID string, stringFormat string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			stringFormat, "WARNING", args))
	}
	ppl.Logger.Warningf("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Notice(logID string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			"", "NOTICE", args))
	}
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Notice("notice", args)
}

func (ppl *GoSDK) Noticef(logID string, stringFormat string, args ...interface{}) {
	if GlobalLogMessage != nil {
		GlobalLogMessage.SendMessage(messageFormat(GlobalLogMessage.Environment, GlobalLogMessage.ServiceName,
			stringFormat, "NOTICE", args))
	}
	ppl.Logger.Noticef("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func messageFormat(environment string, logID string, stringFormat string, messageType string, args ...interface{}) string {

	var bufferText bytes.Buffer
	for _, text := range args {
		bufferText.WriteString(fmt.Sprintf("%s", text))
	}

	message := fmt.Sprintf("[" + environment + "] [" + logID + "]" +
		stringFormat + " " + messageType + " " + bufferText.String())
	return message
}
