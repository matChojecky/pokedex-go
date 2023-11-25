package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeJourney "github.com/matChojecky/pokedexcli/journey"
)


type Config struct {
	PokeJourney *pokeJourney.PokeJourney
}

func Start(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Print("pokedex > ")
		scanner.Scan() 

		cmd := cleanInput(scanner.Text())

		if len(cmd) == 0 {
			continue
		}
		
		args := cmd[1:]
		
		command, exists := getCommands()[cmd[0]];

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		fmt.Println()
		err := command.callback(cfg, args...);
		fmt.Println()
		
		if err != nil {
			fmt.Println(err)
			fmt.Println()
			continue
		}
	}
	
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}