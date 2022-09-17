module bitlog_cli

go 1.18

replace (
	github.com/1uvu/bitlog => ./../../bitlog
	github.com/1uvu/bitlog/collector => ./../../bitlog/collector
	github.com/1uvu/bitlog/parser => ./../../bitlog/parser
	github.com/1uvu/bitlog/storage => ./../../bitlog/storage
)

require (
	github.com/1uvu/bitlog v0.0.0-00010101000000-000000000000
	github.com/1uvu/bitlog/collector v0.0.0-00010101000000-000000000000
	github.com/1uvu/bitlog/parser v0.0.0-00010101000000-000000000000
	github.com/1uvu/bitlog/storage v0.0.0-00010101000000-000000000000
)