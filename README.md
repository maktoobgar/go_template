# What Is This?

<p align="center">
<img src="images/golang.png"></img>
</p>

At this project, I created a fully extendable golang template webservice structure in **Clean** architecture and came up with a reusable, nice and scalable template for any backend project and ofcourse reusable codes and completely separate packages which helps programmers to just focus on implementing their own application. In this template project I tried to imagine different scenarios and test everything for open source community and ofcourse for my future projects.

If you do love to contribute, please do, I appreciate it.

And ofcourse don't forget to read [CONTRIBUTING](/CONTRIBUTING.md) file to know about how to contribute in this project.

## Features

1. [**pkg**](./pkg):
   * [**config**](pkg/config): Inside this package, I implemented a functionality which reads files and fills up passed config structure instances. You can just use output of this package which is inside `g.CFG` structure.
   * [**database**](pkg/database): Simply you just pass your data about your database connections to `New` function and it tries to create database connections and their query builders and return them all.
   * [**errors**](pkg/errors/): If you need to return an error in any where in your `fiber` project (**socket** and **api**), use `New` function and give it an `status code`, an `action` and a `message` and return that error that function gives to you, this makes your error responses to users much more beautiful.
     * Example: `errors.New(errors.InvalidStatus, errors.ReSingIn, g.Trans().TranslateEN("NotIncludedToken"))`
     * errors.InvalidStatus means 400 bad request status code
     * errors.ReSingIn means user has to signin again
     * g.Trans().TranslateEN("NotIncludedToken") tries to translate "NotIncludedToken" and returns the translated value(if translation fails, no error returns and just "NotIncludedToken" returns)
   * [**grpc**](pkg/grpc/): Just simplifies setup for running grpc server.
   * [**logging**](pkg/logging/): Creates four folders inside `/var/log/project` like: `error`, `info`, `panic` and `warning` and if you use this logger, it will record those logs and put them inside their own folders in files.
     * Example: `g.Log().Error(fmt.Sprintf("read: %s", err), FunctionWeAreIn, OptionalMap)`
     * First argument is your error message
     * Second argument, as it's name says `FunctionWeAreIn`, you just pass the function you got your error in them
     * Third argument, is just a map that if you want to provide new information, you can provide your other data in them
   * [**translator**](pkg/translator/): This is the package which reads translations inside your translation folder which by default your translation folder is defined in default config file(which is in `build/config/config.yaml`) inside `build/translations` and you can use this tool to translate.
     * Example: `g.Trans().TranslateEN("RequiresNotProvided")`
     * Example: `g.Trans().Translate(language.English.String(), "RequiresNotProvided")`

2. [**internal**](./internal/):
   * api
   * socket
   * grpc
   * multi database support
   * jwt authentication
   * session authentication
   * csrf
   * users service(sign in, sign up and create users)
   * flexible configuration files
   * not being dependable on which database you use
   * using sql-migrate tool for migrations
   * using fiber as http response handler for api and socket connections
   * running fiber (api, socket) and grpc with a single command
     * **Note**: grpc and fiber can't run on the same port

## Quick Start

1. Install Dependencies:
   * ```
     go mod download -x all
     go get -v github.com/rubenv/sql-migrate/... && git restore go.mod
     ```
   * **Note**: For migrations you most likely need `sql-migrate`.

2. Copy and paste these lines in your terminal when you're inside project root directory:
   * ```bash
      cp dbconfig_example.yml dbconfig.yml
      cp env_example.yml env.yml
      ```
   * Those example files(env_example.yml and dbconfig_example.yml) have ready configurations for a quick start for the project.

3. How to run:
   1. If you want to run **socket**, **api** and **grpc** all in one action. run `cmd/main/main.go` file like:
       * ```bash
          go run ./cmd/main/main.go
          # or
          go run ./cmd/main/main.go -app=0
          # or
          go run ./cmd/main/main.go -app=b
          # or
          go run ./cmd/main/main.go -app=both
          ```
   2. If you want to run **socket** and **api**, run `cmd/main/main.go` file like:
       * ```bash
          go run ./cmd/main/main.go -app=1
          # or
          go run ./cmd/main/main.go -app=f
          # or
          go run ./cmd/main/main.go -app=fiber
          ```

   3. If you want to run **grpc** server, run `cmd/main/main.go` file like:
      * ```bash
        go run ./cmd/main/main.go -app=2
        # or
        go run ./cmd/main/main.go -app=g
        # or
        go run ./cmd/main/main.go -app=grpc
        ```

## Clean Architecture

If you don't have any idea about **clean** architecture, please just take a moment and just have a look at it first and after understanding the architecture, come back here and continue.

**Note**: If you first know about clean architecture, understanding folders usages will make more sense.

## Config Loading

The [`cfg.go`][cfg.go] file (if the [config files](#config-files) are passed just right) fills attributes in [`global.go`][global.go].

Just to load everything right up, you have to import [`cfg.go`][cfg.go] file like:
```
_ "github.com/maktoobgar/go_template/internal/app/load"
```

**Note**: You don't need to do that, I'm just describing what is happening.

## Config Files

There is a file named [config.go][config.go] inside `internal/config` folder which **defines structures** for config files of the project.\
**Main config files**, which have some default configs in them and no coding happens in them, are inside `build` folder and the default config file named [`config.yaml`][config.yaml] is inside `build/config` address.\
Take a look at them and you get whats going on.

## What Is `g`? (important)

`g` stands for **global**. In this project, I created a file named [`global.go`][global.go] inside `internal/global` folder to have all config files for the project in one place and access them as fast as possible.\
In down documentation you may see me mention using `g.Something`, just know that I meant [global.go][global.go] file and I assume you know that you have to import it like `g "github.com/maktoobgar/go_template/internal/global"` to use `g`.

### Update Default Configs

If you want to update default config file([`config.yaml`][config.yaml] file), you may add a file to the root project directory named `env.yaml` or `env.yml` and override any configs you need to and replace you new desired configs.

## Translation

If you need to translate anything in your application, use `g.Translator` object to do it and don't forget to add your translation to `build/translations` folder like other files that you can use as an example in that directory.

**Note**: If you want to be so contractual, you can have interface of `g.Translator` by calling `g.Trans()` function and then do your translations like:
```
g.Trans().TranslateEN("my word")
```

### New Languages

If you need new languages(Persian and English added by default), you can add them by first updating `languages` variable inside [cfg.go][cfg.go] file and then at least create an empty file for that variable inside `build/translations` folder like [translation.fa.json][translation.fa.json] file.

## Query Builder

In this project I added [`goqu`][goqu] query builder, if you need to learn how to use it, read [goqu][goqu] example files inside their repo. (files followed by `_example_test.go` at the end are example files)\
If you take a look at [config.go][config.go] file, you can see that in `Database` structure, there is a `Name` attribute inside them.\
That `Name` attribute will be used to give you an access to your databases, after config files successfully loaded, you can use `Name` attribute of your database to access their query builder to make changes on your databases just by calling `g.Postgres[Name]` or `g.Sqlite[Name]` or `g.MySQL[Name]` or `g.SqlServer[Name]`.
Please attention that a database connection with their `Name` equal to `main`, will be places inside `g.DB` separately from others just for fast access to your main database connection.

**Note**: Connections to these databases are inside `g.PostgresCons`, `g.SqliteCons`, `g.MySQLCons` and `g.SqlServerCons`.

## Fiber Framework

The framework that I'm using in this project is **fiber** and the reasons are simple:
1. fast
2. flexible
3. easy to use

Take a look at [fiber][fiber] documentations.

## Migration

For migrations I'm using [sql-migrate][sql-migrate] tool.\
this tool uses [`dbconfig.yml`](dbconfig.yml) file to create new migration files, migrate new migrations and rollback migrations but you have to create [`dbconfig.yml`](dbconfig.yml) file based on [`dbconfig_example.yml`](dbconfig_example.yml) file to use `sql-migrate` properly.

Read documentations for more information about how to use this tool.

## Adding Service

The services path is `internal/services` and inside that you can create a folder, name it whatever you want and add a file in it and start writing you service. There are examples in `internal/services` folder that you can have a look at them.

**Note**: I strongly recommend you to add a service for every tables you have in your database.

**Note**: If you want to follow up some more recommendations, create your interfaces or your popular structures(you use them somewhere else) for your services inside `internal/contract` folder.

## Adding Handler

Adding handlers:
1. If you are adding a http handler(api handler), use `internal/handlers/http` folder and create your handler inside it.
2. If you are adding a socket handler, use `internal/handlers/socket` folder and create your handler inside it.
3. If you are adding a grpc handler, use `internal/handlers/grpc` folder and create your handler inside it.

### How To Register Handler

1. api: Inside `internal/routers/http.go` file, create a group of your api url and add your handler with a Get or Post or... method.
2. socket: Inside `internal/routers/ws.go` file, create a group of your api url and add your handler with a Get or Post or... method.
3. grpc: Inside `internal/app/grpc.go` file, register your new grpc method. (I can't really explain this too much cause I'm new in this too but just have a look at `internal/handlers/grpc/hello` handler example)

## Adding Middleware

Just take a look at middleware examples inside `internal/middleware` folder and create your middleware in them and you are good to go.

### How To Register Middlewares

Inside `internal/routers/http.go` file for http methods or `internal/routers/ws.go` file for socket methods, create a group of your url address and add your middleware inside it.

[config.yaml]: build/config/config.yaml
[cfg.go]: internal/app/load/cfg.go
[global.go]: internal/global/global.go
[config.go]: internal/config/config.go
[translation.fa.json]: build/translations/translation.fa.json
[goqu]: https://github.com/doug-martin/goqu
[fiber]: https://docs.gofiber.io/
[sql-migrate]: https://github.com/rubenv/sql-migrate
