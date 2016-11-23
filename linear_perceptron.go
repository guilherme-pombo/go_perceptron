package perceptron

import (
	"fmt"
	"math/rand"
)

type LinearPerceptron struct {
	weights   []float64
	activation float64
	learningRate float64
}

// Create Linear Perceptron
func New(dims int, zeroInit bool, learningRate float64) *LinearPerceptron {
	perceptron := new(LinearPerceptron)

	//Either initialize with zero weights or random values near zero
	if zeroInit {
		tmp_weights := make([]float64, numInputs)
		for i := 0; i < numInputs; i++ {
			tmp_weights[i] = 0
		}
	}
	else {
		tmp_weights := make([]float64, numInputs)
		for i := 0; i < numInputs; i++ {
			tmp_weights[i] = rand.Float64()
		}
	}

	perceptron.weights = tmp_weights
	perceptron.activation = rand.Float64()
	perceptron.learningRate = learningRate
	return perceptron
}


// Do the actual training iterations
// TODO: add learning rate
// TODO: use proper activation function
// TODO: add support for real valued inputs
func (perceptron *Perceptron) trainPerceptron(input []int, label int) {
	if len(p.Weights) != len(input) {
		panic(fmt.Sprintf("Dimension mismatch with input dims and input layer")))
	}

	var sum float64
	for i := 0; i < len(p.Weights); i++ {
		sum += p.Weights[i] * float64(input[i])
	}

	var prediction float64
	if sum >= perceptron.activation {
		prediction = 1
	} else {
		prediction = 0
	}

	perceptron.activation = perceptron.activation - (float64(label) - float64(prediction))

	// Update weights
	for i := 0; i < len(p.Weights); i++ {
		perceptron.weights[i] = perceptron.weights[i] + (float64(label)-float64(prediction))*float64(input[i])
	}

}
