package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/roku"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
)

func (d downloader) DASH(video, audio int64) error {
   if d.info {
      fmt.Println(d.Content)
   }
   video_dash := d.Content.DASH()
   res, err := roku.Client.Redirect(nil).Get(video_dash.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = res.Request.URL
   if err := xml.NewDecoder(res.Body).Decode(&d.media); err != nil {
      return err
   }
   reps := d.media.Representations().Codecs("mp4a")
   if err := d.download(audio, reps); err != nil {
      return err
   }
   reps = d.media.Representations().Codecs("avc1")
   return d.download(video, reps)
}

func (d *downloader) set_key() error {
   private_key, err := os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(d.client)
   if err != nil {
      return err
   }
   raw_key_id := d.media.Representations()[0].ContentProtection.Default_KID
   key_ID, err := widevine.Key_ID(raw_key_id)
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      return err
   }
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   play, err := site.Playback(d.Meta.ID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(play)
   if err != nil {
      return err
   }
   d.key = keys.Content().Key
   return nil
}

func (d *downloader) download(bandwidth int64, r dash.Representations) error {
   if bandwidth == 0 {
      return nil
   }
   rep := r.Get_Bandwidth(bandwidth)
   if d.info {
      for _, each := range r {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      if d.key == nil {
         err := d.set_key()
         if err != nil {
            return err
         }
      }
      file, err := os.Create(d.Content.Base()+rep.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      initial, err := d.url.Parse(rep.Initialization())
      if err != nil {
         return err
      }
      res, err := roku.Client.Get(initial.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      media := rep.Media()
      pro := os.Progress_Chunks(file, len(media))
      dec := mp4.New_Decrypt(pro)
      if err := dec.Init(res.Body); err != nil {
         return err
      }
      for _, addr := range media {
         addr, err := d.url.Parse(addr)
         if err != nil {
            return err
         }
         res, err := roku.Client.Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if err := dec.Segment(res.Body, d.key); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
