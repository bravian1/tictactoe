package main

import "fmt"

var Board [3][3]int
type Point struct {
	x int
	y int
}
func PrintBoard (x int){
	
	pos := x
	fmt.Println("-----------")
	for i := 0; i < 3; i++{
		for j:=0; j < 3; j++{
			Board[i][j]=pos
			pos++
			fmt.Printf("%d | ",Board[i][j])
		}
		fmt.Println()
		fmt.Println("-----------")
	}
}

func movePos(p Point) map[int][]Point {
	x := p.x
	y := p.y

	post := make(map[int][]Point)
	for i := 1; i <= 9; i++ {
		// Append the new Point to the list at post[i]
		post[i] = append(post[i], Point{x, y})
		x++
		y++
	}
	return post
}
func main(){
	PrintBoard(2)
}