package utils

import (
	"fmt"
	"kits/config"
	"log"
	"time"
)

/** Retry Functions **/
type Handler func(int) error

func RetryWithTimes(times int, f Handler) (err error) {
	for i := 0; i < times; i++ {
		err = f(i)
		if err == nil {
			return nil
		}
	}
	return err
}

func RetryWithDelay(times int, delay time.Duration, f Handler) (err error) {
	for i := 0; i < times; i++ {
		err = f(i)
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return err
}

func RetryInterval(f Handler) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
	}
}

func RetryIntervalWithDelay(delay time.Duration, f Handler) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(delay)
	}
}

/** Log Functions **/

func PrintLog(typelog int, mes interface{}) {
	typeEvent := ""
	color := ""

	switch typelog {
	case 0:
		typeEvent = "[SUCCESS]"
		color = config.Success
	case 1:
		typeEvent = "[WARN]"
		color = config.Warn
	case 2:
		typeEvent = "[ERROR]"
		color = config.Error
	case 3:
		typeEvent = "[DEBUG]"
		color = config.Debug
	case 4:
		typeEvent = "[INFO]"
		color = config.Info
	default:
		typeEvent = "[UNKNOW]"
		color = config.Unknow
	}
	content := fmt.Sprintf("%s:%+v\n", typeEvent, mes)
	log.Printf(color, content)
}
