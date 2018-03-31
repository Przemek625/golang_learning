package main

import (
	"sort"
	"fmt"
)

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

}
