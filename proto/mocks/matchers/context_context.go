package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	context "golang.org/x/net/context"
)

func AnyContextContext() context.Context {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(context.Context))(nil)).Elem()))
	var nullValue context.Context
	return nullValue
}

func EqContextContext(value context.Context) context.Context {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue context.Context
	return nullValue
}
