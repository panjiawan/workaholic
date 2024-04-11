package putils

import "github.com/shopspring/decimal"

func StringToInt64(num string) (int64, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return dec.IntPart(), nil
}

func StringToInt32(num string) (int32, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return int32(dec.IntPart()), nil
}

func StringToInt(num string) (int, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return int(dec.IntPart()), nil
}

func StringToUint64(num string) (uint64, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return uint64(dec.IntPart()), nil
}

func StringToFloat64(num string) float64 {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0
	}

	return dec.InexactFloat64()
}

func StringSliceToUInt64Slice(slice []string) []uint64 {
	iSlice := make([]uint64, 0, len(slice))
	for _, v := range slice {
		if dec, err := decimal.NewFromString(v); err == nil {
			iSlice = append(iSlice, dec.BigInt().Uint64())
		}
	}
	return iSlice
}

func StringSliceToInt64Slice(slice []string) []int64 {
	iSlice := make([]int64, 0, len(slice))
	for _, v := range slice {
		if dec, err := decimal.NewFromString(v); err == nil {
			iSlice = append(iSlice, dec.IntPart())
		}
	}
	return iSlice
}

func StringSliceToFloat64Slice(slice []string) []float64 {
	iSlice := make([]float64, 0, len(slice))
	for _, v := range slice {
		if dec, err := decimal.NewFromString(v); err == nil {
			fnum, _ := dec.Float64()
			iSlice = append(iSlice, fnum)
		}
	}
	return iSlice
}
