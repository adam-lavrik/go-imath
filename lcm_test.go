// Copyright 2020 Adam Lavrik <lavrik.adam@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//  http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.

package imath

import "testing"

func TestLCMi(t *testing.T) {
	args := []int{
		0,		0,
		20,		8,
		20,		7,
		20,		10,
	}
	expected := []uint{
		0,
		40,
		140,
		20,
	}
	for i, e := range expected {
		if r := LCMi(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, lcm, i_, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkTestLCMi(b *testing.B) {
	for i, j := int(b.N), 0; i > 0; i, j = i - 1, j + 1 {
		LCMi(i, j)
	}
}
