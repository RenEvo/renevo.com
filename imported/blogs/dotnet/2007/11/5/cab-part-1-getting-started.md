---
title: CAB - Part 1 Getting Started
published: true
date: 2007-11-05 15:22:16 +0000 UTC
tags: imported cab
original: http://renevo.com/blogs/dotnet/archive/2007/11/05/cab-getting-started.aspx
file: cab-getting-started.aspx
path: /blogs/dotnet/archive/2007/11/05/
author: tom anderson
words: 851
---
Well, here goes nothing.  This will be the first part of a series of blog entries explaining how to get a Composite UI Application Block up and running from scratch.

Where do we start? First thing we will want to do is to [Download the Composite UI Application Block][1] from Microsoft.  This download requires that you login to Microsoft.

The CAB_VB.msi is about 6mb, and installs pretty easily. While installing you will go through a Custom Setup.  You should choose to install the NUnit 2.2 Unit Tests, and not install the Team System Unit Tests (Unless you actually have Visual Studio Team System). Once you complete the installation process, you are ready to go.

 

**Compiling the CAB**

Before we begin, we are going to have to compile the Composite Application Block.  Open up the Solution that was installed above in Visual Studio (Default Location is: C:Program FilesMicrosoft Composite UI App BlockVBCompositeUI.sln).

For now, the only Solution Folder we are concerned about is the "Source" folder. Go ahead and close the others. Change the Build Configuration to Release, right click on the CompositeUI.WinForms project and select "Rebuild". Now navigate to the binRelease folder for the CompositeUI.WinForms project (Default Location is: C:Program FilesMicrosoft Composite UI App BlockVBSourceCompositeUI.WinFormsbinRelease) and verify these assemblies are there:

* Microsoft.Practices.ObjectBuilder.dll 
* Microsoft.Practices.CompositeUI.dll 
* Microsoft.Practices.CompositeUI.WinForms.dll 

Keep this window open, as we are going to want to copy these files over to our project directory for referencing.

 

**Create a new project**

We are going to just jump right in, and create a new Windows Application project in Visual Studio 2005 or VB Express 2005. Lets name it "RenEvo.Blogs.CAB".

The first thing we want to do is delete the Form1.vb that is created by the project wizard, next, open up the Project Properties then uncheck the "Enable application framework" and change the Startup object: to Sub Main.  Now is also a good time to go ahead and set your Assembly Information.

Next, lets setup our References, from the binRelease folder we have open from above (CompositeUI.WinForms), lets copy those three assemblies and their debug files to a new directory inside of our solution directory called "CAB Assemblies". Open up the Add Reference dialog and click on the Browse tab.  While using the "up arrow", go up one directory, and into the CAB Assemblies directory.  Select all 3 Assemblies in this directory. Go ahead and leave the default reference properties for now.

Next, we are going to want to add a Startup Object, Add a new Class to your project, and call it "Startup.vb". Create a Shared Sub Main which will act as our entry point for the application and we are going to add the STAThreadAttribute to this method.  Your class file should look like below.

 

Public Class Startup 

    <STAThread()> _   
    Public Shared Sub Main() 

    End Sub 

End Class

 

As you will notice, if you press F5 to run, the application will simply open and close, this is intentional for now.

Next we are going to add a new Form, this will be our Main GUI for the application, Create a new folder called "Forms", then add a new Form called "CabApplicationForm". For now, close the form and go back to the Startup class.

With the startup class, we are going to call and initialize a reference to its self, but first we need to add some inheritance and imports to the CompositeUI Application Block.  Add the following imports to the Startup.vb.

Microsoft.Practices.CompositeUI   
Microsoft.Practices.CompositeUI.WinForms

Next, inherit the FormShellApplication with a type constructor of WorkItem and CabApplicationForm as well as adding an implementation of IDisposable (I prefer to do this with just about everything). And finally, lets create an instance of the Startup class in the Sub Main, and call the Run method.  Your Startup.vb should look like below now.

 

Imports Microsoft.Practices.CompositeUI   
Imports Microsoft.Practices.CompositeUI.WinForms 

Public Class Startup   
    Inherits FormShellApplication(Of WorkItem, CabApplicationForm) : Implements IDisposable 

    <STAThread()> _   
    Public Shared Sub Main()   
        Using appRunner As New Startup()   
            appRunner.Run()   
        End Using   
    End Sub 

#Region " IDisposable Support "

    Private disposedValue As Boolean = False        ' To detect redundant calls 

    ' IDisposable   
    Protected Overridable Sub Dispose(ByVal disposing As Boolean)   
        If Not Me.disposedValue Then   
            If disposing Then 

            End If 

        End If   
        Me.disposedValue = True   
    End Sub 

    ' This code added by Visual Basic to correctly implement the disposable pattern.   
    Public Sub Dispose() Implements IDisposable.Dispose   
        ' Do not change this code.  Put cleanup code in Dispose(ByVal disposing As Boolean) above.   
        Dispose(True)   
        GC.SuppressFinalize(Me)   
    End Sub 

#End Region 

End Class

 

Now when you press F5 to launch the application, the Form will display, and we have initialized the CAB framework. Granted at this time, it isn't really doing much.

**Conclusion**

This is the first step, the first step to many. In the next blogs I will go into actually adding some functionality to the CAB block, this was just a quick way to get your environment going.

[Download the Solution][2]

![][3]

[1]: http://www.microsoft.com/downloads/details.aspx?familyid=B619276A-2E9E-47C6-8893-F8E5F88FD8DD&displaylang=en
[2]: http://www.renevo.com/files/folders/articles_vbnet/entry1582.aspx
[3]: http://renevo.com/aggbug.aspx?PostID=1583

