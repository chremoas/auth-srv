package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoGetRolesRequest() *proto.GetRolesRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetRolesRequest))(nil)).Elem()))
	var nullValue *proto.GetRolesRequest
	return nullValue
}

func EqPtrToProtoGetRolesRequest(value *proto.GetRolesRequest) *proto.GetRolesRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetRolesRequest
	return nullValue
}
