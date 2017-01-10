---
title: Dante broke the egg
published: true
date: 2006-12-14 04:46:00 +0000 UTC
tags: imported 
---
Far Cry's crygame.dll - the comstock load for all the game logic in Far Cry - has come into focus for the C&C Far Cry project.   FINALLY.  The tears you see: they are real.  It's taken me awhile to get my composure straightened out.  I'm sooo happy.  *gasps for air*

Give me a minute... talks amongst yourselves. 

Ok, back to reality.  Pardon the stupidity of that opening.  Being a software engineer (just half a year shy of receiving my bachelors of science) I am most comfortable in the C programming language.  I can solve any problem, write any rule, and do it very lean and fast.  What that means for you is a better game in a shorter amount of time.  Up until now, all of our code has been situated on the Lua (an external scripting language comparable to mIRC scripting only a lot more open ended) for several reasons, namely to speed up development and to have the code out in the open for others to build off of and modify.  For us, having a strong mod community is important, and Lua lets us build a strong game that at the same time is very accessible for you to build a strong mod. 

Now just because we're migrating some of our code into C does not mean we are compromising our principles.  It means just the opposite.  Our intentions are to hide only the nitty-gritty details that are heavy in math and computational paradigms.  We can speed up these items, making the game faster and overall better for you.  With access to the net code, rendering methods, and script objects, we can improve the performance of the game while also putting newer items into Lua that you - the modder - can use.  Items which would otherwise be totally inaccessible for the average mod team.

For example: Far Cry has no actual tanks anywhere within the shipped version of the game.  They are not part of the game. They are in C&C Far Cry.  Because of this, turrets on vehicles update their orientations unrealisticly when compared to how a tank works.  Where ever you look, that is where your turret points.  By gaining access to the DLL, we will be able to implant a capped movement speed on turrets, forcing them to move towards where you are looking at at a defined speed, making it look and feel more realistic.  We - and by that includes you - would not have been able to do this from the available resources in the Lua environment for this game.  What we will do besides add this ability is also extract the speed setting for a turret on a vehicle to the Lua script for that vehicle, enabling you to define this max speed from inside the Lua script.  Now something that you couldn't do can now be done very, very easily.

It's 3 in the morning.  Cut me some slack here. 

Over the holiday period, we will be porting a lot of foundation code into the DLL and perform some clean up in our Lua stock code.  Some core features - like the one mentioned above - are also planned.  C&C FarCry is going to prove to be a testing ground for some of the concepts we have in mind for C&C: The Dead 6.  There are going to be some pretty exciting things going on in the next few months, so stick around.  

![][1]

[1]: http://renevo.com/aggbug.aspx?PostID=413

