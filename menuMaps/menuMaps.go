package menuMaps

func GetCharacter(val string) string {
	// return Url for relevant api
	switch val {
	// Character
	case "0":
		return "Guile"
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
func GetAction(val string, char string) string {
	switch val {
	case "0":
		return ""
	case "1":
		return Attack(char)
	default:
		return ""
	}
}

func Attack(char string) string {
	return "https://api.artifactsmmo.com/my/DaGuile/action/fight"
}
