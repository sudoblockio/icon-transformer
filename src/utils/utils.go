package utils

import (
	"fmt"
	"go.uber.org/zap"
	"math/big"
	"reflect"
	"runtime"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringHexToFloat64(hex string, base int) float64 {
	valueDecimal := float64(0)

	var negative bool
	if hex[:1] == "-" {
		hex = hex[1:]
		negative = true
	}

	valueBigInt, success := new(big.Int).SetString(hex[2:], 16)
	if success == false {
		zap.S().Warn("Set String Error: hex=", hex)
		return 0
	}

	baseBigFloatString := "1"
	for i := 0; i < base; i++ {
		baseBigFloatString += "0"
	}
	baseBigFloat, success := new(big.Float).SetString(baseBigFloatString) // 10^(base)
	if success == false {
		zap.S().Warn("Set String Error: base=", base)
		return 0
	}

	valueBigFloat := new(big.Float).SetInt(valueBigInt)
	valueBigFloat = valueBigFloat.Quo(valueBigFloat, baseBigFloat)

	valueDecimal, _ = valueBigFloat.Float64()

	if negative {
		valueDecimal = -1 * valueDecimal
	}

	return valueDecimal
}

func StringHexToInt64(i string) int64 {
	o := new(big.Int)
	_, err := fmt.Sscan(i, o)
	if err != nil {
		zap.S().Warn("String conversion failure Error: ", err.Error())
		return 0
	}

	return o.Int64()
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
