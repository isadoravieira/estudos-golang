package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {

	app := cli.NewApp()

	app.Name = "Command line application"
	app.Usage = "Searching for IPs and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "IP",
			Usage: "Searching for IP addresses on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com/isadoravieira",
				},
			},
			Action: searchingIPS,
		},
		{
			Name:  "servers",
			Usage: "Search the name of the servers on the internet.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com/isadoravieira",
				},
			},
			Action: searchingServers,
		},
	}

	return app
}

func searchingIPS(ctx *cli.Context) {
	host := ctx.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchingServers(ctx *cli.Context) {
	host := ctx.String("host")

	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
