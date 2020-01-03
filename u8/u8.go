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

package u8

import "unsafe"

const (
	Size = unsafe.Sizeof(uint8(0))
	BitSize = Size << 3
	Minimal = uint8(0)
	Maximal = ^Minimal
)

func DivMod(dividend, divisor uint8) (uint8, uint8) {
	return dividend / divisor, dividend % divisor
}

func GCD(value_0, value_1 uint8) uint8 {
	for value_1 != 0 {
		value_0, value_1 = value_1, value_0 % value_1
	}
	return value_0
}

func Is2Power(value uint8) bool {
	return value != 0 && (value & (value - 1) == 0)
}

func IsOdd(value uint8) bool {
	return (value & 1) != 0
}

func LCM(value_0, value_1 uint8) uint8 {
	if value_0 == 0 || value_1 == 0 {
		return 0
	}
	return value_0 / GCD(value_0, value_1) * value_1
}

func Min(value_0, value_1 uint8) uint8 {
	if value_0 < value_1 {
		return value_0
	}
	return value_1
}

func Mins(value uint8, values ...uint8) uint8 {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func MinSlice(values []uint8) uint8 {
	return Mins(values[0], values[1:]...)
}

func MinSliceChecked(values []uint8) (uint8, bool) {
	if len(values) == 0 {
		return 0, true
	}
	return MinSlice(values), false
}

func Max(value_0, value_1 uint8) uint8 {
	if value_0 > value_1 {
		return value_0
	}
	return value_1
}

func Maxs(value uint8, values ...uint8) uint8 {
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

func MaxSlice(values []uint8) uint8 {
	return Maxs(values[0], values[1:]...)
}

func MaxSliceChecked(values []uint8) (uint8, bool) {
	if len(values) == 0 {
		return 0, true
	}
	return MaxSlice(values), false
}

func MinMax(value_0, value_1 uint8) (uint8, uint8) {
	if value_0 < value_1 {
		return value_0, value_1
	}
	return value_1, value_0
}

func MinMaxs(value uint8, values ...uint8) (uint8, uint8) {
	min := value
	max := value
	for _, v := range values {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxSlice(values []uint8) (uint8, uint8) {
	return MinMaxs(values[0], values[1:]...)
}

func MinMaxSliceChecked(values []uint8) (uint8, uint8, bool) {
	if len(values) == 0 {
		return 0, 0, true
	}
	min, max := MinMaxSlice(values)
	return min, max, false
}

func Sign(value uint8) uint8 {
	shift := BitSize - 1
	return value >> shift | uint8((-int8(value &^ (1 << shift)) >> shift) & 1)
}
