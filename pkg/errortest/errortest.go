package errortest

import (
	"github.com/cockroachdb/errors"
)

// PkgError 在 pkg 目錄中產生錯誤
func PkgError() error {
	return errors.New("error from pkg directory")
}

// PkgWrapError 在 pkg 目錄中包裝錯誤
func PkgWrapError(err error) error {
	return errors.Wrapf(err, "wrapped in pkg directory")
}
