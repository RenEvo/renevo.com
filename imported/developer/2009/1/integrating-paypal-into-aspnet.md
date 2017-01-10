---
title: Integrating Paypal into ASP.Net
published: true
date: 2009-01-11 22:53:00 +0000 UTC
tags: imported 
---
There's no reason to write a whole lot of backend code, workarounds, or use third party Paypal controls just to get Paypal to work with ASP.Net. Actually, it only takes two lines of code in the Page_Load event. Let's take a look.

> Form.Action = "https://www.paypal.com/cgi-bin/webscr";

> Form.Method = "post";

Using the above two lines of code in your codebehind, you can now just plop the input markup for the Buy Now, Add to Cart, or Donate buttons in your .aspx page. The reason copying and pasting the Paypal markup doesn't work in the first place, is because ASP.Net is nothing but forms. And you can't have a form inside a form. Of course using this method, the page has to be dedicated to submitting the data to Paypal, but you wouldn't normally have more than one form on a page anyway, so it doesn't really matter. 

This is the most elegant solution as far as I know. After my many hours of Googling around, I never found an easy solution until I realized I could just manipulate the main form.

**Important things to know**

You only need to use the Page_Load event if you are using MasterPages, otherwise you can just use the form tag in the individual .aspx page.

![][1]

[1]: http://renevo.com/aggbug.aspx?PostID=2132

