package main

import (
  "net"
  "github.com/codegangsta/cli"
  "os"
  "strconv"
)

func main() {

  app := cli.NewApp()
  app.Name = "Kickback-Agent"
  app.Usage = "Agent program to Kichback Server"
  app.Flags = []cli.Flag {
    cli.StringFlag {
      Name: "ip",
      Value: "NULL",
      Usage: "Destination IP Address",
    },
    cli.IntFlag {
      Name: "port",
      Value: 0,
      Usage: "Destination port",
    },
  }
  app.Action = func(c *cli.Context) {
    ip := ""
    port := 0
    ip = c.String("ip")
    port = c.Int("port")
    if ip == "" || port == 0 {
      println("Usage: kickback_agent --ip=ip --port=port")
    } else {
      println(ip, port)
      if dial(ip, port) {
        println("Test succeeded")
      } else {
        println("Test failed")
      }
    }
  }
  app.Run(os.Args)
}

func dial(ip string, port int) bool {
  conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
  if err != nil {
    return false
  } else {
    conn.Close()
    return true
  }
}
