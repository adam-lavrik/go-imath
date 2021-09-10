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
package ux

// Fibonacci(index) returns Fibonacci sequence member with corresponding index:
// - Fibonacci(0) == 0
// - Fibonacci(1) == Fibonacci(2) == 1
// - Fibonacci(3) == 2
// - Fibonacci(4) == 3
// ...
// Result is calculated via matrix ((1, 1), (1, 0)) fast raising to `index` power.
func Fibonacci(index uint) uint {
	v0, v1 := uint(0), uint(1) // Result vector

	for m00, m01, m10, m11 := v1, v1, v1, v0; index > 0; index >>= 1 { // `index` fast division by 2
		if IsOdd(index) {
			v0, v1 = v0 * m00 + v1 * m10, v0 * m01 + v1 * m11 // If power is odd then multiply result vector by matrix
		}
		m00, m01, m10, m11 = m00 * m00 + m01 * m10, m00 * m01 + m01 * m11, m10 * m00 + m11 * m10, m10 * m01 + m11 * m11 // Square the matrix
	}
	return v0
}
