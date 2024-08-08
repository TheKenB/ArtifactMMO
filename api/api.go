package api

import (
	parse "artifactMMO/logParsers"
	response "artifactMMO/responseType"
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

func Move(char string, key string) {
	var x string
	var y string

	fmt.Println("Enter x y location")
	_, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println("X,Y Error:", err)
	}
	data := `{"x":` + x + `,"y":` + y + `}`
	resp := response.ActionResponse{Url: "https://api.artifactsmmo.com/my/" + char + "/action/move", PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func Fight(char string, key string) {
	resp := response.ActionResponse{Url: "https://api.artifactsmmo.com/my/" + char + "/action/fight", PostGet: "POST", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func UnequipItem(char string, key string) {
	slot := ItemSlot()
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/unequip"
	data := `{"slot": "` + slot + `" }`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func Gathering(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/DaGuile/action/gathering"
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func Crafting(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/crafting"
	var item string
	var quantity string
	fmt.Println("Enter item and quantity")
	_, err := fmt.Scan(&item, &quantity)
	if err != nil {
		fmt.Println("Crafting Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func DepositBank(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/bank/deposit"
	var item string
	var quantity string
	fmt.Println("Enter item and quantity")
	_, err := fmt.Scan(&item, &quantity)
	if err != nil {
		fmt.Println("Bank Deposit Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func DepositBankGold(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/bank/deposit/gold"
	var gold string
	fmt.Println("Enter gold quantity deposit")
	_, err := fmt.Scan(&gold)
	if err != nil {
		fmt.Println("Bank Deposit Error:", err)
	}
	data := `{"quantity":` + gold + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func Recycling(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/recycling"
	var item string
	var quantity string
	fmt.Println("Enter item and quantity")
	_, err := fmt.Scan(&item, &quantity)
	if err != nil {
		fmt.Println("Bank Deposit Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func WithdrawBank(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/bank/withdraw"
	var item string
	var quantity string
	fmt.Println("Enter item and quantity")
	_, err := fmt.Scan(&item, &quantity)
	if err != nil {
		fmt.Println("Bank Withdraw Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func WithdrawBankGold(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/bank/withdraw/gold"
	var gold string
	fmt.Println("Enter gold quantity to withdraw")
	_, err := fmt.Scan(&gold)
	if err != nil {
		fmt.Println("Bank Deposit Error:", err)
	}
	data := `{"quantity":` + gold + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func GeBuyItem(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/ge/buy"
	var item string
	var quantity string
	var price string
	fmt.Println("Enter item, quantity, price")
	_, err := fmt.Scan(&item, &quantity, &price)
	if err != nil {
		fmt.Println("Ge Item Buy Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `,"price":` + price + `}`
	fmt.Println(data)
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func GeSellItem(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/ge/sell"
	var item string
	var quantity string
	var price string
	fmt.Println("Enter item, quantity, price")
	_, err := fmt.Scan(&item, &quantity, &price)
	if err != nil {
		fmt.Println("Ge Item Buy Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `,"price":` + price + `}`
	fmt.Println(data)
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func AcceptNewTask(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/task/new"
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func CompleteTask(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/task/complete"
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func ExchangeTask(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/task/exchange"
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

func DeleteItem(char string, key string) {
	tempUrl := "https://api.artifactsmmo.com/my/" + char + "/action/delete"
	var item string
	var quantity string
	fmt.Println("Enter item and quantity")
	_, err := fmt.Scan(&item, &quantity)
	if err != nil {
		fmt.Println("Delete Item Error:", err)
	}
	data := `{"code":` + `"` + item + `"` + `,"quantity":` + quantity + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "POST", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, char)
}

// Character agnostic
func GetAllMaps(key string) {
	var page string
	fmt.Println("Enter page number")
	_, err := fmt.Scan(&page)
	if err != nil {
		fmt.Println("X,Y Error:", err)
	}
	resp := response.ActionResponse{Url: "https://api.artifactsmmo.com/maps/?page=" + page + "&size=100", PostGet: "GET", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, "None")
}

func GetMyCharacters(key string) {
	resp := response.ActionResponse{Url: "https://api.artifactsmmo.com/my/characters", PostGet: "GET", Data: ""}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, "None")
}

func GetAllCharLogs(key string) {
	tempUrl := "https://api.artifactsmmo.com/my/logs"
	var page string
	var size string
	fmt.Println("Enter page and page size")
	_, err := fmt.Scan(&page, &size)
	if err != nil {
		fmt.Println("Get Character Logs Error:", err)
	}
	data := `{"page":` + `"` + page + `"` + `,"size":` + size + `}`
	resp := response.ActionResponse{Url: tempUrl, PostGet: "GET", Data: data}
	body := ExecQuery(resp)
	parse.ParseOperator(key, body, "None")
}

// Helpers
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
