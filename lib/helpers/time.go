package helpers

import (
	"fmt"
	"time"
)

// TimeTrack measure time function
func TimeTrack(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", functionName, elapsed)
}
