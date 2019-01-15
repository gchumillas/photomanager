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
    },
    validationLevel: "strict",
    validationAction: "error"
  }
);

print('Create categories collection');
db.createCollection(
  'categories',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['name', 'imageIds'],
        properties: {
          name: {
            bsonType: 'string'
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
    validationLevel: "strict",
    validationAction: "error"
  }
);

print('Create users collection');
db.createCollection(
  'users',
  {
    validator: {
      $jsonSchema: {
        bsonType: 'object',
        required: ['name'],
        properties: {
          name: {
            bsonType: 'string'
          },
          categoryIds: {
            bsonType: 'array',
            items: {
              bsonType: 'objectId'
            }
          }
        }
      }
    },
    validationLevel: "strict",
    validationAction: "error"
  }
);
