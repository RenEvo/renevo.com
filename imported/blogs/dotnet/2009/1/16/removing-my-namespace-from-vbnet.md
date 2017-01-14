---
title: Removing My Namespace from VB.Net
published: true
date: 2009-01-16 17:45:49 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2009/01/16/removing-my-namespace-from-vb-net.aspx
file: removing-my-namespace-from-vb-net.aspx
path: /blogs/dotnet/archive/2009/01/16/
author: tom anderson
words: 789
---
Recently a question arose on [Stackoverflow][1] that asked [if you could remove the My namespace from vb.net][2].

So, before I get going with this article, I want to state that the [My Namespace does have a few uses][3], it provides instant access to resources, settings, and quick environment settings.

I also want to state that **THIS IS NOT A REQUIRED FEATURE IN VB.NET**. Did I stress that yet? It is however a default feature in VB.Net.

Anyway, lets get on with it.

First things first, backup your project directory, I don't want to be responsible for you deleting any of your settings or resources because you want to remove something you are using.

Now that you did that (right?) lets move to the solution explorer and click on the option to "Show all files".

![image][4]

Expand the "My Project" Node and select the "Application.myapp", "Resources.resx", and "Settings.settings" nodes. When I say nodes, that means items below "My Project", this is a treeview.

![image][5]

Now, hit delete. This will remove any of the "My" code that has automatically already been added to your project. Go ahead and click on the "Show all files" button again to get back to a clean view.

Next, double click on "My Project" and navigate to the Compile tab, and click on the "Advanced Compile Options". This dialog has all kinds of fun stuff in it, but we are only worried about one particular setting. Go ahead and click on "Enable Optimizations".

![image][6]

Click "OK" and we are about 50% done.

Now, in the solution explorer, right click on your project and select "Unload Project".

![image][7]

This will unload your project from the IDE, but retain the reference to it, another great thing about it is it allows us to edit the .vbproj file directly instead of through the UI, which is what is required for us to do the next step.

![image][8]

Look for the <MyType> xml tag, we need to set this to "Empty", not Empty, but with the value of "Empty". You may also need to change the <StartupObject> tag to reflect your main form if it is currently set to "My.Application".

Below is the XML from the first PropertyGroup after modifying it.
    
    
       1:    <PropertyGroup>
    
    
       2:      <Configuration Condition=" '$(Configuration)' == '' ">Debug</Configuration>
    
    
       3:      <Platform Condition=" '$(Platform)' == '' ">AnyCPU</Platform>
    
    
       4:      <ProductVersion>9.0.30729</ProductVersion>
    
    
       5:      <SchemaVersion>2.0</SchemaVersion>
    
    
       6:      <ProjectGuid>{99D23E3F-D6D5-467F-AB1D-A594E40F4378}</ProjectGuid>
    
    
       7:      <OutputType>WinExe</OutputType>
    
    
       8:      <StartupObject>RemoveMyNamespace.Form1</StartupObject>
    
    
       9:      <RootNamespace>RemoveMyNamespace</RootNamespace>
    
    
      10:      <AssemblyName>RemoveMyNamespace</AssemblyName>
    
    
      11:      <FileAlignment>512</FileAlignment>
    
    
      12:      **<MyType>Empty</MyType>**
    
    
      13:      <TargetFrameworkVersion>v3.5</TargetFrameworkVersion>
    
    
      14:      <OptionExplicit>On</OptionExplicit>
    
    
      15:      <OptionCompare>Binary</OptionCompare>
    
    
      16:      <OptionStrict>Off</OptionStrict>
    
    
      17:      <OptionInfer>On</OptionInfer>
    
    
      18:    </PropertyGroup>

I bolded the change.

Save the file and close it, now right click on the project in the solution explorer and reload the project.

![image][9]

Compile it, and you can now view it in [Reflector][10] to see that the "My" namespace is completely removed.

![image][11]

We have now successfully remove the "My" namespace.

![image][12]

The keyword still exists, but it is now removed completely from your project.

![kick it on DotNetKicks.com][13]![][14]

[1]: http://stackoverflow.com
[2]: http://stackoverflow.com/questions/451273/why-does-my-vb-net-class-library-show-my-and-my-resources-namespaces-in-reflecto
[3]: http://msdn.microsoft.com/en-us/magazine/cc163972.aspx
[4]: http://www.renevo.com/blogs/dotnet/image_thumb_3CE067BF.png "image"
[5]: http://www.renevo.com/blogs/dotnet/image_thumb_67BB5C82.png "image"
[6]: http://www.renevo.com/blogs/dotnet/image_thumb_583FFCC6.png "image"
[7]: http://www.renevo.com/blogs/dotnet/image_thumb_1014D19B.png "image"
[8]: http://www.renevo.com/blogs/dotnet/image_thumb_16CA6BCF.png "image"
[9]: http://www.renevo.com/blogs/dotnet/image_thumb_4CCEEADC.png "image"
[10]: http://www.red-gate.com/products/reflector/
[11]: http://www.renevo.com/blogs/dotnet/image_thumb_63249021.png "image"
[12]: http://www.renevo.com/blogs/dotnet/image_thumb_3539CE1A.png "image"
[13]: http://www.dotnetkicks.com/Services/Images/KickItImageGenerator.ashx?url=http://www.renevo.com/blogs/dotnet/archive/2009/01/16/removing-my-namespace-from-vb-net.aspx
[14]: http://renevo.com/aggbug.aspx?PostID=2138

