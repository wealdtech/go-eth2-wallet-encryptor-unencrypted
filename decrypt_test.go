// Copyright © 2020 Weald Technology Trading
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

package unencrypted_test

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	unencrypted "github.com/wealdtech/go-eth2-wallet-encryptor-unencrypted"
)

func TestDecrypt(t *testing.T) {
	circular := make(map[string]interface{})
	circular["here"] = &circular

	tests := []struct {
		name   string
		input  map[string]interface{}
		output []byte
		err    string
	}{
		{
			name: "Nil",
			err:  `no data supplied`,
		},
		{
			name:  "Circular",
			input: circular,
			err:   `failed to parse keystore: json: unsupported value: encountered a cycle via *map[string]interface {}`,
		},
		{
			name:  "KeyMissing",
			input: map[string]interface{}{},
			err:   `key missing`,
		},
		{
			name: "KeyWrongType",
			input: map[string]interface{}{
				"key": true,
			},
			err: `failed to parse keystore: json: cannot unmarshal bool into Go struct field unencrypted.key of type string`,
		},
		{
			name: "KeyInvalid",
			input: map[string]interface{}{
				"key": "invalid",
			},
			err: `failed to decode key: encoding/hex: invalid byte: U+0069 'i'`,
		},
		{
			name: "Good",
			input: map[string]interface{}{
				"key": "0x25295f0d1d592a90b333e26e85149708208e9f8e8bc18f6c77bd62f8ad7a6866",
			},
			output: []byte{
				0x25, 0x29, 0x5f, 0x0d, 0x1d, 0x59, 0x2a, 0x90, 0xb3, 0x33, 0xe2, 0x6e, 0x85, 0x14, 0x97, 0x08,
				0x20, 0x8e, 0x9f, 0x8e, 0x8b, 0xc1, 0x8f, 0x6c, 0x77, 0xbd, 0x62, 0xf8, 0xad, 0x7a, 0x68, 0x66,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encryptor := unencrypted.New()
			output, err := encryptor.Decrypt(test.input, "")
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, test.output, output)
			}
		})
	}
}
