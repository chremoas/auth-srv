package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoRoleAdminRequest() *proto.RoleAdminRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.RoleAdminRequest))(nil)).Elem()))
	var nullValue *proto.RoleAdminRequest
	return nullValue
}

func EqPtrToProtoRoleAdminRequest(value *proto.RoleAdminRequest) *proto.RoleAdminRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.RoleAdminRequest
	return nullValue
}
