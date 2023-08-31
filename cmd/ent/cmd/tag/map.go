//
// Copyright 2023 The Ent Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tag

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/google/ent/proto"
	"github.com/spf13/cobra"
)

var (
	remoteFlag string
)

var TagCmd = &cobra.Command{
	Use: "tag",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	TagCmd.AddCommand(setCmd)
	TagCmd.AddCommand(getCmd)
}

func ValidateEntry(m *pb.Tag, ecpk *ecdsa.PublicKey, signature []byte) error {
	entryBytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	if !ecdsa.VerifyASN1(ecpk, entryBytes, signature) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
