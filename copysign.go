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

// If `target` is not 0 and signs of `target` and `source` are opposite, return -`target`;
// otherwise return `target` as is.
// If `source` is 0, its sign is considered positive (0).
// Implemented only for signed integer types ().

func Copysigni(target, source int) int {
	if target == 0 {
		return 0
	}
	source = (target ^ source) >> (IntBitSize - 1)
	return (target ^ source) + (source & 1)
}

func Copysigni8(target, source int8) int8 {
	if target == 0 {
		return 0
	}
	source = (target ^ source) >> (Int8BitSize - 1)
	return (target ^ source) + (source & 1)
}

func Copysigni16(target, source int16) int16 {
	if target == 0 {
		return 0
	}
	source = (target ^ source) >> (Int16BitSize - 1)
	return (target ^ source) + (source & 1)
}

func Copysigni32(target, source int32) int32 {
	if target == 0 {
		return 0
	}
	source = (target ^ source) >> (Int32BitSize - 1)
	return (target ^ source) + (source & 1)
}

func Copysigni64(target, source int64) int64 {
	if target == 0 {
		return 0
	}
	source = (target ^ source) >> (Int64BitSize - 1)
	return (target ^ source) + (source & 1)
}
