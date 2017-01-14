---
title: Combo Box Enumerations with Titles
published: true
date: 2008-08-06 20:08:38 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2008/08/06/combo-box-enumerations-with-titles.aspx
file: combo-box-enumerations-with-titles.aspx
path: /blogs/dotnet/archive/2008/08/06/
author: tom anderson
words: 4605
---
Have you ever just wanted to populate a combo box with values from an enumeration, but hated the fact that it dealt with the name of the item in the enumeration, instead of some snazzy string? Take the following Enum as an example:
    
    
       1:  Public Enum TitledValues
    
    
       2:      FirstName
    
    
       3:      LastName
    
    
       4:      Address1
    
    
       5:      Address2
    
    
       6:      City
    
    
       7:      State
    
    
       8:      ZipCode
    
    
       9:      PhoneNumber
    
    
      10:      EmailAddress
    
    
      11:  End Enum

For some of those properties displaying "City" would be fine, but who wants to show a user "FirstName" ?

I have figured out a little technique using a custom type converter and attributes to be able at code time set the titles from the Enum instead of having to do huge select cases when handling the drop down events.

The first thing to do is to create a new attribute called "EnumTitleAttribute", this attribute is pretty straight forward, it simply contains a Title property.
    
    
       1:  <AttributeUsage(AttributeTargets.Field)> _
    
    
       2:  Public Class EnumTitleAttribute
    
    
       3:      Inherits Attribute
    
    
       4:   
    
    
       5:      Private m_Title As String = String.Empty
    
    
       6:      Public Property Title() As String
    
    
       7:          Get
    
    
       8:              Return m_Title
    
    
       9:          End Get
    
    
      10:          Set(ByVal value As String)
    
    
      11:              m_Title = value
    
    
      12:          End Set
    
    
      13:      End Property
    
    
      14:   
    
    
      15:  End Class

The AttributeUsage attribute is restricting usage of this attribute to fields, which is what enum values are stored as in the object.

The next step was to create a type converter to work with these new custom attributes, since we want this to be a global type of enum, we are going to implement it as a generic.
    
    
       1:  Public Class EnumTitleTypeConverter(Of T)
    
    
       2:      Inherits TypeConverter
    
    
       3:   
    
    
       4:      Public Overrides Function ConvertTo(ByVal context As ITypeDescriptorContext, ByVal culture As CultureInfo, ByVal value As Object, ByVal destinationType As System.Type) As Object
    
    
       5:          If value.GetType Is GetType(T) AndAlso destinationType Is GetType(String) Then
    
    
       6:              If value.GetType.GetField(value.ToString).GetCustomAttributes(GetType(EnumTitleAttribute), True).Length > 0 Then
    
    
       7:                  Return DirectCast(value.GetType.GetField(value.ToString).GetCustomAttributes(GetType(EnumTitleAttribute), True)(0), EnumTitleAttribute).Title
    
    
       8:              Else
    
    
       9:                  Return value.ToString
    
    
      10:              End If
    
    
      11:          Else
    
    
      12:              Return MyBase.ConvertTo(context, culture, value, destinationType)
    
    
      13:          End If
    
    
      14:      End Function
    
    
      15:  End Class

The only functionality we really care about in this class is the "ConvertTo" method.  We want to check to see if the value being converted is the type of our generic (of T) and that the destination type is string.

The basics of this method are check the input parameters for the proper conversion types, if they match, see if the field in the type has the EnumTitleAttribute assigned to it, and if so, we are going to get the first one only and return the Title property, otherwise, we are simply going to return the default ToString on the value object.  If none of these match, we will simply let the standard type converter try to deal with it, where it generally just returns the ToString of the object.

Now, to apply the title's to an enumeration, we simply add the attributes to the enum fields and set the TypeConverter.
    
    
       1:  <TypeConverter(GetType(EnumTitleTypeConverter(Of TitledValues)))> _
    
    
       2:  Public Enum TitledValues
    
    
       3:      <EnumTitle(Title:="First Name")> _
    
    
       4:      FirstName
    
    
       5:      <EnumTitle(Title:="Last Name")> _
    
    
       6:      LastName
    
    
       7:      <EnumTitle(Title:="Address 1")> _
    
    
       8:      Address1
    
    
       9:      <EnumTitle(Title:="Address 2")> _
    
    
      10:      Address2
    
    
      11:      City
    
    
      12:      State
    
    
      13:      <EnumTitle(Title:="Zip Code")> _
    
    
      14:      ZipCode
    
    
      15:      <EnumTitle(Title:="Phone Number")> _
    
    
      16:      PhoneNumber
    
    
      17:      <EnumTitle(Title:="Email Address")> _
    
    
      18:      EmailAddress
    
    
      19:  End Enum

So the first blaring question is, how do you use it?  Create a form (or use an existing one), add a combo box (for this example, name it uxFieldNames) and set the DropDownStyle to DropDownList, then in the form load, do the following code.
    
    
       1:  Me.uxFieldNames.DataSource = [Enum].GetValues(GetType(TitledValues))

And then in the SelectedIndexChanged event for the ComboBox, you can find out the selected item with the following code
    
    
       1:          If Me.uxFieldNames.SelectedIndex > -1 Then
    
    
       2:              Select Case DirectCast(uxFieldNames.SelectedItem, TitledValues)
    
    
       3:                  Case TitledValues.FirstName
    
    
       4:                      'do something with the first name
    
    
       5:                  Case TitledValues.LastName
    
    
       6:                      'do something with the last name
    
    
       7:              End Select
    
    
       8:          End If

Without hardly any additional effort, especially from your UI coding, you end up with this result:

![image][1]

instead of this:

![image][2]

Good luck, and happy coding!

 

Full Source Code:
    
    
       1:  Imports System.ComponentModel
    
    
       2:  Imports System.Globalization
    
    
       3:   
    
    
       4:  <TypeConverter(GetType(EnumTitleTypeConverter(Of TitledValues)))> _
    
    
       5:  Public Enum TitledValues
    
    
       6:      <EnumTitle(Title:="First Name")> _
    
    
       7:      FirstName
    
    
       8:      <EnumTitle(Title:="Last Name")> _
    
    
       9:      LastName
    
    
      10:      <EnumTitle(Title:="Address 1")> _
    
    
      11:      Address1
    
    
      12:      <EnumTitle(Title:="Address 2")> _
    
    
      13:      Address2
    
    
      14:      City
    
    
      15:      State
    
    
      16:      <EnumTitle(Title:="Zip Code")> _
    
    
      17:      ZipCode
    
    
      18:      <EnumTitle(Title:="Phone Number")> _
    
    
      19:      PhoneNumber
    
    
      20:      <EnumTitle(Title:="Email Address")> _
    
    
      21:      EmailAddress
    
    
      22:  End Enum
    
    
      23:   
    
    
      24:  <AttributeUsage(AttributeTargets.Field)> _
    
    
      25:  Public Class EnumTitleAttribute
    
    
      26:      Inherits Attribute
    
    
      27:   
    
    
      28:      Private m_Title As String = String.Empty
    
    
      29:      Public Property Title() As String
    
    
      30:          Get
    
    
      31:              Return m_Title
    
    
      32:          End Get
    
    
      33:          Set(ByVal value As String)
    
    
      34:              m_Title = value
    
    
      35:          End Set
    
    
      36:      End Property
    
    
      37:   
    
    
      38:  End Class
    
    
      39:   
    
    
      40:  Public Class EnumTitleTypeConverter(Of T)
    
    
      41:      Inherits TypeConverter
    
    
      42:   
    
    
      43:      Public Overrides Function ConvertTo(ByVal context As ITypeDescriptorContext, ByVal culture As CultureInfo, ByVal value As Object, ByVal destinationType As System.Type) As Object
    
    
      44:          If value.GetType Is GetType(T) AndAlso destinationType Is GetType(String) Then
    
    
      45:              If value.GetType.GetField(value.ToString).GetCustomAttributes(GetType(EnumTitleAttribute), True).Length > 0 Then
    
    
      46:                  Return DirectCast(value.GetType.GetField(value.ToString).GetCustomAttributes(GetType(EnumTitleAttribute), True)(0), EnumTitleAttribute).Title
    
    
      47:              Else
    
    
      48:                  Return value.ToString
    
    
      49:              End If
    
    
      50:          Else
    
    
      51:              Return MyBase.ConvertTo(context, culture, value, destinationType)
    
    
      52:          End If
    
    
      53:      End Function
    
    
      54:  End Class

 

 

***Edit:** Added C# Code below
    
    
       1:  using System;
    
    
       2:  using System.Collections.Generic;
    
    
       3:  using System.Text;
    
    
       4:  using System.ComponentModel;
    
    
       5:  using System.Globalization;
    
    
       6:   
    
    
       7:  [TypeConverter(typeof(EnumTitleTypeConverter<TitledValues>))]
    
    
       8:  enum TitledValues
    
    
       9:  {
    
    
      10:      [EnumTitle(Title="First Name")]
    
    
      11:      FirstName,
    
    
      12:      [EnumTitle(Title = "Last Name")]
    
    
      13:      LastName,
    
    
      14:      [EnumTitle(Title = "Address 1")]
    
    
      15:      Address1,
    
    
      16:      [EnumTitle(Title = "Address 2")]
    
    
      17:      Address2,
    
    
      18:      City,
    
    
      19:      State,
    
    
      20:      [EnumTitle(Title = "Zip Code")]
    
    
      21:      ZipCode,
    
    
      22:      [EnumTitle(Title = "Phone Number")]
    
    
      23:      PhoneNumber,
    
    
      24:      [EnumTitle(Title = "Email Address")]
    
    
      25:      EmailAddress
    
    
      26:  }
    
    
      27:   
    
    
      28:  [AttributeUsage(AttributeTargets.Field)]
    
    
      29:  public class EnumTitleAttribute : Attribute
    
    
      30:  {
    
    
      31:      private string m_Title = string.Empty;
    
    
      32:      public string Title
    
    
      33:      {
    
    
      34:          get { return m_Title; }
    
    
      35:          set { m_Title = value; }
    
    
      36:      }
    
    
      37:  }
    
    
      38:   
    
    
      39:  public class EnumTitleTypeConverter<T> : TypeConverter
    
    
      40:  {
    
    
      41:      public override object ConvertTo(ITypeDescriptorContext context, CultureInfo culture, object value, Type destinationType)
    
    
      42:      {
    
    
      43:          if ((value.GetType() == typeof(T)) && (destinationType == typeof(String)))
    
    
      44:          {
    
    
      45:              if (value.GetType().GetField(value.ToString()).GetCustomAttributes(typeof(EnumTitleAttribute), true).Length > 0)
    
    
      46:              {
    
    
      47:                  return (value.GetType().GetField(value.ToString()).GetCustomAttributes(typeof(EnumTitleAttribute),true)[0] as EnumTitleAttribute).Title;
    
    
      48:              }
    
    
      49:              else
    
    
      50:              {
    
    
      51:                  return value.ToString();
    
    
      52:              }
    
    
      53:          } 
    
    
      54:          else 
    
    
      55:          {
    
    
      56:              return base.ConvertTo(context, culture, value, destinationType);
    
    
      57:          }
    
    
      58:      }
    
    
      59:  }

![][3]

[1]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/ComboBoxEnumerationswithTitles_C75C/image_thumb.png
[2]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/ComboBoxEnumerationswithTitles_C75C/image_thumb_1.png
[3]: http://renevo.com/aggbug.aspx?PostID=1980

