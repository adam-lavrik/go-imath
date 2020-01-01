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

// Exponentiation of integer `base` and non-negative `exponent`.
// Warning: if `exponent` is 0, result is always 1, even if `base` is 0 too.
package imath

func Powi(base int, exponent uint) int {
	result := int(1)
	for exponent > 0 {
		if IsOddu(exponent) {
			result *= base
		}
		base *= base
		exponent >>= 1
	}
	return result
}

func Powi64(base int64, exponent uint) int64 {
	result := int64(1)
	for exponent > 0 {
		if IsOddu(exponent) {
			result *= base
		}
		base *= base
		exponent >>= 1
	}
	return result
}

func Powu(base uint, exponent uint) uint {
	result := uint(1)
	for exponent > 0 {
		if IsOddu(exponent) {
			result *= base
		}
		base *= base
		exponent >>= 1
	}
	return result
}

func Powu64(base uint64, exponent uint) uint64 {
	result := uint64(1)
	for exponent > 0 {
		if IsOddu(exponent) {
			result *= base
		}
		base *= base
		exponent >>= 1
	}
	return result
}
