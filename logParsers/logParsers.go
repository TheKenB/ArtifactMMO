package logParsers

import (
	response "artifactMMO/responseType"
	"encoding/json"
	"fmt"
)

func ParseOperator(val string, rawLog []byte, character string) {
	switch character {
	case "None":
		NoCharacterParseOperator(val, rawLog)
	default:
		CharacterParseOperator(val, rawLog)
	}
}

func CharacterParseOperator(val string, rawLog []byte) {
	switch val {
	case "0":
		fmt.Println("KeepMoving")
	case "1":
		ParseCombatLog(rawLog)
	case "2":
		ParseCooldown(rawLog)
	case "3":
		ParseUnequip(rawLog)
	default:
		fmt.Println("Some Action Taken")
	}
}

func NoCharacterParseOperator(val string, rawLog []byte) {
	switch val {
	case "0":
		ParseAllMaps(rawLog)
	default:
		fmt.Println("NoCharacterParse log")
	}
}

func ParseCombatLog(rawLog []byte) {
	var result response.CombatResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Character.Name == "" {
		fmt.Println("Can't Attack")
	}
	fmt.Println(result.Data.Character.Name)
	for _, rec := range result.Data.Fight.Logs {
		fmt.Println(rec)
	}
}

func ParseAllMaps(rawLog []byte) {
	var result response.AllMapsResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if len(result.Data) == 0 {
		fmt.Println("Can't Get All Map")
	}
	for _, rec := range result.Data {
		fmt.Println(rec)
	}
}

func ParseCooldown(rawLog []byte) {
	var result response.GetMyCharactersResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if len(result.Data) == 0 {
		fmt.Println("Can't Get Characters")
	}
	for _, rec := range result.Data {
		fmt.Println("Cooldown: ", rec.Cooldown)
		fmt.Println("Cooldown Exp: ", rec.Cooldown_expiration)
	}
}

func ParseUnequip(rawLog []byte) {
	var result response.EquipingResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println("Unequiped")
}
