# PhotoManager

An application to manage pictures from Dropbox.

The application connects to a Dropbox account and allows the user to organize their images into categories.

## Create the database

This application uses a MongoDB database. To create it simply execute the following command from your terminal:
```bash
mongo [ENTER YOUR DATABASE NAME] dbschema.js
```

## Configuration
Rename `config-example.json` to `config.json` and change the variables.

## Start the server
execute the following command from your terminal:
```bash
./photomanager
```
