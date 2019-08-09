---
title: Power Testing
published: true
date: 2008-10-13 17:29:51 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2008/10/13/power-testing.aspx
file: power-testing.aspx
path: /blogs/developer/archive/2008/10/13/
author: tom anderson
words: 689
---
What happens to your company when you lose power for say, 2 minutes?

Until today, that question was never really asked at our company, like many others, we assumed full faith in Edison that our power will be stable, well, at least that's the measures that we took to prepare.

This morning, around 9:00am, our power flickered off then back on, probably due to the massive winds gusting up to 65mph. The experience thus far has literally been a nightmare. First off, the "few" computers that have APC battery backups beeped like mad, that is what they are supposed to do, but when the power came back on, they kept beeping, the batteries, having finally been used, have been destroyed.  Most of these backups are 5-10 years old, and have been used once or twice at most. This was the start of what I like to call, productivity hell.

1. Phone Systems are down, not a big deal, but the computers that run them did not restart with the power coming back on (those old computers with non-state power buttons). 
2. Domains servers are down, one of them is not even on an APC, had one of those nagging "non-state" power buttons, you know the type, a home made PC that was used as a domain controller, and even though our company has upgraded almost every other PC in the building, these didn't get updated. 
3. DNS is gone, once again, one of those "computers that stay off" when they lose power. Turns out our DHCP server (generally goes with the domain controller) took a dump when the system was shutdown improperly. Did I mention our Domains are still running Server 2000? 
4. Internet down, ok, this is an expected one, but when you have a help desk that uses internet to access the customers for support (via remote connections), as well as perform remote upgrades, then you have an entire development department that relies on internet and intranet for source control and bug tracking, it hurts. Not to mention the development that depends on 3rd party web services, which cannot be accessed. Not to mention those Hosted Solutions, about fifteen customers without access to the software we host for them. 
5. Source Control Down, as mentioned above, this one while not being effected by the nasty non-state power button on the computer, was completely locked out because of some "still unknown" router/switch issue that is preventing access to the source control, even by IP address. 

So, here it is nearly 7.5 hours after the "tiny" power flux and development is at a halt as we have a policy against working offline, due to junior developers never connecting back up, then submitting fixed bugs, and "oops I did a full get and overwrote my changes", we have no internet to browse to at least do research while we are unable to program (ever hear that story about those things called books?  most of our developers refuse to read something that may cut them.) Our tech support and help desk departments are at a near halt, our phones are working, but the internet is up for a minute or two, then down for five to ten minutes at a time.

The moral of the story? I would like to see power tests about once a quarter, monthly would be nice, but the scheduling requirements for that have a pretty big effect. What happens during the power test?

* Power is cycled to the building for 20 seconds in the middle of the night during off-peak hours (three of four am) 
* On restoration of power, document every step to get the systems back and running to 100%. 
* After steps are documented meet with the nerds of the company to figure out how to reduce those steps and/or automate them. 
* Implement changes to reduce down time. 

With this type of preparing you are going to know what to expect, so when it happens during peek hours you have reduced and documented the full restoration of the system to 100%.

Oh, Solitaire is a great filler for time!

 

*Edit: Turns out this was an all day affair.





