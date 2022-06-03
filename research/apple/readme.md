# Apple TV

- https://github.com/TDenisM/APPLE-TV-4K-Downloader/blob/main/appletv.py
- https://github.com/ytdl-org/youtube-dl/issues/30808

Using this video:

https://tv.apple.com/us/episode/biscuits/umc.cmc.45cu44369hb2qfuwr3fihnr8e

HLS:

~~~
https://play.itunes.apple.com/WebObjects/MZPlay.woa/hls/subscription/stream/playlist.m3u8?
cc=US&g=230&cdn=vod-ap2-aoc.tv.apple.com&a=1484589502&p=461374806&st=1821682575&a=1625486472&p=461370051&st=1821645191&a=1622268591&p=461372307&st=1821659224&
a=1613450761&p=461480166&st=1822491467&a=1522961240&p=377679659&st=1490983814&a=1524197777&p=368330428&st=1449517784&a=1524197722&p=368330432&st=1449518254&
a=1524198082&p=368330370&st=1449517587&a=1525078430&p=368329706&st=1449509871&a=1524197604&p=368330236&st=1449518699&a=1524197554&p=368330322&st=1449518442&
a=1524197773&p=368330253&st=1449517917&a=1539152595&p=368283705&st=1449199894
~~~

Next we need the Widevine [1] PSSH from the HLS file:

~~~xml
#EXT-X-KEY:URI="data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgu8rYzAgICAgICBI88aJmwY=",
KEYFORMAT="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed",KEYFORMATVERSIONS="1",METHOD=SAMPLE-AES
~~~

1. <https://dashif.org/identifiers/content_protection>

License request:

~~~
POST https://play.itunes.apple.com/WebObjects/MZPlay.woa/web/video/subscription/license HTTP/2.0
content-type: application/json
x-apple-music-user-token: Ao8j9GzJD8Ga+/wgUa+Sh3tp858DCTyWIOG27Uk3AmTI4ys6e+sb3ENUqCEOPnxfDg/QUU1NO0kE/YwULulZqGfKyzaml8X76eCSJY2Lxo3YuNQ3qk2oF9Rz2dehfUY/sr8pf/YkIrL+YQ/aGP/QWim5aCL/avQuy5r3dmSuu2mxvySyULHoDmjBGqMMVzgiyOpVgnWMtIxduDxlciu7c/8EdxtbCdaHeoCKqbrRGetu7zPhPA==
authorization: Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IldlYlBsYXlLaWQifQ.eyJpc3MiOiJBTVBXZWJQbGF5IiwiaWF0IjoxNjUzMzI4NTUwLCJleHAiOjE2Njg4ODA1NTAsInJvb3RfaHR0cHNfb3JpZ2luIjpbImFwcGxlLmNvbSJdfQ.g0i-Loh2MNf0KQGMNjupeXFdvafnPyiAmhrUUxFGHGXspqxgmT7vH3LqFxLJ21MXtbeUo_HTqV3sXXaj2mxx_g

{
   "challenge": "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",
   "extra-server-parameters": {
      "adamId": "1522961240",
      "svcId": "tvs.vds.4105"
   },
   "key-system": "com.widevine.alpha",
   "uri": "data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgu8rYzAgICAgICBI88aJmwY="
}
~~~

Authorization comes from response HTML:

https://tv.apple.com