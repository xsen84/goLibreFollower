# goLibreFollower by Alexandru Mihai
<p align="center">
  <img width="800" height="500" src="https://github.com/xsen84/goLibreFollower/blob/main/doc/main-window.png">
</p>

- the app is meant ro be run on Windows or MacOS.
- it periodically retrieves glucose readings from the LibreLinkUp servers using a Follower account.

# Pre Requirements
- windows or mac computer
- a valid libre follower account https://librelinkup.com/
  
# Installation

## Windows
- download `goLibreFollower-WINDOWS_[VERSION].zip` from here: https://github.com/xsen84/goLibreFollower/releases
- extract/unzip zip archive and run the program
- on first run the settings window will open. See [Settings](https://github.com/xsen84/goLibreFollower/edit/main/readme.md#settings) for how to configure the application.


## Mac OSX
- download `goLibreFollower-MAC_[VERSION].zip` from here: https://github.com/xsen84/goLibreFollower/releases
- extract/unzip zip archive. Assuming you are in `Downloads` open a terminal and run:
```
cd ~/Downloads/
xattr -cr goLibreFollower-MAC.app
chmod +x goLibreFollower-MAC.app/Contents/MacOS/goLibreFollower
```
- if the operating system is blocking the app, please allow an exception acording to this article : https://support.apple.com/en-en/guide/mac-help/mh40616/mac 
- on first run the settings window will open. See [Settings](https://github.com/xsen84/goLibreFollower/edit/main/readme.md#settings) for how to configure the application.


# Settings
- SETTINGS ARE NOT APPLIED IN REAL TIME - settings will be applied after close/open (restart) of the application
- on first run, a settings window will be opened. You will need to provide following info:
<p align="center">
  <img width="600" height="350" src="https://github.com/xsen84/goLibreFollower/blob/main/doc/settings-window.png">
</p>

  1.    username: is the email address of the libre follower account - MANDATORY
  2.    password: is the password of the libre follower account - MANDATORY
  3.    region: two letter code for the region/country where the account id registered. You can see that in the LibreLinkUP application > About > Registered Country. If none provided, will use `de`
  
<p align="center">
  <img width="400" height="250" src="https://github.com/xsen84/goLibreFollower/blob/main/doc/llu-region.png">
</p>


  5.    Text Size: how big the text/window will be. You can use samller size for smaller windows, or big ones if you want a bigger window. Play with the values and choose the one that fits.
  6.    Refresh Interval: how often the application will query the LibreLinkUP servers. Libre readings are uploaded every 60 seconds. Recomended value is 30-60 seconds. Don't set a value too low as you will overwhelm the server.
  7.    SAVE settings and restart application
       
- you can always change settings by pressing the settings button from the app. Remember, you need to restart application for changes to take effeect.
- the application will show different colors for in range/out of range values. The current limits are 
```
Low Glucose is under 70
High Glucose is over 180
```
