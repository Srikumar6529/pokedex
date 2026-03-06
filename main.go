package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		inputString := scanner.Text()
		firstWord := cleanInput(inputString)[0]
		fmt.Println("Your command was:",firstWord)
	}
}

