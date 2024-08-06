package menuMaps

import (
	"fmt"
	"strings"
)

const (
	Slot0  string = "Weapon"
	Slot1  string = "Shield"
	Slot2  string = "Helmet"
	Slot3  string = "Body_armor"
	Slot4  string = "Leg_armor"
	Slot5  string = "Boots"
	Slot6  string = "Ring1"
	Slot7  string = "Ring2"
	Slot8  string = "Amulet"
	Slot9  string = "Artifact1"
	Slot10 string = "Artifact2"
	Slot11 string = "Artifact3"
	Slot12 string = "Consumable1"
	Slot13 string = "Consumable2"
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
[3] Unequip
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
		return GetCharacters(char)
	case "3":
		return UnequipItem(char)
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

func GetCharacters(char string) ActionResponse {
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

func UnequipItem(char string) ActionResponse {
	slot := ItemSlot()
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/unequip"
	data := `{"slot": "` + slot + `" }`
	return ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
}

func ItemSlot() string {
	slot := ""
	fmt.Println(`
---------------------
[0]		` + Slot0 + `
[1]		` + Slot1 + `
[2] 	` + Slot2 + `
[3] 	` + Slot3 + `
[4] 	` + Slot4 + `
[5] 	` + Slot5 + `
[6] 	` + Slot6 + `
[7] 	` + Slot7 + `
[8] 	` + Slot8 + `
[9] 	` + Slot9 + `
[10]	` + Slot10 + `
[11] 	` + Slot11 + `
[12] 	` + Slot12 + `
[13] 	` + Slot13 + `
---------------------`)
	_, err := fmt.Scan(&slot)
	if err != nil {
		fmt.Println("Item Slot Error:", err)
	}

	switch slot {
	case "0":
		return strings.ToLower(Slot0)
	case "1":
		return strings.ToLower(Slot1)
	case "2":
		return strings.ToLower(Slot2)
	case "3":
		return strings.ToLower(Slot3)
	case "4":
		return strings.ToLower(Slot4)
	case "5":
		return strings.ToLower(Slot5)
	case "6":
		return strings.ToLower(Slot6)
	case "7":
		return strings.ToLower(Slot7)
	case "8":
		return strings.ToLower(Slot8)
	case "9":
		return strings.ToLower(Slot9)
	case "10":
		return strings.ToLower(Slot10)
	case "11":
		return strings.ToLower(Slot11)
	case "12":
		return strings.ToLower(Slot12)
	case "13":
		return strings.ToLower(Slot13)
	default:
		return ""
	}
}
