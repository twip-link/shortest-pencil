# Shortest Pencil (sp.exe)

Opens a text editor on a filename that makes sense without thinking about it.
There isn't any magic here. (For example, from where the utility executes, you must ensure a folder named for the current year exists.)

What this gives me when I execute my keyboard shortcut is a file in the editor passed, named like this:
`notes/yyyy/yyyy-mm-dd-hhmm.txt`

If it were run *today*, something like this:
`notes/2023/2023-12-04-0842.txt`

## Build

`go build -o sp.exe`

## Example setup

1. Put the executable (matching your OS) in a folder called `notes`
1. Configure a shortcut for execution with either `subl` or `notepad` as the only argument.
1. Add a folder for the current year as a child of the folder where the executable is.

For me, in my user directory:
`notes/sp.exe`
`notes/2023/`

## Running the utility

How you call the shortcut is up to you; what's shown here worked for me.

### Ubuntu

On Ubuntu, I added a shortcut with the following command:
`gnome-terminal -e  '/home/username/notes/sp.exe subl'`

Under `Set Custom Shortcut`:
`Settings > Keyboard > Keyboard Shortcuts - View and Customize Shortcuts > Custom Shortcuts`

### Windows

On Windows 11, shortcuts using shortcut keys can get pretty convoluted. I went with AutoHotKey instead:
```
#Requires AutoHotkey v2.0

^!+.::{
    Run('C:\Users\username\notes\sp.exe notepad', 'C:\Users\username\notes\', 'Min')
    Sleep 100
    Send '{Enter}'
}
```

Finding a shortcut that won't overlap with your apps is important. 

I went with `Ctrl + Alt + Shift + .` 

And then I added a macro for `Function + Space` on my keyboard.
