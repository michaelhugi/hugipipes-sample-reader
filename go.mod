module hugipipes-sample

go 1.18

require (
	github.com/eripe970/go-dsp-utils v0.0.0-20220123162022-4563116e558a
	github.com/informaticon/lib.go.base.data-types v0.0.13
	github.com/informaticon/lib.go.base.test-utils v0.0.11
	github.com/mattetti/audio v0.0.0-20190404201502-c6aebeb78429
	github.com/michaelhugi/go-hugipipes-musical-notes v1.0.2
	github.com/michaelhugi/go-hugipipes-signal-drawer v0.0.2
	github.com/mjibson/go-dsp v0.0.0-20180508042940-11479a337f12
	golang.org/x/image v0.0.0-20220321031419-a8550c1d254a
)

replace (
	github.com/michaelhugi/go-hugipipes-signal-drawer => C:\workspace\Go\go-hugipipes-signal-drawer
)
require (
	github.com/goccmack/godsp v0.1.1 // indirect
	github.com/goccmack/goutil v0.4.0 // indirect
)
