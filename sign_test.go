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

func TestSigni(t *testing.T) {
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
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		-1,
	}
	for i, e := range expected {
		if r := Signi(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, i_, args[i], r, e, i)
		}
	}
}

func BenchmarkSigni(b *testing.B) {
	for i := int(b.N); i > 0; i-- {
		Signi(-i)
		Signi(i)
	}
}


func TestSigni8(t *testing.T) {
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
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		-1,
	}
	for i, e:= range expected {
		if r := Signi8(args[i]); r != e{
			t.Errorf(testFailMessage1, sgn, i8, args[i], r, e, i)
		}
	}
}

func BenchmarkSigni8(b *testing.B) {
	for i := int8(b.N); i > 0; i-- {
		Signi8(-i)
		Signi8(i)
	}
}

func TestSigni16(t *testing.T) {
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
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		-1,
	}
	for i, e := range expected {
		if r := Signi16(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, i16, args[i], r, e, i)
		}
	}
}

func BenchmarkSigni16(b *testing.B) {
	for i := int16(b.N); i > 0; i-- {
		Signi16(-i)
		Signi16(i)
	}
}

func TestSigni32(t *testing.T) {
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
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		-1,
	}
	for i, e := range expected {
		if r := Signi32(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, i32, args[i], r, e, i)
		}
	}
}

func BenchmarkSigni32(b *testing.B) {
	for i := int32(b.N); i > 0; i-- {
		Signi32(-i)
		Signi32(i)
	}
}

func TestSigni64(t *testing.T) {
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
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		1,
		-1,
		-1,
	}
	for i, e := range expected {
		if r := Signi64(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, i64, args[i], r, e, i)
		}
	}
}

func BenchmarkSigni64(b *testing.B) {
	for i := int64(b.N); i > 0; i-- {
		Signi64(-i)
		Signi64(i)
	}
}

func TestSignu(t *testing.T) {
	args := []uint{
		0,
		1,
		42,
		65,
		MaxUint,
	}
	expected := []uint{
		0,
		1,
		1,
		1,
		1,
	}
	for i, e := range expected {
		if r := Signu(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, u_, args[i], r, e, i)
		}
	}
}

func BenchmarkSignu(b *testing.B) {
	for i := uint(b.N); i > 0; i-- {
		Signu(i)
	}
}

func TestSignu8(t *testing.T) {
	args := []uint8{
		0,
		1,
		42,
		65,
		uint8(MaxInt8),
		MaxUint8,
	}
	expected := []uint8{
		0,
		1,
		1,
		1,
		1,
		1,
	}
	for i, e := range expected {
		if r := Signu8(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, u8, args[i], r, e, i)
		}
	}
}

func BenchmarkSignu8(b *testing.B) {
	for i := uint8(b.N); i > 0; i-- {
		Signu8(i)
	}
}

func TestSignu16(t *testing.T) {
	args := []uint16{
		0,
		1,
		42,
		65,
		16145,
		32000,
		uint16(MaxInt16),
		MaxUint16,
	}
	expected := []uint16{
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
	}
	for i, e := range expected {
		if r := Signu16(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, u16, args[i], r, e, i)
		}
	}
}

func BenchmarkSignu16(b *testing.B) {
	for i := uint16(b.N); i > 0; i-- {
		Signu16(i)
	}
}

func TestSignu32(t *testing.T) {
	args := []uint32{
		0,
		1,
		42,
		65,
		16145,
		32000,
		6613371,
		12345678,
		uint32(MaxInt32),
		MaxUint32,
	}
	expected := []uint32{
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
	}
	for i, e := range expected {
		if r := Signu32(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, u32, args[i], r, e, i)
		}
	}
}

func BenchmarkSignu32(b *testing.B) {
	for i := uint32(b.N); i > 0; i-- {
		Signu32(i)
	}
}

func TestSignu64(t *testing.T) {
	args := []uint64{
		0,
		1,
		42,
		65,
		16145,
		32000,
		6613371,
		12345678,
		2203820234,
		1173840299,
		uint64(MaxInt64),
		MaxUint64,
	}
	expected := []uint64{
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
	}
	for i, e := range expected {
		if r := Signu64(args[i]); r != e {
			t.Errorf(testFailMessage1, sgn, u64, args[i], r, e, i)
		}
	}
}

func BenchmarkSignu64(b *testing.B) {
	for i := uint64(b.N); i > 0; i-- {
		Signu64(i)
	}
}
