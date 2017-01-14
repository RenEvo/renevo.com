---
title: New code snippet plugin for Live Writer
published: true
date: 2008-12-04 20:39:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2008/12/04/new-code-snippet-plugin-for-live-writer.aspx
file: new-code-snippet-plugin-for-live-writer.aspx
path: /blogs/developer/archive/2008/12/04/
author: tom anderson
words: 228
---
I am trying out a new [code snippet plug-in][1] for [Live Writer][2].
    
    
    Public Class CommitDB
        Public Function GetCommitStatusAll() As DataSet
            ' Create Instance of Connection and Command Object
            Dim myConnection As New SqlConnection(ConfigurationManager.AppSettings("NorthstarConnectionString"))
            Dim myCommand As New SqlDataAdapter("GetCommitStatusAll", myConnection)
    
            ' Mark the Command as a SPROC
            myCommand.SelectCommand.CommandType = CommandType.StoredProcedure
    
            ' Create and Fill the DataSet
            Dim myDataSet As New DataSet
            myCommand.Fill(myDataSet)
    
            ' Return the DataSet
            Return myDataSet
        End Function
    End Class

Lets see how it looks on the blogs!![][3]

[1]: http://gallery.live.com/liveItemDetail.aspx?li=d8835a5e-28da-4242-82eb-e1a006b083b9&bt=9&pl=8
[2]: http://windowslivewriter.spaces.live.com/
[3]: http://renevo.com/aggbug.aspx?PostID=2101

