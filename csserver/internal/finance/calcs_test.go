package finance

import (
	"fmt"
	"testing"
)

// https://www.calculatestuff.com/financial/npv-calculator

var npvTestCases = []struct {
	ok     bool
	r      float64
	values []float64
	want   float64
}{
	{
		ok:     true,
		r:      0.281,
		values: []float64{-100, 39, 59, 55, 20},
		want:   -0.00847859163845488,
	},
	{
		ok:     false,
		r:      0.281,
		values: nil,
		want:   0,
	},
	{
		ok:     false,
		r:      0.281,
		values: []float64{42},
		want:   0,
	},
	{
		ok:     true,
		r:      0.07, //---7 percent
		values: []float64{-96000.0, 1000.0, 1000.0, 2000.0, 1000.0, 4000.0},
		want:   -88944.54614871055,
	},
}

func TestNPV(t *testing.T) {
	for _, tt := range npvTestCases {
		got, err := NPV(tt.r, tt.values)

		if !tt.ok {
			if err == nil || got != tt.want {
				t.Errorf("NPV(%v, %v) = %v, %v; want %v, %v",
					tt.r,
					tt.values,
					got,
					err,
					tt.want,
					"<error>",
				)
			}
			continue
		}

		if got != tt.want {
			t.Errorf("NPV(%v, %v) = %v, %v; want %v, %v",
				tt.r,
				tt.values,
				got,
				err,
				tt.want,
				nil,
			)
		}
	}
}

var npvBenchResult float64

func BenchmarkNPV(b *testing.B) {
	var res float64

	r := 0.281
	values := []float64{-100, 39, 59, 55, 20}

	for n := 0; n < b.N; n++ {
		res, _ = NPV(r, values)
	}

	npvBenchResult = res
}

func ExampleNPV() {
	initialInvestment := 100.0

	cashFlowYear1 := 39.0
	cashFlowYear2 := 59.0
	cashFlowYear3 := 55.0
	cashFlowYear4 := 20.0

	discountRate := 0.281

	npv, err := NPV(discountRate, []float64{-initialInvestment, cashFlowYear1, cashFlowYear2, cashFlowYear3, cashFlowYear4})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("NPV is: %f", npv)
	// Output:
	// NPV is: -0.008479
}

var irrTestCases = []struct {
	ok     bool
	values []float64
	want   float64
}{
	{
		ok:     true,
		values: []float64{-100, 39, 59, 55, 20},
		want:   0.2809484211599531,
	},
	{
		ok:     false,
		values: nil,
		want:   0,
	},
	{
		ok:     false,
		values: []float64{42},
		want:   0,
	},
}

func TestIRR(t *testing.T) {
	for _, tt := range irrTestCases {
		got, err := IRR(tt.values)

		if !tt.ok {
			if err == nil || got != tt.want {
				t.Errorf("IRR(%v) = %v, %v; want %v, %v",
					tt.values,
					got,
					err,
					tt.want,
					"<error>",
				)
			}
			continue
		}

		if got != tt.want {
			t.Errorf("IRR(%v) = %v, %v; want %v, %v",
				tt.values,
				got,
				err,
				tt.want,
				nil,
			)
		}
	}
}

var irrBenchResult float64

func BenchmarkIRR(b *testing.B) {
	var r float64

	values := []float64{-100, 39, 59, 55, 20}

	for n := 0; n < b.N; n++ {
		r, _ = IRR(values)
	}

	irrBenchResult = r
}

func ExampleIRR() {
	initialInvestment := 100.0

	cashFlowYear1 := 39.0
	cashFlowYear2 := 59.0
	cashFlowYear3 := 55.0
	cashFlowYear4 := 20.0

	irr, err := IRR([]float64{-initialInvestment, cashFlowYear1, cashFlowYear2, cashFlowYear3, cashFlowYear4})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("IRR is: %f", irr)
	// Output:
	// IRR is: 0.280948
}

var fvTestCases = []struct {
	presentValue       float64
	interestRate       float64
	periods            float64
	compoundedAnnually bool
	want               float64
}{
	{
		presentValue:       100.0,
		interestRate:       0.07,
		periods:            10.0,
		compoundedAnnually: false,
		want:               170.00000000000003,
	},
	{
		presentValue:       100.0,
		interestRate:       0.07,
		periods:            10.0,
		compoundedAnnually: true,
		want:               196.71513572895657,
	},
}

func TestFV(t *testing.T) {
	for _, tt := range fvTestCases {
		got := FV(tt.presentValue, tt.interestRate, tt.compoundedAnnually, tt.periods)
		if got != tt.want {
			t.Errorf("FV(%v, %v, %v) = %v; want %v",
				tt.presentValue,
				tt.interestRate,
				tt.periods,
				got,
				tt.want,
			)
		}
	}
}

var fvBenchResult float64

func BenchmarkFV(b *testing.B) {
	var r float64

	presentValue := 100.0
	interestRate := 0.07
	periods := 10.0
	compoundedAnnually := false

	for n := 0; n < b.N; n++ {
		r = FV(presentValue, interestRate, compoundedAnnually, periods)
	}

	fvBenchResult = r
}

func ExampleFV() {
	presentValue := 100.0
	interestRate := 0.07
	periods := 10.0
	compoundedAnnually := false

	fv := FV(presentValue, interestRate, compoundedAnnually, periods)
	fmt.Printf("FV is: %f", fv)
	// Output:
	// FV is: 170.000000
}
