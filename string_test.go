package main

import (
	"github.com/magiconair/properties/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings_Len(t *testing.T) {

	length := 1

	w := &Strings{[]string{"Rusya"}}

	length_jack := w.Len()

	assert.Equal(t,length_jack,length)
}

func TestStrings_Swap(t *testing.T) {

	s3 := "a"

	s := &Strings{[]string{"c", "d", "b", "a"}}

	s.Swap(3,0)

	assert.Equal(t,s.MyString[0], s3)
}

func TestStrings_Less(t *testing.T) {

	a := true

	s := &Strings{[]string{"c", "d", "b", "a"}}

	check := s.Less(3,0)

	assert.Equal(t,a,check)
}

