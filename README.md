# Glok
Glok! a simple cross-platform  clock for the command line, written in go

Special thanks to https://github.com/Henry0w For the exellent Documentation, and he has also been a huge help overall :)


# Usage
Open a command shell, type glok

and you've got a simple clock for your desktop!

# Installation:

Grab a binary from the [releases section](https://github.com/edward-riley/Glok/releases):


## Unix-like:

### System Wide
```sh
sudo chmod +x /path/to/glok
sudo cp /path/to/glok /usr/local/bin
```

### Locally:

Any changes to the path variable are not permanent, be sure to put them in your bash/zshrc.

#### Linux/BSD: 
```sh
chmod +x /path/to/glok
cp /path/to/glok $HOME/.local/bin
```
If it doesn't work after running this, run 

```sh
export PATH=$PATH:"$HOME/.local/bin"
```
#### Mac

Not recommended, you should install system-wide on mac

```sh
chmod +x /path/to/glok
cp /path/to/glok $HOME/bin/
export PATH=$PATH:"$HOME/.local/bin"
```


## Windows:
Place the .exe file in  C:\Program Files\Glok\

### Command Prompt:
```batch
SET "PATH=C:\Program Files\Glok"     
```
### Powershell:
```powershell
$env:path = $env:path + ";C:\Program Files\Glok"
```
# Building from source
## Dependencies:  
Go (any version)  
Git

## Building

Run `go build` inside the project directory, then follow the steps inside the install section to allow you to use the newly built binary.



I know for a fact that this works currently on Linux and windows, but as i have never used macOS,

i would be extremely gratefull if somebody would test it for me on that platform!

## CREDITS:
https://github.com/inancgumus/  For his extremely usefull screen library

Bevan on stack overflow for this comment that helped me add a binary to the windows path varible
https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows

