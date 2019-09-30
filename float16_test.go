/*

 go-float16 - IEEE 754 binary16 half precision format
 Written in 2013 by h2so5 <mail@h2so5.net>

 To the extent possible under law, the author(s) have dedicated all copyright and
 related and neighboring rights to this software to the public domain worldwide.
 This software is distributed without any warranty.
 You should have received a copy of the CC0 Public Domain Dedication along with this software.
 If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.

*/

package half

import (
	"math"
	"testing"
)

func getFloatTable() map[Float16]float32 {
	table := map[Float16]float32{
		0x3c00: 1,
		0x4000: 2,
		0xc000: -2,
		0x7bfe: 65472,
		0x7bff: 65504,
		0xfbff: -65504,
		0x0000: 0,
		0x8000: float32(math.Copysign(0, -1)),
		0x7c00: float32(math.Inf(1)),
		0xfc00: float32(math.Inf(-1)),
		0x5b8f: 241.875,
		0x48c8: 9.5625,
	}
	return table
}
var float32array=[]float32{1,2,-2,65472,65504,-65504,0, float32(math.Copysign(0, -1)),float32(math.Inf(1)), float32(math.Inf(-1)),241.875, 9.5625}
var float16array=[]Float16{0x3c00,0x4000 ,0xc000 ,0x7bfe ,0x7bff ,0xfbff , 0x0000,0x8000 , 0x7c00,0xfc00,0x5b8f,0x48c8}
func TestFloat32(t *testing.T) {
	for k, v := range getFloatTable() {
		f := k.Float32()
		if f != v {
			t.Errorf("ToFloat32(%d) = %f, want %f.", k, f, v)
		}
	}
}
func TestFloat16Array(t *testing.T){
new16:=NewFloat16Array(float32array)
	for i:=range float32array{
	
		if	new16[i]!=float16array[i]{
			t.Errorf("FromFloat32(%f) = %d, want %d.", float32array[i], new16[i], float16array[i])	
		}

	}

}
func TestFloat32Array(t *testing.T){
	new32:=ToFloat32(float16array)
	for i:=range float16array{
	
		if new32[i]!=float32array[i]{
			t.Errorf("ToFloat32(%d) = %f, want %f.", float16array[i], new32[i], float32array[i])	
		}

	}
}
func TestNewFloat16(t *testing.T) {
	for k, v := range getFloatTable() {
		i := NewFloat16(v)
		if i != k {
			t.Errorf("FromFloat32(%f) = %d, want %d.", v, i, k)
		}
	}
}
