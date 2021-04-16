package linearAlgebra

func ZeroVectors(input ...int) interface{} {

	if len(input) == 1 {

		vs1D := []float64{}

		for i := 0; i < input[0]; i++ {
			vs1D = append(vs1D, 0)
		}
		return vs1D

	} else if len(input) == 2 {
		vs2D := make([][]float64, input[0])
		for i := 0; i < input[0]; i++ {
			vs2D[i] = make([]float64, input[1])
			for j := 0; j < input[1]; j++ {
				vs2D[i][j] = 0
			}
		}
		return vs2D

	} else if len(input) == 3 {
		vs3D := make([][][]float64, input[0])
		for i := 0; i < input[0]; i++ {
			vs3D[i] = make([][]float64, input[1])
			for j := 0; j < input[1]; j++ {
				vs3D[i][j] = make([]float64, input[2])
				for k := 0; k < input[2]; k++ {
					vs3D[i][j][k] = 0
				}
			}
		}

		return vs3D
	}
	return nil
}
