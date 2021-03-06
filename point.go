package point

import (
	"fmt"
	"math"
)

type Point struct {
	X      float64
	Y      float64
	Marked bool
}

type IntPoint struct {
	X int64
	Y int64
}

func (this *IntPoint) String() string {
	return fmt.Sprintf("{X:%d,Y:%d}", this.X, this.Y)
}
func pointDistance(a, b *Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}
func Bulge2Arc(startP, endP *Point, bulge float64) []*Point {
	if bulge == 0 {
		return []*Point{&Point{X: startP.X, Y: startP.Y}, &Point{X: endP.X, Y: endP.Y}}
	}
	c := (1/bulge - bulge) / 2
	//# Calculate the centre point (Micke's formula!)
	O := &Point{
		X: (startP.X + endP.X - (endP.Y-startP.Y)*c) / 2,
		Y: (startP.Y + endP.Y + (endP.X-startP.X)*c) / 2,
	}
	r := pointDistance(startP, O)
	startDegree := math.Atan2(startP.Y-O.Y, startP.X-O.X) * 180 / math.Pi
	endDegree := math.Atan2(endP.Y-O.Y, endP.X-O.X) * 180 / math.Pi

	// for endDegree < 0 {
	// 	endDegree += 360
	// }
	// for startDegree < 0 {
	// 	startDegree += 360
	// }

	// log.Println("X:", O.X)
	// log.Println("Y:", O.Y)
	// log.Println("r:", r)
	// log.Println("bulge:", bulge)

	// log.Println("startDegree:", startDegree)
	// log.Println("endDegree:", endDegree)
	// log.Println("arc:", endDegree-startDegree)

	if bulge > 0 {
		for endDegree <= startDegree {
			endDegree += 360
		}
		return NewArc(O.X, O.Y, r, startDegree, endDegree)
	} else {

		for startDegree <= endDegree {
			startDegree += 360
		}

		poly := NewArc(O.X, O.Y, r, endDegree, startDegree)
		reverse(poly)
		return poly
	}

}
func NewArc(x, y float64, r float64, startDegree, endDegree float64) []*Point {
	if endDegree <= startDegree {
		panic("")
	}
	poly := []*Point{}
	var num = int(math.Ceil((2 * math.Pi) / math.Acos(1-(2/r))))
	num = int(float64(num)*(endDegree-startDegree)/360 + 1)
	if num < 10 {
		num = 10
	}
	alltheta := (endDegree - startDegree) / 360 * (2 * math.Pi)
	starttheta := startDegree / 360 * (2 * math.Pi)
	for i := 0; i <= num; i++ {
		theta := float64(i)*(alltheta/float64(num)) + starttheta
		point := &Point{
			X: r*math.Cos(theta) + x,
			Y: r*math.Sin(theta) + y,
		}
		poly = append(poly, point)
	}
	return poly
}
func reverse(s []*Point) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func NewCircle(x, y float64, r float64) []*Point {
	poly := []*Point{}
	var num = int(math.Ceil((2 * math.Pi) / math.Acos(1-(2/r))))
	if num < 30 {
		num = 30
	}
	for i := 0; i < num; i++ {
		theta := float64(i) * ((2 * math.Pi) / float64(num))
		point := &Point{
			X: r*math.Cos(theta) + x,
			Y: r*math.Sin(theta) + y,
		}
		poly = append(poly, point)
	}
	return poly
}
