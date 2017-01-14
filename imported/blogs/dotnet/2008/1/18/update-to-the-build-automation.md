---
title: Update to the Build Automation
published: true
date: 2008-01-18 00:52:00 +0000 UTC
tags: imported 
original: http://renevo.com/blogs/dotnet/archive/2008/01/17/update-to-the-build-automation.aspx
file: update-to-the-build-automation.aspx
path: /blogs/dotnet/archive/2008/01/17/
author: tom anderson
words: 1270
---
Well, after a few months working on this off and on, I thought I would share the current status of our operations.

First and foremost, we got Cruise Control up and running, this made a huge step forward with the process, and it is open source, so that allowed me to create a custom UI for it to run.  This resulted in us having a custom application that runs very similarly to the Cruise Control Tray Application (CCTray), except I implemented grouping by categories, displaying the categories, as well as auto-discovery of projects, rather then having to set them up.

Additionally, because we don't need everyone in our department building every piece of software all the time, I created a pseudo permission system that is bound into an xml file, like the example below.

<?xml version="1.0" encoding="utf-8" standalone="yes" ?>   
<security>   
    <add user="tanderson" viewConfig="true" projects="*" />   
    <add user="jsmith" viewConfig="false" projects="RecipeEditor_VB6;Rocket;QReports;" />   
</security>

As you can see, it is a very simplistic permissions system, the user is the currently logged in computer user, viewConfig is a small permission to see if they can view the servers ccnet.config, and the projects are a semi-colon separated list of the names of the projects they can see status of, or for ease of people like me, an asterisk for all projects.

Below is a screen shot of the current version of the software.

![image][1]

I have made quite a bit of headway with the system in getting more and more projects completed and into the daily build.  The "DailyBuild" is the only project that is built nightly, and it builds everyone of the other projects.

The actual process of creating all of this mayhem was a few steps.

**Creating an Environment**

The first thing I had to do, was get an environment going that was as clean as possible to do the builds, this is a remote connection only off the domain standalone box that only myself and the VP of software development have access to.  The computer is loaded with Server 2003 Standard Edition, SQL 2000, SQL 2005, Visual Basic 6, VS 2003, VS 2005, and RedGate toolkit.

**Starting the batches**

When someone told me I would do a crap load of batch files as a quick solution to get a process done, I would have laughed at them, but in this case, it seemed to be the quickest route to the finish line. In total for what you see above, I have roughly 38 batch files.

The first ones that I built, where simplistic, setup the compiling environments, and then get the latest code from source control. I created two separate batch files for this to ease the calling of them from other batch files.

_SetupEnvironment.cmd_

@rem Setup all variables here   
@SET SSCMLOGIN=user:password   
@SET SSCMSERVER=server:4900   
@SET CBSSoftware="C:BuildCBS Software"   
@SET CBSLegacy="C:BuildCBS Legacy"   
@SET CBSProto="C:BuildCBS Prototype" 

@SET VB6EXE="C:Program FilesMicrosoft Visual StudioVB98vb6.exe"   
@SET VBCEXE="C:windowsmicrosoft.netframeworkv2.0.50727vbc.exe"   
@SET VS2003="c:program filesMicrosoft Visual Studio .NET 2003common7idedevenv.exe"   
@SET VS2005="C:Program filesmicrosoft visual studio 8common7idedevenv.exe"   
@SET VS2008="C:Program Filesmicrosoft visual studio 9.0common7idedevenv.exe" 

@rem Build utilities here   
@call Utilities/BuildUtilities 

@EXIT /B %ERRORLEVEL%

I added a utilities folder to store a lot of extra files that I didn't want in the main batch folder, such as:

* Icons 
* Sign Keys 
* Winrar.exe 
* XCopy exclusion lists for different types of software 
* Cruise Control Log Files 

I also added a BuildUtilities batch file that allows me to compile and use stand alone class files in .Net, such as one that I use very frequently, the List.exe.

_Utilities/BuildUtilities.cmd_

@rem add any utility compiles here   
@echo Building Utilities 

@rem Building List   
@"%VBCEXE%" /target:exe /nologo "C:BuildUtilitieslist.vb" 

@EXIT /B %ERRORLEVEL%

List.exe is a very simple executable that basically outputs a file from disc to the console.

The next batch file that I built was to pull all of the code for our products from source control.

_GetSource.cmd_

@rem Retrieve all source   
@echo Retrieving latest source code   
@mkdir SSoftware%   
@mkdir SLegacy%   
@mkdir SProto% 

@sscm workdir ""SSoftware%"" "CBS Software" -y"%SSCMLOGIN%" -z"%SSCMSERVER%"   
@sscm get *.* -d""SSoftware%"" -p"CBS Software" -r -e -q -wreplace -y"%SSCMLOGIN%" -z"%SSCMSERVER%" 

@sscm workdir ""SLegacy%"" "CBS Legacy" -y"%SSCMLOGIN%" -z"%SSCMSERVER%"   
@sscm get *.* -d""SLegacy%"" -p"CBS Legacy" -r -e -q -wreplace -y"%SSCMLOGIN%" -z"%SSCMSERVER%" 

@sscm workdir ""SProto%"" "CBS Prototype" -y"%SSCMLOGIN%" -z"%SSCMSERVER%"   
@sscm get *.* -d""SProto%"" -p"CBS Prototype" -r -e -q -wreplace -y"%SSCMLOGIN%" -z"%SSCMSERVER%" 

@EXIT /B %ERRORLEVEL%

We use Surround SCM from Seapine software, so your command lines may differ quite a bit.

**Creating a compile project**

To do anything, I needed to be able to compile the projects, this was done mostly via command line to the proper visual studio, I wanted to do the direct calls to the compilers, but there was entirely way too much legwork that would have had to be done first.

Although, in order to make my giant "DailyBuild" batch file as well as a by-product batch file, I needed to separate some of the logic from the build to only do certain things, for instance, I don't want to get the source for all projects each time I build one project.  Instead I created two files, one with the product name "Application.cmd" and another with "BuildApplication.cmd".  The Build*.cmd sets up the environment, gets the source, then calls the application.cmd batch at the end.

_Iconifier.cmd_

@echo Building Iconifier   
@echo ------------------------------------------------------------------   
@echo This build requires .Net 2.0, Visual Studio 2005   
@echo ------------------------------------------------------------------ 

@rmdir /S /Q "./Output/Iconifier/"   
@mkdir "Output/Iconifier/"   
@rmdir /S /Q "./Output/Iconifier/Icon Templates"   
@mkdir "Output/Iconifier/Icon Templates/" 

@echo Building Iconifier 

@%VS2005% "./CBS Software/Desktop/Utilities/NorthStar.Desktop.Iconifier/NorthStar.Desktop.Iconifier.sln" /rebuild Release /out OutputIconifierbuild.log   
@"./Utilities/List" ./Output/Iconifier/Build.log   
@del outputIconifierbuild.log   
@IF NOT EXIST "./CBS Software/Desktop/Utilities/NorthStar.Desktop.Iconifier/NorthStar.Desktop.Iconifier/bin/release/northstar.Desktop.Iconifier.exe" EXIT /B 1 

@rem copy over all the files   
@xcopy "CBS SoftwareDesktopUtilitiesNorthStar.Desktop.IconifierNorthStar.Desktop.Iconifierbinrelease*.exe" OutputIconifier /S /Q /Y   
@xcopy "CBS SoftwareDesktopUtilitiesNorthStar.Desktop.IconifierNorthStar.Desktop.IconifierIcon Templates*.png" "OutputIconifierIcon Templates" /S /Q /Y 

@rem Archive   
@"./Utilities/winrar.exe" a -r -idcdp -sfx.Utilitiesdefault.sfx -ibck -ep1 -iimg.Utilitiessfx_logo.bmp -ts -iicon".utilitiesdefault.ico" -z".UtilitiesSFX_default.txt" ".OutputNorthStar.Iconifier.Daily.exe" "OutputIconifier*.*"   
@"./Utilities/winrar.exe" a -ag-MM-DD-YYYY -r -idcdp -sfx.Utilitiesdefault.sfx -ibck -ep1 -iimg.Utilitiessfx_logo.bmp -ts -iicon".utilitiesdefault.ico" -z".UtilitiesSFX_default.txt" ".OutputNorthStar.Iconifier.Daily.exe" "OutputIconifier*.*" 

@EXIT /B %ERRORLEVEL%

Quick recap, remove existing directories (clean), create them, build the solution file, output the log to the console for logging purposes, copy all the files from the build directory to the output directory, archive using winrar (giant command lines for the win) a daily file, then archive again with a date stamped file.

_BuildIconifier.cmd_

@cls   
@rem Start Iconifier Daily Build 

@echo ------------------------------------------------------------------   
@echo Daily  Iconifier Script Initialized   
@echo Version: 1.0   
@echo Author:  Tom Anderson   
@echo ------------------------------------------------------------------ 

@call SetupEnvironment   
@IF NOT %ERRORLEVEL%  == 0 EXIT /B %ERRORLEVEL% 

@call GetSource   
@IF NOT %ERRORLEVEL%  == 0 EXIT /B %ERRORLEVEL% 

@echo ------------------------------------------------------------------ 

@call Iconifier   
@IF NOT %ERRORLEVEL%  == 0 EXIT /B %ERRORLEVEL% 

@echo ------------------------------------------------------------------   
@echo Deploying to <http://servername/Downloads/NorthStar.Iconifier.Daily.exe>   
@copy OutputNorthStar.Iconifier.Daily.exe C:InetpubwwwrootDownloads /Y   
@echo ------------------------------------------------------------------ 

@ECHO ERROR LEVEL: %ERRORLEVEL% 

@EXIT /B %ERRORLEVEL%

And here is our setup and build batch file, which can be called stand alone.  Only thing unique is that I copy the daily.exe to the web server hosted on this machine for download ability within the company.

The project pasted above is actually a pretty quick and simple one, some of our projects have upwards of 12 compiles to make the final product, and quite a few supporting files.

**Wrap up**

It is still a work in progress, and I still have about 6 projects to get in that are the most complex, and a database to create a daily build off of, but the progress is going pretty well.

![][2]

[1]: http://www.renevo.com/blogs/dotnet/WindowsLiveWriter/UpdatetotheBuildAutomation_FC2A/image_thumb.png
[2]: http://renevo.com/aggbug.aspx?PostID=1698

