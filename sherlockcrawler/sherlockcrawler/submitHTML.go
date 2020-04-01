package sherlockcrawler

import (
	"context"

	"github.com/micro/go-micro"
	sender "github.com/ob-algdatii-20ss/leistungsnachweis-dievierausrufezeichen/sherlockcrawler/proto"
	"github.com/pkg/errors"
)

/*
size of a chunk. important to minimize the amount of bytes sent at once
*/
const chunkSize int = 1024

/*
Sender
*/
type ClientGRPC struct {
	client    sender.SenderService
	chunkSize int
}

/*
creates a new sender
*/
func NewClientGRPC(service micro.Service) (c ClientGRPC) {
	c.chunkSize = chunkSize
	c.client = sender.NewSenderService("client", service.Client())

	return c
}

/*
help method
*/
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
cuts byte array in slices of chunksize and sends them to the analyzer
*/
func (c *ClientGRPC) UploadFile(ctx context.Context, ltask CrawlerTaskRequest) error {
	stream, err := c.client.Upload(ctx)

	if err != nil {
		return errors.New("failed to create upload stream for file")
	}

	var lengthByteArray int = len(ltask.getResponseBodyInBytes())

	for i := 0; i < lengthByteArray; i += chunkSize {
		buf := ltask.getResponseBodyInBytes()[i:min(i+chunkSize, lengthByteArray)]

		err = stream.Send(&sender.Chunk{
			Content: buf,
		})

		if err != nil {
			return errors.New("error while streaming")
		}
	}

	//receive fehlt

	err = stream.Close()

	if err != nil {
		return errors.New("error while closing stream")
	}
	return nil
}
