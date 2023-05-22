db = new Mongo().getDB("bookinfo");
db.books.drop();
db.books.createIndex( { "isbn": 1 }, { unique: true } )

db.books.insertMany([
  {
    "isbn": "9780141439563",
    "title": "Pride and Prejudice",
    "author": "Jane Austen",
    "price": 9.99,
    "publishDate": "1813-01-28"
  },
  {
    "isbn": "9780061120084",
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "price": 12.5,
    "publishDate": "1960-07-11"
  },
  {
    "isbn": "9780060256654",
    "title": "Where the Wild Things Are",
    "author": "Maurice Sendak",
    "price": 7.99,
    "publishDate": "1963-04-19"
  },
  {
    "isbn": "9780743273565",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "price": 10.99,
    "publishDate": "1925-04-10"
  },
  {
    "isbn": "9780545010221",
    "title": "Harry Potter and the Sorcerer's Stone",
    "author": "J.K. Rowling",
    "price": 14.99,
    "publishDate": "1997-06-26"
  },
  {
    "isbn": "9781400079988",
    "title": "The Kite Runner",
    "author": "Khaled Hosseini",
    "price": 11.99,
    "publishDate": "2003-05-29"
  },
  {
    "isbn": "9780140449334",
    "title": "1984",
    "author": "George Orwell",
    "price": 8.99,
    "publishDate": "1949-06-08"
  },
  {
    "isbn": "9780062315007",
    "title": "The Fault in Our Stars",
    "author": "John Green",
    "price": 9.99,
    "publishDate": "2012-01-10"
  },
  {
    "isbn": "9780060891541",
    "title": "The Catcher in the Rye",
    "author": "J.D. Salinger",
    "price": 10.5,
    "publishDate": "1951-07-16"
  },
  {
    "isbn": "9780307588364",
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "price": 12.99,
    "publishDate": "1937-09-21"
  },
  {
    "isbn": "9780141182801",
    "title": "To the Lighthouse",
    "author": "Virginia Woolf",
    "price": 9.99,
    "publishDate": "1927-05-05"
  },
  {
    "isbn": "9781400033423",
    "title": "Atonement",
    "author": "Ian McEwan",
    "price": 11.99,
    "publishDate": "2001-09-04"
  },
  {
    "isbn": "9780441172719",
    "title": "Dune",
    "author": "Frank Herbert",
    "price": 13.99,
    "publishDate": "1965-08-01"
  },
  {
    "isbn": "9780679723165",
    "title": "One Hundred Years of Solitude",
    "author": "Gabriel Garcia Marquez",
    "price": 10.99,
    "publishDate": "1967-06-05"
  },
  {
    "isbn": "9780399501487",
    "title": "The Shining",
    "author": "Stephen King",
    "price": 9.99,
    "publishDate": "1977-01-28"
  },
  {
    "isbn": "9780062562584",
    "title": "The Alchemist",
    "author": "Paulo Coelho",
    "price": 10.99,
    "publishDate": "1988-01-01"
  },
  {
    "isbn": "9780679732262",
    "title": "Beloved",
    "author": "Toni Morrison",
    "price": 11.5,
    "publishDate": "1987-09-02"
  },
  {
    "isbn": "9780062316097",
    "title": "Gone Girl",
    "author": "Gillian Flynn",
    "price": 11.99,
    "publishDate": "2012-06-05"
  },
  {
    "isbn": "9780812986219",
    "title": "The Help",
    "author": "Kathryn Stockett",
    "price": 10.99,
    "publishDate": "2009-02-10"
  },
  {
    "isbn": "9780143105985",
    "title": "The Book Thief",
    "author": "Markus Zusak",
    "price": 12.99,
    "publishDate": "2005-03-14"
  },
  {
    "isbn": "9780671027032",
    "title": "Ender's Game",
    "author": "Orson Scott Card",
    "price": 9.99,
    "publishDate": "1985-01-15"
  }
]
)
