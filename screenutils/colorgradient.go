package screenutils

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/tpfeiffer67/console/screen"
)

func ColorToColorful(color screen.Color) colorful.Color {
	fr, fg, fb, _ := color.GetRGBfloat64()
	return colorful.Color{fr, fg, fb}
}

func BuildColorGradient(startColor, endColor screen.Color, steps int) []screen.Color {
	var gradients []screen.Color
	startcolorfull := ColorToColorful(startColor)
	endcolorfull := ColorToColorful(endColor)

	for step := 0; step < steps; step++ {
		f := float64(step) / float64(steps-1)
		gradColor := startcolorfull.BlendHcl(endcolorfull, f).Clamped()
		//r, g, b := gradColor.RGB255()
		gradients = append(gradients, screen.NewTrueColorRGB(gradColor.RGB255()))
	}
	return gradients
}

func BuildColorGradientMatrix(topLeft, topRight, bottomLeft, bottomRight screen.Color, vertSteps, horizSteps int) [][]screen.Color {
	var matrixGrad [][]screen.Color

	leftGradients := BuildColorGradient(topLeft, bottomLeft, vertSteps)
	rightGradients := BuildColorGradient(topRight, bottomRight, vertSteps)

	for i := 0; i < vertSteps; i++ {
		horizGradients := BuildColorGradient(leftGradients[i], rightGradients[i], horizSteps)
		matrixGrad = append(matrixGrad, horizGradients)
	}

	return matrixGrad
}
