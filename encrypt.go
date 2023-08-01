// Copyright © 2020, 2023 Weald Technology Trading.
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

package unencrypted

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Encrypt encrypts data.
func (e *Encryptor) Encrypt(secret []byte, _ string) (map[string]any, error) {
	if secret == nil {
		return nil, errors.New("no secret")
	}

	// Build the output.
	output := &unencrypted{
		Key: fmt.Sprintf("%#x", secret),
	}

	// We need to return a generic map; go to JSON and back to obtain it.
	bytes, err := json.Marshal(output)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal output")
	}
	res := make(map[string]any)
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal output")
	}

	return res, nil
}
