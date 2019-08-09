---
title: Wiisis Updates & WiiMod
published: true
date: 2008-05-10 19:12:24 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/community_blogs/archive/2008/05/10/wiisis-updates-wiimod.aspx
file: wiisis-updates-wiimod.aspx
path: /blogs/community_blogs/archive/2008/05/10/
author: tom anderson
words: 505
---
We have a large update to deliver from WiiLabs today and it is chock full of many goodies for all you modders and players alike. We (finally!) have the release of WiiMod available for download. We also have an update to the [WRLib C API][1] and the [Wiisis][2] modification.

_**WiiMod Release**_   
It took a bit longer than we had expected, but available now for downloading is WiiMod. WiiMod is an asset for mod teams who wish to incorporate the Nintendo Wii Remote and Nunchuk controllers into their own custom modifications, but do not have the means or desire to deal with C . WiiMod is a pre-compiled custom GameDll and file package that adds support for the Wii Remote and Nunchuk controllers to your custom modification by means of exclusive Flowgraph Nodes and an expandable Lua scripting framework.

**Features of WiiMod include: **

* Over 50 custom Flowgraph Nodes that let you control up to 4 Wii Remotes and get button, motion and sensor input from them.
* Complete Lua framework that gives you equal power and accessibility of the WRLib C API.
* Customizable Xml file to control basic mod features, including intro videos and starting Singleplayer map.

_**WiiMod Demonstration Modification**_   
Also available is a small mini-mod that utilizes WiiMod and demonstrates its abilities. It comes packaged with a small level with several mini-games that use the Wii Remote. All of these mini-games were made using WiiMod’s custom Flowgraph nodes and Lua framework. One of the puzzles involves “constructing and using a light switch” using the Wii Remote’s motion controls in a style similar to Metroid Prime and Zelda puzzles.

_**Wiisis and WRLib C API Update**_   
One of the top priorities when constructing WiiMod was to correct any major problems that arose after Wiisis’ release. Wiisis is a modification released last month that brings the Wii Remote and Nunchuk controllers into Crysis. The WRLib C API is the source code package that adds support for the Wii Remote and Nunchuk controllers to the Crysis C Mod SDK.   
These two items have been updated to improve the support and overall functionality. The change log for Wiisis is as followed: 

* [Bug] Fixed bug that caused Nunchuk gestures to not be recognized very well
* [Bug] Fixed bug that sometimes prevented LED lights from refreshing on start
* [Bug] Fixed bug that caused the default Windows cursor to be used
* [Bug] Fixed bug that caused the Wii Remote Manager to think it was in the main menu while playing in Editor mode
* [Change] Reworked Wii Remote Warning video
* [Change] Reworked Nano Suit Menu controls to use the IR Sensor cursor instead of the Nunchuk analog stick for selection usage

_**Related Links:**_   
\- [WiiMod Download][3]   
\- [WiiMod Demo Download][4]   
\- [Wiisis Download][2]   
\- [WRLib C API Download][1]   
\- [RenEvo Website][5]   
\- [Wiisis Website][6]

  
We hope you enjoy these releases and we are looking forward to seeing any mods that add support for the Wii Remote controller.



[1]: http://crymod.com/filebase.php?fileid=1097&lim=10
[2]: http://crymod.com/filebase.php?fileid=285
[3]: http://crymod.com/filebase.php?fileid=1484
[4]: http://crymod.com/filebase.php?fileid=1485
[5]: http://www.renevo.com/
[6]: http://www.wiisis.com/


