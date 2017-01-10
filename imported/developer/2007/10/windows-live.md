---
title: Windows "Live"
published: true
date: 2007-10-26 13:31:00 +0000 UTC
tags: imported 
---
Originally some time last year when Windows Live started getting all of this attention, I was drawn in being a Microsoft Junkie and all.  Windows Live Messenger, OneCare, Live Writer, Hotmail Live, I was googie over them all.

Then, came the dreaded "You can not install this program on that operating system".

Lets back up a bit... As soon as Vista went RTM, I installed it, see above about how much of a Microsoft Junkie I am. Having a Dual Dual Core Opterons, installing Vista 64 bit seemed the obvious choice, so that I could take advantage of the extra processing capabilities of the 64 bit and cpus.  First message I get when installing all of my old software "Windows OneCare cannot run on this operating system.".  Are you kidding me?  A Microsoft Product not working on a brand new Operating System, specifically a **brand new** anti-virus system built **for** windows? So, I resorted to installing Avast, which seemed at the time to be the only 64 bit capable Vista anti-virus.

Now, lets move forward to today... Once again, I got an itch to try out the Windows Live Writer, I must admit that I had been a bit lazy at posting to this blog, so I thought I would get back on track.  I head over to live spaces to pick up the newest Beta for Windows Live Writer to get the same message I got with OneCare. "You cannot install this program on this operating system".

Ok, lets look at this logically, and with a bit of thought.

Windows Live Writer is a .Net 2.0 application, most likely written in C#. I am a .Net 2.0 Developer that works in Vista Ultimate 64. Why the hell can't this "special" .Net 2.0 program run in Vista 64?  The .Net framework manages the mode (be-it 32 or 64 bit) when the program is JIT (Just In Time) Compiled from the IL (Intermediate Language).  So what the hell?  Even then, if they noticed some problems with the JIT, why couldn't they have just specified x86 mode, instead of "AnyCpu" configuration on compile?  I am currently developing an x86 compiled application right now that I have 0 problems running, debugging, and building on x64.

I am getting a bit "iffy" about this whole Microsoft's direction with "Live", whats next, the newest patch to Windows Live Messenger won't run on x64?  Now that would be a killer, considering I use Live Messenger for communicating with co-workers on and off site.  When will Microsoft understand that x64 is a mainstream operating system, and they need to get thier apps up to thier own "Microsoft Certified Application" specifications

_Application must run in both 32 and 64 bit modes_

Anyway, Enough about live, lets see how my registration for "Live Custom Domains" goes, still have no idea what the hell it is, but it has an SDK, and I had a free domain name to play with.

**Links**  
[http://www.live.com][1]  
[http://dev.live.com][2]  
<http://dev.live.com/customdomains/>  
<http://dev.live.com/writer/>  

![][3]

[1]: http://www.live.com/
[2]: http://dev.live.com/
[3]: http://renevo.com/aggbug.aspx?PostID=1514

