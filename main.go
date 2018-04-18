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
	"reflect"
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

func (v Vertex ) say_cords()  {
	fmt.Println(v.X, v.Y)

}

//We use composition in here
type MoreAdvancedVertex struct{
	Vertex
	Z int

}

func (v MoreAdvancedVertex ) say_cords() {
	fmt.Println(v.X, v.Y, v.Z)

}

//Interfaces allow polymorphism
type Cords interface {
	say_cords()
}

func say_cords(c Cords)  {
	c.say_cords()

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

	vex := Vertex{1, 2}
	mav := MoreAdvancedVertex{Vertex{1, 2}, 3}
	mav.say_cords()

	//If say_cords was taking a pointer of MoreAdvancedVertex
	//There would be an error...

	//There we use polymorphism
	say_cords(vex)
	say_cords(mav)

	//Coping unique values
	var set []int
	array := []int{1, 1, 1, 2, 2, 3}

	for _, v := range array{
		in := false
		for _, v2 := range set{
			if v == v2{
				in = true
			}
		}

		if in == false{
			set = append(set, v)
		}
	}
	fmt.Println(array)
	fmt.Println(set)

	//A slice has both a length and a capacity.
	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice.
	//The length and capacity of a slice s can be obtained using the expressions
	// len(s) and cap(s).
	//You can extend a slice's length by re-slicing it, provided it has sufficient
	// capacity. Try changing one of the slice operations in the example program
	//to extend it beyond its capacity and see what happens.

	//For example, make([]int, 0, 10) allocates an underlying array
	//of size 10 and returns a slice of length 0 and capacity 10 that is
	//backed by this underlying array
	xs := make([]int, 0, 10)
	//length of xs is 0
	fmt.Println(len(xs))
	//capacity of xs is 10
	fmt.Println(cap(xs))

	//it will panic
	//xs[3] = 10
	//fmt.Println(xs[0])

	//https://tour.golang.org/moretypes/11
	//Extending slice length
	xs = xs[:10]
	//now the length is 10
	fmt.Println(len(xs))
	//prints [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(xs)
	//panic: runtime error: slice bounds out of range
	//xs = xs[:20]

	// Drop its first two values.
	xs = xs[2:]
	fmt.Println(xs)


	point := Vertex{1, 1}
	//Operand & generates a pointer to Vertex
	point2 := &point

	//https://tour.golang.org/moretypes/1

	//main.Vertex
	fmt.Println(reflect.TypeOf(point))
	//*main.Vertex
	fmt.Println(reflect.TypeOf(point2))
	//main.Vertex
	//The * operator denotes the pointer's underlying value.
	fmt.Println(reflect.TypeOf(*point2))

	//This is invalid
	//fmt.Println(reflect.TypeOf(*

	//Reading the value through pointer
	fmt.Println((*point2).X)
	//To access the field X of a struct when we have the struct pointer p we could
	//write (*p).X. However, that notation is cumbersome,
	//so the language permits us instead to write just p.X, without the explicit dereference.
	fmt.Println(point2.X)

	var pp *int
	ii := 42
	pp = &ii

	fmt.Println(*pp) // read i through the pointer p
	*pp = 21         // set i through the pointer p

	abc := "AbcdEfg"

	for i := 0; i < len(abc);  i++{
		char := string(abc[i])
		if char != strings.ToUpper(char){
			fmt.Println(char, "is not upper.")
		} else {
			fmt.Println(char, "is upper.")
		}
	}

	abcd := "QQ@#$$@$@#^$%*&^)()123abcde"

	//QQ@#$$@$@#^$%*&^)()123ABCDE
	fmt.Println(strings.ToUpper(abcd))

	//Go allows not giving any values during creating instances
	emty_vertex := Vertex{}
	//{0 0}
	fmt.Println(emty_vertex)

}
