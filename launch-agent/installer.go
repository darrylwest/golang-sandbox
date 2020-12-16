// main.go
package main

import (
    "fmt"
	"log"
	"os"
    "path/filepath"
	"text/template"
)

func Template() string {
	return `<?xml version='1.0' encoding='UTF-8'?>
 <!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd" >
 <plist version='1.0'>
   <dict>
     <key>Label</key><string>{{.Label}}</string>
     <key>Program</key><string>{{.Program}}</string>
     <key>StandardOutPath</key><string>/tmp/{{.Label}}.out.log</string>
     <key>StandardErrorPath</key><string>/tmp/{{.Label}}.err.log</string>
     <key>KeepAlive</key><{{.KeepAlive}}/>
     <key>RunAtLoad</key><{{.RunAtLoad}}/>
   </dict>
</plist>
`
}

func main() {
    home := os.Getenv("HOME")
	data := struct {
		Label     string
		Program   string
		KeepAlive bool
		RunAtLoad bool
	}{
		Label:     "net.rcs.ticker",
		Program:   filepath.Join(home, "bin/ticker"),
		KeepAlive: true,
		RunAtLoad: true,
	}

    log.Printf("%v\n", data)

    plistPath := filepath.Join(home, "/Library/LaunchAgents", fmt.Sprintf("%s.plist", data.Label))
    log.Printf("install agent config to %s\n", plistPath)

    file, err := os.Create(plistPath)
    if err != nil {
        log.Fatalf("error creating file: %s", err)
    }
    defer file.Close()

	t, err := template.New("plistconfig").Parse(Template())
	if err != nil {
		log.Fatalf("Template parse failed: %s", err)
    }

	err = t.Execute(file, data)
	if err != nil {
		log.Fatalf("Template generation failed: %s", err)
	}

    fmt.Println("installation complete")
}

// un-install?
