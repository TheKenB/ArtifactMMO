package menuMaps

import (
	"fmt"
)

type ActionResponse struct {
	Url     string
	PostGet string
	Data    string
}

func NewEmptyActionResponse() ActionResponse {
	return ActionResponse{Url: "", PostGet: "", Data: ""}
}

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

// Get list of action menus based on character condition
func ActionMenuOperator(character string) string {
	switch character {
	case "None":
		return NoCharacterActionMenu()
	default:
		return CharacterActionMenu()
	}
}

func CharacterActionMenu() string {
	return `
---------------------
[0] Walk
[1] Attack
[2] Cooldown
---------------------`
}

func NoCharacterActionMenu() string {
	return `
---------------------
[0] Locations
---------------------`
}

func ActionOperator(val string, char string) ActionResponse {
	switch char {
	case "None":
		return GetNoCharacterAction(val, char)
	default:
		return GetCharacterAction(val, char)
	}
}

// Call the function to build your url
func GetCharacterAction(val string, char string) ActionResponse {
	switch val {
	case "0":
		return Walk(char)
	case "1":
		return Attack(char)
	case "2":
		return Cooldown(char)
	default:
		return NewEmptyActionResponse()
	}
}

func GetNoCharacterAction(val string, char string) ActionResponse {
	switch val {
	case "0":
		return GetAllMaps()
	default:
		return NewEmptyActionResponse()
	}
}

func Attack(char string) ActionResponse {
	return ActionResponse{Url: "https://api.artifactsmmo.com/my/" + char + "/action/fight", PostGet: "POST", Data: ""}
}

func Walk(char string) ActionResponse {
	var x string
	var y string

	fmt.Println("Enter x y location")
	_, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println("X,Y Error:", err)
	}
	data := `{"x":` + x + `,"y":` + y + `}`
	return ActionResponse{Url: "https://api.artifactsmmo.com/my/" + char + "/action/move", PostGet: "POST", Data: data}
}

func Cooldown(char string) ActionResponse {
	return ActionResponse{Url: "https://api.artifactsmmo.com/my/characters", PostGet: "GET", Data: ""}
}

func GetAllMaps() ActionResponse {
	var page string
	fmt.Println("Enter page number")
	_, err := fmt.Scan(&page)
	if err != nil {
		fmt.Println("X,Y Error:", err)
	}
	return ActionResponse{Url: "https://api.artifactsmmo.com/maps/?page=" + page + "&size=100", PostGet: "GET", Data: ""}
}
