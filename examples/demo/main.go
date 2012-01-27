package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	var MaxFPS int
	var demoName string
	var listDemos bool
	var exit bool
	console := NewConsole()
	go console.Run()

	flags := flag.NewFlagSet("demo", flag.ContinueOnError)
	flags.StringVar(&demoName, "load", "", "Loads a demo.")
	flags.BoolVar(&listDemos, "list", false, "Lists all demos.")
	flags.BoolVar(&exit, "exit", false, "Exists the demo.")
	flags.IntVar(&MaxFPS, "fps", 60, "Maximum frames per second.")
	flags.PrintDefaults()

	// Demos
	simpleWindowDemo := NewSimpleWindowDemo()
	primitivesDemo := NewPrimitivesDemo()

	demoManager := NewDemoManager()
	defer demoManager.Unload()
	demoManager.AddDemos(simpleWindowDemo, primitivesDemo)

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ticker := time.NewTicker(time.Second / time.Duration(MaxFPS))
	for !exit {
		select {
		case line := <-console.NewLine:
			if len(line) == 0 {
				continue
			}
			if line[0] != '-' {
				flags.PrintDefaults()
				continue
			}
			args := strings.Split(line, " ")
			err := flags.Parse(args)
			if exit || err != nil {
				continue
			}
			if listDemos {
				fmt.Println("Demos:")
				demos := demoManager.Demos()
				for i := range demos {
					fmt.Printf("%s - %s\n", demos[i].Name(), demos[i].Description())
				}
				listDemos = false
				continue
			}
			if len(demoName) > 0 {
				err := demoManager.Load(demoName)
				if err != nil {
					demoManager.Unload()
					fmt.Println(err)
				}
				demoName = ""
				continue
			}
		case <-ticker.C:
			demoManager.Update()
		}
	}
}

type Console struct {
	NewLine chan string
	stop    chan bool
	reader  *bufio.Reader
}

func NewConsole() *Console {
	return &Console{NewLine: make(chan string), stop: make(chan bool)}
}

func (c *Console) Run() {
	c.reader = bufio.NewReader(os.Stdin)
	for {
		select {
		case stop := <-c.stop:
			if stop {
				return
			}
		default:
			line, _, err := c.reader.ReadLine()
			if err != nil {
				fmt.Printf("Console error: %s\n", err)
				continue
			}
			c.NewLine <- string(line)
		}
	}
}

func (c *Console) Close() {
	c.stop <- true
}
