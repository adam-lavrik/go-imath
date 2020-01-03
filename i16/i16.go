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

package i16

import "unsafe"

const (
	Size = unsafe.Sizeof(int16(0))
	BitSize = Size << 3
	Minimal = int16(-1) << (BitSize - 1)
	Maximal = ^Minimal
)

func Abs(value int16) int16 {
	mask := value >> (BitSize - 1)
	return (value ^ mask) + (mask & 1)
}

func Absu(value int16) uint16 {
	mask := value >> (BitSize - 1)
	return uint16(value ^ mask) + uint16(mask & 1)
}

func Copysign(target, source int16) int16 {
	if target == 0 {
		return 0
	}
	source = (target ^ source) >> (BitSize - 1)
	return (target ^ source) + (source & 1)
}

func DivMod(dividend, divisor int16) (int16, int16) {
	return dividend / divisor, dividend % divisor
}

func GCD(value_0, value_1 int16) uint16 {
	for value_1 != 0 {
		value_0, value_1 = value_1, value_0 % value_1
	}
	return Absu(value_1)
}

func Is2Power(value int16) bool {
	return value > 0 && (value & (value - 1) == 0)
}

func IsOdd(value int16) bool {
	return (value & 1) != 0
}

func LCM(value_0, value_1 int16) uint16 {
	if value_0 == 0 || value_1 == 0 {
		return 0
	}
	return Absu(value_0) / GCD(value_0, value_1) * Absu(value_1)
}

func Min(value_0, value_1 int16) int16 {
	if value_0 < value_1 {
		return value_0
	}
	return value_1
}

func Mins(value int16, values ...int16) int16 {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func MinSlice(values []int16) int16 {
	return Mins(values[0], values[1:]...)
}

func MinSliceChecked(values []int16) (int16, bool) {
	if len(values) == 0 {
		return 0, true
	}
	return MinSlice(values), false
}

func Max(value_0, value_1 int16) int16 {
	if value_0 > value_1 {
		return value_0
	}
	return value_1
}

func Maxs(value int16, values ...int16) int16 {
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

func MaxSlice(values []int16) int16 {
	return Maxs(values[0], values[1:]...)
}

func MaxSliceChecked(values []int16) (int16, bool) {
	if len(values) == 0 {
		return 0, true
	}
	return MaxSlice(values), false
}

func MinMax(value_0, value_1 int16) (int16, int16) {
	if value_0 < value_1 {
		return value_0, value_1
	}
	return value_1, value_0
}

func MinMaxs(value int16, values ...int16) (int16, int16) {
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

func MinMaxSlice(values []int16) (int16, int16) {
	return MinMaxs(values[0], values[1:]...)
}

func MinMaxSliceChecked(values []int16) (int16, int16, bool) {
	if len(values) == 0 {
		return 0, 0, true
	}
	min, max := MinMaxSlice(values)
	return min, max, false
}

func Sign(value int16) int16 {
	shift := BitSize - 1
	return value >> shift | (-value >> shift) & 1
}
