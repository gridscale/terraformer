// Copyright 2021 The Terraformer Authors.
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
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type GridscaleProvider struct { //nolint
	terraformutils.Provider
	uuid   string
	token  string
	apiURL string
}

func (p *GridscaleProvider) Init(args []string) error {
	if os.Getenv("GRIDSCALE_UUID") == "" {
		return errors.New("set GRIDSCALE_UUID env var")
	}
	if os.Getenv("GRIDSCALE_TOKEN") == "" {
		return errors.New("set GRIDSCALE_TOKEN env var")
	}
	p.uuid = os.Getenv("GRIDSCALE_UUID")
	p.token = os.Getenv("GRIDSCALE_TOKEN")
	p.apiURL = os.Getenv("GRIDSCALE_URL")
	return nil
}

func (p *GridscaleProvider) GetName() string {
	return "gridscale"
}

func (p *GridscaleProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (GridscaleProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *GridscaleProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"server":  &ServerGenerator{},
		"ipv4":    &IPv4Generator{},
		"ipv6":    &IPv6Generator{},
		"storage": &StorageGenerator{},
		"network": &NetworkGenerator{},
	}
}

func (p *GridscaleProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("gridscale: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"uuid":   p.uuid,
		"token":  p.token,
		"apiURL": p.apiURL,
	})
	return nil
}
