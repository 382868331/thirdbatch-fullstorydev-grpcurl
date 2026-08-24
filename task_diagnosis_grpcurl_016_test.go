package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGrpcurl016SourceContract(t *testing.T) {
    source, err := os.ReadFile("desc_source.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer root.Close()") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "root.Close()") {
        t.Fatalf("mutated source contract is still present")
    }
}
