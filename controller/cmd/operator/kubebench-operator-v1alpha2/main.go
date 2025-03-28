// Copyright 2019 The Kubeflow Authors
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

package main

import (
	"flag"
	"os"

	"github.com/liayan/kubebench/controller/cmd/operator/kubebench-operator-v1alpha2/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	opt := app.NewServerOption()
	opt.AddFlags(flag.CommandLine)

	flag.Parse()

	if err := app.Run(opt); err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}
}
