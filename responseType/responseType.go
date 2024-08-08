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
type EffectsResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ItemsResponse struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type CraftMapResponse struct {
	Skill    string          `json:"skill"`
	Level    int             `json:"level"`
	Items    []ItemsResponse `json:"items"`
	Quantity int             `json:"quantity"`
}

type BankResponse struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type BankGoldResponse struct {
	Quantity int `json:"quantity"`
}

type ItemMapResponse struct {
	Name        string            `json:"name"`
	Code        string            `json:"code"`
	Level       int               `json:"level"`
	Type        string            `json:"type"`
	Subtype     string            `json:"subtype"`
	Description string            `json:"description"`
	Effects     []EffectsResponse `json:"effect"`
	Craft       struct {
		Skill string `json:"skill"`
		Level int    `json:"level"`
		Items []struct {
			Code     string `json:"code"`
			Quantity int    `json:"quantity"`
		} `json:"items"`
		Quantity int `json:"quantity"`
	} `json:"craft"`
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

type GatherDetailsResponse struct {
	Xp    int             `json:"xp"`
	Items []ItemsResponse `json:"itmes"`
}

type RecycleDetailResponse struct {
	Details struct {
		Items []ItemsResponse `json:"items"`
	} `json:"details"`
}

type TransactionResponse struct {
	Code       string `json:"code"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
}

type TaskResponse struct {
	Code  string `json:"code"`
	Type  string `json:"type"`
	Total int    `json:"total"`
}

type RewardResponse struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type AllCharacters struct {
	Data []struct {
		Character          string    `json:"character"`
		Account            string    `json:"account"`
		Type               string    `json:"type"`
		Description        string    `json:"description"`
		Content            string    `json:"content"`
		Cooldown           int       `json:"cooldown"`
		CooldownExpiration time.Time `json:"cooldown_expiration"`
		CreateAt           time.Time `json:"created_at"`
	}
}

// Api Responses

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

type GatherResponse struct {
	Data struct {
		Cooldown  CooldownMapResponse   `json:"cooldown"`
		Details   GatherDetailsResponse `json:"details"`
		Character JsonCharacterResponse `json:"character"`
	}
}

type CraftResponse struct {
	Data struct {
		Cooldown  CooldownMapResponse   `json:"cooldown"`
		Details   GatherDetailsResponse `json:"details"`
		Character JsonCharacterResponse `json:"character"`
	}
}

type BankDepositResponse struct {
	Data struct {
		Cooldown  CooldownMapResponse   `json:"cooldown"`
		Item      ItemMapResponse       `json:"item"`
		Bank      []BankResponse        `json:"bank"`
		Character JsonCharacterResponse `json:"character"`
	}
}
