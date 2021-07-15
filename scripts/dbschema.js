// The following script creates the structure of the application database.
//
// Open a terminal and execute the following command:
// mongo <database name> dbschema.js
//
// Alternatively you can execute following commands:
//
// mongo
// use <database name>
// load('dbschema.js')

print('Create users collection');
db.getCollection('users').drop();
db.createCollection(
  'users',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['accountId', 'accessToken'],
        additionalProperties: false,
        properties: {
          accountId: {
            bsonType: 'string'
          },
          accessToken: {
            bsonType: 'string'
          }
        }
      }
    },
    validationLevel: 'strict',
    validationAction: 'error'
  }
);

print('Create categories collection');
db.getCollection('categories').drop();
db.createCollection(
  'categories',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['name', 'userId', 'imageIds'],
        additionalProperties: false,
        properties: {
          name: {
            bsonType: 'string'
          },
          userId: {
            bsonType: 'objectId'
          },
          imageIds: {
            bsonType: 'array',
            items: {
              bsonType: 'objectId'
            }
          },
          parentCategoryId: {
            bsonType: 'objectId'
          }
        }
      }
    },
    validationLevel: 'strict',
    validationAction: 'error'
  }
);

print('Create images collection');
db.getCollection('images').drop();
db.createCollection(
  'images',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['userId', 'imageId'],
        additionalProperties: false,
        properties: {
          userId: {
            bsonType: 'objectId'
          },
          imageId: {
            bsonType: 'string'
          }
        }
      }
    },
    validationLevel: 'strict',
    validationAction: 'error'
  }
);
