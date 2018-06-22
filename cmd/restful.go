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
	"dalu/contested/restful"
	"dalu/contested/restful/repo"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/spf13/cobra"
)

// restfulCmd represents the restful command
var restfulCmd = &cobra.Command{
	Use:   "restful",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ms, e := mgo.Dial("localhost")
		if e != nil {
			return e
		}

		r := gin.Default()
		r.SecureJsonPrefix(")]}',\n")

		rp := repo.NewRepo(ms)
		defer rp.Close()
		h := restful.NewHandler(rp)

		v1 := r.Group("/api/v1")

		defaults := v1.Group("/defaults")
		defaults.GET("/setup", h.DefaultSetup)
		defaults.GET("/raid", h.DefaultRaid)
		defaults.GET("/classes", h.RaidClassList)

		census := v1.Group("/census")
		census.POST("/search", h.CensusSearch)

		profile := v1.Group("/profile")
		profile.GET("/", h.ProfileList)
		profile.GET("/:id", h.ProfileGet)
		profile.POST("/", h.ProfilePost)
		profile.PUT("/:id", h.ProfilePut)
		profile.DELETE("/:id", h.ProfileDelete)

		raids := v1.Group("/raids")
		raids.GET("/", h.RaidList)
		raids.GET("/:id", h.RaidGet)
		raids.POST("/", h.RaidPost)
		raids.PUT("/:id", h.RaidPut)
		raids.DELETE("/:id", h.RaidDelete)

		raidnotes := v1.Group("/raidnotes")
		raidnotes.GET("/", h.RaidNoteList)
		raidnotes.GET("/:id", h.RaidNoteGet)
		raidnotes.POST("/", h.RaidNotePost)
		raidnotes.PUT("/:id", h.RaidNotePut)
		raidnotes.DELETE("/:id", h.RaidNoteDelete)

		raiders := v1.Group("/raiders")
		raiders.GET("/", h.RaiderList)
		raiders.GET("/:id", h.RaiderGet)
		raiders.POST("/", h.RaiderPost)
		raiders.PUT("/:id", h.RaiderPut)
		raiders.DELETE("/:id", h.RaiderDelete)

		return r.RunUnix("/tmp/contested.de.sock")
	},
}

func init() {
	rootCmd.AddCommand(restfulCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restfulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restfulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
