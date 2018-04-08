package main

import (
	"sort"
	"fmt"
	"strconv"
	"math/cmplx"
	"math"
// Packages for playing with the database
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

type Book struct {
	Id int `gorm:"PRIMARY_KEY; AUTO_INCREMENT"`
	Name string `gorm:"NOT NULL"`
	Author string `gorm:"NOT NULL"`
}

type Vertex struct {
	//This variable will be visible outside the package
	X int
	Y int
}

type Player struct {
	//This variable won't be visible outside the package
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
	//has type *Vertex
	pointer_vertex := &vertex
	// This means the same (*pointer_vertex). However language allows below notation.
	pointer_vertex.X = 1000
	fmt.Println(vertex)

	//[n]T is an array of n values of type T
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//[]T is a slice with elements of type T
	//creates a slice which includes elements 0 through 3 of primes
	//Slices are like references to arrays. They dose not store any data
	//It just describes a section of an underlying array.
	var s []int = primes[0:4]
	fmt.Println(s)
	s[0] = 10000000
	//If we change value via slice, We change a value of primes.
	//Proof: prints [10000000 3 5 7 11 13]
	fmt.Println(primes)
	//Slice literal
	//A slice literal is like an array literal without the length.
	//Array literal [3]bool{true, true, false}
	//And this creates the same array as above, then builds a slice that references it:
	//[]bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	qq := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(qq)

	joinedStrings := strings.Join([]string{"1", "2", "3", "4"}, "_")
	fmt.Println(joinedStrings)

	//converting int to string
	integer := strconv.Itoa(1234)
	fmt.Println(integer)
	//converting string to int
	numberFromString, err := strconv.Atoi("1234")
	fmt.Println(numberFromString)

	//Playing with the database
	db, err := gorm.Open("mysql", "falcon:falcon@/falcon")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	for i := 0; i < 10;  i++{
		iToString := strconv.Itoa(i)
		book := Book{
			Name: strings.Join([]string{"Book", iToString}, "_"),
			Author: strings.Join([]string{"Author", iToString}, "_"),
			}
		db.Create(&book)
	}

	var book []Book

	db.Find(&book)

	for key, value := range book {
		fmt.Println(key, value)
	}

	var potop Book
	var potop2 Book

	//// SELECT * FROM books WHERE name = "Potop";
	db.Find(&potop, "name = ?", "Potop")
	db.Where("name = ?", "Potop").Find(&potop2)

	fmt.Println(potop)
	fmt.Println(potop2)

	map_people_age := map[string]int{
		"Bob": 24,
		"Susan": 16,
		"Denis": 60,
	}

	for key, value := range map_people_age {
		fmt.Println(key, value)
	}

	fmt.Println(map_people_age["Bob"])
	fmt.Println(map_people_age["NO_KEY"])

	if val, ok := map_people_age["NO_KEY"]; ok {
		fmt.Println("NO_KEY is in the map with value: ", val)
	} else {
		fmt.Println("there is no such key")
	}

}
