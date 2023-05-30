package main

import (
	"terratest-action/terratest"
	"testing"
)

var t testing.T

func main() {
	terratest.TestAzure(&t)
}
