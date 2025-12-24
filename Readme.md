# Redis-Based Rate Limiter (Gin + Go)

A simple **rate limiter middleware** built using **Go (Gin framework)** and **Redis**.  
This project demonstrates how to protect APIs from multiple requests by limiting the number of requests per client within a fixed time window.

---

## Features

- Fixed-window rate limiting
- Redis-backed (distributed & concurrency-safe)
- Gin middleware integration
- Per-client request tracking
- Automatic window reset using Redis TTL
- Graceful rejection with HTTP 429 Too Many Requests

---

## How It Works

- Each incoming request is intercepted by middleware
- A Redis counter is incremented per client (keyed by Requestor-id or IP)
- The first request starts a 1-minute window using Redis EXPIRE
- Requests beyond the configured limit are rejected
- When the TTL expires, the counter resets automatically

---

## Setup
1.Clone this repository within a folder  
2.Add the field REDIS_URL=redis://<host_name>:<port_number>/<db_number> inside the .env file  
3.Run the following commands in the root of the project
```
go mod tidy
```
```
go run ./cmd
```  

4.Send requests consecutive requests to the following endpoint  
**<host_name>/8080/health-check/ping**

## Observations
If you send 200 requests within the period of one minute, in the same time frame you won't be able to send the 201th request.  
  
  All the requsts after the specified number are aborted by the middleware.
