package domain

import "math"

type MetricCalculator struct{}

func (*MetricCalculator) Instability(module Module) float64 {
	fOut := module.CountFanOutDepend()
	fIn := module.CountFanInDepend()
	return float64(fOut) / float64(fOut+fIn)
}

func (*MetricCalculator) Abstractness(module Module) float64 {
	abs := module.CountAbstractMember()
	con := module.CountConcreteMember()
	return float64(abs) / float64(abs+con)
}

func (*MetricCalculator) Distance(instability, abstractness float64) float64 {
	return math.Abs(instability + abstractness - 1)
}
