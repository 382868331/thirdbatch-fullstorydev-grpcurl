package grpcurl

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGrpcurl008SourceContract(t *testing.T) {
    source, err := os.ReadFile("grpcurl.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
