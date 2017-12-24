package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAuthCreateRequest() *proto.AuthCreateRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AuthCreateRequest))(nil)).Elem()))
	var nullValue *proto.AuthCreateRequest
	return nullValue
}

func EqPtrToProtoAuthCreateRequest(value *proto.AuthCreateRequest) *proto.AuthCreateRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AuthCreateRequest
	return nullValue
}
