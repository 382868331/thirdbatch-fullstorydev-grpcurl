package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGrpcurl012SourceContract(t *testing.T) {
    source, err := os.ReadFile("format.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if h.VerbosityLevel > 1 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if h.VerbosityLevel >= 1 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
