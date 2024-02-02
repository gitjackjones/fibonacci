package main

import (
	"testing"
)

// Tests that checkNumber() works as expected
func TestCheckNumber(t *testing.T) {
    tests := []struct {
        number    string
        expectedVal int
        expectedError bool
    }{
        {"5", 5, false},
        {"-1", 0, true},
        {"101", 0, true},
        {"jackisnotanumber", 0, true},
        {"5.5", 0, true},
    }

    for _, tc := range tests {
        result, err := checkNumber(tc.number) 
        gotError := err != nil 
        // If there was an error, report it
        if gotError != tc.expectedError { 
            t.Errorf("checkNumber(%s): expected error=%v, got error=%v", tc.number, tc.expectedError, gotError)
        } else if !gotError && result != tc.expectedVal { 
            t.Errorf("checkNumber(%s): expected result=%v, got result=%v", tc.number, tc.expectedVal, result)
        }
    }
}
