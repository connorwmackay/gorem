# gorem

A Small Go Rest API, currently no client-side.

## Current Features

- API authentication, with a bearer token stored in the .env file
- Password hashing with sha512
- User creation
- User authentication (username, password)
- Posts
  - Creation
  - Get by ID
- Comments (on a post)
  - Creation
  - Get by ID

## Planned Features
- Database integration (Possibly SQLite)
- User sessions
- Posts
  - Add tags to PostResponse
  - Add rating to PostResponse
  - Editing
  - Deleting
  - List by Filtering (by authorId, tags, rating)
- Comments (on a post)
  - Editing
  - Deleting
  - List all by Post Id
- User signing out
