/*
 * Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
 *
 * This program and the accompanying materials are made available under
 * the terms of the under the Apache License, Version 2.0 (the "License”);
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package eureka

import (
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/httpclient"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/serviceutil"
)

func Disable(cliConnection plugin.CliConnection, srInstanceName string, cfAppName string, authenticatedClient httpclient.AuthenticatedClient, instanceIndex *int) (string, error) {
	return OperateOnApplication(cliConnection, srInstanceName, cfAppName, authenticatedClient, instanceIndex, serviceutil.ServiceInstanceURL,
		func(authClient httpclient.AuthenticatedClient, accessToken string, eurekaUrl string, eurekaAppName string, instanceId string) error {
			_, err := authClient.DoAuthenticatedPut(fmt.Sprintf("%seureka/apps/%s/%s/status?value=OUT_OF_SERVICE", eurekaUrl, eurekaAppName, instanceId), accessToken)
			return err
		})
}
