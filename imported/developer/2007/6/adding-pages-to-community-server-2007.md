---
title: Adding pages to Community Server 2007
published: true
date: 2007-06-21 17:34:00 +0000 UTC
tags: imported 
---
Adding Pages to Community Server 2007

Well, here I am again, and I have finally taken a bit of my time to setup a new CS 2007 website. Even though the steps are pretty similar, there are some changes to the new Community Server.  Below I will outline the steps required to add new pages with skinning support to your Community Server 2007 website.  
In the previous article on adding pages to CS 2.1, I would have had you copy files and cut out content, for this article we will be creating everything from scratch.

**Step 1**  
Start off by creating a new directory called "Sample" in your root web directory. (referred to web/sample/ from here on out).

Now create a new text file called "Default.aspx", open up this file with notepad, and add the following lines.

> > _<%@ Page %>  
<%--   
This is my placeholder, my actual file is located in web/themes/[THEME NAME]/sample/sample.aspx  
\--%>  

_

You should now be able to navigate to this page without any issues. (<http://www.yourdomain.com/test/>)

**Step 2 - Creating the themed page**  
_For this article, I will only be covering how to do this on the default theme, you can pretty much apply this to all other themes._

In your web/Themes/ directory, create a new directory called "sample".

Next, create a new text file and name it "sample.master", open up this file with notepad, and add the following lines.

> > _<%@ Master Language="C#" AutoEventWireup="true" MasterPageFile="../Common/master.Master" %>_

> > _<asp:Content ContentPlaceHolderID="HeaderRegion" runat="server" >  
    <CSControl:SelectedNavigation Selected="sample" runat="Server" />  
    <CSControl:ThemeStyle runat="server" href="~/style/forum.css" mce_href="~/style/forum.css" />  
    <CSControl:ThemeStyle runat="server" href="~/style/gallery.css" mce_href="~/style/gallery.css" />  
    <CSControl:ThemeStyle runat="server" href="~/style/forum_print.css" mce_href="~/style/forum_print.css" media="print" />  
    <CSControl:ThemeStyle runat="server" href="~/style/gallery_print.css" mce_href="~/style/gallery_print.css" media="print" />  
</asp:Content>_
>> 
>> _<asp:Content ContentPlaceHolderID="bcr" runat="server">  
    <asp:ContentPlaceHolder id="bcr" runat="server" />  
</asp:Content>_
>> 
>> _<asp:Content ContentPlaceHolderID="rcr" runat="server" >  
    <asp:ContentPlaceHolder ID="rcr" runat="server" />  
</asp:Content>_

Just to explain a bit, we have created a new master page that inherits the master.Master for the website, we have overridden the header region to select the "sample" page, and finally we have created two content regions for the page (CS 2007 style).

At this point, you can add any custom code to the your new pages "main" template, but for now, lets just leave it as is.

Now, lets create another text file, and call it "sample.aspx" and open it up in notepad.

The first thing we want to do in this file, is all of the standard asp.net page definitions, as well as import some common namespaces, useful for coding later on.  This also keeps our "ad support".

> > _<%@ Page EnableViewState="false" Language="C#" AutoEventWireup="true" Inherits="CommunityServer.Controls.CSThemePage" MasterPageFile="sample.Master" %>  
<%@ Import Namespace="CommunityServer.Components" %>  
<%@ Import Namespace="System.Collections.Generic" %>  
<%@ Register TagPrefix="CSUserControl" TagName="AdTop" src="../Common/Ad-Top.ascx" mce_src="../Common/Ad-Top.ascx" %>  
<%@ Register TagPrefix="CSUserControl" TagName="AdBottom" src="../Common/Ad-Bottom.ascx" mce_src="../Common/Ad-Bottom.ascx" %>_

Next, we will set the pages title, this is something that I can't remember seeing in CS 2.1, its quite useful for creating "sub" sites within your main site.  For now we will set the title to the default site name.

> > _<script language="C#" runat="server">_
>> 
>> _    void Page_Load()  
    {  
        SetTitle(CurrentCSContext.SiteSettings.SiteName, false);  
    }_
>> 
>> _</script>_

Now we will move on to adding our bcr area or, Body Content Area.  This is much cleaner then the previous way of adding content parts, as you now don't have to be a html code guru and can simply enter in regular html code.

> > _<asp:Content ID="Content1" ContentPlaceHolderID="bcr" runat="server">  
    <div class="CommonContentArea">  
 <CSControl:AdPart runat = "Server" contentname="StandardTop" ContentCssClass="CommonContentPartBorderOff" ContentHoverCssClass="CommonContentPartBorderOn">  
     <DefaultContentTemplate>  
         <CSUserControl:AdTop runat="server" />  
     </DefaultContentTemplate>  
 </CSControl:AdPart>_
>> 
>> _ <div class="CommonContent">  
            <CSControl:ContentPart ContentName="sample" runat="server" ContentCssClass="CommonContentPartBorderOff" ContentHoverCssClass="CommonContentPartBorderOn">  
                <DefaultContentTemplate>  
                    <h2 class="CommonTitle">Sample Page</h2>  
                    <div class="CommonContent">  
                        <div style="line-height: 140%;">  
           You have created a sample page!  
                        </div>  
                    </div>  
                </DefaultContentTemplate>  
            </CSControl:ContentPart>  
          
 </div>  
   
 <CSControl:AdPart runat="Server" ContentName="StandardBottom" ContentCssClass="CommonContentPartBorderOff" ContentHoverCssClass="CommonContentPartBorderOn">  
     <DefaultContentTemplate>  
         <CSUserControl:AdBottom runat="server" />  
     </DefaultContentTemplate>  
 </CSControl:AdPart>  
    </div>  
</asp:Content>_

Finally, lets add our rcr area or, Right Content Area.

> > _<asp:Content ContentPlaceHolderID="rcr" runat="server">  
    <div class="CommonSidebar">  
 <CSControl:ContentPart ContentName="sampleSidebar1" runat="server" ContentCssClass="CommonContentPartBorderOff" ContentHoverCssClass="CommonContentPartBorderOn">  
            <DefaultContentTemplate>  
                <div class="CommonSidebarArea">  
                 <div class="CommonSidebarRoundTop"><div class="r1"></div><div class="r2"></div><div class="r3"></div><div class="r4"></div></div>  
                 <div class="CommonSidebarInnerArea">  
                  <h4 class="CommonSidebarHeader">Sidebar 1</h4>  
                  <div class="CommonSidebarContent">  
                            <p>  
                                Sign-in with your Admin account and double-click to edit me!  
                            </p>  
                        </div>  
                 </div>  
                 <div class="CommonSidebarRoundBottom"><div class="r1"></div><div class="r2"></div><div class="r3"></div><div class="r4"></div></div>  
                </div>  
            </DefaultContentTemplate>  
        </CSControl:ContentPart>  
    </div>  
</asp:Content>_

Now save the file, and go ahead and close it.

**Adding the page to the SiteUrls**  
Unlike CS 2.1, you must complete this next step for a page to even load in the web application.

Open up the "SiteUrls.config" file with notepad.

Scroll down to around line 562. You will see a comment that says "When adding custom locations, add above here."  We are going to do just that.

Add the following just above the comment.

> > _  <location name="sample" path="/sample/" themeDir="sample">  
    <url name="samplehome" path="" pattern="default.aspx" vanity="{2}" physicalPath="##themeDir##" page="sample.aspx" />  
  </location>_

Now, before you go jumping off to your newly created page, we have a few more things to do, or else you are going to get some errors.

**Adding the link to the top menu**  
With the SiteUrls.config file still open, scroll down the bottom to the "navigation" section.  Just under the "home" link, add the following.

> > _<link name="sample" resourceUrl="samplehome" resourceName="sample" roles="Everyone" />_

Save the file, and go ahead and close the SiteUrls.config.

As you noticed in my previous articles, until you add the resource to the resources.xml, your link will not show up on the main link menu.

Navigate to "web/languages/[YOUR LANGUAGE]/" and open up the file "Resources.xml" with notepad.

Around line 44, you will see the text "<!-- Main Navigation -->"  just below the "Home" resource, add the following code.

> > _<resource name="sample">Sample</resource>_

Save the file.

**Restarting the web application**  
At this point we need to just "touch" the web.config file, this is done by opening the file with notepad, saving, then closing the file.

**Navigate to your newly created page**  
Finally, you should be able to goto your home page, click on the sample link, and see your new page!

If you don't see it, start back at the beginning and find out what you did wrong.  Or post up in my forums for help and I will try to help you out.

**Conclusion**  
This concludes this article, all in all I think that adding and modifying pages in CS 2007 is much easier, and allows much more flexibility then the previous version.  So go out there, be creative, and see what you can do!

**Download Links**

You can download all of the modifications here: [_CommunityServer 2007 Sample Page][1]_

You can see this article in action at: [http://www.swglabs.com][2]

 

![][3]

[1]: http://www.renevo.com/files/folders/articles_cs/entry1129.aspx
[2]: http://www.swglabs.com/
[3]: http://renevo.com/aggbug.aspx?PostID=1128

