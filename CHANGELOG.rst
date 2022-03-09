CHANGELOG
=========

UNRELEASED
----------

* 🐛 fix: default config file is hardcoded in `internal/app/load/cfg.go` cause other solutions are not acceptable

1.1.1 (2022-03-08)
------------------

* fix: g.Translator changed to g.Trans()

1.1.0 (2022-03-08)
------------------

* fix: install.py script in .githooks folder replaced husky
* feat: with no app option, grpc and fiber will run together
* feat: if debug is true, database with `test` name will place in g.DB and if false, database with `main` name will place in g.DB

1.0.2 (2022-02-25)
------------------

* fix: bug fixed on test cases not getting succeed when running at the same time
* fix: CreateUser bug fixed on returning a user object without their id
* test: test cases added for token and user service

1.0.1 (2022-02-24)
------------------

* fix: test files address in pre-push script fixed
* fix: all contract files are in internal/contract as they should be
* docs: documentation on how to setup updated
* refactor: some config example files added to root directory
* refactor: test cases have a separate folder in root directory
* docs: documentation updated on how to run Quick Start section

1.0.0 (2022-02-23)
------------------

* feat: config package added
* feat: logging package added
* feat: errors package added
* feat: translator package added
* feat: database package added
* feat: grpc package added
* feat: fiber framework added as an http, socket and middleware functionality provider
* feat: sign in and sign up feature with jwt and sessions added
* feat: grpc support added
* feat: goqu query builder added
* feat: multi database feature added
* feat: sql-migrate command line for migration automation added
* feat: csrf service for preventing csrf attacks added