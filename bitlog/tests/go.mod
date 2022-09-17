module github.com/1uvu/bitlog/tests

go 1.18

replace (
	github.com/1uvu/bitlog => ./../
	github.com/1uvu/bitlog/collector => ./../collector
)

require (
	github.com/1uvu/bitlog v0.0.0-00010101000000-000000000000
	github.com/1uvu/bitlog/collector v0.0.0-00010101000000-000000000000
	github.com/fsnotify/fsnotify v1.5.1 // indirect
)
