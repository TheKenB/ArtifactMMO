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

func GetCharacter(val string) string {
	// return Url for relevant api
	switch val {
	// Character
	case "0":
		return "DaGuile"
	case "1":
		return "PlaceHolder"
	default:
		return ""
	}
}

func GetCharacterMenu() string {
	return `
---------------------
[0] Guile
[1] PlaceHolder
---------------------`
}

func GetActionsMenu() string {
	return `
---------------------
[0] Walk
[1] Attack
---------------------`
}

// Call the function to build your url
func GetAction(val string, char string) ActionResponse {
	switch val {
	case "0":
		return Walk(char)
	case "1":
		return Attack(char)
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
	data := `\n  \"x\":` + x + `\n  \"y\":` + y + `\n`
	return ActionResponse{Url: "https://api.artifactsmmo.com/my/" + char + "/action/move", PostGet: "POST", Data: data}
}
