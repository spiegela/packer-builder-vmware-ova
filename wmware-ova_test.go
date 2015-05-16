package main

import (
	"github.com/mitchellh/packer/packer"
	"testing"
)

func TestBuilder_ImplementsBuilder(t *testing.T) {
	var _ packer.Builder = new(Builder)
}
