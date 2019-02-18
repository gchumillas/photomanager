# PhotoManager

An application to manage pictures from Dropbox.The application connects to a Dropbox account and allows the user to organize their images into categories. Take a look at the [API documentation](https://documenter.getpostman.com/view/412470/S11BzN24) so see the list of end-points.

## Requirements

[GO](https://golang.org/) is a modern programming language developed by Google and released in 2009. In order to compile and execute this application you must install the [GO compiler](https://golang.org/doc/install). Also this application uses a [MongoDB](https://www.mongodb.com/) database, so you must [install](https://docs.mongodb.com/manual/administration/install-community/) it in your computer.

## Create the database

To create the database collections simply execute the following command from a terminal:
```bash
mongo [ENTER YOUR DATABASE NAME] scripts/dbschema.js
```

## Configuration
Copy `config-example.json` to `config.json` and change the values. You can get the Dropbox keys from your [Dropbox Console Account](https://www.dropbox.com/developers/apps). Create a new project:

  1. **Choose an API**: select the `Dropbox API` option.
  2. **Choose the type of access you need**: select `App folder`.
  3. **Name your app**: type a name for your project.

Then copy the following values:

  * **App Key**: `dropboxAppKey`
  * **App secret**: `dropboxAppSecret`
  * **Redirect URIs**: `dropboxRedirectUri`

In the **Redirect URIs** section you must enter at least a URI pointing to the `login` handler. For example:
http://localhost:8080/v1/auth/login


## Install and start the server

Change to the project's folder and execute the next commands:
```bash
# install the application
go install

# start the server
photomanager
```

Make sure that `$GOPATH/bin` is included in your `$PATH` environment variable.

## Questions
