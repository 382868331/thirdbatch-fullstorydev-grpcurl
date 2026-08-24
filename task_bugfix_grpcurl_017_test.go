package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrpcurl017SourceContract(t *testing.T) {
    source, err := os.ReadFile("grpcurl.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if firstError == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if firstError != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
