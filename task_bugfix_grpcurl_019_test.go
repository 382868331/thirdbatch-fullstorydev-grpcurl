package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrpcurl019SourceContract(t *testing.T) {
    source, err := os.ReadFile("format.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if h.VerbosityLevel > 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
