module github.com/chremoas/auth-srv

go 1.12

require (
	github.com/chremoas/role-srv v1.2.2
	github.com/chremoas/services-common v1.2.3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-micro v1.8.1
	github.com/petergtz/pegomock v2.5.0+incompatible
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
)

replace github.com/chremoas/services-common => /home/wonko/gohack/github.com/chremoas/services-common
