// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"dalu/contested/models"

	"fmt"

	"github.com/globalsign/mgo"
	"github.com/spf13/cobra"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ms, e := mgo.Dial("localhost")
		if e != nil {
			return e
		}
		defer ms.Close()
		classC := ms.DB("contested").C("classes")

		m := new(models.RaidClass)
		m.Name = "Shadowknight"
		m.Roles = []string{"shadowknight", "crusader", "tank"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Paladin"
		m.Roles = []string{"paladin", "crusader", "tank"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Guardian"
		m.Roles = []string{"guardian", "tank", "warrior"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Berserker"
		m.Roles = []string{"berserker", "tank", "warrior"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Bruiser"
		m.Roles = []string{"bruiser", "tank", "brawler"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Monk"
		m.Roles = []string{"monk", "tank", "brawler"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Brigand"
		m.Roles = []string{"brigand", "scout", "dps", "rogue"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Swashbuckler"
		m.Roles = []string{"swashbuckler", "scout", "dps", "rogue"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Assassin"
		m.Roles = []string{"assassin", "scout", "dps", "predator", "hatetransfer"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Ranger"
		m.Roles = []string{"ranger", "scout", "dps", "predator"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Dirge"
		m.Roles = []string{"dirge", "scout", "bard"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Troubador"
		m.Roles = []string{"troubador", "scout", "bard"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Beastlord"
		m.Roles = []string{"beastlord", "scout", "dps"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Defiler"
		m.Roles = []string{"defiler", "healer", "shaman", "warder"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Mystic"
		m.Roles = []string{"mystic", "healer", "shaman", "warder"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Templar"
		m.Roles = []string{"templar", "healer", "cleric"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Inquisitor"
		m.Roles = []string{"inquisitor", "healer", "cleric"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Warden"
		m.Roles = []string{"warden", "healer", "druid"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Fury"
		m.Roles = []string{"fury", "healer", "druid"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Channeller"
		m.Roles = []string{"channeller", "healer"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Necromancer"
		m.Roles = []string{"necromancer", "mage", "dps"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Conjuror"
		m.Roles = []string{"conjuror", "mage", "dps"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Wizard"
		m.Roles = []string{"wizard", "mage", "dps"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Warlock"
		m.Roles = []string{"warlock", "mage", "dps"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Coercer"
		m.Roles = []string{"coercer", "mage", "enchanter"}
		if e := classC.Insert(m); e != nil {
			return e
		}
		m = new(models.RaidClass)
		m.Name = "Illusionist"
		m.Roles = []string{"illusionist", "mage", "enchanter"}
		if e := classC.Insert(m); e != nil {
			return e
		}

		fmt.Println("all done.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixturesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixturesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
