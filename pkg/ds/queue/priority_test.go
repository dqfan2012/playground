package queue_test

import (
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PriorityQueueTestSuite struct {
	suite.Suite
	// Demonstrate that the queue works with multiple data types
	queue       *queue.PriorityQueue[int]
	stringQueue *queue.PriorityQueue[string]
	floatQueue  *queue.PriorityQueue[float64]
}

func TestPriorityQueueTestSuite(t *testing.T) {
	suite.Run(t, new(PriorityQueueTestSuite))
}

func (suite *PriorityQueueTestSuite) TestNewPriorityQueue() {
	suite.queue = queue.NewPriorityQueue[int]()
	item1Value := 5
	item1Priority := 2
	item2Value := 11
	item2Priority := 1
	item3Value := 8
	item3Priority := 3
	expectedLength := 3

	suite.queue.Enqueue(item1Value, item1Priority)
	suite.queue.Enqueue(item2Value, item2Priority)
	suite.queue.Enqueue(item3Value, item3Priority)
	assert.NotNil(suite.T(), suite.queue)
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}

func (suite *PriorityQueueTestSuite) TestNewStringPriorityQueue() {
	suite.stringQueue = queue.NewPriorityQueue[string]()
	item1Value := "Hello"
	item1Priority := 2
	item2Value := "World"
	item2Priority := 1
	item3Value := "Cheese"
	item3Priority := 3
	expectedLength := 3

	suite.stringQueue.Enqueue(item1Value, item1Priority)
	suite.stringQueue.Enqueue(item2Value, item2Priority)
	suite.stringQueue.Enqueue(item3Value, item3Priority)
	assert.NotNil(suite.T(), suite.queue)
	assert.Equal(suite.T(), expectedLength, suite.stringQueue.Len())
}

func (suite *PriorityQueueTestSuite) TestNewFloatPriorityQueue() {
	suite.floatQueue = queue.NewPriorityQueue[float64]()
	item1Value := 3.3
	item1Priority := 2
	item2Value := 7.7
	item2Priority := 1
	item3Value := 11.32
	item3Priority := 3
	expectedLength := 3

	suite.floatQueue.Enqueue(item1Value, item1Priority)
	suite.floatQueue.Enqueue(item2Value, item2Priority)
	suite.floatQueue.Enqueue(item3Value, item3Priority)
	assert.NotNil(suite.T(), suite.queue)
	assert.Equal(suite.T(), expectedLength, suite.floatQueue.Len())
}

func (suite *PriorityQueueTestSuite) TestDequeue() {
	suite.queue = queue.NewPriorityQueue[int]()
	item1Value := 5
	item1Priority := 2
	item2Value := 11
	item2Priority := 1
	item3Value := 8
	item3Priority := 3
	expectedLength := 2

	suite.queue.Enqueue(item1Value, item1Priority)
	suite.queue.Enqueue(item2Value, item2Priority)
	suite.queue.Enqueue(item3Value, item3Priority)

	item, _ := suite.queue.Dequeue()

	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
	assert.Equal(suite.T(), item2Value, item.Value)
}

func (suite *PriorityQueueTestSuite) TestDequeueWithError() {
	suite.queue = queue.NewPriorityQueue[int]()
	_, err := suite.queue.Dequeue()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), queue.ErrEmptyQueue, err)
}

func (suite *PriorityQueueTestSuite) TestEmptyQueueIsEmpty() {
	suite.queue = queue.NewPriorityQueue[int]()
	assert.True(suite.T(), suite.queue.IsEmpty())
}

func (suite *PriorityQueueTestSuite) TestEmptyQueueIsNotEmpty() {
	suite.queue = queue.NewPriorityQueue[int]()
	item1Value := 5
	item1Priority := 2
	item2Value := 11
	item2Priority := 1

	suite.queue.Enqueue(item1Value, item1Priority)
	suite.queue.Enqueue(item2Value, item2Priority)
	assert.False(suite.T(), suite.queue.IsEmpty())
}

func (suite *PriorityQueueTestSuite) TestLen() {
	suite.queue = queue.NewPriorityQueue[int]()
	item1Value := 17
	item1Priority := 2
	item2Value := 3
	item2Priority := 1
	item3Value := 8
	item3Priority := 3
	item4Value := 44
	item4Priority := 4
	item5Value := 66
	item5Priority := 104
	expectedLength := 5

	suite.queue.Enqueue(item1Value, item1Priority)
	suite.queue.Enqueue(item2Value, item2Priority)
	suite.queue.Enqueue(item3Value, item3Priority)
	suite.queue.Enqueue(item4Value, item4Priority)
	suite.queue.Enqueue(item5Value, item5Priority)
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}

func (suite *PriorityQueueTestSuite) TestLenEmptyQueue() {
	suite.queue = queue.NewPriorityQueue[int]()
	expectedLength := 0
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}

func (suite *PriorityQueueTestSuite) TestLenAfterDequeue() {
	suite.queue = queue.NewPriorityQueue[int]()
	item1Value := 5
	item1Priority := 2
	item2Value := 11
	item2Priority := 1
	item3Value := 8
	item3Priority := 3
	expectedLength := 2

	suite.queue.Enqueue(item1Value, item1Priority)
	suite.queue.Enqueue(item2Value, item2Priority)
	suite.queue.Enqueue(item3Value, item3Priority)

	_, _ = suite.queue.Dequeue()
	assert.Equal(suite.T(), expectedLength, suite.queue.Len())
}
