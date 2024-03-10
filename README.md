# Running the Code Locally

This repository contains GoLang code for a simple HTTP server that implements an API for searching documents and generating summaries based on keywords.

## Prerequisites

Before running the code, ensure you have the following installed on your local machine:

- GoLang: [Download and install Go](https://golang.org/doc/install)
- Git: [Download and install Git](https://git-scm.com/downloads)

## Getting Started

To run the code locally, follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/ChetanSureka/OnFinanceAI_task2.git
   ```

2. Navigate to the project directory:

   ```bash
   cd OnFinanceAI_task2
   ```

3. Run the following command to start the HTTP server:

   ```bash
   go run main.go
   ```

4. The server will start running on `http://localhost:8080`.

## Making Requests

Once the server is running, you can make HTTP GET requests to the `/search` endpoint to search for documents and generate summaries based on keywords.

Example request using cURL:

```bash
curl -X GET -H "Content-Type: application/json" -d '{"keyword": "your-search-keyword"}' http://localhost:8080/search
```

Replace `your-search-keyword` with the keyword you want to search for.

## API Response

The API will respond with JSON data containing the original keyword, new keywords, matched chunks, and the summary.


## Example Request

```json
{
  "keyword": "TATA Power"
}
```

## Example Response

```json
{
  "original_keyword": "TATA Power",
  "new_keywords": [],
  "matched_chunks": [
    "TATA Power Company is an Indian electric utility company based in Mumbai, Maharashtra, India and is part of the Tata Group.",
    "Its primary business is to generate, transmit and distribute electricity.",
    "TATA Power has installed capacity of 10957 MW. Its core business includes generation, transmission, distribution and trading of electricity.",
    "The company has operations in India, Singapore, Indonesia, South Africa, and Bhutan."
  ],
  "summary": "Generated summary based on the search results"
}
```