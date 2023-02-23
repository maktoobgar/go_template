package logging

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alecthomas/units"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/maktoobgar/go_template/pkg/colors"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
)

var (
	// Four folders that will be created inside the path you
	// give in `New` function for logs
	folderNames = []string{"info", "warning", "error", "panic"}
)

// Struct that will returns in `New` function
type LogBundle struct {
	inf *logrus.Logger
	war *logrus.Logger
	err *logrus.Logger
	pan *logrus.Logger

	infDebug *logrus.Logger
	warDebug *logrus.Logger
	errDebug *logrus.Logger
	panDebug *logrus.Logger

	debug bool
}

// Takes options needed for logs configs and returns
// a *Logger
//
// If no address in `opt.Path` is provided,
// "/var/log/project" address will be used as default
func New(opt *Option, debug bool) (Logger, error) {
	if opt == nil {
		return nil, errors.New("option can not be nil")
	}

	if opt.Path == "" {
		opt.Path = "/var/log/project"
	}
	err := createAddress(opt.Path)
	if err != nil {
		return nil, err
	}

	l := &LogBundle{
		inf:      logrus.New(),
		war:      logrus.New(),
		err:      logrus.New(),
		pan:      logrus.New(),
		infDebug: logrus.New(),
		warDebug: logrus.New(),
		errDebug: logrus.New(),
		panDebug: logrus.New(),
		debug:    debug,
	}

	l.inf.SetFormatter(&logrus.JSONFormatter{})
	l.war.SetFormatter(&logrus.JSONFormatter{})
	l.err.SetFormatter(&logrus.JSONFormatter{})
	l.pan.SetFormatter(&logrus.JSONFormatter{})
	if debug {
		l.infDebug.SetFormatter(&logrus.JSONFormatter{})
		l.warDebug.SetFormatter(&logrus.JSONFormatter{})
		l.errDebug.SetFormatter(&logrus.JSONFormatter{})
		l.panDebug.SetFormatter(&logrus.JSONFormatter{})
		l.infDebug.SetOutput(os.Stdout)
		l.warDebug.SetOutput(os.Stdout)
		l.errDebug.SetOutput(os.Stdout)
		l.panDebug.SetOutput(os.Stdout)
	}

	for i := 0; i < 4; i++ {
		writer, err := getLoggerWriter(opt, &i)
		if err != nil {
			return nil, err
		}

		if i == 0 {
			l.inf.SetOutput(writer)
		} else if i == 1 {
			l.war.SetOutput(writer)
		} else if i == 2 {
			l.err.SetOutput(writer)
		} else if i == 3 {
			l.pan.SetOutput(writer)
		}
	}

	return l, nil
}

// Returns io.Writer for 4 different logs of
// Info, Warning, Error and Panic in passed address
// by `New` function
func getLoggerWriter(opt *Option, i *int) (io.Writer, error) {
	maxAge, err := str2duration.ParseDuration(opt.MaxAge)
	path := filepath.Join(opt.Path, folderNames[*i], opt.Pattern)
	if err != nil {
		return nil, err
	}

	rotationTime, err := str2duration.ParseDuration(opt.RotationTime)
	if err != nil {
		return nil, err
	}

	rotationSize, err := units.ParseBase2Bytes(opt.RotationSize)
	if err != nil {
		return nil, err
	}

	return rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationSize(int64(rotationSize)),
	)
}

// If required addresses for logs to create do not exist, will create them and
// makes sure they are accessable
func createAddress(address string) error {
	var (
		addr    string = ""
		addrSep string = ""
	)

	addrSep = "/"
	createIfDoesNotExist := func(addr *string) error {
		if _, err := os.Stat(*addr); os.IsNotExist(err) {
			var cmd *exec.Cmd

			// Create folders
			if runtime.GOOS == "linux" {
				cmd = exec.Command("sudo", "mkdir", *addr)
			} else {
				cmd = exec.Command("mkdir", *addr)
			}
			_, err := cmd.Output()

			if err != nil {
				return err
			}

			// Change folders owners (just for linux users)
			if runtime.GOOS == "linux" {
				cmd = exec.Command("whoami")
				whoAmIByte, err := cmd.Output()
				whoAmI := strings.Replace(string(whoAmIByte), "\n", "", -1)

				if err != nil {
					return err
				}
				cmd = exec.Command("sudo", "chown", fmt.Sprintf("%s:%s", whoAmI, whoAmI), *addr)
				_, err = cmd.Output()
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	for _, str := range strings.Split(address, addrSep) {
		addr += str + addrSep

		if addr != ("."+addrSep) && addr != addrSep && len(addr) > 4 {
			err := createIfDoesNotExist(&addr)
			if err != nil {
				return err
			}
		}
	}

	for _, folderName := range folderNames {
		t := addr + folderName
		err := createIfDoesNotExist(&(t))
		if err != nil {
			return err
		}
	}

	return nil
}

// Info message log
func (l *LogBundle) Info(message string, r *http.Request, function any, params ...map[string]any) {
	param := map[string]any{}
	if len(params) > 0 {
		param = params[0]
	}
	data := ""
	if dataBytes, errTemp := io.ReadAll(r.Body); errTemp == nil {
		data = string(dataBytes)
	}
	logFields := logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"url":      r.URL.Path,
		"method":   r.Method,
		"body":     data,
		"params":   param,
	}
	l.inf.WithFields(logFields).Info(message)
	if l.debug {
		fmt.Print(colors.Gray)
		l.infDebug.WithFields(logFields).Info(message)
		fmt.Print(colors.Reset)
	}
}

// Warning message log
func (l *LogBundle) Warning(message string, r *http.Request, function any, params ...map[string]any) {
	param := map[string]any{}
	if len(params) > 0 {
		param = params[0]
	}
	data := ""
	if dataBytes, errTemp := io.ReadAll(r.Body); errTemp == nil {
		data = string(dataBytes)
	}
	logFields := logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"url":      r.URL.Path,
		"method":   r.Method,
		"body":     data,
		"params":   param,
	}
	l.war.WithFields(logFields).Warning(message)
	if l.debug {
		fmt.Print(colors.Purple)
		l.warDebug.WithFields(logFields).Warning(message)
		fmt.Print(colors.Reset)
	}
}

// Error message log
func (l *LogBundle) Error(message string, r *http.Request, function any, params ...map[string]any) {
	param := map[string]any{}
	if len(params) > 0 {
		param = params[0]
	}
	data := ""
	if dataBytes, errTemp := io.ReadAll(r.Body); errTemp == nil {
		data = string(dataBytes)
	}
	logFields := logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"url":      r.URL.Path,
		"method":   r.Method,
		"body":     data,
		"params":   param,
	}
	l.err.WithFields(logFields).Error(message)
	if l.debug {
		fmt.Print(colors.Yellow)
		l.errDebug.WithFields(logFields).Error(message)
		fmt.Print(colors.Reset)
	}
}

// Panic message log
func (l *LogBundle) Panic(err any, r *http.Request, stack string, params ...map[string]any) {
	message := fmt.Sprintf("%v", err)
	param := map[string]any{}
	if len(params) > 0 {
		param = params[0]
	}
	data := ""
	if dataBytes, errTemp := io.ReadAll(r.Body); errTemp == nil {
		data = string(dataBytes)
	}
	logFields := logrus.Fields{
		"url":    r.URL.Path,
		"method": r.Method,
		"body":   data,
		"stack":  stack,
		"params": param,
	}
	l.pan.WithFields(logFields).Error(message)
	if l.debug {
		fmt.Print(colors.Red + message + "\n" + stack + colors.Reset)
	}
}
