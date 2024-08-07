package menuMaps

import (
	api "artifactMMO/api"
	"fmt"
)

func CharacterOperator(val string) string {
	// return Url for relevant api
	switch val {
	// Character
	case "0":
		return "None"
	case "1":
		return "DaGuile"
	default:
		return ""
	}
}

func GetCharactersMenu() string {
	return `
---------------------
[0] None
[1] DaGuile
---------------------`
}

func CharacterActionMenu() string {
	return `
---------------------
[0] Walk
[1] Attack
[2] Cooldown
[3] Unequip
[4] Gather
[5] Craft
---------------------`
}

func NoCharacterActionMenu() string {
	return `
---------------------
[0] Locations
---------------------`
}

// Get list of action menus based on character condition
func ActionMenuOperator(character string) string {
	switch character {
	case "None":
		return NoCharacterActionMenu()
	default:
		return CharacterActionMenu()
	}
}

func ActionOperator(val string, char string) {
	switch char {
	case "None":
		GetNoCharacterAction(val, char)
	default:
		GetCharacterAction(val, char)
	}
}

// Call the function to build your url
func GetCharacterAction(val string, char string) {
	switch val {
	case "0":
		api.Move(char, val)
	case "1":
		api.Fight(char, val)
	case "2":
		api.GetMyCharacters(char, val)
	case "3":
		api.UnequipItem(char, val)
	case "4":
		api.Gathering(char, val)
	case "5":
		api.Crafting(char, val)
	default:
		fmt.Println("No Action")
	}
}

func GetNoCharacterAction(val string, char string) {
	switch val {
	case "0":
		api.GetAllMaps(val)
	default:
		fmt.Println("No Actin")
	}
}
