package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoCharacterAdminRequest() *proto.CharacterAdminRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.CharacterAdminRequest))(nil)).Elem()))
	var nullValue *proto.CharacterAdminRequest
	return nullValue
}

func EqPtrToProtoCharacterAdminRequest(value *proto.CharacterAdminRequest) *proto.CharacterAdminRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.CharacterAdminRequest
	return nullValue
}
