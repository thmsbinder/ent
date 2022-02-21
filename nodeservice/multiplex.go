//
// Copyright 2021 The Ent Authors.
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

package nodeservice

import (
	"context"
	"fmt"
	"log"

	"github.com/google/ent/utils"
)

type Multiplex struct {
	Inner []ObjectGetter
}

func (s Multiplex) Get(ctx context.Context, h utils.Hash) ([]byte, error) {
	for i, ss := range s.Inner {
		b, err := ss.Get(ctx, h)
		if err != nil {
			log.Printf("error fetching from remote %d: %v", i, err)
			continue
		}
		return b, nil
	}
	return nil, fmt.Errorf("not found")
}

func (s Multiplex) Has(ctx context.Context, h utils.Hash) (bool, error) {
	for _, i := range s.Inner {
		b, err := i.Has(ctx, h)
		if err != nil {
			continue
		}
		return b, nil
	}
	return false, nil
}

func (s Multiplex) Put(ctx context.Context, b []byte) (utils.Hash, error) {
	// return s.Inner[0].Put(ctx, b)
	return "", fmt.Errorf("not implemented")
}
