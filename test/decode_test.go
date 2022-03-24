package test

import (
	"github.com/gocollection/go-smile/smile"
	"github.com/gocollection/go-smile/test/testdata"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	filenames, err := testdata.TestFilenames()
	require.NoError(t, err)

	for _, f := range filenames {
		f := f
		t.Run(filepath.Base(f), func(t *testing.T) {
			jsonFile := testdata.LoadTestFile(t, f+".json")
			smileFile := testdata.LoadTestFile(t, f+".smile")

			actualJSON, err := smile.DecodeToJSON(smileFile)
			require.NoError(t, err, "Error while decoding %q", f)

			require.JSONEq(t, string(jsonFile), actualJSON, "Decoding %q didn't produce the expected result", f)

		})
	}
}
