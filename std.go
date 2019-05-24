package stdlog

import (
	"log"
	"os"
)

type stdLogger struct {
	infoL  *log.Logger
	debugL *log.Logger
	warnL  *log.Logger
	errorL *log.Logger
}

var Std = &stdLogger{
	infoL:  log.New(os.Stdout, prefixInfo, log.LstdFlags),
	debugL: log.New(os.Stdout, prefixDebug, log.LstdFlags),
	warnL:  log.New(os.Stderr, prefixWarn, log.LstdFlags),
	errorL: log.New(os.Stderr, prefixError, log.LstdFlags),
}

func (s *stdLogger) Infof(format string, args ...interface{}) {
	s.infoL.Printf(format, args...)
}

func (s *stdLogger) Debugf(format string, args ...interface{}) {
	s.debugL.Printf(format, args...)
}

func (s *stdLogger) Warnf(format string, args ...interface{}) {
	s.warnL.Printf(format, args...)
}

func (s *stdLogger) Errorf(format string, args ...interface{}) {
	s.errorL.Printf(format, args...)
}

func (s *stdLogger) Panicf(format string, args ...interface{}) {
	s.errorL.Panicf(format, args...)
}
