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
package i64

// Pow raises `base` to `exponent` power.
// Fast binary algorithm is used.
// Pow(0, 0) == 1
func Pow(base int64, exponent uint) int64 {
	power := int64(1)
	for exponent > 0 {
		if (exponent & 1) != 0 { // If exponent is odd
			power *= base
		}
		base *= base
		exponent >>= 1 // `exponent` fast division by 2
	}
	return power
}
