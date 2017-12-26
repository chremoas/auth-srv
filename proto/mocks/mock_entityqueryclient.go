// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/chremoas/auth-srv/proto (interfaces: EntityQueryClient)

package authsrv_mocks

import (
	context "golang.org/x/net/context"
	proto "github.com/chremoas/auth-srv/proto"
	client "github.com/micro/go-micro/client"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockEntityQueryClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockEntityQueryClient() *MockEntityQueryClient {
	return &MockEntityQueryClient{fail: pegomock.GlobalFailHandler}
}

func (mock *MockEntityQueryClient) GetAlliances(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) (*proto.AlliancesResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetAlliances", params, []reflect.Type{reflect.TypeOf((**proto.AlliancesResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.AlliancesResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.AlliancesResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityQueryClient) GetCharacters(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) (*proto.CharactersResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetCharacters", params, []reflect.Type{reflect.TypeOf((**proto.CharactersResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.CharactersResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.CharactersResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityQueryClient) GetCorporations(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) (*proto.CorporationsResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetCorporations", params, []reflect.Type{reflect.TypeOf((**proto.CorporationsResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.CorporationsResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.CorporationsResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityQueryClient) GetRoles(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) (*proto.RoleResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetRoles", params, []reflect.Type{reflect.TypeOf((**proto.RoleResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.RoleResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.RoleResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityQueryClient) VerifyWasCalledOnce() *VerifierEntityQueryClient {
	return &VerifierEntityQueryClient{mock, pegomock.Times(1), nil}
}

func (mock *MockEntityQueryClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierEntityQueryClient {
	return &VerifierEntityQueryClient{mock, invocationCountMatcher, nil}
}

func (mock *MockEntityQueryClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierEntityQueryClient {
	return &VerifierEntityQueryClient{mock, invocationCountMatcher, inOrderContext}
}

type VerifierEntityQueryClient struct {
	mock                   *MockEntityQueryClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierEntityQueryClient) GetAlliances(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) *EntityQueryClient_GetAlliances_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetAlliances", params)
	return &EntityQueryClient_GetAlliances_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityQueryClient_GetAlliances_OngoingVerification struct {
	mock              *MockEntityQueryClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityQueryClient_GetAlliances_OngoingVerification) GetCapturedArguments() (context.Context, *proto.EntityQueryRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityQueryClient_GetAlliances_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.EntityQueryRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.EntityQueryRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.EntityQueryRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}

func (verifier *VerifierEntityQueryClient) GetCharacters(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) *EntityQueryClient_GetCharacters_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetCharacters", params)
	return &EntityQueryClient_GetCharacters_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityQueryClient_GetCharacters_OngoingVerification struct {
	mock              *MockEntityQueryClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityQueryClient_GetCharacters_OngoingVerification) GetCapturedArguments() (context.Context, *proto.EntityQueryRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityQueryClient_GetCharacters_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.EntityQueryRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.EntityQueryRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.EntityQueryRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}

func (verifier *VerifierEntityQueryClient) GetCorporations(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) *EntityQueryClient_GetCorporations_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetCorporations", params)
	return &EntityQueryClient_GetCorporations_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityQueryClient_GetCorporations_OngoingVerification struct {
	mock              *MockEntityQueryClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityQueryClient_GetCorporations_OngoingVerification) GetCapturedArguments() (context.Context, *proto.EntityQueryRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityQueryClient_GetCorporations_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.EntityQueryRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.EntityQueryRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.EntityQueryRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}

func (verifier *VerifierEntityQueryClient) GetRoles(_param0 context.Context, _param1 *proto.EntityQueryRequest, _param2 ...client.CallOption) *EntityQueryClient_GetRoles_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetRoles", params)
	return &EntityQueryClient_GetRoles_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityQueryClient_GetRoles_OngoingVerification struct {
	mock              *MockEntityQueryClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityQueryClient_GetRoles_OngoingVerification) GetCapturedArguments() (context.Context, *proto.EntityQueryRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityQueryClient_GetRoles_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.EntityQueryRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.EntityQueryRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.EntityQueryRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}
