// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveChar(t *testing.T) {
	tests := []struct {
		title     string
		candidate string
		chars     string
		expected  string
	}{
		{
			title:     "when we have one char to replace",
			candidate: "hello foobar",
			chars:     "a",
			expected:  "hello foobr",
		},
		{
			title:     "when we have multiple chars to replace",
			candidate: "hello foobar",
			chars:     "all",
			expected:  "heo foobr",
		},
		{
			title:     "when we have no chars to replace",
			candidate: "hello foobar",
			chars:     "x",
			expected:  "hello foobar",
		},
		{
			title:     "when we have an empty string",
			candidate: "",
			chars:     "x",
			expected:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.expected, RemoveChars(test.candidate, test.chars))
		})
	}
}
