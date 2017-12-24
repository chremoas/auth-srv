package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAuthCreateResponse() *proto.AuthCreateResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AuthCreateResponse))(nil)).Elem()))
	var nullValue *proto.AuthCreateResponse
	return nullValue
}

func EqPtrToProtoAuthCreateResponse(value *proto.AuthCreateResponse) *proto.AuthCreateResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AuthCreateResponse
	return nullValue
}
