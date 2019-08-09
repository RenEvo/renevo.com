---
title: Loading Config Files from non-Default Locations
published: true
date: 2008-01-31 21:33:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2008/01/31/loading-config-files-from-non-default-locations.aspx
file: loading-config-files-from-non-default-locations.aspx
path: /blogs/dotnet/archive/2008/01/31/
author: tom anderson
words: 1357
---
I have spent a good deal of my time trying to figure out how to load a *.config file that is compatable with the System.Configuration.Configuration object. Over the years I have done several "fake" .config files, which where nothing more then glorified structured xml files that I manually load.  I have also even created fake 0 byte files so I could load the .config using the OpenExeConfiguration.

Until recently, this has been a viable solution, and since these files where mostly hidden by ClickOnce installations and storing them in the Application Data file structure. Now though, I got a hair up my butt to do it correctly.

Enter the System.Configuration.ExeConfigurationFileMap class.  This lets you specify an exe config file name, and allows you to load it up at runtime without having a valid executable to load off of.  So now, I can save logged in specific users along with windows users configuration files. 

For example, my goal was to do the following:   
Allow the user to setup connections and save them to the Application DataProduct folder under their user account.   
After authenticating with the server, I want to store per-logged in user settings.

Just short of using datasets and/or serialized objects, I was using folders to seperate between the different logged in users.

Now, to the solution:

       1:  Imports System.Configuration    
       2:       
       3:  Public Class LocalConfiguration    
       4:       
       5:      Public Shared Function GetConfig(ByVal Path As String) As Configuration    
       6:          Dim retVal As Configuration = Nothing    
       7:       
       8:          Dim configFileMap As New ExeConfigurationFileMap()    
       9:          configFileMap.ExeConfigFilename = Path    
      10:       
      11:          retVal = ConfigurationManager.OpenMappedExeConfiguration(configFileMap, ConfigurationUserLevel.None)    
      12:       
      13:          Return retVal    
      14:      End Function    
      15:       
      16:  End Class

That is it, now you can simply load a valid application config file and get a System.Configuration.Configuration object back.  All it takes is the following line of code:

       1:          Dim config As Configuration = LocalConfiguration.GetConfig("./data/user.config")

The beuty is, the file doesn't even have to exist. When you save the config, it will automatically create it if it doesn't exist.

The next step, which turned out to be a bit tricky, was actually saving stuff, like appSettings, to the file, specifically if they didn't exist to begin with.

       1:      Private Sub LoadSettings()    
       2:          Dim config As Configuration = LocalConfiguration.GetConfig("./data/user.config")    
       3:          Dim newSection As AppSettingsSection = Nothing    
       4:       
       5:          If config.Sections("MySettings") Is Nothing Then    
       6:              newSection = New AppSettingsSection()    
       7:              newSection.SectionInformation.AllowExeDefinition = ConfigurationAllowExeDefinition.MachineToLocalUser    
       8:              config.Sections.Add("MySettings", newSection)    
       9:          Else    
      10:              newSection = config.Sections("MySettings")    
      11:          End If    
      12:       
      13:          If newSection.Settings("Text") Is Nothing Then    
      14:              newSection.Settings.Add("Text", Me.Text)    
      15:          End If    
      16:       
      17:          Me.Text = newSection.Settings("Text").Value    
      18:       
      19:          config.Save(ConfigurationSaveMode.Minimal)    
      20:       
      21:      End Sub

On line 2 I call the method from the first example to get the configuration object, next I define an AppSettingsSection which is a key/value pair collection that is used in the <appSettings> in normal .config files.  Depeding on whether the section exists or not in the current file I either set the newSection to the existing reference, or create a new one, set the SectionInformation to Allow Machine to Local User, then work with the keys as normal.

Finally, I save the config file, just incase it didn't exist before I read it into the application.

Once saved, the user.config looks like this:

       1:  <?xml version="1.0" encoding="utf-8"?>    
       2:  <configuration>    
       3:      <configSections>    
       4:          <section name="MySettings" type="System.Configuration.AppSettingsSection, System.Configuration, Version=2.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" allowExeDefinition="MachineToLocalUser" />    
       5:      </configSections>    
       6:      <MySettings>    
       7:          <add key="Text" value="My Application" />    
       8:      </MySettings>    
       9:  </configuration>

As you can see, it created a new config section named "MySettings" using the System.Configuration.AppSettingsSection.

Pretty simple once you get the hang of it, and I could very easily see a CAB based, or plugin based application really taking advantage of loading different config sections per module/plugin, and reading other settings from other modules/plugins.

In the next discovery, I will be trying to figure out how to merge multiple configurations into a single config file to ease the reading of multi-configs.





