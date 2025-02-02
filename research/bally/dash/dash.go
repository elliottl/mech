package dash

import (
   "fmt"
   "strconv"
   "strings"
)

func (p Presentation) Representation() Representations {
   var reps []Representation
   for i, per := range p.Period {
      if strings.HasPrefix(per.ID, "src") {
         for j, ada := range per.AdaptationSet {
            for _, rep := range ada.Representation {
               rep.Adaptation = &p.Period[i].AdaptationSet[j]
               if rep.Codecs == "" {
                  rep.Codecs = ada.Codecs
               }
               if rep.ContentProtection == nil {
                  rep.ContentProtection = ada.ContentProtection
               }
               if rep.MimeType == "" {
                  rep.MimeType = ada.MimeType
               }
               if rep.SegmentTemplate == nil {
                  rep.SegmentTemplate = ada.SegmentTemplate
               }
               rep.BaseURL = per.BaseURL
               reps = append(reps, rep)
            }
         }
      }
   }
   return reps
}

var verbose bool

func (r Representation) Media(position *int) []string {
   var start int
   var refs []string
   for i, seg := range r.SegmentTemplate.SegmentTimeline.S {
      if i == 0 {
         start = seg.T
      }
      for seg.R >= 0 {
         ref := r.replace_ID(r.SegmentTemplate.Media)
         ref = strings.Replace(ref, "$Time$", strconv.Itoa(start), 1)
         seg.R--
         var status string
         if start > *position {
            refs = append(refs, ref)
            *position = start
            status = "new"
         } else {
            status = "old"
         }
         if verbose {
            fmt.Println(status, start)
         }
         start += seg.D
      }
   }
   return refs
}

type Representation struct {
   Adaptation *Adaptation
   Bandwidth int64 `xml:"bandwidth,attr"`
   Codecs string `xml:"codecs,attr"`
   ContentProtection *ContentProtection
   Height int64 `xml:"height,attr"`
   ID string `xml:"id,attr"`
   MimeType string `xml:"mimeType,attr"`
   SegmentTemplate *SegmentTemplate
   Width int64 `xml:"width,attr"`
   BaseURL string
}

type Presentation struct {
   Period []struct {
      AdaptationSet []Adaptation
      BaseURL string
      ID string `xml:"id,attr"`
   }
}

func (r Representation) String() string {
   var b []byte
   b = append(b, "ID:"...)
   b = append(b, r.ID...)
   if r.Width + r.Bandwidth >= 1 {
      b = append(b, "\n  "...)
   }
   if r.Width >= 1 {
      b = append(b, "Width:"...)
      b = strconv.AppendInt(b, r.Width, 10)
      b = append(b, " Height:"...)
      b = strconv.AppendInt(b, r.Height, 10)
   }
   if r.Bandwidth >= 1 {
      if r.Width >= 1 {
         b = append(b, ' ')
      }
      b = append(b, "Bandwidth:"...)
      b = strconv.AppendInt(b, r.Bandwidth, 10)
   }
   b = append(b, "\n  MimeType:"...)
   b = append(b, r.MimeType...)
   if r.Codecs != "" {
      b = append(b, " Codecs:"...)
      b = append(b, r.Codecs...)
   }
   if r.Adaptation.Lang != "" {
      b = append(b, " Lang:"...)
      b = append(b, r.Adaptation.Lang...)
   }
   if r.Adaptation.Role != nil {
      b = append(b, " Role:"...)
      b = append(b, r.Adaptation.Role.Value...)
   }
   return string(b)
}

func (r Representation) Initialization() string {
   return r.replace_ID(r.SegmentTemplate.Initialization)
}

type Segment struct {
   D int `xml:"d,attr"` // duration
   R int `xml:"r,attr"` // repeat
   T int `xml:"t,attr"` // time
}

type SegmentTemplate struct {
   Initialization string `xml:"initialization,attr"`
   Media string `xml:"media,attr"`
   SegmentTimeline struct {
      S []Segment
   }
   PresentationTimeOffset *int `xml:"presentationTimeOffset,attr"`
}

func (r Representations) Filter(f func(Representation) bool) Representations {
   var carry []Representation
   for _, item := range r {
      if f(item) {
         carry = append(carry, item)
      }
   }
   return carry
}

func (r Representations) Video() Representations {
   return r.Filter(func(a Representation) bool {
      return a.MimeType == "video/mp4"
   })
}

func (r Representations) Audio() Representations {
   return r.Filter(func(a Representation) bool {
      return a.MimeType == "audio/mp4"
   })
}

func (r Representations) Index(f func(a, b Representation) bool) int {
   carry := -1
   for i, item := range r {
      if carry == -1 || f(r[carry], item) {
         carry = i
      }
   }
   return carry
}

func (r Representations) Bandwidth(v int64) int {
   distance := func(a Representation) int64 {
      if a.Bandwidth > v {
         return a.Bandwidth - v
      }
      return v - a.Bandwidth
   }
   return r.Index(func(carry, item Representation) bool {
      return distance(item) < distance(carry)
   })
}

func (s Segment) Time() string {
   return strconv.Itoa(s.T)
}

type Representations []Representation

type Adaptation struct {
   Codecs string `xml:"codecs,attr"`
   ContentProtection *ContentProtection
   Lang string `xml:"lang,attr"`
   MimeType string `xml:"mimeType,attr"`
   Role *struct {
      Value string `xml:"value,attr"`
   }
   SegmentTemplate *SegmentTemplate
   Representation []Representation
}

type ContentProtection struct {
   Default_KID string `xml:"default_KID,attr"`
}

func (r Representation) Ext() string {
   switch r.MimeType {
   case "video/mp4":
      return ".m4v"
   case "audio/mp4":
      return ".m4a"
   }
   return ""
}

func (r Representation) Role() string {
   if r.Adaptation.Role == nil {
      return ""
   }
   return r.Adaptation.Role.Value
}

func (r Representation) replace_ID(s string) string {
   return strings.Replace(s, "$RepresentationID$", r.ID, 1)
}
