package tool

import (
	"errors"
	"math"

	"github.com/sineycoder/go-bigger/types"
)

/**
 @author: nizhenxian
 @date: 2021/8/12 14:22:00
**/

func Copy(original []types.Int, length types.Int) []types.Int {
	if length >= types.Int(len(original)) {
		return CopyRange(original, 0, length)
	} else {
		dest := make([]types.Int, length)
		CopyRangePosLen(original, 0, dest, 0, types.Int(len(original)))
		return dest
	}
}

func CopyRange(original []types.Int, from, to types.Int) []types.Int {
	newLength := to - from
	if newLength < 0 {
		panic(errors.New("invalid params"))
	}
	l := types.Int(math.Min(float64(len(original)), float64(newLength)))
	cp := make([]types.Int, l)
	copy(cp, original[from:from+l])
	return cp
}

func CopyRangePosLen(src []types.Int, srcPos types.Int, dest []types.Int, destPos, length types.Int) {
	srcLen := types.Int(len(src))
	destLen := types.Int(len(dest))
	if srcPos+length > srcLen || destPos+length > destLen {
		panic("array out of index")
	}
	for ind := types.Int(0); ind < length; ind++ {
		dest[destPos+ind] = src[srcPos+ind]
	}
}

func Fill(a []types.Int, fromIndex, toIndex, val types.Int) {
	for i := fromIndex; i < toIndex; i++ {
		a[i] = val
	}
}

func MaxInt(a, b types.Int) types.Int {
	return MaxLong(a.ToLong(), b.ToLong()).ToInt()
}

func MaxLong(a, b types.Long) types.Long {
	if a >= b {
		return a
	} else {
		return b
	}
}

func MinInt(a, b types.Int) types.Int {
	return MinLong(a.ToLong(), b.ToLong()).ToInt()
}

func MinLong(a, b types.Long) types.Long {
	if a <= b {
		return a
	} else {
		return b
	}
}

func IntEqual(a, b []types.Int) bool {
	if ((a == nil) != (b == nil)) || (len(a) != len(b)) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Digit(a uint8, radix uint8) types.Int {
	if radix < 2 || radix > 36 {
		return -1
	}
	if a >= 48 && a <= 57 {
		if a-48 < radix {
			return types.Int(a - 48)
		}
	} else if a >= 65 && a <= 90 {
		if a-55 < radix {
			return types.Int(a - 55)
		}
	}
	return -1
}

func Arraycopy(src []types.Int, srcPos types.Int, dest []types.Int, destPos, length types.Int) {
	for i := types.Int(0); i < length; i++ {
		dest[destPos+i] = src[srcPos+i]
	}
}

func ByteToInt(val []byte) (ret []types.Int) {
	ret = make([]types.Int, len(val))
	for idx, _ := range val {
		ret[idx] = types.Int(val[idx])
	}
	return ret
}
