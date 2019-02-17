# PhotoManager

An application to manage pictures from Dropbox.

The application connects to a Dropbox account and allows the user to organize their images into categories.

## Create the database

This application uses a MongoDB database. To create it just execute the following command from your terminal (Replace `DATABASE_NAME`):
```bash
mongo DATABASE_NAME dbschema.js
```

## Configuration
Rename `config-example.json` to `config.json` and change the variables.

## Start the server
Simply execute the following command from your terminal:
```bash
./photomanager
```
