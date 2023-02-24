package app

import (
	"fmt"
	"log"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/pkg/colors"
)

func info() {
	fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sSystem Info%s==%s\n", colors.Yellow, colors.Cyan, colors.Reset))
	mainOrTest := "test"
	mainOrTestColor := colors.Red + mainOrTest + colors.Reset
	if !g.CFG.Debug {
		mainOrTest = "main"
		mainOrTestColor = colors.Green + mainOrTest + colors.Reset
	}
	for _, database := range g.CFG.Databases {
		if database.Name == mainOrTest {
			if database.Type == "sqlite3" {
				fmt.Printf("Main Database:\t\t%v, %v (%v)\n", database.Type, database.DBName, mainOrTestColor)
			} else {
				fmt.Printf("Main Database:\t\t%v, %v, %v:%v (%v)\n", database.Type, database.DBName, database.Host, database.Port, mainOrTestColor)
			}
			if g.DB == nil {
				log.Fatal("expected database connection is not assigned as main database")
			}
			break
		}
	}
	if g.CFG.Debug {
		fmt.Printf("Debug:\t\t\t%s%v%s\n", colors.Red, g.CFG.Debug, colors.Reset)
	} else {
		fmt.Printf("Debug:\t\t\t%s%v%s\n", colors.Green, g.CFG.Debug, colors.Reset)
	}
	fmt.Printf("Address:\t\thttp://%s:%s\n", g.CFG.Api.IP, g.CFG.Api.Port)
	fmt.Printf("Allowed Origins:\t%v\n", g.CFG.AllowOrigins)
	if g.CFG.AllowHeaders != "" {
		fmt.Printf("Extra Allowed Headers:\t%v\n", g.CFG.AllowHeaders)
	}
	fmt.Print(colors.Cyan, "===============\n\n", colors.Reset)
}
