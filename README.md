# What Is This?

<p align="center">
<img src="images/golang.png"></img>
</p>

At this project, I created a fully extendable golang template webservice structure in **Clean** architecture and came up with a reusable, nice and scalable template for any backend api project and ofcourse reusable codes and completely separate packages which helps programmers to just focus on implementing their own application. In this template project I tried to ease creating just an api just for myself and golang community.

If you do love to contribute, please do, I appreciate it.

# Features

1. [**pkg**](./pkg):

   - [**config**](pkg/config): Inside this package, I implemented a functionality which reads files and fills up passed config structure instances. You can just use output of this package which is inside `g.CFG` structure.
   - [**database**](pkg/database): Simply you just pass your data about your database connections to `New` function and it tries to create database connections and their query builders and return them all.
   - [**errors**](pkg/errors/): If you need to return an error in any where in your API project, panic by an error from `New` function in error package and give it an `status code`, an `action` and a `message`. this makes your error responses to users much more beautiful.
     - Example: `panic(errors.New(errors.InvalidStatus, errors.ReSignIn, "NotIncludedToken"))`
     - errors.InvalidStatus means 400 bad request status code
     - errors.ReSignIn means user has to signin again
   - [**logging**](pkg/logging/): Creates four folders inside `/var/log/project` like: `error`, `info`, `panic` and `warning` and if you use this logger, it will record those logs and put them inside their own folders in files.
     - Example: `g.Logger.Error(fmt.Sprintf("read: %s", err), FunctionWeAreIn, OptionalMap)`
     - First argument is your error message
     - Second argument, as it's name says `FunctionWeAreIn`, you just pass the function you got your error in them
     - Third argument, is just a map that if you want to provide new information, you can provide your other data in them
   - [**translator**](pkg/translator/): This is the package which reads translations inside your translation folder in build time inside `build/translations` and you can use this tool to translate.
     - Example: `g.Translator.TranslateFunction("en")("translate me")`
     - By default in Translator middleware, a function gets added to context of the request and can be retrieved and used.

2. [**internal**](./internal/):
   - api
   - multi database support
   - jwt authentication
   - users service(sign in, sign up and create users)
   - flexible configuration files
   - not being dependable on which database you use
   - using sql-migrate tool for migrations
   - using net/http as http response handler for api

# Quick Start

There are two ways:

1. You can just run `auto.py` script like:

   ```
   python3 auto.py setup
   ```

2. Or you can do it all by yourself:

   1. Install Dependencies:

      - ```
        go mod download -x all
        ```
      - **Note**: For migrations you most likely need [sql-migrate].

   2. Copy and paste these lines in your terminal when you're inside project root directory:

      - ```bash
         cp dbconfig_example.yml dbconfig.yml
         cp env_example.yml env.yml
        ```
      - Those example files(env_example.yml and dbconfig_example.yml) have ready configurations for a quick start for the project.

   3. You can run `.githooks/install.py` script to activate custom githooks inside `.githooks` folder if you want.

3. How to run:
   - `go run main.go`

[sql-migrate]: https://github.com/rubenv/sql-migrate
