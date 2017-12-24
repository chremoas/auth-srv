package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoEntityAdminResponse() *proto.EntityAdminResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.EntityAdminResponse))(nil)).Elem()))
	var nullValue *proto.EntityAdminResponse
	return nullValue
}

func EqPtrToProtoEntityAdminResponse(value *proto.EntityAdminResponse) *proto.EntityAdminResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.EntityAdminResponse
	return nullValue
}
