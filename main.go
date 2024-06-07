package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName = "cli-reminder"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3{
		fmt.Printf("Input a time value and a message")
		os.Exit(1)
	}

	present := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(os.Args[1], present)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if r == nil {
		fmt.Printf("Unable to get time")
		os.Exit(2)
	}
	if present.After(r.Time){
		fmt.Printf("Your set time has passed")
		os.Exit(3)
	}

	diff := r.Time.Sub(present)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Printf("Reminder will be displayed in: %v\n", diff)
		os.Exit(0)
	}
}