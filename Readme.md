# Mailchamp

## Introduction
Mailchamp is a lightweight, concurrent Go application designed to process and send bulk emails. It reads recipient data from a CSV file and utilizes a worker pool to send templated emails to an SMTP server concurrently.

## Project Architecture
The application uses a **Producer-Consumer** pattern to achieve high concurrency and performance:
- **Producer**: A goroutine reads recipient information from a CSV file (`emails.csv`) line by line and feeds it into a Go channel.
- **Consumer (Worker Pool)**: Multiple worker goroutines listen to the channel. As soon as a recipient is available, a worker picks it up, generates the email message using a Go template, and sends it via the `net/smtp` package to the configured SMTP server.
- **Synchronization**: A `sync.WaitGroup` ensures the main program waits for all workers to finish processing the queue before exiting.

## Directory Structure
```text
Mailchamp/
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
go build -o mailchamp .
```
This will generate an executable binary named `mailchamp` in the root directory.

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
