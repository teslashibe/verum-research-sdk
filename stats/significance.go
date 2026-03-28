package stats

import "math"

func PValue(mean, stdDev float64, n int) float64 {
	if n <= 1 || stdDev == 0 {
		return 1.0
	}
	se := stdDev / math.Sqrt(float64(n))
	t := math.Abs(mean / se)
	df := float64(n - 1)
	return twoTailedTProb(t, df)
}

func IsSignificant(pValue, alpha float64) bool {
	return pValue < alpha
}

func twoTailedTProb(t, df float64) float64 {
	x := df / (df + t*t)
	return regularizedIncompleteBeta(df/2.0, 0.5, x)
}

func regularizedIncompleteBeta(a, b, x float64) float64 {
	if x < 0 || x > 1 {
		return 0
	}
	if x == 0 || x == 1 {
		return x
	}

	lnBeta := lgamma(a) + lgamma(b) - lgamma(a+b)
	front := math.Exp(math.Log(x)*a + math.Log(1-x)*b - lnBeta)

	if x < (a+1)/(a+b+2) {
		return front * betaCF(a, b, x) / a
	}
	return 1 - front*betaCF(b, a, 1-x)/b
}

func betaCF(a, b, x float64) float64 {
	const maxIter = 200
	const epsilon = 1e-14

	qab := a + b
	qap := a + 1
	qam := a - 1
	c := 1.0
	d := 1 - qab*x/qap
	if math.Abs(d) < epsilon {
		d = epsilon
	}
	d = 1 / d
	h := d

	for m := 1; m <= maxIter; m++ {
		mf := float64(m)

		num := mf * (b - mf) * x / ((qam + 2*mf) * (a + 2*mf))
		d = 1 + num*d
		if math.Abs(d) < epsilon {
			d = epsilon
		}
		c = 1 + num/c
		if math.Abs(c) < epsilon {
			c = epsilon
		}
		d = 1 / d
		h *= d * c

		num = -(a + mf) * (qab + mf) * x / ((a + 2*mf) * (qap + 2*mf))
		d = 1 + num*d
		if math.Abs(d) < epsilon {
			d = epsilon
		}
		c = 1 + num/c
		if math.Abs(c) < epsilon {
			c = epsilon
		}
		d = 1 / d
		delta := d * c
		h *= delta

		if math.Abs(delta-1) < epsilon {
			break
		}
	}
	return h
}

func lgamma(x float64) float64 {
	v, _ := math.Lgamma(x)
	return v
}
