package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/vimeo"
   "os"
   "path"
)

func main() {
   var (
      height int
      info bool
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.IntVar(&height, "h", 720, "height")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("vimeo [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   err := vimeo.ValidID(id)
   if err != nil {
      panic(err)
   }
   vimeo.Verbose = true
   cfg, err := vimeo.NewConfig(id)
   if err != nil {
      panic(err)
   }
   // info
   if info {
      for _, f := range cfg.Request.Files.Progressive {
         fmt.Printf("%+v\n", f)
      }
      return
   }
   // download
   for _, f := range cfg.Request.Files.Progressive {
      if f.Height == height {
         download(cfg, f.URL)
      }
   }
}

func download(cfg *vimeo.Config, addr string) {
   name := cfg.Video.Owner.Name + "-" + cfg.Video.Title + path.Ext(addr)
   fmt.Println(name)
}
