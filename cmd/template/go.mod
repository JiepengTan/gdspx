module gd4go-demo

go 1.22.3

toolchain go1.22.5

require github.com/godot-go/godot-go v0.0.0-00010101000000-000000000000

require godot-ext/gd4go v0.0.0-00010101000000-000000000000

require (
	github.com/CannibalVox/cgoalloc v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/exp v0.0.0-20230725093048-515e97ebf090 // indirect
	golang.org/x/text v0.15.0 // indirect
)

replace godot-ext/gd4go => ../../

replace github.com/godot-go/godot-go => ../../../godot-go