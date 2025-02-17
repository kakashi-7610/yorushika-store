/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"yorushika-store/product-manager/config"
	"yorushika-store/product-manager/repository"
	"yorushika-store/product-manager/server"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configs, err := config.NewConfigs()
		if err != nil {
			return err
		}
		outputSettingInfo(configs)

		db, err := repository.NewDatabase(configs.DbConfig)
		if err != nil {
			return err
		}

		// オートマイグレーション用
		// Model更新時のみ使用する
		// err = db.AutoMigrate()
		// if err != err {
		// 	return err
		// }

		s, err := server.NewServer(configs.ServerConfig, db)
		if err != nil {
			return err
		}

		ctx := context.Background()
		s.Run(ctx)
		defer s.Close()

		return nil
	},
}

func outputSettingInfo(cfgs *config.Configs) {
	log.Printf("Log config: %+v", cfgs.LogConfig)
	log.Printf("Server config: %+v", cfgs.ServerConfig)
}

func init() {
	serverCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
