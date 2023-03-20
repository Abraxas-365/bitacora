package myerror

import (
	"fmt"
	"time"
)

type MyError struct {
	err       error
	code      int
	timestamp time.Time
}

type MyErrorJson struct {
	Err       string    `json:"err"`
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

func New(msg string, code int) *MyError {
	return &MyError{
		err:       fmt.Errorf(msg),
		code:      code,
		timestamp: time.Now(),
	}
}

func Wrap(err error, code int) *MyError {
	return &MyError{
		err:       err,
		code:      code,
		timestamp: time.Now(),
	}
}

func (e *MyError) Unwrap() error {
	return e.err
}

func (e *MyError) Error() string {
	return fmt.Sprintf("HTTP %d - %s (%s)", e.code, e.err.Error(), e.timestamp.Format("2006-01-02 15:04:05"))
}

func (e *MyError) Code() int {
	return e.code
}

func (e *MyError) Timestamp() time.Time {
	return e.timestamp
}

func (e *MyError) ToJson() MyErrorJson {
	myErrorJson := MyErrorJson{
		Err:       e.Error(),
		Code:      e.Code(),
		Timestamp: e.Timestamp(),
	}

	return myErrorJson
}
