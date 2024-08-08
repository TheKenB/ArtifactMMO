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
[0]    Walk
[1]    Attack
[3]    Unequip
[4]    Gather
[5]    Craft
[6]    Bank Deposit
[7]    Bank Deposit Gold
[8]    Recycle Item
[9]    Withdraw From Bank
[10]   Withdraw Bank Gold
[11]   Ge Buy Item
[12]   Ge Sell Item
[13]   Accept New Task
[14]   Complete Task
[15]   Exchange Task
[16]   Delete Item
---------------------`
}

func NoCharacterActionMenu() string {
	return `
---------------------
[0] Locations
[1] Get All Character Logs
[2] Get Character List
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
	case "3":
		api.UnequipItem(char, val)
	case "4":
		api.Gathering(char, val)
	case "5":
		api.Crafting(char, val)
	case "6":
		api.DepositBank(char, val)
	case "7":
		api.DepositBankGold(char, val)
	case "8":
		api.DepositBankGold(char, val)
	case "9":
		api.WithdrawBank(char, val)
	case "10":
		api.WithdrawBankGold(char, val)
	case "11":
		api.GeBuyItem(char, val)
	case "12":
		api.GeSellItem(char, val)
	case "13":
		api.AcceptNewTask(char, val)
	case "14":
		api.CompleteTask(char, val)
	case "15":
		api.ExchangeTask(char, val)
	case "16":
		api.DeleteItem(char, val)
	default:
		fmt.Println("No Action")
	}
}

func GetNoCharacterAction(val string, char string) {
	switch val {
	case "0":
		api.GetAllMaps(val)
	case "1":
		api.GetAllCharLogs(val)
	case "2":
		api.GetMyCharacters(val)
	default:
		fmt.Println("No Actin")
	}
}
