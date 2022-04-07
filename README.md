# github.com/mahfuz110244/golang-assignment
GraphQL demo using Golang and MySQL

## Run
###
Create MySQL graphql_demo using below docker command
```
docker-compose up --build
```

### Run the server
```
go run .
```

### GraphQL playground
```
http://localhost:8080/graphql
```

## Query

### Create two authors John and Karim and Get Authors
```
mutation CreateAuthor1 {
  createAuthor(name: "John", biography: "good writer") {
    id
    name
    biography
  }
}

query GetAuthorJohn{
  author(id:1) {
      id
      name
      biography
  }
}


mutation CreateAuthor2 {
  createAuthor(name: "Karim", biography: "good writer") {
    id
    name
    biography
  }
}

query GetAuthorKarim{
  author(id:2) {
      id
      name
      biography
  }
}


query GetAuthorNotFound{
  author(id:9999) {
      id
      name
      biography
  }
}

query GetAllTheAuthors{
  author_list {
      id
      name
      biography
  }
}
```

### Create two books with author John and one book with author Karim and get books
```
mutation CreateBook1 {
  createBook(title: "Demo Book 1 John", price: 1000, isbn_no: "BK98969887", author_id:"1"){
    id
    title
    price
    isbn_no
    authors{
      id
      name
    	biography
    }
  }
}

query GetBooks1 {
  book(id:1) {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
		} 
  }
}

mutation CreateBook2 {
  createBook(title: "Demo Book 2 John", price: 1000, isbn_no: "BK98969888", author_id:"1"){
    id
    title
    price
    isbn_no
    authors{
      id
      name
    	biography
    }
  }
}

mutation CreateBook2 {
  createBook(title: "Demo Book 3 Karim", price: 500, isbn_no: "BK98969889", author_id:"2"){
    id
    title
    price
    isbn_no
    authors{
      id
      name
    	biography
    }
  }
}

query GetBooks2 {
  book(id:2) {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
		} 
  }
}

query GetBooksNotFound {
  book(id:99999) {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
		} 
  }
}
```

### Get all the books of John
```
query GetAllTheBooksOfJohn{
  authors(name: "John") {
    books{
      id
      title
      price
      isbn_no
    }
  }
}
```

### Get all the books of all authors
```
query GetAllBooksWithAuthors {
  books {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
		} 
  }
}
```


## Challenges
As I have no experience in GraphQL previously, the first challenge for me is understanding GraphQL and its structure. Second challenge is wrinting the schema according to
```
https://github.com/graphql-go/graphql
```
this library in golang.
