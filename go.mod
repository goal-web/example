module github.com/goal-web/example

go 1.17

require (
	github.com/goal-web/application v0.1.4
	github.com/goal-web/auth v0.1.5
	github.com/goal-web/bloomfilter v0.1.1
	github.com/goal-web/cache v0.1.0
	github.com/goal-web/config v0.1.1
	github.com/goal-web/console v0.1.1
	github.com/goal-web/container v0.1.5
	github.com/goal-web/contracts v0.1.47
	github.com/goal-web/database v0.1.7
	github.com/goal-web/email v0.1.1
	github.com/goal-web/encryption v0.1.1
	github.com/goal-web/events v0.1.5
	github.com/goal-web/filesystem v0.0.0-20220120135714-dc3dd88b880a
	github.com/goal-web/hashing v0.1.0
	github.com/goal-web/http v0.1.6
	github.com/goal-web/querybuilder v0.1.13
	github.com/goal-web/queue v0.1.1
	github.com/goal-web/ratelimiter v0.1.2
	github.com/goal-web/redis v0.1.3
	github.com/goal-web/serialization v0.1.8
	github.com/goal-web/session v0.1.4
	github.com/goal-web/supports v0.1.18
	github.com/goal-web/validation v0.1.0
	github.com/goal-web/websocket v0.1.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-module/carbon/v2 v2.0.1
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/labstack/echo/v4 v4.6.3
	github.com/qbhy/parallel v1.4.0
	github.com/stretchr/testify v1.7.0
	github.com/xwb1989/sqlparser v0.0.0-20180606152119-120387863bf2
)

require (
	github.com/ClickHouse/clickhouse-go/v2 v2.0.9 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/apex/log v1.9.0 // indirect
	github.com/bits-and-blooms/bitset v1.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-redis/redis/v8 v8.11.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goal-web/collection v0.1.5 // indirect
	github.com/goal-web/pipeline v0.1.6 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20220104163920-15ed2e8cf2bd // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jmoiron/sqlx v1.3.4 // indirect
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	github.com/modood/table v0.0.0-20200225102042-88de94bb9876 // indirect
	github.com/nsqio/go-nsq v1.1.0 // indirect
	github.com/paulmach/orb v0.4.0 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/segmentio/kafka-go v0.4.27 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opentelemetry.io/otel v1.3.0 // indirect
	go.opentelemetry.io/otel/trace v1.3.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20220128200615-198e4374d7ed // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/goal-web/application => ../application
	github.com/goal-web/auth => ../auth
	github.com/goal-web/bloomfilter => ../bloomfilter
	github.com/goal-web/cache => ../cache
	github.com/goal-web/collection => ../collection
	github.com/goal-web/config => ../config
	github.com/goal-web/console => ../console
	github.com/goal-web/container => ../container
	github.com/goal-web/contracts => ../contracts
	github.com/goal-web/database => ../database
	github.com/goal-web/email => ../email
	github.com/goal-web/encryption => ../encryption
	github.com/goal-web/http => ../http
	github.com/goal-web/queue => ../queue
	github.com/goal-web/redis => ../redis
	github.com/goal-web/serialization => ../serialization
	github.com/goal-web/session => ../session
	github.com/goal-web/supports => ../supports
	github.com/goal-web/validation => ../validation
	github.com/goal-web/websocket => ../websocket

)
