package main
//https://www.reddit.com/r/dailyprogrammer/comments/8ewq2e/20180425_challenge_358_intermediate_everyones_a/
import (
	"os"
	"encoding/csv"
	"strconv"
)

type Game struct {
	playerOne string
	playerOneScore int
	playerTwo string
	playerTwoScore int
}

func IsInWinners(winners []string, winner string)  bool {

	for _, value := range winners {
		if value == winner{
			return true
		}
	}
	return false

}

func solution1() int{

	//This variable is going to have 16446 records!!!
	var games []Game
	winners := []string{"Villanova"}

	file, err := os.Open("LDbXGeJn.csv")
	defer file.Close()

	if err == nil{

		reader := csv.NewReader(file)
		cdata, err := reader.ReadAll()

		if err == nil{

			for _, row := range cdata[1:] {
				s1, _ := strconv.Atoi(row[1])
				s2, _ := strconv.Atoi(row[3])
				games = append(games, Game{row[0], s1, row[2], s2})
			}
		}

	}


	for i := 0; i < len(games);  i++{

		for _, game := range games {

			var winner string

			if game.playerOneScore > game.playerTwoScore{
				winner = game.playerOne
			} else if game.playerOneScore <= game.playerTwoScore{
				winner = game.playerTwo
			}

			if !IsInWinners(winners, winner){
				winners = append(winners, winner)
			}

		}
	}
	return len(winners)
}

func main() {

}
