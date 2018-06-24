package models

type RaidSetup struct {
	MT      *RaidGroup  `json:"mt"`
	OT      *RaidGroup  `json:"ot"`
	DPS1    *RaidGroup  `json:"dps_1"`
	DPS2    *RaidGroup  `json:"dps_2"`
	Sitting []*RaidSlot `json:"sitting,omitempty"`
	Sitsize int         `json:"sitsize"`
}

type RaidGroup struct {
	Slot1 string          `json:"slot_1"`
	Slot2 string          `json:"slot_2"`
	Slot3 string          `json:"slot_3"`
	Slot4 string          `json:"slot_4"`
	Slot5 string          `json:"slot_5"`
	Slot6 string          `json:"slot_6"`
	Rules *RaidGroupRules `json:"rules"`
}

type RaidGroupRules struct {
	Slot1 string `json:"slot_1"`
	Slot2 string `json:"slot_2"`
	Slot3 string `json:"slot_3"`
	Slot4 string `json:"slot_4"`
	Slot5 string `json:"slot_5"`
	Slot6 string `json:"slot_6"`
}

type RaidSlot struct {
	Designation string
	Raider      string
	Rule        string
}
