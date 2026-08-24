package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrpcurl011SourceContract(t *testing.T) {
    source, err := os.ReadFile("grpcurl.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if part == \"\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
