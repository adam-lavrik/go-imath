# go-imath
Package containing some integral constants and elementary functions for integral arguments.

## Summary
Go standard `math` package lacks some frequently used integral functions. This package tries to fill the gap, providing their effective and convenient implementations.

Functions are named closely to ones for `float64` type from `math` package (if there are any). Since Go does not allow overloading, function names contain type suffixes based on their arguments and\or result: `i` for int, `u` for `uint`, `i8` for `int8`, `u64` for `uint64` and so on (`uintptr` and type aliases like `byte` or `rune` are not covered by this package). The suffixes are reffered below as:
* `i#` (for `int#` = `int`|`int8`|...);
* `u#` (for `uint#` = `uint`|`uint8`|...);
* `#` (for any of that types).

## Import
```go
import "github.com/adam-lavrik/go-imath"
```

## Constants

### Platform dependent
Name|Type|Value (64-bit)|Value (32-bit)
:-:|:-:|:-:|:-:
MinInt|int|-9223372036854775808|-2147483648
MaxInt|int|9223372036854775807|2147483647
MinUint|uint|0|0
MaxUint|uint|18446744073709551615|4294967295
IntSize|uintptr|8|4
IntBitSize|uintptr|64|32
UintSize|uintptr|8|4
UintBitSize|uintptr|64|32

### Platform independent
Name|Type|Value
:-:|:-:|:-:
MinInt8|int8|-128
MaxInt8|int8|127
MinUint8|uint8|0
MaxUint8|uint8|255
MinInt16|int16|-32768
MaxInt16|int16|32767
MinUint16|uint16|0
MaxUint16|uint16|65535
MinInt32|int32|-2147483648
MaxInt32|int32|2147483647
MinUint32|uint32|0
MaxUint32|uint32|4294967295
MinInt64|int64|-9223372036854775808
MaxInt64|int64|9223372036854775807
MinUint64|uint64|0
MaxUint64|uint64|18446744073709551615
Int8Size|uintptr|1
Int8BitSize|uintptr|8
Uint8Size|uintptr|1
Uint8BitSize|uintptr|8
Int16Size|uintptr|2
Int16BitSize|uintptr|16
Uint16Size|uintptr|2
Uint16BitSize|uintptr|16
Int32Size|uintptr|4
Int32BitSize|uintptr|32
Uint32Size|uintptr|4
Uint32BitSize|uintptr|32
Int64Size|uintptr|8
Int64BitSize|uintptr|64
Uint64Size|uintptr|8
Uint64BitSize|uintptr|64

## Functions

### Absi#(number int#) int#
Absolute value of signed integer. Argument and result have the same type.

__Examples__:
```go
a0 := imath.Absi(-16) // a0 == int(16)
a1 := imath.Absi8(42) // a1 == int8(42)
```

__Important:__ since the lowest negative values of integral signed types have no corresponding positive values, result of `Absi*(<lowest negative value of type i*>)` is falsely equal to the argument:

```go
a2 := imath.Absi16(-32768) // a2 == int16(-32768)
```
### Absu#(number int#) uint#
Absolute value of signed integer. Argument is of signed integer type, and result is of the complementary unsigned integer type thus can hold all positive values corresponding to all possible negative values of argument type.

__Examples__:
```go
a0 := imath.Absu(-16) // a0 == uint(16)
a1 := imath.Absu8(42) // a1 == uint8(42)
a2 := imath.Absu16(-32768) // a2 == uint16(32768)
```
### Copysigni#(target, source int#) int#
Value with magnitude of `target` and sign of `source`. Zero `target` always produces zero result. Non-zero `target` and zero `source` produce positive result.

For an obvious reason these functions take only signed arguments.

__Examples__:
```go
c0 := imath.Copysigni(16, 0) // c0 == int(16)
c1 := imath.Copysigni8(-38, 1) // c1 == int32(38)
c2 := imath.Copysingi64(42, -1) // c2 == int64(-42)
```
__Important:__ since the lowest negative values of integral signed types have no corresponding positive values, their sign cannot be set to positive. This limitation cannot be evaded.

### DivMod#(dividend, divisor #) (#, #)
Quotient and remainder of `dividend` and `divisor`. Zero `divisor` causes division by zero.

__Examples__:
```go
q0, d0 := imath.DivModu(15, 8) // q0 == uint(1), d0 == uint(7)
q1, d1 := imath.DivModi16(320, 64) // q1 == int16(5), d1 == int16(0)
q2, d2 := imath.DivModi64(79, 0) // Error
```

### GCD#(x, y #) uint#
The greatest common divisor of two integers. Arguments can be of signed or unsigned integer type, while result is always unsigned (same for unsigned argument type or complementary unsigned for signed one).

__Examples__:
```go
g0 = imath.GCDi(48, 32) // g0 == uint(16)
g1 = imath.GCDu16(45, 0) // g1 == uint16(45)
g2 = imath.GCDi8(-42, -8) // g2 == uint8(2)
g3 = imath.GCDi32(-28, 21) // g3 == uint32(7)
```
### Is2Power#(number #) bool
Check whether the number is an integer power of 2 (`true`) or not (`false`). Zero and negative numbers always produce `false`.

__Examples__:
```go
i2p0 := imath.Is2Poweri(14) // i2p0 == false
i2p1 := imath.Is2Poweru8(128) // i2p1 == true
i2p2 := imath.Is2Poweri16(-7) // i2p2 == false
```
### IsOdd#(value #) bool
Check whether the number is odd (`true`) or not (`false`). Zero is odd.
```go
io0 := imath.IsOddu(742) // io0 == false
io1 := imath.IsOddu8(0) // io1 == true
io2 := imath.IsOddi32(-11345) // io2 == true
```
### LCM#(x, y #) uint#
The least common multiply of two integers. Arguments can be of signed or unsigned integer type, while result is always unsigned (same for unsigned argument type or complementary unsigned for signed one).

__Examples__:
```go
l0 = imath.LCMi(48, 32) // l0 == uint(96)
l1 = imath.GCDu16(45, 0) // l1 == uint16(0)
l2 = imath.GCDi8(-42, -8) // l2 == uint8(168)
l3 = imath.GCDi32(-28, 21) // l3 == uint32(84)
```

### Min#(x, y #) #
Minimal of two integers.

__Examples__:
```go
mi0 := imath.Mini(13, -42) // mi0 == int(-42)
mi1 := imath.Minu8(255, 255) // mi1 == uint8(255)
```

### Max#(x, y #) #
Maximal of two integers.

__Examples__:
```go
ma0 := imath.Maxi(13, -42) // ma0 == int(13)
ma1 := imath.Maxu8(255, 255) // ma1 == uint8(255)
```

### MinMax#(x, y #) #
Minimal (first) and maximal (second) of two integers.

__Examples__:
```go
mi0, ma0 := imath.MinMaxi(13, -42) // mi0 == int(-42), ma0 == int(13)
mi1, ma1 := imath.MinMaxu8(255, 255) // mi1 == uint8(255), ma1 == uint8(255)
```

### Pow#(base #, exponent uint) #
Exponentiation of integer `base` and non-negative `exponent`.

__Important__:
* since result values grows rapidly, this function is implemented only for `int`, `uint`, `int64` and `uint64` types.
* if `exponent` is 0, result is always 1, even if `base` is 0 too.

__Examples__:
```go
p0 := imath.Poweru(17, 2) // p0 == uint(289)
p1 := imath.Poweri64(-41, 7) // p1 == int64(-194754273881)
p2 := imath.Poweru64(0, 5) // p2 == uint64(0)
p3 := imath.Poweri(-122, 0) // p3 == int(1)
```

### Signi# (number int#) int#
Signum of signed integer value:
* -1 if value < 0
* 0 if value = 0
* 1 if value > 0

__Examples__:
```go
s0 := Signi(99) // s0 == int(1)
s1 := Signi16(0) // s1 == int16(0)
s2 := Signi64(-1234) // s2 == int64(-1)
```

### Signu# (number uint#) uint#
Signum of unsigned integer value:
* 0 if value = 0
* 1 if value > 0

__Examples__:
```go
s0 := Signu(99) // s0 == uint(1)
s1 := Signu16(0) // s1 == uint16(0)
```
