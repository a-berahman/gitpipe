package cmd

import (
	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/utility/rest"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the application",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

// serve handles the serve command
func serve() {

	e := echo.New()
	e.HideBanner = true
	p := prometheus.NewPrometheus("gitpipe", nil)
	p.Use(e)

	config.LoadConfig(configPath)

	rest.Initialize()

	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	rootCMD.AddCommand(serveCmd)
}
