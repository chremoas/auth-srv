package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/auth-srv/proto"
)

func AnyPtrToProtoAlliancesResponse() *proto.AlliancesResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.AlliancesResponse))(nil)).Elem()))
	var nullValue *proto.AlliancesResponse
	return nullValue
}

func EqPtrToProtoAlliancesResponse(value *proto.AlliancesResponse) *proto.AlliancesResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.AlliancesResponse
	return nullValue
}
