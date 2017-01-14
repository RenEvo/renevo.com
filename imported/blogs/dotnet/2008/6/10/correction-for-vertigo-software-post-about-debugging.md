---
title: Correction for Vertigo Software post about debugging
published: true
date: 2008-06-10 18:01:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2008/06/10/correction-for-vertigo-software-post-about-debugging.aspx
file: correction-for-vertigo-software-post-about-debugging.aspx
path: /blogs/dotnet/archive/2008/06/10/
author: tom anderson
words: 312
---
Recently I read the article by [Chris Idzerda][1] over at [Vertigo software][2] that described how to [Detect Visual Studio Debugging][3].  While the method he used was a bit of a "hacky workaround", below is a more sound way to detect it using built in .Net libraries, rather than a string lookup on the running executable.

VB.Net Version
    
    
       1:          If System.Diagnostics.Debugger.IsAttached Then
    
    
       2:              ' do code here for debugging with visual studio
    
    
       3:   
    
    
       4:          End If
    
    
     

C# Version
    
    
       1:              if (System.Diagnostics.Debugger.IsAttached)
    
    
       2:              {
    
    
       3:                  // do code here for debugging with visual studio
    
    
       4:              }

I hope this helps clear up any confusion.

![][4]

[1]: http://blogs.vertigo.com/personal/chris/Blog/Lists/Categories/Category.aspx?Name=Work
[2]: http://www.vertigo.com/
[3]: http://blogs.vertigo.com/personal/chris/Blog/Lists/Posts/Post.aspx?ID=22
[4]: http://renevo.com/aggbug.aspx?PostID=1937

