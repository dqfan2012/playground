package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/linkedlist"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SingleLinkedListTestSuite struct {
	suite.Suite
	// Demonstrate that the linked list works with multiple data types
	singleList       *linkedlist.Single[int]
	singleStringList *linkedlist.Single[string]
	singleFloatList  *linkedlist.Single[float64]
}

func TestSingleLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(SingleLinkedListTestSuite))
}

func (suite *SingleLinkedListTestSuite) TestNewSingleEmpty() {
	expectedLen := 0
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Nil(suite.T(), suite.singleList.GetHead())
	assert.Nil(suite.T(), suite.singleList.GetTail())
}

func (suite *SingleLinkedListTestSuite) TestNewSingleWithHead() {
	expectedHeadValue := 10
	expectedLen := 1
	suite.singleList = linkedlist.NewSingleWithHead[int](expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestNewSingleStringListWithHead() {
	expectedHeadValue := "Hello"
	expectedLen := 1
	suite.singleStringList = linkedlist.NewSingleWithHead[string](expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.singleStringList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.singleStringList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleStringList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestNewSingleFloatListWithHead() {
	expectedHeadValue := float64(14)
	expectedLen := 1
	suite.singleFloatList = linkedlist.NewSingleWithHead[float64](expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.singleFloatList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.singleFloatList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleFloatList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListInsertHead() {
	expectedHeadValue := 5
	expectedLen := 1
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	inserted := suite.singleList.InsertHead(expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestListWithDataInsertHead() {
	expectedOriginalHeadValue := 17
	expectedNewHeadValue := 44
	expectedLen := 2
	suite.singleList = linkedlist.NewSingleWithHead[int](expectedOriginalHeadValue)
	inserted := suite.singleList.InsertHead(expectedNewHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedNewHeadValue, suite.singleList.GetHead().Data)
	assert.Equal(suite.T(), expectedOriginalHeadValue, suite.singleList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListInsertTail() {
	expectedHeadValue := 9
	expectedLen := 1
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	inserted := suite.singleList.InsertTail(expectedHeadValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.True(suite.T(), inserted)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetHead().Data)
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestListWithDataInsertTail() {
	expectedHeadValue := 103
	expectedTailValue := 37
	expectedLen := 2
	suite.singleList = linkedlist.NewSingleWithHead[int](expectedHeadValue)
	_ = suite.singleList.InsertTail(expectedTailValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedHeadValue, suite.singleList.GetHead().Data)
	assert.Equal(suite.T(), expectedTailValue, suite.singleList.GetTail().Data)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListInsertAtInvalidPosition() {
	expectedLen := 0
	position := 5
	valueNotInserted := 4
	errMsg := fmt.Errorf("%v: %d", linkedlist.ErrInvalidPosition, position)
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	inserted, err := suite.singleList.InsertAtPosition(position, valueNotInserted)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.False(suite.T(), inserted)
	assert.ErrorIs(suite.T(), err, linkedlist.ErrInvalidPosition, errMsg)
}

func (suite *SingleLinkedListTestSuite) TestInsertAtFirstPosition() {
	headValue := 99
	expectedInsertAtValue := 388
	expectedLen := 7
	insertPosition := 1
	suite.singleList = linkedlist.NewSingleWithHead[int](headValue)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	suite.singleList.InsertHead(5)
	suite.singleList.InsertHead(6)
	suite.singleList.InsertHead(7)
	inserted, _ := suite.singleList.InsertAtPosition(insertPosition, expectedInsertAtValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.True(suite.T(), inserted)
}

func (suite *SingleLinkedListTestSuite) TestInsertAtNPosition() {
	headValue := 99
	expectedInsertAtValue := 388
	expectedLen := 7
	insertPosition := 3
	suite.singleList = linkedlist.NewSingleWithHead[int](headValue)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	suite.singleList.InsertHead(5)
	suite.singleList.InsertHead(6)
	suite.singleList.InsertHead(7)
	inserted, _ := suite.singleList.InsertAtPosition(insertPosition, expectedInsertAtValue)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.True(suite.T(), inserted)
}

func (suite *SingleLinkedListTestSuite) TestValueIsPresent() {
	headValue := 1
	expectedLen := 4
	suite.singleList = linkedlist.NewSingleWithHead[int](headValue)
	suite.singleList.InsertHead(2)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	present := suite.singleList.IsValuePresent(2)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.True(suite.T(), present)
}

func (suite *SingleLinkedListTestSuite) TestValueIsNotPresent() {
	headValue := 1
	expectedLen := 4
	suite.singleList = linkedlist.NewSingleWithHead[int](headValue)
	suite.singleList.InsertHead(2)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	present := suite.singleList.IsValuePresent(13)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.False(suite.T(), present)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListDeleteHead() {
	expectedLen := 0
	expectedValue := 0
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	actualValue, deleted := suite.singleList.DeleteHead()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.False(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestListWithDataDeleteHead() {
	expectedLen := 0
	expectedValue := 84
	suite.singleList = linkedlist.NewSingleWithHead[int](expectedValue)
	actualValue, deleted := suite.singleList.DeleteHead()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestBiggerListDeleteHead() {
	expectedLen := 4
	expectedValue := 6
	suite.singleList = linkedlist.NewSingleWithHead[int](1)
	suite.singleList.InsertHead(2)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	suite.singleList.InsertHead(expectedValue)
	actualValue, deleted := suite.singleList.DeleteHead()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListDeleteTail() {
	expectedLen := 0
	expectedValue := 0
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	actualValue, deleted := suite.singleList.DeleteTail()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.False(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestListWithDataDeleteTail() {
	expectedLen := 0
	expectedValue := 84
	suite.singleList = linkedlist.NewSingleWithHead[int](expectedValue)
	actualValue, deleted := suite.singleList.DeleteTail()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestBiggerListDeleteTail() {
	expectedLen := 5
	expectedValue := 84
	suite.singleList = linkedlist.NewSingleWithHead[int](expectedValue)
	suite.singleList.InsertHead(2)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	suite.singleList.InsertHead(5)
	suite.singleList.InsertHead(6)
	actualValue, deleted := suite.singleList.DeleteTail()
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListDeleteAtInvalidPosition() {
	expectedLen := 0
	position := 5
	errMsg := fmt.Errorf("%v: %d", linkedlist.ErrInvalidPosition, position)
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	actualValue, deleted, err := suite.singleList.DeleteAtPosition(position)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Nil(suite.T(), actualValue)
	assert.False(suite.T(), deleted)
	assert.ErrorIs(suite.T(), err, linkedlist.ErrInvalidPosition, errMsg)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListDeleteAtFirstPosition() {
	expectedLen := 5
	position := 0
	expectedValue := 6
	suite.singleList = linkedlist.NewSingleWithHead[int](1)
	suite.singleList.InsertHead(2)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	suite.singleList.InsertHead(5)
	suite.singleList.InsertHead(expectedValue)
	actualValue, deleted, err := suite.singleList.DeleteAtPosition(position)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListDeleteAtNFirstPosition() {
	expectedLen := 5
	position := 4
	expectedValue := 2
	suite.singleList = linkedlist.NewSingleWithHead[int](1)
	suite.singleList.InsertHead(expectedValue)
	suite.singleList.InsertHead(3)
	suite.singleList.InsertHead(4)
	suite.singleList.InsertHead(5)
	suite.singleList.InsertHead(6)
	actualValue, deleted, err := suite.singleList.DeleteAtPosition(position)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLen, suite.singleList.Len())
	assert.Equal(suite.T(), expectedValue, actualValue)
	assert.True(suite.T(), deleted)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListIsEmpty() {
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	isEmpty := suite.singleList.IsEmpty()
	assert.True(suite.T(), isEmpty)
}

func (suite *SingleLinkedListTestSuite) TestListWithDataIsEmpty() {
	newHeadValue := 6
	suite.singleList = linkedlist.NewSingleWithHead[int](newHeadValue)
	isEmpty := suite.singleList.IsEmpty()
	assert.False(suite.T(), isEmpty)
}

func (suite *SingleLinkedListTestSuite) TestEmptyListGetLen() {
	expectedLen := 0
	suite.singleList = linkedlist.NewSingleEmpty[int]()
	actualLen := suite.singleList.Len()
	assert.Equal(suite.T(), expectedLen, actualLen)
}

func (suite *SingleLinkedListTestSuite) TestListWithDataGetLen() {
	newHeadValue := 6
	expectedLen := 1
	suite.singleList = linkedlist.NewSingleWithHead[int](newHeadValue)
	actualLen := suite.singleList.Len()
	assert.Equal(suite.T(), expectedLen, actualLen)
}
