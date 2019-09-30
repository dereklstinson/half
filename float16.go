/*

 go-float16 - IEEE 754 binary16 half precision format
 Written in 2013 by h2so5 <mail@h2so5.net>

 To the extent possible under law, the author(s) have dedicated all copyright and
 related and neighboring rights to this software to the public domain worldwide.
 This software is distributed without any warranty.
 You should have received a copy of the CC0 Public Domain Dedication along with this software.
 If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.

*/

// Package half is an IEEE 754 binary16 half precision format.
package half

import "math"

// A Float16 represents a 16-bit floating point number.
type Float16 uint16

// NewFloat16 allocates and returns a new Float16 set to f.
func NewFloat16(f float32) Float16 {
	i := math.Float32bits(f)
	sign := uint16((i >> 31) & 0x1)
	exp := (i >> 23) & 0xff
	exp16 := int16(exp) - 112
	frac := uint16(i>>13) & 0x3ff
/*	if exp == 0 {
		exp16 = 0
	} else if exp == 0xff {
		exp16 = 0x1f
	} else {
		if exp16 > 0x1e {
			exp16 = 0x1f
			frac = 0
		} else if exp16 < 0x01 {
			exp16 = 0
			frac = 0
		}
	}*/
	switch exp{
	case 0:
		exp16 = 0
	case 0xff:
		exp16 = 0x1f
	default:
		if exp16 > 0x1e {
			exp16 = 0x1f
			frac = 0
		} else if exp16 < 0x01 {
			exp16 = 0
			frac = 0
		}
	}

	return (Float16)((sign << 15) | uint16(exp16<<10) | frac)
}
//NewFloat16Array creates a Float16 array from a float32 array
func NewFloat16Array(array []float32)[]Float16{
	var (
		i uint32
		sign uint16
		exp uint32
	exp16 int16
	frac uint16
	)
	array16:=make([]Float16,len(array))
	for idx:=range array{
		i = math.Float32bits(array[idx])
		sign = uint16((i >> 31) & 0x1)
		exp = (i >> 23) & 0xff
		exp16 = int16(exp) -112
		frac = uint16(i>>13) & 0x3ff
		switch exp{
		case 0:
			exp16 = 0
		case 0xff:
			exp16 = 0x1f
		default:
			if exp16 > 0x1e {
				exp16 = 0x1f
				frac = 0
			} else if exp16 < 0x01 {
				exp16 = 0
				frac = 0
			}
		}
	
		array16[idx]= (Float16)((sign << 15) | uint16(exp16<<10) | frac)
	}
	return array16
}
// Float32 returns the float32 representation of f.
func (f Float16) Float32() float32 {
	sign := uint32((f >> 15) & 0x1)
	exp := (f >> 10) & 0x1f
	exp32 := uint32(exp) + 127 - 15
	if exp == 0 {
		exp32 = 0
	} else if exp == 0x1f {
		exp32 = 0xff
	}
	return math.Float32frombits((sign << 31) | (exp32 << 23) | ( (uint32)(f & 0x3ff) << 13))
}
//ToFloat32 takes an []Float16 and returns a []float32 
func ToFloat32(array []Float16)[]float32{
	var (
		sign uint32
		exp uint16
		exp32 uint32
		f uint16	
	)
	array32:=make([]float32,len(array))
	for i:=range array32{
		f=(uint16)(array[i])
		sign = uint32((f >> 15) & 0x1)
	exp = (f >> 10) & 0x1f
	exp32 = uint32(exp) + 127 - 15
	if exp == 0 {
		exp32 = 0
	} else if exp == 0x1f {
		exp32 = 0xff
	}
	array32[i]= math.Float32frombits((sign << 31) | (exp32 << 23) | ( (uint32)(f & 0x3ff) << 13))
	}
	return array32
}
