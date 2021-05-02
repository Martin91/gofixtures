# gofixtures
Ruby on Rails' style test fixtures for Golang

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
    db, err := fixtures.OpenDB("mysql", "root:@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }

    fixtures, err := fixtures.Load(tt.args.path, db)
    if err != nil {
        panic(err)
    }

    // do your test and something else

    // once the program exit or db.Close() is called, fixtures will rollback all database changes
    ```

## NOTICE
This repository is still under active development and the overall design and performance is unstable.

## Features (WIP, please keep looking forward to it)
1. YAML based simple and clean syntax
2. Built-in [Faker](https://github.com/bxcodec/faker/) supported
3. Bundled field evaluators, enable you to custome dynamic data generation
4. Support specifying database and tables
5. Based on standard `sql` package, compatible with different dialects
6. Transaction based database cleaner