package supabase_test

import(
	"context"
	"leet-term/supabase"
	"leet-term/types"
	"testing"
)

func TestLanguageTable_FindAll(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		opts    []supabase.QueryOpt
		want    []types.Language
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var l supabase.LanguageTable
			got, gotErr := l.FindAll(context.Background(), tt.opts...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("FindAll() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("FindAll() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

