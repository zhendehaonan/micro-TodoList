package handlers

import (
	"errors"
	"micro-TodoList/api-gateway/pkg/logging"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}
