---
title: Creating Your Own Community Server Skinned Control
published: true
date: 2006-12-19 17:00:49 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2006/12/19/creating-your-own-community-server-skinned-control.aspx
file: creating-your-own-community-server-skinned-control.aspx
path: /blogs/developer/archive/2006/12/19/
author: tom anderson
words: 2500
---
In my [last article][1], I went into creating your own tabbed page that kept with the theme system of Community Server, today I will take it a step further and begin explaining on how to create a new control that will take advantage of the Community Server "skin" system.  First and foremost, you "should" have a local running copy of CommunityServer if you plan on doing any development for it, if not, now is a good time to set one up really fast.

Lets start out by copying our Default.aspx from our last exercise and naming it SkinnedControl.aspx in the Sample Directory.

Open up the newly created aspx and delete the "samplepageContentPart" CS:ContentPart control.  This will allow us to start from scratch with this article, feel free to leave the side menu, as it will be used later for linking the two pages the easy way.

Next, double click on the right content part and add a link to the Default.aspx and the SkinnedControl.aspx for ease of navigating between the sample controls.  This can be done with the following **html** code.

_<BR>  
<H4 class=CommonSidebarHeader>Sample Page Links</H4>  
<DIV class=CommonSidebarContent>  
<ul>  
<li><a href="Default.aspx">Sample Page</a></li>  
<li><a href="SkinnedControl.aspx">Skinned Control</a></li>  
</ul>  
</DIV>_

Now we have a nice side menu to be able to navigate between our sample pages, in future articles, you can simply add to this link list.

**Creating our Web Class**

This is where a decision needs to be made by you, you can code the control in any of the .Net languages, and for this article I have provided a VB and a C# project for the code.  You can download either [VB Express][2] or [C# Express][3] to complete this article.

Open up Visual Studio, and create a new Class Library Project, name it "RenEvo.CommunityServer.Articles" and click **OK**.

First and foremost, rename "Class1" to "SampleControl". In VB this may or may NOT change the class name in code, if not, go ahead and change it now.

Click on References, add Reference, and then Browse to your CommunityServer website's bin folder from the Browse tab, and select the following assemblies.

* CommunityServer.Components.dll
* CommunityServer.Controls.dll

Once again, Click on References, add Reference, and select System.Web from the .Net tab.

And finally, we need to setup our build environment under C#, open up the properties window, select the Build tab, and browse to your CommunityServer's **bin** directory.  This will be used to essentially deploy your assembly for debugging.  In VB, double click on "My Project", select the Compile tab, and browse for the CommunityServer's **bin** directory.  You can alternatively setup the Release Configuration to build to the same directory, but since this is a sample, I will not go over the recommended flags and settings for release builds.

Lets open up our SampleControl class code.

To save some typing, we are going to import some namespaces into our class.

**_C#_**

_using System;  
using System.Text;  
using System.Web.UI;  
using System.Web.UI.WebControls;  
using System.Collections.Generic;  
using CommunityServer.Controls;  
using CommunityServer.Components;_

**_VB_**

_Imports System.Text  
Imports System.Web.UI  
Imports System.Web.UI.WebControls  
Imports CommunityServer.Controls  
Imports CommunityServer.Components_

Next, we will inherit the CommunityServer's TemplatedWebControl

**_C#_**

_public class SampleControl:TemplatedWebControl_

**_VB_**

_Public Class SampleControl : Inherits TemplatedWebControl_

Create a Region that will contain our Private Declarations, we will store a Title that will be settable in the aspx when creating the control on a page.

**_C#_**

_#region Private Declares  
String m_Title = ""  
#endregion_

**_VB_**

_#Region " Private Declares " _

_Private m_Title As String = "" _

_#End Region_

Before we go and build, we still have a bit more to do, so stay away from that build button.  At this point we have created a new project, added some references to CommunityServer, and added our "codebehind" for our skinned control.  What you have noticed if you have ever developed web pages is that we have not added a .ascx to our project, the reason behind this is that we are using the CommunityServer framework to build and display our control based on some settings that we will apply next for "skinning" a control.  What basically happens is that when a CommunityServer page creates your control, it then calls the "AttachChildControls" function to allow you to bind any internal controls that you have for your control manually via "FindControl".  This allows developers to only put items that they need in any new skins that they create.  It is VERY important to check for nulls on any controls you "think" are on the page, as they might not actually have been included in the skin. 

**Skinning Support**

Now that we have cracked our Class a bit, lets add some "skin" support to the code behind. 

Create a new region (hey, regions are fantastic ways to organize your code) with a title of "Skin Support", create a new protected variable called SkinnedControl of type WebControls.Repeater.  Do not set this to new, this will be our internal reference to our "imaginary" control in the skin. 

Next we will want to override two function in the base class (TemplatedWebControl) _OnInit_ and _AttachChildControls_ Both of these will be used for applying our skin at run time, and informing the framework exactly what files to use. _OnInit_ we will simply specify our ExternalSkinFileName and then call the base class's OnInit.  In _AttachChildControls_ we will attempt to set our SkinnedControl.  Below is the code snippets to accomplish this. 

**_C#_**

_#region Skin Support  
Repeater SkinnedControl; _

_protected override void OnInit(EventArgs e)  
{  
// set our skin name, this is the "./Themes/<currentskinname>/" folder  
base.ExternalSkinFileName = "Skin-SampleControl.ascx"  
base.OnInit(e);  
} _

_protected override void AttachChildControls()  
{  
// attach our reference to the control in the skin (if it exists)  
SkinnedControl = (Repeater)FindControl("SkinnedControl");  
} _

_#endregion_

**_VB_**

_#Region " Skin Support " _

_Protected SkinnedControl As WebControls.Repeater _

_Protected Overrides Sub OnInit(ByVal e As System.EventArgs)  
'set our skin name, this is the "./Themes/<currentskinname>/" folder  
MyBase.ExternalSkinFileName = "Skin-SampleControl.ascx"  
MyBase.OnInit(e)  
End Sub _

_Protected Overrides Sub AttachChildControls()  
'attach our reference to the control in the skin (if it exists)  
SkinnedControl = FindControl("SkinnedControl")  
End Sub _

_#End Region_

**Databinding**

Databinding is a part of everything that makes a website function dynamically, without it, it would be a plethera of hard-coded table generation and response.write statements (remember asp 3.0 and php?).  Since we aren't quite covering data layers etc... yet in this article, we will simply create a fopa table in code and use it as a binding source.  To do this, we are going to override the base class's _Databind_ method, call our base class databind function first, see if our SkinnedControl is set to an instance (to prevent errors, remember), create a sample datatable, then bind it to our control.  And once again, we will include this section inside of a region called "Databinding". 

**_C#_**

_#region Databinding _

_public override void DataBind()  
{  
base.DataBind(); _

_if (SkinnedControl != null)  
{  
// Create a sample table  
System.Data.DataTable dt = new System.Data.DataTable("SampleTable");  
// Add Columns  
dt.Columns.Add("Column1");  
dt.Columns.Add("Column2");  
dt.Columns.Add("Column3");  
// Add Rows  
dt.Rows.Add("1", "2", "3");  
dt.Rows.Add("4", "5", "6");  
dt.Rows.Add("7", "8", "9"); _

_// Bind it to our skinned control  
SkinnedControl.DataSource = dt;  
SkinnedControl.DataBind();  
}  
} _

_#endregion_

**_VB_**

_#Region " Databinding " _

_Public Overrides Sub DataBind()  
MyBase.DataBind() _

_If Not SkinnedControl Is Nothing Then  
'Create a sample table  
Dim dt As New DataTable("SampleTable")  
'Add Columns  
dt.Columns.Add("Column1")  
dt.Columns.Add("Column2")  
dt.Columns.Add("Column3")  
'Add Rows  
dt.Rows.Add(New String() {"1", "2", "3"})  
dt.Rows.Add(New String() {"4", "5", "6"})  
dt.Rows.Add(New String() {"7", "8", "9"}) _

_'Bind it to our skinned control  
SkinnedControl.DataSource = dt  
SkinnedControl.DataBind()  
End If  
End Sub _

_#End Region_

Next, we need to call the databind function, for our purposes, we are going to override the RenderContents method, check for postback, if false, then we will databind, then finally tell our base class to RenderContents.  Once again, we put the Render code inside of a Render Region 

**_C#_**

_#region Render  
protected override void RenderContents(HtmlTextWriter writer)  
{  
// bind our data  
if (Page.IsPostBack == false)  
{  
DataBind();  
} _

_// this will apply our skinned asp:repeater here  
base.RenderContents(writer);  
}  
#endregion_

**_VB_**

_#Region " Render " _

_Protected Overrides Sub RenderContents(ByVal writer As System.Web.UI.HtmlTextWriter)  
'bind our data  
If Page.IsPostBack = False Then  
DataBind()  
End If _

_'this will apply our skinned asp:repeater here  
MyBase.RenderContents(writer) _

_End Sub _

_#End Region_

 

**Build and test**

At this point, we can take a break for all of you impatient people, and see if our control works or not, then we will get into customizing it a bit more.  First things first, hit the build button.  If you followed the code snippets correctly, you should have had a clean build, if not, go back and check what might have went wrong. C# may throw a warning that m_Title is set but never used, this will be rectified near the end of this article. 

Now, we need to create a new ascx for our skin, above we named it "Skin-SampleControl.ascx", so open up your CommunityServer/Themes/default/Skins folder and create a new text document aptly named "Skin-SampleControl.ascx".  If you would like, you can now open this file with your favorite editor, I prefer visual studio myself. 

Since we are doing a simple control without anything special on it, and obviously no code behind, we need to do no special directives at the top of the file.  Lets just add a simple asp:repeater control, name it "SkinnedControl" as per our code looks for this control, and then give it a very simplistic headertemplate that creates a table and column header row, an itemtemplate that just iterates our 3 sample columns, and a footertemplate to close the table.  Below is the ascx code, and how it should look for this sample. 

**_ASP.Net_**

_<!--   
This is a sample control, this specific control only has a repeater on it.  
//-->  
<asp:repeater id="SkinnedControl" runat="server">  
<HeaderTemplate>  
<table border="0" class="CommonSidebarContent">  
<tr>  
<td width="140" align="left" class="CommonBreadCrumbArea">Column1</td><td width="30" align="left" class="CommonBreadCrumbArea">Column2</td><td width="30" align="left" class="CommonBreadCrumbArea">Column3</td>  
</tr>   
</HeaderTemplate>  
<ItemTemplate>  
<tr>  
<td>  
<b><%# DataBinder.Eval(Container.DataItem, "Column1")%>  
</td>  
<td>  
<%# DataBinder.Eval(Container.DataItem, "Column2")%>  
</td>  
<td>  
<%# DataBinder.Eval(Container.DataItem, "Column3")%>  
</td>  
</tr>  
</ItemTemplate>  
<FooterTemplate>  
</table>  
</FooterTemplate>  
</asp:repeater>_

Save the control, and now lets add the control to our SkinnedControl.aspx we created earlier in the article. (CommunityServer/Sample/SkinnedControl.aspx).

**Add the control to the page**

First we want to register a tag prefix for the project, lets add it just above the Import Namespace line.

_<%@ Register TagPrefix="SC" Namespace="RenEvo.CommunityServer.Articles" Assembly="RenEvo.CommunityServer.Articles" %>_

Next, just below our first AdPart, lets add the control to the page.

_<SC:SampleControl runat="server" id="SampleControl1" />_

Pretty easy huh?  Now, the moment you have been waiting for, lets go ahead and visit the SkinnedControl.aspx!

If you have followed this article to the tee, you should have a little table with 3 columns and 4 rows the first one being the column headers.  If not, go back over the article and figure out where you went wrong.  No seriously, you didn't pay attention, start over.

Ok, now I am going to tell you what that m_Title was for, as well as some tricks to make this a bit more snazzy.

**Customizing the Control**

Lets create a property that will wrap our m_Title internal variable. I love me some Regions...

**_C#_**

_#region Properties  
public string Title  
{  
get { return m_Title; }  
set { m_Title = value; }  
}  
#endregion_

**_VB_**

_#Region " Properties " _

_Public Property Title() As String  
Get  
Return m_Title  
End Get  
Set(ByVal value As String)  
m_Title = value  
End Set  
End Property _

_#End Region_

Now, lets go ahead and do a generic way to output the "Title" of our control. In the RenderContents function, add the following code below the databind, but above the call to the base class to RenderContents. 

**_C#_**

_// pre writer stuff  
// optional title  
if (m_Title.Length > 0)  
{  
writer.WriteBeginTag("div");  
writer.WriteAttribute("class", "CommonMessageTitle");  
writer.Write(HtmlTextWriter.TagRightChar);  
writer.Write(m_Title);  
writer.WriteEndTag("div");  
}_

**_VB_**

_'pre writer stuff  
'optional title  
If Me.Title.Length > 0 Then  
writer.WriteBeginTag("div")  
writer.WriteAttribute("class", "CommonMessageTitle")  
writer.Write(HtmlTextWriter.TagRightChar)  
writer.Write(m_Title)  
writer.WriteEndTag("div")  
End If_

Now if we build and refresh our page, you will notice nothing changed, that is because we did not set our Title in the control definition.  Open up SkinnedControl.aspx and add the title="Sample Title" attribute to the SampleControl. When you save and refresh the page, you will notice now that you have a nifty title above the table!  Below is the code, incase you goofed up those easy instructions.

_<SC:SampleControl runat="server" id="SampleControl1" title="Sample Control" />_

That is pretty much it, but to clean it up, and show you a bit more of what is possible, I have added a bit more code to the RenderContents function that fully wraps the table with some custom styling.  Modified RenderContents below

**_C#_**

_#region Render  
protected override void RenderContents(HtmlTextWriter writer)  
{  
// bind our data  
if (Page.IsPostBack == false)  
{  
DataBind();  
} _

_// pre writer stuff  
// optional title  
if (m_Title.Length > 0)  
{  
writer.WriteBeginTag("div");  
writer.WriteAttribute("class", "CommonMessageTitle");  
writer.Write(HtmlTextWriter.TagRightChar);  
writer.Write(m_Title);  
writer.WriteEndTag("div");  
}  
// our wrapper for the content  
writer.WriteBeginTag("div");  
writer.WriteAttribute("class", "CommonMessageContent");  
writer.WriteAttribute("id", this.ID);  
writer.Write(HtmlTextWriter.TagRightChar); _

_// this will apply our skinned asp:repeater here  
base.RenderContents(writer); _

_// close our content wrapper  
writer.WriteEndTag("div");  
}  
#endregion_

**_VB_**

_#Region " Render " _

_Protected Overrides Sub RenderContents(ByVal writer As System.Web.UI.HtmlTextWriter)  
'bind our data  
If Page.IsPostBack = False Then  
DataBind()  
End If _

_'pre writer stuff  
'optional title  
If Me.Title.Length > 0 Then  
writer.WriteBeginTag("div")  
writer.WriteAttribute("class", "CommonMessageTitle")  
writer.Write(HtmlTextWriter.TagRightChar)  
writer.Write(m_Title)  
writer.WriteEndTag("div")  
End If _

_'our wrapper for the content  
writer.WriteBeginTag("div")  
writer.WriteAttribute("class", "CommonMessageContent")  
writer.WriteAttribute("id", Me.ID)  
writer.Write(HtmlTextWriter.TagRightChar) _

_'this will apply our skinned asp:repeater here  
MyBase.RenderContents(writer) _

_'close our content wrapper  
writer.WriteEndTag("div")  
End Sub _

_#End Region_

**End of article wrap up**

Throughout this article you have gotten a decent grasp on how the skin system works with CommunityServer, it is not nearly as overwhelming as some might imagine, and is actually kind of fun once you get the hang of it, something about binding a code blindly to an ascx really got me fired up that in my own product when the need arose to have a secondary view format for a specific web page, I took a similar approach with creating multiple ascx files that had the same code behind, and it worked out wonderfully.

Some key things that I would like to mention about this article is, the original article had rave reviews and fantastic feedback, which prompted me to want to write more, as it seems that I have something to offer to the CommunityServer community.  I have provided with this article all of the files neccessary to get the sample running on your site before even cracking the code which is available in C# and Visual Basic from the download link below.

Again, thanks to everyone for thier support, it is really encouraging to find that you spend time writing an article and it gets such great responses.

[Download Article Code][4]

[See the online demo!][5]

![][6]

[1]: http://renevo.com/files/folders/articles_cs/entry252.aspx "Community Server Sample Page"
[2]: http://renevo.com/files/folders/misc/entry156.aspx "Download VB Express"
[3]: http://renevo.com/files/folders/misc/entry157.aspx "C# Express"
[4]: http://renevo.com/files/folders/articles_cs/entry516.aspx "Get the code!"
[5]: http://renevo.com/Sample/SkinnedControl.aspx "View the online demo"
[6]: http://renevo.com/aggbug.aspx?PostID=517

