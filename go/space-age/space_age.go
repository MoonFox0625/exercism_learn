package space

type Planet string

var factorMap = map[Planet]float64{
	"Mercury": 0.2408467,
	"Earth":   1,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	var orbitalFactor float64 // 轨道系数
	const EarthSeconds = 31557600
	orbitalFactor = factorMap[planet]
	return seconds / EarthSeconds / orbitalFactor
}
