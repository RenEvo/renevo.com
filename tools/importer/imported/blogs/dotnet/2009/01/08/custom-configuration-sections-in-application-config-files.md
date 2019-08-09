---
title: Custom Configuration Sections in Application Config files
published: true
date: 2009-01-08 18:18:03 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2009/01/08/custom-configuration-sections-in-application-config-files.aspx
file: custom-configuration-sections-in-application-config-files.aspx
path: /blogs/dotnet/archive/2009/01/08/
author: tom anderson
words: 14674
---
A colleague of mine approached me the other day on how to build custom configuration section handlers in .Net.  As I have said [in previous articles, I have been using "real" .config files now instead of pseudo .config, .cfg, or .xml files for complex application settings][1], and this has proven very portable for the configurations.

When creating custom configuration sections, there are essentially three steps.

1. Provide a way to load the configuration section 
2. Define properties for the section 
3. Read/Write to the base object's collection 

A bit of knowledge here.  ConfigurationSection and ConfigurationElement are essentially property bags, you will be directing all of your properties in your classes to read and write to the Item collection where your property name is the key, and the value is, you guessed it, the value.

So lets start off easy, lets create the configuration section first and lets make it something simple, like a configuration section for our application windows.

       1:  Imports System.Configuration    
       2:       
       3:  Public Class WindowSettings    
       4:      Inherits ConfigurationSection    
       5:       
       6:  End Class

Basic class, yes, now we want to define some primary properties for the section, lets add "applicationKey" so this is re-usable.

Above I mentioned that we are simply creating wrappers for the internal property bag, we also need to add an attribute to the property specifying the name of the attribute in the actual config file.

       1:      <ConfigurationProperty("applicationKey")> _    
       2:      Public Property ApplicationKey() As String    
       3:          Get    
       4:              Return MyBase.Item("applicationKey")    
       5:          End Get    
       6:          Set(ByVal value As String)    
       7:              MyBase.Item("applicationKey") = value    
       8:          End Set    
       9:      End Property

Pretty simply so far?

Now is an optional step that I like to use for testing, it really helps me work with these sections, below are some helper methods (both static and instanced) that I build for any custom section handlers, a base class is also recommended if you are going to be doing a lot of these.

       1:  #Region " Base Implementation For Load/Save "    
       2:       
       3:      'only initialized internally    
       4:      Protected Sub New()    
       5:       
       6:      End Sub    
       7:       
       8:      ''' <summary>    
       9:      ''' Internal placeholder for configuration file    
      10:      ''' </summary>    
      11:      ''' <remarks></remarks>    
      12:      Protected m_BaseConfiguration As Configuration = Nothing    
      13:       
      14:      ''' <summary>    
      15:      ''' Gets the custom configuration by path specifying a custom section name    
      16:      ''' </summary>    
      17:      ''' <param name="path">Path to the configuration file</param>    
      18:      ''' <param name="section">Section in configuration file to load</param>    
      19:      Public Shared Function LoadConfigSection(ByVal path As String, _    
      20:                                               ByVal section As String) As WindowSettings    
      21:          Dim retVal As New WindowSettings(), config As Configuration = Nothing    
      22:       
      23:          Dim configFileMap As New ExeConfigurationFileMap()    
      24:       
      25:          configFileMap.ExeConfigFilename = path    
      26:          config = ConfigurationManager.OpenMappedExeConfiguration( _    
      27:                                                              configFileMap, _    
      28:                                                              ConfigurationUserLevel.None)    
      29:       
      30:          If config.Sections(section) Is Nothing Then    
      31:              config.Sections.Add(section, New WindowSettings)    
      32:          End If    
      33:       
      34:          retVal = TryCast(config.GetSection(section), WindowSettings)    
      35:          retVal.m_BaseConfiguration = config    
      36:       
      37:          Return retVal    
      38:      End Function    
      39:       
      40:      ''' <summary>    
      41:      ''' Gets the custom configuration by path    
      42:      ''' </summary>    
      43:      ''' <param name="path"></param>    
      44:      ''' <returns></returns>    
      45:      ''' <remarks></remarks>    
      46:      Public Shared Function LoadConfigSection(ByVal path As String) As WindowSettings    
      47:          Return LoadConfigSection(path, "WindowSettings")    
      48:      End Function    
      49:       
      50:      ''' <summary>    
      51:      ''' Gets the custom configuration    
      52:      ''' </summary>    
      53:      Public Shared Function LoadConfigSection() As WindowSettings    
      54:          Return LoadConfigSection(Application.ExecutablePath & ".config")    
      55:      End Function    
      56:       
      57:      ''' <summary>    
      58:      ''' Saves the current configuration    
      59:      ''' </summary>    
      60:      ''' <remarks></remarks>    
      61:      Public Sub Save()    
      62:          m_BaseConfiguration.Save(ConfigurationSaveMode.Minimal)    
      63:      End Sub    
      64:       
      65:  #End Region

This is the big bread and butter code that makes these easy to work with. First it makes the class construtor protected which prevents the class from accidentally being created, adds three overloaded methods for loading the config section, and then provides a save routine.

I am sure somewhere down the road I will be creating a generic implementation of this class so you can simply inherit from ConfigurationSection(of T).

So currently we have implemented all three steps required to create a basic configuration section in a standard .config file.

To load, set, and save simply do the following line of code.

       1:          Dim config As WindowSettings= WindowSettings.LoadConfigSection()    
       2:          config.ApplicationKey = Application.ProductName    
       3:          config.Save()

To retrieve the value, simply use the following one liner.

       1:          MessageBox.Show(WindowSettings.LoadConfigSection().ApplicationKey)

Well, all that is fine and dandy, but what if we need to do some collections in our settings, like saving each forms top, left, width, and height values?

Then we will create a ConfigurationElement and ConfigurationElementCollection.  First with the ConfigurationElement, it is almost identical to the ConfigurationSection as far as wrapping the property bags.  Although since we are setting form settings, lets add some default values to these properties so our forms are 0,0 location and 0,0 sizes.

       1:      Public Class WindowSetting    
       2:          Inherits ConfigurationElement    
       3:       
       4:          <ConfigurationProperty("id")> _    
       5:          Public Property ID() As String    
       6:              Get    
       7:                  Return MyBase.Item("id")    
       8:              End Get    
       9:              Set(ByVal value As String)    
      10:                  MyBase.Item("id") = value    
      11:              End Set    
      12:          End Property    
      13:       
      14:          <ConfigurationProperty("top", DefaultValue:=100)> _    
      15:          Public Property Top() As Integer    
      16:              Get    
      17:                  Return MyBase.Item("top")    
      18:              End Get    
      19:              Set(ByVal value As Integer)    
      20:                  MyBase.Item("top") = value    
      21:              End Set    
      22:          End Property    
      23:       
      24:          <ConfigurationProperty("left", DefaultValue:=100)> _    
      25:          Public Property Left() As Integer    
      26:              Get    
      27:                  Return MyBase.Item("left")    
      28:              End Get    
      29:              Set(ByVal value As Integer)    
      30:                  MyBase.Item("left") = value    
      31:              End Set    
      32:          End Property    
      33:       
      34:          <ConfigurationProperty("height", DefaultValue:=600)> _    
      35:          Public Property Height() As Integer    
      36:              Get    
      37:                  Return MyBase.Item("height")    
      38:              End Get    
      39:              Set(ByVal value As Integer)    
      40:                  MyBase.Item("height") = value    
      41:              End Set    
      42:          End Property    
      43:       
      44:          <ConfigurationProperty("width", DefaultValue:=800)> _    
      45:          Public Property Width() As Integer    
      46:              Get    
      47:                  Return MyBase.Item("width")    
      48:              End Get    
      49:              Set(ByVal value As Integer)    
      50:                  MyBase.Item("width") = value    
      51:              End Set    
      52:          End Property    
      53:      End Class

 

This class is about as straight forward as they come, define properties, route to property bag, define attribute name, set default value.  We also added an ID property which will act as our key for our collection.

Which brings us to our Collection. ConfigurationElementCollections have two must overrides, CreateNewElement() and GetElementKey()  both of these methods are pretty easy to implement, and below is the most basic ConfigurationElementCollection implementation you can get by with. You must also specify an attribute to the class in order to name the element in the config file.

       1:      <ConfigurationCollection(GetType(WindowSetting), AddItemName:="windows")> _    
       2:      Public Class WindowSettingCollection    
       3:          Inherits ConfigurationElementCollection    
       4:       
       5:          Protected Overloads Overrides Function CreateNewElement() _    
       6:                                  As System.Configuration.ConfigurationElement    
       7:              Return New WindowSetting    
       8:          End Function    
       9:       
      10:          Protected Overrides Function GetElementKey( _    
      11:                                  ByVal element As System.Configuration.ConfigurationElement) _    
      12:                                  As Object    
      13:              Return DirectCast(element, WindowSetting).ID    
      14:          End Function    
      15:      End Class

That is it, but how much fun would it be if I just showed you that bit of code, and said deal with figuring out the rest on your own?  Personally, I like to add **lots** of helper methods to this class, again, this is another huge candidate for a generic.

       1:  <ConfigurationCollection(GetType(WindowSetting), AddItemName:="windows")> _    
       2:      Public Class WindowSettingCollection    
       3:          Inherits ConfigurationElementCollection    
       4:       
       5:          ''' <summary>    
       6:          ''' Creates a new element with default properties    
       7:          ''' </summary>    
       8:          Protected Overloads Overrides Function CreateNewElement() _    
       9:                                          As System.Configuration.ConfigurationElement    
      10:              Return New WindowSetting()    
      11:          End Function    
      12:       
      13:          ''' <summary>    
      14:          ''' Creates a new element with the id specified    
      15:          ''' </summary>    
      16:          ''' <param name="id">id of the new element</param>    
      17:          Protected Overloads Function CreateNewElement(ByVal id As String) _    
      18:                                          As System.Configuration.ConfigurationElement    
      19:              Return New WindowSetting() With {.ID = id}    
      20:          End Function    
      21:       
      22:          ''' <summary>    
      23:          ''' Gets an element's key    
      24:          ''' </summary>    
      25:          ''' <param name="element">element to test</param>    
      26:          Protected Overrides Function GetElementKey( _    
      27:                                          ByVal element As System.Configuration.ConfigurationElement) _    
      28:                                          As Object    
      29:              Return DirectCast(element, WindowSetting).ID    
      30:          End Function    
      31:       
      32:          ''' <summary>    
      33:          ''' Adds a new element to the collection    
      34:          ''' </summary>    
      35:          ''' <param name="element"></param>    
      36:          ''' <remarks></remarks>    
      37:          Public Sub Add(ByVal element As WindowSetting)    
      38:              MyBase.BaseAdd(element)    
      39:          End Sub    
      40:       
      41:          ''' <summary>    
      42:          ''' Adds a new element to the collection by id    
      43:          ''' </summary>    
      44:          ''' <param name="id">id of the new element</param>    
      45:          Public Function AddNew(ByVal id As String) As WindowSetting    
      46:              Dim newElement As WindowSetting = Me.CreateNewElement(id)    
      47:              Add(newElement)    
      48:              Return newElement    
      49:          End Function    
      50:       
      51:          ''' <summary>    
      52:          ''' Adds a new element to the collection by form    
      53:          ''' </summary>    
      54:          ''' <param name="form">form for the new element</param>    
      55:          Public Function AddNew(ByVal form As System.Windows.Forms.Form) As WindowSetting    
      56:              Dim newElement As WindowSetting = Me.CreateNewElement(form.Name)    
      57:              newElement.Top = form.Top    
      58:              newElement.Left = form.Left    
      59:              newElement.Height = form.Height    
      60:              newElement.Width = form.Width    
      61:              Add(newElement)    
      62:              Return newElement    
      63:          End Function    
      64:       
      65:          ''' <summary>    
      66:          ''' Removes an element by id    
      67:          ''' </summary>    
      68:          ''' <param name="id">id of the element to remove</param>    
      69:          Public Sub Remove(ByVal id As String)    
      70:              MyBase.BaseRemove(id)    
      71:          End Sub    
      72:       
      73:          ''' <summary>    
      74:          ''' Clears all elements from the collection    
      75:          ''' </summary>    
      76:          Public Sub Clear()    
      77:              MyBase.BaseClear()    
      78:          End Sub    
      79:       
      80:          ''' <summary>    
      81:          ''' Retrieves an item from the collection by index    
      82:          ''' </summary>    
      83:          ''' <param name="index">index of the element</param>    
      84:          Default Public Overloads ReadOnly Property Item(ByVal index As Integer) As WindowSetting    
      85:              Get    
      86:                  Return MyBase.BaseGet(index)    
      87:              End Get    
      88:          End Property    
      89:       
      90:          ''' <summary>    
      91:          ''' Retrieves an item from the collection by id    
      92:          ''' </summary>    
      93:          ''' <param name="id">id of the element</param>    
      94:          Default Public Overloads ReadOnly Property Item(ByVal id As String) As WindowSetting    
      95:              Get    
      96:                  'auto-add if not exists    
      97:                  If MyBase.BaseGet(id) Is Nothing Then    
      98:                      Me.AddNew(id)    
      99:                  End If    
     100:       
     101:                  Return Me.BaseGet(id)    
     102:              End Get    
     103:          End Property    
     104:       
     105:      End Class

All of those additional methods will give you all the tools you need to create a CRUD interface to the collection. I also added in an AddNew that simply takes a form object.

Finally to add the collection to the section, simply add a property for it.

       1:      <ConfigurationProperty("windows")> _    
       2:      Public ReadOnly Property Windows() As WindowSettingCollection    
       3:          Get    
       4:              Return Me.Item("windows")    
       5:          End Get    
       6:      End Property

I also added IDisposable support to the WindowSettings class for ease of use as well as two helper methods for loading and saving form settings.

       1:      ''' <summary>    
       2:      ''' Helper method to load form settings    
       3:      ''' </summary>    
       4:      ''' <param name="form">form to load</param>    
       5:      Public Sub LoadFormSettings(ByVal form As System.Windows.Forms.Form)    
       6:          With form    
       7:              .Top = Me.Windows(form.Name).Top    
       8:              .Left = Me.Windows(form.Name).Left    
       9:              .Height = Me.Windows(form.Name).Height    
      10:              .Width = Me.Windows(form.Name).Width    
      11:          End With    
      12:      End Sub    
      13:       
      14:      ''' <summary>    
      15:      ''' Helper method to save form settings    
      16:      ''' </summary>    
      17:      ''' <param name="form"></param>    
      18:      ''' <param name="saveConfig">Optional parameter, when true will save the configuration file</param>    
      19:      ''' <remarks></remarks>    
      20:      Public Sub SaveFormSettings(ByVal form As System.Windows.Forms.Form, Optional ByVal saveConfig As Boolean = False)    
      21:          With form    
      22:              Me.Windows(form.Name).Top = .Top    
      23:              Me.Windows(form.Name).Left = .Left    
      24:              Me.Windows(form.Name).Height = .Height    
      25:              Me.Windows(form.Name).Width = .Width    
      26:          End With    
      27:       
      28:          If saveConfig Then    
      29:              Me.Save()    
      30:          End If    
      31:      End Sub

Now to use it!  In our forms Load event handler, add the following code:

       1:          Using config As WindowSettings = WindowSettings.LoadConfigSection()    
       2:              config.ApplicationKey = Application.ProductName    
       3:              config.LoadFormSettings(Me)    
       4:          End Using

In our forms closing event handler, add the following code:

       1:          Using config As WindowSettings = WindowSettings.LoadConfigSection()    
       2:              config.ApplicationKey = Application.ProductName    
       3:              config.SaveFormSettings(Me, True)    
       4:          End Using

 

And for the curious, here is the resulting app.config.

       1:  <configuration>    
       2:      <configSections>    
       3:          <section name="WindowSettings" type="SandBoxVB.WindowSettings, SandBoxVB, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null" />    
       4:      </configSections>    
       5:      <WindowSettings applicationKey="SandBoxVB">    
       6:          <windows>    
       7:              <windows id="Form1" top="215" left="313" height="334" width="479" />    
       8:          </windows>    
       9:      </WindowSettings>    
      10:  </configuration>

 

Below is the full code for the WindowSettings Section Handler.

       1:  Imports System.Configuration    
       2:       
       3:  Public Class WindowSettings    
       4:      Inherits ConfigurationSection    
       5:      Implements IDisposable    
       6:       
       7:  #Region " Base Implementation For Load/Save "    
       8:       
       9:      'only initialized internally    
      10:      Protected Sub New()    
      11:       
      12:      End Sub    
      13:       
      14:      ''' <summary>    
      15:      ''' Internal placeholder for configuration file    
      16:      ''' </summary>    
      17:      ''' <remarks></remarks>    
      18:      Protected m_BaseConfiguration As Configuration = Nothing    
      19:       
      20:      ''' <summary>    
      21:      ''' Gets the custom configuration by path specifying a custom section name    
      22:      ''' </summary>    
      23:      ''' <param name="path">Path to the configuration file</param>    
      24:      ''' <param name="section">Section in configuration file to load</param>    
      25:      Public Shared Function LoadConfigSection(ByVal path As String, _    
      26:                                               ByVal section As String) As WindowSettings    
      27:          Dim retVal As New WindowSettings(), config As Configuration = Nothing    
      28:       
      29:          Dim configFileMap As New ExeConfigurationFileMap()    
      30:       
      31:          configFileMap.ExeConfigFilename = path    
      32:          config = ConfigurationManager.OpenMappedExeConfiguration( _    
      33:                                                              configFileMap, _    
      34:                                                              ConfigurationUserLevel.None)    
      35:       
      36:          If config.Sections(section) Is Nothing Then    
      37:              config.Sections.Add(section, New WindowSettings)    
      38:          End If    
      39:       
      40:          retVal = TryCast(config.GetSection(section), WindowSettings)    
      41:          retVal.m_BaseConfiguration = config    
      42:       
      43:          Return retVal    
      44:      End Function    
      45:       
      46:      ''' <summary>    
      47:      ''' Gets the custom configuration by path    
      48:      ''' </summary>    
      49:      ''' <param name="path"></param>    
      50:      ''' <returns></returns>    
      51:      ''' <remarks></remarks>    
      52:      Public Shared Function LoadConfigSection(ByVal path As String) As WindowSettings    
      53:          Return LoadConfigSection(path, "WindowSettings")    
      54:      End Function    
      55:       
      56:      ''' <summary>    
      57:      ''' Gets the custom configuration    
      58:      ''' </summary>    
      59:      Public Shared Function LoadConfigSection() As WindowSettings    
      60:          Return LoadConfigSection(Application.ExecutablePath & ".config")    
      61:      End Function    
      62:       
      63:      ''' <summary>    
      64:      ''' Saves the current configuration    
      65:      ''' </summary>    
      66:      ''' <remarks></remarks>    
      67:      Public Sub Save()    
      68:          m_BaseConfiguration.Save(ConfigurationSaveMode.Minimal)    
      69:      End Sub    
      70:       
      71:  #End Region    
      72:       
      73:      ''' <summary>    
      74:      ''' Application key for the project,     
      75:      ''' Useful when loading multiples from the save file.    
      76:      ''' </summary>    
      77:      <ConfigurationProperty("applicationKey")> _    
      78:      Public Property ApplicationKey() As String    
      79:          Get    
      80:              Return MyBase.Item("applicationKey")    
      81:          End Get    
      82:          Set(ByVal value As String)    
      83:              MyBase.Item("applicationKey") = value    
      84:          End Set    
      85:      End Property    
      86:       
      87:      ''' <summary>    
      88:      ''' Retrieves a collection of windows in the application    
      89:      ''' </summary>    
      90:      <ConfigurationProperty("windows")> _    
      91:      Public ReadOnly Property Windows() As WindowSettingCollection    
      92:          Get    
      93:              Return Me.Item("windows")    
      94:          End Get    
      95:      End Property    
      96:       
      97:      ''' <summary>    
      98:      ''' Helper method to load form settings    
      99:      ''' </summary>    
     100:      ''' <param name="form">form to load</param>    
     101:      Public Sub LoadFormSettings(ByVal form As System.Windows.Forms.Form)    
     102:          With form    
     103:              .Top = Me.Windows(form.Name).Top    
     104:              .Left = Me.Windows(form.Name).Left    
     105:              .Height = Me.Windows(form.Name).Height    
     106:              .Width = Me.Windows(form.Name).Width    
     107:          End With    
     108:      End Sub    
     109:       
     110:      ''' <summary>    
     111:      ''' Helper method to save form settings    
     112:      ''' </summary>    
     113:      ''' <param name="form"></param>    
     114:      ''' <param name="saveConfig">    
     115:      ''' Optional parameter, when true will save the configuration file    
     116:      ''' </param>    
     117:      ''' <remarks></remarks>    
     118:      Public Sub SaveFormSettings(ByVal form As System.Windows.Forms.Form, _    
     119:                                      Optional ByVal saveConfig As Boolean = False)    
     120:          With form    
     121:              Me.Windows(form.Name).Top = .Top    
     122:              Me.Windows(form.Name).Left = .Left    
     123:              Me.Windows(form.Name).Height = .Height    
     124:              Me.Windows(form.Name).Width = .Width    
     125:          End With    
     126:       
     127:          If saveConfig Then    
     128:              Me.Save()    
     129:          End If    
     130:      End Sub    
     131:       
     132:      Public Class WindowSetting    
     133:          Inherits ConfigurationElement    
     134:       
     135:          <ConfigurationProperty("id")> _    
     136:          Public Property ID() As String    
     137:              Get    
     138:                  Return MyBase.Item("id")    
     139:              End Get    
     140:              Set(ByVal value As String)    
     141:                  MyBase.Item("id") = value    
     142:              End Set    
     143:          End Property    
     144:       
     145:          <ConfigurationProperty("top", DefaultValue:=100)> _    
     146:          Public Property Top() As Integer    
     147:              Get    
     148:                  Return MyBase.Item("top")    
     149:              End Get    
     150:              Set(ByVal value As Integer)    
     151:                  MyBase.Item("top") = value    
     152:              End Set    
     153:          End Property    
     154:       
     155:          <ConfigurationProperty("left", DefaultValue:=100)> _    
     156:          Public Property Left() As Integer    
     157:              Get    
     158:                  Return MyBase.Item("left")    
     159:              End Get    
     160:              Set(ByVal value As Integer)    
     161:                  MyBase.Item("left") = value    
     162:              End Set    
     163:          End Property    
     164:       
     165:          <ConfigurationProperty("height", DefaultValue:=600)> _    
     166:          Public Property Height() As Integer    
     167:              Get    
     168:                  Return MyBase.Item("height")    
     169:              End Get    
     170:              Set(ByVal value As Integer)    
     171:                  MyBase.Item("height") = value    
     172:              End Set    
     173:          End Property    
     174:       
     175:          <ConfigurationProperty("width", DefaultValue:=800)> _    
     176:          Public Property Width() As Integer    
     177:              Get    
     178:                  Return MyBase.Item("width")    
     179:              End Get    
     180:              Set(ByVal value As Integer)    
     181:                  MyBase.Item("width") = value    
     182:              End Set    
     183:          End Property    
     184:      End Class    
     185:       
     186:      <ConfigurationCollection(GetType(WindowSetting), AddItemName:="windows")> _    
     187:      Public Class WindowSettingCollection    
     188:          Inherits ConfigurationElementCollection    
     189:       
     190:          ''' <summary>    
     191:          ''' Creates a new element with default properties    
     192:          ''' </summary>    
     193:          Protected Overloads Overrides Function CreateNewElement() _    
     194:                                          As System.Configuration.ConfigurationElement    
     195:              Return New WindowSetting()    
     196:          End Function    
     197:       
     198:          ''' <summary>    
     199:          ''' Creates a new element with the id specified    
     200:          ''' </summary>    
     201:          ''' <param name="id">id of the new element</param>    
     202:          Protected Overloads Function CreateNewElement(ByVal id As String) _    
     203:                                          As System.Configuration.ConfigurationElement    
     204:              Return New WindowSetting() With {.ID = id}    
     205:          End Function    
     206:       
     207:          ''' <summary>    
     208:          ''' Gets an element's key    
     209:          ''' </summary>    
     210:          ''' <param name="element">element to test</param>    
     211:          Protected Overrides Function GetElementKey( _    
     212:                                  ByVal element As System.Configuration.ConfigurationElement) _    
     213:                                  As Object    
     214:              Return DirectCast(element, WindowSetting).ID    
     215:          End Function    
     216:       
     217:          ''' <summary>    
     218:          ''' Adds a new element to the collection    
     219:          ''' </summary>    
     220:          ''' <param name="element"></param>    
     221:          ''' <remarks></remarks>    
     222:          Public Sub Add(ByVal element As WindowSetting)    
     223:              MyBase.BaseAdd(element)    
     224:          End Sub    
     225:       
     226:          ''' <summary>    
     227:          ''' Adds a new element to the collection by id    
     228:          ''' </summary>    
     229:          ''' <param name="id">id of the new element</param>    
     230:          Public Function AddNew(ByVal id As String) As WindowSetting    
     231:              Dim newElement As WindowSetting = Me.CreateNewElement(id)    
     232:              Add(newElement)    
     233:              Return newElement    
     234:          End Function    
     235:       
     236:          ''' <summary>    
     237:          ''' Adds a new element to the collection by form    
     238:          ''' </summary>    
     239:          ''' <param name="form">form for the new element</param>    
     240:          Public Function AddNew(ByVal form As System.Windows.Forms.Form) _    
     241:                                                                  As WindowSetting    
     242:              Dim newElement As WindowSetting = Me.CreateNewElement(form.Name)    
     243:              newElement.Top = form.Top    
     244:              newElement.Left = form.Left    
     245:              newElement.Height = form.Height    
     246:              newElement.Width = form.Width    
     247:              Add(newElement)    
     248:              Return newElement    
     249:          End Function    
     250:       
     251:          ''' <summary>    
     252:          ''' Removes an element by id    
     253:          ''' </summary>    
     254:          ''' <param name="id">id of the element to remove</param>    
     255:          Public Sub Remove(ByVal id As String)    
     256:              MyBase.BaseRemove(id)    
     257:          End Sub    
     258:       
     259:          ''' <summary>    
     260:          ''' Clears all elements from the collection    
     261:          ''' </summary>    
     262:          Public Sub Clear()    
     263:              MyBase.BaseClear()    
     264:          End Sub    
     265:       
     266:          ''' <summary>    
     267:          ''' Retrieves an item from the collection by index    
     268:          ''' </summary>    
     269:          ''' <param name="index">index of the element</param>    
     270:          Default Public Overloads ReadOnly Property Item(ByVal index As Integer) _    
     271:                                                                      As WindowSetting    
     272:              Get    
     273:                  Return MyBase.BaseGet(index)    
     274:              End Get    
     275:          End Property    
     276:       
     277:          ''' <summary>    
     278:          ''' Retrieves an item from the collection by id    
     279:          ''' </summary>    
     280:          ''' <param name="id">id of the element</param>    
     281:          Default Public Overloads ReadOnly Property Item(ByVal id As String) _    
     282:                                                                      As WindowSetting    
     283:              Get    
     284:                  'auto-add if not exists    
     285:                  If MyBase.BaseGet(id) Is Nothing Then    
     286:                      Me.AddNew(id)    
     287:                  End If    
     288:       
     289:                  Return Me.BaseGet(id)    
     290:              End Get    
     291:          End Property    
     292:       
     293:      End Class    
     294:       
     295:  #Region " IDisposable Support "    
     296:       
     297:      Private disposedValue As Boolean = False        ' To detect redundant calls    
     298:       
     299:      ' IDisposable    
     300:      Protected Overridable Sub Dispose(ByVal disposing As Boolean)    
     301:          If Not Me.disposedValue Then    
     302:              If disposing Then    
     303:       
     304:              End If    
     305:       
     306:          End If    
     307:          Me.disposedValue = True    
     308:      End Sub    
     309:       
     310:      ' This code added by Visual Basic to correctly implement the disposable pattern.    
     311:      Public Sub Dispose() Implements IDisposable.Dispose    
     312:          ' Do not change this code.    
     313:          Dispose(True)    
     314:          GC.SuppressFinalize(Me)    
     315:      End Sub    
     316:       
     317:  #End Region    
     318:       
     319:  End Class    
     320:       
     321:   

 

C# Version provided by request.

 

> _*Sorry about formatting, I don't normally use underscores to break lines other than for attributes, but the page is skinny and I didn't want any overrun._



[1]: http://www.renevo.com/blogs/dotnet/archive/2008/01/31/loading-config-files-from-non-default-locations.aspx


