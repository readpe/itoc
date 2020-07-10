package main

import "math"

// relay interface is satisfied with opTime method
type relay interface {
	opTime(float64) float64
}

// itoc curve parameters
// formula opTime = timedial * ((A / math.Pow(M, P)) + B)
type curve struct {
	A, B, C, P float64
}

// US curves
var (
	U1 = curve{A: 0.0104, B: 0.2256, C: 1.08, P: 0.02}
	U2 = curve{A: 5.95, B: 0.180, C: 5.95, P: 2.0}
	U3 = curve{A: 3.88, B: 0.0963, C: 3.88, P: 2.0}
	U4 = curve{A: 5.67, B: 0.352, C: 5.67, P: 2.0}
	U5 = curve{A: 0.00342, B: 0.00262, C: 0.323, P: 0.02}
)

// IEC curves
var (
	C1 = curve{A: 0.14, B: 0, C: 13.5, P: 0.02}
	C2 = curve{A: 13.5, B: 0, C: 47.3, P: 2.0}
	C3 = curve{A: 80, B: 0, C: 80, P: 2.0}
	C4 = curve{A: 120, B: 0, C: 120, P: 2.0}
	C5 = curve{A: 0.05, B: 0, C: 4.85, P: 0.04}
)

// Inverse time overcurrent relay
type itoc struct {
	pickup   float64
	timedial float64
	curve    curve
}

// opTime calculates operation time based on secondary current amps
func (r itoc) opTime(amps float64) float64 {
	M := amps / r.pickup
	return r.timedial * ((r.curve.A / math.Pow(M, r.curve.P)) + r.curve.B)
}

// tdCalc calculates timedial for given operation time and secondary amps
func (r itoc) tdCalc(t float64, amps float64) float64 {
	M := amps / r.pickup
	return t / ((r.curve.A / math.Pow(M, r.curve.P)) + r.curve.B)
}

// Instantaneous overcurrent relay
type ioc struct {
	pickup float64
	delay  float64
}

// opTime calculates operation time given amps for IOC
func (r ioc) opTime(amps float64) float64 {
	if amps >= r.pickup {
		return r.delay
	}
	return math.MaxFloat64
}
