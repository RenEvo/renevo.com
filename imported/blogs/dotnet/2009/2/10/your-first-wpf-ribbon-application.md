---
title: Your first WPF Ribbon Application
published: true
date: 2009-02-10 19:11:02 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2009/02/10/your-first-wpf-ribbon-application.aspx
file: your-first-wpf-ribbon-application.aspx
path: /blogs/dotnet/archive/2009/02/10/
author: tom anderson
words: 11586
---
Well, in the spirit of those "great" first programs, I thought that I would post this one up as a [follow up to my last news post about how I found this little gem in the WPF Futures][1] site.

So, without further ado, lets get started.

First things first, you need to head over to [codeplex][2] and [download the WPF Ribbon Preview][3], this will require you to fill out an Office 2007 licensing agreement, it is free, and basically says that you aren't going to compete with Microsoft if you use this UI, and that you will adhere to the Ribbon Standards, which the control does a really good job of enforcing for you.

In the download, the file you are going to be most worried about is the "RibbonControlsLibrary.dll", this is the assembly that contains all the Ribbon goodness. Extract that to a location where you can easily find it, and startup Visual Studio 2008.

We want to create a new WPF Application, the language you choose at this point is totally irrelevant, as most of this article is going to be in XAML. I named mine "RibbonSample".

![image][4]

First things first, add a reference to the RibbonControlsLibrary.dll. You will need to browse to the file, the place you extracted it above. Also be sure that the "Copy Local" property for the reference is set to True.

Now lets dig into some XAML, I prefer in Visual Studio to just remove the preview pane all together when starting my XAML work.

We are going to have to add a few new schema references to the XAML, specifically for the Ribbon Control.  Lets prefix this with an "r" to keep our XAML sane. If you have never added any XAML references, it is done by simply declaring a new CLR schema reference. The code below has the added code italic and bolded.
    
    
    <Window x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    **_    _****_xmlns:r_****_="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"_**
        Title="Window1" Height="300" Width="300">
        <Grid>
            
        </Grid>
    </Window>

Now some major changes, we are going to change the main window class to a ribbon window, modify the size, startup location, resize mode, min height, and min width. For easier layout, I am going to also change the Grid to a DockPanel. Inside the DockPanel, we also need to add an actual Ribbon, or else our app will just show up as a big black mass. This is done by simply placing a Ribbon control inside of the DockPanel. And finally, to fill up the area that we don't need with the ribbon, lets just dump in a RichTextBox control.

After doing all of the above, the XAML to get a Ribbon on a Form up and running is only this amount of XAML.
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto"></RichTextBox>
        </DockPanel>
    </r:RibbonWindow>

![image][5]

Pretty simple eh?

It doesn't really do much at this point, as we haven't added anything to it, but that will come in the next step.

**Adding some images   
**Ribbon controls are very graphical, so in order to fill this thing up, lets add some images to our projects, create an images directory and find some nice png files, or use the ones [provided in the download below][6]. I used 48x48 PNG files with transparency.

Starting at the top, we are going to want to start placing some images on our form. Just like regular forms, you can add the form icon in the root window declaration.

Next we want to add a button to the QAT (Quick Access Toolbar) that is located on the form's title bar. This is done by simply adding to the QuickAccessToolBar element for the Ribbon. For this one, we will simply add a command button that has an icon and is clickable. This part is a bit new to me, but we need to create some static resources in the form (or you can link them) for the Ribbon Commands, this makes them very re-usable, and easier to work with.  Below is the XAML for the new QAT button and its placement. You will however have to implement the "CanExecute" event in order for the button to be clickable, simply return True for this function.
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
        
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto"></RichTextBox>
        </DockPanel>
    </r:RibbonWindow>

![image][7]

Now that we have implemented a QAT button, lets add an image to our "Start Button" for the Ribbon, this is pretty straightforward. For this image I used a 24x24 as it doesn't streatch into the area.
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
                <r:Ribbon.ApplicationMenu>
                    <r:RibbonApplicationMenu>
                        <r:RibbonApplicationMenu.Command>
                            <r:RibbonCommand SmallImageSource="Imagesbox.png" LargeImageSource="Imagesbox.png" />
                        </r:RibbonApplicationMenu.Command>
                    </r:RibbonApplicationMenu>
                </r:Ribbon.ApplicationMenu>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto"></RichTextBox>
        </DockPanel>
    </r:RibbonWindow>
    

The next step will be to add a few menu item commands, and the first one we will add a few sub-items to so you can see how the rollout menus work. As part of a work around, we will need to also add a sized rectangle to the RecentItemList of the application menu so that it will draw large enough for us to have sub items without scrolling (similar to how office works).
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 1" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 1" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesfiles.png" LargeImageSource="Imagesfiles.png" />
                <r:RibbonCommand x:Key="MenuItem2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 2" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 2" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 3" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 3" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesprint.png" LargeImageSource="Imagesprint.png" />
                <r:RibbonCommand x:Key="MenuItem4" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 4" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 4" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesdiagnostic.png" LargeImageSource="Imagesdiagnostic.png" />
                
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
                <r:Ribbon.ApplicationMenu>
                    <r:RibbonApplicationMenu>
                        <r:RibbonApplicationMenu.Command>
                            <r:RibbonCommand SmallImageSource="Imagesbox.png" LargeImageSource="Imagesbox.png" />
                        </r:RibbonApplicationMenu.Command>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem1}">
                            <TextBlock Text="Item 1 in the list" />
                            <TextBlock Text="Item 2 in the list" />
                            <TextBlock Text="Item 3 in the list" />
                            <TextBlock Text="Item 4 in the list" />
                        </r:RibbonApplicationMenuItem>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem2}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem3}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem4}" />
                        <r:RibbonApplicationMenu.RecentItemList>
                            <Rectangle Height="300" />
                        </r:RibbonApplicationMenu.RecentItemList>
                    </r:RibbonApplicationMenu>
                </r:Ribbon.ApplicationMenu>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto"></RichTextBox>
        </DockPanel>
    </r:RibbonWindow>

![image][8]

Now we can move along to the fun parts, and add a few ribbon tabs, and a few groups of buttons on those tabs.
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 1" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 1" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesfiles.png" LargeImageSource="Imagesfiles.png" />
                <r:RibbonCommand x:Key="MenuItem2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 2" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 2" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 3" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 3" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesprint.png" LargeImageSource="Imagesprint.png" />
                <r:RibbonCommand x:Key="MenuItem4" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 4" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 4" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesdiagnostic.png" LargeImageSource="Imagesdiagnostic.png" />
                
                <r:RibbonCommand x:Key="HomeButton1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calculator" LabelDescription="Calc This!" ToolTipTitle="Calculator" ToolTipDescription="Used to do math and stuff" SmallImageSource="Imagescalculator.png" LargeImageSource="Imagescalculator.png" />
                <r:RibbonCommand x:Key="HomeButton2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calendar" LabelDescription="Schedule This!" ToolTipTitle="Calendar" ToolTipDescription="Schedule and remind yourself of stuff" SmallImageSource="Imagescalendar.png" LargeImageSource="Imagescalendar.png" />
                <r:RibbonCommand x:Key="HomeButton3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Computer" LabelDescription="Format This!" ToolTipTitle="Computer" ToolTipDescription="Where you store your naked pictures" SmallImageSource="Imagescomputer.png" LargeImageSource="Imagescomputer.png" />
                
                <r:RibbonCommand x:Key="MediaEject" CanExecute="RibbonCommand_CanExecute" LabelTitle="Eject" LabelDescription="Eject" ToolTipTitle="Eject" ToolTipDescription="Open the cup holder" SmallImageSource="Imagesbt_eject.png" LargeImageSource="Imagesbt_eject.png" />
                <r:RibbonCommand x:Key="MediaBackward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Previous" LabelDescription="Previous" ToolTipTitle="Previous" ToolTipDescription="Previous Tune" SmallImageSource="Imagesbt_skip_backward.png" LargeImageSource="Imagesbt_skip_backward.png" />
                <r:RibbonCommand x:Key="MediaPlay" CanExecute="RibbonCommand_CanExecute" LabelTitle="Play" LabelDescription="Play" ToolTipTitle="Play" ToolTipDescription="Play Tune" SmallImageSource="Imagesbt_play.png" LargeImageSource="Imagesbt_play.png" />
                <r:RibbonCommand x:Key="MediaStop" CanExecute="RibbonCommand_CanExecute" LabelTitle="Stop" LabelDescription="Stop" ToolTipTitle="Stop" ToolTipDescription="Stop the music" SmallImageSource="Imagesbt_stop.png" LargeImageSource="Imagesbt_stop.png" />
                <r:RibbonCommand x:Key="MediaForward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Next" LabelDescription="Next" ToolTipTitle="Next" ToolTipDescription="Next Tune" SmallImageSource="Imagesbt_skip_forward.png" LargeImageSource="Imagesbt_skip_forward.png" />
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
                <r:Ribbon.ApplicationMenu>
                    <r:RibbonApplicationMenu>
                        <r:RibbonApplicationMenu.Command>
                            <r:RibbonCommand SmallImageSource="Imagesbox.png" LargeImageSource="Imagesbox.png" />
                        </r:RibbonApplicationMenu.Command>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem1}">
                            <TextBlock Text="Item 1 in the list" />
                            <TextBlock Text="Item 2 in the list" />
                            <TextBlock Text="Item 3 in the list" />
                            <TextBlock Text="Item 4 in the list" />
                        </r:RibbonApplicationMenuItem>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem2}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem3}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem4}" />
                        <r:RibbonApplicationMenu.RecentItemList>
                            <Rectangle Height="300" />
                        </r:RibbonApplicationMenu.RecentItemList>
                    </r:RibbonApplicationMenu>
                </r:Ribbon.ApplicationMenu>
                <r:RibbonTab Label="Home">
                    <r:RibbonGroup>
                        <r:RibbonButton Command="{StaticResource HomeButton1}" />
                        <r:RibbonButton Command="{StaticResource HomeButton2}" />
                        <r:RibbonButton Command="{StaticResource HomeButton3}" />
                    </r:RibbonGroup>
                </r:RibbonTab>
                <r:RibbonTab Label="Media">
                    <r:RibbonGroup>
                        <r:RibbonButton Command="{StaticResource MediaEject}" />
                        <r:RibbonButton Command="{StaticResource MediaBackward}" />
                        <r:RibbonButton Command="{StaticResource MediaPlay}" />
                        <r:RibbonButton Command="{StaticResource MediaStop}" />
                        <r:RibbonButton Command="{StaticResource MediaForward}" />
                    </r:RibbonGroup>
                </r:RibbonTab>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto">
                <FlowDocument>
                    <Paragraph>
                        <Hyperlink NavigateUri="http://www.renevo.com">RenEvo Software &amp; Designs</Hyperlink>
                    </Paragraph>
                </FlowDocument>
            </RichTextBox>
        </DockPanel>
    </r:RibbonWindow>

And the Preview:

![image][9] ![image][10]

Now, lets work on those groups a bit, lets add a title to both groups, and make the icons for the media in a group.
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 1" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 1" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesfiles.png" LargeImageSource="Imagesfiles.png" />
                <r:RibbonCommand x:Key="MenuItem2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 2" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 2" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 3" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 3" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesprint.png" LargeImageSource="Imagesprint.png" />
                <r:RibbonCommand x:Key="MenuItem4" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 4" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 4" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesdiagnostic.png" LargeImageSource="Imagesdiagnostic.png" />
    
                <r:RibbonCommand x:Key="HomeButton1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calculator" LabelDescription="Calc This!" ToolTipTitle="Calculator" ToolTipDescription="Used to do math and stuff" SmallImageSource="Imagescalculator.png" LargeImageSource="Imagescalculator.png" />
                <r:RibbonCommand x:Key="HomeButton2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calendar" LabelDescription="Schedule This!" ToolTipTitle="Calendar" ToolTipDescription="Schedule and remind yourself of stuff" SmallImageSource="Imagescalendar.png" LargeImageSource="Imagescalendar.png" />
                <r:RibbonCommand x:Key="HomeButton3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Computer" LabelDescription="Format This!" ToolTipTitle="Computer" ToolTipDescription="Where you store your naked pictures" SmallImageSource="Imagescomputer.png" LargeImageSource="Imagescomputer.png" />
    
                <r:RibbonCommand x:Key="MediaEject" CanExecute="RibbonCommand_CanExecute" LabelTitle="Eject" LabelDescription="Eject" ToolTipTitle="Eject" ToolTipDescription="Open the cup holder" SmallImageSource="Imagesbt_eject.png" LargeImageSource="Imagesbt_eject.png" />
                <r:RibbonCommand x:Key="MediaBackward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Previous" LabelDescription="Previous" ToolTipTitle="Previous" ToolTipDescription="Previous Tune" SmallImageSource="Imagesbt_skip_backward.png" LargeImageSource="Imagesbt_skip_backward.png" />
                <r:RibbonCommand x:Key="MediaPlay" CanExecute="RibbonCommand_CanExecute" LabelTitle="Play" LabelDescription="Play" ToolTipTitle="Play" ToolTipDescription="Play Tune" SmallImageSource="Imagesbt_play.png" LargeImageSource="Imagesbt_play.png" />
                <r:RibbonCommand x:Key="MediaStop" CanExecute="RibbonCommand_CanExecute" LabelTitle="Stop" LabelDescription="Stop" ToolTipTitle="Stop" ToolTipDescription="Stop the music" SmallImageSource="Imagesbt_stop.png" LargeImageSource="Imagesbt_stop.png" />
                <r:RibbonCommand x:Key="MediaForward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Next" LabelDescription="Next" ToolTipTitle="Next" ToolTipDescription="Next Tune" SmallImageSource="Imagesbt_skip_forward.png" LargeImageSource="Imagesbt_skip_forward.png" />
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
                <r:Ribbon.ApplicationMenu>
                    <r:RibbonApplicationMenu>
                        <r:RibbonApplicationMenu.Command>
                            <r:RibbonCommand SmallImageSource="Imagesbox.png" LargeImageSource="Imagesbox.png" />
                        </r:RibbonApplicationMenu.Command>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem1}">
                            <TextBlock Text="Item 1 in the list" />
                            <TextBlock Text="Item 2 in the list" />
                            <TextBlock Text="Item 3 in the list" />
                            <TextBlock Text="Item 4 in the list" />
                        </r:RibbonApplicationMenuItem>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem2}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem3}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem4}" />
                        <r:RibbonApplicationMenu.RecentItemList>
                            <Rectangle Height="300" />
                        </r:RibbonApplicationMenu.RecentItemList>
                    </r:RibbonApplicationMenu>
                </r:Ribbon.ApplicationMenu>
                <r:RibbonTab Label="Home">
                    <r:RibbonGroup>
                        <r:RibbonGroup.Command>
                            <r:RibbonCommand LabelTitle="Programs" />
                        </r:RibbonGroup.Command>
                        <r:RibbonGroup.GroupSizeDefinitions>
                            <r:RibbonGroupSizeDefinitionCollection>
                                <r:RibbonGroupSizeDefinition>
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                </r:RibbonGroupSizeDefinition>
                            </r:RibbonGroupSizeDefinitionCollection>
                        </r:RibbonGroup.GroupSizeDefinitions>
    
                        <r:RibbonButton Command="{StaticResource HomeButton1}" />
                        <r:RibbonButton Command="{StaticResource HomeButton2}" />
                        <r:RibbonButton Command="{StaticResource HomeButton3}" />
                    </r:RibbonGroup>
                </r:RibbonTab>
                <r:RibbonTab Label="Media">
                    <r:RibbonGroup>
                        <r:RibbonGroup.Command>
                            <r:RibbonCommand LabelTitle="Media Controls" />
                        </r:RibbonGroup.Command>
                        <r:RibbonControlGroup>
                            <r:RibbonButton Command="{StaticResource MediaEject}" />
                            <r:RibbonButton Command="{StaticResource MediaBackward}" />
                            <r:RibbonButton Command="{StaticResource MediaPlay}" />
                            <r:RibbonButton Command="{StaticResource MediaStop}" />
                            <r:RibbonButton Command="{StaticResource MediaForward}" />
                        </r:RibbonControlGroup>
                    </r:RibbonGroup>
                </r:RibbonTab>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto">
                <FlowDocument>
                    <Paragraph>
                        <Hyperlink NavigateUri="http://www.renevo.com">RenEvo Software &amp; Designs</Hyperlink>
                    </Paragraph>
                </FlowDocument>
            </RichTextBox>
        </DockPanel>
    </r:RibbonWindow>
    

![image][11] ![image][12]

And now lets add a single extra item to the QAT Drop down.

![image][13]
    
    
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 1" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 1" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesfiles.png" LargeImageSource="Imagesfiles.png" />
                <r:RibbonCommand x:Key="MenuItem2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 2" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 2" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 3" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 3" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesprint.png" LargeImageSource="Imagesprint.png" />
                <r:RibbonCommand x:Key="MenuItem4" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 4" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 4" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesdiagnostic.png" LargeImageSource="Imagesdiagnostic.png" />
    
                <r:RibbonCommand x:Key="HomeButton1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calculator" LabelDescription="Calc This!" ToolTipTitle="Calculator" ToolTipDescription="Used to do math and stuff" SmallImageSource="Imagescalculator.png" LargeImageSource="Imagescalculator.png" />
                <r:RibbonCommand x:Key="HomeButton2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calendar" LabelDescription="Schedule This!" ToolTipTitle="Calendar" ToolTipDescription="Schedule and remind yourself of stuff" SmallImageSource="Imagescalendar.png" LargeImageSource="Imagescalendar.png" />
                <r:RibbonCommand x:Key="HomeButton3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Computer" LabelDescription="Format This!" ToolTipTitle="Computer" ToolTipDescription="Where you store your naked pictures" SmallImageSource="Imagescomputer.png" LargeImageSource="Imagescomputer.png" />
    
                <r:RibbonCommand x:Key="MediaEject" CanExecute="RibbonCommand_CanExecute" LabelTitle="Eject" LabelDescription="Eject" ToolTipTitle="Eject" ToolTipDescription="Open the cup holder" SmallImageSource="Imagesbt_eject.png" LargeImageSource="Imagesbt_eject.png" />
                <r:RibbonCommand x:Key="MediaBackward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Previous" LabelDescription="Previous" ToolTipTitle="Previous" ToolTipDescription="Previous Tune" SmallImageSource="Imagesbt_skip_backward.png" LargeImageSource="Imagesbt_skip_backward.png" />
                <r:RibbonCommand x:Key="MediaPlay" CanExecute="RibbonCommand_CanExecute" LabelTitle="Play" LabelDescription="Play" ToolTipTitle="Play" ToolTipDescription="Play Tune" SmallImageSource="Imagesbt_play.png" LargeImageSource="Imagesbt_play.png" />
                <r:RibbonCommand x:Key="MediaStop" CanExecute="RibbonCommand_CanExecute" LabelTitle="Stop" LabelDescription="Stop" ToolTipTitle="Stop" ToolTipDescription="Stop the music" SmallImageSource="Imagesbt_stop.png" LargeImageSource="Imagesbt_stop.png" />
                <r:RibbonCommand x:Key="MediaForward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Next" LabelDescription="Next" ToolTipTitle="Next" ToolTipDescription="Next Tune" SmallImageSource="Imagesbt_skip_forward.png" LargeImageSource="Imagesbt_skip_forward.png" />
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                        <r:RibbonButton Command="{StaticResource MediaEject}" r:RibbonQuickAccessToolBar.Placement="InCustomizeMenu" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
                <r:Ribbon.ApplicationMenu>
                    <r:RibbonApplicationMenu>
                        <r:RibbonApplicationMenu.Command>
                            <r:RibbonCommand SmallImageSource="Imagesbox.png" LargeImageSource="Imagesbox.png" />
                        </r:RibbonApplicationMenu.Command>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem1}">
                            <TextBlock Text="Item 1 in the list" />
                            <TextBlock Text="Item 2 in the list" />
                            <TextBlock Text="Item 3 in the list" />
                            <TextBlock Text="Item 4 in the list" />
                        </r:RibbonApplicationMenuItem>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem2}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem3}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem4}" />
                        <r:RibbonApplicationMenu.RecentItemList>
                            <Rectangle Height="300" />
                        </r:RibbonApplicationMenu.RecentItemList>
                    </r:RibbonApplicationMenu>
                </r:Ribbon.ApplicationMenu>
                <r:RibbonTab Label="Home">
                    <r:RibbonGroup>
                        <r:RibbonGroup.Command>
                            <r:RibbonCommand LabelTitle="Programs" />
                        </r:RibbonGroup.Command>
                        <r:RibbonGroup.GroupSizeDefinitions>
                            <r:RibbonGroupSizeDefinitionCollection>
                                <r:RibbonGroupSizeDefinition>
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                </r:RibbonGroupSizeDefinition>
                            </r:RibbonGroupSizeDefinitionCollection>
                        </r:RibbonGroup.GroupSizeDefinitions>
    
                        <r:RibbonButton Command="{StaticResource HomeButton1}" />
                        <r:RibbonButton Command="{StaticResource HomeButton2}" />
                        <r:RibbonButton Command="{StaticResource HomeButton3}" />
                    </r:RibbonGroup>
                </r:RibbonTab>
                <r:RibbonTab Label="Media">
                    <r:RibbonGroup>
                        <r:RibbonGroup.Command>
                            <r:RibbonCommand LabelTitle="Media Controls" />
                        </r:RibbonGroup.Command>
                        <r:RibbonControlGroup>
                            <r:RibbonButton Command="{StaticResource MediaEject}" />
                            <r:RibbonButton Command="{StaticResource MediaBackward}" />
                            <r:RibbonButton Command="{StaticResource MediaPlay}" />
                            <r:RibbonButton Command="{StaticResource MediaStop}" />
                            <r:RibbonButton Command="{StaticResource MediaForward}" />
                        </r:RibbonControlGroup>
                    </r:RibbonGroup>
                </r:RibbonTab>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto">
                <FlowDocument>
                    <Paragraph>
                        <Hyperlink NavigateUri="http://www.renevo.com">RenEvo Software &amp; Designs</Hyperlink>
                    </Paragraph>
                </FlowDocument>
            </RichTextBox>
        </DockPanel>
    </r:RibbonWindow>
    

Finally, if you want to switch to the Office 2007 look & feel, instead of the Windows 7 look & feel, simply change the Resource Dictionary (you can also create customized ones). you will want to remove the form's icon if you use the office 2007 theme though.
    
    
    <!-- Remove the Icon Property if you are going to use the Office 2007 themes-->
    <r:RibbonWindow x:Class="Window1"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        xmlns:r="clr-namespace:Microsoft.Windows.Controls.Ribbon;assembly=RibbonControlsLibrary"
        Title="My First Ribbon Form" ResizeMode="CanResizeWithGrip" WindowStartupLocation="CenterScreen"
        Icon="Imagesapp.png"
        Height="600" Width="800" MinHeight="300" MinWidth="400">
    
        <r:RibbonWindow.Resources>
            <ResourceDictionary>
                <r:RibbonCommand x:Key="QATButton" CanExecute="RibbonCommand_CanExecute" LabelTitle="QAT Button" LabelDescription="This is a sample QAT Button" ToolTipTitle="QAT Button" ToolTipDescription="This is a sample QAT Button, it doesn't do anything" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 1" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 1" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesfiles.png" LargeImageSource="Imagesfiles.png" />
                <r:RibbonCommand x:Key="MenuItem2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 2" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 2" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagessave.png" LargeImageSource="Imagessave.png" />
                <r:RibbonCommand x:Key="MenuItem3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 3" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 3" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesprint.png" LargeImageSource="Imagesprint.png" />
                <r:RibbonCommand x:Key="MenuItem4" CanExecute="RibbonCommand_CanExecute" LabelTitle="Menu Item 4" LabelDescription="This is a sample menu item" ToolTipTitle="Menu Item 4" ToolTipDescription="This is a sample menu item" SmallImageSource="Imagesdiagnostic.png" LargeImageSource="Imagesdiagnostic.png" />
    
                <r:RibbonCommand x:Key="HomeButton1" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calculator" LabelDescription="Calc This!" ToolTipTitle="Calculator" ToolTipDescription="Used to do math and stuff" SmallImageSource="Imagescalculator.png" LargeImageSource="Imagescalculator.png" />
                <r:RibbonCommand x:Key="HomeButton2" CanExecute="RibbonCommand_CanExecute" LabelTitle="Calendar" LabelDescription="Schedule This!" ToolTipTitle="Calendar" ToolTipDescription="Schedule and remind yourself of stuff" SmallImageSource="Imagescalendar.png" LargeImageSource="Imagescalendar.png" />
                <r:RibbonCommand x:Key="HomeButton3" CanExecute="RibbonCommand_CanExecute" LabelTitle="Computer" LabelDescription="Format This!" ToolTipTitle="Computer" ToolTipDescription="Where you store your naked pictures" SmallImageSource="Imagescomputer.png" LargeImageSource="Imagescomputer.png" />
    
                <r:RibbonCommand x:Key="MediaEject" CanExecute="RibbonCommand_CanExecute" LabelTitle="Eject" LabelDescription="Eject" ToolTipTitle="Eject" ToolTipDescription="Open the cup holder" SmallImageSource="Imagesbt_eject.png" LargeImageSource="Imagesbt_eject.png" />
                <r:RibbonCommand x:Key="MediaBackward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Previous" LabelDescription="Previous" ToolTipTitle="Previous" ToolTipDescription="Previous Tune" SmallImageSource="Imagesbt_skip_backward.png" LargeImageSource="Imagesbt_skip_backward.png" />
                <r:RibbonCommand x:Key="MediaPlay" CanExecute="RibbonCommand_CanExecute" LabelTitle="Play" LabelDescription="Play" ToolTipTitle="Play" ToolTipDescription="Play Tune" SmallImageSource="Imagesbt_play.png" LargeImageSource="Imagesbt_play.png" />
                <r:RibbonCommand x:Key="MediaStop" CanExecute="RibbonCommand_CanExecute" LabelTitle="Stop" LabelDescription="Stop" ToolTipTitle="Stop" ToolTipDescription="Stop the music" SmallImageSource="Imagesbt_stop.png" LargeImageSource="Imagesbt_stop.png" />
                <r:RibbonCommand x:Key="MediaForward" CanExecute="RibbonCommand_CanExecute" LabelTitle="Next" LabelDescription="Next" ToolTipTitle="Next" ToolTipDescription="Next Tune" SmallImageSource="Imagesbt_skip_forward.png" LargeImageSource="Imagesbt_skip_forward.png" />
    
                <!-- Uncomment below for Office 2007 Blue -->
                <ResourceDictionary.MergedDictionaries>
                    <ResourceDictionary Source="/RibbonControlsLibrary;component/Themes/Office2007Blue.xaml" />
                </ResourceDictionary.MergedDictionaries>
                <!-- Uncomment below for Office 2007 Silver -->
                <!--<ResourceDictionary.MergedDictionaries>
                    <ResourceDictionary Source="/RibbonControlsLibrary;component/Themes/Office2007Silver.xaml" />
                </ResourceDictionary.MergedDictionaries>-->
                <!-- Uncomment below for Office 2007 Black -->
                <!--<ResourceDictionary.MergedDictionaries>
                    <ResourceDictionary Source="/RibbonControlsLibrary;component/Themes/Office2007Black.xaml" />
                </ResourceDictionary.MergedDictionaries>-->
            </ResourceDictionary>
        </r:RibbonWindow.Resources>
    
        <DockPanel>
            <r:Ribbon DockPanel.Dock="Top" Title="My First Ribbon Form" x:Name="mainRibbon">
                <r:Ribbon.QuickAccessToolBar>
                    <r:RibbonQuickAccessToolBar>
                        <r:RibbonButton Command="{StaticResource QATButton}" />
                        <r:RibbonButton Command="{StaticResource MediaEject}" r:RibbonQuickAccessToolBar.Placement="InCustomizeMenu" />
                    </r:RibbonQuickAccessToolBar>
                </r:Ribbon.QuickAccessToolBar>
                <r:Ribbon.ApplicationMenu>
                    <r:RibbonApplicationMenu>
                        <r:RibbonApplicationMenu.Command>
                            <r:RibbonCommand SmallImageSource="Imagesbox.png" LargeImageSource="Imagesbox.png" />
                        </r:RibbonApplicationMenu.Command>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem1}">
                            <TextBlock Text="Item 1 in the list" />
                            <TextBlock Text="Item 2 in the list" />
                            <TextBlock Text="Item 3 in the list" />
                            <TextBlock Text="Item 4 in the list" />
                        </r:RibbonApplicationMenuItem>
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem2}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem3}" />
                        <r:RibbonApplicationMenuItem Command="{StaticResource MenuItem4}" />
                        <r:RibbonApplicationMenu.RecentItemList>
                            <Rectangle Height="300" />
                        </r:RibbonApplicationMenu.RecentItemList>
                    </r:RibbonApplicationMenu>
                </r:Ribbon.ApplicationMenu>
                <r:RibbonTab Label="Home">
                    <r:RibbonGroup>
                        <r:RibbonGroup.Command>
                            <r:RibbonCommand LabelTitle="Programs" />
                        </r:RibbonGroup.Command>
                        <r:RibbonGroup.GroupSizeDefinitions>
                            <r:RibbonGroupSizeDefinitionCollection>
                                <r:RibbonGroupSizeDefinition>
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                    <r:RibbonControlSizeDefinition ImageSize="Large" IsLabelVisible="True" />
                                </r:RibbonGroupSizeDefinition>
                            </r:RibbonGroupSizeDefinitionCollection>
                        </r:RibbonGroup.GroupSizeDefinitions>
    
                        <r:RibbonButton Command="{StaticResource HomeButton1}" />
                        <r:RibbonButton Command="{StaticResource HomeButton2}" />
                        <r:RibbonButton Command="{StaticResource HomeButton3}" />
                    </r:RibbonGroup>
                </r:RibbonTab>
                <r:RibbonTab Label="Media">
                    <r:RibbonGroup>
                        <r:RibbonGroup.Command>
                            <r:RibbonCommand LabelTitle="Media Controls" />
                        </r:RibbonGroup.Command>
                        <r:RibbonControlGroup>
                            <r:RibbonButton Command="{StaticResource MediaEject}" />
                            <r:RibbonButton Command="{StaticResource MediaBackward}" />
                            <r:RibbonButton Command="{StaticResource MediaPlay}" />
                            <r:RibbonButton Command="{StaticResource MediaStop}" />
                            <r:RibbonButton Command="{StaticResource MediaForward}" />
                        </r:RibbonControlGroup>
                    </r:RibbonGroup>
                </r:RibbonTab>
            </r:Ribbon>
            <RichTextBox Height="auto" Width="auto">
                <FlowDocument>
                    <Paragraph>
                        <Hyperlink NavigateUri="http://www.renevo.com">RenEvo Software &amp; Designs</Hyperlink>
                    </Paragraph>
                </FlowDocument>
            </RichTextBox>
        </DockPanel>
    </r:RibbonWindow>
    

![image][14]

And that is it, a very simple look into creating and using the new WPF Ribbon Control.  There are a lot more features, but this is just a taste.

[You can download the code used in this article][6] if you like which includes a whole host of images to play with.  You will however need to get the Ribbon Control and re-adjust the reference to your location before compiling (licensing and all).

Also, Microsoft has [posted an introduction][15] to this which I kind of learned from on [windowsclient.net][16]

![kick it on DotNetKicks.com][17]

![][18]

[1]: http://www.renevo.com/blogs/community_blogs/archive/2009/01/18/windows-7-style-wpf-ribbon.aspx
[2]: http://www.codeplex.com
[3]: http://www.codeplex.com/wpf/Wiki/View.aspx?title=WPF Ribbon Preview
[4]: http://www.renevo.com/blogs/dotnet/image_thumb_178DA5C5.png "image"
[5]: http://www.renevo.com/blogs/dotnet/image_thumb_296CCA43.png "image"
[6]: http://www.renevo.com/files/folders/articles_vbnet/entry2155.aspx
[7]: http://www.renevo.com/blogs/dotnet/image_thumb_3FC26F88.png "image"
[8]: http://www.renevo.com/blogs/dotnet/image_thumb_015CABD5.png "image"
[9]: http://www.renevo.com/blogs/dotnet/image_thumb_17B2511A.png "image"
[10]: http://www.renevo.com/blogs/dotnet/image_thumb_6E3E0FD9.png "image"
[11]: http://www.renevo.com/blogs/dotnet/image_thumb_39A0444F.png "image"
[12]: http://www.renevo.com/blogs/dotnet/image_thumb_741B4423.png "image"
[13]: http://www.renevo.com/blogs/dotnet/image_thumb_587948DE.png "image"
[14]: http://www.renevo.com/blogs/dotnet/image_thumb_3141905A.png "image"
[15]: http://windowsclient.net/wpf/wpf35/wpf-35sp1-ribbon-walkthrough.aspx
[16]: http://windowsclient.net
[17]: http://www.dotnetkicks.com/Services/Images/KickItImageGenerator.ashx?url=http://www.renevo.com/blogs/dotnet/archive/2009/02/10/your-first-wpf-ribbon-application.aspx
[18]: http://renevo.com/aggbug.aspx?PostID=2156

