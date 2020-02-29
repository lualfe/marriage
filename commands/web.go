package commands

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lualfe/casamento/app"
	"github.com/lualfe/casamento/web"
	"github.com/spf13/cobra"
)

var webCommand = &cobra.Command{
	Use: "web",
	RunE: func(cmd *cobra.Command, args []string) error {
		e := echo.New()
		a, err := app.InitApp()
		foreignKey, _ := cmd.Flags().GetString("foreignkey")
		switch foreignKey {
		case "bind":
			a.Cockroach.Migrate()
			a.Cockroach.BindForeignKeys()
		default:
			a.Cockroach.Migrate()
		}
		defer a.Cockroach.Instance.Close()
		if err != nil {
			e.Logger.Printf("error initializing database: %v", err)
			panic(err)
		}
		w, err := web.New(a)
		if err != nil {
			e.Logger.Printf("error initializing database: %v", err)
			panic(err)
		}
		e.Use(middleware.CORS())
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "<-- Method: ${method} | URI: ${uri} | Status: ${status} | Latency: ${latency_human} -->\n",
		}))
		w.InitRoutes(e)
		e.Logger.Fatal(e.Start(":1323"))
		return nil
	},
}

func init() {
	cmd.AddCommand(webCommand)
	webCommand.Flags().StringP("foreignkey", "k", "", "Add or remove foreign keys")
}
