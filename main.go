package main

import (
	"errors"
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/codenjoyme/codenjoy-go-client/games/clifford"
	"github.com/codenjoyme/codenjoy-go-client/games/mollymage"
	"github.com/codenjoyme/codenjoy-go-client/games/sample"
	"github.com/codenjoyme/codenjoy-go-client/games/verland"
	"log"
	"os"
	"fmt"
)

func main() {
	game := "mollymage"
	url := "http://127.0.0.1:8080/codenjoy-contest/board/player/0?code=000000000000"
    source := "Runner"

	if len(os.Args) == 3 {
		game = os.Args[1]
		url = os.Args[2]
        source = "Environment"
	}

	fmt.Printf("Got from %s:\n" +
            "\t 'GAME': '%s'\n" +
            "\t 'URL':  '%s'\n",
            source,
            game,
            url)
  
	s, err := gameSolver(game)
	if err != nil {
		log.Fatalln(err)
	}
  engine.NewWebSocketRunner(url).Run(s)
}

func gameSolver(game string) (engine.Solver, error) {
	switch game {
	case "sample":
		return sample.NewSolver()
	case "mollymage":
		return mollymage.NewSolver()
	case "clifford":
		return clifford.NewSolver()
	case "verland":
		return verland.NewSolver()
	default:
		return nil, errors.New("unable to determine game type")
	}
}
