# go-imath
Package containing some integral constants and elementary functions for integral arguments.

## Summary
Go standard `math` package lacks some frequently used integral functions. This package tries to fill the gap, providing their effective and convenient implementations.

Functions have the same or similar names as their `float64` equivalents from `math` package (if there are any). Since Go does not allow overloading, constants and functions related to each of builtin integral types (except `uintptr` and type aliases like `byte` or `rune`, which are not covered by this package) are grouped within subpackage with a corresponding name (`ix` for `int`, `i8` for `int8`, ...; `ux` for `uint`, `u8` for `uint8`, ...). Functions defined only for few of the types (like `Power` functions) are kept in the root `imath` package.

## Import
```go
import (
	"github.com/adam-lavrik/go-imath/ix" // int-related functions
	"github.com/adam-lavrik/go-imath/u32" // uint32-related function
	...
)
```

## Types
* `T` - integer type, package is centered on (`uint8` for `u8`, `int` for `ix`, ...)
* `UT` - unsigned type, complementary to signed `T` (`uint8` for `int8`, `uint` for `int`, ...)
* `ST` - signed type, complementary to unsigned `T` (`int8` for `uint8`, `int` for `uint`, ...)

## Constants
* `Size` - size of type value in bytes
* `BitSize` - same in bits
* `Minimal` - minimal possible value of type
* `Maximal` - maximal possible value of type

### Platform dependent
Name|Type|Value (32-bit)|Value (64-bit)
-|:-:|:-:|:-:
`ix.Size`|`uintptr`|4|8
`ix.BitSize`|`uintptr`|32|64
`ix.Minimal`|`int`|-2147483648|-9223372036854775808
`ix.Maximal`|`int`|2147483647|9223372036854775807
`ux.Size`|`uintptr`|4|8
`ux.BitSize`|`uintptr`|32|64
`ux.Minimal`|`uint`|0|0
`ux.Maximal`|`uint`|4294967295|18446744073709551615

### Platform independent
Name|Type|Value
-|:-:|:-:
`i8.Size`|`uintptr`|1
`i8.BitSize`|`uintptr`|8
`i8.Minimal`|`int8`|-128
`i8.Maximal`|`int8`|127
`u8.Size`|`uintptr`|1
`u8.BitSize`|`uintptr`|8
`u8.Minimal`|`uint8`|0
`u8.Maxnimal`|`uint8`|255
`i16.Size`|`uintptr`|2
`i16.BitSize`|`uintptr`|16
`i16.Minimal`|`int16`|-32768
`i16.Maximal`|`int16`|32767
`u16.Size`|`uintptr`|2
`u16.BitSize`|`uintptr`|16
`u16.Minimal`|`uint16`|0
`u16.Maximal`|`uint16`|65535
`i32.Size`|`uintptr`|4
`i32.BitSize`|`uintptr`|32
`i32.Minimal`|`int32`|-2147483648
`i32.Maximal`|`int32`|2147483647
`u32.Size`|`uintptr`|4
`u32.BitSize`|`uintptr`|32
`u32.Minimal`|`uint32`|0
`u32.Maximal`|`uint32`|4294967295
`i64.Size`|`uintptr`|8
`i64.BitSize`|`uintptr`|64
`i64.Minimal`|`int64`|-9223372036854775808
`i64.Maximal`|`int64`|9223372036854775807
`u64.Size`|`uintptr`|8
`u64.BitSize`|`uintptr`|64
`u64.Minimal`|`uint64`|0
`u64.Maximal`|`uint64`|18446744073709551615

## Functions

### i#.Abs(value int#) int#
Absolute value of signed integer. Argument and result have the same type.

__Examples__:
```go
a0 := ix.Abs(-16) // a0 == int(16)
a1 := i8.Abs(42) // a1 == int8(42)
```

__Important:__ since the lowest negative values of integral signed types have no corresponding positive values, result of `Abs(i#.Minimal)` is falsely equal to the argument:

```go
a2 := i16.Abs(-32768) // a2 == int16(-32768)
```
### i#.Absu(value int#) uint#
Absolute value of signed integer. Argument is of signed integer type, and result is of the complementary unsigned integer type thus can hold all positive values corresponding to all possible negative values of argument type.

__Examples__:
```go
a0 := ix.Absu(-16) // a0 == uint(16)
a1 := i8.Absu(42) // a1 == uint8(42)
a2 := i16.Absu(-32768) // a2 == uint16(32768)
```
### i#.Copysign(target, source int#) int#
Value with magnitude of `target` and sign of `source`. Zero `target` always produces zero result. Non-zero `target` and zero `source` produce positive result.

For an obvious reason these functions take only signed arguments.

__Examples__:
```go
c0 := ix.Copysign(16, 0) // c0 == int(16)
c1 := i8.Copysign(-38, 1) // c1 == int8(38)
c2 := i64.Copysing(42, -1) // c2 == int64(-42)
```
__Important:__ since the lowest negative values of integral signed types have no corresponding positive values, their sign cannot be set to positive. This limitation cannot be avoided.

### #.DivMod(dividend, divisor #) (#, #)
Quotient and remainder of `dividend` and `divisor`. Zero `divisor` causes division by zero error.

__Examples__:
```go
q0, d0 := ux.DivMod(15, 8) // q0 == uint(1), d0 == uint(7)
q1, d1 := i16.DivMod(320, 64) // q1 == int16(5), d1 == int16(0)
q2, d2 := i64.DivMod(79, 0) // Error
```

### #.GCD(value_0, value_1 #) uint#
Greatest common divisor of two integers. Arguments can be of signed or unsigned integer type, while result is always unsigned (same type for unsigned arguments or complementary unsigned for signed ones).

__Examples__:
```go
g0 = ix.GCD(48, 32) // g0 == uint(16)
g1 = u16.GCD(45, 0) // g1 == uint16(45)
g2 = i8.GCD(-42, -8) // g2 == uint8(2)
g3 = i32.GCD(-28, 21) // g3 == uint32(7)
```

### #.Is2Power(value #) bool
Check whether the value is an integer power of 2 (`true`) or not (`false`). Zero and negative values always produce `false`.

__Examples__:
```go
i2p0 := ix.Is2Power(14) // i2p0 == false
i2p1 := u8.Is2Power(128) // i2p1 == true
i2p2 := i16.Is2Power(-7) // i2p2 == false
```

### #.IsOdd(value #) bool
Check whether the value is odd (`true`) or not (`false`).
```go
io0 := ux.IsOdd(742) // io0 == false
io1 := u8.IsOdd(0) // io1 == true
io2 := i32.IsOdd(-11345) // io2 == true
```

### #.LCM(value_0, value_1 #) uint#
Least common multiply of two integers. Arguments can be of signed or unsigned integer type, while result is always unsigned (same type for unsigned arguments or complementary unsigned for signed ones).

__Examples__:
```go
l0 = ix.LCM(48, 32) // l0 == uint(96)
l1 = u16.LCM(45, 0) // l1 == uint16(0)
l2 = i8.LCM(-42, -8) // l2 == uint8(168)
l3 = i32.LCM(-28, 21) // l3 == uint32(84)
```

### #.Min(value_0, value_1 #) #
Minimal of two integers.

__Examples__:
```go
mi0 := ix.Min(13, -42) // mi0 == int(-42)
mi1 := u8.Min(255, 255) // mi1 == uint8(255)
```

### #.Mins(value #, values ...#) #
Minimal of one or more integers.

__Examples__:
```go
mis0 := ix.Mins(-42) // mis0 == int(-42)
mis1 := u8.Mins(255, 255, 14) // mis1 == uint8(14)
```

### #.MinSlice(values []#) #
Minimal of slice items. Empty slice causes error.

__Examples__:
```go
mis := ix.MinSlice([]int{1, - 3, -42}) // mis == int(-42)
```

### #.MinSliceChecked(values []#) (#, bool)
Minimal of slice items. For an empty slice second result is `true`. Otherwise it is `false`, and first result is meaningful.

__Examples__:
```go
misc, empty := ix.MinSliceChecked([]int{1, - 3, -42}) // misc == int(-42), empty == false
```

### #.Max(value_0, value_1 #) #
Maximal of two integers.

__Examples__:
```go
ma0 := ix.Max(13, -42) // ma0 == int(13)
ma1 := u8.Max(255, 255) // ma1 == uint8(255)
```

### #.Maxs(value #, values ...#) #
Maximal of one or more integers.

__Examples__:
```go
mas0 := ix.Maxs(-42) // mas0 == int(-42)
mas1 := u8.Maxs(255, 255, 14) // mas1 == uint8(255)
```

### #.MaxSlice(values []#) #
Maximal of slice items. Empty slice causes error.

__Examples__:
```go
mas := ix.MaxSlice([]int{1, - 3, -42}) // mas == int(1)
```

### #.MaxSliceChecked(values []#) (#, bool)
Maximal of slice items. For an empty slice second result is `true`. Otherwise it is `false`, and first result is meaningful.

__Examples__:
```go
masc, empty := ix.MaxSliceChecked([]int{1, - 3, -42}) // masc == int(1), empty == false
```

### #.MinMax(value_0, value_1 #) (#, #)
Minimal (first) and maximal (second) of two integers.

__Examples__:
```go
min0, max0 := ix.MinMax(13, -42) // min0 == int(-42), max0 == int(13)
min1, max1 := u8.MinMax(255, 255) // min1 == uint8(255), max1 == uint8(255)
```

### #.MinMaxs(value #, values ...#) (#, #)
Minimal (first) and maximal (second) of one or more integers.

__Examples__:
```go
mins1, maxs1 := u8.MinMaxs(255, 255, 14) // mins1 == uint8(14), maxs1 == uint8(255)
```

### #.MinMaxSlice(values []#) (#, #)
Minimal (first) and maximal (second) of slice items. Empty slice causes error.

__Examples__:
```go
mins, maxs := ix.MinMaxSlice([]int{1, - 3, -42}) // mins == int(-42), maxs == int(1)
```

### #.MinMaxSliceChecked(values []#) (#, #, bool)
Minimal (first) and maximal (second) of slice items. For an empty slice third result is `true`. Otherwise it is `false`, and first and second results are meaningful.

__Examples__:
```go
minsc, maxsc, empty := ix.MaxSliceChecked([]int{1, - 3, -42}) // minsc == int(-42), maxsc == int(1), empty == false
```

### i#.Sign (value int#) int#
Signum of signed integer:
* -1 if value < 0
* 0 if value = 0
* 1 if value > 0

__Examples__:
```go
s0 := ix.Sign(99) // s0 == int(1)
s1 := i16.Sign(0) // s1 == int16(0)
s2 := i64.Sign(-1234) // s2 == int64(-1)
```

### u#.Sign (value uint#) uint#
Signum of unsigned integer:
* 0 if value = 0
* 1 if value > 0

__Examples__:
```go
s0 := ux.Sign(99) // s0 == uint(1)
s1 := u16.Sign(0) // s1 == uint16(0)
```

### i#.SignBit (value int#) int#
Result of the same signed type as `value` with all bits equal to `value`'s highest bit:
* -1 if value < 0
* 0 if value >= 0

__Examples__:
```go
s0 := ix.SignBit(99) // s0 == int(0)
s1 := i16.SignBit(0) // s1 == int16(0)
s2 := i64.SignBit(-1234) // s2 == int64(-1)
```

### ix.Fibonacci(index int) int
### i64.Fibonacci(index int) int64
Fibonacci sequence member with corresponding `index`. `index` can be zero or negative.

__Examples__:
```go
fi0 := ix.Fibonacci(0) // fi0 == int(0)
fi1 := i64.Fibonacci(5) // fi1 == int64(5)
fi2 := ix.Fibonacci(-4) // fi2 == int(-3)
fi3 := i64.Fibonacci(-5) // fi3 == int64(5)
```

### ux.Fibonacci(index uint) uint
### u64.Fibonacci(index uint) uint64
Fibonacci sequence member with corresponding `index`. `index` can be zero.

__Examples__:
```go
fu0 := ix.Fibonacci(0) // fu0 == int(0)
fu1 := i64.Fibonacci(6) // fu == int64(8)
```

### ix.Pow(base int, exponent uint) int
### i64.Pow(base int64, exponent uint) int64
### ux.Pow(base uint, exponent uint) uint
### u64.Pow(base uint64, exponent uint) uint64
Exponentiation of integral `base` and non-negative `exponent`.

__Important__:
* if `exponent` is 0, result is always 1, even if `base` is 0 too;
* since result values grow rapidly, this function is implemented only for `int`, `uint`, `int64` and `uint64` types.

__Examples__:
```go
p0 := ux.Pow(17, 2) // p0 == uint(289)
p1 := i64.Pow(-41, 7) // p1 == int64(-194754273881)
p2 := u64.Pow(0, 5) // p2 == uint64(0)
p3 := ix.Pow(-122, 0) // p3 == int(1)
```

(c) Adam Lavrik <lavrik.adam@gmail.com>, 2020
