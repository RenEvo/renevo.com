---
title: Creating a ListView for Vista (Tiles Extended Support)
published: true
date: 2007-10-29 18:07:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2007/10/29/creating-a-listview-for-vista-tiles-extended-support.aspx
file: creating-a-listview-for-vista-tiles-extended-support.aspx
path: /blogs/dotnet/archive/2007/10/29/
author: tom anderson
words: 1229
---
Have you ever wondered why Windows Vista's ListViews and .Net 2.0's ListViews just don't match?  Well, after a lot of hunting on the web, and finding the best implementations, I have put together this little tutorial to turn this:

![ListView.NoStyle][1]

Into this:

![ListView.Style][2]

**Building the project**

Create a new Class Library Project somewhere on your hard drive, for now, lets name it _"RenEvo.Articles.Controls"_. As I state always, lets get rid of the Class1.vb that was created with the new project by right clicking and deleting it.

****

**Adding the needed references**

Next lets double click on "My Project" and bring up the Project Properties Editor. Click on the References tab, and add the following two references from the .Net tab.

System.Windows.Forms   
System.Drawing

At this time you can do a global import on the namespaces, but for this article I am going to be assuming you didn't.

**Creating our custom control**

Now that we have our environment setup, lets add a new Class by right clicking on the project "Add New Item", select _Class_ from the items in the list, and name it _VistaListView.vb._ You should now have a blank class.

Lets go ahead and add some imports to the top of the class.

Imports System.Windows.Forms   
Imports System.ComponentModel   
Imports System.Drawing

We are now going to add some inheritance to the class by adding "Inherits ListView" under the class declaration.

****

**Adding a toolbox bitmap**

Since we don't like the look of the default gear icon in the Design Toolbox, lets implement the Icon from the Listview to show. Add the following attribute tag above the class declaration.

<ToolboxBitmap(GetType(ListView))> _   
Public Class VistaListView

As you can see, we have hadded the "ToolboxBitmap" attribute, and used the Type constructor to select the ToolboxBitmap from the ListView object. There are a few other overloads for the ToolboxBitmap attribute, but we aren't going to cover those here.

**Declaring an API**

Next we are going to quickly declare an API to update the style on the control, the method that we are looking for is in the uxtheme.dll and is called "SetWindowTheme". Below is the declaration.

<System.Runtime.InteropServices.DllImport("uxtheme.dll", Charset:=System.Runtime.InteropServices.CharSet.Unicode)> _   
Private Shared Function SetWindowTheme(ByVal Handle As IntPtr, ByVal Theme As String, ByVal SubIdList As String) As Integer   
End Function

We have made this method private, as we don't need to access it externally.  We will call this function with the handle of the control to set, the theme to apply, and then we don't need to worry about the third argument at this time and will pass a null value.

**Creating a Vista Check**

We are going to want to detect wether we are running on Vista, or an earlier build of Windows.  This is to check and see if we should even try to call the API, or just ignore it and let the regular styles be.  Vista is version 6.* of windows, so we will simply return if the OS version is 6 or greater.

Private Function IsVistaOrLater() As Boolean   
    Return (Environment.OSVersion.Version.Major >= 6)   
End Function

Again, we made this function private, since we don't need to use it anywhere else, later on you may wish to implement these two functions into a Utilities or API class.

**Overrides**

We are going to want to override one very key method, the _OnHandleCreated_ this method is called right after the class is created by the managed code. We will call the base method, check our os version, then if it is Vista or higher, we will set the theme to "explorer".  There are several themes, but this is the one we want to achieve.

Protected Overrides Sub OnHandleCreated(ByVal e As System.EventArgs)   
    MyBase.OnHandleCreated(e) 

    If IsVistaOrLater() Then   
        SetWindowTheme(Me.Handle, "explorer", Nothing)   
    End If   
End Sub

Now we could drop this control on a form and run it right now, and everything would be great, but there is one more feature we should adopt.

**Tile View Support**

We are going to add another override real fast that will force a set on the View property, I will explain this in a moment.

Protected Overrides Sub OnVisibleChanged(ByVal e As System.EventArgs)   
    MyBase.OnVisibleChanged(e)   
    'enforce a set of the view property   
    View = View   
End Sub

Next we are going to shadow the View property, check to see if we are not running on vista, if we are setting Tile, and that we are not in design mode. Then if all of these are correct, we will set the view property to LargeIcon.  This will do the check internally, instead of letting windows do it for you and cause a bit of an internal mess. Personally, I like knowing why things happen, and not let any assumptions go through. We are also going to extend the current description of the property and mention that Tile is not supported on Windows XP.

<Description("Selects one of five different views that items can be shown in. (Tile is not supported in Windows XP).")> _   
Public Shadows Property View() As System.Windows.Forms.View   
    Get   
        Return MyBase.View   
    End Get   
    Set(ByVal value As System.Windows.Forms.View)   
        If value = Windows.Forms.View.Tile AndAlso IsVistaOrLater() = False AndAlso Me.DesignMode = False Then   
            value = Windows.Forms.View.LargeIcon   
        End If   
        MyBase.View = value   
    End Set   
End Property

**An explanation of Shadows**

Shadows is a declaration type that overrides properties and functions by completely replacing the base properties, methods, etc...  This is useful for properties and such that don't have and Overridable declaration.  One such method that I like to use is to modify collection to return the newly added object on "Collection.Add".

**Conclusion**

Now that we have contructed the class, you should be able to compile it and drop the ListView on a form.  If running on Vista you will notice that the tiles are now properly fully selected, and if running on XP you will see the items in LargeIcon view.

**Full Class** \- [Code 2 Html][3]

Imports System.Windows.Forms   
Imports System.ComponentModel   
Imports System.Drawing   
  
<ToolboxBitmap(GetType(ListView))> _   
Public Class VistaListView   
    Inherits ListView   
  
    <System.Runtime.InteropServices.DllImport("uxtheme.dll", Charset:=System.Runtime.InteropServices.CharSet.Unicode)> _   
    Private Shared Function SetWindowTheme(ByVal Handle As IntPtr, ByVal Theme As String, ByVal SubIdList As String) As Integer   
    End Function   
  
    Private Function IsVistaOrLater() As Boolean   
        Return (Environment.OSVersion.Version.Major >= 6)   
    End Function   
  
    Protected Overrides Sub OnHandleCreated(ByVal e As System.EventArgs)   
        MyBase.OnHandleCreated(e)   
  
        If IsVistaOrLater() Then   
            SetWindowTheme(Me.Handle, "explorer", Nothing)   
        End If   
    End Sub   
  
    Protected Overrides Sub OnVisibleChanged(ByVal e As System.EventArgs)   
        MyBase.OnVisibleChanged(e)   
        'xp overrides of view   
        View = View   
    End Sub   
  
    <Description("Selects one of five different views that items can be shown in. (Tile is not supported in Windows XP).")> _   
    Public Shadows Property View() As System.Windows.Forms.View   
        Get   
            Return MyBase.View   
        End Get   
        Set(ByVal value As System.Windows.Forms.View)   
            If value = Windows.Forms.View.Tile AndAlso IsVistaOrLater() = False AndAlso Me.DesignMode = False Then   
                value = Windows.Forms.View.LargeIcon   
            End If   
            MyBase.View = value   
        End Set   
    End Property   
  
End Class   



[1]: ./creating-a-listview-for-vista-tiles-extended-support/ListView.NoStyle_thumb.jpg
[2]: ./creating-a-listview-for-vista-tiles-extended-support/ListView.Style_thumb.jpg
[3]: http://puzzleware.net/CodeHtmler/default.aspx


