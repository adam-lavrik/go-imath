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

func TestPowi(t *testing.T) {
	args0 := []int{
		0,
		0,
		0,
		1,
		-1,
		2,
		-2,
		2,
		-2,
		2,
		-2,
	}
	args1 := []uint{
		0,
		3,
		4,
		0,
		0,
		1,
		1,
		2,
		2,
		5,
		5,
	}
	expected := []int{
		1,
		0,
		0,
		1,
		1,
		2,
		-2,
		4,
		4,
		32,
		-32,
	}
	for i, e := range expected {
		if r := Powi(args0[i], args1[i]); r != e {
			t.Errorf(testFailMessage2, pow, i_, args0[i], args1[i], r, e, i)
		}
	}
}

func BenchmarkPowi(b *testing.B) {
	for i := uint(b.N); i > 0; i-- {
		Powi(5, i)
	}
}


func TestPowi64(t *testing.T) {
	args0 := []int64{
		0,
		0,
		0,
		1,
		-1,
		2,
		-2,
		2,
		-2,
		2,
		-2,
	}
	args1 := []uint{
		0,
		3,
		4,
		0,
		0,
		1,
		1,
		2,
		2,
		5,
		5,
	}
	expected := []int64{
		1,
		0,
		0,
		1,
		1,
		2,
		-2,
		4,
		4,
		32,
		-32,
	}
	for i, e := range expected {
		if r := Powi64(args0[i], args1[i]); r != e {
			t.Errorf(testFailMessage2, pow, i64, args0[i], args1[i], r, e, i)
		}
	}
}

func BenchmarkPowi64(b *testing.B) {
	for i := uint(b.N); i > 0; i-- {
		Powi64(5, i)
	}
}


func TestPowu(t *testing.T) {
	args0 := []uint{
		0,
		0,
		0,
		1,
		2,
		2,
		2,
	}
	args1 := []uint{
		0,
		3,
		4,
		0,
		1,
		2,
		5,
	}
	expected := []uint{
		1,
		0,
		0,
		1,
		2,
		4,
		32,
	}
	for i, e := range expected {
		if r := Powu(args0[i], args1[i]); r != e {
			t.Errorf(testFailMessage2, pow, u_, args0[i], args1[i], r, e, i)
		}
	}
}

func BenchmarkPowu(b *testing.B) {
	for i := uint(b.N); i > 0; i-- {
		Powu(5, i)
	}
}

func TestPowu64(t *testing.T) {
	args0 := []uint64{
		0,
		0,
		0,
		1,
		2,
		2,
		2,
	}
	args1 := []uint{
		0,
		3,
		4,
		0,
		1,
		2,
		5,
	}
	expected := []uint64{
		1,
		0,
		0,
		1,
		2,
		4,
		32,
	}
	for i, e := range expected {
		if r := Powu64(args0[i], args1[i]); r != e {
			t.Errorf(testFailMessage2, pow, u64, args0[i], args1[i], r, e, i)
		}
	}
}

func BenchmarkPowu64(b *testing.B) {
	for i := uint(b.N); i > 0; i-- {
		Powu64(5, i)
	}
}
