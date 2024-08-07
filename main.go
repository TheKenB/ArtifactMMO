package main

import (
	menus "artifactMMO/menuMaps"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	Run()
}

func Run() {
	var running int = 1
	var character string = ""

	for running == 1 {

		// Print Character list
		fmt.Println(menus.GetCharactersMenu())
		character = HandleCharacterStep()
		keyPressed := HandleActionChoice(character)
		menus.ActionOperator(keyPressed, character)
	}
}

func HandleCharacterStep() string {
	// Call the reader to read user's input
	reader := bufio.NewReader(os.Stdin)
	key, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if len(strings.TrimSpace(key)) > 0 {
		key = strings.TrimSpace(key)
		char := menus.CharacterOperator(key)
		return char
	} else {
		return HandleCharacterStep()
	}
}

func HandleActionChoice(char string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(menus.ActionMenuOperator(char))
	key2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(key2)
}
