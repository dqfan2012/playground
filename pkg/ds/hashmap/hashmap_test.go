package hashmap_test

import (
	"testing"

	"github.com/dqfan2012/playground/pkg/ds/hashmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HashMapTestSuite struct {
	suite.Suite
	// Demonstrate that the HashMap works with multiple data types
	hashmap       *hashmap.HashMap[int]
	stringHashmap *hashmap.HashMap[string]
	floatHashmap  *hashmap.HashMap[float64]
}

func TestHashMapTestSuite(t *testing.T) {
	suite.Run(t, new(HashMapTestSuite))
}

func (suite *HashMapTestSuite) TestNew() {
	mapSize := 3
	suite.hashmap = hashmap.NewHashMap[int](mapSize)

	suite.hashmap.Insert("one", 1)
	suite.hashmap.Insert("two", 2)
	suite.hashmap.Insert("three", 3)

	assert.NotNil(suite.T(), suite.hashmap)
}

func (suite *HashMapTestSuite) TestNewStringHashmap() {
	mapSize := 2
	suite.stringHashmap = hashmap.NewHashMap[string](mapSize)

	suite.stringHashmap.Insert("first", "Hello")
	suite.stringHashmap.Insert("second", "World")

	assert.NotNil(suite.T(), suite.stringHashmap)
}

func (suite *HashMapTestSuite) TestNewFloatHashmap() {
	mapSize := 2
	suite.floatHashmap = hashmap.NewHashMap[float64](mapSize)

	suite.floatHashmap.Insert("pi", 3.14)
	suite.floatHashmap.Insert("e", 2.718)

	assert.NotNil(suite.T(), suite.floatHashmap)
}

func (suite *HashMapTestSuite) TestGet() {
	mapSize := 3
	suite.hashmap = hashmap.NewHashMap[int](mapSize)

	suite.hashmap.Insert("one", 1)
	suite.hashmap.Insert("two", 2)
	suite.hashmap.Insert("three", 3)

	value, exists := suite.hashmap.Get("two")

	assert.True(suite.T(), exists)
	assert.Equal(suite.T(), 2, *value)

	value, exists = suite.hashmap.Get("four")

	assert.False(suite.T(), exists)
	assert.Nil(suite.T(), value)
}

func (suite *HashMapTestSuite) TestDelete() {
	mapSize := 3
	suite.hashmap = hashmap.NewHashMap[int](mapSize)

	suite.hashmap.Insert("one", 1)
	suite.hashmap.Insert("two", 2)
	suite.hashmap.Insert("three", 3)

	suite.hashmap.Delete("two")

	value, exists := suite.hashmap.Get("two")
	assert.False(suite.T(), exists)
	assert.Nil(suite.T(), value)

	value, exists = suite.hashmap.Get("one")
	assert.True(suite.T(), exists)
	assert.Equal(suite.T(), 1, *value)
}

func (suite *HashMapTestSuite) TestUpdate() {
	mapSize := 3
	suite.hashmap = hashmap.NewHashMap[int](mapSize)

	suite.hashmap.Insert("one", 1)
	suite.hashmap.Insert("two", 2)

	suite.hashmap.Insert("two", 22)

	value, exists := suite.hashmap.Get("two")
	assert.True(suite.T(), exists)
	assert.Equal(suite.T(), 22, *value)
}

func (suite *HashMapTestSuite) TestCollision() {
	mapSize := 3
	suite.hashmap = hashmap.NewHashMap[int](mapSize)

	// Keys "one" and "three" may result in the same index
	suite.hashmap.Insert("one", 1)
	suite.hashmap.Insert("three", 3)

	value, exists := suite.hashmap.Get("one")
	assert.True(suite.T(), exists)
	assert.Equal(suite.T(), 1, *value)

	value, exists = suite.hashmap.Get("three")
	assert.True(suite.T(), exists)
	assert.Equal(suite.T(), 3, *value)
}

func (suite *HashMapTestSuite) TestEmptyHashMap() {
	mapSize := 3
	suite.hashmap = hashmap.NewHashMap[int](mapSize)

	value, exists := suite.hashmap.Get("nonexistent")
	assert.False(suite.T(), exists)
	assert.Nil(suite.T(), value)
}
