package code

import (
  "testing"
  "github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	file_name := "testdata/regular"
	file_size := "21B"

	testing_size, _ := GetPathSize(file_name)
	
	require.Equal(t, file_size, testing_size, "File sizes should be equal") 

}

func TestGetPathSize_Dir(t *testing.T) {
	dir_name := "testdata/"
	dir_size := "41B"

	testing_size, _ := GetPathSize(dir_name)
	
	require.Equal(t, dir_size, testing_size, "Dir sizes should be equal")
}

