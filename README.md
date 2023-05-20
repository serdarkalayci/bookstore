# bookstore
Bookstore is a simple series of web services which forms a book store.

## info
Info service uses a mongodb database to store book information. It exposes a REST API to list of books and get the details of a book.

## stock
Stock service uses a redis database to store book stock information. It exposes a REST API to get the stock amount of a book and update it.

## basket
Basket service uses a mongodb database to store basket information. It exposes a REST API to add a book to the basket, remove a book from the basket and list the books in the basket.

## payment
Payment service uses a redis database to store payment information. It exposes a REST API to make a payment.