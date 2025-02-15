// Copyright 2021The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gridscale

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gridscale/gsclient-go/v3"
)

type IPv4Generator struct {
	GridscaleService
}

func (g IPv4Generator) createResources(ipAddrList []gsclient.IP) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for idx, ipAddr := range ipAddrList {
		if ipAddr.Properties.Family != 4 {
			continue
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			ipAddr.Properties.ObjectUUID,
			fmt.Sprintf("%s-%d", ipAddr.Properties.Name, idx),
			"gridscale_ipv4",
			"gridscale",
			[]string{}))
	}
	return resources
}

func (g *IPv4Generator) InitResources() error {
	client := g.generateClient()
	ipAddrList, err := client.GetIPList(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(ipAddrList)
	return nil
}
