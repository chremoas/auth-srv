package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAuthConfirmRequest() *proto.AuthConfirmRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AuthConfirmRequest))(nil)).Elem()))
	var nullValue *proto.AuthConfirmRequest
	return nullValue
}

func EqPtrToProtoAuthConfirmRequest(value *proto.AuthConfirmRequest) *proto.AuthConfirmRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AuthConfirmRequest
	return nullValue
}
