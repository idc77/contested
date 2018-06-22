package models

type RaidSetup struct {
	MT      *RaidGroup  `json:"mt"`
	OT      *RaidGroup  `json:"ot"`
	DPS1    *RaidGroup  `json:"dps_1"`
	DPS2    *RaidGroup  `json:"dps_2"`
	Sitting []*RaidSlot `json:"sitting"`
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

type MTGroup struct {
	Tank                string   `json:"tank"`
	TemplarOrHealer     string   `json:"templar_or_healer"`
	DefilerOrWarder     string   `json:"defiler_or_warder"`
	CoercerOrEnchanter  string   `json:"coercer_or_enchanter"`
	DirgeOrBard         string   `json:"dirge_or_bard"`
	SwashOrHatetransfer string   `json:"swash_or_hatetransfer"`
	Rules               *MTRules `json:"rules"`
}

type MTRules struct {
	Tank                string `json:"tank"`
	TemplarOrHealer     string `json:"templar_or_healer"`
	DefilerOrWarder     string `json:"defiler_or_warder"`
	CoercerOrEnchanter  string `json:"coercer_or_enchanter"`
	DirgeOrBard         string `json:"dirge_or_bard"`
	SwashOrHatetransfer string `json:"swash_or_hatetransfer"`
}

type OTGroup struct {
	Tank               string   `json:"tank"`
	ClericOrHealer     string   `json:"cleric_or_healer"`
	ShamanOrHealer     string   `json:"shaman_or_healer"`
	CoercerOrEnchanter string   `json:"coercer_or_enchanter"`
	DirgeOrBard        string   `json:"dirge_or_bard"`
	DPSOrHatetransfer  string   `json:"dps_or_hatetransfer"`
	Rules              *OTRules `json:"rules"`
}

type OTRules struct {
	Tank               string `json:"tank"`
	ClericOrHealer     string `json:"cleric_or_healer"`
	ShamanOrHealer     string `json:"shaman_or_healer"`
	CoercerOrEnchanter string `json:"coercer_or_enchanter"`
	DirgeOrBard        string `json:"dirge_or_bard"`
	DPSOrHatetransfer  string `json:"dps_or_hatetransfer"`
}

type DPSGroup struct {
	Bard        string    `json:"bard"`
	Enchanter   string    `json:"enchanter"`
	HealerOrDPS string    `json:"healer_or_dps"`
	Healer      string    `json:"healer"`
	DPS         string    `json:"dps"`
	DPSOrTank   string    `json:"dps_or_tank"`
	Rules       *DPSRules `json:"rules"`
}

type DPSRules struct {
	Bard        string `json:"bard"`
	Enchanter   string `json:"enchanter"`
	HealerOrDPS string `json:"healer_or_dps"`
	Healer      string `json:"healer"`
	DPS         string `json:"dps"`
	DPSOrTank   string `json:"dps_or_tank"`
}

type RaidSlot struct {
	Who string
}
