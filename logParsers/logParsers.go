package logParsers

import (
	"encoding/json"
	"fmt"
	"time"
)

type CombatResponse struct {
	Data struct {
		Cooldown struct {
			Total_seconds     float32 `json:"total_seconds"`
			Remaining_seconds float32 `json:"remaining_seconds"`
			TotalSeconds      float32 `json:"totalSeconds"`
			RemainingSeconds  float32 `json:"remainingSeconds"`
			Expiration        time.Time
			Reason            string `json:"reason"`
		} `json:"cooldown"`
		Fight struct {
			Xp    int `json:"xp"`
			Gold  int `json:"gold"`
			Drops []struct {
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"drops"`
			Turns                int `json:"turns"`
			Monster_blocked_hits struct {
				Fire  int `json:"fire"`
				Earth int `json:"earth"`
				Water int `json:"water"`
				Air   int `json:"air"`
				Total int `json:"total"`
			} `json:"monster_blocked_hits"`
			Player_blocked_hits struct {
				Fire  int `json:"fire"`
				Earth int `json:"earth"`
				Water int `json:"water"`
				Air   int `json:"air"`
				Total int `json:"total"`
			} `json:"player_blocked_hits"`
			Logs   []string `json:"logs"`
			Result string   `json:"result"`
		} `json:"fight"`
		Character struct {
			Name                      string    `json:"name"`
			Skin                      string    `json:"skin"`
			Level                     int       `json:"level"`
			Xp                        int       `json:"xp"`
			Max_xp                    int       `json:"max_xp"`
			Total_xp                  int       `json:"total_xp"`
			Gold                      int       `json:"gold"`
			Speed                     int       `json:"speed"`
			Mining_level              int       `json:"mining_level"`
			Mining_xp                 int       `json:"mining_xp"`
			Mining_max_xp             int       `json:"mining_max_xp"`
			Woodcutting_level         int       `json:"woodcutting_level"`
			Woodcutting_xp            int       `json:"woodcutting_xp"`
			Woodcutting_max_xp        int       `json:"woodcutting_max_xp"`
			Fishing_level             int       `json:"fishing_level"`
			Fishing_xp                int       `json:"fishing_xp"`
			Fishing_max_xp            int       `json:"fishing_max_xp"`
			Weaponcrafting_level      int       `json:"weaponcrafting_level"`
			Weaponcrafting_xp         int       `json:"weaponcrafting_xp"`
			Weaponcrafting_max_xp     int       `json:"weaponcrafting_max_xp"`
			Gearcrafting_level        int       `json:"gearcrafting_level"`
			Gearcrafting_xp           int       `json:"gearcrafting_xp"`
			Gearcrafting_max_xp       int       `json:"gearcrafting_max_xp"`
			Jewelrycrafting_level     int       `json:"jewelrycrafting_level"`
			Jewelrycrafting_xp        int       `json:"jewelrycrafting_xp"`
			Jewelrycrafting_max_xp    int       `json:"jewelrycrafting_max_xp"`
			Cooking_level             int       `json:"cooking_level"`
			Cooking_xp                int       `json:"cooking_xp"`
			Cooking_max_xp            int       `json:"cooking_max_xp"`
			Hp                        int       `json:"hp"`
			Haste                     int       `json:"haste"`
			Critical_strike           int       `json:"critical_strike"`
			Stamina                   int       `json:"stamina"`
			Attack_fire               int       `json:"attack_fire"`
			Attack_earth              int       `json:"attack_earth"`
			Attack_water              int       `json:"attack_water"`
			Attack_air                int       `json:"attack_air"`
			Dmg_fire                  int       `json:"dmg_fire"`
			Dmg_earth                 int       `json:"dmg_earth"`
			Dmg_water                 int       `json:"dmg_water"`
			Dmg_air                   int       `json:"dmg_air"`
			Res_fire                  int       `json:"res_fire"`
			Res_earth                 int       `json:"res_earth"`
			Res_water                 int       `json:"res_water"`
			Res_air                   int       `json:"res_air"`
			X                         int       `json:"x"`
			T                         int       `json:"y"`
			Cooldown                  int       `json:"cooldown"`
			Cooldown_expiration       time.Time `json:"cooldown_expiration"`
			Weapon_slot               string    `json:"weapon_slot"`
			Shield_slot               string    `jsong:"shield_slot"`
			Helmet_slot               string    `json:"helmet_slot"`
			Body_armor_slot           string    `json:"body_armor_slot"`
			Leg_armor_slot            string    `json:"leg_armor_slot"`
			Boots_slot                string    `json:"boots_slot"`
			Ring1_slot                string    `json:"ring1_slot"`
			Ring2_slot                string    `json:"ring2_slot"`
			Amulet_slot               string    `json:"amulet_slot"`
			Artifact1_slot            string    `json:"artifact1_slot"`
			Artifact2_slot            string    `json:"artifact2_slot"`
			Artifact3_slot            string    `json:"artifact3_slot"`
			Consumable1_slot          string    `json:"consumable1_slot"`
			Consumable1_slot_quantity int       `json:"consumable1_slot_quantity"`
			Consumable2_slot          string    `json:"consumable2_slot"`
			Consumable2_slot_quantity int       `json:"consumable2_slot_quantity"`
			Task                      string    `json:"task"`
			Task_type                 string    `json:"task_type"`
			Task_progress             int       `json:"task_progress"`
			Task_total                int       `json:"task_total"`
			Inventory_max_items       int       `json:"inventory_max_items"`
			Inventory                 []struct {
				Slot     int    `json:"slot"`
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"inventory"`
			Inventory_slot1           string `json:"inventory_slot1"`
			Inventory_slot1_quantity  int    `json:"inventory_slot1_quantity"`
			Inventory_slot2           string `json:"inventory_slot2"`
			Inventory_slot2_quantity  int    `json:"inventory_slot2_quantity"`
			Inventory_slot3           string `json:"inventory_slot3"`
			Inventory_slot3_quantity  int    `json:"inventory_slot3_quantity"`
			Inventory_slot4           string `json:"inventory_slot4"`
			Inventory_slot4_quantity  int    `json:"inventory_slot4_quantity"`
			Inventory_slot5           string `json:"inventory_slot5"`
			Inventory_slot5_quantity  int    `json:"inventory_slot5_quantity"`
			Inventory_slot6           string `json:"inventory_slot6"`
			Inventory_slot6_quantity  int    `json:"inventory_slot6_quantity"`
			Inventory_slot7           string `json:"inventory_slot7"`
			Inventory_slot7_quantity  int    `json:"inventory_slot7_quantity"`
			Inventory_slot8           string `json:"inventory_slot8"`
			Inventory_slot8_quantity  int    `json:"inventory_slot8_quantity"`
			Inventory_slot9           string `json:"inventory_slot9"`
			Inventory_slot9_quantity  int    `json:"inventory_slot9_quantity"`
			Inventory_slot10          string `json:"inventory_slot10"`
			Inventory_slot10_quantity int    `json:"inventory_slot10_quantity"`
			Inventory_slot11          string `json:"inventory_slot11"`
			Inventory_slot11_quantity int    `json:"inventory_slot11_quantity"`
			Inventory_slot12          string `json:"inventory_slot12"`
			Inventory_slot12_quantity int    `json:"inventory_slot12_quantity"`
			Inventory_slot13          string `json:"inventory_slot13"`
			Inventory_slot13_quantity int    `json:"inventory_slot13_quantity"`
			Inventory_slot14          string `json:"inventory_slot14"`
			Inventory_slot14_quantity int    `json:"inventory_slot14_quantity"`
			Inventory_slot15          string `json:"inventory_slot15"`
			Inventory_slot15_quantity int    `json:"inventory_slot15_quantity"`
			Inventory_slot16          string `json:"inventory_slot16"`
			Inventory_slot16_quantity int    `json:"inventory_slot16_quantity"`
			Inventory_slot17          string `json:"inventory_slot17"`
			Inventory_slot17_quantity int    `json:"inventory_slot17_quantity"`
			Inventory_slot18          string `json:"inventory_slot18"`
			Inventory_slot18_quantity int    `json:"inventory_slot18_quantity"`
			Inventory_slot19          string `json:"inventory_slot19"`
			Inventory_slot19_quantity int    `json:"inventory_slot19_quantity"`
			Inventory_slot20          string `json:"inventory_slot20"`
			Inventory_slot20_quantity int    `json:"inventory_slot20_quantity"`
		} `json:"character"`
	} `json:"data"`
}

func ParseCombatLog(rawLog []byte) {
	var result CombatResponse
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
