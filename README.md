# Classify Number API

## Overview
This API classifies a given number based on its mathematical properties. It checks whether the number is prime, perfect, or an Armstrong number. It also calculates the digit sum and fetches a fun fact about the number.

## API Endpoint
### **GET /api/classify-number**



#### **Response Format:**
```json
{
  "number": 477,
  "is_prime": false,
  "is_perfect": false,
  "properties": ["odd"],
  "digit_sum": 18,
  "fun_fact": "477 is the smallest number whose cube contains four 3's."
}
```

## Features
- Checks if a number is **prime**.
- Checks if a number is **perfect**.
- Checks if a number is an **Armstrong number**.
- Calculates the **digit sum**.
- Determines if the number is **even or odd**.
- Fetches a **fun fact** about the number.



### **Setup and Run**
```sh
git clone https://github.com/yourusername/classify-number-api.git
cd classify-number-api
go mod tidy
go run cmd/main.go
```

## Testing the API
You can test the API using **curl** or Postman:
```sh
curl -X GET "http://localhost:8080/api/classify-number?number=477"
```

