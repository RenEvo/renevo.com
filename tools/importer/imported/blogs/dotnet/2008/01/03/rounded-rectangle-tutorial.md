---
title: Rounded Rectangle Tutorial
published: true
date: 2008-01-03 01:18:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2008/01/02/rounded-rectangle-tutorial.aspx
file: rounded-rectangle-tutorial.aspx
path: /blogs/dotnet/archive/2008/01/02/
author: tom anderson
words: 816
---
After seeing a lot of samples on the web, and a lot of different methods, I have decided to write up my own little tutorial on getting Rounded Rectangles in .Net

First off, lets start by opening up Visual Studio and creating a new Windows Forms Application.

Next, resize your main form and add a new picture box to the form. Resize the picture box to 400,400 and adjust your form size to where it looks something like this.

![clip_image001][1]

Double click on the form to access the Form Load event handler, and lets start some small bits of code.

Instead of just showing you how to do this, I am going to first demonstrate exactly how each part of what we are doing actually works.

Start off by creating a new Bitmap object, I decreased the size a bit to allow for the borders and padding of the picture box.

![clip_image002][2]

Now lets setup our picture box and display the new bitmap in it at runtime.

![clip_image003][3]

Next, we are going to create a new Graphics object (System.Drawing.Graphics).

![clip_image004][4]

You may also want to set the smoothing mode of the Graphics object, this will produce much cleaner lines while drawing shapes.

![clip_image005][5]

With the graphics object you can draw and fill all kinds of wonderful shapes and colors.  We are going to first concern ourselves with the Graphics.DrawArc method.

To get a good feel, lets first start with the obvious, draw a single arc on the image.

![clip_image006][6]

The method above is drawing an arc with a black pen at position 10,10 in a 10x10 size a start angle of 0 degrees, and an sweep angle of 90 degrees, which is a right triangle.

If you run the application, you should see something like this.

![clip_image007][7]

Not very impressive, but a good start to the rounded rectangle we are going to want to create.  Working around in quarter circles, we can then display all 4 sides of the circle.

![clip_image008][8]

When you run the application now, you should see a complete circle like so

![clip_image009][9]

Neat huh? Comment out the lines one at a time to see which angle is which, just to save you some time, here are the mappings for our corners.

Starting Angle 0 = Bottom Right

Starting Angle 90 = Bottom Left

Starting Angle 180 = Top Left

Starting Angle 270 = Top Right

That is a very basic understanding of the Arc drawing, lets make a rounded rectangle now!

Stub out the following method

![clip_image010][10]

This method will create a GraphicsPath Object that defines the shape of our rounded rectangle.  A graphics path can best be described as a way to draw piece by piece the lines, points, etc… of a shape, and then "fill in the blanks" by connecting all the points.  A lot of samples I have seen, and used, usually draw all the internal lines, or create a bunch of boxes, etc…  But in this tutorial, we are going to do it the "easy way".

Lets first take care of a small issue that I have noticed with this method, For some reason the right side and bottom always gets clipped off, so I manually add a padding of 1,1,2,2; that's top 1, left 1, right 2, bottom 2.  This will center up the rounded rectangle quite nicely.  And as you noticed in the constructor, I added a way to put in a padding for the actual drawing so that you can grow shrink it using the padding.

![clip_image011][11]

That should be pretty self explanatory.

Now, using the logic that we built above, we are going to create arcs at each of the four corners of the image, except using the GraphicsPath.AddArc method (same as the Graphics.DrawArc, except it doesn't contain a color in the arguments).

![clip_image012][12]

Following the code, you will see that we start in the upper left hand corner to create an arc, move to the upper right hand corner, bottom right hand corner, then finally the bottom left hand corner, now we will want to take a big shortcut

![clip_image013][13]

Here we tell the GraphicsPath to "connect the dots", basically it takes the current path and fills in the blanks to create a solid shape.

With the method complete, it should look like this

![clip_image014][14]

Only thing left to do is make a call to this method to get a path, and then use the Graphics object to "fill the path" - back in our form load under the smoothing mode set

![clip_image015][15]

Big call, simple understanding.  Call the Graphics.FillPath, set the brush color, then pass in the return from our call to MakeRoundedRectanglePath that we created above.

Now run the application, and look at your rounded rectangle goodness.

![clip_image016][16]

Now, lets dress this up a bit, to make it a bit more attractive. Use a System.Drawing.Drawing2D.LinearGradientBrush instead of a standard SolidBrush.

![clip_image017][17]

And now our newly applied fill path method

![clip_image018][18]

And, our final result:

![clip_image019][19]

Pretty simple!  I hope you enjoyed this tutorial, below is the full code for this tutorial.  With some practice and modifications of the function calls, you can very easily get this result

![clip_image020][20]

Source

![clip_image021][21]

*Yes, they are images, you should type to learn.



[1]: ./rounded-rectangle-tutorial/clip_image001_thumb.png
[2]: ./rounded-rectangle-tutorial/clip_image002_thumb.png
[3]: ./rounded-rectangle-tutorial/clip_image003_thumb.png
[4]: ./rounded-rectangle-tutorial/clip_image004_thumb.png
[5]: ./rounded-rectangle-tutorial/clip_image005_thumb.png
[6]: ./rounded-rectangle-tutorial/clip_image006_thumb.png
[7]: ./rounded-rectangle-tutorial/clip_image007_thumb.png
[8]: ./rounded-rectangle-tutorial/clip_image008_thumb.png
[9]: ./rounded-rectangle-tutorial/clip_image009_thumb.png
[10]: ./rounded-rectangle-tutorial/clip_image010_thumb.png
[11]: ./rounded-rectangle-tutorial/clip_image011_thumb.png
[12]: ./rounded-rectangle-tutorial/clip_image012_thumb.png
[13]: ./rounded-rectangle-tutorial/clip_image013_thumb.png
[14]: ./rounded-rectangle-tutorial/clip_image014_thumb.png
[15]: ./rounded-rectangle-tutorial/clip_image015_thumb.png
[16]: ./rounded-rectangle-tutorial/clip_image016_thumb.png
[17]: ./rounded-rectangle-tutorial/clip_image017_thumb.png
[18]: ./rounded-rectangle-tutorial/clip_image018_thumb.png
[19]: ./rounded-rectangle-tutorial/clip_image019_thumb.png
[20]: ./rounded-rectangle-tutorial/clip_image020_thumb.png
[21]: ./rounded-rectangle-tutorial/clip_image021_thumb.png


