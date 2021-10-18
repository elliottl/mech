package bandcamp

import (
   "fmt"
   "testing"
)

const band = "https://schnaussandmunk.bandcamp.com"

func TestBand(t *testing.T) {
   Verbose(true)
   typ, id, err := Head(band)
   if err != nil {
      t.Fatal(err)
   }
   if typ != 'i' {
      t.Fatal(typ)
   }
   if id != 3454424886 {
      t.Fatal(id)
   }
   b, err := NewBand(id)
   if err != nil {
      t.Fatal(err)
   }
   a, err := b.Discography[0].Tralbum()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}
