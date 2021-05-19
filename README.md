# gofixtures
Ruby on Rails' style test fixtures for Golang, use database transation to keep database clean state

## Getting started

Install the package by:

```sh
go get -u github.com/Martin91/gofixtures
```

## Example
1. Reference to [dummy](./dummy) to see demo yaml files.
2. In your project, initalize database connection like:
    ```go
    import (
        _ "github.com/go-sql-driver/mysql"
        "github.com/Martin91/gofixtures"
    )

    // fixtures is dependent on https://github.com/DATA-DOG/go-txdb to rollback fixtures automatically,
    //  so it is required to setup a transational *sql.DB by fixtures
    db := fixtures.OpenDB("mysql", "root:@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local")
    fixtures := fixtures.Load(tt.args.path, db)

    // do your test and something else

    // once the program exit, fixtures will rollback all database changes automatically
    ```

## NOTICE
This repository is still under active development and the overall design and performance is unstable.

## Features (WIP, please keep looking forward to it)
[x] YAML based simple and clean syntax
[x] Built-in [Faker](https://github.com/bxcodec/faker/) supported
[x] Bundled field evaluators, enable you to customize dynamic data generation
[x] Support specifying database and tables
[x] Based on standard `sql` package, compatible with different dialects
[x] Transaction based database cleaner
[ ] Templates to support batch data

## TODOs
1. Test with different databases
2. Complete test cases with high test coverage
3. Review the overall architecture design
