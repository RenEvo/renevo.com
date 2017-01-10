---
title: The importance of Virtual PC
published: true
date: 2007-11-01 19:01:42 +0000 UTC
tags: imported 
---
Up until about 2 months ago I had not realized what I was missing out on.

[Microsoft Virtual PC][1] is not only extremely useful, it is Free. Granted you have to pay for the operating systems you install, that is still one hell of a deal when you compare the cost of a test machine, the power to run it, as well as the KVM or storage locations for it.

So, what is Virtual PC?  Virtual PC is a computer emulator that you can run on your desktop computers.  It loads up like a computer, it uses virtual hardware, and you can install Operating Systems on it.

Why is this so important you might ask?  Well lets think of this phrase "It works on my machine...".  This has been a software mantra for quite a few years, and admitidly I used to over use it myself.

During the software development process you don't find a lot of developers who will take the time to test it on multiple operating systems, clean machines, or other configurations that are common for your target audience.  Instead, you fix the bug, run it from the IDE, if it doesn't crash again from the Replication steps (if you had any), then you sign it off as good and either deliver a patch to the customer, or send it to QA. Let us hope you are doing the latter.

Now, lets investigate this situation.  You do the following on your Windows Vista box, you send it to your QA guy, who also happens to be using Windows Vista, he signs it off as fixed, and it is deployed to the customer. Except the customer is using Windows XP, and that bug still exists in Windows XP because for some reason the Users folders are in totally different locations.  You never tested it in that environment, so you have no idea what is going to happen, yet again the customer lets you know that it is broken, and it comes back to you.

How do you solve this customer loop back?  Very easily actually, you simply use Virtual PC to test on a few different environments before you send it to QA or the customer.

Below is my list of Virtual PC's that I have setup, and use daily.

* Blank install of Windows XP SP2 - Always up to date.
* Windows XP SP2 with the Point Of Sale we integrate into.
* Windows XP SP2 with Office 2003 - Always up to date.
* Blank install of Windows Server 2003 R2 w/IIS & SQL Server 2000
* Blank install of Windows Server 2003 R2 w/IIS & SQL Server 2005
* Blank install of Windows Server 2003 w/IIS (non-R2)
* Blank install of Windows Vista Business - (for 32bit testing, as I run Vista 64 bit)

As you can see from the list above, it targets my entire client base.  We restrict our software to only be installed on XP SP2, Server 2003, Server 2003 R2, and Vista. At any time I can load up one of these PC's and do a quick test of the issue that is being reports on the operating system configuration that the customer is complaining about, and find the issue much faster, and make sure it doesn't happen on ANY of the configurations again.

I suggest taking the time, as well as the Hard Drive space (each OS is about 5 gigs) and get this setup, once you have your targets, do a monthly burn to DVD to keep a clean backup with the latest updates so you can always regress back to a non-damaged environment.

_Note: Game development, or graphic heavy development will most likely require real computers for the graphic processing speeds, these are for more business development needs._

Good luck!

![][2]

[1]: http://www.microsoft.com/windows/products/winfamily/virtualpc/default.mspx
[2]: http://renevo.com/aggbug.aspx?PostID=1534

