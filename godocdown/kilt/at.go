// This file was AUTOMATICALLY GENERATED by kilt-import (smuggol) from github.com/robertkrimen/kilt

package kilt

import (
	"os"
)

// At
func (self Kilt) At(there string, fn func() error) []error {
	return At(there, fn)
}

func At(there string, fn func() error) (err []error) {
	origin, err_ := os.Getwd()
	if err_ != nil {
		origin = ""
		return []error{err_, nil}
	}
	err_ = os.Chdir(there)
	if err_ != nil {
		origin = ""
		return []error{err_, nil}
	}
	defer func() {
		if origin != "" {
			err_ = os.Chdir(origin)
			if err != nil {
				err[0] = err_
			} else {
				err = []error{err_, nil}
			}
		}
	}()
	err_ = fn()
	if err_ != nil {
		return []error{nil, err_}
	}
	return nil
}
