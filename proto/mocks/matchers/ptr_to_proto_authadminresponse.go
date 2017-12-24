package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAuthAdminResponse() *proto.AuthAdminResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AuthAdminResponse))(nil)).Elem()))
	var nullValue *proto.AuthAdminResponse
	return nullValue
}

func EqPtrToProtoAuthAdminResponse(value *proto.AuthAdminResponse) *proto.AuthAdminResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AuthAdminResponse
	return nullValue
}
