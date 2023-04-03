package convertt

func convertTemperature(celsius float64) []float64 {
	k:= 273.15
    mass := [2]float64{}
    mass[0] = k+celsius
    mass[1] = celsius*1.80+32.00

    return mass[:]
}