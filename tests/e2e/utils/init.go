/*
Copyright (c) Advanced Micro Devices, Inc. All rights reserved.
Licensed under the Apache License, Version 2.0 (the \"License\");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an \"AS IS\" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"log"
	"os"
)

var (
	nodeAppImage string
)

func init() {
	var ok bool
	// read e2e related env variables
	nodeAppImage, ok = os.LookupEnv("E2E_NODEAPP_IMG")
	if !ok {
		log.Fatalf("E2E_NODEAPP_IMG is not defined. Please determine an image repo to temporarily save node app image, which will be used for running e2e test.")
	}
}
