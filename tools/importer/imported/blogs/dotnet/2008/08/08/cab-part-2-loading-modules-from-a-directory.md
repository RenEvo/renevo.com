---
title: CAB - Part 2 Loading Modules from a directory
published: true
date: 2008-08-08 19:22:59 +0000 UTC
tags: imported cab
original: http://renevo.com/blogs/dotnet/archive/2008/08/08/cab-part-2-loading-modules-from-a-directory.aspx
file: cab-part-2-loading-modules-from-a-directory.aspx
path: /blogs/dotnet/archive/2008/08/08/
author: tom anderson
words: 539
---
In the [first part of the CAB articles][1], I explained a bit about getting the CAB working in a generic shell, and simply initiating the framework, in this article I will explain a little known ability with the CAB to dynamically load all modules in a directory, instead of the "ProductCatalog.xml" that is used in the CAB samples.  This is something that I do personally in all 3 of our CAB shells, and have found it to be much easier to deploy new modules to customers.

In our Startup class, we simply need to override one of the base WorkItem methods for adding services "AddServices", remove teh IModuleEnumerator that is currently loaded (that productcatalog.xml loader) and replace it with a new ReflectionModuleEnumerator instead.

       1:      Protected Overrides Sub AddServices()    
       2:          'setup base services    
       3:          MyBase.AddServices()    
       4:       
       5:          'remove the FileModule for ProductCatalog.xml    
       6:          MyBase.RootWorkItem.Services.Remove(Of IModuleEnumerator)()    
       7:       
       8:          'create a new reflection enumerator    
       9:          Dim reflectionEnumerator As New ReflectionModuleEnumerator()    
      10:       
      11:          'make sure our path exists for modules    
      12:          Directory.CreateDirectory(Path.GetDirectoryName(Application.ExecutablePath) & "Modules")    
      13:       
      14:          'set the working path to ./Modules/    
      15:          reflectionEnumerator.BasePath = Path.GetDirectoryName(Application.ExecutablePath) & "Modules"    
      16:       
      17:          'add it back to the services    
      18:          MyBase.RootWorkItem.Services.Add(GetType(IModuleEnumerator), reflectionEnumerator)    
      19:       
      20:      End Sub

In the next article, I will go over how to add a Splash Screen to the main form, since we have not enabled the "Application Framework".

[Download Solution][2]



[1]: http://www.renevo.com/blogs/dotnet/archive/2007/11/05/cab-getting-started.aspx
[2]: http://www.renevo.com/files/folders/articles_vbnet/entry1985.aspx


