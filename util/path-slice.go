package main

import (
	"bytes"
	"fmt"
)

type path []byte

// modifies the slice in place
func (p *path) TruncateAtFinalSlashInPlace() {
	idx := bytes.LastIndex(*p, []byte("/"))

	if idx >= 0 {
		*p = (*p)[:idx]
	}
}

// creates a modified version of the slice
func (p path) TruncateAtFinalSlash() path {
	idx := bytes.LastIndex(p, []byte("/"))

	if idx >= 0 {
		p = p[:idx]
	}

	return p
}

// modifies the array in place without a pointer
func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

// modifies the array in place with a pointer
func (p *path) ToLower() {
	for i, b := range *p {
		if 'A' <= b && b <= 'Z' {
			(*p)[i] = b + 'a' - 'A'
		}
	}
}

func main() {
	pathName := path("/usr/bin/tso")

	// leaves the pathName intact
	pn := pathName.TruncateAtFinalSlash()

	fmt.Printf("%s %s\n", pathName, pn)

	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)

	pathName.ToLower()
	fmt.Printf("%s\n", pathName)

	// modifies pathName
	pathName.TruncateAtFinalSlashInPlace()
	fmt.Printf("%s\n", pathName)

}
