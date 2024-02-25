package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

func configStyles() {
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorBlack,
		ContrastBackgroundColor:     tcell.ColorDarkBlue,
		MoreContrastBackgroundColor: tcell.ColorGreen,
		BorderColor:                 tcell.ColorWhite,
		TitleColor:                  tcell.ColorWhite,
		GraphicsColor:               tcell.ColorWhite,
		PrimaryTextColor:            tcell.ColorGhostWhite,
		SecondaryTextColor:          tcell.ColorYellow,
		TertiaryTextColor:           tcell.ColorGreen,
		InverseTextColor:            tcell.ColorDeepSkyBlue,
		ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
	}
}

var cfgFile string

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./business/data/schema/sql", "directory with migration files")
)

type MigrationFiles struct {
	name   string
	widget tview.Primitive
}

var rootCmd = &cobra.Command{
	Use:   "dp",
	Short: " ",
}

// func init() {
// 	fmt.Println()
// 	cobra.OnInitialize(initConfig)

// 	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config files ")
// }

func main() {

	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// 	fmt.Println("---", time.Now().Format(""))
	// }

	// fmt.Println(cfgFile)
	// // box := tview.NewBox().SetBorder(true).SetTitle("App Migrator")
	// // if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// // 	panic(err)
	// // }

	// app := tview.NewApplication()
	// tasks := tasks(app)

	// pages := createPages(tasks)

	// flags.Parse(os.Args[1:])

	// args := flags.Args()

	// if len(args) < 3 {
	// 	flags.Usage()
	// 	return
	// }

	// dbstring, command := args[1], args[2]

	// db, err := goose.OpenDBWithDriver("sqlite", dbstring)
	// goose.Create(db, "/business/data/schema/sql", "test", "up")
	// if err != nil {
	// 	log.Fatalf("goose: failed to open DB: %v\n", err)
	// }

	// defer func() {
	// 	if err := db.Close(); err != nil {
	// 		log.Fatalf("goose: failed to close DB: %v\n", err)
	// 	}
	// }()

	// arguments := []string{}
	// if len(args) > 3 {
	// 	arguments = append(arguments, args[3:]...)
	// }

	// if err := goose.Run(command, db, *dir, arguments...); err != nil {
	// 	log.Fatalf("goose %v: %v", command, err)
	// }

	configStyles()

	app := tview.NewApplication()

	files := []MigrationFiles{
		{name: "test"},
		{name: "test 2"},
	}
	fileWindow := createFilesWindows(files)

	sidebar := createSideWindows(files, fileWindow)

	// set flex
	flex := tview.NewFlex()
	done := make(chan struct{})

	// sidebar.AddPage("Migrations ", tview.NewTextView().SetText("Test"), true, true)

	// sidebar.SetBorder(true)
	fileWindow.SetBorder(true)

	flex.AddItem(sidebar, 0, 1, true)
	flex.AddItem(fileWindow, 0, 2, false)
	app.SetRoot(flex, true)

	focusingMenu := false

	app.EnableMouse(true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()

		if key != tcell.KeyTab && key != tcell.KeyBacktab {
			return event
		}

		if focusingMenu {
			app.SetFocus(sidebar)
		} else {
			app.SetFocus(fileWindow)
		}
		focusingMenu = !focusingMenu
		return nil

	})

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(
			signals,
			syscall.SIGILL, syscall.SIGINT, syscall.SIGTERM,
		)
		<-signals
		app.Stop()
	}()

	<-done

}

func initConfig() {
	fmt.Println(cfgFile)
	if cfgFile != "" {

		fmt.Println("file found")

	} else {

		fmt.Println("Error reading config files")

	}
}

func createFilesWindows(files []MigrationFiles) *tview.Pages {

	pages := tview.NewPages()
	for i, f := range files {
		var view tview.Primitive
		if f.widget != nil {
			view = f.widget
		} else {
			view = tview.NewTextView().SetText(f.name)
		}
		pages.AddPage(f.name, view, true, i == 0)
	}

	pages.SetBorder(true)
	return pages

}

func createSideWindows(files []MigrationFiles, pages *tview.Pages) tview.Primitive {
	menu := createMenu(files, pages)

	frame := tview.NewFrame(menu)
	frame.SetBorder(true)
	frame.AddText("Migrations Pilot", true, tview.AlignLeft, tcell.ColorWhite)
	return frame

}

func createMenu(tasks []MigrationFiles, pages *tview.Pages) *tview.List {
	menu := tview.NewList()
	menuWidth := 0
	for i, t := range tasks {
		if len(t.name) > menuWidth {
			menuWidth = len(t.name)
		}
		menu.AddItem(t.name, "", rune(i+'0'), nil)
	}

	menu.ShowSecondaryText(false)
	menu.SetChangedFunc(func(_ int, task string, _ string, _ rune) {
		pages.SwitchToPage(task)
	})

	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() != tcell.KeyRune {
			return event
		}

		if event.Rune() == 'j' {
			return tcell.NewEventKey(tcell.KeyDown, event.Rune(), event.Modifiers())
		}

		if event.Rune() == 'k' {
			return tcell.NewEventKey(tcell.KeyUp, event.Rune(), event.Modifiers())
		}

		return event
	})

	return menu

}
