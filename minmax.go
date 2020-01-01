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

// The lesser of two values

func Mini(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Mini8(x, y int8) int8 {
    if x < y {
        return x
    }
    return y
}

func Mini16(x, y int16) int16 {
    if x < y {
        return x
    }
    return y
}

func Mini32(x, y int32) int32 {
    if x < y {
        return x
    }
    return y
}

func Mini64(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}

func Minu(x, y uint) uint {
    if x < y {
        return x
    }
    return y
}

func Minu8(x, y uint8) uint8 {
    if x < y {
        return x
    }
    return y
}

func Minu16(x, y uint16) uint16 {
    if x < y {
        return x
    }
    return y
}

func Minu32(x, y uint32) uint32 {
    if x < y {
        return x
    }
    return y
}

func Minu64(x, y uint64) uint64 {
    if x < y {
        return x
    }
    return y
}

// The greater of two values

func Maxi(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Maxi8(x, y int8) int8 {
    if x > y {
        return x
    }
    return y
}

func Maxi16(x, y int16) int16 {
    if x > y {
        return x
    }
    return y
}

func Maxi32(x, y int32) int32 {
    if x > y {
        return x
    }
    return y
}

func Maxi64(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}

func Maxu(x, y uint) uint {
    if x > y {
        return x
    }
    return y
}

func Maxu8(x, y uint8) uint8 {
    if x > y {
        return x
    }
    return y
}

func Maxu16(x, y uint16) uint16 {
    if x > y {
        return x
    }
    return y
}

func Maxu32(x, y uint32) uint32 {
    if x > y {
        return x
    }
    return y
}

func Maxu64(x, y uint64) uint64 {
    if x > y {
        return x
    }
    return y
}

// Both values, first is lesser, second is greater

func MinMaxi(x, y int) (int, int) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxi8(x, y int8) (int8, int8) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxi16(x, y int16) (int16, int16) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxi32(x, y int32) (int32, int32) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxi64(x, y int64) (int64, int64) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxu(x, y uint) (uint, uint) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxu8(x, y uint8) (uint8, uint8) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxu16(x, y uint16) (uint16, uint16) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxu32(x, y uint32) (uint32, uint32) {
    if x < y {
        return x, y
    }
    return y, x
}

func MinMaxu64(x, y uint64) (uint64, uint64) {
    if x < y {
        return x, y
    }
    return y, x
}
