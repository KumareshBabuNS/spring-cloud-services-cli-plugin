package instance_test

import (
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/httpclient/httpclientfakes"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/instance"
)

var _ = DescribeCommandTest("Stop", "stop",
	func(fakeAuthClient *httpclientfakes.FakeAuthenticatedClient, serviceInstanceAdminURL string,
		accessToken string) (string, error) {

		return instance.Stop(fakeAuthClient, serviceInstanceAdminURL, accessToken)
	})
