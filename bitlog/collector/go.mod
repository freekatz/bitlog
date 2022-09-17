module github.com/1uvu/bitlog/collector

go 1.18

replace (
	github.com/1uvu/bitlog => ../
)

require (
	github.com/1uvu/bitlog v0.0.0-00010101000000-000000000000
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/fsnotify/fsnotify v1.5.4
)

