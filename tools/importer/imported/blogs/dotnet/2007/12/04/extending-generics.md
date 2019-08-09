---
title: Extending Generics
published: true
date: 2007-12-04 19:23:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2007/12/04/extending-generics.aspx
file: extending-generics.aspx
path: /blogs/dotnet/archive/2007/12/04/
author: tom anderson
words: 2905
---
I am a huge fan of generics, but sometimes you just don't have that extra functionality that you need.  After using the CAB framework, I have found a few additions that I make to the root Generics that make using them a lot better.

First, I will inherit a System.Collections.Generic.List(of T) and create my own generic, from there I will add a few constructors, and Add with a return value, and AddNew method, and IndexOf, and an Item overload.  This will allow me to use this generic, specify an "IndexProperty" that will allow me to say, find an object in the list by the property "Name", "Text", or whatever else my heart desires.

Below is the code for the actual new class with comments.  Notice this is in the My namespace so will only be available from within your assembly, unless you change the namespace yourself.

       1:  Imports System.ComponentModel    
       2:       
       3:  Namespace My.Generics    
       4:       
       5:      ''' <summary>    
       6:      ''' Extension class for the Generic.List Object    
       7:      ''' </summary>    
       8:      Public Class List(Of T)    
       9:          Inherits Generic.List(Of T)    
      10:       
      11:          ''' <summary>    
      12:          ''' Default constructor    
      13:          ''' </summary>    
      14:          Public Sub New()    
      15:          End Sub    
      16:       
      17:          ''' <summary>    
      18:          ''' Constructor to set the IndexProperty    
      19:          ''' </summary>    
      20:          Public Sub New(ByVal IndexProperty As String)    
      21:              Me.IndexProperty = IndexProperty    
      22:          End Sub    
      23:       
      24:          Private m_IndexProperty As String = String.Empty    
      25:          ''' <summary>    
      26:          ''' Gets or Sets the property name to Index on    
      27:          ''' </summary>    
      28:          Public Property IndexProperty() As String    
      29:              Get    
      30:                  Return m_IndexProperty    
      31:              End Get    
      32:              Set(ByVal value As String)    
      33:                  m_IndexProperty = value    
      34:              End Set    
      35:          End Property    
      36:       
      37:          ''' <summary>    
      38:          ''' Returns the newly added item    
      39:          ''' </summary>    
      40:          Public Shadows Function Add(ByVal Item As T) As T    
      41:              MyBase.Add(Item)    
      42:              Return Item    
      43:          End Function    
      44:       
      45:          ''' <summary>    
      46:          ''' Uses the Activator Class to create a new instance of the object    
      47:          ''' </summary>    
      48:          Public Function AddNew() As T    
      49:              Dim newItem As T = Activator.CreateInstance(GetType(T))    
      50:              Return Add(newItem)    
      51:          End Function    
      52:       
      53:          ''' <summary>    
      54:          ''' Returns the index of an object based on the IndexProperty    
      55:          ''' </summary>    
      56:          Public Overloads Function IndexOf(ByVal Value As Object) As Integer    
      57:              Dim retVal As Integer = -1    
      58:       
      59:              If m_IndexProperty.Length > 0 Then    
      60:                  For Each obj As T In Me    
      61:                      If obj.GetType.GetProperty(m_IndexProperty) IsNot Nothing Then    
      62:                          If obj.GetType.GetProperty(m_IndexProperty).GetValue(obj, Nothing) Is Value Then    
      63:                              retVal = IndexOf(obj)    
      64:                              Exit For    
      65:                          End If    
      66:                      End If    
      67:                  Next    
      68:              End If    
      69:       
      70:              Return retVal    
      71:          End Function    
      72:       
      73:          ''' <summary>    
      74:          ''' Returns an object based on the IndexProperty    
      75:          ''' </summary>    
      76:          Default Public Overloads Property Item(ByVal PropertyValue As Object) As T    
      77:              Get    
      78:                  Return Item(IndexOf(PropertyValue))    
      79:              End Get    
      80:              Set(ByVal value As T)    
      81:                  If IndexOf(value) > -1 Then    
      82:                      Item(IndexOf(value)) = value    
      83:                  Else    
      84:                      Add(value)    
      85:                  End If    
      86:              End Set    
      87:          End Property    
      88:       
      89:          ''' <summary>    
      90:          ''' Removes an object based on the IndexProperty    
      91:          ''' </summary>    
      92:          Public Overloads Sub Remove(ByVal Value As Object)    
      93:              Remove(Item(Value))    
      94:          End Sub    
      95:       
      96:      End Class    
      97:       
      98:  End Namespace

As you can see, there isn't really a lot of extra coding in there to get a LOT of functionality.  Below is a sample implementation class for the new List object.

       1:      Public Class NamedFormCollection    
       2:          Inherits My.Generics.List(Of Form)    
       3:       
       4:          Public Sub New()    
       5:              MyBase.New("Name")    
       6:          End Sub    
       7:       
       8:      End Class

This allows me to now do the following, which i could never have done with a simple List(of Form)

       1:  Public Class MainForm    
       2:       
       3:      Private Sub MainForm_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load    
       4:          Dim nfc As New NamedFormCollection()    
       5:          With nfc.AddNew()    
       6:              .Name = "Hello"    
       7:              .Text = "Hello Form"    
       8:          End With    
       9:       
      10:          nfc.Item("Hello").ShowDialog(Me)    
      11:       
      12:      End Sub    
      13:       
      14:  End Class     
     

That to me, is way worth the extra effort of fleshing out some helper classes.





