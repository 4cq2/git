package diff

import (
   "github.com/go-git/go-git/v5/plumbing"
   "github.com/go-git/go-git/v5/plumbing/filemode"
   "github.com/go-git/go-git/v5/plumbing/format/diff"
)

type chunk struct{}

var _ diff.Chunk = chunk{}

func (chunk) Content() string {
   return ""
}

func (chunk) Type() diff.Operation {
   return 0
}

type file struct{}

var _ diff.File = file{}

func (file) Hash() plumbing.Hash {
   return [20]byte{}
}

func (file) Mode() filemode.FileMode {
   return 0
}

func (file) Path() string {
   return ""
}

type file_patch struct{}

var _ diff.FilePatch = file_patch{}

func (file_patch) Chunks() []diff.Chunk {
   return nil
}

func (file_patch) Files() (diff.File, diff.File) {
   return nil, nil
}

func (file_patch) IsBinary() bool {
   return false
}

type patch struct{}

var _ diff.Patch = patch{}

func (patch) FilePatches() []diff.FilePatch {
   return nil
}

func (patch) Message() string {
   return ""
}
