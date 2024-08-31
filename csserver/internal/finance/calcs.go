package finance

/*
https://github.com/orcaman/financial/tree/master
https://medium.com/@_orcaman/package-financial-for-golang-the-math-behind-the-irr-function-1eedf225d9f
*/

import (
	"errors"
	"math"
)

const (
	irrMaxInterations = 20
	irrAccuracy       = 1e-7
	irrInitialGuess   = 0
)

// IRR returns the Internal Rate of Return (IRR).
func IRR(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("values must include the initial investment (usually negative number) and period cash flows")
	}
	x0 := float64(irrInitialGuess)
	var x1 float64
	for i := 0; i < irrMaxInterations; i++ {
		fValue := float64(0)
		fDerivative := float64(0)
		for k := 0; k < len(values); k++ {
			fk := float64(k)
			fValue += values[k] / math.Pow(1.0+x0, fk)
			fDerivative += -fk * values[k] / math.Pow(1.0+x0, fk+1.0)
		}
		x1 = x0 - fValue/fDerivative
		if math.Abs(x1-x0) <= irrAccuracy {
			return x1, nil
		}
		x0 = x1
	}
	return 0, errors.New("could not find irr for the provided values")
}

// FV gets the present value, the interest rate (compounded annually if needed) and the number of periods and returns the future value.
func FV(pv, r float64, compoundedAnnually bool, n float64) float64 {
	if !compoundedAnnually {
		return pv * (1 + (r * n))
	}
	return pv * math.Pow(1+r, n)
}

// NPV returns the NPV (Net Present Value) of a cash flow series given an interest rate r, or an error.
func NPV(discountRate float64, values []float64) (float64, error) {
	if len(values) < 2 {
		return 0, errors.New("values must include the initial investment (usually negative number) and period cash flows")
	}

	npv := values[0]
	values = values[1:]

	for i, v := range values {
		npv += (v / math.Pow(1+discountRate, float64(i+1)))
	}

	return npv, nil
}
