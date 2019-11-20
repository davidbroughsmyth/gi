// Code generated by "stringer -type=FileNodeVcsStates"; DO NOT EDIT.

package giv

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FileNodeNotInVcs-0]
	_ = x[FileNodeVcsAdded-1]
	_ = x[FileNodeInVcs-2]
	_ = x[FileNodeVcsModified-3]
	_ = x[FileNodeVcsStatesN-4]
}

const _FileNodeVcsStates_name = "FileNodeNotInVcsFileNodeVcsAddedFileNodeInVcsFileNodeVcsModifiedFileNodeVcsStatesN"

var _FileNodeVcsStates_index = [...]uint8{0, 16, 32, 45, 64, 82}

func (i FileNodeVcsStates) String() string {
	if i < 0 || i >= FileNodeVcsStates(len(_FileNodeVcsStates_index)-1) {
		return "FileNodeVcsStates(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FileNodeVcsStates_name[_FileNodeVcsStates_index[i]:_FileNodeVcsStates_index[i+1]]
}

func (i *FileNodeVcsStates) FromString(s string) error {
	for j := 0; j < len(_FileNodeVcsStates_index)-1; j++ {
		if s == _FileNodeVcsStates_name[_FileNodeVcsStates_index[j]:_FileNodeVcsStates_index[j+1]] {
			*i = FileNodeVcsStates(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FileNodeVcsStates")
}
