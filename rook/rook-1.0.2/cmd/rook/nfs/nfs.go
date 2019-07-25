/*
Copyright 2018 The Rook Authors. All rights reserved.

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

package nfs

import (
	"github.com/coreos/pkg/capnslog"
	"github.com/spf13/cobra"

	"github.com/rook/rook/cmd/rook/rook"
	"github.com/rook/rook/pkg/clusterd"
	"github.com/rook/rook/pkg/util/exec"
)

var Cmd = &cobra.Command{
	Use:   "nfs",
	Short: "Main command for NFS operator and daemons.",
}

var (
	logger = capnslog.NewPackageLogger("github.com/rook/rook", "nfscmd")
)

func init() {
	Cmd.AddCommand(operatorCmd)
	Cmd.AddCommand(provisonerCmd)
	Cmd.AddCommand(serverCmd)
}

func createContext() *clusterd.Context {
	executor := &exec.CommandExecutor{}
	return &clusterd.Context{
		Executor: executor,
		LogLevel: rook.Cfg.LogLevel,
	}
}
