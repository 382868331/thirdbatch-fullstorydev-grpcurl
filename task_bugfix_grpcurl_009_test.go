package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrpcurl009SourceContract(t *testing.T) {
    source, err := os.ReadFile("format.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(b) > 0 && b[len(b)-1] == textSeparatorChar {") {
        t.Fatalf("expected source contract is missing")
    }
}
