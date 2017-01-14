---
title: CAB - Part 4 Adding a WorkSpace
published: true
date: 2008-08-08 21:15:29 +0000 UTC
tags: imported cab
original: http://renevo.com/blogs/dotnet/archive/2008/08/08/cab-part-4-adding-a-workspace.aspx
file: cab-part-4-adding-a-workspace.aspx
path: /blogs/dotnet/archive/2008/08/08/
author: tom anderson
words: 1440
---
In the [previous articles on CAB][1] we have gone over creating the basics with a CAB, right now our application really has no "Composite" to it, just "Application Block".  In this article we will go over creating our first WorkSpace for modules to dock into, which will lead us to our next article on creating a basic module with a single workspace.

The first thing we need to do, in order to work with the CAB WinForms controls, is to create a new Tab on our Forms Toolbox, and load up the Microsoft.Practices.CompositeUI.WinForms.dll controls to it.  If you need help on this, you can see many articles on the web on how to add items to your toolbox, like [this one][2].

You should now have 7 new controls, the one we are going to concentrate on for this article is the DeckWorkspace.  Simply drag it onto our form, set the Dock to Fill, name it uxPrimaryWorkspace, then clear the Text property.

Your form in designer should look like this:

![image][3]

The size of the form is un-important, later we will go about saving the forms settings, etc... but for now, lets leave it as is.

Now that we have our form setup, we need to add some ways to access this workspace, as well as register the WorkSpace with the CAB framework.  This is usually done via id strings, but since we want our users to properly access this workspace, we will want to create constants for access.  Instead of the modules accessing and referencing the main executable, we will add another project with the name ".Interfaces" that will contain all the constants for extension sites, workspaces, as well as internal services.

Add a Class Library project to the solution named "RenEvo.Blogs.Cab.Interfaces" with the default directory specified.  Next delete the Class1.vb that is automaticaly created.

Create a new folder called Common, then a new class file inside of that folder called WorkSpaces.  In the class file you need to implement a namespace, as well as a few constants inside of the class.  The namespace is only for better organization, and separation of objects.
    
    
       1:  Namespace Common.Constants
    
    
       2:   
    
    
       3:      ''' <summary>
    
    
       4:      ''' Class to store all of the workspace names
    
    
       5:      ''' </summary>
    
    
       6:      Public Class WorkSpaces
    
    
       7:          ''' <summary>
    
    
       8:          ''' The Primary Workspace for the CAB application
    
    
       9:          ''' </summary>
    
    
      10:          Public Const PrimaryWorkSpace As String = "PrimaryWorkSpace"
    
    
      11:   
    
    
      12:      End Class
    
    
      13:   
    
    
      14:  End Namespace

 

For now our class really doesn't cover a lot of constants, but in a larger application this could get quite huge.  Plus we no longer will have casing or mis-spelling issues with the other developers.

Now add a reference to the new project to the main Shell project, setting the copy to True.

Next, back in our form, we need to access the Constructor, and after the InitializeComponent, we want to reset the name of the uxPrimaryWorkspace in runtime to the name of the string in the preceding class.
    
    
       1:      Public Sub New()
    
    
       2:   
    
    
       3:          ' This call is required by the Windows Form Designer.
    
    
       4:          InitializeComponent()
    
    
       5:   
    
    
       6:          ' Add any initialization after the InitializeComponent() call.
    
    
       7:          Me.uxPrimaryWorkspace.Name = Interfaces.Common.Constants.WorkSpaces.PrimaryWorkSpace
    
    
       8:   
    
    
       9:      End Sub

 

This is a really easy implementation, probably the easiest in the entire application.  To test it, we can do the following in our Startup.vb class.

In the ShellShown event handlers, add the following lines of code to the bottom of the method.
    
    
       1:      Protected Sub ShellShown(ByVal sender As Object, ByVal e As EventArgs)
    
    
       2:          'remove the handler
    
    
       3:          RemoveHandler Shell.Shown, AddressOf ShellShown
    
    
       4:   
    
    
       5:          'set the cursor back to normal
    
    
       6:          m_Splash.Cursor = Cursors.Default
    
    
       7:   
    
    
       8:          'hide, dispose, and kill the splash 
    
    
       9:          m_Splash.Hide()
    
    
      10:          m_Splash.Dispose()
    
    
      11:          m_Splash = Nothing
    
    
      12:   
    
    
      13:          'test to see if we have a workspace
    
    
    **  14:          MessageBox.Show("Workspace Initialized: " & Boolean.Parse(Me.RootWorkItem.Workspaces(Interfaces.Common.Constants.WorkSpaces.PrimaryWorkSpace) IsNot Nothing))**
    
    
      15:      End Sub

When ran, the application will display "Workspace Initialized: True" in a message box just after the splash screen closes.

Now remove that line of code, as we don't want that message box to continue displaying.

As stated before, in the next article I will be covering the creation of a very simplistic module that will load into this newly created workspace.

[Download the Solution][4]

![][5]

[1]: http://www.renevo.com/blogs/dotnet/archive/tags/CAB/default.aspx
[2]: http://msdn.microsoft.com/en-us/library/ms165355(VS.80).aspx
[3]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/CABPart3AddingaWorkSpace_C821/image_thumb.png
[4]: http://www.renevo.com/files/folders/articles_vbnet/entry1992.aspx
[5]: http://renevo.com/aggbug.aspx?PostID=1993

