/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/api/unversioned"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

func NewCmdDeploymentConfigs(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "deployment-configurations",
		Aliases: []string{"dcs"},
		Short:   "Print  the deployment configurations (flags) used to standup the k8s stack.",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunDeploymentConfigs(f, out)
			cmdutil.CheckErr(err)
		},
	}
	return cmd
}

func RunDeploymentConfigs(f *cmdutil.Factory, w io.Writer) error {

	if len(os.Args) > 0 {
		for i, a := range os.Args {
			fmt.Fprintln(w, fmt.Sprintf("at %d we have argument: %s", i, a))
		}
	}

	var module string
	if len(os.Args) >= 2 {
		module = os.Args[2]
	}

	if module == "" {
		module = "all"
	}

	fmt.Fprintln(w, fmt.Sprintf("module to get configurations from: %s", module))

	client, err := f.Client()
	if err != nil {
		return err
	}

	// would not be surprised if this is loaded once.
	groupList, err := client.Discovery().ServerGroups()
	if err != nil {
		return fmt.Errorf("Couldn't get available deployment-configurations from server: %v\n", err)
	}
	apiVersions := unversioned.ExtractGroupVersions(groupList)
	sort.Strings(apiVersions)
	for _, v := range apiVersions {
		fmt.Fprintln(w, v)
	}
	return nil
}
