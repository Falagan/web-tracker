# Web Tracker

## Context

We have a website on the Internet and we would like to get some very simple indication
of how visitors navigate the pages. For that purpose, we managed to configure our
website to send an event every time a visitor navigates to a page. Our website is
capable of generating unique identifiers for visitors as a string of characters.
The system generating that event is able to talk to a REST HTTP interface and
represents each individual event as a JSON document containing two attributes: the
unique identifier of the visitor and the URL of the visited page.
Our product team is starting a new sprint. We are picking the following user story:
As a digital marketeer, I need to know how many distinct visitors navigated to a page,
knowing its URL.

## Task
Build a GoLang web service capable of:
Ingesting user navigation JSON events via a REST HTTP endpoint. Each event is
to be ingested via a separate HTTP request (i.e. no batch and no streaming
ingestion).
Serving the number of distinct visitors for any given page via another REST HTTP
endpoint. The page URL we are interested in should be a query parameter of the
HTTP request. The number of distinct visitors for that URL is returned in a JSON
object.

## Constraints
There is no need for persistence to a database. Everything can be kept in memory.
The web service must be capable of handling concurrent requests on both
endpoints.
Don't solve the data access concurrency problem using an external library

----

# HightLevel Plan

## 1. Analysis

1. **Understand the problem**
2. **Define required features and constrains**
3. **Analyze different solutions for each feature**
4. **Select solution based on requirements and balancing results vs trade-offs**
5. **Select best fit architecture**

## 2. Scaffolding

- **Define go module** and base main hello world

- **Setup Makefile** for basic commands:
  - Run: Locally/Dev env/Production.
  - Test: ensure code compliance.
  - Build: Production build optimization.
  - Lint: Ensure code quality.
  - Performance: local performace Analysis

- **Implement base architecture**

## 3. Feature Development

1. **Feature: Ingest web visitor events**

2. **Feature: Serve data about unique visitors by URL**

----
# LowLevel Plan

## Architecture

### Vertical Slicing
- `cmd` > app shell to expose features. In this case a Basic HTTP Server
- `internal` 
  - `domain` > core business
  - `features` > explicit features
  - `infra` > data repositories
- root files


### Common Base 'must-have':
  - Envs
  - Debugger config
  - Logger
  - Health Check
  - Context Global Timeout
  - Graceful shutdown
  - Rate limit
  - HTTPS

## Constraints
1. Concurrency support
2. In memory storage
3. Thread-safe structures
4. Memory limits/optimization

### Extras to analyze convenience:
1. TTL data strategy
2. Cache strategies

## Development

### 0. Setup base HTTP server

### 1. Domain
- Define Ingest entities
- Define Analytics entities 
  
### 2. Infra: Repositories
- Define track repository

### 3. Feature: Ingest web visitor events
- POST endpoint to receive events
- Request Mapper
- Request Validation
- In-memory storage with thread-safe structures (sync.Map/sync.RWMux)
- Response Mapper
- Proper Errors Handling
- Unit Test

### 4. Feature: Serve data about unique visitors by URL
- GET endpoint for analytics queries
- Request Mapper
- Request/Query Validation
- Params query support
- In-memory storage with thread-safe structures (sync.Map/sync.RWMux)
- Response Mapper with structured JSON responses
- Proper Errors Handling
- Unit Test

### 5. Add extra features to HTTP Server