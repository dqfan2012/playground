package set_test

import (
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SetTestSuite struct {
	suite.Suite
	// Demonstrate that the queue works with multiple data types
	set       *set.Set[int]
	stringSet *set.Set[string]
	floatSet  *set.Set[float64]
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}

func (suite *SetTestSuite) TestNew() {
	suite.set = set.New[int]()
	expectedLength := 3

	suite.set.Add(3)
	suite.set.Add(5)
	suite.set.Add(9)

	assert.NotNil(suite.T(), suite.set)
	assert.Equal(suite.T(), expectedLength, suite.set.Len())
}

func (suite *SetTestSuite) TestNewStringSet() {
	suite.stringSet = set.New[string]()
	expectedLength := 2

	suite.stringSet.Add("Hello")
	suite.stringSet.Add("World")

	assert.NotNil(suite.T(), suite.stringSet)
	assert.Equal(suite.T(), expectedLength, suite.stringSet.Len())
}

func (suite *SetTestSuite) TestNewFloatSet() {
	suite.floatSet = set.New[float64]()
	expectedLength := 3

	suite.floatSet.Add(3.1)
	suite.floatSet.Add(5.207)
	suite.floatSet.Add(9.88)

	assert.NotNil(suite.T(), suite.floatSet)
	assert.Equal(suite.T(), expectedLength, suite.floatSet.Len())
}

func (suite *SetTestSuite) TestRemove() {
	suite.set = set.New[int]()
	expectedLength := 1

	suite.set.Add(3)
	suite.set.Add(5)
	suite.set.Add(9)

	suite.set.Remove(3)
	suite.set.Remove(5)

	assert.Equal(suite.T(), expectedLength, suite.set.Len())
}

func (suite *SetTestSuite) TestContains() {
	suite.set = set.New[int]()
	expectedLength := 4

	suite.set.Add(3)
	suite.set.Add(5)
	suite.set.Add(9)
	suite.set.Add(12)
	suite.set.Add(14)
	suite.set.Add(18)

	suite.set.Remove(3)
	suite.set.Remove(5)

	assert.True(suite.T(), suite.set.Contains(9))
	assert.True(suite.T(), suite.set.Contains(12))
	assert.False(suite.T(), suite.set.Contains(3))
	assert.False(suite.T(), suite.set.Contains(5))

	assert.Equal(suite.T(), expectedLength, suite.set.Len())
}

func (suite *SetTestSuite) TestUnion() {
	suite.set = set.New[int]()

	suite.set.Add(3)
	suite.set.Add(5)
	suite.set.Add(9)

	setB := set.New[int]()

	setB.Add(7)
	setB.Add(9)
	setB.Add(12)
	setB.Add(14)

	expectedSet := set.New[int]()

	expectedSet.Add(3)
	expectedSet.Add(5)
	expectedSet.Add(7)
	expectedSet.Add(9)
	expectedSet.Add(12)
	expectedSet.Add(14)

	actualSet := suite.set.Union(setB)

	assert.Equal(suite.T(), expectedSet, actualSet)
}

func (suite *SetTestSuite) TestIntersection() {
	suite.set = set.New[int]()

	suite.set.Add(3)
	suite.set.Add(5)
	suite.set.Add(7)
	suite.set.Add(9)

	setB := set.New[int]()

	setB.Add(7)
	setB.Add(9)
	setB.Add(12)
	setB.Add(14)

	expectedSet := set.New[int]()

	expectedSet.Add(7)
	expectedSet.Add(9)

	actualSet := suite.set.Intersection(setB)

	assert.Equal(suite.T(), expectedSet, actualSet)
}

func (suite *SetTestSuite) TestDifference() {
	suite.set = set.New[int]()

	suite.set.Add(3)
	suite.set.Add(5)
	suite.set.Add(7)
	suite.set.Add(9)

	setB := set.New[int]()

	setB.Add(7)
	setB.Add(9)
	setB.Add(12)
	setB.Add(14)

	expectedSet := set.New[int]()

	expectedSet.Add(3)
	expectedSet.Add(5)
	expectedSet.Add(12)
	expectedSet.Add(14)

	actualSet := suite.set.Difference(setB)

	assert.Equal(suite.T(), expectedSet, actualSet)
}

func (suite *SetTestSuite) TestIsSubset() {
	suite.set = set.New[int]()

	suite.set.Add(1)
	suite.set.Add(2)
	suite.set.Add(3)
	suite.set.Add(4)
	suite.set.Add(5)
	suite.set.Add(6)
	suite.set.Add(7)
	suite.set.Add(8)

	subset1 := set.New[int]()
	subset2 := set.New[int]()
	subset3 := set.New[int]()
	notSubset1 := set.New[int]()
	notSubset2 := set.New[int]()

	// subset1 is a subset of suite.set
	subset1.Add(3)
	subset1.Add(4)
	subset2.Add(5)

	// subset2 is a subset of suite.set
	subset2.Add(6)
	subset2.Add(7)
	subset2.Add(8)

	// subset3 is a subset of suite.set
	subset3.Add(1)
	subset3.Add(5)
	subset3.Add(7)

	// notSubest1 is not a subset of suite.set
	notSubset1.Add(1)
	notSubset1.Add(5)
	notSubset1.Add(9)

	// notSubset2 is not a subset of suite.set
	notSubset2.Add(3)
	notSubset2.Add(4)
	notSubset2.Add(0)

	assert.Equal(suite.T(), true, suite.set.IsSubset(subset1))
	assert.Equal(suite.T(), true, suite.set.IsSubset(subset2))
	assert.Equal(suite.T(), true, suite.set.IsSubset(subset3))
	assert.Equal(suite.T(), false, suite.set.IsSubset(notSubset1))
	assert.Equal(suite.T(), false, suite.set.IsSubset(notSubset2))
}

func (suite *SetTestSuite) TestClear() {
	suite.set = set.New[int]()
	expectedLength := 0

	suite.set.Add(1)
	suite.set.Add(2)
	suite.set.Add(3)
	suite.set.Add(4)
	suite.set.Add(5)
	suite.set.Add(6)
	suite.set.Add(7)
	suite.set.Add(8)

	suite.set.Clear()

	assert.Equal(suite.T(), expectedLength, suite.set.Len())
}

func (suite *SetTestSuite) TestEmptySetIsEmpty() {
	suite.set = set.New[int]()

	assert.True(suite.T(), suite.set.IsEmpty())
}

func (suite *SetTestSuite) TestNotEmptySetIsEmpty() {
	suite.set = set.New[int]()

	suite.set.Add(1)
	suite.set.Add(2)
	suite.set.Add(3)
	suite.set.Add(4)
	suite.set.Add(5)
	suite.set.Add(6)
	suite.set.Add(7)
	suite.set.Add(8)

	assert.False(suite.T(), suite.set.IsEmpty())
}

func (suite *SetTestSuite) TestLen() {
	suite.set = set.New[int]()
	expectedLength := 8

	suite.set.Add(1)
	suite.set.Add(2)
	suite.set.Add(3)
	suite.set.Add(4)
	suite.set.Add(5)
	suite.set.Add(6)
	suite.set.Add(7)
	suite.set.Add(8)

	assert.Equal(suite.T(), expectedLength, suite.set.Len())
}

func (suite *SetTestSuite) TestEmptySetLen() {
	suite.set = set.New[int]()
	expectedLength := 0

	assert.Equal(suite.T(), expectedLength, suite.set.Len())
}
