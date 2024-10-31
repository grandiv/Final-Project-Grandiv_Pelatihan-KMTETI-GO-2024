# Final-Project-Grandiv_Pelatihan-KMTETI-GO-2024

## Bookstore APi Documentation

A REST API server for managing books and employees in a bookstore. Built with Go and MongoDB. Assisted by GitHub Copilot (Claude 3.5 Sonnet and GPT 4o)

## Setup

Clone the project

```bash
  git clone https://github.com/grandiv/Final-Project-Grandiv_Pelatihan-KMTETI-GO-2024.git
```

Go to the project directory

```bash
  cd Final-Project-Grandiv_Pelatihan-KMTETI-GO-2024
```

Install dependencies

```bash
  go mod download
  go mod tidy
```

Create `env` file in root directory:

```bash
  MONGODB="your_mongodb_connection_string"
```

Start the server locally

```bash
  ./dist/server.exe
```

## Collections and Fields Definition in English

| **Collections**     | **Definition**                                               |
| ------------------- | ------------------------------------------------------------ |
| buku                | Book                                                         |
| karyawan            | Employee                                                     |
| **Fields**          | **Definition**                                               |
| judul               | Title of the Book                                            |
| penulis             | Author of the Book                                           |
| tahun               | The Year a Book was Published                                |
| stok                | Current Stock Quantity of a Book                             |
| harga               | Book Price                                                   |
| nama                | Employee Name                                                |
| nik                 | National Identification Number (NIN) of an Employee          |
| pendidikan_terakhir | Last Education of an Employee                                |
| tanggal_masuk       | Date Employed                                                |
| status_kerja        | Binary values of "Kontrak" (contract) or "Tetap" (full time) |

## API Endpoints

**There are 7 API endpoints which are specifically created for the final project:**

### 1. Retrieve All Books (with fields "judul", "penulis", and "harga")

```
GET http://localhost:8080/api/buku/
```

### 2. Retrieve a Detailed Book Info (with fields "judul", "penulis", "tahun", "stok", and "harga")

```
GET http://localhost:8080/api/buku/{id}
```

### 3. Add a New Book to the Database

```
POST http://localhost:8080/api/buku/
Body (example):
{
    "judul": "Harry Potter and the Philosopher's Stone",
    "penulis": "J.K. Rowling",
    "tahun": 1997,
    "stok": 12,
    "harga": 350000
}
```

### 4. Update a Book's Stock and Price

```
PUT http://localhost:8080/api/buku/{id}
Body (example):
{
    "stok": 10,
    "harga": 345000
}
```

### 5. Delete a Specific Book

```
DELETE http://localhost:8080/api/buku/{id}
```

### 6. Retrieve All Employees (with fields "nama", "tanggal_masuk", and "status_kerja")

```
GET http://localhost:8080/api/karyawan/
```

### 7. Add a New Employee

```
POST http://localhost:8080/api/karyawan/
Body (example):
{
    "nama": "grandiv",
    "nik": "1337",
    "pendidikan_terakhir": "S1",
    "tanggal_masuk": "2022-08-10",
    "status_kerja": "Kontrak"
}
```

## Error Responses

The API returns appropriate HTTP status codes:

- 200: Success
- 201: Created successfully
- 400: Bad request / Invalid input
- 404: Resource not found
- 405: Method not allowed
- 500: Internal server error

Error responses include an error message in the response body.
