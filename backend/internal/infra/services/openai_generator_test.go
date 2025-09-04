package services

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var generator *OpenAIGenerator

// TestMain runs before all tests
func TestMain(m *testing.M) {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		println("⚠️  OPENAI_API_KEY is not set, skipping OpenAI tests")
		os.Exit(0) // skip tests if no key
	}

	generator = NewOpenAIGenerator(apiKey, "gpt-3.5-turbo")

	os.Exit(m.Run())
}

func TestGenerateSummary(t *testing.T) {
	// Arrange
	rawPost := `
Go (or Golang) is an open-source programming language designed at Google.
It is statically typed, compiled, and known for simplicity, concurrency,
and excellent performance. It is widely used for backend systems,
cloud-native applications, and developer tooling.
`

	// Act
	summary, err := generator.GenerateSummary(context.Background(), rawPost)

	// Assert
	if err != nil {
		t.Fatalf("GenerateSummary failed: %v", err)
	}
	if summary == "" {
		t.Fatal("expected non-empty summary")
	}

	t.Logf("✅ Summary: %s", summary)
}

func TestGenerateTags(t *testing.T) {
	// Arrange
	rawPost := `
Go (or Golang) is an open-source programming language designed at Google.
It is statically typed, compiled, and known for simplicity, concurrency,
and excellent performance. It is widely used for backend systems,
cloud-native applications, and developer tooling.
`

	// Act
	tags, err := generator.GenerateTags(context.Background(), rawPost)

	// Assert
	if err != nil {
		t.Fatalf("GenerateTags failed: %v", err)
	}
	if len(tags) == 0 {
		t.Fatal("expected at least one tag")
	}

	t.Logf("✅ Tags: %v", tags)
}
