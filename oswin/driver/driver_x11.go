// Copyright 2018 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// based on golang.org/x/exp/shiny:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,!android dragonfly openbsd

package driver

import (
	"github.com/goki/gi/oswin"
	//	"github.com/goki/gi/oswin/driver/gldriver"
	"github.com/goki/gi/oswin/driver/x11driver"
)

func main(f func(oswin.App)) {
	//     	gldriver.Main(f)
	x11driver.Main(f)
}
