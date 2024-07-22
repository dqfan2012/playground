package stack_test

import (
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StackTestSuite struct {
	suite.Suite
	// Demonstrate that the Stack works with multiple data types
	stack       *stack.Stack[int]
	stringStack *stack.Stack[string]
	floatStack  *stack.Stack[float64]
}

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}

func (suite *StackTestSuite) TestNew() {
	suite.stack = stack.New[int]()
	expectedLength := 3

	suite.stack.Push(3)
	suite.stack.Push(5)
	suite.stack.Push(9)

	assert.NotNil(suite.T(), suite.stack)
	assert.Equal(suite.T(), expectedLength, suite.stack.Len())
}

func (suite *StackTestSuite) TestNewStringStack() {
	suite.stringStack = stack.New[string]()
	expectedLength := 3

	suite.stringStack.Push("Hello,")
	suite.stringStack.Push("World!")
	suite.stringStack.Push("You are beautiful!")

	assert.NotNil(suite.T(), suite.stringStack)
	assert.Equal(suite.T(), expectedLength, suite.stringStack.Len())
}

func (suite *StackTestSuite) TestNewFloatStack() {
	suite.floatStack = stack.New[float64]()
	expectedLength := 3

	suite.floatStack.Push(5.5)
	suite.floatStack.Push(7.7)
	suite.floatStack.Push(12.32)

	assert.NotNil(suite.T(), suite.floatStack)
	assert.Equal(suite.T(), expectedLength, suite.floatStack.Len())
}

func (suite *StackTestSuite) TestPop() {
	suite.stack = stack.New[int]()

	suite.stack.Push(3)
	suite.stack.Push(5)
	suite.stack.Push(9)

	item1, _ := suite.stack.Pop()
	item2, _ := suite.stack.Pop()
	item3, _ := suite.stack.Pop()

	assert.Equal(suite.T(), 9, *item1)
	assert.Equal(suite.T(), 5, *item2)
	assert.Equal(suite.T(), 3, *item3)
}

func (suite *StackTestSuite) TestPopError() {
	suite.stack = stack.New[int]()

	_, err := suite.stack.Pop()

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), stack.ErrEmptyStack, err)
}

func (suite *StackTestSuite) TestPeek() {
	suite.stack = stack.New[int]()

	suite.stack.Push(3)
	suite.stack.Push(5)
	suite.stack.Push(9)

	item1, _ := suite.stack.Peek()

	assert.Equal(suite.T(), 9, *item1)
}

func (suite *StackTestSuite) TestPeekError() {
	suite.stack = stack.New[int]()

	_, err := suite.stack.Peek()

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), stack.ErrEmptyStack, err)
}

func (suite *StackTestSuite) TestEmptyStackIsEmpty() {
	suite.stack = stack.New[int]()
	assert.True(suite.T(), suite.stack.IsEmpty())
}

func (suite *StackTestSuite) TestNotEmptyStackIsEmpty() {
	suite.stack = stack.New[int]()

	suite.stack.Push(3)
	suite.stack.Push(5)
	suite.stack.Push(9)

	assert.False(suite.T(), suite.stack.IsEmpty())
}
