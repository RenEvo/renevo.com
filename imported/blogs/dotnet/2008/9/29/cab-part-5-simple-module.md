---
title: CAB – Part 5 Simple Module
published: true
date: 2008-09-29 05:56:00 +0000 UTC
tags: imported cab
original: http://renevo.com/blogs/dotnet/archive/2008/09/29/cab-part-5-simple-module.aspx
file: cab-part-5-simple-module.aspx
path: /blogs/dotnet/archive/2008/09/29/
author: tom anderson
words: 1813
---
In the [previous articles on CAB][1], we have gone over some of the basics of getting our shell application ready to load up some modules.

In this article, we will actually create a simplistic Module to display a "SmartPart" in our primary workspace.

The first thing we need to do, is to add a new "Class Library" project to our current solution. Name the new project "RenEvo.Blogs.Cab.SimpleModule" and hit ok. Rename the "Class1.vb" to "SimpleModuleInit.vb".  This class will be our entry point into the actual module.

At this point, we are going to make a few modifications to our current projects, as the build directories are kind of all over the place, and since we have a defined structure we need to adhere to (example: ./Modules/ for any modules loaded into the Shell).

In the RenEvo.Blogs.Cab project, set the Build output Path in the Compile tab of the project properties to "..binDebug" for Debug configuration, and "..binRelease" for the Release configuration.  Do the same thing to the RenEvo.Blogs.Cab.Interfaces project.

For the new RenEvo.Blogs.Cab.SimpleModule, set the output paths to "..binDebugModules" and "..binReleaseModules" respectively.

Build the entire solution, and verify that there is a bin directory in the Solution folder that contains all of the projects output assemblies in the per-configuration folder.

Back to our Simple Module project, we need to add some key references. Add the following references, and be sure to set the "Copy Local" property to False for each reference.

* Microsoft.Practices.CompositeUI
* Microsoft.Practices.CompositeUI.WinForms
* Microsoft.Practices.ObjectBuilder
* System.Windows.Forms
* System.Drawing

And then finally add a project reference to the RenEvo.Blogs.Cab.Interfaces, also settings the "Copy Local" property to false.

![image][2]

Next, create a few new folders inside of the Simple Module project.

* Services
* SmartParts
* WorkItems

The first thing we need to do is to setup the SimpleModuleInit class so that the shell will load it up. The first thing we will do is to import a few namespaces, then add an assembly attribute "Module" which will be the attribute that the Shell will look for when loading modules via reflection. Next we will inherit ModuleInit in our class.  At this point, your code will look like this:
    
    
       1:  Imports Microsoft.Practices.CompositeUI
    
    
       2:  Imports Microsoft.Practices.CompositeUI.Services
    
    
       3:   
    
    
       4:  <Assembly: [Module]("Simple Module")> 
    
    
       5:   
    
    
       6:  Public Class SimpleModuleInit
    
    
       7:      Inherits ModuleInit
    
    
       8:   
    
    
       9:   
    
    
      10:  End Class

We will have to setup our Dependency Injection property in order to integrate into the Shell Application. If you are not familiar with [Dependency Injection][3], I suggest you read up on it before continuing.  Essentially for this implementation of Injection, we will use Property Attributes to notify the base framework that we have a service dependency, and then the property declared type will provide the service type we are accessing, in this case "WorkItem" type.
    
    
       1:  #Region " Root Work Item Dependency Injection "
    
    
       2:   
    
    
       3:      Private m_WorkItem As WorkItem = Nothing
    
    
       4:      <ServiceDependency()> _
    
    
       5:      Public Property RootWorkItem() As WorkItem
    
    
       6:          Get
    
    
       7:              Return m_WorkItem
    
    
       8:          End Get
    
    
       9:          Set(ByVal value As WorkItem)
    
    
      10:              m_WorkItem = value
    
    
      11:          End Set
    
    
      12:      End Property
    
    
      13:   
    
    
      14:  #End Region

This is basically setting this property by doing the following line of code (not used in our code, but deep behind the scenes).
    
    
       1:  SimpleModuleInit.RootWorkItem = Services.Get(Of WorkItem)()

By using Dependency Injection, we don't need to know what is going on behind the scenes, and in some later examples of implementing our own services, it is quite useful to just have what we need automatically populated, rather than trying to find it manually.

Moving along, the next step we will need to do, is to override the ModuleInit.Load() sub in our SimpleModuleInit. Within this Method, lets just simple do the MessageBox.Show("SimpleModule Loaded") to verify that our module is being loaded.

Be sure that the RenEvo.Blogs.Cab project is set as the Startup Project, and run the application.  While the Splash Screen is displayed, you should see a message box stating that your SimpleModule Loaded.

Ok, moving along quickly here, lets remove that message box line. Now in the SmartParts folder, add a new UserControl named "SimpleSmartPart.vb" and in the WorkItems folder, add a new Class named "SimpleWorkItem.vb"

Code for SimpleSmartPart.vb
    
    
       1:  Imports Microsoft.Practices.CompositeUI.SmartParts
    
    
       2:   
    
    
       3:  <SmartPart()> _
    
    
       4:  Public Class SimpleSmartPart
    
    
       5:   
    
    
       6:  End Class

What we are doing in the class above is just tagging this user control as a SmartPart, this allows the WorkSpace objects to communicate with them properly, which we will go over in a later article.

For the sake of being able to see the smart part easier, drop a label control in the design view on the SimpleSmartPart, name it "uxSmartPartName" and set its Text property to "Simple Smart Part Loaded". Where you put the label is irrelevant, mine looks like this:

![image][4] 

Now, the code for the SimpleWorkItem.vb is as follows:
    
    
       1:  Imports Microsoft.Practices.CompositeUI
    
    
       2:  Imports Microsoft.Practices.CompositeUI.Commands
    
    
       3:  Imports RenEvo.Blogs.Cab.Interfaces.Common
    
    
       4:   
    
    
       5:  Imports System.Windows.Forms
    
    
       6:   
    
    
       7:  Public Class SimpleWorkItem
    
    
       8:      Inherits WorkItem
    
    
       9:   
    
    
      10:      Protected Overrides Sub OnRunStarted()
    
    
      11:          MyBase.OnRunStarted()
    
    
      12:   
    
    
      13:          Me.RootWorkItem.Workspaces(Constants.WorkSpaces.PrimaryWorkSpace).Show(New SimpleSmartPart)
    
    
      14:   
    
    
      15:      End Sub
    
    
      16:   
    
    
      17:  End Class

The code above inherits from WorkItem and overrides the OnRunStarted method, we then Show a new SimpleSmartPart in our PrimaryWorkSpace.

Finally, to get the work item to actually "run" we need to modify the SimpleModuleInit.vb by adding a workitem to the RootWorkItem property, and calling the "Run" method.  The code below goes in the Load() method in SimpleModuleInit.vb
    
    
       1:  Me.RootWorkItem.WorkItems.AddNew(Of SimpleWorkItem).Run()

In the above code we are using the "AddNew" method with a generic typing of SimpleWorkitem and then calling the method "Run" on the newly created WorkItem.

Save everything, and run the application.  When the Shell interface displays after the splash screen, you will see the label with the text "Simple Smart Part Loaded" in our interface.

![image][5]

And there you have it, we have loaded up a module into the Shell, and displayed our First Smart Part.

In the next article, we will be adding a Controller to the SmartPart for the MVC (Model View Controller) design pattern, and implementing a few commands within that Controller.

[Download the Solution][6]

![][7]

[1]: http://www.renevo.com/blogs/dotnet/archive/tags/CAB/default.aspx
[2]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/CABPart5SimpleModule_14E26/image_thumb.png "image"
[3]: http://www.martinfowler.com/articles/injection.html
[4]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/CABPart5SimpleModule_14E26/image_thumb_1.png "image"
[5]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/CABPart5SimpleModule_14E26/image_thumb_2.png "image"
[6]: http://www.renevo.com/files/folders/articles_vbnet/entry2027.aspx
[7]: http://renevo.com/aggbug.aspx?PostID=2028

