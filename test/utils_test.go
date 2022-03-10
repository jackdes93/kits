package test

import (
	"errors"
	"fmt"
	utils "kits"
	"testing"
	"time"
)

func TestRetryWithTime(t *testing.T) {
	err := utils.RetryWithTimes(8, func(times int) (err error) {
		if times < 3 {
			err = errors.New("FAIL")
		}
		fmt.Printf("TIME at %d\n", times)
		return err
	})
	if err != nil {
		t.Fail()
	}
}

func TestRetryWitDelay(t *testing.T) {
	startTime := time.Now().UnixNano()
	delay := 200 * time.Millisecond
	err := utils.RetryWithDelay(2, delay, func(times int) (err error) {
		if times == 0 {
			err = errors.New("")
		}
		return err
	})
	dt := (time.Now().UnixNano() - startTime - int64(delay)) / (int64(time.Millisecond) / int64(time.Nanosecond))
	maxErr := int64(50)
	if err != nil || (dt < -maxErr || dt > maxErr) {
		t.Fail()
	}
}

type MEssage struct {
	ID   string
	Code int
}

func TestPrintLog(t *testing.T) {
	message := MEssage{ID: "DREAA", Code: 901}
	utils.PrintLog(0, message)
	utils.PrintLog(1, "TESTING")
	utils.PrintLog(2, "TESTING")
	utils.PrintLog(3, "TESTING")
	utils.PrintLog(4, "TESTING")
	utils.PrintLog(5, "TESTING")
}
