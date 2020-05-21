# Glok
Glok! a simple cross-platform  clock for the command line, written in go


# Installation:
## Install from Binary
To install Glok, you can grab the binary release here:

https://github.com/edward-riley/Glok/tree/master/glok/Binaries

### Linux / Mac Installation:
You can input this command to make glok avalible on the command line:

sudo cp /path/to/glok/binary/ /usr/bin/

### Windows:
I reccomend just placing the windows exe in somewhere like C:\Program Files\Glock\
Then, if you're on CommandPrompt:
SET "PATH=C:\Program Files\Glok"     
if you're on Powershell:
$env:path = $env:path + ";C:\Program Files\Glok"

## Install from source (probably best for people who know what they are doing)
Requirements: 
https://github.com/inancgumus/screen (see the link on how to install)
Go (any version)
Git

#### in your terminal/cmd/powershell type: git clone https://github.com/edward-riley/Glok/
this should download glok into  C:\users\$USER     on windows
and  /home/$USER/     on Linux/Mac

#### Be sure to place the glok folder in your go src path, which on linux/macos should be:
#### /home/$USER/go/src/

#### and on windows should be:
#### C:\User\ $USER\go\src

If there is now src directory inside your gopath, just make one

Now on all platforms you should open up cmd/powershell/terminal and type:

#### go build -o glok

you should then see the glok binary in your glok directory

#### Linux / Mac Installation:
You can input this command to make glok avalible on the command line:
sudo cp /path/to/glok/binary/ /usr/bin/

#### Windows:
I reccomend just placing the windows exe in somewhere like C:\Program Files\Glock\
Then, if you're on CommandPrompt:
SET "PATH=C:\Program Files\Glok"   
if you're on Powershell:
$env:path = $env:path + ";C:\Program Files\Glok"



I know for a fact that this works currently on Linux and windows, but as i have never used macOS,
i would be extremely gratefull if somebody would test it for me on that platform!

CREDITS:
https://github.com/inancgumus/  For his extremely usefull screen library
Bevan on stack overflow for this comment that helped me add a binary to the windows path varible
https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows
