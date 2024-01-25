/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"BookStoreApi-Go/Controller"
	"BookStoreApi-Go/Routes"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "BookStoreApi",
	Long:  `BookStore Api Server.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		Controller.TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)
		Routes.Start()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
