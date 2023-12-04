# Shortest Pencil (sp.exe)

Opens a text editor on a filename that makes sense without thinking about it.
There isn't any magic here. (For example, from where the utility executes, you must ensure a folder named for the current year exists.)

What this gives me when I execute my keyboard shortcut is a file in the editor passed, named like this:
`notes/yyyy/yyyy-mm-dd-hhmm.txt`

If it were run *today*, something like this:
`notes/2023/2023-12-04-0842.txt`

## Build

`go build -o sp.exe`

One way to use it:

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
`gnome-terminal -e "~/notes/sp.exe subl"`

Under `Set Custom Shortcut`:
`Settings > Keyboard ? Keyboard Shortcuts - View and Customize Shortcuts > Custom Shortcuts`

### Windows

On Windows 11, add a BAT file with the following line, and configure a shortcut for it:
`start ~/notes/sp.exe notepad`

1. Right-click > Create Shortcut
1. With the shotcut file selected: `Alt + Enter`
1. In **Properties**, click **Shortcut key**
1. Type your shortcut
1. Click **OK**

Windows 11 doesn't allow `Space`, `Backspace`, or `Super` in these shortcuts. I went with `Ctrl + Alt + Shift + .` And then I added a macro for `Function + Space` on my keyboard.








