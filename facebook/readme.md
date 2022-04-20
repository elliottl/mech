# Facebook

https://github.com/ytdl-org/youtube-dl/issues/30872

## Android

- https://android.stackexchange.com/questions/246303/facebook-whitehat-settings
- https://play.google.com/store/apps/details?id=com.facebook.katana

HTTP Toolkit:

Install Facebook app. Go to Menu, Settings & Privacy, Whitehat Settings. Click
"allow user installed certificates". FORCE STOP app. Start HTTP Toolkit. Start
app. Notice that requests are not being captured, and Certificate rejected
errors are occurring.

https://github.com/httptoolkit/frida-android-unpinning/issues/18

MITM Proxy:

Install user certificate. Turn off proxy. Install Facebook app. Go to Menu,
Settings & Privacy, Whitehat Settings. Click "allow user installed
certificates". FORCE STOP app. Turn on proxy. Start app. Notice that requests
are not being captured. Go back to Whitehat Settings, and click "Proxy for
Platform API requests". Enter 127.0.0.1:8080. FORCE STOP app. Start app. Notice
that requests are not being captured.

https://github.com/mitmproxy/mitmproxy/discussions/5271

## Web

Using this video:

https://www.facebook.com/FromTheBasementPage/videos/309868367063220

some projects just parse the HTML:

~~~
GET /FromTheBasementPage/videos/309868367063220 HTTP/1.1
Host: www.facebook.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3766.0 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close
content-length: 0
~~~

others use DASH:

~~~
GET /video/playback/dash_mpd_debug.mpd?v=309868367063220&dummy=.mpd HTTP/1.1
Host: www.facebook.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.41 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Accept-Encoding: gzip, deflate
Connection: close
content-length: 0
~~~

Is the quality the same? Yes. Next we can try monitoring this page:

https://www.facebook.com/video.php?v=309868367063220

Same result. Next we can try monitoring this page:

<https://www.facebook.com/video/embed?video_id=309868367063220>

That only returns HTML, no JSON. Next we can try monitoring this page:

https://www.facebook.com/watch

4561733853932056:

~~~
POST https://www.facebook.com/api/graphql/ HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
content-type: application/x-www-form-urlencoded
x-fb-friendly-name: CometVideoHomeLOEVideoPermalinkAuxiliaryRootQuery
x-fb-lsd: AVqz-uQgZmM
content-length: 1066
origin: https://www.facebook.com
dnt: 1
referer: https://www.facebook.com/watch
cookie: datr=TELWYW7HP4RtCGikVySUp6x-
cookie: fr=0W6VmibpnyLPd9tKK.AWVlKEWfTBue7bA9KbNc1fyBDFc.BiWv6G.wl.AAA.0.0.BiWwxw.AWW1MTA4ySM
cookie: sb=y2kMYljEFad1tU7CPu3dtnFa
cookie: dpr=1.25
cookie: wd=1536x699
cookie: locale=en_US
cookie: m_pixel_ratio=1
cookie: usida=eyJ2ZXIiOjEsImlkIjoiQXJhZzJxcm91YXlmMSIsInRpbWUiOjE2NTAxMzM1NDZ9
te: trailers

av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cw5MKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q0GpovU19pobodEGdw46wbS16Awzw&__csr=gRsy6SzYZHihRArhnqGjAKaGFbXBUjBGqWx2djjyk268VFaV9oix28ByA59VU8bDyUC2emuuElBhoC2GbyEdEKqdx2F8W7USUeo8HwGw0VLxq05-U1dE6500_Ow0aDi8Gq9nO1u3-04oo0qjw8uu0CV80Ju050opw3QU0apU0dYE0DKlo0LG0pF0d-7UV0BDw2GIYU0ibw&__req=9&__hs=19098.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005366256&__s=jt7voe%3Anshyhl%3Au8k4e5&__hsi=7087289471768853946-0&__comet_req=1&lsd=AVqz-uQgZmM&jazoest=21008&__spin_r=1005366256&__spin_b=trunk&__spin_t=1650138169&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometVideoHomeLOEVideoPermalinkAuxiliaryRootQuery&variables=%7B%22SEOInfoTriggerData%22%3A%7B%22video_id%22%3A%22893177214675696%22%7D%2C%22relatedPagesTriggerData%22%3A%7B%22video_id%22%3A%22893177214675696%22%7D%2C%22scale%22%3A1.5%2C%22triggerData%22%3A%7B%22video_id%22%3A%22893177214675696%22%7D%7D&server_timestamps=true&doc_id=4561733853932056

{"data":{"video_home_www_related_videos_section":{"section_components":{"edges":[
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
~~~

4987920667970827:

~~~
POST https://www.facebook.com/api/graphql/ HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
content-type: application/x-www-form-urlencoded
x-fb-friendly-name: CometVideoHomeNewPermalinkHeroUnitQuery
x-fb-lsd: AVrSNpSxeEQ
content-length: 1647
origin: https://www.facebook.com
dnt: 1
referer: https://www.facebook.com/watch
cookie: dpr=1.25
cookie: datr=MhtgYqqBJE_lz5CiRKb-ZJAA
te: trailers

av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cw5MKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q0GpovU19pobodEGdw46wbS16Awzw&__csr=gBkAAy5FPGyleK8hripaLlrFDAG9Ahp9uijyZ2axqEDAyoC5p9KibDGdCGqi2i5EkyV8V5zSeAyEgx_V8gKbwHwPWHxW4osCDAzQ9U6-03sW023ow0fFE02A1dx60oG888yU0iSw1E-1-g7O780v2K0dYw0tZo0xy4ofk1BhoGE0GUhe04go&__req=a&__hs=19102.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1&__ccg=EXCELLENT&__rev=1005380421&__s=o8bjfq%3A375vwd%3A9kay20&__hsi=7088695716733192194-0&__comet_req=1&lsd=AVrSNpSxeEQ&jazoest=2992&__spin_r=1005380421&__spin_b=trunk&__spin_t=1650465586&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometVideoHomeNewPermalinkHeroUnitQuery&variables=%7B%22displayCommentsContextEnableComment%22%3Anull%2C%22displayCommentsContextIsAdPreview%22%3Anull%2C%22displayCommentsContextIsAggregatedShare%22%3Anull%2C%22displayCommentsContextIsStorySet%22%3Anull%2C%22displayCommentsFeedbackContext%22%3Anull%2C%22focusCommentID%22%3Anull%2C%22privacySelectorRenderLocation%22%3A%22COMET_STREAM%22%2C%22renderLocation%22%3A%22video_home%22%2C%22scale%22%3A1%2C%22useDefaultActor%22%3Afalse%2C%22videoID%22%3A%222031900003587453%22%2C%22UFI2CommentsProvider_commentsKey%22%3A%22CometVideoHomeNewPermalinkHeroUnitQuery%22%2C%22caller%22%3A%22TAHOE%22%2C%22channelEntryPoint%22%3A%22TAHOE%22%2C%22channelID%22%3A%22%22%2C%22feedbackSource%22%3A41%2C%22feedLocation%22%3A%22TAHOE%22%2C%22isLoggedOut%22%3Atrue%2C%22skipFetchingChaining%22%3Atrue%2C%22streamChainingSection%22%3Afalse%2C%22videoChainingContext%22%3Anull%7D&server_timestamps=true&doc_id=4987920667970827

{"data":{"video":{"story":{"attachments":[{"media":{"__typename":"Video","preferr
{"label":"CometVideoHomeHeroUnitPlayerSurface_video$defer$CometTahoeUpNextOverlay
{"label":"CometVideoHomeHeroUnitPlayerSurface_video$defer$VideoPlayerWithVideoCar
{"label":"CometVideoHomeHeroUnit_story$defer$CometVideoHomeHeroUnitSidePane_story
{"label":"CometTahoeVideoContextSectionBody_video$defer$CometTahoeSidePaneAttachm
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"CometVideoHomeHeroUnit_story$defer$CometVideoHomeHeroUnitLeftBottomSect
{"label":"CometVideoHomeHeroUnitSidePane_story$defer$CometTahoeUFIChainingSection
~~~

5214388728679644:

~~~
POST https://www.facebook.com/api/graphql/ HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
content-type: application/x-www-form-urlencoded
x-fb-friendly-name: CometVideoHomeLOEHomeRootQuery
x-fb-lsd: AVqz-uQg_ak
content-length: 935
origin: https://www.facebook.com
dnt: 1
referer: https://www.facebook.com/ESPN/videos/watch-brett-phillips-hits-home-run-with-his-biggest-fan-in-attendance/501982491404746/
cookie: datr=TELWYW7HP4RtCGikVySUp6x-
cookie: fr=0W6VmibpnyLPd9tKK.AWVlKEWfTBue7bA9KbNc1fyBDFc.BiWv6G.wl.AAA.0.0.BiWwxw.AWW1MTA4ySM
cookie: sb=y2kMYljEFad1tU7CPu3dtnFa
cookie: dpr=1.25
cookie: wd=1186x615
cookie: locale=en_US
cookie: m_pixel_ratio=1
cookie: usida=eyJ2ZXIiOjEsImlkIjoiQXJhZzJxcm91YXlmMSIsInRpbWUiOjE2NTAxMzM1NDZ9
te: trailers

av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cwfG12wOKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q3q1OBx_y8179obodEGdwda3e0Lo4qifxe&__csr=gZ8wBf4Tl99eBSQkgnSiAaGFf-nUyrVoCWwzjyo8KrV9aV4KdAx28x15UOuu4-VHDgkG22muuEkhkm9wGyUGawHyVESeWAzElyuVXwVx68DLwEw1wG0QmEhwrU1uqw1ci0Oomw1vK0jq1xg0L6ew0PNw5cw089S0H83uw1LS8Gq9nO2V9EfU0h3xS0aIg0-mU4O3qu0CV8iwn-0qS0J81oo0M-2u1DCxa2-0bUwaG0nO8w15W04wE0mEw0x9w2uVlw4Nwem682Tx62S0mN0d-7UV0BDw2GIYU1x8W0sO9w4ww&__req=m&__hs=19098.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005366273&__s=hpahh1%3A9t1dg4%3Ahfftu1&__hsi=7087307216950888986-0&__comet_req=1&lsd=AVqz-uQg_ak&jazoest=21031&__spin_r=1005366273&__spin_b=trunk&__spin_t=1650142301&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometVideoHomeLOEHomeRootQuery&variables=%7B%22scale%22%3A1.5%7D&server_timestamps=true&doc_id=5214388728679644

{"data":{"video_home_www_logged_out_home":{"video_home_sections":{"edges":[{"node
{"label":"CometVideoHomeLOEHomeSectionsList_query$stream$CometVideoHomeLOEHomeSec
{"label":"CometVideoHomeLOEHomeSectionsList_query$defer$CometVideoHomeLOEHomeSect
~~~

5242228839129775:

~~~
POST https://www.facebook.com/api/graphql/ HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
content-type: application/x-www-form-urlencoded
x-fb-friendly-name: CometTahoeUpNextEndCardQuery
x-fb-lsd: AVrSNpSxeEQ
content-length: 970
origin: https://www.facebook.com
dnt: 1
referer: https://www.facebook.com/CookistWow/videos/egg-white-breakfast-cups-ready-in-only-15-minutes/2031900003587453/
cookie: dpr=1.25
cookie: datr=MhtgYqqBJE_lz5CiRKb-ZJAA
te: trailers

av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cwfG12wOKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q3q1OBx_y8179obodEGdwda3e0Lo4qifxe&__csr=gR9969PT7Hy4mCyfjRnWpV5yrpp9uiLV_gzgmGm-i9yolACV8KuETAGqi2h6zEkyVaDihoYwSiax28xrV4fKbxSdz89KGXxGicxmqp6AzQ9U4ycxC04z80TTnwmWw1l2cw0wS80cwzU0cto1B807_S19Dw1NMSbGE1yEwwybw1bq360oB01dai1bwJg7O6bwhU25w5bw27o0NW3C7Uh-0cowoU0Fu8w3ME0a580bSU31wdW6EiyU1TEhwZg6l5yGw2Hx4U1x-0aTw&__req=e&__hs=19102.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1&__ccg=EXCELLENT&__rev=1005380421&__s=o8bjfq%3A375vwd%3A9kay20&__hsi=7088695716733192194-0&__comet_req=1&lsd=AVrSNpSxeEQ&jazoest=2992&__spin_r=1005380421&__spin_b=trunk&__spin_t=1650465586&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometTahoeUpNextEndCardQuery&variables=%7B%22upNextVideoID%22%3A%22%22%2C%22scale%22%3A1%2C%22currentID%22%3A%222031900003587453%22%7D&server_timestamps=true&doc_id=5242228839129775

{"data":{"upNextVideoData":null,"currentVideo":{"__typename":"Video","video_chann
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
~~~

5258095424210922:

~~~
POST https://www.facebook.com/api/graphql/ HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
content-type: application/x-www-form-urlencoded
x-fb-friendly-name: CometVideoHomeNewPermalinkHeroUnitQuery
x-fb-lsd: AVqz-uQgZmM
content-length: 1673
origin: https://www.facebook.com
dnt: 1
referer: https://www.facebook.com/watch
cookie: datr=TELWYW7HP4RtCGikVySUp6x-
cookie: fr=0W6VmibpnyLPd9tKK.AWVlKEWfTBue7bA9KbNc1fyBDFc.BiWv6G.wl.AAA.0.0.BiWwxw.AWW1MTA4ySM
cookie: sb=y2kMYljEFad1tU7CPu3dtnFa
cookie: dpr=1.25
cookie: wd=1536x699
cookie: locale=en_US
cookie: m_pixel_ratio=1
cookie: usida=eyJ2ZXIiOjEsImlkIjoiQXJhZzJxcm91YXlmMSIsInRpbWUiOjE2NTAxMzM1NDZ9
te: trailers

av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cw5MKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q0GpovU19pobodEGdw46wbS16Awzw&__csr=gRsy6SzYZHihRArhnqGjAKaGFbXBUjBGqWx2djjyk268VFaV9oix28ByA59VU8bDyUC2emuuElBhoC2GbyEdEKqdx2F8W7USUeo8HwGw0VLxq05-U1dE6500_Ow0aDi8Gq9nO1u3-04oo0qjw8uu0CV80Ju050opw3QU0apU0dYE0DKlo0LG0pF0d-7UV0BDw2GIYU0ibw&__req=a&__hs=19098.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005366256&__s=jt7voe%3Anshyhl%3Au8k4e5&__hsi=7087289471768853946-0&__comet_req=1&lsd=AVqz-uQgZmM&jazoest=21008&__spin_r=1005366256&__spin_b=trunk&__spin_t=1650138169&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometVideoHomeNewPermalinkHeroUnitQuery&variables=%7B%22displayCommentsContextEnableComment%22%3Anull%2C%22displayCommentsContextIsAdPreview%22%3Anull%2C%22displayCommentsContextIsAggregatedShare%22%3Anull%2C%22displayCommentsContextIsStorySet%22%3Anull%2C%22displayCommentsFeedbackContext%22%3Anull%2C%22focusCommentID%22%3Anull%2C%22privacySelectorRenderLocation%22%3A%22COMET_STREAM%22%2C%22renderLocation%22%3A%22video_home%22%2C%22scale%22%3A1.5%2C%22useDefaultActor%22%3Afalse%2C%22videoID%22%3A%22893177214675696%22%2C%22UFI2CommentsProvider_commentsKey%22%3A%22CometVideoHomeNewPermalinkHeroUnitQuery%22%2C%22caller%22%3A%22TAHOE%22%2C%22channelEntryPoint%22%3A%22TAHOE%22%2C%22channelID%22%3A%22%22%2C%22feedbackSource%22%3A41%2C%22feedLocation%22%3A%22TAHOE%22%2C%22isLoggedOut%22%3Atrue%2C%22skipFetchingChaining%22%3Atrue%2C%22streamChainingSection%22%3Afalse%2C%22videoChainingContext%22%3Anull%7D&server_timestamps=true&doc_id=5258095424210922

{"data":{"video":{"story":{"attachments":[{"media":{"__typename":"Video","preferr
{"label":"CometVideoHomeHeroUnitPlayerSurface_video$defer$CometTahoeUpNextOverlay
{"label":"CometVideoHomeHeroUnitPlayerSurface_video$defer$VideoPlayerWithVideoCar
{"label":"CometVideoHomeHeroUnit_story$defer$CometVideoHomeHeroUnitSidePane_story
{"label":"CometTahoeVideoContextSectionBody_video$defer$CometTahoeSidePaneAttachm
{"label":"CometVideoHomeHeroUnit_story$defer$CometVideoHomeHeroUnitLeftBottomSect
{"label":"VideoPlayerRelay_video$defer$InstreamVideoAdBreaksPlayer_video","path":
{"label":"CometVideoHomeHeroUnitSidePane_story$defer$CometTahoeUFIChainingSection
~~~

I found this:

<https://facebook.com/video/video_data?video_id=309868367063220>

but it only contains the media, not metadata, so its only marginally useful.
Same for this:

<https://www.facebook.com/video/video_data_async/?video_id=309868367063220&__a=1>

and this:

<https://www.facebook.com/pages/profile/cover_video_data/?video_id=309868367063220&__a=1>