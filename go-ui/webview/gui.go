package main

import (
    "fmt"
    "os"
    "github.com/zserge/webview"
)

const Version = "2018.12.16"

func main() {
    sites := map[string]string{
        "map":"https://maps.google.com",
        "raw":"https://rawstory.com",
        "you":"https://youtube.com",
        "crook":"https://crooksandliars.com",
        "news":"https://www.reuters.tv",
        "post":"https://www.washingtonpost.com",
        "vox":"https://www.vox.com",
        "gate":"https://www.sfgate.com/bayarea",
        "blav":"https://blavity.com",
        "nyt":"https://www.nytimes.com",
        "flick":"https://www.flickr.com/photos/darrylwest",
        "time":"https://www.worldtimeserver.com",
        "pan":"https://www.pandora.com",
        "wet":"https://www.wunderground.com/weather/us/ca/metro-oakland-international",
        "gce":"https://console.cloud.google.com/home/dashboard?project=api-project-903136704055",
    }

    args := os.Args[1:]
    url := "https://www.google.com"

    if len(args) > 0 {
        if site, ok := sites[args[0]]; ok {
            url = site
        } else if args[0] == "help" {
            fmt.Printf("Version: %s\n", Version)

            for k, v := range sites {
                fmt.Printf("  %s\t%s\n", k, v)
            }

            os.Exit(0)
        } else {
            url = args[0]
        }
    }
    
    webview.Open("Minimal webview example", url, 1200, 800, true)
}
