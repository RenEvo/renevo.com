---
title: "Not my code..."
published: true
date: 2007-10-30 13:33:40 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2007/10/30/not-my-code.aspx
file: not-my-code.aspx
path: /blogs/developer/archive/2007/10/30/
author: tom anderson
words: 389
---
By definition in agile software development, no one owns the code.

What does this mean exactly?  Simply stated, if you find a bug, you fix it, if you can refactor a process without breaking it, you refactor, if you find something that can use some updates, you update it.  There is no architecture meeting, no code change approval process, and definitely no "I found a bug in your code".

I was faced with a situation where a junior developer had sent me an email during regular business hours (mind you I sit about 10 ft away from him) letting me know that there was an error in the code.

![MenuId.Error][1] 

From the code above any junior level developer could easily notice that the @MenuId was used instead of the @ButtonPropertiesId.  Instead of just changing the string name, I received an email at 5pm stating that the code was broken, and this might be the spot.  Why not just fix it?

Let me restate a few principles from the [Agile Manifesto][2], which our team tries to live by.

_The most efficient and effective method of conveying information to and within a development team is face-to-face conversation._

This basically means talk to each-other, instead of sending an email to someone 10ft away from you, get up, stretch your legs, relax your eyeballs, and talk to the other developers.

_Working software is the primary measure of progress._

So if the code works, you are progressing, if it does not work, then the progress does not continue until it does work.

_Continuous attention to technical excellence and good design enhances agility. _

If you spend some time doing it right, and "just fixing it", the progress is going to increase, hence the previously stated principle.

In short, if you find something that is broken, you see how to fix it, you have the knowledge to fix it, why not just fix it?

Now I do realize however that there are environments where the developers and programmers are confined to tiny spaces of code and can not dream of touching other portions of the codebase, but when you are encouraged to, and the team you are working with pushes the fact that the agile methodologies work, then double click on that variable and paste it above, then check it in and call Tom out on it, cause you just fixed a Senior developers mistake as a Junior.

![][3]

[1]: http://www.renevo.com/blogs/developer/WindowsLiveWriter/Notmycode_6A3E/MenuId.Error_thumb.jpg
[2]: http://www.agilemanifesto.org
[3]: http://renevo.com/aggbug.aspx?PostID=1529

