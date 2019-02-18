# PhotoManager

An application to manage pictures from Dropbox.

The application connects to a Dropbox account and allows the user to organize their images into categories.

## Requirements

[GO](https://golang.org/) is a modern programming language developed by Google and released in 2009. In order to compile and execute this application you must install the [GO compiler](https://golang.org/doc/install).

Also this application uses a [MongoDB](https://www.mongodb.com/) database, so you must [install](https://docs.mongodb.com/manual/administration/install-community/) it in your computer.

## Create the database

To create the database collections simply execute the following command from a terminal:
```bash
mongo [ENTER YOUR DATABASE NAME] scripts/dbschema.js
```

## Configuration
Copy `config-example.json` to `config.json` and change the values.

## Install and start the server

Change to the project's folder and execute the next commands:
```bash
# install the application
go install

# start the server
photomanager
```

Make sure that `$GOPATH/bin` is included in your `$PATH` environment variable.
