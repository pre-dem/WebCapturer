package pandora

import (
	"errors"
	"time"

	. "github.com/qiniu/pandora-go-sdk/base"
	"github.com/qiniu/pandora-go-sdk/base/config"
	"github.com/qiniu/pandora-go-sdk/pipeline"
	"log"
)

type Client struct {
	logger       Logger
	client       pipeline.PipelineAPI
	repo         string
	pointChannel chan []pipeline.Point
}

const (
	bufferSize    = 16 * 1024
	batchSize     = 128
	maxWaitSecond = 2
	timeout       = 10
)

var (
	ErrSinkTimeout = errors.New("Sink timeout")
)

func NewClient(ak, sk, repo string) *Client {
	logger := NewDefaultLogger()
	cfg := config.NewConfig().
		WithEndpoint(p_endpoint).
		WithAccessKeySecretKey(ak, sk).
		WithLogger(logger).
		WithLoggerLevel(LogDebug)

	c, _ := pipeline.New(cfg)
	dataCh := make(chan []pipeline.Point, bufferSize)

	client := &Client{logger, c, repo, dataCh}
	go client.startLoop()
	go client.pingChan()
	return client
}

func (c *Client) Sink(points []pipeline.Point) error {
	select {
	case c.pointChannel <- points:
		return nil
	case <-time.After(time.Duration(timeout) * time.Millisecond):
		return ErrSinkTimeout
	}
}

func (c *Client) startLoop() {
	buffer := make(pipeline.Points, bufferSize)
	buffer = buffer[:0]
	last := time.Now()
	for data := range c.pointChannel {
		if data != nil {
			buffer = append(buffer, data...)
		}

		if len(buffer) >= batchSize || time.Now().After(last.Add(time.Second*time.Duration(maxWaitSecond))) {
			if err := c.post(buffer); err != nil {
				// TODO retry
				log.Println("loopSend failed -", err)
			}

			buffer = buffer[:0]
			last = time.Now()
		}
	}
}

func (c *Client) post(points pipeline.Points) error {
	if len(points) == 0 {
		return nil
	}
	postDataInput := &pipeline.PostDataInput{
		RepoName: c.repo,
		Points:   points,
	}
	err := c.client.PostData(postDataInput)
	if err != nil {
		log.Println(err)
		c.logger.Error(c.repo, err)
	}
	return err
}

// 在某些没更新的chan里的sendBuf里可能有数据
// 需要有个后台 goruntine 不停的发送nil，来触发发送当作
func (c *Client) pingChan() {
	for range time.Tick(time.Second) {
		c.pointChannel <- nil
	}
}
