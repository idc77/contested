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

	defaultSetup.MT = new(models.RaidGroup)
	defaultSetup.MT.Rules = new(models.RaidGroupRules)
	defaultSetup.MT.Rules.Slot1 = "tank"
	defaultSetup.MT.Rules.Slot2 = "dirge"
	defaultSetup.MT.Rules.Slot3 = "coercer"
	defaultSetup.MT.Rules.Slot4 = "templar"
	defaultSetup.MT.Rules.Slot5 = "defiler"
	defaultSetup.MT.Rules.Slot6 = "hatetransfer"

	defaultSetup.OT = new(models.RaidGroup)
	defaultSetup.OT.Rules = new(models.RaidGroupRules)
	defaultSetup.OT.Rules.Slot1 = "tank"
	defaultSetup.OT.Rules.Slot2 = "dirge"
	defaultSetup.OT.Rules.Slot3 = "coercer"
	defaultSetup.OT.Rules.Slot4 = "healer"
	defaultSetup.OT.Rules.Slot5 = "shaman"
	defaultSetup.OT.Rules.Slot6 = "scout,hatetransfer"

	defaultSetup.DPS1 = new(models.RaidGroup)
	defaultSetup.DPS1.Rules = new(models.RaidGroupRules)
	defaultSetup.DPS1.Rules.Slot1 = "dps,tank"
	defaultSetup.DPS1.Rules.Slot2 = "troubador"
	defaultSetup.DPS1.Rules.Slot3 = "illusionist"
	defaultSetup.DPS1.Rules.Slot4 = "mage"
	defaultSetup.DPS1.Rules.Slot5 = "dps,healer"
	defaultSetup.DPS1.Rules.Slot6 = "healer"

	defaultSetup.DPS2 = new(models.RaidGroup)
	defaultSetup.DPS2.Rules = new(models.RaidGroupRules)
	defaultSetup.DPS2.Rules.Slot1 = "dps,tank"
	defaultSetup.DPS2.Rules.Slot2 = "bard"
	defaultSetup.DPS2.Rules.Slot3 = "enchanter"
	defaultSetup.DPS2.Rules.Slot4 = "dps"
	defaultSetup.DPS2.Rules.Slot5 = "healer"
	defaultSetup.DPS2.Rules.Slot6 = "dps,healer"
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
	m.Roles = []string{"swashbuckler", "scout", "dps", "rogue", "hatetransfer"}
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
