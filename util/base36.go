//
// convert the current nano second time to a base 36
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-14 17:50:41
//

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now().UTC()
	s := strconv.FormatInt(now.UnixNano(), 36)

	fmt.Println(s, len(s))
}
