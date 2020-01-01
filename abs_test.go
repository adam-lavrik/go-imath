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

func TestAbsi(t *testing.T) {
	args := []int{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		MaxInt,
		-MaxInt,
		MinInt,
	}
	expected := []int{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		MaxInt,
		MaxInt,
		MinInt,
	}
	for i, e := range expected {
		if r := Absi(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, i_, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsi(b *testing.B) {
	for i := int(b.N); i > 0; i-- {
		Absi(-i)
		Absi(i)
	}
}


func TestAbsi8(t *testing.T) {
	args := []int8{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		MaxInt8,
		-MaxInt8,
		MinInt8,
	}
	expected := []int8{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		MaxInt8,
		MaxInt8,
		MinInt8,
	}
	for i, e := range expected {
		if r := Absi8(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, i8, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsi8(b *testing.B) {
	for i := int8(b.N); i > 0; i-- {
		Absi8(-i)
		Absi8(i)
	}
}

func TestAbsi16(t *testing.T) {
	args := []int16{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		16145,
		-16145,
		32000,
		-32000,
		MaxInt16,
		-MaxInt16,
		MinInt16,
	}
	expected := []int16{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		16145,
		16145,
		32000,
		32000,
		MaxInt16,
		MaxInt16,
		MinInt16,
	}
	for i, e := range expected {
		if r := Absi16(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, i16, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsi16(b *testing.B) {
	for i := int16(b.N); i > 0; i-- {
		Absi16(-i)
		Absi16(i)
	}
}

func TestAbsi32(t *testing.T) {
	args := []int32{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		16145,
		-16145,
		32000,
		-32000,
		6613371,
		-6613371,
		12345678,
		-12345678,
		MaxInt32,
		-MaxInt32,
		MinInt32,
	}
	expected := []int32{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		16145,
		16145,
		32000,
		32000,
		6613371,
		6613371,
		12345678,
		12345678,
		MaxInt32,
		MaxInt32,
		MinInt32,
	}
	for i, e := range expected {
		if r := Absi32(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, i32, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsi32(b *testing.B) {
	for i := int32(b.N); i > 0; i-- {
		Absi32(-i)
		Absi32(i)
	}
}

func TestAbsi64(t *testing.T) {
	args := []int64{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		16145,
		-16145,
		32000,
		-32000,
		6613371,
		-6613371,
		12345678,
		-12345678,
		2203820234,
		-2203820234,
		1173840299,
		-1173840299,
		MaxInt64,
		-MaxInt64,
		MinInt64,
	}
	expected := []int64{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		16145,
		16145,
		32000,
		32000,
		6613371,
		6613371,
		12345678,
		12345678,
		2203820234,
		2203820234,
		1173840299,
		1173840299,
		MaxInt64,
		MaxInt64,
		MinInt64,
	}
	for i, expect := range expected {
		if r := Absi64(args[i]); r != expect {
			t.Errorf(testFailMessage1, abs, i64, args[i], r, expect, i)
		}
	}
}

func BenchmarkAbsi64(b *testing.B) {
	for i := int64(b.N); i > 0; i-- {
		Absi64(-i)
		Absi64(i)
	}
}

func TestAbsu(t *testing.T) {
	args := []int{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		MaxInt,
		-MaxInt,
		MinInt,
	}
	expected := []uint{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		uint(MaxInt),
		uint(MaxInt),
		uint(MaxInt) + 1,
	}
	for i, e := range expected {
		if r := Absu(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, u_, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsu(b *testing.B) {
	for i := int(b.N); i > 0; i-- {
		Absu(i)
	}
}


func TestAbsu8(t *testing.T) {
	args := []int8{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		MaxInt8,
		-MaxInt8,
		MinInt8,
	}
	expected := []uint8{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		uint8(MaxInt8),
		uint8(MaxInt8),
		uint8(MaxInt8) + 1,
	}
	for i, e := range expected {
		if r := Absu8(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, u8, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsu8(b *testing.B) {
	for i := int8(b.N); i > 0; i-- {
		Absu8(i)
	}
}

func TestAbsu16(t *testing.T) {
	args := []int16{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		16145,
		-16145,
		32000,
		-32000,
		MaxInt16,
		-MaxInt16,
		MinInt16,
	}
	expected := []uint16{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		16145,
		16145,
		32000,
		32000,
		uint16(MaxInt16),
		uint16(MaxInt16),
		uint16(MaxInt16) + 1,
	}
	for i, e := range expected {
		if r := Absu16(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, u16, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsu16(b *testing.B) {
	for i := int16(b.N); i > 0; i-- {
		Absu16(i)
	}
}

func TestAbsu32(t *testing.T) {
	args := []int32{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		16145,
		-16145,
		32000,
		-32000,
		6613371,
		-6613371,
		12345678,
		-12345678,
		MaxInt32,
		-MaxInt32,
		MinInt32,
	}
	expected := []uint32{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		16145,
		16145,
		32000,
		32000,
		6613371,
		6613371,
		12345678,
		12345678,
		uint32(MaxInt32),
		uint32(MaxInt32),
		uint32(MaxInt32) + 1,
	}
	for i, e := range expected {
		if r := Absu32(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, u32, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsu32(b *testing.B) {
	for i := int32(b.N); i > 0; i-- {
		Absu32(i)
	}
}

func TestAbsu64(t *testing.T) {
	args := []int64{
		0,
		1,
		-1,
		42,
		-42,
		65,
		-65,
		16145,
		-16145,
		32000,
		-32000,
		6613371,
		-6613371,
		12345678,
		-12345678,
		2203820234,
		-2203820234,
		1173840299,
		-1173840299,
		MaxInt64,
		-MaxInt64,
		MinInt64,
	}
	expected := []uint64{
		0,
		1,
		1,
		42,
		42,
		65,
		65,
		16145,
		16145,
		32000,
		32000,
		6613371,
		6613371,
		12345678,
		12345678,
		2203820234,
		2203820234,
		1173840299,
		1173840299,
		uint64(MaxInt64),
		uint64(MaxInt64),
		uint64(MaxInt64) + 1,
	}
	for i, e := range expected {
		if r := Absu64(args[i]); r != e {
			t.Errorf(testFailMessage1, abs, u64, args[i], r, e, i)
		}
	}
}

func BenchmarkAbsu64(b *testing.B) {
	for i := int64(b.N); i > 0; i-- {
		Absu64(i)
	}
}
