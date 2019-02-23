# PhotoManager (in progress)

Photomanager is a REST application aimed primarily at photographers who want to organize their photos into categories, so that the same image can belong to one or more categories.

The user must have a [Dropbox](https://www.dropbox.com) account, so that the application can manage their images. The user can upload an image and then assign it to one or more categories. So simple!

Please take a look at the [API documentation](https://documenter.getpostman.com/view/412470/S11BzN24) to learn more.

## Requirements

1. [MongoDB](https://www.mongodb.com/): this application uses a [NoSQL](https://en.wikipedia.org/wiki/NoSQL) database.
2. [GO compiler](https://golang.org/doc/install): this application was written in [GO](https://golang.org/), a modern language developed by Google and realeased in 2009.

## Create the database

Open a terminal and execute the following command from the application directory:
```bash
$ mongo [ENTER YOUR DATABASE NAME] scripts/dbschema.js
```

## Configuration
Copy `config-example.toml` to `config.toml` and change the key values. To configure the `[dropbox]` section simply follow these steps:

1. [Sign in](https://www.dropbox.com/login) into your Dropbox account.
2. Open the [Dropbox Application Console](https://www.dropbox.com/developers/apps) and create a new project. Select the following options:
  * **Choose an API**: Dropbox API
  * **Choose the type of access you need**: App folder
  * **Name your app**: Any available name.
3. Add a new `Redirect URI`: `http://localhost:8080/v1/auth/login` (in case your `serverAddr` is `localhost:8080`)
4. Then copy the following values into your configuration file:
  * `appKey`: App key
  * `appSecret`: App secret
  * `redirectUri`: Redirect URIs

And that's all.

## Install and start the server

Change to the project's folder and execute the next commands:
```bash
# install the application
$ go install

# start the server
$ photomanager
2019/02/20 19:32:59 Server started at port localhost:8080
```

Make sure that `$GOPATH/bin` is included in your `$PATH` environment variable.
