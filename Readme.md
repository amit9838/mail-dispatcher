# Mail Dispatcher

## Introduction
Mail Dispatcher is a lightweight, concurrent Go application designed to process and send bulk emails. It reads recipient data from a CSV file and utilizes a worker pool to send templated emails to an SMTP server concurrently.

## Project Architecture
The application uses a **Producer-Consumer** pattern to achieve high concurrency and performance:
- **Producer**: A goroutine reads recipient information from a CSV file (`emails.csv`) line by line and feeds it into a Go channel.
- **Consumer (Worker Pool)**: Multiple worker goroutines listen to the channel. As soon as a recipient is available, a worker picks it up, generates the email message using a Go template, and sends it via the `net/smtp` package to the configured SMTP server.
- **Synchronization**: A `sync.WaitGroup` ensures the main program waits for all workers to finish processing the queue before exiting.

## Directory Structure
```text
mail-dispatcher/
├── Readme.md      # Project documentation
├── consumer.go    # Defines the worker logic (emailWorker) for sending emails
├── email.tmpl     # Go template for formatting the email headers and body
├── emails.csv     # Input data file containing recipient names and emails
├── go.mod         # Go module dependency definitions
├── main.go        # Main entry point; orchestrates channels, workers, and template execution
└── producer.go    # Logic for parsing the CSV file and pushing data to the channel
```

## Dependencies
This project relies solely on the **Go Standard Library** (e.g., `net/smtp`, `encoding/csv`, `html/template`, `sync`). No external third-party Go packages are required.

**External Requirement:**
To test the email dispatching locally without sending actual emails, you need a mock SMTP server running on `localhost:1025` (such as [MailHog](https://github.com/mailhog/MailHog) or [Mailpit](https://github.com/axllent/mailpit)).

## Build
To compile the application into a standalone executable, run:
```bash
go build -o mail-dispatcher .
```
This will generate an executable binary named `mail-dispatcher` in the root directory.

## Test
To run the project tests (if any are added to `*_test.go` files), use the standard Go test command:
```bash
go test -v ./...
```

## Running Locally
To run the application directly from the source code, ensure your local SMTP server is running on port `1025`, then execute:
```bash
go run .
```
This will trigger the producer to read `emails.csv` and the worker pool to dispatch the emails concurrently.

## Checklist for improvements
To make this application production-ready and adhere to industry standards, the following improvements should be considered. They are categorized by difficulty:

### Easy (Quick Wins)
- [ ] **Configuration Management**: Externalize configurations (SMTP host, port, worker count, sender email) using environment variables (e.g., `.env`) or command-line flags instead of hardcoding them.
- [ ] **Template Caching**: Parse the email template (`email.tmpl`) once at startup rather than on every email dispatch to significantly improve CPU utilization and latency.
- [ ] **Input Validation**: Add structural validation to ensure recipient email addresses are well-formed before attempting delivery.
- [ ] **Modularization**: Organize code into separate sub-packages (e.g., `internal/worker`, `internal/producer`) instead of keeping everything in `main.go` to improve maintainability.

### Medium (Architecture & Stability)
- [ ] **Structured Logging**: Replace standard `fmt.Printf` with a structured logging library (like `log/slog`, `logrus`, or `zap`) for better observability.
- [ ] **Memory-Efficient CSV Processing**: Refactor `producer.go` to stream the CSV file line-by-line (`r.Read()`) instead of loading all records into memory at once (`r.ReadAll()`), allowing it to handle massive files.
- [ ] **Context & Graceful Shutdown**: Implement `context.Context` to handle OS signals (SIGTERM, SIGINT) and ensure workers can finish processing inflight emails before shutting down gracefully.
- [ ] **Containerization & Tooling**: Create a `Dockerfile`, `docker-compose.yml`, and a `Makefile` to streamline deployment, local environment setup, and developer onboarding.
- [ ] **Testing & CI/CD**: Add robust unit tests, mock SMTP handlers for integration tests, and a CI/CD pipeline (e.g., GitHub Actions) for automated verification.

### Hard (Scalability & Resilience)
- [ ] **Connection Pooling**: Use a persistent SMTP connection pool (e.g., maintaining `net/smtp.Client` instances) instead of opening and closing a new TCP connection for every single email via `smtp.SendMail`.
- [ ] **Error Handling & Retries**: Add a robust retry mechanism (e.g., exponential backoff) and a persistent dead-letter queue (DLQ) for tracking and re-processing emails that fail to send.
- [ ] **Rate Limiting**: Implement a rate limiter to respect upstream SMTP server sending limits and avoid being blocked or throttled.
- [ ] **Metrics & Monitoring**: Expose a lightweight HTTP server to serve Prometheus metrics (e.g., tracking total successful, failed, and queued emails) or health checks.
