package subscriber

import (
	"context"

	"github.com/bitrainforest/pulsar/internal/cache"
	"github.com/bitrainforest/pulsar/internal/dao"
)

const (
	DefaultWorkPoolNum = 2000
	MaxWorkPoolNum     = 3000
)

type (
	OptFn func(opts *Opts)
	Opts  struct {
		workPoolNum      int64
		appSubDao        dao.UserAppSubDao
		addressMarkCache cache.AddressMark
	}
)

func defaultOpts() Opts {
	return Opts{
		addressMarkCache: cache.NewAddressMark(context.Background()),
		appSubDao:        dao.NewUserAppSubDao(),
		workPoolNum:      DefaultWorkPoolNum,
	}
}

// WithUserAppSubDao set useAppSubDao
func WithUserAppSubDao(appSub dao.UserAppSubDao) OptFn {
	return func(opts *Opts) {
		if appSub != nil {
			opts.appSubDao = appSub
		}
	}
}

// WithAddressMarkCache set address mark cache
func WithAddressMarkCache(mark cache.AddressMark) OptFn {
	return func(opts *Opts) {
		if mark != nil {
			opts.addressMarkCache = mark
		}
	}
}

// WithWorkPoolNum set work pool num
func WithWorkPoolNum(num int64) OptFn {
	return func(opts *Opts) {
		opts.workPoolNum = num
	}
}
