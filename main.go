package main
import (
	"fmt"
	"bufio"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex")
		inputString := scanner.Scan().Text()
		firstWord := cleanInput(inputString)[0]
	}
}

