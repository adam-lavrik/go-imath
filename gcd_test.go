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

func TestGCDi(t *testing.T) {
	args := []int{
		0,		0,
		10,		0,
		-10,	0,
		0,		10,
		0,		-10,
		42,		21,
		42,		-21,
		-42,	21,
		-42,	-21,
		65,		30,
		65,		-30,
		-65,    30,
		-65,    -30,
		MaxInt,	MaxInt,
		MaxInt,	-MaxInt,
		MaxInt,	MinInt,
		MinInt,	MaxInt,
		MinInt,	-MaxInt,
		MinInt,	MinInt,
	}
	expected := []uint{
		0,
		10,
		10,
		10,
		10,
		21,
		21,
		21,
		21,
		5,
		5,
		5,
		5,
		uint(MaxInt),
		uint(MaxInt),
		1,
		1,
		1,
		Absu(MinInt),
	}
	for i, e := range expected {
		if r := GCDi(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, gcd, i_, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkTestGCDi(b *testing.B) {
	for i, j := int(b.N), 0; i > 0; i, j = i - 1, j + 1 {
		GCDi(i, j)
	}
}

func TestGCDi8(t *testing.T) {
	args := []int8{
		0,		0,
		10,		0,
		-10,	0,
		0,		10,
		0,		-10,
		42,		21,
		42,		-21,
		-42,	21,
		-42,	-21,
		65,		30,
		65,		-30,
		-65,    30,
		-65,    -30,
		MaxInt8,	MaxInt8,
		MaxInt8,	-MaxInt8,
		MaxInt8,	MinInt8,
		MinInt8,	MaxInt8,
		MinInt8,	-MaxInt8,
		MinInt8,	MinInt8,
	}
	expected := []uint8{
		0,
		10,
		10,
		10,
		10,
		21,
		21,
		21,
		21,
		5,
		5,
		5,
		5,
		uint8(MaxInt8),
		uint8(MaxInt8),
		1,
		1,
		1,
		Absu8(MinInt8),
	}
	for i, e := range expected {
		if r := GCDi8(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, gcd, i8, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkTestGCDi8(b *testing.B) {
	for i, j := int8(b.N), int8(0); i > 0; i, j = i - 1, j + 1 {
		GCDi8(i, j)
	}
}

func TestGCDi16(t *testing.T) {
	args := []int16{
		0,		0,
		10,		0,
		-10,	0,
		0,		10,
		0,		-10,
		42,		21,
		42,		-21,
		-42,	21,
		-42,	-21,
		65,		30,
		65,		-30,
		-65,    30,
		-65,    -30,
		MaxInt16,	MaxInt16,
		MaxInt16,	-MaxInt16,
		MaxInt16,	MinInt16,
		MinInt16,	MaxInt16,
		MinInt16,	-MaxInt16,
		MinInt16,	MinInt16,
	}
	expected := []uint16{
		0,
		10,
		10,
		10,
		10,
		21,
		21,
		21,
		21,
		5,
		5,
		5,
		5,
		uint16(MaxInt16),
		uint16(MaxInt16),
		1,
		1,
		1,
		Absu16(MinInt16),
	}
	for i, e := range expected {
		if r := GCDi16(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, gcd, i16, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkTestGCDi16(b *testing.B) {
	for i, j := int16(b.N), int16(0); i > 0; i, j = i - 1, j + 1 {
		GCDi16(i, j)
	}
}
