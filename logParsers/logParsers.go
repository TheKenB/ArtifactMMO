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
		fmt.Println("Moved")
	case "1":
		ParseCombatLog(rawLog)
	case "3":
		ParseUnequip(rawLog)
	case "4":
		ParseGather(rawLog)
	case "5":
		ParseCraft(rawLog)
	case "6":
		ParseDepositBank(rawLog)
	case "7":
		ParseDepositBankGold(rawLog)
	case "8":
		ParseRecyling(rawLog)
	case "9":
		ParseWithdrawBank(rawLog)
	case "10":
		ParseWithdrawBankGold(rawLog)
	case "11":
		ParseGeBuyItem(rawLog)
	case "12":
		ParseGeSellItem(rawLog)
	case "13":
		ParseAcceptNewTask(rawLog)
	case "14":
		ParseCompleteTask(rawLog)
	case "15":
		ParseExchangeTask(rawLog)
	default:
		fmt.Println("Some Action Taken")
	}
}

func NoCharacterParseOperator(val string, rawLog []byte) {
	switch val {
	case "0":
		ParseAllMaps(rawLog)
	case "1":
		ParseAllCharacterLog(rawLog)
	case "2":
		ParseMyCharacters(rawLog)
	case "3":
		ParseGetBankItems(rawLog)
	case "4":
		ParseGetBankGold(rawLog)
	case "5":
		ParseGetCharacter(rawLog)
	case "6":
		ParseGetMap(rawLog)
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
	} else {
		fmt.Println(result.Data.Character.Name)
		for _, rec := range result.Data.Fight.Logs {
			fmt.Println(rec)
		}
	}
}

func ParseAllMaps(rawLog []byte) {
	var result response.AllMapsResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if len(result.Data) == 0 {
		fmt.Println("Can't Get All Map")
	} else {
		for _, rec := range result.Data {
			fmt.Println(rec)
		}
	}
}

func ParseMyCharacters(rawLog []byte) {
	var result response.GetMyCharactersResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if len(result.Data) == 0 {
		fmt.Println("Can't Get Characters")
	} else {
		for _, rec := range result.Data {
			fmt.Println("Name: ", rec.Name)
			fmt.Println("Cooldown: ", rec.Cooldown)
			fmt.Println("Cooldown Exp: ", rec.Cooldown_expiration)
		}
	}
}

func ParseUnequip(rawLog []byte) {
	var result response.EquipingResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	} else {
		fmt.Println("Unequiped")
	}
}

func ParseGather(rawLog []byte) {
	var result response.GatherResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Gather")
	} else {
		fmt.Println(result.Data.Character.Name + " Gathering")
		fmt.Println("Xp: " + strconv.Itoa(result.Data.Details.Xp))
		// Looks like nothing is returned currently
		for _, rec := range result.Data.Details.Items {
			fmt.Println("Resource: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
		}
	}
}

func ParseCraft(rawLog []byte) {
	var result response.GatherResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Craft")
	} else {
		fmt.Println(result.Data.Character.Name + " Crafting")
		fmt.Println("Xp: " + strconv.Itoa(result.Data.Details.Xp))
		// Looks like nothing is returned currently
		for _, rec := range result.Data.Details.Items {
			fmt.Println("Item: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
		}
	}
}

func ParseDepositBank(rawLog []byte) {
	var result response.BankDepositResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Bank Deposit")
	} else {
		fmt.Println(result.Data.Character.Name + " Depositing")
		for _, rec := range result.Data.Bank {
			fmt.Println("Item: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
		}
	}
}

func ParseDepositBankGold(rawLog []byte) {
	var result response.BankDepositGoldResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Bank Deposit")
	} else {
		fmt.Println(result.Data.Character.Name + " Depositing")
		fmt.Println("Gold: " + strconv.Itoa(result.Data.Bank.Quantity))
	}
}

func ParseRecyling(rawLog []byte) {
	var result response.RecyclingResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Recycle")
	} else {
		fmt.Println(result.Data.Character.Name + " Depositing")
		for _, rec := range result.Data.Details.Items {
			fmt.Println("Item: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
		}
	}
}

func ParseWithdrawBank(rawLog []byte) {
	var result response.BankWithdrawResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Bank Deposit")
	} else {
		fmt.Println(result.Data.Character.Name + " Withdrawing")
		for _, rec := range result.Data.Bank {
			fmt.Println("Item: " + rec.Code + " Quantity: " + strconv.Itoa(rec.Quantity))
		}
	}
}

func ParseWithdrawBankGold(rawLog []byte) {
	var result response.BankWithdrawGoldResponse
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Get Bank Gold Deposit")
	} else {
		fmt.Println(result.Data.Character.Name + " Withdrawing Gold")
		fmt.Println("Gold: " + strconv.Itoa(result.Data.Bank.Quantity))
	}
}

func ParseGeBuyItem(rawLog []byte) {
	var result response.GeBuyItem
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Buy Item")
	} else {
		fmt.Println(result.Data.Character.Name + " Buying Item")
		fmt.Println("Item: " + result.Data.Transaction.Code + " Quantity: " + strconv.Itoa(result.Data.Transaction.Quantity))
	}
}

func ParseGeSellItem(rawLog []byte) {
	var result response.GeSellItem
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Sell Item")
	} else {
		fmt.Println(result.Data.Character.Name + " Selling Item")
		fmt.Println("Item: " + result.Data.Transaction.Code + " Quantity: " + strconv.Itoa(result.Data.Transaction.Quantity))
	}
}

func ParseAcceptNewTask(rawLog []byte) {
	var result response.AcceptNewTask
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Accept Task")
	} else {
		fmt.Println(result.Data.Character.Name + " Accepted Task")
		fmt.Println("Task: " + result.Data.Task.Code + " Type: " + result.Data.Task.Type)
	}
}

func ParseCompleteTask(rawLog []byte) {
	var result response.CompleteTask
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Accept Task")
	} else {
		fmt.Println(result.Data.Character.Name + " Accepted Task")
		fmt.Println("Reward: " + result.Data.Reward.Code + " Quantity: " + strconv.Itoa(result.Data.Reward.Quantity))
	}
}

func ParseExchangeTask(rawLog []byte) {
	var result response.ExchangeTask
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Exchange Task")
	} else {
		fmt.Println(result.Data.Character.Name + " Exchanged Task")
		fmt.Println("Reward: " + result.Data.Reward.Code + " Quantity: " + strconv.Itoa(result.Data.Reward.Quantity))
	}
}

func ParseDeleteItem(rawLog []byte) {
	var result response.DeleteItem
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if result.Data.Cooldown.Reason == "" {
		fmt.Println("Can't Delete Item")
	} else {
		fmt.Println(result.Data.Character.Name + " Deleted Item")
		fmt.Println("Reward: " + result.Data.Item.Code + " Quantity: " + strconv.Itoa(result.Data.Item.Quantity))
	}
}

func ParseAllCharacterLog(rawLog []byte) {
	var result response.AllCharacterLogs
	if err := json.Unmarshal(rawLog, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if len(result.Data) == 0 {
		fmt.Println("Can't Get Logs")
	} else {
		fmt.Println(" -- Logs --")
		fmt.Println("Total: " + strconv.Itoa(result.Total) + " Page: " + strconv.Itoa(result.Page) + " Size: " + strconv.Itoa(result.Size) + " Pages: " + strconv.Itoa(result.Pages))
		for _, rec := range result.Data {
			fmt.Println("")
			fmt.Println("/////////////////////////////////////")
			fmt.Println("Character: " + rec.Character + " Account: " + rec.Account + " Type: " + rec.Type + " Description: " + rec.Description)
			fmt.Println("Content: " + rec.Content + " Cooldown: " + strconv.Itoa(rec.Cooldown) + " Cooldown Exp: " + rec.CooldownExpiration.String() + " Created At: " + rec.CreatedAt.String())
		}
	}
}

func ParseGetBankItems(rawLog []byte) {
	var result response.GetBankItems
	if err := json.Unmarshal(rawLog, &result); err != nil {
		fmt.Println("Can't unmarshal JSON")
	}
	if len(result.Data) == 0 {
		fmt.Println("Can't get bank items")
	} else {
		for _, rec := range result.Data {
			fmt.Println("Item: " + rec.Code + " Quantity:" + strconv.Itoa(rec.Quantity))
		}
	}
}

func ParseGetBankGold(rawLog []byte) {
	var result response.GetBankGold
	if err := json.Unmarshal(rawLog, &result); err != nil {
		fmt.Println("Can't unmarshal JSON")
	}
	if result.Data.Quantity == "0" {
		fmt.Println("Can't get gold from bank")
	} else {
		fmt.Println("Withdrawed: " + result.Data.Quantity)
	}
}

func ParseGetCharacter(rawLog []byte) {
	var result response.GetCharacter
	if err := json.Unmarshal(rawLog, &result); err != nil {
		fmt.Println("Can't unmarshal JSON")
	} else {
		fmt.Println("Name: " + result.Data.Name + " Level: " + strconv.Itoa(result.Data.Level) + " Xp: " + strconv.Itoa(+result.Data.Xp))
	}
}

func ParseGetMap(rawLog []byte) {
	var result response.GetMap
	if err := json.Unmarshal(rawLog, &result); err != nil {
		fmt.Println("Can't unmarshal JSON")
	} else {
		fmt.Println("Name: " + result.Data.Skin + "  X: " + strconv.Itoa(result.Data.X) + " Y: " + strconv.Itoa(+result.Data.Y) + " Info: " + result.Data.Content.Code)
	}
}
