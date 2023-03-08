package diff

import (
   "github.com/go-git/go-git/v5/plumbing"
   "github.com/go-git/go-git/v5/plumbing/filemode"
   "github.com/go-git/go-git/v5/plumbing/format/diff"
)

func (chunk) Content() string {
   return ""
}

type patch struct{}

func (patch) Message() string {
   return ""
}

func (patch) FilePatches() []diff.FilePatch {
   return nil
}

type file_patch struct{}

func (file_patch) Chunks() []diff.Chunk {
   return nil
}

func (file_patch) Files() (diff.File, diff.File) {
   return nil, nil
}

func (file_patch) IsBinary() bool {
   return false
}

type chunk struct{}

func (chunk) Type() diff.Operation {
   return 0
}

type file struct{}

func (file) Hash() plumbing.Hash {
   return [20]byte{}
}

func (file) Mode() filemode.FileMode {
   return 0
}

func (file) Path() string {
   return ""
}
