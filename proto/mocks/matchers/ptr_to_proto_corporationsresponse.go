package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoCorporationsResponse() *proto.CorporationsResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.CorporationsResponse))(nil)).Elem()))
	var nullValue *proto.CorporationsResponse
	return nullValue
}

func EqPtrToProtoCorporationsResponse(value *proto.CorporationsResponse) *proto.CorporationsResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.CorporationsResponse
	return nullValue
}
