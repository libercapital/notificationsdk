// Code generated by mockery v2.35.3. DO NOT EDIT.

package notificationsdk

import mock "github.com/stretchr/testify/mock"

// SnsClient is an autogenerated mock type for the SnsClient type
type SnsClient struct {
	mock.Mock
}

// NewSnsClient creates a new instance of SnsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSnsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SnsClient {
	mock := &SnsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
