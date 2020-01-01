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

// Functions returning signed (of the same type as `number`):

func Absi(number int) int {
	mask := number >> (IntBitSize - 1)
	return (number ^ mask) + (mask & 1)
}

func Absi8(number int8) int8 {
	mask := number >> (Int8BitSize - 1)
	return (number ^ mask) + (mask & 1)
}

func Absi16(number int16) int16 {
	mask := number >> (Int16BitSize - 1)
	return (number ^ mask) + (mask & 1)
}

func Absi32(number int32) int32 {
	mask := number >> (Int32BitSize - 1)
	return (number ^ mask) + (mask & 1)
}

func Absi64(number int64) int64 {
	mask := number >> (Int64BitSize - 1)
	return (number ^ mask) + (mask & 1)
}

// Functions returning unsigned (complementary to signed type of `number`):

func Absu(number int) uint {
	mask := number >> (UintBitSize - 1)
	return uint(number ^ mask) + uint(mask & 1)
}

func Absu8(number int8) uint8 {
	mask := number >> (Uint8BitSize - 1)
	return uint8(number ^ mask) + uint8(mask & 1)
}

func Absu16(number int16) uint16 {
	mask := number >> (Uint16BitSize - 1)
	return uint16(number ^ mask) + uint16(mask & 1)
}

func Absu32(number int32) uint32 {
	mask := number >> (Uint32BitSize - 1)
	return uint32(number ^ mask) + uint32(mask & 1)
}

func Absu64(number int64) uint64 {
	mask := number >> (Uint64BitSize - 1)
	return uint64(number ^ mask) + uint64(mask & 1)
}
