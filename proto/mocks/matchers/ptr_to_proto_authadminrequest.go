package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAuthAdminRequest() *proto.AuthAdminRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AuthAdminRequest))(nil)).Elem()))
	var nullValue *proto.AuthAdminRequest
	return nullValue
}

func EqPtrToProtoAuthAdminRequest(value *proto.AuthAdminRequest) *proto.AuthAdminRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AuthAdminRequest
	return nullValue
}
