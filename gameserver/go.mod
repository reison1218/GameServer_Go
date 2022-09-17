module gameserver

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gomodule/redigo v1.8.5
	google.golang.org/protobuf v1.27.1
	tools v0.0.0
)

require (
	github.com/json-iterator/go v1.1.11
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1
)

replace tools => ../tools
