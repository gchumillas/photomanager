// The following script creates the structure of the application database.
//
// Open a terminal and execute the following command:
// mongo <database name> create_db_schema.js
//
// Alternatively you can execute following commands:
//
// mongo
// use photomanager
// load('create_db_schema.js')

print('Create images collection');
db.createCollection(
  'images',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['path'],
        properties: {
          path: {
            bsonType: 'string'
          }
        }
      }
    }
  }
);

print('Create categories collection');
db.createCollection(
  'categories',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['name'],
        properties: {
          name: {
            bsonType: 'string'
          },
          imageIds: {
            bsonType: 'array',
            items: {
              bsonType: 'objectId'
            }
          }
        }
      }
    }
  }
);

print('Create users collection');
db.createCollection(
  'users',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['name', 'categoryId'],
        properties: {
          name: {
            bsonType: 'string'
          },
          categoryId: {
            bsonType: 'objectId'
          }
        }
      }
    }
  }
);
