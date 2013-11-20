/*

 go-float16 - IEEE 754 binary16 half precision format
 Written in 2013 by h2so5 <mail@h2so5.net>

 To the extent possible under law, the author(s) have dedicated all copyright and
 related and neighboring rights to this software to the public domain worldwide.
 This software is distributed without any warranty.
 You should have received a copy of the CC0 Public Domain Dedication along with this software.
 If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.

*/

package float16

import (
	"math"
)

type Float16 uint16

func ToFloat32(f16 Float16) float32 {
	sign := uint32((f16 >> 15) & 0x1)
	exp := (f16 >> 10) & 0x1f
	exp32 := uint32(exp) + 127 - 15
	if exp == 0 {
		exp32 = 0
	} else if exp == 0x1f {
		exp32 = 0xff
	}
	frac := uint32(f16 & 0x3ff)
	i := (sign << 31) | (exp32 << 23) | (frac << 13)
	return math.Float32frombits(i)
}

func FromFloat32(f float32) Float16 {
	i := math.Float32bits(f)
	sign := uint16((i >> 31) & 0x1)
	exp := (i >> 23) & 0xff
	exp16 := int16(exp) - 127 + 15
	frac := uint16(i>>13) & 0x3ff
	if exp == 0 {
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
	}
	f16 := (sign << 15) | uint16(exp16<<10) | frac
	return Float16(f16)
}
