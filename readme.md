# goLibreFollower by Alexandru Mihai

- the app is meant ro be run on Windows and MacOS.
- it periodically retrieves glucose readings from the LibreLinkUP servers using a Follower account.

# Installation
## Windows
- download `goLibreFollower-WINDOWS_[VERSION].zip` from here: https://github.com/xsen84/goLibreFollower/releases
- extract/unzip zip archive and run the program
- on first run the settings window will open. See Settings for how to configure the application.


## Mac OSX
- download `goLibreFollower-MAC_[VERSION].zip` from here: https://github.com/xsen84/goLibreFollower/releases
- extract/unzip zip archive. Assumig you are in `Downloads` open a terminal and run:
```
cd ~/Downloads/
xattr -cr goLibreFollower.app
chmod +x goLibreFollower.app/Contents/MacOS/goLibreFollower
```
- if the operating system is blocking the app, please allow an exception acording to this: https://support.apple.com/en-en/guide/mac-help/mh40616/mac 
- on first run the settings window will open. See Settings for how to configure the application.


# Settings
- seettings are not applied real time - settings will be applied after close/open (restart) of the application
- on first run a settings window will be opened. You will need to provide followinf indo:
    1.  username: is the email address of the libre follower account - MANDATORY
    2.  password: is the password of the libre follower account - MANDATORY
    3.  region: two letter code for the region/country where the account id registered. You can see that in the LibreLinkUP application > About > Registered Country. If none provided, will use `de`
    4.  Text Size: how big the text/window will be. You can use samller size for smaller windows, or big ones if you want a bigger window. Play with the values and choose the one that fits.
    5.  Refresh Interval: how often the application will query the LibreLinkUP servers. Libre readings are uploaded every 60 seconds. Recomended value is 30-60 seconds. Don't set a value too low as you will overwhelm the server.
    6. SAVE settings and restart application
- you can always change settings by pressing the settings button from the app. Remember, you need to restart application for changes to take effeect.
- the application will show different colors for in range/out of range values. The current limits are 
```
Low Glucose is under 70
High Glucose is over 180
```
