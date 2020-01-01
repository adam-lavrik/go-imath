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

// Functions of signed (returning -1 for `number` < 0, 0 for `number` == 0, 1 for `number` > 0):

func Signi(number int) int {
	shift := IntBitSize - 1
	return number >> shift | (-number >> shift) & 1
}

func Signi8(number int8) int8 {
	shift := Int8BitSize - 1
	return number >> shift | (-number >> shift) & 1
}

func Signi16(number int16) int16 {
	shift := Int16BitSize - 1
	return number >> shift | (-number >> shift) & 1
}

func Signi32(number int32) int32 {
	shift := Int32BitSize - 1
	return number >> shift | (-number >> shift) & 1
}

func Signi64(number int64) int64 {
	shift := Int64BitSize - 1
	return number >> shift | (-number >> shift) & 1
}

// Functions of unsigned (returning 0 for `number` == 0, 1 for `number` > 0):

func Signu(number uint) uint {
	shift := UintBitSize - 1
	return number >> shift | uint((-int(number &^ (1 << shift)) >> shift) & 1)
}

func Signu8(number uint8) uint8 {
	shift := Uint8BitSize - 1
	return number >> shift | uint8((-int8(number &^ (1 << shift)) >> shift) & 1)
}

func Signu16(number uint16) uint16 {
	shift := Uint16BitSize - 1
	return number >> shift | uint16((-int16(number &^ (1 << shift)) >> shift) & 1)
}

func Signu32(number uint32) uint32 {
	shift := Uint32BitSize - 1
	return number >> shift | uint32((-int32(number &^ (1 << shift)) >> shift) & 1)
}

func Signu64(number uint64) uint64 {
	shift := Uint64BitSize - 1
	return number >> shift | uint64((-int64(number &^ (1 << shift)) >> shift) & 1)
}
