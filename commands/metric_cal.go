package commands

import (
	"math"

	"github.com/khanhtc1202/chio/entity"
)

type MetricCalculator struct{}

func (*MetricCalculator) Instability(module entity.Module) float64 {
	fOut := module.GetFanOutDepend()
	fIn := module.GetFanInDepend()
	return float64(fOut) / float64(fOut+fIn)
}

func (*MetricCalculator) Abstractness(module entity.Module) float64 {
	abs := module.GetAbstractMember()
	con := module.GetConcreteMember()
	return float64(abs) / float64(abs+con)
}

func (*MetricCalculator) Distance(instability, abstractness float64) float64 {
	return math.Abs(instability + abstractness - 1)
}
