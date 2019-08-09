---
title: Reason for the Zune 30gb Bug
published: true
date: 2009-01-02 18:17:04 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2009/01/02/reason-for-the-zune-30gb-bug.aspx
file: reason-for-the-zune-30gb-bug.aspx
path: /blogs/developer/archive/2009/01/02/
author: tom anderson
words: 1098
---
Well, the [source code has been found][1] that caused the [December 31st Zune Brick][2].

       1:  //------------------------------------------------------------------------------    
       2:  //    
       3:  // Function: ConvertDays    
       4:  //    
       5:  // Local helper function that split total days since Jan 1, ORIGINYEAR into     
       6:  // year, month and day    
       7:  //    
       8:  // Parameters:    
       9:  //    
      10:  // Returns:    
      11:  //      Returns TRUE if successful, otherwise returns FALSE.    
      12:  //    
      13:  //------------------------------------------------------------------------------    
      14:  BOOL ConvertDays(UINT32 days, SYSTEMTIME* lpTime)    
      15:  {    
      16:      int dayofweek, month, year;    
      17:      UINT8 *month_tab;    
      18:       
      19:      //Calculate current day of the week    
      20:      dayofweek = GetDayOfWeek(days);    
      21:       
      22:      year = ORIGINYEAR;    
      23:       
      24:      while (days > 365)    
      25:      {    
      26:          if (IsLeapYear(year))    
      27:          {    
      28:              **__if (days > 366)__**    
      29:              {    
      30:                  days -= 366;    
      31:                  year  = 1;    
      32:              }    
      33:          }    
      34:          else    
      35:          {    
      36:              days -= 365;    
      37:              year  = 1;    
      38:          }    
      39:      }    
      40:       
      41:       
      42:      // Determine whether it is a leap year    
      43:      month_tab = (UINT8 *)((IsLeapYear(year))? monthtable_leap : monthtable);    
      44:       
      45:      for (month=0; month<12; month  )    
      46:      {    
      47:          if (days <= month_tab[month])    
      48:              break;    
      49:          days -= month_tab[month];    
      50:      }    
      51:       
      52:      month  = 1;    
      53:       
      54:      lpTime->wDay = days;    
      55:      lpTime->wDayOfWeek = dayofweek;    
      56:      lpTime->wMonth = month;    
      57:      lpTime->wYear = year;    
      58:       
      59:      return TRUE

If you didn't catch it, December 31st was day 366, therefore the there was never any code in what to do if days == 366.  Woops…



[1]: http://pastie.org/349916
[2]: http://www.associatedcontent.com/article/1350870/zune_locked_up_and_frozen_happy_new.html


