// +build linux,cgo windows,cgo,darwin cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofutuapi

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/libcpp/Bin/Ubuntu16.04 -Wl,-rpath,${SRCDIR}/libcpp/Bin/Ubuntu16.04 -lFTAPIChannel -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/libcpp/include
#cgo darwin LDFLAGS: -fPIC -L${SRCDIR}/libcpp/Bin/Mac/Release -Wl,-rpath,${SRCDIR}/libcpp/Bin/Mac/Release -lFTAPIChannel
#cgo darwin CPPFLAGS: -fPIC -I${SRCDIR}/libcpp/include
*/
import "C"
