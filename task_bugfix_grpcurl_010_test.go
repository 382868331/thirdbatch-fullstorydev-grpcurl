package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGrpcurl010SourceContract(t *testing.T) {
    source, err := os.ReadFile("grpcurl.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "envVarName := result[2 : len(result)-1] // strip leading `${` and trailing `}`") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "envVarName := result[2 : len(result)+ 1] // strip leading `${` and trailing `}`") {
        t.Fatalf("mutated source contract is still present")
    }
}
