// Copyright 2020 Adam Lavrik <lavrik.adam@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//  http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package imath

import "unsafe"

const (
	// Limit values of integer types

	MinUint = uint(0)
	MaxUint = ^MinUint
	MaxInt = int(MaxUint >> 1)
	MinInt = -MaxInt - 1

	MinUint8 = uint8(0)
	MaxUint8 = ^MinUint8
	MaxInt8 = int8(MaxUint8 >> 1)
	MinInt8 = -MaxInt8 - 1

	MinUint16 = uint16(0)
	MaxUint16 = ^MinUint16
	MaxInt16 = int16(MaxUint16 >> 1)
	MinInt16 = -MaxInt16 - 1

	MinUint32 = uint32(0)
	MaxUint32 = ^MinUint32
	MaxInt32 = int32(MaxUint32 >> 1)
	MinInt32 = -MaxInt32 - 1

	MinUint64 = uint64(0)
	MaxUint64 = ^MinUint64
	MaxInt64 = int64(MaxUint64 >> 1)
	MinInt64 = -MaxInt64 - 1

	// Sizes of integer types

	IntSize = unsafe.Sizeof(int(0))
	Int8Size = unsafe.Sizeof(int8(0))
	Int16Size = unsafe.Sizeof(int16(0))
	Int32Size = unsafe.Sizeof(int32(0))
	Int64Size = unsafe.Sizeof(int64(0))

	IntBitSize = IntSize << 3
	Int8BitSize = Int8Size << 3
	Int16BitSize = Int16Size << 3
	Int32BitSize = Int32Size << 3
	Int64BitSize = Int64Size << 3

	UintSize = unsafe.Sizeof(uint(0))
	Uint8Size = unsafe.Sizeof(uint8(0))
	Uint16Size = unsafe.Sizeof(uint16(0))
	Uint32Size = unsafe.Sizeof(uint32(0))
	Uint64Size = unsafe.Sizeof(uint64(0))

	UintBitSize = UintSize << 3
	Uint8BitSize = Uint8Size << 3
	Uint16BitSize = Uint16Size << 3
	Uint32BitSize = Uint32Size << 3
	Uint64BitSize = Uint64Size << 3
)
