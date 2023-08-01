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
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

// Decrypt decrypts the data provided, returning the secret.
func (e *Encryptor) Decrypt(data map[string]any, _ string) ([]byte, error) {
	if data == nil {
		return nil, errors.New("no data supplied")
	}
	// Marshal the map and unmarshal it back in to a format we can work with.
	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse keystore")
	}
	ks := &unencrypted{}
	err = json.Unmarshal(b, &ks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse keystore")
	}

	if ks.Key == "" {
		return nil, errors.New("key missing")
	}

	key, err := hex.DecodeString(strings.TrimPrefix(ks.Key, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode key")
	}

	return key, nil
}
