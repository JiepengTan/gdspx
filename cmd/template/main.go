package main

import "C"
import (
	"godot-ext/gd4go/pkg/gd4go"
	"unsafe"

	"gd4go-demo/pkg/autogen"

	. "github.com/godot-go/godot-go/pkg/core"
	. "github.com/godot-go/godot-go/pkg/ffi"
	"github.com/godot-go/godot-go/pkg/log"
)

func main() {
	// log.Trace("this application is meant to be run as a plugin to godot")
}

//export GdExtentionEnterPoint
func GdExtentionEnterPoint(p_get_proc_address unsafe.Pointer, p_library unsafe.Pointer, r_initialization unsafe.Pointer) bool {
	log.Debug("GdExtentionEnterPoint called")
	initObj := NewInitObject(
		(GDExtensionInterfaceGetProcAddress)(p_get_proc_address),
		(GDExtensionClassLibraryPtr)(p_library),
		(*GDExtensionInitialization)(unsafe.Pointer(r_initialization)),
	)

	initObj.RegisterSceneInitializer(func() {
		gd4go.RegisterEngineTypes()
		autogen.RegisterGameTypes()
	})

	initObj.RegisterSceneTerminator(func() {
	})

	return initObj.Init()
}