module github.com/lc/gau/v2

go 1.17

require (
	github.com/bobesa/go-domain-util v0.0.0-20190911083921-4033b5f7dd89
	github.com/valyala/bytebufferpool v1.0.0
	tinygo.org/x/drivers v0.23.0
)

require (
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
)

retract (
	v2.0.7
	v2.0.3
	v2.0.2
	v2.0.1
)
