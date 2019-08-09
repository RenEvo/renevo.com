---
title: Adding Pages To Community Server
published: true
date: 2006-11-30 21:35:08 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2006/11/30/adding-pages-to-community-server.aspx
file: adding-pages-to-community-server.aspx
path: /blogs/developer/archive/2006/11/30/
author: tom anderson
words: 1270
---
As some people have seen on this site, I have added an IRC page.  In this blog I will briefly go over how I created the new page, and then implemented it in the site to show up on the tab list.

**First Step**

The first thing that I did was create a new directory and copy the Default.aspx from the root Community Server folder into the newly created folder.  For this example, we will call the folder "Sample".  Next, navigate to the new page <http://www.renevo.com/Sample/>

Now that we have our new page, we need to add it to the menu, and get rid of the default content that is on the page.

**Removing The "Home Page" Content**

Open up the new Default.aspx and edit the Page_Load function at the top to remove some "Home Page" specific items.  Remove everything below the UsersOnline.SetLocation("Home"); Modify that string in the function call to "Sample Page".

Next, scroll down and remove the items inside the <div class="CommonContent">  (you can leave the adpart if you like), next remove all of the components inside of the <div class="CommonSideBar"> section.

Finally it is also a good idea to remove the tagprefix's and imports at the top of the page that you won't be using, specifically the Blogs.Components, Disucssions.Components, and Galleries.Components imports and the Galleries, Blog, and CSD TagPrefix's.

Now if you navigate to your page, you will have a blank page to work with.  At the end of this blog is a link to download the default.aspx used in this article, so you don't have to worry about making your own.  This is just an explanation on how I implemented the system.

**Adding Your Own Content Locations**

Inside of the first <div class="CommonContent"> element, lets add a ContentPart, which is editable by double clicking on it to bring up the modal editor similar to the home pages welcome box you get out of the box.

<CS:ContentPart runat="server" contentname="samplepage" id="samplepageContentPart" text=""&lt;p&gt;&lt;/div&gt;Sample Editor&lt;br /&gt;&lt;/p&gt;" />

_Woh woh woh... hold on a tick... wtf did you just do?_

What I did was create a new asp.net control, specified that it runs on the server, the "contentname" is used for storage in the database, the "id" is the name of the control, and the "text" is the default text if you haven't previously configured any text for it (via the double click edit mentioned previously).

Alright, lets save it, and navigate back to that page.  Double click the control, and add the following text. "My First Community Server Page".

**Adding Your Own Sidebar**

This step is totally optional, but if you want to create a sidebar for the page to have the same look and feel as the rest of the website, lets add a new sidebar editor.

Inside the <div class="CommonSidebar"> towards the bottom of the page, insert the following code.

<CS:ContentPart runat="Server" contentname="samplesidebar" id="samplesidebarContentPart" text="&lt;br /&gt;&lt;h4 class=CommonSidebarHeader&gt;Title&lt;/h4&gt;&lt;div class=CommonSidebarContent&gt;Content&lt;/div&gt;" />

As you can see, this is the same type of control that we used for the top section, we simply used the html style of the sidebar for the default.  After you navigate to the newly saved page, you will see a sidebar with the header of "Title" and the inside of that content box will be "Content".  You can double click on this control to edit its text, just be sure to watch your delete key, as long as you don't delete the line between the bold text and the "Content" you will retain your style, if you do edit it, I am sure you are clever enough to figure out how to edit it in html mode.

**Review**

Ok, so up to this point we have created a new directory and put our new default.aspx in it, modified the page to look like we want, and added some custom content editor areas to display some text.  Whats next?

**_Integrating Into The Web Site_**

**SiteUrls.config**

Here is the fun part, lets open up SiteUrls.config in notepad. _Crap that file is huge and confusing..._ Yes, I agree, it is huge, and it does look confusing at first, but lets just focus on the important parts.  Be wary, you SHOULD create a backup before doing anything in this file, you mess up here, your site goes boom.  Ok, warning of screwing up everything is over, lets move along.

At the top of the file, you will find a section call "locations", lets create a new location in that sub section like so:

<location name="sample" path="/sample/" exclude="true" />

This simply creates a new location for the website, specifies the directory, and that it should be excluded from Url Rewriting Changes. (i.e. static)

Don't save just yet, you will mess up stuff, people will complain, and your not done yet lazy.

Next, scroll down to the <urls> tag, and lets add a new url.

<url name="samplehome" location="sample" path="default.aspx" />

Yes, that was easy, and all of the tags do exactly what they specify, "name" is the internal name of the url, "location" is the location name that we created in the previous step, and "path" is the file to navigate to in that location.

_So is that it?_ Not quite, we need to add it to the actual tabs, so hands off that CTRL S key combo...

Find the <navigation> tag, and lets add a new link.

<link name="sample" resourceUrl="samplehome" resourceName="sample" roles="Everyone" />

That wasn't so hard now was it?  the "name" once again is the internal name of the link, "resourceUrl" is that url that we just setup in the previous step, resourceName is something I will cover in the next step (good to just match it to the link name), and the "roles" are the allowed roles for this page, seperated by comma's.

Ok, save the file, and lets move on.

_Dude, it didn't work, nothing changed on my page tabs...._  I didn't say we where done, lets move on...

**Language Configuration**

You didn't think that we where going to break the multi-language support did you?  Shame on you...

Lets navigate to <install>/Languages/en-US/ and open the Resources.xml in notepad.

This file pretty much has all of your english strings for the Community Server in it, there are more, but this is the bulk of them.

Lets do a CTRL F and look for "<!-- Main Navigation -->" which shouldn't be that far down, around line 40 or so.

Just above the <!-- End MainNavigation --> lets add the following line.

<resource name="sample">Sample Page</resource>

**Touch Web.config**

For those of you that don't know, you will have to "touch" web.config to restart the web application, as resources.xml seems to be read on application_start.  Simply open up web.config in notepad, press CTRL S, and close the file.  This will force a restart of the CommunityServer web application.

If you followed my instructions, you will now have a link to "Sample Page" on your portal.

_Hey its there! and it works... but, it doesn't seem to be selected when I navigate to it, it still looks like I am on the homepage._

Thats why I have the next step!

**Creating Your Own Master**

Lets find a new directory, and head into <install>/Themes/default/Masters/.  Copy the HomeMaster.ascx and name it SampleMaster.ascx.

Lets open up SampleMaster.ascx in notepad, and modify the following line:

<CS:SelectedNavigation Selected = "home" runat="Server" id="Selectednavigation1" />

TO

<CS:SelectedNavigation Selected = "sample" runat="Server" id="Selectednavigation1" />

Save and close the file.

_That didn't do anything_

Thats because we didn't specify our page to use this master yet, lets re-open our <install>/Sample/Default.aspx in notepad, and modify it to use our new master page. Modify the following line:

<CS:MPContainer runat="server" id="Mpcontainer1" ThemeMasterFile = "HomeMaster.ascx" >

TO

<CS:MPContainer runat="server" id="Mpcontainer1" ThemeMasterFile = "SampleMaster.ascx" >

And now, navigate to your page, and see some new page goodness.

**Review**

Ok, we copied a file, modified it, created some new edit regions, added the link, then bound it to a new master theme.  That wasn't too bad was it?

You can download all of the modifications [Here][1]

Hope this article was helpful, if you have any questions, feel free to ask in the forums!



[1]: http://renevo.com/files/folders/articles_cs/entry252.aspx "Download Sample Code"


