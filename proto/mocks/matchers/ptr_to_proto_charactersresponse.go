package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoCharactersResponse() *proto.CharactersResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.CharactersResponse))(nil)).Elem()))
	var nullValue *proto.CharactersResponse
	return nullValue
}

func EqPtrToProtoCharactersResponse(value *proto.CharactersResponse) *proto.CharactersResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.CharactersResponse
	return nullValue
}
