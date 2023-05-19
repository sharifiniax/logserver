### Install the required dependencies and run:

<pre><code> 
cd log-api
go mod download
go run main.go
</code></pre>

## Usage

### Sending Log Messages

To send log messages to the API, make a POST request to the /logs endpoint with the following JSON payload in port 8080:

```json
{
  "level": "info",  // DEBUG, INFO, WARNING, ERROR
  "message": "This is an info log message",
  "tag": "myTag"
}
```


The level field specifies the log level, which can be one of the following values: DEBUG, INFO, WARNING, or ERROR. The message field contains the log message, and the optional tag field allows you to provide additional context or categorize the log entries.







