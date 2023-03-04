package main

import (
   "fmt"
   "github.com/go-git/go-billy/v5/osfs"
   "github.com/go-git/go-git/v5"
   "github.com/go-git/go-git/v5/plumbing"
   "github.com/go-git/go-git/v5/plumbing/cache"
   "github.com/go-git/go-git/v5/plumbing/filemode"
   "github.com/go-git/go-git/v5/plumbing/format/index"
   "github.com/go-git/go-git/v5/plumbing/object"
   "github.com/go-git/go-git/v5/storage/filesystem"
   "github.com/go-git/go-git/v5/utils/diff"
   "github.com/sergi/go-diff/diffmatchpatch"
   "io"
   "os"
   fdiff "github.com/go-git/go-git/v5/plumbing/format/diff"
)

func main() {
   root := `D:\GitHub\git`
   r, err := git.Open(
      filesystem.NewStorage(
         osfs.New(root + "/.git"),
         cache.NewObjectLRUDefault(),
      ),
      osfs.New(root),
   )
   if err != nil {
      panic(err)
   }
   idx, err := r.Storer.Index()
   if err != nil {
      panic(err)
   }
   iter := &indexEntriesIter{
      idx: idx,
      r:   r,
   }
   if err := worktreeDiff(os.Stdout, iter, root); err != nil {
      panic(err)
   }
}

func worktreeDiff(w io.Writer, iter fileIter, root string) error {
   var filePatches []fdiff.FilePatch
   for {
      file, err := iter.Next()
      if err == io.EOF {
         break
      }
      if err != nil {
         return err
      }
      fromContent, err := file.Contents()
      if err != nil {
         return err
      }
      b, err := os.ReadFile(root + "/" + file.Name)
      if err != nil {
         if !os.IsNotExist(err) {
            return err
         }
         b = nil
      }
      toContent := string(b)
      if fromContent != toContent {
         fp, err := fileDiff(os.Stdout, file, fromContent, toContent)
         if err != nil {
            return err
         }
         filePatches = append(filePatches, fp)
      }
   }
   ue := fdiff.NewUnifiedEncoder(w, fdiff.DefaultContextLines)
   return ue.Encode(&gigPatch{
      filePatches: filePatches,
   })
}

type fileIter interface {
   Next() (*object.File, error)
}

type indexEntriesIter struct {
   idx *index.Index
   r   *git.Repository
   k   int
}

func (i *indexEntriesIter) Next() (*object.File, error) {
   entries := i.idx.Entries
   if i.k >= len(entries) {
      return nil, io.EOF
   }
   if i.k < 0 {
      return nil, fmt.Errorf("index %v out of range", i.k)
   }
   e := entries[i.k]
   b, err := i.r.BlobObject(e.Hash)
   if err != nil {
      return nil, err
   }
   i.k++
   return object.NewFile(e.Name, e.Mode, b), nil
}

func fileDiff(w io.Writer, f *object.File, a, b string) (fdiff.FilePatch, error) {
   diffs := diff.Do(a, b)
   var chunks []fdiff.Chunk
   for _, d := range diffs {
      var op fdiff.Operation
      switch d.Type {
      case diffmatchpatch.DiffEqual:
         op = fdiff.Equal
      case diffmatchpatch.DiffDelete:
         op = fdiff.Delete
      case diffmatchpatch.DiffInsert:
         op = fdiff.Add
      }
      chunks = append(chunks, &gigChunk{content: d.Text, op: op})
   }

   isBinary, err := f.IsBinary()
   if err != nil {
      return nil, err
   }
   fp := &gigFilePatch{
      isBinary: isBinary,
      from: &gigFile{
         hash: f.Hash,
         mode: f.Mode,
         path: f.Name,
      },
      to: &gigFile{
         hash: f.Hash, // TODO
         mode: f.Mode, // TODO
         path: f.Name,
      },
      chunks: chunks,
   }
   return fp, nil
}

type gigPatch struct {
   message     string
   filePatches []fdiff.FilePatch
}

func (p *gigPatch) FilePatches() []fdiff.FilePatch { return p.filePatches }
func (p *gigPatch) Message() string                { return p.message }

type gigFilePatch struct {
   isBinary bool
   from, to *gigFile
   chunks   []fdiff.Chunk
}

func (fp *gigFilePatch) IsBinary() bool               { return fp.isBinary }
func (fp *gigFilePatch) Files() (from, to fdiff.File) { return fp.from, fp.to }
func (fp *gigFilePatch) Chunks() []fdiff.Chunk        { return fp.chunks }

type gigFile struct {
   hash plumbing.Hash
   mode filemode.FileMode
   path string
}

func (f *gigFile) Hash() plumbing.Hash     { return f.hash }
func (f *gigFile) Mode() filemode.FileMode { return f.mode }
func (f *gigFile) Path() string            { return f.path }

type gigChunk struct {
   content string
   op      fdiff.Operation
}

func (c *gigChunk) Content() string       { return c.content }
func (c *gigChunk) Type() fdiff.Operation { return c.op }
