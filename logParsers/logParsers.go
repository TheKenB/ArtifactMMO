package logParsers

import (
	response "artifactMMO/responseType"
	"encoding/json"
	"fmt"
	"strconv"
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
	case "4":
		ParseGather(rawLog)
	case "5":
		ParseCraft(rawLog)
	case "6":
		ParseDepositBank(rawLog)
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

func ParseGather(rawLog []byte) {
	var result response.GatherResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Gather")
	}

	fmt.Println(result.Data.Character.Name + " Gathering")
	fmt.Println("Xp: " + strconv.Itoa(result.Data.Details.Xp))
	// Looks like nothing is returned currently
	for _, rec := range result.Data.Details.Items {
		fmt.Println("Resource: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
	}
}

func ParseCraft(rawLog []byte) {
	var result response.GatherResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Craft")
	}

	fmt.Println(result.Data.Character.Name + " Crafting")
	fmt.Println("Xp: " + strconv.Itoa(result.Data.Details.Xp))
	// Looks like nothing is returned currently
	for _, rec := range result.Data.Details.Items {
		fmt.Println("Item: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
	}
}

func ParseDepositBank(rawLog []byte) {
	var result response.BankDepositResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Bank Deposit")
	}

	fmt.Println(result.Data.Character.Name + " Depositing")
	for _, rec := range result.Data.Bank {
		fmt.Println("Item: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
	}
}
