/*
Copyright 2020...2021 Adam Lavrik <lavrik.adam@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific
language governing permissions and limitations under the License.
*/
package ix

import "unsafe"

type T = int
type UT = uint

const (
	Size = unsafe.Sizeof(T(0))
	BitSize = Size << 3
	Maximal = T(^UT(0) >> 1)
	Minimal = ^Maximal
)

func Abs(value T) T {
	signBit := SignBit(value)
	return (value ^ signBit) + (signBit & 1)
}

func Absu(value T) UT {
	signBit := SignBit(value)
	return UT(value ^ signBit) + UT(signBit & 1)
}

func Copysign(target, source T) T {
	if target == 0 {
		return 0
	}
	source = SignBit(target ^ source)
	return (target ^ source) + (source & 1)
}

func DivMod(dividend, divisor T) (T, T) {
	return dividend / divisor, dividend % divisor
}

func GCD(value_0, value_1 T) UT {
	for value_1 != 0 {
		value_0, value_1 = value_1, value_0 % value_1
	}
	return Absu(value_0)
}

func Is2Power(value T) bool {
	return value > 0 && (value & (value - 1) == 0)
}

func IsOdd(value T) bool {
	return (value & 1) != 0
}

func LCM(value_0, value_1 T) UT {
	if value_0 == 0 || value_1 == 0 {
		return 0
	}
	return Absu(value_0) / GCD(value_0, value_1) * Absu(value_1)
}

func Min(value_0, value_1 T) T {
	if value_0 < value_1 {
		return value_0
	}
	return value_1
}

func Mins(value T, values ...T) T {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func MinSlice(values []T) T {
	return Mins(values[0], values[1:]...)
}

func MinSliceChecked(values []T) (T, bool) {
	if len(values) == 0 {
		return 0, true
	}
	return MinSlice(values), false
}

func Max(value_0, value_1 T) T {
	if value_0 > value_1 {
		return value_0
	}
	return value_1
}

func Maxs(value T, values ...T) T {
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

func MaxSlice(values []T) T {
	return Maxs(values[0], values[1:]...)
}

func MaxSliceChecked(values []T) (T, bool) {
	if len(values) == 0 {
		return 0, true
	}
	return MaxSlice(values), false
}

func MinMax(value_0, value_1 T) (T, T) {
	if value_0 < value_1 {
		return value_0, value_1
	}
	return value_1, value_0
}

func MinMaxs(value T, values ...T) (T, T) {
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

func MinMaxSlice(values []T) (T, T) {
	return MinMaxs(values[0], values[1:]...)
}

func MinMaxSliceChecked(values []T) (T, T, bool) {
	if len(values) == 0 {
		return 0, 0, true
	}
	min, max := MinMaxSlice(values)
	return min, max, false
}

func Sign(value T) T {
	return SignBit(value) | SignBit(-value) & 1
}

func SignBit(value T) T {
	return value >> (BitSize - 1)
}

