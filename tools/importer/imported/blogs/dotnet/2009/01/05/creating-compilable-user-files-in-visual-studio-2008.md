---
title: Creating Compilable User Files in Visual Studio 2008
published: true
date: 2009-01-05 22:30:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2009/01/05/creating-compilable-user-files-in-visual-studio-2008.aspx
file: creating-compilable-user-files-in-visual-studio-2008.aspx
path: /blogs/dotnet/archive/2009/01/05/
author: tom anderson
words: 819
---
Back in my c days when working on projects with other people in source control etc… I would do all kinds of weird things, things like include files for coding only for me.

Example:

#ifdef (TOM)   
include "Tom.h";   
#endif

This allowed me to call code that I didn't quite want out of my grasp yet, but I wanted in source control.  Naturally the "Tom.h" and its associated files where not in source control, or even included in the project. I kind of missed that feature, and investigated today how to get the same functionality.

Behold the *.proj.user files in Visual Studio 2008 (might work in 2005, who knows, I didn't try).

Normal source control will keep the .user files out of source control, as these are generally used to store debug information, local configurations, etc… There is nothing in the rulebook that stated I couldn't use this file to also compile "user" code. You can create a .user file simply by adding it in the same directory as your *.csproj file with the <Project> tags.  The *.csproj files (and *.vbproj files) are simply msbuild file format, so you can check up on the [MSDN documentation][1] on how to work with it.

My "Sandbox.csproj.user" file:

       1:  <Project xmlns="http://schemas.microsoft.com/developer/msbuild/2003">    
       2:    <PropertyGroup Condition=" '$(Configuration)|$(Platform)' == 'Debug|AnyCPU' ">    
       3:      <UserConfig>Tom</UserConfig>    
       4:      <DefineConstants>DEBUG;TRACE;TOM</DefineConstants>    
       5:    </PropertyGroup>    
       6:    <ItemGroup Condition=" '$(UserConfig)' == 'Tom'">    
       7:      <Compile Include="Tom.cs" />    
       8:    </ItemGroup>    
       9:  </Project>

So, basically what is going on here is that I added a new property in the Debug configuration for "UserConfig", this is used later for a conditional, and thats about it.  I then added "TOM" to the defined constants.

I then added an ItemGroup with the condition of the UserConfig being "Tom", clever eh?  Inside this ItemGroup I simply added a Compile for Tom.cs.  This file is pretty barebones, and does nothing, but here it is anyway.

       1:  using System;    
       2:  using System.Collections.Generic;    
       3:  using System.Linq;    
       4:  using System.Text;    
       5:       
       6:  namespace Sandbox    
       7:  {    
       8:      class Tom    
       9:      {    
      10:      }    
      11:  }

Then in my main application, I simply do an #if on the TOM constant and insert my code in that block.

       1:  #if (TOM)    
       2:              Tom t = new Tom();    
       3:  #endif

Pretty nifty eh?

What is even better, if I switch to release mode Line 2 above grays out (letting me know it isn't compiled) and my "Tom" class isn't included in the compile either (verified with Reflector).

The only issue is that it doesn't show up in the solution explorer, but neither did the c counter parts, this also has the nice beauty of not adding the file to source control, so now I can do all that "tom machine" specific coding I need to to bang out that gong feature.

![image][2]

So, if you want to do some code and keep it checked in without screwing everyone else up, or have your own set of tools you want to be stingy with, this method might work great for you!

 

***Warning**: Code in the #if (TOM) blocks will be visible in source control, no one will have the compiled version of it, and unless they declare the constant, their code will be safe as code in conditional blocks that don't meet the condition don't get compiled.



[1]: http://msdn.microsoft.com/en-us/library/wea2sca5.aspx
[2]: ./creating-compilable-user-files-in-visual-studio-2008/image_thumb_634A38D6.png "image"


