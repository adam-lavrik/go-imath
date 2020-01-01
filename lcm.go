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

// Least common multiply of two integers

func LCMi(x, y int) uint {
	if x == 0 || y == 0 {
		return 0
	}
	return Absu(x) / GCDi(x, y) * Absu(y)
}

func LCMi8(x, y int8) uint8 {
	if x == 0 || y == 0 {
		return 0
	}
	return Absu8(x) / GCDi8(x, y) * Absu8(y)
}

func LCMi16(x, y int16) uint16 {
	if x == 0 || y == 0 {
		return 0
	}
	return Absu16(x) / GCDi16(x, y) * Absu16(y)
}

func LCMi32(x, y int32) uint32 {
	if x == 0 || y == 0 {
		return 0
	}
	return Absu32(x) / GCDi32(x, y) * Absu32(y)
}

func LCMi64(x, y int64) uint64 {
	if x == 0 || y == 0 {
		return 0
	}
	return Absu64(x) / GCDi64(x, y) * Absu64(y)
}

func LCMu(x, y uint) uint {
	if x == 0 || y == 0 {
		return 0
	}
	return x / GCDu(x, y) * y
}

func LCMu8(x, y uint8) uint8 {
	if x == 0 || y == 0 {
		return 0
	}
	return x / GCDu8(x, y) * y
}

func LCMu16(x, y uint16) uint16 {
	if x == 0 || y == 0 {
		return 0
	}
	return x / GCDu16(x, y) * y
}

func LCMu32(x, y uint32) uint32 {
	if x == 0 || y == 0 {
		return 0
	}
	return x / GCDu32(x, y) * y
}

func LCMu64(x, y uint64) uint64 {
	if x == 0 || y == 0 {
		return 0
	}
	return x / GCDu64(x, y) * y
}
