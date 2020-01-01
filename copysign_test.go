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

func TestCopysigni(t *testing.T) {
	args := []int{
		0, 0,
		0, 1,
		0, -1,
		1, 0,
		-1, 0,
		42, 24,
		42, -7,
		-42, 3,
		-42, -9,
		MaxInt, 0,
		MaxInt, -1,
		-MaxInt, 34,
		-MaxInt, -1,
		MinInt, 0,
		MinInt, -45,
	}
	expected := []int{
		0,
		0,
		0,
		1,
		1,
		42,
		-42,
		42,
		-42,
		MaxInt,
		-MaxInt,
		MaxInt,
		-MaxInt,
		MinInt,
		MinInt,
	}
	for i, e := range expected {
		if r := Copysigni(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, copysign, i_, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkCopysigni(b *testing.B) {
	for i := int(b.N); i > 0; i-- {
		Copysigni(i, 0)
		Copysigni(i, -1)
	}
}

func TestCopysigni8(t *testing.T) {
	args := []int8{
		0, 0,
		0, 1,
		0, -1,
		1, 0,
		-1, 0,
		42, 24,
		42, -7,
		-42, 3,
		-42, -9,
		MaxInt8, 0,
		MaxInt8, -1,
		-MaxInt8, 34,
		-MaxInt8, -1,
		MinInt8, 0,
		MinInt8, -45,
	}
	expected := []int8{
		0,
		0,
		0,
		1,
		1,
		42,
		-42,
		42,
		-42,
		MaxInt8,
		-MaxInt8,
		MaxInt8,
		-MaxInt8,
		MinInt8,
		MinInt8,
	}
	for i, e := range expected {
		if r := Copysigni8(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, copysign, i8, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkCopysigni8(b *testing.B) {
	for i := int8(b.N); i > 0; i-- {
		Copysigni8(i, 0)
		Copysigni8(i, -1)
	}
}

func TestCopysigni16(t *testing.T) {
	args := []int16{
		0, 0,
		0, 1,
		0, -1,
		1, 0,
		-1, 0,
		42, 24,
		42, -7,
		-42, 3,
		-42, -9,
		MaxInt16, 0,
		MaxInt16, -1,
		-MaxInt16, 34,
		-MaxInt16, -1,
		MinInt16, 0,
		MinInt16, -45,
	}
	expected := []int16{
		0,
		0,
		0,
		1,
		1,
		42,
		-42,
		42,
		-42,
		MaxInt16,
		-MaxInt16,
		MaxInt16,
		-MaxInt16,
		MinInt16,
		MinInt16,
	}
	for i, e := range expected {
		if r := Copysigni16(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, copysign, i16, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkCopysigni16(b *testing.B) {
	for i := int16(b.N); i > 0; i-- {
		Copysigni16(i, 0)
		Copysigni16(i, -1)
	}
}

func TestCopysigni32(t *testing.T) {
	args := []int32{
		0, 0,
		0, 1,
		0, -1,
		1, 0,
		-1, 0,
		42, 24,
		42, -7,
		-42, 3,
		-42, -9,
		MaxInt32, 0,
		MaxInt32, -1,
		-MaxInt32, 34,
		-MaxInt32, -1,
		MinInt32, 0,
		MinInt32, -45,
	}
	expected := []int32{
		0,
		0,
		0,
		1,
		1,
		42,
		-42,
		42,
		-42,
		MaxInt32,
		-MaxInt32,
		MaxInt32,
		-MaxInt32,
		MinInt32,
		MinInt32,
	}
	for i, e := range expected {
		if r := Copysigni32(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, copysign, i32, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkCopysigni32(b *testing.B) {
	for i := int32(b.N); i > 0; i-- {
		Copysigni32(i, 0)
		Copysigni32(i, -1)
	}
}

func TestCopysigni64(t *testing.T) {
	args := []int64{
		0, 0,
		0, 1,
		0, -1,
		1, 0,
		-1, 0,
		42, 24,
		42, -7,
		-42, 3,
		-42, -9,
		MaxInt64, 0,
		MaxInt64, -1,
		-MaxInt64, 34,
		-MaxInt64, -1,
		MinInt64, 0,
		MinInt64, -45,
	}
	expected := []int64{
		0,
		0,
		0,
		1,
		1,
		42,
		-42,
		42,
		-42,
		MaxInt64,
		-MaxInt64,
		MaxInt64,
		-MaxInt64,
		MinInt64,
		MinInt64,
	}
	for i, e := range expected {
		if r := Copysigni64(args[i + i], args[i + i + 1]); r != e {
			t.Errorf(testFailMessage2, copysign, i64, args[i + i], args[i + i + 1], r, e, i)
		}
	}
}

func BenchmarkCopysigni64(b *testing.B) {
	for i := int64(b.N); i > 0; i-- {
		Copysigni64(i, 0)
		Copysigni64(i, -1)
	}
}
