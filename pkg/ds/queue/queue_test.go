package queue_test

import (
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
	// Demonstrate that the queue works with multiple data types
	queue       *queue.Queue[int]
	stringQueue *queue.Queue[string]
	floatQueue  *queue.Queue[float64]
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}

func (suite *QueueTestSuite) TestNew() {
	suite.queue = queue.New[int]()
	expectedLength := 3

	suite.queue.Enqueue(5)
	suite.queue.Enqueue(7)
	suite.queue.Enqueue(9)
	assert.NotNil(suite.T(), suite.queue)
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}

func (suite *QueueTestSuite) TestNewStringQueue() {
	suite.stringQueue = queue.New[string]()
	expectedLength := 2

	suite.stringQueue.Enqueue("hello")
	suite.stringQueue.Enqueue("world")
	assert.NotNil(suite.T(), suite.stringQueue)
	assert.Equal(suite.T(), expectedLength, suite.stringQueue.Len())
}

func (suite *QueueTestSuite) TestNewFloatQueue() {
	suite.floatQueue = queue.New[float64]()
	expectedLength := 3

	suite.floatQueue.Enqueue(5.5)
	suite.floatQueue.Enqueue(7.7)
	suite.floatQueue.Enqueue(9.9)
	assert.NotNil(suite.T(), suite.floatQueue)
	assert.Equal(suite.T(), expectedLength, suite.floatQueue.Len())
}

func (suite *QueueTestSuite) TestDequeue() {
	suite.queue = queue.New[int]()
	expectedValue := 5

	suite.queue.Enqueue(5)
	suite.queue.Enqueue(7)
	suite.queue.Enqueue(9)
	actualValue, _ := suite.queue.Dequeue()
	assert.Equal(suite.T(), expectedValue, *actualValue)
}

func (suite *QueueTestSuite) TestDequeueWithError() {
	suite.queue = queue.New[int]()
	_, err := suite.queue.Dequeue()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), queue.ErrEmptyQueue, err)
}

func (suite *QueueTestSuite) TestEmptyQueueIsEmpty() {
	suite.queue = queue.New[int]()
	assert.True(suite.T(), suite.queue.IsEmpty())
}

func (suite *QueueTestSuite) TestEmptyQueueIsNotEmpty() {
	suite.queue = queue.New[int]()

	suite.queue.Enqueue(5)
	suite.queue.Enqueue(7)
	assert.False(suite.T(), suite.queue.IsEmpty())
}

func (suite *QueueTestSuite) TestLen() {
	suite.queue = queue.New[int]()
	expectedLength := 5

	suite.queue.Enqueue(5)
	suite.queue.Enqueue(7)
	suite.queue.Enqueue(9)
	suite.queue.Enqueue(9)
	suite.queue.Enqueue(9)
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}

func (suite *QueueTestSuite) TestLenEmptyQueue() {
	suite.queue = queue.New[int]()
	expectedLength := 0
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}

func (suite *QueueTestSuite) TestLenAfterDequeue() {
	suite.queue = queue.New[int]()
	expectedLength := 4

	suite.queue.Enqueue(5)
	suite.queue.Enqueue(7)
	suite.queue.Enqueue(9)
	suite.queue.Enqueue(9)
	suite.queue.Enqueue(9)
	_, _ = suite.queue.Dequeue()
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}
