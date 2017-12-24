package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoCorporationAdminRequest() *proto.CorporationAdminRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.CorporationAdminRequest))(nil)).Elem()))
	var nullValue *proto.CorporationAdminRequest
	return nullValue
}

func EqPtrToProtoCorporationAdminRequest(value *proto.CorporationAdminRequest) *proto.CorporationAdminRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.CorporationAdminRequest
	return nullValue
}
