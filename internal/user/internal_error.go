package user

import (
	"github.com/cockroachdb/errors"
)

// InternalError 在 internal 目錄中產生錯誤
func InternalError() error {
	return errors.New("error from internal directory")
}

// InternalWrapError 在 internal 目錄中包裝錯誤
func InternalWrapError(err error) error {
	return errors.Wrapf(err, "wrapped in internal directory")
}
