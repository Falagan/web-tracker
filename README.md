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