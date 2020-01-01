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

// Check whether the `number` is power of 2 (true) or not (false).
// Negative values or zeros always produce false.
package imath

func Is2Poweri(number int) bool {
	return number > 0 && (number & (number - 1) == 0)
}

func Is2Poweri8(number int8) bool {
	return number > 0 && (number & (number - 1) == 0)
}

func Is2Poweri16(number int16) bool {
	return number > 0 && (number & (number - 1) == 0)
}

func Is2Poweri32(number int32) bool {
	return number > 0 && (number & (number - 1) == 0)
}

func Is2Poweri64(number int64) bool {
	return number > 0 && (number & (number - 1) == 0)
}

func Is2Poweru(number uint) bool {
	return number != 0 && (number & (number - 1) == 0)
}

func Is2Poweru8(number uint8) bool {
	return number != 0 && (number & (number - 1) == 0)
}

func Is2Poweru16(number uint16) bool {
	return number != 0 && (number & (number - 1) == 0)
}

func Is2Poweru32(number uint32) bool {
	return number != 0 && (number & (number - 1) == 0)
}

func Is2Poweru64(number uint64) bool {
	return number != 0 && (number & (number - 1) == 0)
}
