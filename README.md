# csv-file-store

## Description
This is a simple http server for storing csv files.

## Usage
To start the server, run the following command:
```go run ./cmd/csv-file-store```

The server will start on port 8080 by default.

## Endpoints

### POST /API/upload
This endpoint is used to upload a csv file to the server.
The request should contain a file in the body. This will then be queued for processing.

### GET /API/download/{id}
This endpoint is used to download a csv file from the server.
The request should contain the id of the file to download.
This needs to be in the format of a UUID (e.g. 123e4567-e89b-12d3-a456-426614174000).

## Processing
When a file is uploaded, it is queued for processing. 
Processing involves detecting whether each row contains an email address and then writing a boolean value to a new column in the file.
The server will then process the file and store the data in the local file system. 
Once the file has been processed, it will be available for download.

## Testing
To run the tests, run the following command:
```go test ./...```

## Future Improvements
- keep a history of file uploads
- add a database to store file data
- add the ability to remove files
- add the ability to update files
- add an expiry time for files
- add a limit to files that can be uploaded