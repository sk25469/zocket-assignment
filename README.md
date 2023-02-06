# Zocket-Assignment

## Zocket CRUD API

### Base address - <http://20.204.236.223:3000>

GET - /book returns all the books

POST - /book takes book properties as follows

```json
{
    "name": <book_name>,
    "publication": <book_publication>,
    "count": <book_count>
}
```

PUT - /book/<book_id> updates the book with given id

DELETE - /book/<book_id> deletes the book with given id
