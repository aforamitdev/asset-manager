package main

import (
	"changeme/business/sqllite"
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/anchore/go-logger"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// var log *logger.Logger
	// Create an instance of the app structure
	var log *logger.Logger

	db, err := sqllite.NewDatabase()

	defer db.Close()

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	var id int
	rows, err := db.Query("SELECT * FROM test")
	fmt.Println(err, "s")
	for rows.Next() {
		rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)

	}

	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "asset-manager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	// response, err := http.Get("https://www.nseindia.com/api/marketStatus")

	// if err != nil {
	// 	fmt.Println(err.Error())

	// }
	// responseData, err := io.ReadAll(response.Body)
	// fmt.Println(string(responseData))

	// if err != nil {
	// 	println("Error:", err.Error())
	// }
}

func run(ctx context.Context, log *log.Logger) error {

	log.in

}
