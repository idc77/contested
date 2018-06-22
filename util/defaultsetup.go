package util

import "dalu/contested/models"

var (
	raidClasses  []*models.RaidClass
	defaultSetup *models.RaidSetup
)

func init() {
	initDefaultSetup()
	initRaidClasses()
}

func DefaultSetup() *models.RaidSetup {
	return defaultSetup
}

func initDefaultSetup() {
	defaultSetup = new(models.RaidSetup)
	defaultSetup.Sitsize = 4

	defaultSetup.MT = new(models.MTGroup)
	defaultSetup.MT.Rules = new(models.MTRules)
	defaultSetup.MT.Rules.Tank = "tank"
	defaultSetup.MT.Rules.DirgeOrBard = "dirge"
	defaultSetup.MT.Rules.CoercerOrEnchanter = "coercer"
	defaultSetup.MT.Rules.TemplarOrHealer = "templar"
	defaultSetup.MT.Rules.DefilerOrWarder = "defiler"
	defaultSetup.MT.Rules.SwashOrHatetransfer = "hatetransfer"

	defaultSetup.OT = new(models.OTGroup)
	defaultSetup.OT.Rules = new(models.OTRules)
	defaultSetup.OT.Rules.Tank = "tank"
	defaultSetup.OT.Rules.CoercerOrEnchanter = "coercer"
	defaultSetup.OT.Rules.DirgeOrBard = "dirge"
	defaultSetup.OT.Rules.ClericOrHealer = "healer"
	defaultSetup.OT.Rules.ShamanOrHealer = "shaman"
	defaultSetup.OT.Rules.DPSOrHatetransfer = "dps"

	defaultSetup.DPS1 = new(models.DPSGroup)
	defaultSetup.DPS1.Rules = new(models.DPSRules)
	defaultSetup.DPS1.Rules.Bard = "troubador"
	defaultSetup.DPS1.Rules.Enchanter = "illusionist"
	defaultSetup.DPS1.Rules.DPS = "mage"

	defaultSetup.DPS2 = new(models.DPSGroup)
	defaultSetup.DPS2.Rules = new(models.DPSRules)
}

func DefaultRaid() *models.Raid {
	r := new(models.Raid)
	return r
}

func initRaidClasses() {
	m := new(models.RaidClass)
	m.Name = "Shadowknight"
	m.Roles = []string{"shadowknight", "crusader", "tank"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Paladin"
	m.Roles = []string{"paladin", "crusader", "tank"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Guardian"
	m.Roles = []string{"guardian", "tank", "warrior"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Berserker"
	m.Roles = []string{"berserker", "tank", "warrior"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Bruiser"
	m.Roles = []string{"bruiser", "tank", "brawler"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Monk"
	m.Roles = []string{"monk", "tank", "brawler"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Brigand"
	m.Roles = []string{"brigand", "scout", "dps", "rogue"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Swashbuckler"
	m.Roles = []string{"swashbuckler", "scout", "dps", "rogue"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Assassin"
	m.Roles = []string{"assassin", "scout", "dps", "predator", "hatetransfer"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Ranger"
	m.Roles = []string{"ranger", "scout", "dps", "predator"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Dirge"
	m.Roles = []string{"dirge", "scout", "bard"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Troubador"
	m.Roles = []string{"troubador", "scout", "bard"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Beastlord"
	m.Roles = []string{"beastlord", "scout", "dps"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Defiler"
	m.Roles = []string{"defiler", "healer", "shaman", "warder"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Mystic"
	m.Roles = []string{"mystic", "healer", "shaman", "warder"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Templar"
	m.Roles = []string{"templar", "healer", "cleric"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Inquisitor"
	m.Roles = []string{"inquisitor", "healer", "cleric"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Warden"
	m.Roles = []string{"warden", "healer", "druid"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Fury"
	m.Roles = []string{"fury", "healer", "druid"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Channeller"
	m.Roles = []string{"channeller", "healer"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Necromancer"
	m.Roles = []string{"necromancer", "mage", "dps"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Conjuror"
	m.Roles = []string{"conjuror", "mage", "dps"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Wizard"
	m.Roles = []string{"wizard", "mage", "dps"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Warlock"
	m.Roles = []string{"warlock", "mage", "dps"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Coercer"
	m.Roles = []string{"coercer", "mage", "enchanter"}
	raidClasses = append(raidClasses, m)

	m = new(models.RaidClass)
	m.Name = "Illusionist"
	m.Roles = []string{"illusionist", "mage", "enchanter"}

}

func RaidClasses() []*models.RaidClass {
	return raidClasses
}
