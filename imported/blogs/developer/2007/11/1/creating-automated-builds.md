---
title: Creating Automated Builds
published: true
date: 2007-11-01 18:44:03 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2007/11/01/creating-automated-builds.aspx
file: creating-automated-builds.aspx
path: /blogs/developer/archive/2007/11/01/
author: tom anderson
words: 1216
---
One of the stresses of working in software development is getting the applications built on a regular basis.  Generally builds are done on a Sr. Developers workstation, then sent off to testing attached to the bug software, via email, or a link to a shared resource.

Getting automated builds up an running is a tricky thing, and with some complex projects requires a lot of scripting and knowledge of the internal software as well as some licensing issues with the software.

When requested to get "Automated Nightly Builds" working by the VP of Software Development in a recent meeting, I looked at our 10 software products and just shuddered. Below is a list of software that we use.

[Microsoft Visual Studio][1] 2003, 2005, and Visual Basic 6 for Code   
[Seapine Surround Source Control Management][2] (SCM)   
[Seapine Test Track Pro][3] for Bug Tracking   
[Install Shield 12][4] for some of our software installation media   
[WinRar][5] for patch packaging   
[MSBuild][6] from the .Net 2.0 Framework for OneClick installation generation and Satellite Resource File Compilations

As you can see, building integration between a lot of those products and kicking off event driven automated builds is going to be tricky.  Although I have recently downloaded and implemented [Cruise Control .Net][7] for remote builds of one of our products, putting together the rest are going to be tricky. 

Let me explain the game plan that I will be attempting, and I will blog my progress with the pitfalls of the system.

**Pulling from source control**

The hardest part of this is going to be getting the code out of source control, granted it is just a quick pull of the files into a directory, but managing which branch to work out of is the trickiest.  Surround SCM allows you to do multiple branches of source code and we have implemented an implementation that allows our Junior developers to work on sub-branches, while I, as the only Senior developer at the moment, work in the main branch of the code.  On a daily basis, or as needed, I review the code and promote it from the Junior level sub-branch, retrieve the code, then check for a stable build.  This is obviously a process that can not be automated.  So firstly, I am going to have to create two builds, one for the in-testing code, which would be from the sub-branch, and a second for release code, which is pulled from the main branch.

**Compiling the code**

With the third party tools that we are currently using, some unique licensing has to be in place in order to compile the applications.  Although using command line compilers and not running out of the IDE might be able to just imbed current licensing, I am pretty sure that the licensing scheme of some of our assemblies will prove to require additional licensing.  A good example of this is the activation based Red Gate tools that we are using.

**Building the applications**

We have three different scenarios with the products that we produce. The first being the easiest, is a completely self contained application, all of the source code is stand alone, and requires no other shared projects.  The second scenario are web sites, these are going to require that we only pull the releasable files, such as ascx, aspx, asax, and binaries for instance. And the third is the most difficult, this one is our framework based applications.  We have a shared framework based on the Composite UI Application Block from Microsoft so that we can accommodate modular based application design.  When we want to add a new product, we simply create a new set of modules that load into the application framework.  Then we can distribute them and only turn them on if the client has licensing for them.  This is also the way that we are trying to move all of our applications into, our "OneSource" for applications, which also happens to be the name of our framework.

**Creating the OneClick Installs**

With our newest framework applications, we are using the OneClick technology from Microsoft that is built into .Net 2.0.  This is an outstanding way to get updates out to our customers without calling or emailing them to update their software. The largest problem with OneClick installations though is that they are statically bound to an install location so if you configure the OneClick to install from <http://www.domain.com> then you have to have to deploy it FROM that location.  This extra step is going to require that we additionally setup multiple staging environments for the OneClick installations.  The first staging area will be our internal developer only install.  This is used primarily for internal review of the latest software by the VP, Analysts, as well as Technical writers.  Secondly, we will want to create a Beta, or Pre-Release tab, this will need to be publicly accessible and could potentially be given to customers, the CEO, Project Managers, and the Help Desk for training on new features, as well as approval and additional QA.  The final staging that will need to be built is the Release staging.  This will basically be the released version of the software for everyone to install and will be made available via a public domain.

**Building the Install Media**

This one is by far one of the trickiest ones.  Due to the fact that files could be added or removed from the installations, knowing what to add, as well as adding it to the installer in the correct place before compiling, is going to prove to be extremely tricky.  To be honest, this is the area that I am least knowledgeable about how to do, and will be doing a lot of research on it when I get there.

**Building the Patch Media**

This one is pretty easy, basically we take the build directories for the software, add it to an SFX via a huge command line to WinRar.exe and it creates a nice little patch for us.  I could probably take it one step further generating a CRC32 compare list and only adding changed files, but this is a nice way to "Magic Zip" update previous installations that might not be completely up to date.

**Notifications**

Of coarse this entire process needs to be logged and sent to the right people if something goes wrong. Luckily this is where Cruise Control .Net is going to come into play, as it will log the entire process (by redirecting the standard output when it calls all of the executables) and save them to XML in the working directories of the builds. Myself and the VP also run the Cruise Control client application in the Tray that displays the status of each build, last build time, and any additional messages.  When a build is finished, we get a little popup in the bottom right hand of the screen letting us know if it was successful or not.  Obviously this is ok for working hour builds, but I would prefer to have more information.  So I will be setting up SMS alerting and Email logs to the end of the build.

 

Basically that's it... not much work, if you like this sort of thing, but unfortunately I still have to maintain the constant stream of new features, updates, and bug fixes to the current software.

I will keep you up to date.

![][8]

[1]: http://msdn2.microsoft.com/en-us/vstudio/default.aspx
[2]: http://www.seapine.com/surroundscm.html
[3]: http://www.seapine.com/ttpro.html
[4]: http://www.macrovision.com/products/installation/installshield.htm
[5]: http://www.win-rar.com/
[6]: http://msdn2.microsoft.com/en-us/library/0k6kkbsd.aspx
[7]: http://confluence.public.thoughtworks.org/display/CCNET/Welcome to CruiseControl.NET
[8]: http://renevo.com/aggbug.aspx?PostID=1533

