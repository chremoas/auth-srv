// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/chremoas/auth-srv/proto (interfaces: EntityAdminClient)

package authsrv_mocks

import (
	context "golang.org/x/net/context"
	proto "github.com/chremoas/auth-srv/proto"
	client "github.com/micro/go-micro/client"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockEntityAdminClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockEntityAdminClient() *MockEntityAdminClient {
	return &MockEntityAdminClient{fail: pegomock.GlobalFailHandler}
}

func (mock *MockEntityAdminClient) AllianceUpdate(_param0 context.Context, _param1 *proto.AllianceAdminRequest, _param2 ...client.CallOption) (*proto.EntityAdminResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AllianceUpdate", params, []reflect.Type{reflect.TypeOf((**proto.EntityAdminResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.EntityAdminResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.EntityAdminResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityAdminClient) CharacterUpdate(_param0 context.Context, _param1 *proto.CharacterAdminRequest, _param2 ...client.CallOption) (*proto.EntityAdminResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CharacterUpdate", params, []reflect.Type{reflect.TypeOf((**proto.EntityAdminResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.EntityAdminResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.EntityAdminResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityAdminClient) CorporationUpdate(_param0 context.Context, _param1 *proto.CorporationAdminRequest, _param2 ...client.CallOption) (*proto.EntityAdminResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CorporationUpdate", params, []reflect.Type{reflect.TypeOf((**proto.EntityAdminResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.EntityAdminResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.EntityAdminResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityAdminClient) RoleUpdate(_param0 context.Context, _param1 *proto.RoleAdminRequest, _param2 ...client.CallOption) (*proto.EntityAdminResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RoleUpdate", params, []reflect.Type{reflect.TypeOf((**proto.EntityAdminResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.EntityAdminResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.EntityAdminResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEntityAdminClient) VerifyWasCalledOnce() *VerifierEntityAdminClient {
	return &VerifierEntityAdminClient{mock, pegomock.Times(1), nil}
}

func (mock *MockEntityAdminClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierEntityAdminClient {
	return &VerifierEntityAdminClient{mock, invocationCountMatcher, nil}
}

func (mock *MockEntityAdminClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierEntityAdminClient {
	return &VerifierEntityAdminClient{mock, invocationCountMatcher, inOrderContext}
}

type VerifierEntityAdminClient struct {
	mock                   *MockEntityAdminClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierEntityAdminClient) AllianceUpdate(_param0 context.Context, _param1 *proto.AllianceAdminRequest, _param2 ...client.CallOption) *EntityAdminClient_AllianceUpdate_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AllianceUpdate", params)
	return &EntityAdminClient_AllianceUpdate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityAdminClient_AllianceUpdate_OngoingVerification struct {
	mock              *MockEntityAdminClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityAdminClient_AllianceUpdate_OngoingVerification) GetCapturedArguments() (context.Context, *proto.AllianceAdminRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityAdminClient_AllianceUpdate_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.AllianceAdminRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.AllianceAdminRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.AllianceAdminRequest)
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

func (verifier *VerifierEntityAdminClient) CharacterUpdate(_param0 context.Context, _param1 *proto.CharacterAdminRequest, _param2 ...client.CallOption) *EntityAdminClient_CharacterUpdate_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CharacterUpdate", params)
	return &EntityAdminClient_CharacterUpdate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityAdminClient_CharacterUpdate_OngoingVerification struct {
	mock              *MockEntityAdminClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityAdminClient_CharacterUpdate_OngoingVerification) GetCapturedArguments() (context.Context, *proto.CharacterAdminRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityAdminClient_CharacterUpdate_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.CharacterAdminRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.CharacterAdminRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.CharacterAdminRequest)
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

func (verifier *VerifierEntityAdminClient) CorporationUpdate(_param0 context.Context, _param1 *proto.CorporationAdminRequest, _param2 ...client.CallOption) *EntityAdminClient_CorporationUpdate_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CorporationUpdate", params)
	return &EntityAdminClient_CorporationUpdate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityAdminClient_CorporationUpdate_OngoingVerification struct {
	mock              *MockEntityAdminClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityAdminClient_CorporationUpdate_OngoingVerification) GetCapturedArguments() (context.Context, *proto.CorporationAdminRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityAdminClient_CorporationUpdate_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.CorporationAdminRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.CorporationAdminRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.CorporationAdminRequest)
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

func (verifier *VerifierEntityAdminClient) RoleUpdate(_param0 context.Context, _param1 *proto.RoleAdminRequest, _param2 ...client.CallOption) *EntityAdminClient_RoleUpdate_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RoleUpdate", params)
	return &EntityAdminClient_RoleUpdate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type EntityAdminClient_RoleUpdate_OngoingVerification struct {
	mock              *MockEntityAdminClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *EntityAdminClient_RoleUpdate_OngoingVerification) GetCapturedArguments() (context.Context, *proto.RoleAdminRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *EntityAdminClient_RoleUpdate_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.RoleAdminRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.RoleAdminRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.RoleAdminRequest)
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
