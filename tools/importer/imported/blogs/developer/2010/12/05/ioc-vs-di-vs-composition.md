---
title: IOC vs. DI vs. Composition
published: true
date: 2010-12-05 07:57:51 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/developer/archive/2010/12/05/ioc-vs-di-vs-composition.aspx
file: ioc-vs-di-vs-composition.aspx
path: /blogs/developer/archive/2010/12/05/
author: tom anderson
words: 1710
---
First, lets define some terms in the basics:

**IOC**: [Inversion of Control][1]   
IOC is quite simply, handing off responsibility to another system based on a known contract.

**DI**: [Dependency Injection][2]   
DI is the act of, for lack of a better term, injecting those known contracts at runtime.

**Composite**:   
[Composition][3] is the taking many of those known contracts, and acting on them all as one.

As the years go by, it comes to me that more and more, these terms get mixed together, without any clarification of what they each really do. To help illustrate the differences in how these work, I am going to use a known system that utilizes all three of these. Hopefully at the end of this little sample, you will see what pattern/practice you really need to be using.

**Asp.Net Membership.**

Anyone who has ever done a website should have known about this specific static method call:

       1:  Membership.GetUser();

This is a perfect example of Inversion of Control.  Behind the scenes, you are simply talking to a known contract (MembershipProvider) that then implements the abstract method "GetUser". Instead of accessing the helper method "GetUser" you could talk directly to the default provider to validate a user:

       1:  bool validUser = Membership.Provider.ValidateUser("username", "password");

This is no different, and no matter what [MembershipProvider][4] is being implemented, you know that this will return an expected result.

This then begs the question, how does the Membership.Provider get populated with a MembershipProvider? This is where a minor version of dependency injection occurs.

Deep in the delves of your machine.config, there is a setting that says the default Membership Provider is named "AspNetSqlMembershipProvider" and is of type: "[System.Web.Security.SqlMembershipProvider][5]". If for any chance you went in your web.config and in the system.web/membership/provider collection and did a <clear /> and attempted to call the Membership.GetUser, it would throw an exception stating that there isn't a default Membership Provider. In this case, the Membership static class is acting as a DI container for the Provider.

From [machine.config][6] (C:windowsmicrosoft.netframeworkv4.0.30319Configmachine.config):

       1:          <membership>    
       2:              <providers>    
       3:                  <add name="AspNetSqlMembershipProvider" type="System.Web.Security.SqlMembershipProvider    
       4:              </providers>    
       5:          </membership>

Now comes my favorite, Composition. 

The [Membership][7] static class provides not just DI and IOC, but it also provides for Composition, which is taking multiple implementations of a single contract and treating them as one. To take advantage of the Composite implementation of MembershipProvider you would use this type of code:

       1:              bool validUser = false;    
       2:              foreach (MembershipProvider provider in Membership.Providers)    
       3:                  try    
       4:                  {    
       5:                      if (provider.ValidateUser("username", "password"))    
       6:                          validUser = true;    
       7:                  }    
       8:                  catch (Exception)    
       9:                  {    
      10:                      // TODO: Log Here    
      11:                  }

Essentially, we are asking ALL of the membership providers if they can validate this user, if so awesome, if not the user is not valid.

*Adding a new membership provider through DI for composition/IOC

       1:      <membership defaultProvider="MyProvider">    
       2:        <providers>    
       3:          <add name="MyProvider" type="MyNamespace.MyProvider"/>    
       4:        </providers>    
       5:      </membership>

The above sample would now attempt to call the AspNetSqlMembershipProvider as well as MyProvider to validate a user, while the Membership.ValidateUser would only call MyProvider.

***

For this article I just wanted to take a few minutes to briefly touch on the subtle differences between these patterns/practices, as well as clearly separating out what IOC was. I have seen in many many cases where DI was used for Composition (Unfortunately most frameworks mix these two via Resolve<> and ResolveAll<>) or IOC.

***

If you are just trying to make your code testable, use IOC and provide the required contracts in your constructor, with an optional default constructor that assembles them from working implementations.

       1:       
       2:          public class RequiredInterface : IRequiredInterface    
       3:          { }    
       4:       
       5:          public interface IRequiredInterface    
       6:          {    
       7:       
       8:          }    
       9:       
      10:          public class RequiredInterfaceConsumer    
      11:          {    
      12:              private IRequiredInterface _required;    
      13:       
      14:              public RequiredInterfaceConsumer()    
      15:              {    
      16:                  _required = new RequiredInterface();    
      17:              }    
      18:       
      19:              public RequiredInterfaceConsumer(IRequiredInterface required)    
      20:              {    
      21:                  if (required == null)    
      22:                      throw new ArgumentNullException("required");    
      23:       
      24:                  _required = required;    
      25:              }    
      26:          }

This is all that is required to make a class testable with dependencies, the IRequiredInterface can be mocked using any type of manual process, or mocking framework.

Only use DI where you expect frequent enough changes to the implementation of the contract. You should always prefer [orchestration][8] over DI whenever possible.

Only use Composition where it makes sense such as when you need to treat multiple objects like a single object.

***

Once again, Composition/IOC do not require DI.



[1]: http://en.wikipedia.org/wiki/Inversion_of_control
[2]: http://en.wikipedia.org/wiki/Dependency_injection
[3]: http://en.wikipedia.org/wiki/Composite_pattern
[4]: http://msdn.microsoft.com/en-us/library/system.web.security.membershipprovider.aspx
[5]: http://msdn.microsoft.com/en-us/library/system.web.security.sqlmembershipprovider.aspx
[6]: http://stackoverflow.com/questions/2325473/where-is-machine-config
[7]: http://msdn.microsoft.com/en-us/library/system.web.security.membership.aspx
[8]: http://en.wikipedia.org/wiki/Orchestration_(computers)


