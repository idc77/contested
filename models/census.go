package models

type CensusResult struct {
	CharacterList []*CensusCharacter `json:"character_list"`
}

type CensusParams struct {
	Id             string `url:"id,omitempty"`
	NameFirst      string `url:"name.first,omitempty"`
	NameFirstLower string `url:"name.first_lower,omitempty"`
	World          string `url:"locationdata.world,omitempty"`
}

type CensusCharacter struct {
	Stats struct {
		Sta struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"sta"`
		Power struct {
			Max   int `json:"max"`
			Regen int `json:"regen"`
		} `json:"power"`
		Weapon struct {
			Primarydelay       float64 `json:"primarydelay"`
			Primarymindamage   int     `json:"primarymindamage"`
			Primarymaxdamage   int     `json:"primarymaxdamage"`
			Rangeddelay        float64 `json:"rangeddelay"`
			Rangedmindamage    int     `json:"rangedmindamage"`
			Rangedmaxdamage    int     `json:"rangedmaxdamage"`
			Secondarydelay     float64 `json:"secondarydelay"`
			Secondarymindamage int     `json:"secondarymindamage"`
			Secondarymaxdamage int     `json:"secondarymaxdamage"`
		} `json:"weapon"`
		Wis struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"wis"`
		Runspeed float64 `json:"runspeed"`
		Int      struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"int"`
		Defense struct {
			Armor     int `json:"armor"`
			Avoidance int `json:"avoidance"`
			Block     int `json:"block"`
			Parry     int `json:"parry"`
		} `json:"defense"`
		Health struct {
			Max   int `json:"max"`
			Regen int `json:"regen"`
		} `json:"health"`
		Str struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"str"`
		Agi struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"agi"`
		Combat struct {
			Toughness                  float64 `json:"toughness"`
			MitigationElemental        float64 `json:"mitigation_elemental"`
			Aeautoattackchance         float64 `json:"aeautoattackchance"`
			Basemodifier               float64 `json:"basemodifier"`
			Abilitymod                 float64 `json:"abilitymod"`
			Baseavoidancebonus         float64 `json:"baseavoidancebonus"`
			Blockchance                float64 `json:"blockchance"`
			Attackspeed                float64 `json:"attackspeed"`
			MitigationNoxious          float64 `json:"mitigation_noxious"`
			Critbonus                  float64 `json:"critbonus"`
			Outofcombatsavageryregen   float64 `json:"outofcombatsavageryregen"`
			Abilitydoubleattackchance  float64 `json:"abilitydoubleattackchance"`
			Dissonancegainmod          float64 `json:"dissonancegainmod"`
			MitigationPhysical         float64 `json:"mitigation_physical"`
			Fervor                     float64 `json:"fervor"`
			Dissonanceregen            float64 `json:"dissonanceregen"`
			Savageryregen              float64 `json:"savageryregen"`
			Accuracy                   float64 `json:"accuracy"`
			Savagerygainmod            float64 `json:"savagerygainmod"`
			Strikethrough              float64 `json:"strikethrough"`
			Lethality                  float64 `json:"lethality"`
			Incombatsavageryregen      float64 `json:"incombatsavageryregen"`
			Flurry                     float64 `json:"flurry"`
			Outofcombatdissonanceregen float64 `json:"outofcombatdissonanceregen"`
			Dps                        float64 `json:"dps"`
			Doubleattackchance         float64 `json:"doubleattackchance"`
			Weapondamagebonus          float64 `json:"weapondamagebonus"`
			Spelldoubleattackchance    float64 `json:"spelldoubleattackchance"`
			Resolve                    float64 `json:"resolve"`
			Incombatdissonanceregen    float64 `json:"incombatdissonanceregen"`
			Maxsavagerylevel           float64 `json:"maxsavagerylevel"`
			Critchance                 float64 `json:"critchance"`
			Hategainmod                float64 `json:"hategainmod"`
			MitigationArcane           float64 `json:"mitigation_arcane"`
		} `json:"combat"`
		PersonalStatusPoints int `json:"personal_status_points"`
		Ability              struct {
			Spelltimereusepct       float64 `json:"spelltimereusepct"`
			Spelltimerecoverypct    float64 `json:"spelltimerecoverypct"`
			Spelltimereusespellonly float64 `json:"spelltimereusespellonly"`
			Spelltimecastpct        float64 `json:"spelltimecastpct"`
		} `json:"ability"`
	} `json:"stats"`
	Dbid       int     `json:"dbid"`
	Ts         float64 `json:"ts"`
	LastUpdate float64 `json:"last_update"`
	Version    int     `json:"version"`
	Resists    struct {
		Noxious struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"noxious"`
		Arcane struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"arcane"`
		Elemental struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"elemental"`
		Physical struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"physical"`
	} `json:"resists"`
	Type struct {
		Classid      int    `json:"classid"`
		AaLevel      int    `json:"aa_level"`
		TsLevel      int    `json:"ts_level"`
		Raceid       int    `json:"raceid"`
		Level        int    `json:"level"`
		Gender       string `json:"gender"`
		TsClass      string `json:"ts_class"`
		BirthdateUtc int    `json:"birthdate_utc"`
		Race         string `json:"race"`
		Deity        string `json:"deity"`
		Class        string `json:"class"`
		Alignment    int    `json:"alignment"`
	} `json:"type"`
	Deityabilitysystem struct {
		Availablepoints int `json:"availablepoints"`
	} `json:"deityabilitysystem"`
	Displayname  string `json:"displayname"`
	Locationdata struct {
		Homefaction int    `json:"homefaction"`
		World       string `json:"world"`
	} `json:"locationdata"`
	Name struct {
		Prefix     string `json:"prefix"`
		FirstLower string `json:"first_lower"`
		Last       string `json:"last"`
		Suffix     string `json:"suffix"`
		First      string `json:"first"`
	} `json:"name"`
	Account struct {
		Age int `json:"age"`
	} `json:"account"`
	Guild struct {
		Status  int    `json:"status"`
		Name    string `json:"name"`
		Level   int    `json:"level"`
		Joined  int    `json:"joined"`
		Rank    int    `json:"rank"`
		Guildid int    `json:"guildid"`
		ID      int64  `json:"id"`
	} `json:"guild,omitempty"`
	ID int64 `json:"id"`
}
