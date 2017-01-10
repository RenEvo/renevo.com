---
title: A bad approach to adding RemoveAt to a Queue<T>
published: true
date: 2009-02-10 06:48:00 +0000 UTC
tags: imported .net framework performance
---
Tonight on [Stackoverflow][1] a gentlemen asked [if it would be possible to add an extension method to the Queue<T> class to remove an item by index][2]. Anyone who has worked with this class knows that there is no Add or Remove methods, but instead a Dequeue and Enqueue pair of methods. Well, the problem with Dequeue is that it only removes an item from the end of the queue, so that's not really going to work for this guy.

> "I thought of a quick and dirty solution: Inherit from the Queue class and add your own RemoveAt method."

Hey! Bloody brilliant, right? Wrong. The catch 22 is that there is no way to remove an item from the queue class itself except from the Dequeue method, which simply doesn't cut it. So to get around this I had to enumerate through all the items, add each item to a new queue, skipping the index we want to remove. Well, this is a neat solution, however it creates a new queue everytime you need to remove an item by index.

It doesn't sound so bad, but consider the code:

> **Custom Queue<T> class**

> ![queuea][3]
> 
> **Test Code in Main()**

> ![queueb][4]

 

I knew as soon as I wrote this that it would be a significant performance hit, but I didn't realize just how bad until I looked a bit deeper. Essentially this means that each time RemoveAt is called, the queue is enumerated, and a new queue is created. If we have a queue with 5000 items, and we needed to remove 2500 of them, this would result in 12,502,500 calls to the Count property, and 12,410,000 calls to the Enqueue method. A lot, much? I didn't get this data out of thin air though, I used Visual Studios performance tools, so let's take a look at the actual benchmarks:

 

![perfs][5]

 

Ouch! Approximately 9 minutes to remove 2500 items from a queue with 5000 items.. That is definitely not performance savvy. So I ran a comparison using a generic list. 

 

 

![perfs2][6]

 

4.36 milliseconds… What a time difference! So you can see that a small "innocent" "work-around" can be absolutely detrimental to your applications performance. Oops! Luckily I thought ahead of time and told the guy someone else would probably post a better solution, and it was as simple as [recommending a generic List<T>][7]. :P

![][8]

[1]: http://www.stackoverflow.com/
[2]: http://stackoverflow.com/questions/531191/c-adding-a-removeint-index-method-to-the-net-queue-class
[3]: http://renevo.com/blogs/developer/queuea_thumb_584D6EB2.png "queuea"
[4]: http://renevo.com/blogs/developer/queueb_thumb_0BA9320F.png "queueb"
[5]: http://renevo.com/blogs/developer/perfs_thumb_256D4271.png "perfs"
[6]: http://renevo.com/blogs/developer/perfs2_thumb_418AD4D1.png "perfs2"
[7]: http://stackoverflow.com/questions/531191/c-adding-a-removeint-index-method-to-the-net-queue-class/531231#531231
[8]: http://renevo.com/aggbug.aspx?PostID=2154

