package main

/**
 * this watcher uses polling rather than the standard fsnotify...
 */

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
    logo = `
             _      _            
__ __ ____ _| |_ __| |_  ___ _ _ 
\ V  V / _' |  _/ _| ' \/ -_) '_|
 \_/\_/\__,_|\__\__|_||_\___|_|  
`
    VERSION = "2018.02.26"
)

type FileStat struct {
	path    string
	size    int64
	updated time.Time
	hash    string
}

type FileMap map[string]FileStat

type WatcherContext struct {
	filemap    FileMap
	root       string
	extensions []string
	wait       time.Duration
	action     string
}


var (
	ctx     *WatcherContext
	verbose = false
	files   []string
)

func createStat(f os.FileInfo) *FileStat {
	stat := new(FileStat)

	stat.path = f.Name()
	stat.size = f.Size()
	stat.updated = f.ModTime()

	return stat
}

func (w *WatcherContext) isInExtensions(path string) bool {
	for _, ext := range w.extensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}

	return false
}

// walk the tree and get the stats for all files
func (w *WatcherContext) walkpath(path string, f os.FileInfo, err error) error {
	if f.IsDir() == false && w.isInExtensions(path) {
		// try to find in the hashmap

		if cachedStat, ok := ctx.filemap[path]; ok {
			if cachedStat.size != f.Size() || cachedStat.updated != f.ModTime() {
				files = append(files, path)
				ctx.filemap[path] = (*createStat(f))
			}
		} else {
			// dont queue new files; wait for a change first
			stat := createStat(f)

			fmt.Printf("watching %s %d...\n", stat.path, stat.size)
			ctx.filemap[path] = (*stat)
		}
	}

	return nil
}

func (w *WatcherContext) process(jobs []string) {
	for _, f := range jobs {
		fmt.Printf("process %s\n", f)
		if w.action != "" {
			cmd := exec.Command(w.action, f)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Start(); err != nil {
				fmt.Println(err)
			} else {
				if err := cmd.Wait(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func NewDefaultWatcher() *WatcherContext {
	ctx = new(WatcherContext)

	ctx.extensions = []string{".go"}
	ctx.wait = 750 * time.Millisecond
	ctx.root = "./"
	ctx.action = "./run.sh"

	ctx.filemap = make(FileMap)

	return ctx
}

func parseArgs() *WatcherContext {
	dflt := NewDefaultWatcher()

    suffixes := flag.String("ext", strings.Join(dflt.extensions, ","), "a comma-separated list of extension suffixes to watch, e.g., '.java,.go'")
    // delay := flag.Int("delay", int(dflt.wait), "the number of millisectonds to wait between file scans")
    dir := flag.String("dir", dflt.root, "the folder to watch")
    action := flag.String("action", dflt.action, "the default command to run agains the watched file when it changes")

	flag.Parse()

    if err := os.Chdir(*dir); err != nil {
        panic(err)
    }

    ctx := WatcherContext{
        extensions: strings.Split(*suffixes, ","),
        wait: dflt.wait,
        root: dflt.root,
        action: *action,
        filemap: dflt.filemap,
    }

	return &ctx
}

func main() {
    fmt.Println(logo)
    fmt.Printf("Version: %s\n", VERSION)

	ctx := parseArgs()

	for {
		filepath.Walk(ctx.root, ctx.walkpath)

		if len(files) > 0 {
			ctx.process(files)
			files = make([]string, 0)
		}

		time.Sleep(ctx.wait)
	}
}
