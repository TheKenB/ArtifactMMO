package responseType

import "time"

type ActionResponse struct {
	Url     string
	PostGet string
	Data    string
}

type JsonCharacterResponse struct {
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
	Woodcutting_xp            int       `JSON:"WOODCUTTING_XP"`
	Woodcutting_max_xp        int       `JSON:"WOODCUTTING_MAX_XP"`
	Fishing_level             int       `JSON:"FISHING_LEVEL"`
	Fishing_xp                int       `JSON:"FISHING_XP"`
	Fishing_max_xp            int       `JSON:"FISHING_MAX_XP"`
	Weaponcrafting_level      int       `JSON:"WEAPONCRAFTING_LEVEL"`
	Weaponcrafting_xp         int       `JSON:"WEAPONCRAFTING_XP"`
	Weaponcrafting_max_xp     int       `JSON:"WEAPONCRAFTING_MAX_XP"`
	Gearcrafting_level        int       `JSON:"GEARCRAFTING_LEVEL"`
	Gearcrafting_xp           int       `JSON:"GEARCRAFTING_XP"`
	Gearcrafting_max_xp       int       `JSON:"GEARCRAFTING_MAX_XP"`
	Jewelrycrafting_level     int       `JSON:"JEWELRYCRAFTING_LEVEL"`
	Jewelrycrafting_xp        int       `JSON:"JEWELRYCRAFTING_XP"`
	Jewelrycrafting_max_xp    int       `JSON:"JEWELRYCRAFTING_MAX_XP"`
	Cooking_level             int       `JSON:"COOKING_LEVEL"`
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
	Ring0_slot                string    `json:"ring1_slot"`
	Ring1_slot                string    `json:"ring2_slot"`
	Amulet_slot               string    `json:"amulet_slot"`
	Artifact0_slot            string    `json:"artifact1_slot"`
	Artifact1_slot            string    `json:"artifact2_slot"`
	Artifact2_slot            string    `json:"artifact3_slot"`
	Consumable0_slot          string    `json:"consumable1_slot"`
	Consumable0_slot_quantity int       `json:"consumable1_slot_quantity"`
	Consumable1_slot          string    `json:"consumable2_slot"`
	Consumable1_slot_quantity int       `json:"consumable2_slot_quantity"`
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
}

type CooldownMapResponse struct {
	Total_seconds     float32   `json:"total_seconds"`
	Remaining_seconds float32   `json:"remaining_seconds"`
	StartedAt         time.Time `json:"started_at"`
	Expiration        time.Time `jsong:"experation"`
	Reason            string    `json:"reason"`
}

type ItemMapResponse struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Level       int    `json:"level"`
	Type        string `json:"type"`
	Subtype     string `json:"subtype"`
	Description string `json:"description"`
	Effects     []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"effect"`
	Craft struct {
		Skill string `json:"skill"`
		Level int    `json:"level"`
		Items []struct {
			Code     string `json:"code"`
			Quantity int    `json:"quantity"`
		} `json:"items"`
		Quantity int `json:"quantity"`
	} `json:"craft"`
}

type CombatResponse struct {
	Data struct {
		CooldownMapResponse `json:"cooldown"`
		Fight               struct {
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
		Character JsonCharacterResponse `json:"Character"`
	} `json:"data"`
}

type AllMapsResponse struct {
	Data []struct {
		Name    string `json:"name"`
		Skin    string `json:"skin"`
		X       int    `json:"x"`
		Y       int    `json:"y"`
		Content struct {
			Type string `json:"type"`
			Code string `json:"code"`
		} `json:"content"`
	} `json:"data"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	Pages int `json:"pages"`
}

type DropResponse struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type BlockedHitsResponse struct {
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Water int `json:"water"`
	Air   int `json:"air"`
	Total int `json:"total"`
}

type FightResponse struct {
	Xp             int                 `json:"xp"`
	Gold           int                 `json:"gold"`
	Drops          []DropResponse      `json:"drops"`
	Turns          int                 `json:"turns"`
	MonsterBlocked BlockedHitsResponse `json:"monster_blocked_hits"`
	PlayerBlocked  BlockedHitsResponse `json:"player_blocked_hits"`
	Logs           []string            `json:"logs"`
	Result         string              `json:"result"`
}

type GetMyCharactersResponse struct {
	Data []JsonCharacterResponse `json:"data"`
}

type EquipingResponse struct {
	Data struct {
		Cooldown CooldownMapResponse `json:"cooldown"`
		Slot     string              `json:"slot"`
		Item     ItemMapResponse     `json:"item"`
	} `json:"data"`
	Character JsonCharacterResponse `json:"character"`
}

type ItemsResponse struct {
	Code     string `json:"code"`
	Quantity int    `jsong:"quantity"`
}

type GatherDetails struct {
	Xp int `json:"xp"`
}
