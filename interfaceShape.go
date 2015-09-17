package main

import ("fmt"; "math")

type Shape interface{

  calculatePerimeter() float64
}

/*
func totalPerimeter(shapes ...Shape) float64{

var perimeter float64

	for _, s := range shapes {

	   perimeter += s.calculatePerimeter()
	}
	return perimeter
}
*/

type Circle struct{

	radius float64
}

func (c Circle) calculatePerimeter() float64{

	return 2 * math.Pi * c.radius
}

type Rectangle struct{

	length,breadth float64
}

func (r Rectangle) calculatePerimeter() float64{

	return 2 * r.length + 2 * r.breadth
}

func displayPerimeter(shapes Shape){

	fmt.Println(shapes.calculatePerimeter())


}

func main(){

	c := Circle{7}
	r := Rectangle{6,6}

	fmt.Print("Perimeter of a circle with radius 7 is - ")
        displayPerimeter(c)
	//fmt.Println("")

	fmt.Print("Perimeter of a rectangle with length 5 and breadth 6 is - ")
	displayPerimeter(r)
	//fmt.Println("")


}
