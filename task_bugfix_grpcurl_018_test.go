package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrpcurl018SourceContract(t *testing.T) {
    source, err := os.ReadFile("desc_source.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "i++") {
        t.Fatalf("expected source contract is missing")
    }
}
