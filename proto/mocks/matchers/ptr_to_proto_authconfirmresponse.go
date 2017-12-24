package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAuthConfirmResponse() *proto.AuthConfirmResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AuthConfirmResponse))(nil)).Elem()))
	var nullValue *proto.AuthConfirmResponse
	return nullValue
}

func EqPtrToProtoAuthConfirmResponse(value *proto.AuthConfirmResponse) *proto.AuthConfirmResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AuthConfirmResponse
	return nullValue
}
