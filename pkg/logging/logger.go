package logging

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alecthomas/units"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
)

var (
	// Four folders that will be created inside the path you
	// give in `New` function for logs
	folderNames                    = []string{"info", "warning", "error", "panic"}
	errOperationSystemNotSupported = fmt.Errorf("%s operation system not supported", runtime.GOOS)
)

// Struct that will returns in `New` function
type logBundle struct {
	inf *logrus.Logger
	war *logrus.Logger
	err *logrus.Logger
	pan *logrus.Logger
}

// Takes options needed for logs configs and returns
// a *Logger
//
// If no address in `opt.Path` is provided, for linux users:
// "/var/log/project" and for windows users "c:\\logs\project"
// address will be used as default path
func New(opt *Option) (Logger, error) {
	if opt == nil {
		return nil, errors.New("option can not be nil")
	}

	if opt.Path == "" {
		if runtime.GOOS == "linux" {
			opt.Path = "/var/log/project"
		} else if runtime.GOOS == "windows" {
			opt.Path = "c:\\\\logs\\project"
		} else {
			return nil, errOperationSystemNotSupported
		}
	}
	err := createAddress(opt.Path)
	if err != nil {
		return nil, err
	}

	l := &logBundle{
		inf: logrus.New(),
		war: logrus.New(),
		err: logrus.New(),
		pan: logrus.New(),
	}

	l.inf.SetFormatter(&logrus.TextFormatter{})
	l.war.SetFormatter(&logrus.TextFormatter{})
	l.err.SetFormatter(&logrus.TextFormatter{})
	l.pan.SetFormatter(&logrus.TextFormatter{})

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

	if runtime.GOOS == "windows" {
		addrSep = "\\"
	} else if runtime.GOOS == "linux" {
		addrSep = "/"
	} else {
		return errOperationSystemNotSupported
	}

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
//
// Put your message in `message` variable and write your current function in `function`
// and if you need to print some more parameters, put them inside `params` variable
func (l *logBundle) Info(message string, function interface{}, params map[string]interface{}) {
	l.inf.WithFields(logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"params":   params,
	}).Info(message)
}

// Warning message log
//
// Put your message in `message` variable and write your current function in `function`
// and if you need to print some more parameters, put them inside `params` variable
func (l *logBundle) Warning(message string, function interface{}, params map[string]interface{}) {
	l.war.WithFields(logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"params":   params,
	}).Warning(message)
}

// Error message log
//
// Put your message in `message` variable and write your current function in `function`
// and if you need to print some more parameters, put them inside `params` variable
func (l *logBundle) Error(message string, function interface{}, params map[string]interface{}) {
	l.err.WithFields(logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"params":   params,
	}).Error(message)
}

// Panic message log
//
// Put your message in `message` variable and write your current function in `function`
// and if you need to print some more parameters, put them inside `params` variable
func (l *logBundle) Panic(message string, function interface{}, params map[string]interface{}) {
	l.pan.WithFields(logrus.Fields{
		"package":  getPackageName(function),
		"function": getFunctionName(function),
		"params":   params,
	}).Panic(message)
}
