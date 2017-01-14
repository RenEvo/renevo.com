---
title: Upcoming Site Developments
published: true
date: 2009-01-07 01:32:18 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2009/01/06/upcoming-site-developments.aspx
file: upcoming-site-developments.aspx
path: /blogs/developer/archive/2009/01/06/
author: tom anderson
words: 391
---
I have two websites that I am about to embark on doing, below is a scratch list of the technology that I will be using, as well as any additional brain dump info on them, there will be much more information coming later on these, but for now… Enjoy my tech.

#### Base Framework 

    **ASP.Net MVC**   
<http://www.asp.net/mvc/default.aspx>

    **LinqToSql   
**<http://msdn.microsoft.com/en-us/library/bb425822.aspx>

    **jQuery**   
<http://jquery.com/>   
<http://plugins.jquery.com/>

    _**Possible**_?   
<http://ui.jquery.com/>

    **WMD**   
<http://wmd-editor.com/>

    **HTML Sanitation**   
<http://refactormycode.com/codes/333-sanitize-html>   
        _Allowed Tags_   
            <a>   
                href=""   
                title=""   
            <b>   
            <blockquote>   
            <code>   
            <del>   
            <dd>   
            <dl>   
            <dt>   
            <em>   
            <h1>, <h2>, <h3>   
            <i>   
            <li>   
            <ol>   
            <p>   
            <pre>   
            <s>   
            <sup>   
            <sub>   
            <strong>   
            <strike>   
            <ul>   
            <br/>   
            <hr/>   
            <img>   
                src=""   
                width="" (up to 999)   
                height="" (up to 999)   
                alt=""   
                title=""

    **Prettify**   
<http://code.google.com/p/google-code-prettify/>

        _jQuery to Prettify   
_            function styleCode() {   
                var hascode = false;   
                $("pre code").parent().each(function() {   
                    if (!$(this).hasClass('prettyprint')) {   
                        $(this).addClass('prettyprint');   
                        hascode = true;   
                    }   
                });   
                if (hascode) { prettyPrint(); }   
            }; 

#### Authentication 

    **.Net Open Id   
**<http://www.codeplex.com/dotnetopenid>   
<http://www.hanselman.com/blog/TheWeeklySourceCode25OpenIDEdition.aspx>

    **RpxLib**   
<http://code.google.com/p/rpxlib/>   
<http://www.blechie.com/wtilton/archive/2009/01/03/Implementing-OpenId-and-more-with-ASP.NET-MVC-and-RPXNow.aspx>

#### Spam Protection 

    **ReCaptcha**   
<http://recaptcha.net/>   
<http://recaptcha.net/plugins/aspnet/>

    **Akismet**   
<http://akismet.com/>   
<http://www.codeplex.com/AkismetApi>

#### User Customization

    **Gravatar**   
<http://www.gravatar.com>   
<http://en.gravatar.com/site/implement/url>

        _Sample_:   
[http://www.gravatar.com/avatar/1822257954fd46d7ab7732998776a80a?s=128&d=identicon&r=PG][1]

        _Construct_:   
[http://www.gravatar.com/avatar/md5ofemail?s=size&d=defaulticon&r=rating][2]

        \--Probably use the wavatar 

    _Default User Stuff   
_        Name (Pre-populated from openid)   
        Email (Not displayed, used for gravatar)   
        Real Name (Pre-populated from openid)   
        Website   
        Location   
        Birthday (never displayed, used to show age)   
        About Me (Markdown)

![][3]

[1]: http://www.gravatar.com/avatar/1822257954fd46d7ab7732998776a80a?s=128&d=identicon&r=PG
[2]: http://www.gravatar.com/avatar/md5ofemail?s=size&d=defaulticon&r=rating
[3]: http://renevo.com/aggbug.aspx?PostID=2127

