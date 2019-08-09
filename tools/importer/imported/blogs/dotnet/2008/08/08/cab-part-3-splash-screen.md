---
title: CAB - Part 3 Splash Screen
published: true
date: 2008-08-08 19:38:51 +0000 UTC
tags: imported cab
original: http://renevo.com/blogs/dotnet/archive/2008/08/08/cab-part-3-splash-screen.aspx
file: cab-part-3-splash-screen.aspx
path: /blogs/dotnet/archive/2008/08/08/
author: tom anderson
words: 984
---
In the [previous articles for CAB][1] we have gone over creating our base form, as well as implementing a new way to load our modules (once we create some), in this article we are going to go over another simple implementation (that is specifically not covered in the documentation) to make our application much more user friendly.

Splash Screens.  As you know, in Visual Studio 2005 and 2008 you have the ability to enable the Application Framework, and simply select a Splash Screen from the project properties, in a CAB shell, you must run through the Sub Main() instead of simply using a Form for startup.  We will add one real fast that displays while the CAB is loading its assemblies and modules.

First, add a new form to our project (under Forms->Dialogs in our directory structure), select the Splash Screen form template, and name it CabSplashScreen.  For now we aren't going to modify the Splash Screen, just use the default template.

Just close the Splash Screen, and forget about it for now.

In our Startup class, we need to create a new private field for the splash screen, add creation to it in the sub new (non-shared instance), override the AfterShellCreated, as well as create an event handler for ShellShown.

Below is the code to do just that:

       1:      Private m_Splash As CabSplashScreen = Nothing    
       2:       
       3:      Public Sub New()    
       4:          'create the splash    
       5:          m_Splash = New CabSplashScreen    
       6:       
       7:          'show and update the splash    
       8:          m_Splash.Show()    
       9:          m_Splash.Update()    
      10:       
      11:          'set the cursor to the hourglass    
      12:          m_Splash.Cursor = Cursors.WaitCursor    
      13:       
      14:          'let splash screen events process    
      15:          Application.DoEvents()    
      16:      End Sub    
      17:       
      18:      Protected Overrides Sub AfterShellCreated()    
      19:          MyBase.AfterShellCreated()    
      20:       
      21:          'add an event handler for the ShellForm.Show event (this is when we will kill our splash)    
      22:          AddHandler Shell.Shown, AddressOf ShellShown    
      23:       
      24:      End Sub    
      25:       
      26:      Protected Sub ShellShown(ByVal sender As Object, ByVal e As EventArgs)    
      27:          'remove the handler    
      28:          RemoveHandler Shell.Shown, AddressOf ShellShown    
      29:       
      30:          'set the cursor back to normal    
      31:          m_Splash.Cursor = Cursors.Default    
      32:       
      33:          'hide, dispose, and kill the splash     
      34:          m_Splash.Hide()    
      35:          m_Splash.Dispose()    
      36:          m_Splash = Nothing    
      37:      End Sub

The code is pretty straight forward, and I will leave you to the Comments to figure out what it is doing (I am sure that if you are venturing into creating a CAB shell then this is nothing new to you)

In our next article, we will go over creating our first WorkSpace.

[Download the Solution][2]



[1]: http://www.renevo.com/blogs/dotnet/archive/tags/CAB/default.aspx
[2]: http://www.renevo.com/files/folders/articles_vbnet/entry1988.aspx


