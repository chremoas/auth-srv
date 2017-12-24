package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoEntityQueryRequest() *proto.EntityQueryRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.EntityQueryRequest))(nil)).Elem()))
	var nullValue *proto.EntityQueryRequest
	return nullValue
}

func EqPtrToProtoEntityQueryRequest(value *proto.EntityQueryRequest) *proto.EntityQueryRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.EntityQueryRequest
	return nullValue
}
