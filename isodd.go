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

func IsOddi(number int) bool {
	return (number & 1) != 0
}

func IsOddi8(number int8) bool {
	return (number & 1) != 0
}

func IsOddi16(number int16) bool {
	return (number & 1) != 0
}

func IsOddi32(number int32) bool {
	return (number & 1) != 0
}

func IsOddi64(number int64) bool {
	return (number & 1) != 0
}

func IsOddu(number uint) bool {
	return (number & 1) != 0
}

func IsOddu8(number uint8) bool {
	return (number & 1) != 0
}

func IsOddu16(number uint16) bool {
	return (number & 1) != 0
}

func IsOddu32(number uint32) bool {
	return (number & 1) != 0
}

func IsOddu64(number uint64) bool {
	return (number & 1) != 0
}
