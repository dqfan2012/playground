package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/linkedlist"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DoubleLinkedListTestSuite struct {
	suite.Suite
	// Demonstrate that the linked list works with multiple data types
	doubleList       *linkedlist.Double[int]
	doubleStringList *linkedlist.Double[string]
	doubleFloatList  *linkedlist.Double[float64]
}

func TestDoubleLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(DoubleLinkedListTestSuite))
}

func (suite *DoubleLinkedListTestSuite) TestNewDoubleEmpty() {
	expectedLen := 0
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Nil(suite.T(), suite.doubleList.GetHead())
	assert.Nil(suite.T(), suite.doubleList.GetTail())
}

func (suite *DoubleLinkedListTestSuite) TestNewDoubleWithHead() {
	expectedHeadValue := 10
	expectedLen := 1
	suite.doubleList = linkedlist.NewDoubleWithHead[int](expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetTail().Data)
}

func (suite *DoubleLinkedListTestSuite) TestNewDoubleStringListWithHead() {
	expectedHeadValue := "Hello"
	expectedLen := 1
	suite.doubleStringList = linkedlist.NewDoubleWithHead[string](expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleStringList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleStringList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleStringList.GetTail().Data)
}

func (suite *DoubleLinkedListTestSuite) TestNewDoubleFloatListWithHead() {
	expectedHeadValue := float64(14)
	expectedLen := 1
	suite.doubleFloatList = linkedlist.NewDoubleWithHead[float64](expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleFloatList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleFloatList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleFloatList.GetTail().Data)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListInsertHead() {
	expectedHeadValue := 5
	expectedLen := 1
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	inserted := suite.doubleList.InsertHead(expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetTail().Data)
}

func (suite *DoubleLinkedListTestSuite) TestListWithDataInsertHead() {
	expectedOriginalHeadValue := 17
	expectedNewHeadValue := 44
	expectedLen := 2
	suite.doubleList = linkedlist.NewDoubleWithHead[int](expectedOriginalHeadValue)
	inserted := suite.doubleList.InsertHead(expectedNewHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedNewHeadValue, suite.doubleList.GetHead().Data)
	assert.Equal(suite.T(), expectedOriginalHeadValue, suite.doubleList.GetTail().Data)
	assert.Equal(suite.T(), suite.doubleList.GetHead().Next, suite.doubleList.GetTail())
	assert.Equal(suite.T(), suite.doubleList.GetTail().Prev, suite.doubleList.GetHead())
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListInsertTail() {
	expectedHeadValue := 9
	expectedLen := 1
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	inserted := suite.doubleList.InsertTail(expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetTail().Data)
}

func (suite *DoubleLinkedListTestSuite) TestListWithDataInsertTail() {
	expectedHeadValue := 103
	expectedTailValue := 37
	expectedLen := 2
	suite.doubleList = linkedlist.NewDoubleWithHead[int](expectedHeadValue)
	inserted := suite.doubleList.InsertTail(expectedTailValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedHeadValue, suite.doubleList.GetHead().Data)
	assert.Equal(suite.T(), expectedTailValue, suite.doubleList.GetTail().Data)
	assert.Equal(suite.T(), suite.doubleList.GetHead().Next, suite.doubleList.GetTail())
	assert.Equal(suite.T(), suite.doubleList.GetTail().Prev, suite.doubleList.GetHead())
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListInsertAtInvalidPosition() {
	expectedLen := 0
	position := 5
	valueNotInserted := 4
	errMsg := fmt.Errorf("%v: %d", linkedlist.ErrInvalidPosition, position)
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	inserted, err := suite.doubleList.InsertAtPosition(position, valueNotInserted)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.False(suite.T(), inserted)
	assert.ErrorIs(suite.T(), err, linkedlist.ErrInvalidPosition, errMsg)
}

func (suite *DoubleLinkedListTestSuite) TestInsertAtFirstPosition() {
	headValue := 99
	expectedInsertAtValue := 388
	expectedLen := 7
	insertPosition := 1
	suite.doubleList = linkedlist.NewDoubleWithHead[int](headValue)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	suite.doubleList.InsertHead(5)
	suite.doubleList.InsertHead(6)
	suite.doubleList.InsertHead(7)
	inserted, _ := suite.doubleList.InsertAtPosition(insertPosition, expectedInsertAtValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), inserted)
}

func (suite *DoubleLinkedListTestSuite) TestInsertAtNPosition() {
	headValue := 99
	expectedInsertAtValue := 388
	expectedLen := 7
	insertPosition := 3
	suite.doubleList = linkedlist.NewDoubleWithHead[int](headValue)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	suite.doubleList.InsertHead(5)
	suite.doubleList.InsertHead(6)
	suite.doubleList.InsertHead(7)
	inserted, _ := suite.doubleList.InsertAtPosition(insertPosition, expectedInsertAtValue)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), inserted)
}

func (suite *DoubleLinkedListTestSuite) TestValueIsPresent() {
	headValue := 1
	expectedLen := 4
	suite.doubleList = linkedlist.NewDoubleWithHead[int](headValue)
	suite.doubleList.InsertHead(2)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	present := suite.doubleList.IsValuePresent(2)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.True(suite.T(), present)
}

func (suite *DoubleLinkedListTestSuite) TestValueIsNotPresent() {
	headValue := 1
	expectedLen := 4
	suite.doubleList = linkedlist.NewDoubleWithHead[int](headValue)
	suite.doubleList.InsertHead(2)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	present := suite.doubleList.IsValuePresent(13)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.False(suite.T(), present)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListDeleteHead() {
	expectedLen := 0
	expectedValue := 0
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	actualValue, deleted := suite.doubleList.DeleteHead()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.False(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestListWithDataDeleteHead() {
	expectedLen := 0
	expectedValue := 84
	suite.doubleList = linkedlist.NewDoubleWithHead[int](expectedValue)
	actualValue, deleted := suite.doubleList.DeleteHead()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestBiggerListDeleteHead() {
	expectedLen := 4
	expectedValue := 6
	suite.doubleList = linkedlist.NewDoubleWithHead[int](1)
	suite.doubleList.InsertHead(2)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	suite.doubleList.InsertHead(expectedValue)
	actualValue, deleted := suite.doubleList.DeleteHead()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListDeleteTail() {
	expectedLen := 0
	expectedValue := 0
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	actualValue, deleted := suite.doubleList.DeleteTail()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.False(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestListWithDataDeleteTail() {
	expectedLen := 0
	expectedValue := 84
	suite.doubleList = linkedlist.NewDoubleWithHead[int](expectedValue)
	actualValue, deleted := suite.doubleList.DeleteTail()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestBiggerListDeleteTail() {
	expectedLen := 5
	expectedValue := 84
	suite.doubleList = linkedlist.NewDoubleWithHead[int](expectedValue)
	suite.doubleList.InsertHead(2)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	suite.doubleList.InsertHead(5)
	suite.doubleList.InsertHead(6)
	actualValue, deleted := suite.doubleList.DeleteTail()
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListDeleteAtInvalidPosition() {
	expectedLen := 0
	position := 5
	errMsg := fmt.Errorf("%v: %d", linkedlist.ErrInvalidPosition, position)
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	actualValue, deleted, err := suite.doubleList.DeleteAtPosition(position)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Nil(suite.T(), actualValue)
	assert.False(suite.T(), deleted)
	assert.ErrorIs(suite.T(), err, linkedlist.ErrInvalidPosition, errMsg)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListDeleteAtFirstPosition() {
	expectedLen := 5
	position := 0
	expectedValue := 6
	suite.doubleList = linkedlist.NewDoubleWithHead[int](1)
	suite.doubleList.InsertHead(2)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	suite.doubleList.InsertHead(5)
	suite.doubleList.InsertHead(expectedValue)
	actualValue, deleted, err := suite.doubleList.DeleteAtPosition(position)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListDeleteAtNFirstPosition() {
	expectedLen := 5
	position := 4
	expectedValue := 2
	suite.doubleList = linkedlist.NewDoubleWithHead[int](1)
	suite.doubleList.InsertHead(expectedValue)
	suite.doubleList.InsertHead(3)
	suite.doubleList.InsertHead(4)
	suite.doubleList.InsertHead(5)
	suite.doubleList.InsertHead(6)
	actualValue, deleted, err := suite.doubleList.DeleteAtPosition(position)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLen, suite.doubleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListIsEmpty() {
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	isEmpty := suite.doubleList.IsEmpty()
	assert.True(suite.T(), isEmpty)
}

func (suite *DoubleLinkedListTestSuite) TestListWithDataIsEmpty() {
	newHeadValue := 6
	suite.doubleList = linkedlist.NewDoubleWithHead[int](newHeadValue)
	isEmpty := suite.doubleList.IsEmpty()
	assert.False(suite.T(), isEmpty)
}

func (suite *DoubleLinkedListTestSuite) TestEmptyListGetLen() {
	expectedLen := 0
	suite.doubleList = linkedlist.NewDoubleEmpty[int]()
	actualLen := suite.doubleList.Len()
	assert.Equal(suite.T(), expectedLen, actualLen)
}

func (suite *DoubleLinkedListTestSuite) TestListWithDataGetLen() {
	newHeadValue := 6
	expectedLen := 1
	suite.doubleList = linkedlist.NewDoubleWithHead[int](newHeadValue)
	actualLen := suite.doubleList.Len()
	assert.Equal(suite.T(), expectedLen, actualLen)
}
