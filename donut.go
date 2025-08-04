package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	width, height := 80, 24
	thetaSpacing := 0.07
	phiSpacing := 0.02
	chars := ".,-~:;=!*#$@"

	A, B := 0.0, 0.0

	for {
		z := make([]float64, width*height)
		b := make([]rune, width*height)
		for i := range b {
			b[i] = ' '
		}

		for j := 0.0; j < 2*math.Pi; j += thetaSpacing {
			for i := 0.0; i < 2*math.Pi; i += phiSpacing {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e
				x := int(float64(width)/2 + 30*D*(l*h*m-t*n))
				y := int(float64(height)/2 + 15*D*(l*h*n+t*m))
				o := x + width*y
				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))
				if y >= 0 && y < height && x >= 0 && x < width && D > z[o] {
					z[o] = D
					if N > 0 {
						if N > len(chars)-1 {
							N = len(chars) - 1
						}
						b[o] = rune(chars[N])
					} else {
						b[o] = rune(chars[0])
					}
				}
			}
		}

		clear()
		for k := 0; k < len(b); k++ {
			if k%width == 0 {
				fmt.Print("\n")
			}
			fmt.Printf("%c", b[k])
		}
		A += 0.04
		B += 0.08
		time.Sleep(30 * time.Millisecond)
	}
}
