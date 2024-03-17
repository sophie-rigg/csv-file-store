package csv

import (
	"context"
	"encoding/csv"
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sophie-rigg/csv-file-store/utils"
)

type Client struct {
	*csv.Writer
	*csv.Reader
	zerolog.Logger
}

func New(ctx context.Context, writer io.Writer, reader io.Reader) (*Client, error) {
	csvReader := csv.NewReader(reader)
	csvWriter := csv.NewWriter(writer)

	// Assume the first row is the header
	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	// Add the email_exists header
	headers = append(headers, "email_exists")

	err = csvWriter.Write(headers)
	if err != nil {
		return nil, err
	}

	return &Client{
		Writer: csvWriter,
		Reader: csvReader,
		Logger: log.Ctx(ctx).With().Fields(map[string]interface{}{
			"writer": "csv",
		}).Logger(),
	}, nil
}

func (c *Client) Write() error {
	for {
		record, err := c.Reader.Read()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		// Check if the record contains an email and add the email_exists field
		if utils.RecordContainsEmail(record) {
			record = append(record, "true")
		} else {
			record = append(record, "false")
		}

		err = c.Writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Close() {
	c.Writer.Flush()
}
