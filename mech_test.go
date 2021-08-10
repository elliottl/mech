package mech

import (
   "fmt"
   "os"
   "strings"
   "testing"
)

func TestScan(t *testing.T) {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   s := mech.NewScanner(f)
   s.ScanAttr("type", "application/ld+json")
   fmt.Println(s.Attr("data-rh"))
   s.ScanText()
   fmt.Println(s.Data)
}
