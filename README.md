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

# Web Tracker - Plan

## 1. Problem Analysis

### Core Problem
- REST API to track web visitor events
- Serve unique visitors analytics by URL
- In-memory storage with concurrency support

### Features
- **Feature 1**: Ingest visitor events via REST endpoint
- **Feature 2**: Query unique visitors analytics via REST endpoint in JSON format

### Event Ingestion
**Endpoint**: `POST /api/events`
```json
{
  "uid": "uuid-visitor",
  "url": "https://example.com/page"
}
```

### Analytics Query
**Endpoint**: `GET /api/unique-visitors?url=https://example.com/page`
```json
{
  "url": "https://example.com/page",
  "unique_visitors": 1250
}
```

## 3. Technical Constraints

- HTTP Server REST
- In-memory storage
- Concurrency support
- Thread-safe structures
- No external libs on concurrency data access

## 4. Technology Stack

- **Go 1.24**: Performance + concurrency
- **gorilla/mux**: REST routing
- **Memory**: sync.Map, sync.RWMux for thread-safe storage
- **testify**: Unit testing

## 5. Architecture

### Memory Store Design

- Possible solutions:
  
  1. Map storage
  2. Bloom filters

### Project Structure: Vertical Slicing aproach
```
cmd/
├── envs
├── http-server
├── main.go
internal/
├── domain/
│   ├── visitor-event.go
│   ├── visitor-analytic.go
│   └── visitor-repository.go
├── features/
│   ├── ingest-visitors-events/
│   │   └── controller.go
│   │   └── mapper.go
│   │   └── validator.go
│   │   └── command.go
│   └── get-visitors-analytics/
│       └── controller.go
│       └── mapper.go
│       └── validator.go
│       └── query.go
├── infra/
│   └── in-memory/
│       └── visitor-repository.go
Makefile
README.md 
...
```

## 6. Implementation Plan

### Phase 0: Base Setup
- Go module + Makefile
- Basic HTTP server with gorilla/mux
- Health check endpoint
- 
### Phase 1: Domain
- visitors event
- visitors analytic
- visitors repository

### Phase 2: Events API
- `POST /web-tracker/new-visitor` controller
- mapper
- validator
- command
- command-handler
- test

### Phase 3: Analytics API
- `GET /web-tracker/analytics` controller
- mapper
- validator
- query
- query-handler
- test

### Phase 4: In-Memory repository
- add unique visitor
- is unique visitor
- get analytics 
- test

### Phase 4: Production Ready
- Rate limiting config
- Graceful shutdown
- Performance optimization

## API Endpoints
```
POST   /web-tracker/new-visitor
GET    /web-tracker/analytics
GET    /health
```
## Makefile Commands
```makefile
run:    # Run locally
test:   # Unit tests + coverage
build:  # Production build
lint:   # Code quality checks
performace: # Performance measurements
```
## Future Improvements
- TTL strategy for old events cleanup
- Open-api specification
- Docker contenerization