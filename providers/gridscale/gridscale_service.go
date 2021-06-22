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
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gridscale/gsclient-go/v3"
)

const (
	defaultAPIURL                    = "https://api.gridscale.io"
	defaultGSCDelayIntervalMilliSecs = 1000
	defaultGSCMaxNumberOfRetries     = 1
)

type GridscaleService struct { //nolint
	terraformutils.Service
}

func (s *GridscaleService) generateClient() *gsclient.Client {
	apiURL := defaultAPIURL
	if s.Args["apiURL"] != "" {
		apiURL = s.Args["apiURL"].(string)
	}
	cfg := gsclient.NewConfiguration(
		apiURL,
		s.Args["uuid"].(string),
		s.Args["token"].(string),
		os.Getenv("GSC_LOG") != "",
		true,
		defaultGSCDelayIntervalMilliSecs,
		defaultGSCMaxNumberOfRetries,
	)
	client := gsclient.NewClient(cfg)
	return client
}
