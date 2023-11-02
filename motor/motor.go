package motor

import (
	"math/rand"
	"time"
)

var (
	source        = rand.NewSource(time.Now().UnixNano())
	randGenerator = rand.New(source)
)
