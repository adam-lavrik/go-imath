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

// Greatest common divisor of two integers

func GCDi(x, y int) uint {
	for y != 0 {
		x, y = y, x % y
	}
	return Absu(x)
}

func GCDi8(x, y int8) uint8 {
	for y != 0 {
		x, y = y, x % y
	}
	return Absu8(x)
}

func GCDi16(x, y int16) uint16 {
	for y != 0 {
		x, y = y, x % y
	}
	return Absu16(x)
}

func GCDi32(x, y int32) uint32 {
	for y != 0 {
		x, y = y, x % y
	}
	return Absu32(x)
}

func GCDi64(x, y int64) uint64 {
	for y != 0 {
		x, y = y, x % y
	}
	return Absu64(x)
}

func GCDu(x, y uint) uint {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func GCDu8(x, y uint8) uint8 {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func GCDu16(x, y uint16) uint16 {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func GCDu32(x, y uint32) uint32 {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func GCDu64(x, y uint64) uint64 {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}
