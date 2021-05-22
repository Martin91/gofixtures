module github.com/Martin91/gofixtures

go 1.15

replace github.com/DATA-DOG/go-txdb v0.1.4 => github.com/martin91/go-txdb v0.0.0-20210522103453-a6eaffbdb2e9

require (
	github.com/DATA-DOG/go-txdb v0.1.4
	github.com/bxcodec/faker/v3 v3.6.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lib/pq v1.10.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v2 v2.4.0
)
