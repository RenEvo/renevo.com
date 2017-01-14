---
title: Facebook Site Thumbnails Tutorial
published: true
date: 2009-07-20 07:54:29 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2009/07/20/facebook-site-thumbnails-tutorial.aspx
file: facebook-site-thumbnails-tutorial.aspx
path: /blogs/developer/archive/2009/07/20/
author: tom anderson
words: 190
---
Here is a quick tutorial for all you [Facebook][1] users.

It annoyed me that when I added links from my website, the images that were given were "eh" at best, and usually from some random blog post. I had followed a tutorial that I found on using the image_src meta, but it just didn't work for me.

Here is my process

First thing I did was open up my site in my browser, I then "zoomed" out a bit to get more content in the view (I use a widescreen monitor).  I then took a decent centered screenshot (OneNote screen clipping is awesome), in Photoshop I tweaked it a bit to my liking so that it was nice and centered.  I then exported as a 200x175 png file and named it "renevo.thumb.png"

![][2]

Uploaded to my web site, then added this tiny bit of HTML to the site header.

> _<img src="/images/__renevo__.thumb.png" alt="" style="display:none;" />_

This doesn't effect my website's layout at all, and since it is set to display:none "most" browsers will not download the image.

And, for the result when sharing a link on [Facebook][1].

![image][3]

So… super easy way to add a 100% available website thumbnail. 

![][4]

[1]: http://www.facebook.com/Tom.D.Anderson
[2]: http://www.renevo.com/images/renevo.thumb.png
[3]: http://www.renevo.com/blogs/developer/image_4F55F0B8.png "image"
[4]: http://renevo.com/aggbug.aspx?PostID=2239

