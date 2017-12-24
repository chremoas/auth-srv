package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoRoleResponse() *proto.RoleResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.RoleResponse))(nil)).Elem()))
	var nullValue *proto.RoleResponse
	return nullValue
}

func EqPtrToProtoRoleResponse(value *proto.RoleResponse) *proto.RoleResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.RoleResponse
	return nullValue
}
