// Copyright © 2020, 2023 Weald Technology Trading
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

// Encryptor is an encryptor that stores data unencrypted.
type Encryptor struct{}

type unencrypted struct {
	Key string `json:"key"`
}

// New creates a new null encryptor.
func New() *Encryptor {
	return &Encryptor{}
}

// String returns the string for this encryptor.
func (e *Encryptor) String() string {
	return "unencryptedv1"
}

// Name returns the name of this encryptor.
func (e *Encryptor) Name() string {
	return "unencrypted"
}

// Version returns the version of this encryptor.
func (e *Encryptor) Version() uint {
	return 1
}
