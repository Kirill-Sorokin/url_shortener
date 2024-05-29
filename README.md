# URL Shortener

This is a simple URL shortener service built with Golang for the backend and HTML/CSS/JavaScript for the frontend. The service allows users to shorten long URLs and retrieve the original URLs using the shortened versions.

## Features

- **URL Shortening**: Converts long URLs into short, unique URLs.
- **URL Redirection**: Redirects the user to the original URL when they visit the shortened URL.
- **Concurrency**: Utilizes Go's concurrency features to handle multiple requests efficiently.

## Project Structure

- **Backend**: Golang, Gorilla Mux for routing.
- **Frontend**: HTML/CSS/JavaScript.
- **Database**: In-memory data store (for simplicity).

## How to Run

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)
- A web browser

### Steps

1. **Clone the repository**:
   ```sh
   git clone https://github.com/Kirill-Sorokin/url_shortener.git
   cd url_shortener
