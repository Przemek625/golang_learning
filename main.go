package main

import (
	"sort"
	"fmt"
	"strconv"
	"math/cmplx"
	"math"
)

type Vertex struct {
	X int
	Y int
}

type Player struct {
	nickname string
	points   int
}

func NewPlayer(nickname string) *Player{

	player := new(Player)
	player.nickname = nickname
	player.points = 0
	return player
}

func (player *Player ) addPoint()  {
	player.points +=1
}

func (player *Player) addPoints(points int)  {
	player.points += points
}

func (player *Player ) subtractPoint()  {
	player.points -=1
}

func (player *Player) getPoints() int{
	return player.points
}

func getWinner(pl_one Player, pl_two Player) Player{
	players := []Player{pl_one, pl_two}
	sort.Slice(players, func(i, j int) bool {
		return players[i].points > players[j].points
	})

	return players[0]
}

func main() {

	player1 := NewPlayer("player1")
	player2 := NewPlayer("player2")
	player1.addPoint()
	player1.addPoint()
	player1.addPoint()
	player1.addPoint()
	player2.addPoint()
	winner := getWinner(*player1, *player2)
	fmt.Println(winner.nickname)

	player2.addPoints(100)
	fmt.Println(player2.points)
	fmt.Println(getWinner(*player1, *player2))

	var text string = "1231231"
	//converting string to int
	w, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w)
	//trying to convert not number to int
	d, err := strconv.Atoi("not_number")

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

	//variable declaration with initializers
	var i, j int = 1, 2

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Println(ToBe, MaxInt, z)

	fmt.Println(i, j)
	//short variable declaration
	k := 100

	fmt.Println(k)

	const someConstant = 1234
	fmt.Println(someConstant)

	//for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var x, n, lim float64 = 3, 2, 10

	//Variables declared by the statement are only in scope until the end of the if
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
	} else {
		fmt.Println("v > lim")
	}

	for i := 0; i < 100; i++{
		if i % 5 == 0 && i % 3 == 0{
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	//Pointers
	l := 0
	fmt.Println(l)
	//pointer_l points to l
	pointer_l := &l
	//change the value of l through pointer_l
	*pointer_l = 2
	fmt.Println(l)

	vertex := Vertex{1, 1}
	vertex.X = 100
	pointer_vertex := &vertex
	// This means the same (*pointer_vertex). However language allows below notation.
	pointer_vertex.X = 1000
	fmt.Println(vertex)

}
