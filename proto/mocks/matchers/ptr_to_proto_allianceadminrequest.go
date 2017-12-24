package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAllianceAdminRequest() *proto.AllianceAdminRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AllianceAdminRequest))(nil)).Elem()))
	var nullValue *proto.AllianceAdminRequest
	return nullValue
}

func EqPtrToProtoAllianceAdminRequest(value *proto.AllianceAdminRequest) *proto.AllianceAdminRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AllianceAdminRequest
	return nullValue
}
