package main
import ("fmt"; "math")

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) distance (x1, y1, x2, y2 float64) (float64, float64) {
  a := r.x2 - r.x1
  b := r.y2 - r.y1
  
  return math.Sqrt(a*a ), math.Sqrt (b*b)
  //return math.Sqrt (b*b)
}

func (r *Rectangle) area() (float64, float64) {
	t, l := r.distance(r.x1, r.y1, r.x1, r.y2)
	m, w := r.distance(r.x1, r.y1, r.x2, r.y1)
	fmt.Println("value1:",t)
	fmt.Println("value2:",m)
	return l * w, t * m
}

func main(){
	r := Rectangle{0, 0, 10, 10}
    fmt.Println(r.area())


}