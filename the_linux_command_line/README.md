# The Linux Command Line

By William E. Shotts. JR.

2013 - Second Internet Edition

[toc]

## Part 1 - Learning The Shell

### 1 - What Is The Shell?

#### Cursor Movement

> **A Few Words About Mice and Focus**

By highlighting some text, it is copied into a buffer maintained by X.
Pressing the middle mouse button will cause the text to be pasted at
the cursor location.

### 2 - Navigation

| Command | Description                             |
| ------- | --------------------------------------- |
| `pwd`   | Print name of current working directory |
| `cd`    | Change directory                        |
| `ls`    | List directory contents                 |

#### Change The Current Working Directory

##### Some Helpful Shortcuts

| Shortcut        | Result                                                           |
| --------------- | ---------------------------------------------------------------- |
| `cd`            | Change the working directory to your home directory              |
| `cd -`          | Change the working directory to the previous working directory   |
| `cd ~user_name` | Changes the working directory to the home directory of user_name |

### 3 - Exploring The System

| Command | Description             |
| ------- | ----------------------- |
| `ls`    | List directory contents |
| `file`  | Determine file type     |
| `less`  | View file contents      |

#### `ls`

| Option | Long Option      | Description                                                                                                                                                                                                                   |
| ------ | ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -a     | --all            | List all files, even those with names that begin with a period, which are normally not listed (i.e., hidden).                                                                                                                 |
| -A     | --almost-all     | Like the -a option above except it does not list . (current directory) and .. (parent directory).                                                                                                                             |
| -d     | --directory      | Ordinarily, if a directory is specified, ls will list the contents of the directory, not the directory itself. Use this option in conjunction with the -l option to see details about the directory rather than its contents. |
| -F     | --classify       | This option will append an indicator character to the end of each listed name. For example, a “/” if the name is a directory.                                                                                                 |
| -h     | --human-readable | In long format listings, display file sizes in human readable format rather than in bytes.                                                                                                                                    |
| -l     |                  | Display results in long format.                                                                                                                                                                                               |
| -r     | --reverse        | Display the results in reverse order. Normally, ls displays its results in ascending alphabetical order.                                                                                                                      |
| -S     | -t               | Sort results by file size. Sort by modification time.                                                                                                                                                                         |

`ls` Long Listing Fields:

```bash
-rw-r--r-- 1 root root 3576296 2007-04-03 11:05 Experience ubuntu.ogg
-rw-r--r-- 1 root root 1186219 2007-04-03 11:05 kubuntu-leaflet.png
-rw-r--r-- 1 root root
-rw-r--r-- 1 root root
-rw-r--r-- 1 root root
-rw-r--r-- 1 root root
-rw-r--r-- 1 root root  159744 2007-04-03 11:05 oo-derivatives.doc
-rw-r--r-- 1 root root   27837 2007-04-03 11:05 oo-maxwell.odt
-rw-r--r-- 1 root root   98816 2007-04-03 11:05 oo-trig.xls
-rw-r--r-- 1 root root  453764 2007-04-03 11:05 oo-welcome.odt
-rw-r--r-- 1 root root  358374 2007-04-03 11:05 ubuntu Sax.ogg
```

| Field            | Meaning                                                                                                                                                                                                                                                                                                                                                                                                              |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -rw-r--r--       | Access rights to the file. The first character indicates the type of file. Among the different types, a leading dash means a regular file, while a “d” indicates a directory. The next three characters are the access rights for the file's owner, the next three are for members of the file's group, and the final three are for everyone else. The full meaning of this is discussed in Chapter 9 – Permissions.                                                                   |
| 1             | File's number of hard links. See the discussion of links later in this chapter.                                                                                                                                                                                                                                                                                                                                      |
| root             | The username of the file's owner.                                                                                                                                                                                                                                                                                                                                                                                    |
| 32059            | The name of the group which owns the file. Size of the file in bytes.                                                                                                                                                                                                                                                                                                                                                |
| 2007-04-03 11:05 | Date and time of the file's last modification.                                                                                                                                                                                                                                                                                                                                                                       |
| oo-cd-cover.odf  | Name of the file.                                                                                                                                                                                                                                                                                                                                                                                                    |

#### `file`

Determining A File's Type With `file`

```bash
$ file filename
$ file picture.jpg
picture.jpg: JPEG image data, JFIF standard 1.01
```

#### A Guided Tour

| Directory      | Comments                                                                                                                                                                                                                                                                                                                                            |
| -------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| /              | The root directory. Where everything begins.                                                                                                                                                                                                                                                                                                        |
| /bin           | Contains binaries (programs) that must be present for the system to boot and run.                                                                                                                                                                                                                                                                   |
| /boot          | Contains the Linux kernel, initial RAM disk image (for drivers needed at boot time), and the boot loader.                                                                                                                                                                                                                                           |
|                | Interesting files:                                                                                                                                                                                                                                                                                                                                  |
|                | ● /boot/grub/grub.conf or menu.lst, which are used to configure the boot loader.                                                                                                                                                                                                                                                                    |
|                | ● /boot/vmlinuz, the Linux kernel                                                                                                                                                                                                                                                                                                                   |
| /dev           | This is a special directory which contains device nodes. “Everything is a file” also applies to devices. Here is where the kernel maintains a list of all the devices it understands.                                                                                                                                                               |
| /etc           | The /etc directory contains all of the system-wide configuration files. It also contains a collection of shell scripts which start each of the system services at boot time. Everything in this directory should be readable text.                                                                                                                  |
|                | Interesting files: While everything in /etc is interesting, here are some of my all-time favorites:                                                                                                                                                                                                                                                 |
|                | ● `/etc/crontab`, a file that defines when automated jobs will run.                                                                                                                                                                                                                                                                                 |
|                | ● `/etc/fstab`, a table of storage devices and their associated mount points.                                                                                                                                                                                                                                                                       |
|                | ● `/etc/passwd`, a list of the user accounts                                                                                                                                                                                                                                                                                                        |
| /home          | In normal configurations, each user is given a directory in /home. Ordinary users can only write files in their home directories. This limitation protects the system from errant user activity.                                                                                                                                                    |
| /lib           | Contains shared library files used by the core system programs. These are similar to DLLs in Windows.                                                                                                                                                                                                                                               |
| /lost+found    | Each formatted partition or device using a Linux file system, such as ext3, will have this directory. It is used in the case of                                                                                                                                                                                                                     |
|                | a partial recovery from a file system corruption event. Unless something really bad has happened to your system, this directory will remain empty.                                                                                                                                                                                                  |
| /media         | On modern Linux systems the /media directory will contain the mount points for removable media such as USB drives, CD-ROMs, etc. that are mounted automatically at insertion.                                                                                                                                                                       |
| /mnt           | On older Linux systems, the /mnt directory contains mount points for removable devices that have been mounted manually.                                                                                                                                                                                                                             |
| /opt           | The /opt directory is used to install “optional” software. This is mainly used to hold commercial software products that may be installed on your system.                                                                                                                                                                                           |
| /proc          | The /proc directory is special. It's not a real file system in the sense of files stored on your hard drive. Rather, it is a virtual file system maintained by the Linux kernel. The “files” it contains are peepholes into the kernel itself. The files are readable and will give you a picture of how the kernel sees your computer.             |
| /root          | This is the home directory for the root account.                                                                                                                                                                                                                                                                                                    |
| /sbin          | This directory contains “system” binaries. These are programs that perform vital system tasks that are generally reserved for the superuser.                                                                                                                                                                                                        |
| /tmp           | The /tmp directory is intended for storage of temporary, transient files created by various programs. Some configurations cause this directory to be emptied each time the system is rebooted.                                                                                                                                                      |
| /usr           | The /usr directory tree is likely the largest one on a Linux system. It contains all the programs and support files used by regular users.                                                                                                                                                                                                          |
| /usr/bin       | /usr/bin contains the executable programs installed by your Linux distribution. It is not uncommon for this directory to hold thousands of programs.                                                                                                                                                                                                |
| /usr/lib       | The shared libraries for the programs in /usr/bin.                                                                                                                                                                                                                                                                                                  |
| /usr/local     | The /usr/local tree is where programs that are not included with your distribution but are intended for system-wide use are installed. Programs compiled from source code are normally installed in /usr/local/bin. On a newly installed Linux system, this tree exists, but it will be empty until the system administrator puts something in it.  |
| /usr/sbin      | Contains more system administration programs.                                                                                                                                                                                                                                                                                                       |
| /usr/share     | /usr/share contains all the shared data used by programs in /usr/bin. This includes things like default configuration files, icons, screen backgrounds, sound files, etc.                                                                                                                                                                           |
| /usr/share/doc | Most packages installed on the system will include some kind of documentation. In /usr/share/doc, we will find documentation files organized by package.                                                                                                                                                                                            |
| /var           | With the exception of /tmp and /home, the directories we have looked at so far remain relatively static, that is, their contents don't change. The /var directory tree is where data that is likely to change is stored. Various databases, spool files, user mail, etc. are located here.                                                          |
| /var/log       | /var/log contains log files, records of various system activity. These are very important and should be monitored from time to time. The most useful one is /var/log/messages. Note that for security reasons on some systems, you must be the superuser to view log files .                                                                        |

### 4 - Manipulating Files And Directories

| Command | Description                       |
| ------- | --------------------------------- |
| cp      | Copy files and directories        |
| mv      | Move/rename files and directories |
| mkdir   | Create directories                |
| rm      | Remove files and directories      |
| ln      | Create hard and symbolic links    |

#### Wildcards

Using wildcards (which is also known as globbing) allow you to select filenames based on patterns of characters. **Wildcards always expand in sorted order**.

| Wildcard     | Meaning                                                          |
| ------------ | ---------------------------------------------------------------- |
| \*           | Matches any characters                                           |
| ?            | Matches any single character                                     |
| [character]  | Matches any character that is a member of the set characters     |
| [!character] | Matches any character that is not a member of the set characters |
| [\[:class:]] | Matches any character that is a member of the specified class    |

| Character Class | Meaning                            |
| --------------- | ---------------------------------- |
| [:alnum:]       | Matches any alphanumeric character |
| [:alpha:]       | Matches any alphabetic character   |
| [:digit:]       | Matches any numeral                |
| [:lower:]       | Matches any lowercase letter       |
| [:upper:]       | Matches any uppercase letter       |

#### `mkdir` - Create Directories

```bash
$ mkdir dir1
$ mkdir dir2 dir3 dir4
```

#### `cp` - Copy Files and Directories

```bash
$ cp item1 item2
$ cp item... directory # copy multiple items into a directory
```

| Option            | Meaning                                                                                                                                                                            |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -a, --archive     | Copy the files and directories and all of their attributes, including ownerships and permissions. Normally, copies take on the default attributes of the user performing the copy. |
| -i, --interactive | Before overwriting an existing file, prompt the user for confirmation. If this option is not specified, cp will silently overwrite files.                                          |
| -r, --recursive   | Recursively copy directories and their contents. This option (or the -a option) is required when copying directories.                                                              |
| -u, --update      | When copying files from one directory to another, only copy files that either don't exist, or are newer than the existing corresponding files, in the destination directory.       |
| -v, --verbose     | Display informative messages as the copy is performed.                                                                                                                             |

#### `mv` - Move And Rename Files

```bash
$ mv item1 item2 # rename file
$ mv item... directory
```

| Option            | Meaning                                                                                                                                                                    |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -i, --interactive | Before overwriting an existing file, prompt the user for confirmation. If this option is not specified, mv will silently overwrite files.                                  |
| -u, --update      | When moving files from one directory to another, only move files that either don't exist, or are newer than the existing corresponding files in the destination directory. |
| -v, --verbose     | Display informative messages as the move is                                                                                                                                |

#### `rm` - Remove Files And Directories

```bash
$ rm item...
```

| Option            | Meaning                                                                                                                                                                 |
| ----------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -i, --interactive | Before deleting an existing file, prompt the user for confirmation. If this option is not specified, rm will silently delete files.                                     |
| -r, --recursive   | Recursively delete directories. This means that if a directory being deleted has subdirectories, delete them too. To delete a directory, this option must be specified. |
| -f, --force       | Ignore nonexistent files and do not prompt. This overrides the --interactive option.                                                                                    |
| -v, --verbose     | Display informative messages as the deletion is performed.                                                                                                              |

> **Be careful with rm!**

Here is a useful tip. Whenever you use wildcards with rm (besides carefully checking your typing!), test the wildcard first with ls. This will let you see the files that will be deleted. Then press the up arrow key to recall the command and replace the ls with rm.

#### `ln` - Create Links

```bash
$ ln file link # hard link eg. ln -s dir1 dir1-sym
$ ln -s item link # symbolic link
```

##### Hard Links

Hard links have two important limitation:

1. A hard link cannot reference a file outside its own file system.
   This means a link cannot reference a file that is not on the same
   disk partition as the link itself.

2. A hard link may not reference a directory.

When a hard link is deleted, the link is removed but the contents
of the file itself continue to exist (that is, its space is
not deallocated) until all links to the file are deleted.

##### Symbolic Links

A file pointed to by a symbolic link, and the symbolic link itself are largely indistinguishable
from one another. For example, if you write something to the symbolic link, the referenced
file is written to. However when you delete a symbolic link,
only the link is deleted, not the file itself.

### 5 - Working With Commands

| Command | Description                                       |
| ------- | ------------------------------------------------- |
| type    | Indicate how a command name is interpreted        |
| which   | Display which executable program will be executed |
| help    | Get help for shell builtins                       |
| man     | Display a commands manual page                    |
| apropos | Display a list of appropriate commands            |
| info    | Display a command's info entry                    |
| whatis  | Display a very brief description of a command     |
| alias   | Create an alias for a command                     |

#### What Exactly Are Commands?

A command can be one of four different things:

1. An executable program: compiled binaries

2. A command built into the shell itself. e.g. `cd`

3. A shell function.

4. An alias.

#### Identifying Commands

##### `type` - Display A Command's Type

```bash
$ type command
```

```bash
[me@linuxbox ~]$ type type
type is a shell builtin
[me@linuxbox ~]$ type ls
ls is aliased to `ls --color=tty'
[me@linuxbox ~]$ type cp
cp is /bin/cp
```

##### `which` - Display An Executable's Location

```bash
$ which ls
/bin/ls
```

It only works for executable programs, not builtins nor aliases.

#### Getting A Command's Documentation

##### `help` - Get Help For Shell Builtins

```bash
$ help cd
```

##### `--help` - Disply Usage information

```bash
$ mkdir --help
```

Some programs don't support the “--help” option, but try it anyway. Often it results
in an error message that will reveal the same usage information.

##### `man` - Display A Program's Manual Page

```bash
$ man program
$ man ls
```

Manual is broken into sections and not only covers user commands but also
system administration commands, programming interfaces, file formats and
more. The table below describes the layout of the manual:

```bash
$ man section search_term
$ man 5 passwd
```

| Section | Contents                                       |
| ------- | ---------------------------------------------- |
| 1       | User commands                                  |
| 2       | Programming interfaces kernel system calls     |
| 3       | Programming interfaces to the C library        |
| 4       | Special files such as device nodes and drivers |
| 5       | File formats                                   |
| 6       | Games and amusements such as screen savers     |
| 7       | Miscellaneous                                  |
| 8       | System administration commands                 |

##### `apropos` - Display Appropriate Command

It is also possible to search the list of man pages for possible matches based
on a search term:

```bash
$ apropos passwd
chgpasswd (8)        - update group passwords in batch mode
chpasswd (8)         - update passwords in batch mode
gpasswd (1)          - administer /etc/group and /etc/gshadow
grub-mkpasswd-pbkdf2 (1) - generate hashed password for GRUB
openssl-passwd (1ssl) - compute password hashes
pam_localuser (8)    - require users to be listed in /etc/passwd
passwd (1)           - change user password
passwd (1ssl)        - OpenSSL application commands
passwd (5)           - the password file
update-passwd (8)    - safely update /etc/passwd, /etc/shadow and /etc/group
```

##### `whatis` - Display A Very Brief Description Of A Command

The `whatis` program displays the name and a one line description of a man page
matching a specified keyword:

```bash
$ whatis ls
ls (1)               - list directory contents
```

##### `info` - Display A Program's Info Entry

The GNU Project provides an alternative to man pages for their programs,
called "info". Info pages are displayed with a reader program named, appropriately
enough, `info`. Info pages are _hyperlinked_ much like web pages.

| Command           | Action                                                                        |
| ----------------- | ----------------------------------------------------------------------------- |
| ?                 | Display command help                                                          |
| PgUp or Backspace | Display previous page                                                         |
| PgDn or Space n   | Display next page                                                             |
| p                 | Next - Display the next node                                                  |
| u                 | Previous - Display the previous node                                          |
| Enter             | Up - Display the parent node of the currently displayed node, usually a menu. |
| q                 | Follow the hyperlink at the cursor location Quit                              |

```bash
$ info coreutils
```

##### README And Other Program Documentation Files

Many software packages installed on your system have documentation files residing in the /usr/share/doc directory.
Most of these are stored in plain text format and can be viewed with `less`.

#### Creating Your Own Commands With `alias`

**Use type to find out if the command is being used**

```bash
$ alias name='string'
$ unalias name
$ alias # to list all the aliases
```

### 6 - Redirection

| Command | Description                                                        |
| ------- | ------------------------------------------------------------------ |
| `cat`   | Concatenate files                                                  |
| `sort`  | Sort lines of text                                                 |
| `uniq`  | Report or omit repeated lines                                      |
| `grep`  | Print lines matching a pattern. **It exits with 1 when not found** |
| `wc`    | Print newline, word, and byte counts for each file                 |
| `head`  | Output the first part of a file                                    |
| `tail`  | Output the last part of a file                                     |
| `tee`   | Read from standard input and write to standard output and files    |

#### Standard Input, Output, And Error

Keeping with the Unix theme of “everything is a file,” programs such as ls actually send their results
to a special file called standard output (often expressed as stdout) and their status messages to another file
called standard error (stderr). By default, both standard output and standard error are linked to
the screen and not saved into a disk file.

In addition, many programs take input from a facility called standard input (stdin) which is,
by default, attached to the keyboard.

I/O redirection allows us to change where output goes and where input comes from. Normally, output goes to the screen and input comes from the keyboard, but with I/O redirection, we can change that.

#### Redirecting Standard Output

To redirect standard output to another file instead of the screen, we use the “>” redirection operator followed by the name of the file.

```bash
$ ls -l /usr/bin > ls-output.txt
$ ls -l ls-output.txt
-rw-rw-r-- 1 me me 167878 2008-02-01 15:07 ls-output.txt
```

However, we have a zero length when there is an error. The redirection operation started to rewrite the file and then stopped because of the error, resulting in its truncation.

```bash
$ ls -l /bin/usr > ls-output.txt
ls: cannot access /bin/usr: No such file or directory
$ ls -l ls-output.txt
-rw-rw-r-- 1 me me 0 2008-02-01 15:08 ls-output.txt
```

To append, use `>>`:

```bash
$ ls -l /usr/bin >> ls-output.txt
```

#### Redirecting Standard Error

To redirect standard error we must refer to its file descriptor. Program
can produce output on any of several numbered file streams. While we have referred to the first
three of these file streams as standard input, output and error, the
shell references them internally as file descriptors 0, 1 and 2, respectively.

Since standard error is the same as file descriptor number 2:

```bash
$ ls -l /bin/usr 2> ls-error.txt
```

##### Redirecting Standard Output And Standard Error To One File

To do this, we must redirect both standard output and standard error at the same time.
There are two ways to do this. First, the traditional way, which works
with old versions of the shell:

```bash
$ ls -l /bin/usr > ls-output.txt 2>&1
```

First we redirect standard output to the file ls-output.txt and then
we redirect file descriptor 2 (standard error) to file descriptor
one (standard output) using the notation 2>&1.

**Notice that the order of the redirections is significant.**: This is wrong:
`2>&1 > ls-output.txt`, which standard error is directed to the screen

Recent versions of bash provide a second, more streamlined method for performing this
combined redirection:

```bash
$ ls -l /bin/usr &> ls-output.txt
$ ls -l /bin/usr &>> ls-output.txt
```

##### Disposing Of Unwanted Output

The system provides a way to do this by redirecting output to a special file
called “/dev/null”. This file is a system device called a bit bucket
which accepts input and does nothing with it.

```bash
$ ls -l /bin/usr 2> /dev/null
```

#### Redirecting Standard Input

##### `cat` - Concatenate Files

The cat command reads one or more files and copies them to standard output like so:

```bash
cat [file...]
```

We could join file back together with this command:

```bash
$ cat movie.mpeg.0* > movie.mpeg
```

In the absence of filename arguments, cat copies standard input to standard output:

```bash
$ cat > lazy_dog.txt
The quick brown fox jumped over the lazy dog.
```

#### Pipelines

Using the pipe operator “|” (vertical bar), the
standard output of one command can be piped into the standard input of another:

```bash
command1 | command2
$ ls -l /usr/bin | less
```

> **The Difference Between > (redirection operator) and | (pipeline operator)**

Redirection operator silently creates or overwrites files, so you need to treat it with a lot of respect.

##### `wc` - Print Line, Word, and Byte Counts

The wc (word count) command is used to display the number of lines, words, and bytes
contained in files. For example:

```bash
$ wc ls-output.txt
7902 64566 503634 ls-output.txt
$ ls /bin /usr/bin | sort | uniq | wc -l  # output only lines
2728
```

##### `grep` - Print Lines Matching A Pattern

```bash
$ grep pattern [file...]
$ ls /bin /usr/bin | sort | uniq | grep zip
```

##### `head/tail` - Print First / Last Part Of Files

Sometimes you don't want all the output from a command.
You may only want the first few lines or the last few lines.

```bash
$ head -n 5 ls-output.txt
$ tail -n 5 ls-output.txt
$ ls /usr/bin | tail -n 5
```

`tail` has an option which allows you to view files in real-time:

```bash
$ tail -f /var/log/messages
```

##### `tee` - Read From Stdin And Output To Stdout And Files

The tee program reads standard input and copies it to both standard output (allowing
the data to continue down the pipeline) and to one or more files.

```bash
$ ls /usr/bin | tee ls.txt | grep zip bunzip2
bzip2
gunzip
gzip
```

### 7 - Seeing The World As The Shell Sees It

| Command | Description            |
| ------- | ---------------------- |
| `echo`  | Display a line of text |

#### Expansion

```bash
$ echo *
a ls.txt
```

The `*` character means match any characters in a filename. The shell
expands the "\*" into something else (in this instance, the names of the files
in the current working directory) before `echo` command is executed

##### Pathname Expansion

```bash
$ echo ls*
ls.txt

$ echo *t
ls.txt

$ echo [[:lower:]]*
a ls.txt

$ echo /usr/*/share
/usr/local/share
```

**Tip**: The filenames that begin with a period character are hidden.
`echo .*` almost works, but we will noticed "." and ".." there too.
There are workaround way such as `echo .[!.]*`, but it won't include
filenames with multiple leading periods. Use `ls -A` instead.

##### Tilde Expansion

If no user is named, the home directory of the current user was expanded:

```bash
$ echo ~
/home/me

$ echo ~foo
/home/foo
```

##### Arithmetic Expansion

This allow us to use the shell prompt as the calculator:

```bash
$ echo $((2 + 2))
4

# Tip: Single parentheses may be used to group multiple subexpressions
$ echo $(($((5**2))*3))
75
$ echo $(((5**2)*3))
75
$ FIVE=5
$ echo $((FIVE*2))
10
```

Arithmetic expansion only supports integers (whole numbers, no decimals),
but can perform quite a number of different operations:

| Operator | Descripion                                                                                      |
| -------- | ----------------------------------------------------------------------------------------------- |
| +        | Addition                                                                                        |
| -        | Substrction                                                                                     |
| \*       | Multiplication                                                                                  |
| /        | Division (but remember, since expansion only supports integer arithmetic, results are integers) |
| %        | Modulo, which simply means, "remainder"                                                         |
| \*\*     | Exponentation (powered of)                                                                      |

**Tip**: In Mini Mac, decimals is working.

##### Brace Expansion

You can create multiple text strings from a pattern containing braces:

```bash
$ echo Front-{A,B,C}-Back
Front-A-Back Front-B-Back Front-C-Back

$ echo Number_{1..5}
Number_1 Number_2 Number_3 Number_4 Number_5
```

Integers may also be zero-padded like so:

```bash
$ echo {01..15}
01 02 03 04 05 06 07 08 09 10 11 12 13 14 15

$ echo {001..15}
001 002 003 004 005 006 007 008 009 010 011 012 013 014 015
```

Reversing

```bash
$ echo {015..1}
015 014 013 012 011 010 009 008 007 006 005 004 003 002 001

$ echo {Z..A}
Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
```

Nested

```bash
$ echo {A{1,2},B{3,4}}b
A1b A2b B3b B4b

$ mkdir {2007..2009}-{01..12}
```

##### Parameter Expansion

```bash
$ echo $USER # me
```

To see a list of available variable: `printenv | less`

###### Command Substitution

```bash
$ echo $(ls)
$ ls -l $(which cp)
-rwxr-xr-x  1 root  wheel   150K Sep 16 21:28 /bin/cp

$ file $(ls -d /usr/bin/* | grep zip)
```

There is an alternate syntax for command substitution in older
shell programs, which use back-quotes instead of the dollar sign and parentheses:

```bash
$ ls -l `which cp`
```

#### Quoting

We'll learn how to control below cases

```bash
$ echo this is a     test
this is a test

$ echo The total is $100.00
The total is .00
```

##### Double Quotes

If you place text inside double quotes, all the special characters by the shell
**lose their special meaning and are treated as ordinary characters**. **EXCEPT "$",
"\\" and "`"**

It means, word-splitting, pathname expansion, tilde expansion, and brace expansion
are suppressed. But parameter expansion, arithmetic expansion, and command substitution
are still carried out.

```bash
$ echo "$USER $((2+2)) $(cal)"
me 4    November 2023
Su Mo Tu We Th Fr Sa
          1  2  3  4
 5  6  7  8  9 10 11
12 13 14 15 16 17 18
19 20 21 22 23 24 25
26 27 28 29 30
```

##### Single Quotes

If we need to suppress **ALL expansions**, we use single quotes.

> Comparison between single and double quotes

```bash
$ echo text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER
text /Users/alexsoon/ls-output.txt a b foo 4 alexsoon

$ echo "text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER"
text ~/*.txt {a,b} foo 4 alexsoon

$ echo 'text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER'
text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER
```

##### Escaping Characters

Use a backslash inside double quotes to selectively prevent
an expansion:

```bash
$ echo "The balance for user $USER is: \$5.00"
```

You may use it to escape "$", "!", "&", " ", "\\", and others without inside double
quotes

```bash
$ cat bad\ filename
```

> The first 32 characters in ASCII coding scheme are used to transmit commands
> to teletype-like devices

| Escape Sequence | Meaning                                      |
| --------------- | -------------------------------------------- |
| \a              | Bell ("Alert" - causes the computer to beep) |
| \b              | Backspace, "a\bb" -> b                       |
| \n              | Newline                                      |
| \r              | Carriage return                              |
| \t              | Tab                                          |

**Tip**: Adding an "-e" option to `echo` will enable interpretation
of escape sequences. You may also place them inside $' '.

```bash
$ echo -e "Time's up\a"
$ echo "Time's up" $'\a'
```

### 8 - Advanced Keyboard Tricks

| Command   | Description                              |
| --------- | ---------------------------------------- |
| `clear`   | Clear the screen                         |
| `history` | Display the contents of the history list |

#### Command Line Editing

##### Cursor Movement

| Key    | Action                                                                                             |
| ------ | -------------------------------------------------------------------------------------------------- |
| Ctrl-a | Move cursor to the beginning of the line.                                                          |
| Ctrl-e | Move cursor to the end of the line.                                                                |
| Ctrl-f | Move cursor forward one character, same as the right arrow key                                     |
| Ctrl-b | Move cursor backward one character, same as the left arrow key                                     |
| Alt-f  | Move cursor forward one word                                                                       |
| Alt-b  | Move cursor backward one word                                                                      |
| Ctrl-l | Clear the screen and move the cursor to the top left corner. The clear command does the same thing |

##### Modifying Text

| Key    | Action                                                                              |
| ------ | ----------------------------------------------------------------------------------- |
| Ctrl-d | Delete the character at the cursor location                                         |
| Ctrl-t | Transpose(exchange) the **character** at cursor location with the one preceding it  |
| Alt-d  | Transpose the **word** at the cursor location with the one preceding it             |
| Alt-l  | Convert the characters from the cursor location to the end of the word to lowercase |
| Alt-u  | Convert the characters from the cursor location to the end of the word to uppercase |

##### Cutting And Pasting (Killing and Yanking) Text

| Key           | Action                                                                                                                                         |
| ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| Ctrl-k        | Kill text from the cursor location to the end of line.                                                                                         |
| Ctrl-u        | Kill text from the cursor location to the beginning of the line.                                                                               |
| Alt-d         | Kill text from the cursor location to the end of the current word.                                                                             |
| Alt-Backspace | Kill text from the cursor location to the beginning of the current word. If the cursor is at the beginning of the word, kill the previous word |
| Ctrl-w        | same                                                                                                                                           |
| Ctrl-y        | Yank text from the kill-ring and insert it at the cursor location                                                                              |

**Tip**: While the `Alt` key serves as the meta key on modern keyboards, you can also press and release the `Esc` key to get
the same effect as holding down the `Alt` key if you're still using
a terminal

#### Completion

Completion will work on

- pathnames
- variable (if the beginning of word is a "$")
- Usernames (if the word begins with "~")
- hostnames (if the beginning of word is "@")

There are a number of control and meta key sequences that are
associated with completion:

| Key    | Action                                                                                                                                   |
| ------ | ---------------------------------------------------------------------------------------------------------------------------------------- |
| Alt-?  | Display a list of possible completions (**On most system you can also do this by pressing the tab a second time, which is much easier**) |
| Alt-\* | Insert all possile completions. This is useful when you want to use more than one possible match                                         |

#### Using History

The list of commands that have been entered is kept in your
home directory in a file called **.bash_history**

**Tip**: Some environment variable related:

| Environment Variable | Action                                                                                    | Example                       |
| -------------------- | ----------------------------------------------------------------------------------------- | ----------------------------- |
| HISTCONTROL          | Cause history recording feature to ignore a command if the same command was just recorded | export HISTCONTROL=ignoredups |
| HISTSIZE             | Size of command history (default 500)                                                     | export HISTSIZE=1000          |

##### Searching History

```bash
$ history | less
$ history | grep /usr/bin
```

To use our discovered line, use below example.

```bash
$ !88
```

`bash` also provides the ability to search the history list incrementally. To start
incremental search, press `Ctrl-r` followed by the text you are looking for. To find
the next occurrence of the text (moving "up" the history list), press `Ctrl-r` again.

Press `Enter` to execute or press `Ctrl-j` to copy the line. To quit searching,
press either `Ctrl-g` or `Ctrl-c`

##### History Expansion

| Sequence | Action                                                                     |
| -------- | -------------------------------------------------------------------------- |
| !!       | Repeat the last command. It is probably easier to press up arrow and enter |
| !number  | Repeat history list item number.                                           |
| !string  | Repeat last history list item starting with string                         |
| !?string | Repeat last history list item containing string                            |

Caution against using the "!string" and "!?string" forms unless you are absolutely
sure of the contents of the history list items.

**Tip**: The HISTORY EXPANSION section of the `bash` man page goes into all the gory
details. Feel free to explore!

**Tip**: In addition to the command history feature in `bash`, most Linux distributions
include a program called `script` that can be **used to record an entire shell session
and store it in a file**. The basic syntax of the command is: `script [file]`,
where file is the name of the file used for storing the recording. If no file
is specified, the file `typescript` is used.

### 9 - Permissions

Operating system in the Unix tradition differ from those in the MS-DOS tradition
in that they are not only multitasking systems, but also _multi-user systems_, as well.

It means that more than one person can be using the computer at the same time. For example,
if a computer is attached to a network or the Internet, remote users can log in via
`ssh` and operate the computer.

| Command  | Description                       |
| -------- | --------------------------------- |
| `id`     | Display user identity             |
| `chmod`  | Change a file's mode              |
| `umask`  | Set the default file permissions  |
| `su`     | Run a shell as another use        |
| `sudo`   | Execute a command as another user |
| `chown`  | Change a file's owner             |
| `chgrp`  | Change a file's group ownership   |
| `passwd` | Change a user's passwd            |

#### Owners, Group Members, And Everybody Else

The reason for `Permission denied` error message is that, as regular user,
we do not have permission to read this file.

In the Unix security model, a user may own files and directories. When a user
owns a file or directory, the user has control over its access. Users can,
in turn, belong to a group consisting of one or more users who are given
access to files and directories by their owners.

To find out your identity, use `id`

**Tip**: Fedora starts its numbering of regular user accounts at 500, while
Ubuntu starts at 1000.

User accounts are defined in the `/etc/passwd` file and groups are defined in
the `/etc/group`. You will notice there are accounts for the superuser (uid 0) and
various other system users.

#### Reading, Writing, And Executing

```bash
$ ls -l foo.txt
-rw-rw-r-- 1 me me 0 2008-03-06 14:52 foo.txt
```

The first ten characters of the listing are the _file attributes_.
The first of these characters is the **file type**.

| Attribute | File Type                                                                                                                                                |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -         | A regular file                                                                                                                                           |
| d         | A directory                                                                                                                                              |
| l         | A symbolic link, notice the remaining file attributes are always `rwxrwxrwx`. The real file attributes are those of the file the symbolic link points to |
| c         | A character special file. This file type refers to a device that handles data as a stream of bytes, such as a terminal or modem                          |
| b         | A block special file. This file type refers to a device that handles data in blocks such as a hard drive or CD-ROM drive                                 |

The remaining nine characters of the file attributes, called the file mode,
represent the read, write, and execute permissions for the file's owner,
the file's group owner, and everybody else:

| Owner | Group | World |
| ----- | ----- | ----- |
| rwx   | rwx   | rwx   |

| Attribute | Files                                                                                                                                                                                          | Directories                                                                                              |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| r         | Allows a file to be opened and read.                                                                                                                                                           | Allows a directory's contents to be listed if the execute attribute is also set.                         |
| w         | Allows a file to be written to or truncated, however this attribute does not allow files to be renamed or deleted. The ability to delete or rename files is determined by directory attributes | Allows files within a directory to be created, deleted, and renamed if the execute attribute is also set |
| x         | Allows a file to be treated as a program and executed. Program files written in scripting languages must also be set as readable to be executed.                                               | Allows a directory to be entered, e.g. cd directory                                                      |

| File Attributes | Meaning                                                                                                                                                                                                  |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -rwx------      | A regular file that is readable, writable, and executable by the file's owner. No one else has any access                                                                                                |
| -rw-------      | A regular file that is readable and writable by the file's owner. No one else has any access.                                                                                                            |
| -rw-r--r--      | A regular file that is readable and writable by the file's owner. Members of the file's owner group may read the file. This file is world-readable                                                       |
| -rwxr-xr-x      | A regular file that is readable, writable, and executable by the file's owner. The file may be read and executed by everybody else                                                                       |
| -rw-rw----      | A regular file that is readable and writable by the file's owner and members of the file's group owner only.                                                                                             |
| lrwxrwxrwx      | A symbolic link. All symbolic links have "dummy" permissions. The real permissions are kept with the actual file pointed to by the symbolic link.                                                        |
| drwxrwx---      | A directory. The owner and the members of the owner group may enter the directory and create, rename, and remove files within the directory.                                                             |
| drwxr-x---      | A directory. The owner may enter the directory and create, rename, and reletes files within the directory. Members of the owner group may enter the directory but cannot create, delete or rename files. |

##### `chmod` - Change File Mode

To change the mode (permissions) of a file. Only the file's owner or the superuser
can change the mode of a file or directory.

`chmod` supports two distinct ways of specifying mode changes:

- octal number representation

  - ```bash
    $ chmod 600 foo.txt
    $ ll foo*
      -rw-------@ 1 alexsoon  staff     0B Nov 23 17:09 foo.txt
    ```

  - | Octal | Binary | File Mode |
    | ----- | ------ | --------- |
    | 0     | 000    | **---**   |
    | 1     | 001    | --x       |
    | 2     | 010    | -w-       |
    | 3     | 011    | -wx       |
    | 4     | 100    | **r--**   |
    | 5     | 101    | **r-x**   |
    | 6     | 110    | **rw-**   |
    | 7     | 111    | **rwx**   |

- symbolic representation
  - | Symbol | Meaning                                            |
    | ------ | -------------------------------------------------- |
    | u      | Short for "user"(the file or directory owner)      |
    | g      | Group owner                                        |
    | o      | Short for "others", but means world                |
    | a      | Short for "all." The combination of "u", "g", "o". |
  - If no character is specified, "all" will be assumed.
  - | Notation   | Meaning                                                                                                                                                                      |
    | ---------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
    | u+x        | Add execute permission for the owner.                                                                                                                                        |
    | u-x        | Remove execute permission from the owner.                                                                                                                                    |
    | +x         | Add execute permission for the owner, group, and world. (a+x)                                                                                                                |
    | o-rw       | Remove the read and write permission from anyone besides the owner and group owner                                                                                           |
    | go=rw      | Set the group owner and anyone besides the owner to have read and write permission. If either the group owner or would previously had execute permissions, they are removed. |
    | u+x, go=rx | Add execute permission for the owner and set the permissions for the group and others to read and execute. Multiple specifications may be separated by commas                |

**Tip**: "--recursive" option acts on both files and directories, but we rarely want
files and directories to have the same permissions.

##### `umask` - Set Default Permissions

It controls the default permissions given to a file when it is created.
It uses `octal` notation to express a mask of bits **to be removed** from a
file's mode attributes:

```bash
$ umask 022
$ > foo.txt
$ ls -l foo.txt
-rw-r--r--@ 1 me  me     1B Nov 24 09:29 foo.txt
```

| Original file mode | --- rw- rw- rw- |
| ------------------ | --------------- |
| Mask               | 000 000 010 010 |
| Result             | 000 rw- r-- r-- |

```bash
$ umask 0002
$ > foo.txt
$ ls -l foo.txt
-rw-rw-r--@ 1 me  me     1B Nov 24 09:29 foo.txt
```

Most of the time you won't have to change the mask; the default
provided by your distribution will be fine. In some high-security
situations, however, you will want to control it.

##### Some Special Permissions

Though we usually see an octal permission mask expressed as a three digit number,
it is more technically correct to express it in four digits. Why? Because, in
addition to read, write, and execute permission, there are some other, less used,
permission settings.

- The first of these is the **setuid bit** (octal 4000).
  - It sets the _effective user ID_ from that of the real user ( the user running the program) to that of the program's owner.
  - When an ordinary user runs a program that is "setuid root", the program runs with the
    effective privileges of the superuser.
  - Clearly, because this raises security concerns, the number of setuid programs
    must be held to an absolute minimum.
- Te second less-used setting is the **setgid bit** (octal 2000).
  - If the setgid bit is set on a directory, newly created files in the directory will
    be given the group ownership of the directory rather than group ownership of the file's
    creator.
  - This is useful in a shared directory when members of a common group need access
    to all the files in the directory, regardless of the file owner's primary group.
- The third is called the sticky bit (octal 1000).
  - On files, Linux ignores the sticky bit
  - On directory, it prevents users from deleting or renaming files unless
    the user is either the owner of the directory, the owner of the file, or the superuser.
  - It is often used to control access to a shared directory, such as /tmp

```bash
$ chmod u+s program
# => -rws......
$ chmod g+s dir
# => ....rws...
$ chmod +t dir
# => .......rwt
$ chmod 7000 dir
# => ---S--S--T
```

#### Changing Identities

There are three ways to take on an alternate identity:

1. Log out and log back in as the alternate user.
2. Use the `su` command.
   - From within our own shell session, `su` command allow you to
     assume the identity of another user, and either start a new shell session
     with that user's IDs, or issue a single command as that user.
3. Use the `sudo` command.
   - Allows an administrator to set up a configuration file called `/etc/sudoers`
     and define specific commands that particular users are permitted to execute
     under an assumed identity.

##### `su` - Run a Shell With Substitute User and Group IDs

```bash
su [-[l]] [user]
```

`-l`: the resulting shell session is a login shell for the specified user. This
means that the user's environment is loaded and the working directory is
changed to the user's home directory. This is usually what we want.

**Tip**: If the user is not specified, the superuser is assumed.

**Tip**: Notice that (strangely) the "-l" may be abbreviated "-", which is how it is
most often used.

```bash
$ su -
Password:
[root@..]# exit
```

**Tip**: Superuser privileges has the trailing '#' rather than a '$'

It is also possible to execute a single command rather than starting a new
interactive command:

```bash
$ su -c 'command'
$ su -c 'ls -l /root/*'
```

##### `sudo` - Execute a Command As Another User

The administrator can configure `sudo` to allow an ordinary user to execute
commands as a different user (usually the superuser) in a very controlled way.

Another important difference is that the use of `sudo` does not require access
to the superuser's password. To authenticate using `sudo`, the user
uses his/her own password.

Another difference is `sudo` does not start a new shell, nor does it load another
user's environment. This means that commands do not need to be quoted.

```bash
$ sudo -l # list what privileges are granted by `sudo`
User me may run the following commands on this host:
    (ALL) ALL
```

**Tip**: Notice that after first use of sudo, you were not prompted for password.
This is because `sudo`, in most configurations, "trusts" you for several minutes until
its timer runs out.

##### `chown` - Change File Owner And Group

```bash
$ chown [owner][:[group]] file...
```

Change the owner and group owner of a file or directory.
Superuser privileges are required to use this command.

Here are some examples:

| Argument  | Results                                                                                                                  |
| --------- | ------------------------------------------------------------------------------------------------------------------------ |
| bob       | Changes the ownership of the file from its current owner to user `bob`                                                   |
| bob:users | Changes the ownership of the file from its current owner to user `bob` and changes the file group owner to group `users` |
| :admins   | Changes the group owner to the group admins. The file owner is unchanged                                                 |
| bob:      | Change the file owner from current owner to user `bob` and changes the group owner to the login group of user `bob`      |

Example: `janet` wants to copy a file to `tony` home directory and make `tony` owner

```bash
$ sudo cp myfile.txt ~tony
Password:
$ sudo chown tony: ~tony/myfile.txt
$ sudo ls -l ~tony/myfile.txt
 -rw-r--r-- 1 tony  tony  8031 2008-03-20 14:30 /home/tony/myfile.txt
```

##### `chgrp` - Change Group Ownership

In older versions of Unix, the `chown` command only changed file ownership, not
group ownership. It works much the same way as `chown`, except for being more limited.

#### Changing Your Password

```bash
$ passwd [user]
```

To change your password, just enter the `passwd` command.

If you have superuser privileges, you can specify a username as an argument
to set password for another user.

Other options are available to the superuser to allow account locking, password expiration,
etc. See the `passwd` man page for details.

#### Further Reading -- Create and Maintain Users

- `adduser`
- `useradd`
- `groupadd`

### 10 - Processes

Modern operating systems are usually multitasking, meaning
that they create the illusion of doing more than one thing at once by
rapidly switching from one executing program to another. The Linux
kernel manages this through the use of processes. Processes
are how Linux organizes the different programs waiting for their turn
at the CPU.

| Command    | Description                            |
| ---------- | -------------------------------------- |
| `ps`       | Report a snapshot of current processes |
| `top`      | Display tasks                          |
| `jobs`     | List active jobs                       |
| `bg`       | Place a job in the background          |
| `fg`       | Place a job in the foreground          |
| `kill`     | Send a signal to a process             |
| `killall`  | Kill processes by name                 |
| `shutdown` | Shutdown or reboot the system          |

#### How A Process Works

When a system starts up, the kernel initiates a few of its own activities
as processes and launches a program called `init`. `init`, in turn, runs
a series of shell scripts (located in /etc) called `init` scripts, which
start all the system services. Many of these services are implemented
as _daemon programs_, which just sit in the background and do their thing
without having any user interface. so even if we are not logged in,
the system is at least a little busy performing routine stuff.

The kernel maintains information about each process to help keep things
organized. For example, each process is assigned a number called a
process ID or IPD. PIDs are assigned in ascending order, with `init` always
getting PID 1. The kernel also keep track of the memory assigned to each
process, as well as the processes' readiness to resume execution.
Like files, processes also have owners and user IDs, effective user IDs, etc.

#### Viewing Processes

```bash
$ ps
  PID TTY           TIME CMD
62266 ttys000    0:00.00 bash
62272 ttys000    0:00.00 ps
```

TTY is short for "Teletype" and refers to the controlling terminal for the process.
TIME field is the amount of CPU time consumed by the process.

```bash
$ ps x
  PID   TT  STAT      TIME COMMAND
  378   ??  Ss     0:00.09 /System/Library/Frameworks/Metal.framework/Versions/A
  380   ??  Ss     0:03.99 /System/Library/Frameworks/Metal.framework/Versions/A
  455   ??  S      1:11.19 /usr/sbin/distnoted agent
  479   ??  S      0:57.05 /usr/sbin/cfprefsd agent
  480   ??  S      0:16.60 /usr/libexec/lsd
  481   ??  S      0:00.09 /System/Library/Frameworks/CoreServices.framework/Ver
  482   ??  S      0:18.98 /usr/libexec/secd
  487   ??  S      4:10.82 /usr/libexec/trustd --agent
  490   ??  SN     1:03.50 /System/Library/Frameworks/CoreServices.framework/Fra
```

`x` tells `ps` to show all of our processes regardless of what terminal (if any)
they are controlled by. The presence of a "?" in the TTY column indicates
no controlling terminal. Use this option to list every process that we own.

STAT is short for "state" and reveals the current status of the process:

| State | Meaning                                                                                                                                                                                                                                                                                         |
| ----- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| R     | Running. This means that the process is running or ready to run.                                                                                                                                                                                                                                |
| S     | Sleeping. The process is not running; rather, it is waiting for an event, such as keystroke or network packet.                                                                                                                                                                                  |
| D     | Uninterruptible Sleep. Process is waiting for I/O such as a disk drive.                                                                                                                                                                                                                         |
| T     | Stopped. Process has been instructed to stop. More on this later.                                                                                                                                                                                                                               |
| Z     | A defunct or "zombie" process. This is child process that has terminated, but has not been cleaned up by its parent.                                                                                                                                                                            |
| <     | A high priority process. It's possible to grant more important to a process, giving it more time on the CPU. This property of a process is called nicess. A process with high priority is said to be less nice because it's taking more of the CPU'stime, which leaves less for everybody else. |
| N     | ("Nice") A low priority process. It will only get processor time after other processes with higher priority have been serviced.                                                                                                                                                                 |

**Tip**: Without the leading dash invokes the command with "BSD style" behavior.

`ps aux`

| Header | Meaning                                                                                   |
| ------ | ----------------------------------------------------------------------------------------- |
| USER   | User ID. This is the owner of the process.                                                |
| %CPU   | CPU usage in percent.                                                                     |
| %MEM   | Memory usage in percent.                                                                  |
| VSZ    | Virtual memory size                                                                       |
| RSS    | Resident Set Size. The amount of physical memory (RAM) the process is using in kilobytes. |
| START  | Time when the process started. For values over 24 hours, a date is used                   |

##### Viewing Processes Dynamically With `top`

`top` displays a continuously updating (by default, every 3 seconds) display of the system processes
listed in order of process activity. The name "top" comes from the fact that the `top` program
is used to see the "top" processes on the system.

It accepts a number of keyboards commands:

- `h` for help
- `q` for quit

> top information fields

![Top_Info_Fields](images/top_info_fields.png)

| Row | Field         | Meaning                                                                                                                                                                                                                                                                                                                                                                                 |
| --- | ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | top           | Name of the program.                                                                                                                                                                                                                                                                                                                                                                    |
|     | 14:59:20      | Current time of day.                                                                                                                                                                                                                                                                                                                                                                    |
|     | up 6:30       | This is called _uptime_. It is the amount of time since the machine was last booted. Six and a half hours.                                                                                                                                                                                                                                                                              |
|     | 2 users       | There are two users logged in.                                                                                                                                                                                                                                                                                                                                                          |
|     | load average: | The number of processes that are waiting to run, that is the number of processes that are in a runnable state and are sharing the CPU. Three values are shown, each for a different period of time. The first is the average for the last 60 seconds, the next the previous 5 minutes, and finally the previous 15 minutes. **Values under 1.0 indicate that the machine is not busy.** |
| 2   | Tasks:        | This summarizes the number of processes and their various process states                                                                                                                                                                                                                                                                                                                |
| 3   | Cpu(s):       | This row describes the character of the activities that the CPU is performing.                                                                                                                                                                                                                                                                                                          |
|     | 0.7%us        | 0.7% of the CPU is being used for user processes. This means processes outside of the kernel itself                                                                                                                                                                                                                                                                                     |
|     | 1.0%sy        | 1.0% of the CPU is being used for system (kernel) process                                                                                                                                                                                                                                                                                                                               |
|     | 0.0%ni        | 0.0% of the CPU is being used by "nice" (low priority) processes                                                                                                                                                                                                                                                                                                                        |
|     | 98.3%id       | 98.3% of the CPU is idle.                                                                                                                                                                                                                                                                                                                                                               |
|     | 0.0%wa        | 0.0% of the CPU is waiting for I/O.                                                                                                                                                                                                                                                                                                                                                     |
| 4   | Mem:          | Shows how physical RAM is being used                                                                                                                                                                                                                                                                                                                                                    |
| 5   | Swap:         | Show how swap space (virtual memory) is being used                                                                                                                                                                                                                                                                                                                                      |

#### Controlling Processes

##### Putting A Process In The Background

```bash
$ <program> &
[1] 28237
```

- [1] -- we have started job number 1
- 28237 -- PID

Use `jobs` to list the jobs that have been launched from **our terminal**

```bash
$ jobs
[1]+ Running    <program> &
```

##### Returning A Process To The Foreground

```bash
$ jobs
[1]+ Running     <program> &
$ fg %1
```

**Tip**: If we only have one background job, the jobspec is optional. To terminate, press `Ctrl-c`

##### Stopping (Pausing) A Process

To stop a foreground process, press Ctrl-z.

```bash
$ <program>
[1]+ Stopped    <program>
```

We can:

1. Restore the program to foreground with `fg`
2. Move the program to the background with the `bg` command

   ```bash
   $ bg %1
   ```

It is handy if we launch a program from the command but forget to place it in the background by appending
the "&"

#### Signals

```bash
$ <program> &
[1] 28401
$ kill 28401
[1]+ Terminated         <program>
# or kill %1
```

##### Sending Signals To Multiple Processes With `kill`

Below used kill send signals, **TERM**(Terminate) is sent by default if no signal is specified on the command line

```bash
$ kill [-signal] PID...
$ kill -1 PID
$ kill -HUP PID
```

| Number |        | Name  | Meaning                                                                                                                                                                                                                   |
| ------ | ------ | ----- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1      |        | HUP   | Hangup. The effect of tis signal can be demonstrated by closing a terminal session. The fg program running on the terminal will be sent the signal and will terminate.                                                    |
|        |        |       | This signal is also used by many daemon programs to cause a reinitialization. When a daemon is sent this signal, it will restart and re-read its configuration file (e.g. Apache web server)                              |
| 2      | Ctrl-c | INT   | Interrupt. Terminate the program                                                                                                                                                                                          |
| 9      |        | KILL  | The `KILL` signel is never actually sent to the target program. Rather the kernel immediately terminates the process.                                                                                                     |
|        |        |       | When a process is terminated in this manner, it is given no opportunity to "clean up" after itself or save its work. For this reason, the `KILL` signal should only as a last resort when other termination signals fail. |
| 15     |        | TERM  | Terminate. If a program is still "alive" enough to receive signals, it will terminate.                                                                                                                                    |
| 18     |        | CONT  | Continue. This will restore a process after a STOP signal.                                                                                                                                                                |
| 19     |        | STOP  | Stop. This signal causes a process to pause without terminating. Like the `KILL` signal, it is not sent to the target process, and thus it cannot be ignored.                                                             |
| 3      |        | QUIT  | Quit.                                                                                                                                                                                                                     |
| 11     |        | SEGV  | Segmentation Violation. This signal is sent if a program makes illegal use of memory, that is, it tried to write somewhere it was not allowed to                                                                          |
| 20     | Ctrl-z | TSTP  | Terminal stop. Unlike the `STOP` signal, the `TSTP` signal is received by the program but the program may choose to ignore it.                                                                                            |
| 28     |        | WINCH | Window Change. This is a signal sent by the system when a window changes size. Some programs, like `top` and `less` will respond to this signal by redrawing themselves to fit the new window dimensions.                 |

**Tip**: To see a complete list of signals `kill -l`

##### Sending Signals To Multiple Processes With `killall`

Kill matched program or username

```bash
killall [- user] [-signal] name...
e.g
killall go run main.go
```

#### More Process Related Commands

| Command  | Description                                                                                                                                                       |
| -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `pstree` | Outputs a process list arranged in a tree-like pattern showing parent/child relationships between processes                                                       |
| `vmstat` | Outputs a snapshot of system resource usage including, memory, swap, and disk I/O. To see a continuous display, follow the command with a time delay (in seconds) |
|          | e.g `vmstat 5`                                                                                                                                                    |
| `xload`  | A graphical program that draws a graph showing system load over time                                                                                              |
| `tload`  | Similar to `xload` program, but draws the graph in the terminal.                                                                                                  |

## Part 2 - Configuration And The Environment

### 11 - The Environment

| Command    | Description                                          |
| ---------- | ---------------------------------------------------- |
| `printenv` | Print part or all of the environment                 |
| `set`      | Set shell options                                    |
| `export`   | Export environment to subsequently executed programs |
| `alias`    | Create an alias for a command                        |

#### What Is Stored In The Environment?

| Type                 | Description                         |
| -------------------- | ----------------------------------- |
| shell variable       | Bits of data placed there by `bash` |
| environment variable | everything else                     |

##### Examining The Environment

To see what is stored in the environment we can use

- `set` -- show both the shell and environment variables
- `printenv` -- show environment variables.

```bash
$ printenv | less
$ printenv USER # same with echo $USER
$ set | less
$ alias
```

##### Some Interesting Variables

| Variable | Contents                                                                                                                                                                                                      |
| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| DISPLAY  | The name of your display if you are running a graphical environment. Usually this is ":0", meaning the first display generated by the X server.                                                               |
| EDITOR   | The name of the program to be used for text editing                                                                                                                                                           |
| SHELL    | The name of your shell program.                                                                                                                                                                               |
| HOME     | The pathname of your home directory                                                                                                                                                                           |
| LANG     | Defines the character set and collation order of your language. (e.g en_US.UTF-8)                                                                                                                             |
| OLD_PWD  | The previous working directory                                                                                                                                                                                |
| PAGER    | The name of the program to be used for paging output. This is often set to /usr/bin/less                                                                                                                      |
| PATH     | A colon-separated list of directories that are searched when you enter the name of a executable program.                                                                                                      |
| PS1      | Prompt String 1. This defines the contents of your shell prompt. As we will later see, this can be extensively customized                                                                                     |
| PWD      | The current working directory                                                                                                                                                                                 |
| TERM     | The name of your terminal type (e.g xterm-256color). Unix-like systems support many terminal protocols; this variable sets the protocol to be used with your terminal emulator                                |
| TZ       | Specifies your timezone. Most Unix-like system maintain the computer's internal clock in (Coordinated Universal Time (UTC)) and then displays the local time by applying an offset specified by this variable |
| USER     | Your username                                                                                                                                                                                                 |

**Tip**: Don't worry if some of these values are missing. They vary by distribution.

#### How Is The Environment Established?

The exact sequence depends on the type of shell session being started:

- login shell session - which we are prompted for our username and password
- non-login shell session - launch a terminal session in the GUI

Login Shells read one or more startup files as shown:

| File            | Contents                                                                                                                                                   |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| /etc/profile    | A global configuration script that applies to all users                                                                                                    |
| ~/.bash_profile | A user's personal startup file. Can be used to extend or override settings in the global configuration script                                              |
| ~/.bash_login   | If ~/.bash_profile is not found, bash attempts to read this script.                                                                                        |
| ~/.profile      | if neither ~/.bash_profile nor ~/.bash_login is found, bash attempts to read this file. This is the default in Debian-based distributions, such as Ubuntu. |

Non-login shell sessions read the following startup files:

| File             | Contents                                                                                                      |
| ---------------- | ------------------------------------------------------------------------------------------------------------- |
| /etc/bash.bashrc | A global configuration script that applies to all users                                                       |
| ~/.bashrc        | A user's personal startup file. Can be used to extend or override settings in the global configuration script |

In addition to reading the startup files above, non-login shells also inherit the environment from their parent process, usually a login shell.

The ~/.bashrc file is probably the most important startup file from the ordinary user's point of view, since it is almost
always read. Non-login shells read it by default and most startup files for login shells are written in such a way as to read
the ~/.bashrc file as well

##### What's In A Startup File?

- Load .bashrc file (typical .bash_profile)

  ```bash
  if [ -f ~/.bashrc ]; then
    . ~/.bashrc
  fi
  ```

- Set PATH

  ```bash
  PATH=$PATH:$HOME/bin
  ```

- Export PATH, tells the shell to make the contents of PATH variable to child processes of this shell
  ```bash
  export PATH
  ```

#### Modifying The Environment

##### Which Files Should We Modify?

As a general rule:

| Action                                 | File                                |
| -------------------------------------- | ----------------------------------- |
| Add directories to your PATH           | .bash_profile (.profile for Ubuntu) |
| Define additional environment variable | .bash_profile (.profile for Ubuntu) |
| Other (e.g alias, export VAR, umask)   | .bashrc                             |

##### Text Editors

Two basic categories: graphical and text based

- Graphical: GNOME (`gedit`) and KDE (`kedit`, `kwrite`, `kate`)
- Text Based: `nano`, `vi`, `emacs` (`emacs` is by Richard Stallman!)

```bash
$ gedit some_file
```

**Tip**: Whenever we edit an important configuration file, it is always a good idea to create a backup copy of the file first.

Popular ways of indicating a backup file:

- .bak
- .sav
- .old
- .orig

```bash
$ cp .bashrc .bashrc.bak
```

###### Nano

| Event  | Action |
| ------ | ------ |
| Ctrl-x | exit   |
| Ctrl-o | save   |

##### Activating Our Changes

```bash
$ source .bashrc
```

### 12 - A Gentle Introduction To vi -- Skipped

### 13 - Customizing The Prompt

`PS1` short for prompt string one

```bash
[me@linuxbox ~]$ echo $PS1
[\u@\h \W]\$
```

#### Trying Some Alternative Prompt Designs

| Sequence | Value Displayed                                                                                                             |
| -------- | --------------------------------------------------------------------------------------------------------------------------- |
| \a       | ASCII bell. This makes the computer beep when it is encountered.                                                            |
| \d       | Current date in day, month, date format. For example, "Mon May 26"                                                          |
| \h       | Hostname of the local machine minus the training domain name                                                                |
| \H       | Full hostname                                                                                                               |
| \j       | Number of jobs running in the current shell session.                                                                        |
| \l       | Name of the current terminal device.                                                                                        |
| \n       | A newline character                                                                                                         |
| \r       | A carriage return                                                                                                           |
| \s       | Name of the shell program.                                                                                                  |
| \t       | Current time in 24 hour hours:minutes:seconds format.                                                                       |
| \T       | Current time in 12 hour format                                                                                              |
| \@       | Current time in 12 hour AM/PM format                                                                                        |
| \A       | Current time in 24 hour hours:minutes format                                                                                |
| \u       | username of the current user.                                                                                               |
| \v       | Version number of the shell                                                                                                 |
| \V       | Version and release numbers of the shell                                                                                    |
| \w       | Name of the current working directory                                                                                       |
| \W       | Last part of the current working directory name                                                                             |
| \!       | History number of the current command                                                                                       |
| \#       | Number of commands entered during this shell session                                                                        |
| \$       | This displays a "$" character unless you have superuser privileges. In that case, it displays a "#" instead.                |
| \[       | Signals the start of a series of one or more non-printing characters. This is used to embed non-printing control characters |
|          | which manipulate the terminal emulator in some way, such as moving the cursor or changing text color.                       |
| \]       | Signals the end of a non-printing character sequence                                                                        |

```bash
$ ps_old="$PS1"

$ PS1="\[a\]\$ "
$ # with beep sound

$ PS1="\A \h \$ "
17.33 linuxbox $

$ PS1="$ps_old"
```

#### Adding Color

The control code does not "print out" on the display, rather it is interpreted by the terminal
as an instruction. An ANSI escape code begins with an octal `033` (the code generated by the escape key),
followed by an optional character attribute, followed by an instruction. For example, the code
to set the text color to normal (attribute = 0), black text is:

```
\033[0;30m
```

Here is a table of available text colors. Notice that the colors are divided into two groups,
differentiated by the application of the bold character attribute (1) which creates the appearance
of "light" colors:

| Sequence   | Text Color | Sequence   | Text Color   |
| ---------- | ---------- | ---------- | ------------ |
| \033[0;30m | Black      | \033[1;30m | Dark Gray    |
| \033[0;31m | Red        | \033[1;31m | Light Red    |
| \033[0;32m | Green      | \033[1;32m | Light Green  |
| \033[0;33m | Brown      | \033[1;33m | Yellow       |
| \033[0;34m | Blue       | \033[1;34m | Light Blue   |
| \033[0;35m | Purple     | \033[1;35m | Light Purple |
| \033[0;36m | Cyan       | \033[1;36m | Light Cyan   |
| \033[0;37m | Light Grey | \033[1;37m | White        |

Follow change the prompt to red color:

```bash
PS1="\[\033[0;31m\]<\u@\h \W> \$"

# change back to previous color at the end
PS1="\[\033[0;31m\]<\u@\h \W> \$\[\033[0m\]"
```

To set background color:

| Sequence   | Background Color | Sequence   | Background Color |
| ---------- | ---------------- | ---------- | ---------------- |
| \033[0;40m | Black            | \033[0;44m | Blue             |
| \033[0;41m | Red              | \033[0;45m | Purple           |
| \033[0;42m | Green            | \033[0;46m | Cyan             |
| \033[0;43m | Brown            | \033[0;47m | Light Grey       |

```bash
PS1="\[\033[0;41m\]<\u@\h \W>\$ \[\033[0m\]"
```

**Note**: Besides the normal (0) and bold (1) character attributes, text may
also be given underscore (4), blinking (5), and inverse (7) attributes
as well. In the interests of good taste, many terminal emulators
refuse to honor the blinking attribute, however.

#### Moving The Cursor

This is commonly used to provide a clock or some other kind of information at a
different location on the screen, such as an upper corner each time the prompt is drawn.

| Escape Code | Action                                                                           |
| ----------- | -------------------------------------------------------------------------------- |
| \033[l;cH   | Move the cursor to line _l_ and column _c_                                       |
| \033[nA     | Move the cursor up _n_ lines                                                     |
| \033[nB     | Move the cursor down _n_ lines                                                   |
| \033[nC     | Move the cursor forward _n_ characters                                           |
| \033[nD     | Move the cursor backward _n_ characters                                          |
| \033[2J     | Clear the screen and move the cursor to the upper left corner (line 0, column 0) |
| \033[K      | Clear from the cursor position to the end of the current line                    |
| \033[s      | Store the current cursor position                                                |
| \033[u      | Recall the stored cursor position                                                |

Below example draws a red bar at the top of the screen containing a clock (rendered in yellow text)
each time the prompt is displayed.

```bash
PS1="\[\033[s\033[0;0H\033[0;41m\033[K\033[1;33m\t\033[0m\033[u\]<\u@\h \W>\$ "
```

## Part 3 - Common Tasks And Essential Tools

### 14 - Package Management

#### Packaging Systems

Different distributions use different packaging systems and as a general rule,
a package intended for one distribution is not compatible with another distribution.
Most distributions fall into one of two camps of packaging technologies:

- Debian ".deb" camp
  - Debian, Ubuntu, Xandros, Linspire
- Red Hat ".rpm" camp
  - Fedora, CentOS, Red Hat Enterprise Linux, OpenSUSE, Mandriva, PCLinuxOS

#### How a Package System Works

Virtually all software for a Linux system will be found on the
Internet. Most of it will be provided by the distribution vendor in the form
of `package file` and the rest will be available in source code that can be
installed manually.

##### Package Files

In addition to the files to be installed, the package file also includes
metadata about the package, such as

- test description of the package and its contents
- pre-installation scripts that perform configuration tasks before the package installation
- post-installation scripts that perform configuration tasks after the package installation

##### Repositories

Packages are made available to the users of a distribution in central repositories that may
contain many thousands of packages, each specially built and maintained for the
distribution.

A distribution may maintain several different repositories for different stages of the software
development life cycle. For example, there will usually be a "testing" repository
that contains packages that have just been built and are intended for use by brave
souls who are looking for bugs before they are released for general distribution.

##### Dependencies

Programs seldom "standalone"; rather they rely on the presence of other software components to get
their work done. Modern package management systems all provide some method of `dependency resolution`
to ensure that when a package is installed, all of its dependencies are installed, too.

##### High and Low-level Package Tools

Package management systems usually consist of two types of tools: low-level tools
which handle tasks such as installing and removing package files,
and high-level tools that perform metadata searching and dependency resolution.

| Distributions            | Low-Level Tools | High-Level Tools      |
| ------------------------ | --------------- | --------------------- |
| Debian-Style             | `dpkg`          | `apt-get`, `aptitude` |
| Fedora, Red Hat          | `rpm`           | `yum`                 |
| Enterprise Linux, CentOS |                 |                       |

#### Common Package Management Tasks

In the discussion below, the term `package _name` refers to the file that contains the package file.

##### Finding A Package In A Repository

| Style   | Command(s)                       |
| ------- | -------------------------------- |
| Debian  | `apt-get update`                 |
|         | `apt-cache search search_string` |
| Red Hat | `yum search search_string`       |

##### Installing A Package From A Repository

| Style   | Command(s)                     |
| ------- | ------------------------------ |
| Debian  | `apt-get update`               |
|         | `apt-get install package_name` |
| Red Hat | `yum install package_name`     |

##### Installing A Package From A Package File

Using a low-level tool:

| Style   | Command(s)                    |
| ------- | ----------------------------- |
| Debian  | `dpkg --install package_file` |
| Red Hat | `rpm -i package_file`         |

**Note**: Since this technique uses the low-level `rpm` program to perform
the installation, **no dependency resolution is performed**. If `rpm` discovers
a missing dependency, **`rpm` will exit with an error**.

##### Removing A Package

High-level tools:

| Style   | Command(s)                    |
| ------- | ----------------------------- |
| Debian  | `apt-get remove package_name` |
| Red Hat | `yum erase package_name`      |

##### Updating Packages From a Repository

Keeping the system up-to-date with the latest packages.

| Style   | Command(s)                        |
| ------- | --------------------------------- |
| Debian  | `apt-get update; apt-get upgrade` |
| Red Hat | `yum update`                      |

##### Upgrading A Package From A Package File

If an updated version of a package has been downloaded from a non-repository source,
it can be installed, replacing the previous version:

| Style   | Command(s)                    |
| ------- | ----------------------------- |
| Debian  | `dpkg --install package_file` |
| Red Hat | `rpm -U package_file`         |

##### Listing Installed Packages |

| Style   | Command(s)    |
| ------- | ------------- |
| Debian  | `dpkg --list` |
| Red Hat | `rpm -qa`     |

##### Determining If A Package Is Installed

| Style   | Command(s)                   |
| ------- | ---------------------------- |
| Debian  | `dpkg --status package_name` |
| Red Hat | `rpm -q package_name`        |

##### Displaying Info About An Installed Package

| Style   | Command(s)                    |
| ------- | ----------------------------- |
| Debian  | `apt-cache show package_name` |
| Red Hat | `yum info package_name`       |

##### Finding Which Package Installed A File

To determine what package is responsible for the installation of a particular file

| Style   | Command(s)                |
| ------- | ------------------------- |
| Debian  | `dpkg --search file_name` |
| Red Hat | `rpm -qf file_name`       |

```
$ rpm -qf /usr/bin/vim
vim-enhanced-7.4.160-1.el7_3.1.x86_64
```

### 15 - Storage Media

| Command                 | Description                                    |
| ----------------------- | ---------------------------------------------- |
| `df`                    | Display free disk space                        |
| `free`                  | Display free memory                            |
| `mount`                 | Mount a file system                            |
| `umount`                | Unmount a file system                          |
| `fsck`                  | Check and repair a file system                 |
| `fdisk`                 | Partition table manipulator                    |
| `mkfs`                  | Create a file system                           |
| `fdformat`              | Format a floppy disk                           |
| `dd`                    | Write block oriented data directly to a device |
| `genisoimage (mkisofs)` | Create an ISO 9660 image file                  |
| `wodim (cdrecord)`      | Write data to optical storage media            |
| `md5sum`                | Calculate an MD5 checksum                      |

#### Mounting and Unmounting Storage Devices

The first step in managing a storage device is attaching the device to the file
system tree. This process, called `mounting`, allows
the device to participate with the operating system.

A file named `/etc/fstab` lists the devices (typically hard disk partitions)
that are to be mounted at boot time

```bash
$ cat /etc/fstab
/dev/vda1            /                    ext3       noatime,acl,user_xattr 1 1
proc                 /proc                proc       defaults               0 0
sysfs                /sys                 sysfs      noauto                 0 0
debugfs              /sys/kernel/debug    debugfs    noauto                 0 0
devpts               /dev/pts             devpts     mode=0620,gid=5        0 0
/dev/vdb             /data1               ext4       defaults               1 2
```

These are the hard disk partitions. Each line of the file consists of six fields, as follows:

| Field | Contents         | Description                                                                                                              |
| ----- | ---------------- | ------------------------------------------------------------------------------------------------------------------------ |
| 1     | Device           | Actual name of device file associated with the physical device,                                                          |
|       |                  | such as /dev/hda1 (the first partition of the master                                                                     |
|       |                  | device on the first IDE channel). But with today's computers, which have                                                 |
|       |                  | many devices that are hot pluggable (like USD drives), many modern Linux distributions                                   |
|       |                  | associate a device with a text label instead. This label                                                                 |
|       |                  | is read by the operating system when the device is attached to the system. That way,                                     |
|       |                  | no matter which device file is assigned to the actual physical device, it can still be correctly identified              |
| 2     | Mount Point      | The directory where the device is attached to the file system tree.                                                      |
| 3     | File System Type | Most native Linux file systems are `ext3`, but many others are supported, such as `FAT16(msdos)`, `FAT32(vfat)`          |
|       |                  | , `NTFS (ntfs)`, `CD-ROM (iso9660)`, etc                                                                                 |
| 4     | Options          | File systems can be mounted with various options. It is possible, for example, to mount file systems as read-only, or to |
|       |                  | prevent any programs from being executed from them (a useful security feature for removable media).                      |
| 5     | Frequency        | A single number that specifies if and when file system is to be backed up with the `dump` command                        |
| 6     | Order            | A single number that specifies in what order file systems should be checked with the `fsck` command                      |

##### Viewing A List Of Mounted File Systems

The `mount` command is used to mount file systems. Entering the command without arguments will display a list
of the file systems currently mounted:

```bash
$ mount
/dev/vda1 on / type ext3 (rw,noatime,data=ordered)
```

The format is: `device` on `mount_point` type `file_system_type` `(options)`

**Tip**: type `vfat` is SD memory card

**Tip**: type `nfs4` is network drive

**Tip**: type `iso9660` is CD-ROM disc

> Exercise:

1. Insert CD-ROM disc

2. `umount <device>` e.g device is `/dev/hdc`

3. The next step is to create a new mount point for the disk

   - ```bash
     $ mkdir /mnt/cdrom
     $ mkdir -t iso9660 /dev/hdc/ /mnt/cdrom
     ```

   - It doesn't even have to be an empty directory, though **if you mount a device on a non-empty directory, you will
     not be able to see the directory's previous contents until you unmount the device**.

4. Let's try

   - ```bash
     $ cd /mnt/cdrom
     $ ls
     $ umount /dev/hdc
     umount: /mnt/cdrom: device is busy
     ```

   - It is because the device is being used by someone or some process.

   - ```bash
     $ cd
     $ umount /dev/hdc
     ```

> **Why Unmounting Is Important**

If you look at the output of the `free` command, which display statistics about memory usage,
you will see a statistic called "buffers". Computer systems are designed to go as fast as possible.
For example, printer prints extremely slow. If you were working on a spreadsheet or text document,
the computer would stop and become unavailable every time you printed.

This problem was solved by the advert of the `printer buffer`, a device containing some
RAM memory that would sit between the computer and the printer. With the printer buffer
in place, the computer would send the printer output to the buffer and it would quickly be stored
in the fast RAM so the computer could go back to work without waiting. Meanwhile, the printer
buffer would slowly spool the data to the printer from the buffer's memory at the speed at which
the printer could accept it.

This idea of buffering is used extensively in computers to make them faster. On a Linux
system for example, you will notice that the system seems to fill up memory the longer
it is used. This does not mean Linux is "using" all the memory, it means that Linux is
taking advantage of all the available memory to do as much buffering as it can.

Unmounting a device entails writing all the remaining data to the device so that it
can be safely removed. If the device is removed without unmounting it first, the
possibility exists that not all the data destined for the device has been transferred.
In some cases, this data may include vital directory updates, which will lead to file
system corruption, one of the worst things that can happen on a computer.

##### Determining Device Names

If we list the contents of the /dev directory, we can see that there are lots of devices.

The contents of this listing reveal some patterns of device naming. Here are a few:

| Pattern   | Device                                                                                                                                             |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| /dev/fd\* | Floppy disk drives                                                                                                                                 |
| /dev/hd\* | IDE(PATA) disks on older system.                                                                                                                   |
| /dev/lp\* | Printers                                                                                                                                           |
| /dev/sd\* | SCSI disks. On recent Linux systems, the kernel treats all disk like devices (PATA/SATA hard disks, flash drives, USB mass storage devices such as |
|           | portable music players, and digital cameras) as SCSI disks                                                                                         |
| /dev/sr\* | Optical drives (CD/DVD readers and burners)                                                                                                        |

If you are working on a system that does not automatically mount removable devices, you can
use the following technique to determine how the removable device is named when it is attached:

1. `sudo tail -f /var/log/messages` or `sudo tail -f /var/log/syslog`

2. Next, plug the removable device in and observe the log

After that, we can now mount the flash drive:

```bash
$ sudo mkdir /mnt/flash
$ sudo mount /dev/sdb1 /mnt/flash
$ df
```

#### Creating New File Systems

If we want to reformat the flash drive with a Linux native file system, rather than the FAT32 system it has now:

1. (optional) create a new partition layout if the existing one is not to our liking

2. Create a new, empty file system on the drive.

This process of partitioning and formatting can be used anytime additional storage devices are
added to the system.

##### Manipulating Partitions With `fdisk`

The `fdisk` program allows us to interact directly with disk-like devices at a very low level.

```bash
# First, umount the flash drive
$ sudo umount /dev/sdb1
# Specify the entire device instead of partiion number
$ sudo fdisk /dev/sdb
Command (m for help): m
Command action
   a   toggle a bootable flag
   b   edit bsd disklabel
   c   toggle the dos compatibility flag
   d   delete a partition
   l   list known partition types
   m   print this menu
   n   add a new partition
   o   create a new empty DOS partition table
   p   print the partition table
   q   quit without saving changes
   s   create a new empty Sun disklabel
   t   change a partition\'s system id
   u   change display/entry units
   v   verify the partition table
   w   write table to disk and exit
   x   extra functionality (experts only)

Command (m for help): p
Disk /dev/sdb: 16 MB, 16006656 bytes
1 heads, 31 sectors/track, 1008 cylinders
Units = cylinders of 31 * 512 = 15872 bytes
   Device Boot      Start         End      Blocks   Id  System
/dev/sdb1               2        1008       15608+   b  W95 FAT32
```

In above example, we see a 16 MB device with a single partition (1) that uses
1006 of the available of 1008 cylinders on the device. The partition is identified
as a Windows 95 FAT32 partition. Some programs will use this identifier to limit the kinds
of operation that can be done to the disk, but most of the time it is not critical to change it.
However, we will change it to indicate a Linux partition below.

Enter `l` to list possible types of partition types. We can find `83` for Linux. `b` is
our existing partition type.

Enter `t` to change a partition's system id to `83`

Up to this point, the device has been untouched (all the changes have been stored in memory,
not on the physical device).

Enter `w` to write the modified partition table to device and exit.

##### Creating A New File System With `mkfs`

`mkfs` (make file system) can create file systems in a variety of formats.

- `-t` to specify the system type

```bash
$ sudo mkfs -t ext3 /dev/sdb1
```

To re-format the device to its original FAT32 file system, specify "vfat" as the file system type

```bash
$ sudo mkfs -t vfat /dev/sdb1
```

#### Testing And Repairing File Systems

In /etc/fstab, we saw some mysterious digits at the end of each line. Each time the system
boots, it routinely checks the integrity of the file systems before mounting them.
This is done by the `fsck` (file system check) program. The last number in each `fstab`
entry specifies the order in which the devices are to be checked. Devices with
a zero as the last digit are not routinely checked.

In addition, `fsck` can also repair corrupt file system with varying degrees of success,
depending on the amount of damage. On Unix-like file systems, recovered portions of files
are placed in the `/lost+found` directory.

```bash
$ sudo fsck /dev/sdb1
```

#### Formatting Floppy Disks

```bash
$ sudo fdformat /dev/fd0
$ sudo mkfs -t msdos /dev/fd0
```

#### Moving Data Directly To/From Devices

The `dd` program performs the task. I copies blocks of data from one place to another. It
uses a unique syntax (for historical reasons) and is usually used by this way:

```bash
dd if=input_file of=output_file [bs=block_size [count=blocks]]
```

Let's say we want to exactly copy the first drive to the second exactly same size USB flash drives.

```bash
dd if=/dev/sdb of=/dev/sdc
```

Alternately, if only the first device were attached to the computer, we could copy its content to an
ordinary file for later restoration or copying:

```bash
dd if=/dev/sdb of=flash_drive.img
```

**Warning**: Though its name derives from "data definition", it is sometimes called "destroy disk" because
users often mistype either `if` or `of` specifications. **Always double check before pressing enter**

#### Creating CD-ROM Images

`CD-R` or `CD-RW`

Steps:

1. Constructing an iso image file that is the exact file system image of the CD-ROM
2. Writing the image file onto the CD-ROM media.

##### Creating an Image Copy Of A CD-ROM

```bash
$ dd if=/dev/cdrom of=ubuntu.iso
```

This technique works for data DVDs as well, but will not work for audio CDs, as
they do not use a file system for storage. For audio CDs, look at the `cdrdao` command.

##### Creating an Image From A Collection Of Files

```bash
$ genisoimage -o cd-rom.iso -R -J ~/cd-rom-files
```

The "-R" option adds metadata for the _Rock Ridge_ extensions, which
allows the use of long filenames and POSIX style file permissions.

The "-J" option enables the _Joliet_ extensions, which permit long filenames for Windows.

> **A Program By Any Other Name...**

`mkisofs` and `cdrecord` programs can be used for creating and burning optical media like
CD-ROMs and DVDs. These programs were part of a popular package called "cdr-tools"
authored by Jörg Schilling. In the summer of 2006, Mr. Schilling made a license change
to a portion of "cdrtools" package which, in the opinion of many in the Linux community,
created a license incompatibility with the GNU GPL.

As a result, a _fork_ of the cdrtools project was started that now includes replacement
programs for `cdrecord` and `mkisofs` named `wodim` and `genisoimage` respectively.

#### Writing CD-ROM Images

After we have an image file, we can burn it onto our optical media.

##### Mounting An ISO Image Directly

By adding the "-o loop" option to `mount` (with required "-t iso9660" file system type), we can mount the image file
as though it were a device and attach it to the file system tree:

```bash
$ mkdir /mnt/iso_image
$ mount -t iso9660 -o loop image.iso /mnt/iso_image
```

**Tip**: Remember to unmount the image when it is no longer needed.

##### Blanking A Re-Writable CD-ROM

The `wodim` program offers several types. The most minimal (and fastest) is the "fast" type:

```bash
$ wodim dev=/dev/cdrw blank=fast
```

##### Writing An Image

```bash
$ wodim dev=/dev/cdrw image.iso
```

In addition to the device name and image file.
`wodim` supports a very large set of options.
Two common ones are "-v" for verbose output, and "-dao" which writes the disc
in _disc-at-once_ mode, this mode should be used if you are preparing a disc
for commercial reproduction. The default mode for `wodim` is track-at-once
which is useful for recording tracks.

#### Verifying the Integrity of an ISO image

In most cases, a distributor of an iso image will also supply a checksum file,
which is the result of an exotic mathematical calculation resulting in
a number that represents the content of the target file. **If the contents
of the file change by even one bit**, the resulting check sum will be much
different.

The most common method of checksum generation uses the `md5sum` program.
When you use `md5sum`, it produces a unique hexadecimal number:

```bash
$ md5sum image.iso
```

After you download an image, you should run `md5sum` against it and
compare the results with the `md5sum` value supplied by the publisher.

### 16 - Networking

| Command      | Description                                                                                                        |
| ------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ping`       | Send a ICMP ECHO_REQUEST to network hosts                                                                          |
| `traceroute` | Print the route packets trace to a network host                                                                    |
| `netstat`    | Print network connections, routing tables, interface statistics, masquerade connections, and multicast memberships |
| `ftp`        | Internet file transfer program                                                                                     |
| `wget`       | Non-interactive network downloader                                                                                 |
| `ssh`        | OpenSSH SSH client (remote login program)                                                                          |

#### Pre-knowledge

> IP (Internet Protocol) address

A numerical label such as 192.0.2.1 that is connected to a computer network that uses the
Internet Protocol for communication. An IP address serves two main functions:

- network interface identification

- location addressing

Network administrators assign an IP address to each device connected to a network. Such
assignments may be on a static (fixed or permanent) or dynamic basis.

**Private addresses**

Private networks are widely used and typically connect to the Internet with network address translation (NAT), when needed.

Three non-overlapping ranges of IPV4 addresses for private networks are reserved.
These addresses are not routed on the Internet and thus their use need not be coordinated with an
IP address registry.

| Name         | CIDR block     | Address range                 | Number of addresses | Classful Description                   |
| ------------ | -------------- | ----------------------------- | ------------------- | -------------------------------------- |
| 24-bit block | 10.0.0.0/8     | 10.0.0.0 - 10.255.255.255     | 16 777 216          | Single class A                         |
| 20-bit block | 172.16.0.0/12  | 172.16.0.0 - 172.31.255.255   | 1 048 576           | Contiguous range of 16 Class B blocks  |
| 16-bit block | 192.168.0.0/16 | 192.168.0.0 - 192.168.255.255 | 65 536              | Contiguous range of 256 Class C blocks |

**Tip**: In a home situation, a public IP address is the IP address assigned to the home's network by the ISP.
The ISP usually assigns a dynamic IP for home. Else it is more likely to be abused by customers who host websites from
home, or by hackers who can try the same IP address over and over until they breach a network.

> Host and domain name

A hostname is a label that is assigned to a device connected to a computer network
and is used to identify the device in various forms of electronic communication, such as the World Wide Web.

Internet hostnames may have appended the name of a Domain Name System (DNS) domain,
separated from the host-specific label by a period ("dot"). In the latter form,
a hostname is also called a domain name. If the domain name is completely specified,
including a top-level domain of the internet, then the hostname is said to be a full qualified
domain name (FQDM). Hostnames that include DNS domains are often stored in the Domain
Name System together with the IP address of the host they represent
for the purpose of mapping the hostname to an address, or the reverse process.

The internet hostnames is translated into an IP addresses via the local hosts file, or the
Domain Name System (DNS) resolver. It is possible for a single host computer to have several hostnames.

> URI (Uniform Resource Identifier)

It is a unique sequence of characters that identifies a logical or physical resource used
by web technologies.

A Uniform Resource Locators (URLS) provide the location of the resource. A URI identifies
the resource by name at the specified location or URL.

The URI generic syntax consists of five components:

```bash
URI = scheme ":" ["//" authority] path ["?" query] ["#" fragment]
```

Examples of popular schemes include `http` `https` `ftp` `mailto` `file` `data` `irc`

The authority component consists of subcomponents:

```bash
authority = [userinfo "@"] host [":" port]
```

ftp://username:password@example.com/file.txt

scheme: ftp

path : /file.txt

#### Examining And Monitoring A Network

> `ping`

The `ping` command sends a special network packet called an IMCP ECHO_REQUEST to a specified host.
Most network devices receiving this package will reply to it, allowing the network connection to be verified.

```bash
$ ping google.com
PING google.com (64.233.170.102): 56 data bytes
64 bytes from 64.233.170.102: icmp_seq=0 ttl=64 time=13.406 ms
64 bytes from 64.233.170.102: icmp_seq=1 ttl=64 time=12.971 ms
^C
--- google.com ping statistics ---
2 packets transmitted, 2 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 12.971/13.189/13.406/0.217 ms
```

A properly performing network will exhibit zero percent packet loss.

**Note**: It is possible to configure most network devices (including Linux hosts) to ignore these packets.
This is usually done for security reasons, to partially obscure a host from a potential attacker.
It is also common for firewalls to be configured to block IMCP traffic.

> `traceroute`

`traceroute` program displays a listing of all the "hops" network traffic takes to get from the local system to a specific host.

```bash
$ traceroute slashdot.org
...
16  slashdot.org (216.34.181.45)  42.727 ms  42.016 ms  41.437 ms
```

In the output, we can see that connecting requires traversing sixteen routers. For routers that provided identifying information,
we see their hostnames, IP addresses, and performance data, which includes three samples of
round-trip time from the local system to the router. For routers that do not provide identifying information
(because of router configuration, network congestion, firewalls, etc), we see asterisks like `* * *`.

> `netstat`

The `netstat` program is used to examine various network settings and statistics. Using the "-ie" option,
we can examine the network interfaces in our system:

```bash
$ netstat -ie
eth0 Link encap:Ethernet HWaddr 00:1d:09:9b:99:67
        inet addr:192.168.1.2 Bcast:192.168.1.255 Mask:255.255.255.0
        inet6 addr: fe80::21d:9ff:fe9b:9967/64 Scope:Link
        UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
        RX packets:238488 errors:0 dropped:0 overruns:0 frame:0
        TX packets:403217 errors:0 dropped:0 overruns:0 carrier:0
        collisions:0 txqueuelen:100
        RX bytes:153098921 (146.0 MB)  TX bytes:261035246 (248.9 MB)
        Memory:fdfc0000-fdfe0000
lo      Link encap:Local Loopback
        inet addr:127.0.0.1  Mask:255.0.0.0
        inet6 addr: ::1/128 Scope:Host
        UP LOOPBACK RUNNING  MTU:16436  Metric:1
        RX packets:2208 errors:0 dropped:0 overruns:0 frame:0
        TX packets:2208 errors:0 dropped:0 overruns:0 carrier:0
        collisions:0 txqueuelen:0
        RX bytes:111490 (108.8 KB)  TX bytes:111490 (108.8 KB)
```

In the example above, we see that our test system has two network interfaces. The first called `eth0`,
is the Ethernet interface and the second, called `lo`, is the `loopback` interface, a virtual
interface that the system uses to "talk to itself".

Presence of valid IP address in the `inet addr` field on the second line.

Using the "-r" option will display the kernel's kernel's network routing table. This shows how
the network is configured to send packets from network to network:

```bash
$ netstat -r
Destination Gateway Genmask Flags MSS Window irtt Iface
192.168.1.0  *           255.255.255.0 U         0 0          0 eth0
default      192.168.1.1 0.0.0.0       UG        0 0          0 eth0
```

In this example, we see a typical routing table for a client machine on a LAN (Local Area Network)
behind a firewall/router. The first line of the listing shows the destination `192.168.1.0` IP addresses
that **end in zero refer to networks rather than individual host**, so this destination means any host
on the LAN.

The next field, Gateway, is the name or IP address of the gateway (router) used to go from the
current host to the destination network. An asterisk in this field indicates
that no gateway is needed.

The first line contains the destination `default`. This means any traffic destined for a
network that is not otherwise listed in the table.

#### Transporting Files Over a Network

##### `ftp` File Transfer Protocol

FTP (in its original form) is not secure, because it sends account names and passwords
in _cleartext_. This means that they are not encrypted and anyone sniffing the network
can see them. Because of this, almost all FTP done over the Internet is done by anonymous FTP
servers which allows anyone to login using the login name "anonymous" and a meaningless password.

Below example fileserver is an FTP server

```bash
$ ftp fileserver.
Name: anonymous
Password:
ftp> cd pub/cd_images/Ubuntu-8.04
ftp> ls
ftp> lcd Desktop
ftp> get ubuntu-8.04...iso
ftp> bye
```

##### `wget`

Downloading content from both web and FTP sites. Single files, multiple files and
even entire sites can be downloaded.

To download the first page of `linuxcommand.org`

```bash
$ wget http://linuxcommand.org/index.php
```

#### Secure Communication With Remote Hosts

> `ssh`

A new protocol called SSH (Secure Shell) solves two basic problems
of secure communication with a remote host.

- authenticates the remote host is who it says it is (preventing so-called "man in the middle" attack)
- encrypts all of the communications between the local and remote hosts.

SSH consists of two parts. An SSH server runs on the remote host, listening for incoming connections on port 22,
while an SSH client is used on the local system to communicate with the remote server.

Most Linux distributions ships an implementation of SSH called OpenSSH from the OpenBSD
project. Some distributions include both the client and the server packages by default (for example, Red Had),
while others only supply the client. To enable a system to receive remote connections, it must have
the `OpenSSH-server` package installed.

**Tip**: If you don't have a remote system to connect to but want to try these examples, make sure
you have installed SSH server and use `localhost` as the name of the remote host.

```bash
$ ssh remote-sys
```

The remote shell session continues until the user enters the `exit` command at the remote shell prompt, thereby closing the remote connection.

It is also possible to connect to remote systems using a different username.

```bash
$ ssh bob@remote-sys
```

Besides opening a shell session on a remote system, `ssh` also allows us to execute a single command:

```bash
$ ssh remote-sys free
$ ssh remote-sys 'ls *' > dirlist.txt
```

> **Tunneling With SSH**

Part of what happens when you establish a connection with a remote host via SSH is that an encrypted tunnel is created
between the local and remote systems. This tunnel is used to allow commands typed at the local system to be
transmitted safely to the remote system.

#### `scp` And `sftp`

The OpenSSH package also includes two progams that can make use of an SSH-encrypted tunnel to copy files across the network.

> `scp` (secure copy)

Used much like the familiar `cp` to copy files. The notable difference is that the source or destination pathnames may be
preceded with the name of a remote host, followed by a colon character.

For example, if we wanted to copy a document named `document.txt` from our home directory on the remote system,
`remote-sys`, to the current working directory on our local system:

```bash
$ scp remote-sys:document.txt .
```

> `sftp`, a secure replacement for the `ftp` program.

Instead of transmitting everything in cleartext, it uses an SSH encrypted tunnel.
`sftp` has an important advantage over conventional `ftp` in that it does not require
an FTP server to be running on the remote host. It only requires the SSH server.

```bash
$ sftp remote-sys
sftp> ls
sftp> lcd Desktop
sftp> get ubuntu-8.04-desktop-i386.iso
sftp> bye
```

### 17 - Searching For Files

| Command  | Description                                         |
| -------- | --------------------------------------------------- |
| `locate` | Find files by name                                  |
| `find`   | Search for files in a directory hierarchy           |
| `xargs`  | Build and execute command lines from standard input |
| `touch`  | Change file times                                   |
| `stat`   | Display file or file system status                  |

#### `locate` - Find Files The Easy Way

```bash
$ locate bin/zip
```

It performs a rapid database search of pathnames.

If the search requirement is not so simple, `locate` can be combined
with other tools such as `grep` to design more interesting searches:

```bash
$ locate zip | grep bin
$ locate --regex 'bin/(bz|gz|zip)'
```

> **Where Does The `locate` Database Come From?**

The `locate` database is created by another program named `updatedb`.
Usually, it is run periodically as a cron job; that is,
a task performed at regular intervals by the cron daemon.

Most systems equipped with `locate` run `updatedb` once a day.
Since the database is not updated continuously, you will notice
that very recent files do not show up when using `locate`.

To overcome this, it's possible to run the `updatedb` program manually.

#### `find` - Find Files The Hard Way

While the `locate` program can find a file based solely on its name,
the `find` program searches a given directory (and its subdirectories)
for files based on a variety of attributes.

```bash
$ find ~ # produce a list of our home directory
$ find ~ -type d # directories
$ find ~ -type f # file
$ find ~ -type f -name "*.JPG" -size +1M # file, with .JPG extension, file size larger than 1MB
```

| File Type | Description                   |
| --------- | ----------------------------- |
| b         | Block special device file     |
| c         | Character special device file |
| d         | Directory                     |
| f         | Regular file                  |
| l         | Symbolic link                 |

| Character | Size Unit                                                    |
| --------- | ------------------------------------------------------------ |
| b         | 512-byte blocks. This is the default if no unit is specified |
| c         | Bytes                                                        |
| w         | 2-byte words                                                 |
| k         | Kilobytes (units of 1024 bytes)                              |
| M         | Megabytes (units of 1024 KB)                                 |
| G         | Gigabytes (units of 1024 MB)                                 |

> Another find Tests

| Test             | Description                                                                                                                                                                                                                                               |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -regex '<regex>' |                                                                                                                                                                                                                                                           |
| -cmin n          | last modified exactly n minutes ago, use + or - to specify more or lesser than                                                                                                                                                                            |
| -cnewer file     | last modified more recently than those of file                                                                                                                                                                                                            |
| -ctime           | last modified **n hours ago**                                                                                                                                                                                                                            |
| -empty           | Match empty files and directories                                                                                                                                                                                                                         |
| -group name      | Match file or directories belonging to group. group may be expressed as either a group name or as a numeric group id                                                                                                                                      |
| -iname pattern   | Like the -name test but case insensitive                                                                                                                                                                                                                  |
| -inum n          | Match files with inode number n. This is helpful for finding all the hard links to a particular inode                                                                                                                                                     |
| -mmin n          | last modified n minutes ago                                                                                                                                                                                                                               |
| -mtime n         | last modified n\*24 hours ago                                                                                                                                                                                                                             |
| -name pattern    |                                                                                                                                                                                                                                                           |
| -newer file      | last modified more recently than the specified file. This is very useful when writing shell scripts that perform file backups.Each time you make a backup, update a file, and then use `find` to determine which files have changed since the last update |
| -nouser          | that do not belong to a valid user. This can be used to file files belonging to deleted accounts or to detect activity by attackers                                                                                                                       |
| -nogroup         | that do not belong to a valid group                                                                                                                                                                                                                       |
| -perm mode       |                                                                                                                                                                                                                                                           |
| -samefile name   | Similar to -inum, matches files that share the same inode number as file _name_                                                                                                                                                                           |
| -size n          | Match files of size _n_                                                                                                                                                                                                                                   |
| -type c          | Match files of type _c_                                                                                                                                                                                                                                   |
| -user name       | belong to (case sensitive) user _name_.                                                                                                                                                                                                                   |

##### Operators

`find` provides a way to combine tests using logical operators:

```bash
$ find ~ \( -type f -not -perm 0600 \) -or \( -type d -not -perm 0700 \)
```

Since the parentheses have special meaning to the shell, we must escape them
to prevent the shell from trying to interpret them.

| Results of expr1 | Operator | expr2 is...      |
| ---------------- | -------- | ---------------- |
| True             | -and     | Always performed |
| False            | -and     | Never performed  |
| True             | -or      | Never performed  |
| False            | -or      | Always performed |

##### Predefined Actions

Perform actions based on the search results

| Action        | Description                                                                |
| ------------- | -------------------------------------------------------------------------- |
| -delete       | Delete currently matching file                                             |
| -ls           | Perform the equivalent of ls -dils                                         |
| -print        | Output the full pathname of the matching file. (Default if no other action |
|               | is specified)                                                              |
| -quit         | Quit once a match has been made                                            |

**Warning**: Use extreme caution when using the -delete option. Always test
the command first by substituting the -print action

##### User-Defined Actions

Invoke arbitrary commands:

```bash
-exec command {} ;
-exec rm '{}' ';'
# where {} is a symbolic representation of the current path name
```

By using the `-ok` action in place of `-exec`, the user is prompted
before execution of each specified command:

```bash
find ~ -type f -name 'foo*' -ok ls -l '{}' ';'
```

##### Improving Efficiency

When the `-exec` action is used, it launches a new instance of
the specified command each time a matching file is found.
There are times when we might **prefer to combine all of the search
results and launch a single instance of the command**.

For example:

```bash
# we prefer
ls -l file1 file2
# than
ls -l file1
ls -l file2
```

There are two ways:

- `xargs`

- by `find` feature

  ```bash
  # from
  $ find ~ -type f -name 'foo*' -ok ls -l '{}' ';'
  # to
  $ find ~ -type f -name 'foo*' -ok ls -l '{}' +
  ```

##### `xargs`

It accepts input from standard input and converts it into an
**argument list** for a specified command.

```bash
$ find ~ -type f -name 'foo*' -print | xargs ls -l
```

**Note**: The number of arguments that can be placed into a command line is not
unlimited. When a command line exceeds the maximum length supported by the
system, `xargs` executes the specified command with the max number of
arguments possible and then repeats this process until standard input is
exhausted. To see the maximum size of the command line, execute `xargs --show-limits`

##### Options

| Option           | Description                                                                                                                                                                    |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| -depth           | Direct `find` to process a directory's files before the directory itself                                                                                                       |
| -maxdepth levels | Set the maximum number of levels that find will descend into a directory tree when performing tests and actions                                                                |
| -mindepth levels | Set the minimum number of levels that find will descend into a directory tree when performing tests and actions                                                                |
| -noleaf          | Direct find not to optimize its search based on the assumption that it is searching a Unix-like file system. This is needed when scanning DOS/Windows file systems and CD-ROMs |

### 18 - Archiving And Backup

| Command | Description                               |
| ------- | ----------------------------------------- |
| `gzip`  | Compress or expand files                  |
| `bzip2` | A block sorting file compressor           |
| `tar`   | Tape archiving utility                    |
| `zip`   | Package and compress files                |
| `rsync` | Remote file and directory synchronization |

#### Compressing Files

Data compression is the process of removing redundancy from data.
Imagine an entirely black picture file with the dimensions of 100
pixels by 100 pixels. In terms of data storage (assuming 24 bits,
or 3 bytes per pixel), the image will occupy 30,000 bytes of storage.

An image that is all one color contains entirely redundant data. Instead of
storing a block of data containing 30,000 zeros (black), we could compress
the data into the number 10,000, followed by a zero to represent our data.
-- such a data compression scheme is called _run-length encoding_.

_Compression algorithms_ fall into two general categories: _lossless_ and _lossy_.
Lossless compression preserves all the data contained in the original. This means
that when a file is restored from a compressed version, the restored file is
exactly the same as the original, uncompressed version.

Lossy compression (e.g JPEG, and MP3), removes data as the compression is performed, to allow
more compression to be applied. When a lossy file is restored, it does not match
the original version; rather, it is a close approximation.

##### `gzip` and `gunzip`

**Replaces the original file with a compressed version of the original**.

```bash
$ ls -l /etc > foo.txt
-rw-r--r--@ 1 alexsoon  staff  8644 Dec 21 10:16 foo.txt
$ ls -l foo.*
$ gzip foo.txt
$ ls -l foo.*
-rw-r--r--@ 1 alexsoon  staff  3832 Dec 21 10:16 foo.txt.gz
$ gunzip foo.txt
$ ls -l foo.*
-rw-r--r--@ 1 alexsoon  staff  8644 Dec 21 10:16 foo.txt
```

`gzip` can also be used in interesting ways via standard input and output:

```bash
$ ls -l /etc | gzip > foo.txt.gz
```

If our goal were only to view the contents of a compressed text file, we could
do this:

```bash
$ gunzip -c foo.txt | less
```

Alternatively, there is a program supplied with `gzip`, called `zcat` which
is equivalent to `gunzip -c`.

```bash
$ zcat foo.txt | less
```

Or `zless` program

```bash
$ zless foo.txt.gz
```

| Option  | Description                                                                                                                                                                    |
| ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| -c      | Write output to standard output and keep original files. Also `--stdout` and `--to-stdout`                                                                                     |
| -d      | Decompress. This cause `gzip` to act like `gunzip`. Also `--decompress` or `--uncompress`                                                                                      |
| -f      | Force compression even if a compressed version of the original file already exists                                                                                             |
| -h      | Display usage information                                                                                                                                                      |
| -l      | List compression statistics for each file compressed                                                                                                                           |
| -r      | If one of more arguments on the command line are directories, recursively compress files contained within them                                                                 |
| -t      | Test the integrity of a compressed file                                                                                                                                        |
| -v      | Display verbose message while compressing                                                                                                                                      |
| -number | Set amount of compression. `1` (fastest, least compression) to `9` (slowest, most compression). 1 and 9 may also be expressed as `--fast` and `--best`. The default value is 6 |

```bash
$ gzip foo.txt -l
  compressed uncompressed  ratio uncompressed_name
        3832         8644  55.6% foo.txt
```

##### `bzip2`

`bzip2` program, by Julian Seward, is similar to `gzip`, but uses
a different compression algorithm that achieve higher levels of compression
at the cost of compression speed. A file compressed with `bzip2` is denoted
with the extension `.bz2`:

```bash
$ ls -l /etc > foo.txt
$ ls -l foo.txt
-rw-r--r--@ 1 alexsoon  staff  59710 Dec 21 10:50 foo.txt
$ bzip2 foo.txt
$ ls -l foo.txt.bz2
-rw-r--r--@ 1 alexsoon  staff  7066 Dec 21 10:50 foo.txt.bz2
$ bunzip2 foo.txt.bz2
```

All the options (except for -r) that we discussed in `gzip`.

> **Don't Be Compressive Compulsive**

I occasionally see people attempt to compress a file, that has
already been compressed with an effective compression algorithm,
by doing something like this:

```bash
$ gzip picture.jpg
```

Don't do it. You're probably just wasting time and space! If
you apply compression to a file that is already compressed, you
will actually end up a larger file.

#### Archiving Files

A common file-management task often used in conjunction with compression
is archiving. Archiving is the **process of gathering up many files and
bundling them together into single larger file**.

##### `tar`

Its name, short for `tape archive`, reveals its roots as a tool
for making backup tapes. We often see filenames that end with
the extension `.tar` or `.tgz`, which indicate a plain tar archive
and a gzipped archive, respectively.

There is one caveat of extracting an achieve: unless you are
operating as the superuser, files and directories extracted from
archives take on the ownership of the user performing the restoration,
**rather than the original owner**.

Another interesting behavior of `tar` is the way it handles pathnames
in archives. The default for pathnames is relative rather than absolute.
It removes the leading slash (root) when creating the archive.

```bash
$ tar cf playground2.tar ~/playground
$ cd foo
$ tar xf ../playground2.tar
$ ls
home
$ ls home/me
playground
```

Sometimes it's actually useful, for example imagine we want to
copy the home directory to its contents from one system to another system
and we have a large USB hard drive that we can use for the transfer.

```bash
$ sudo tar cf /media/BigDisk/home.tar /home
# change system..
$ cd /
$ sudo tar xf /media/BigDisk/home.tar
```

```bash
tar mode[options] pathname...
```

When extracting an archive, it's possible to limit what is extracted from
the archive:

- Extracting single file from an archive

  ```bash
  tar xf archive.tar pathname
  # multiple pathnames may be specified
  tar xf archive.tar pathname1 pathname2
  ```

- Wildcards support (GNU version of `tar`, the version most often found in Linux Distribution)
  , normally wildcards are not support.

  ```bash
  $ tar xf archive.tar --wildcards '*/*5*''
  ```

`tar` is often used in conjunction with `find` to produce archives.
In this example, we will use `find` to produce a set of files
to include in an archive:

```bash
$ find playground -name 'file-A' -exec tar rf playground.tar '{}' '+'
```

Using `tar` with `find` is a good way of creating _incremental backups_ of
a directory tree or an entire system. By using `find` to match files
newer than a timestamp file.

```bash
$ find playground -ctime 1 -exec tar rf playground.tar '{}' '+'
```

`tar` can also make use of both standard input and output. If the filename "-"
is specified, it is taken to mean standard input. The `--file-from` option
(which may also be specified as `-T`) causes `tar` to read its list of
pathnames from a file rather than the command line.

```bash
$ cd
$ find playground -name 'file-A' | tar cf - --files-from=. | gzip > playground.tgz
```

We can simplify it to:

```bash
$ find playground -name 'file-A' | tar czf playground.tgz -T - # gzip
$ find playground -name 'file-A' | tar cjf playground.tbz -T - # bzip
```

Imagine that we had two machines, we can:

```bash
$ mkdir remote-stuff
$ cd remote-stuff
$ ssh remote-sys 'tar cf - Documents' | tar xf -
```

| Mode | Description                                              |
| ---- | -------------------------------------------------------- |
| c    | Create an archive from a list of files and/or directores |
| x    | Extract an archieve                                      |
| r    | Append specified pathnames file to the end of an archive |
| t    | List the contents of an archive                          |

```bash
$ tar tf <archived file>
$ tar tvf <archived file> # more detail: show the time...
$ tar cf xx.tar <file>
$ tar zcvf xx.tar.gz <file>
$ cd folder && tar xf ../playground.tar && ls # playground
$ tar zxvf <archived file>
```

| Option | Description                     |
| ------ | ------------------------------- |
| f      | Specify name of the tar archive |

##### `zip`

It is both a compression tool and an archiver. The file format
used by the program is familiar to Windows users: `.zip` files.

```bash
zip options zipfile file...

e.g.
zip -r playground.zip playground
```

**Tip**: Unless we include the `-r` option for recursion, only the directory
(but none of its contents) is stored.

To extract the contents of zip

```bash
$ cd foo
$ unzip ../playground.zip
```

One thing to note about `zip` (as opposed to `tar`) is that if an existing
archive is specified, it is updated rather than replaced. **This means that
the existing archive is preserved, but new files are added and
matching files are replaced.**

Extract selectively from a zip archive:

```bash
$ unzip -l playground.zip playground/dir-087/file-Z
$ unzip playground.zip playground/dir-087/file-Z
```

`-l` to merely list the contents of the archive without extracting the file.
If no file(s) are specified, `unzip` will list all files in the archive.

`-v` option can be added to increase the verbosity of the listing.

Like `tar`, `zip` can make use of standard input and output:

```bash
$ cd
$ find playground -name "file-A" | zip -@ file-A.zip
```

Unfortunately, the `unzip` program does not accept standard input. This
prevents `zip` and `unzip` from being used together to
perform network file coping like `tar`.

However, `zip` and accept standard input:

```bash
$ ls -l /etc/ | zip ls-etc.zip -
```

The `unzip` program allows its output to be sent to standard output:

```bash
$ unzip -p ls-etc.zip | less
```

#### Synchronizing Files And Directories

Synchronizing the directory of local system or a remote system.

In the Unix-like world, the preferred tool for this task is `rsync`, it
use the _rsync remote-update protocol_, which allows `rsync`
to quickly detect the differences between two directories and perform
the minimum amount of copying required to bring them into sync.

```bash
rsync options source destination
```

Where source and destination are one of the following:

- A local file or directory (**either source or destination must be**)

- A remote file or directory in the form of [user@]host:path

- A remote rsync server specified with a URI of \_rsync://[user@]host[:port]/path

**Tip**: Remote-to-remote copying is not supported

```bash
$ rsync -av playground foo
```

| Option     | Description                                                                        |
| ---------- | ---------------------------------------------------------------------------------- |
| `a`        | for archiving -- causing recursion and preservation of file attribute              |
| `v`        | verbose output                                                                     |
| `--delete` | remove files that may have existed on the destination that is not in source device |

We can perform a useful system backup by first creating a directory named
/backup on the external drive, and then use `rsync` to copy the most important
stuff from our system to the external drive:

```bash
$ sudo rsync -av --delete /etc/ /home /usr/local /media/BigDisk/backup
$ mkdir /media/BigDisk/backup
```

In this example, we copied the /etc, /home, and /usr/local directories
from our system to our imaginary storage device.

**Tip**: you can create an alias and add it to `.bashrc` file

##### Using `rsync` Over A Network

```bash
$ sudo rsync -av --delete --rsh=ssh /etc /home /usr/local remote-sys:/backup
```

`--rsh` instruct `rsync` to use the `ssh` program as its remote shell. In this way,
we were able to use an ssh encrypted tunnel to securely transfer the data
from the local system to the remote host.

The second way is by using a `rsync` server. `rsync` can be configured to run
as a daemon and listen to incoming requests for synchronization. This
is often done to allow mirroring of a remote system. For example, Red Hat
Software maintains a large repository of software packages under development
for its Fedora distribution. It is useful for software testers to mirror
this collection during the testing phase of the distribution release cycle.
Since files in the repository change frequently (often more than once a day),
it is desirable to maintain a local mirror by periodic synchronization, rather
than by bulk copying of the repository. We could mirror it using
our local copy of `rsync` and their rsync server like this:

```bash
$ mkdir fedora-devel
$ rsync -av -delete rsync://gtlib.gatech.edu/fedora-linux-core/development/i386/os fedora-devel
```

### 19 - Regular Expressions

#### `grep`

"global regular expression print". It searches text files
for the occurrence of a specified regular expression and outputs
any line containing a match to standard output.

```bash
$ grep [options] regex [file...]

$ ls /usr/bin | grep zip
$ grep bzip dirlist*.txt
```

| Option | Description                                                                       |
| ------ | --------------------------------------------------------------------------------- |
| `-i`   | Ignore case                                                                       |
| `-v`   | Invert match                                                                      |
| `-c`   | Print the number of matches (line)                                                |
| `-l`   | Print the name of each file that contains a match instead of the lines themselves |
| `-L`   | Invert of `-l`                                                                    |
| `-n`   | For multi-file searches, suppress the output of filenames                         |

**Tip**: `zgrep` program provides a front end for `grep`, allowing
it to read compressed files. `zgrep `

```bash
$ zgrep -El 'regex|regular expression' /usr/share/man/man1/*.gz
/usr/share/man/man1/less.1.gz
/usr/share/man/man1/more.1.gz
/usr/share/man/man1/rsync.1.gz
```

#### Metacharacters And Literals

Regular expression metacharacters consist of the following:

`^ $ . [ ] { } - ? * + ( ) | \`

All other characters are considered literals,

#### The Any Character

`.`

```bash
[me@linuxbox ~]$ grep -h '.zip' dirlist*.txt
bunzip2
bzip2
bzip2recover
gunzip
gzip
funzip
gpg-zip
preunzip
prezip
prezip-bin
unzip
unzipsfx
```

**The `zip` program was not found because the dot metacharacter in our
regular expression increased the length of the required match to four characters**

**Tip**: `/usr/share/dict/words` is a dictionary, you can find the words using regex
for crossword puzzle question : `grep -i '^..j.r$' /usr/share/dict/words`

#### Bracket Expressions And Character Classes

Match a single character from a specified set of characters by using
_bracket expressions_.

```bash
$ grep -h '[bg]zip' dirlist*.txt
bzip2
bzip2recover
gzip2
```

##### Negation

With `^`, the remaining characters are taken to be a set of characters
that must not be present at the given characters position.

```bash
$ grep -h '[^bg]zip' dirlist*.txt
bunzip2
gunzip
```

##### Traditional Character Range

```bash
$ grep -h '^[A-Z]' dirlist*.txt
$ grep -h '^[A-Za-z0-9]' dirlist*.txt
```

##### POSIX Character Classes

The traditional character ranges are an easily understood and effective way
to handle the problem of quickly specifying sets of characters.
Unfortunately, they don't always work. Try `ls /usr/sbin/[ABCDWFGHIJKLMNOPQRSTUVWXYZ]*`
and `ls /usr/sbin/[A-Z]*`.

> Why?

When Unix was first developed, it only knew about ASCII characters, and
this feature reflects that fact. In ASCII, the first 32 characters (numbers 0-31)
are control codes (thing like tabs, backspaces, and carriage returns).
The next (32-63) contain printable characters, including
most punctuation characters and the numerals zero through nine. The next
32 (numbers 64-95) contains the uppercase letters and a few more punctuation
symbols. The final 31 (numbers 96-127) contain the lowercase letters and
yet more punctuation symbols.

Based on this arrangement , systems using ASCII used a _collation order_:

`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`

This differs from proper dictionary order, which is like this:

`aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ`

As the popularity of Unix spread beyond the United Stated, there grew a
need to support characters not found in U.S. English. The ASCII
table was expanded to use a full eight bits, adding characters numbers 128-255,
which accommodated many more languages. To support this ability,
the POSIX standards introduced a concept called a _local_, which could
be adjusted to select the character set needed for a particular location.
We can see the language setting of our system using this command:

```bash
$ echo $LANG
en_US.UTF-8
```

**Tip**: `export LANG=POSIX` coverts the system to use ASCII for its character set.
You can also make this change parmanent by adding this line to your .bashrc file

With this setting, POSIX compliant applications will use a dictionary collation
order rather than ASCII order. This explains the behavior of the commands above.
A character range of \[A-Z] when interpreted in dictionary order includes all of
the alphabetic characters except the lowercase "a", hence our results.

To partially work around this problem, the POSIX standard includes a number
of character classes which provide useful range of characters:

| Character Class | Description                                                                              |
| --------------- | ---------------------------------------------------------------------------------------- |
| [:alnum:]       | [A-Za-z0-9]                                                                              |
| [:word:]        | Same with above with an addition of the `_`                                              |
| [:alpha:]       | [A-Za-z]                                                                                 |
| [:blank:]       | Incldues the space and tab characters                                                    |
| [:cntrl:]       | The ASCII control codes. Includes the ASCII characters 0 - 31 and 127                    |
| [:digit:]       | The numerals zero through nine.                                                          |
| [:graph:]       | The visible characters. In ASCII, it includes characters 33 - 126                        |
| [:lower:]       | The lowercase letters                                                                    |
| [:punct:]       | The puntuation characters. In ASCII, equivalent to [-!"#$%&'()\*+,./:;<=>?@[\\\]_`{\|}~] |
| [:print:]       | All the characters in [:graph:] plus the space character                                 |
| [:space:]       | [ \t\r\n\v\f]                                                                            |
| [:upper:]       |                                                                                          |
| [:xdigit:]      | Hexadecimal numbers, [0-9A-Fa-f]                                                         |

**Tip**: Even with the character classes, there is still no convenient way
to express partial ranges like [A-M]

**Important**: This is not example of a regular expression, rather it is
**the shell performing pathname expansion**.

#### POSIX Basic Vs. Extended Regular Expressions

POSIX also splits regular expression implementations into two kinds:

- basic regular expressions

- extended regular expression

The difference between BRE and ERE:

- BRE:

  - ^ $ . [ ] \* are metacharacters

  - Preceding any ERE added metacharacters are treated as a metacharacters in BRE

- ERE:

  - ( ) { } ? + | are added to metacharacters

  - Preceding any metacharacter with a backslash causes it to be treated as a literal.

**Tip**: Traditionally, ERE has been performed by the `egrep` program, but the
GNU version of `grep` also supports ERE when the -E option is used.

> **POSIX**

During the 1980's, Unix became a very popular commercial operating
system, but by 1988, the Unix world was in turmoil. Many computer manufacturers
had licensed the Unix source code from its creators, AT&T, and were supplying
various versions of the operating system with their systems. However, in their
efforts to create product differentiation, each manufacturer added proprietary changes
and extensions. This started to limit the compatibility of the software.

Enter the IEEE (Institute of Electrical and Electronics Engineers). In the
mid 1980s, the IEEE began developing a set of standards that would
define how Unix (and Unix-like) systems would perform. These standards,
formally known as IEEE 1003, define the _application programming interface_ (APIs),
shell and uitilities that are to be found on a standard Unix-like system.
The name POSIX, which stands for Portable Operating System Interface (with the
"X" added to the end for extra snappiness), was suggested by Richard Stallman,
and was adopted by the IEEE.

#### Alternation

`Extended feature` -- There is a must to use grep with `-E` option, or `egrep`

Vertical-bar metacharacter:

```bash
$ echo "AAA" | grep -E 'AAA|BBB'
AAA
$ echo "BBB" | grep -E 'AAA|BBB'
BBB
$ echo "CCC" | grep -E 'AAA|BBB'
```

To combine alternation with other regular expression elements, use ()

```bash
$ grep -Eh '^(bz|gz|zip)' dirlist*.txt
bzip
gzip
zip
zipcloak
```

Below

```bash
$ grep -Eh '^bz|gz|zip' dirlist*.txt
bzip
funzip
gunzip
gzcat
```

#### Quantifiers

`Extended feature`

##### ? - Match An Element Zero Or One Time

```bash
# to match
# (nnn) nnn-nnnn
# or
# nnn nnn-nnnn
$ echo "(555) 123-4567" | grep -E '^\(?[[:digit:]]{3}\)? [[:digit:]]{3}-[[:digit:]]{4}$'
(555) 123-4567
```

##### \* - Match An Element Zero Or More Times

To match a string that starts with an uppercase letter, then
contains any number of upper and lowercase letters and spaces, and ends with a period

```bash
$ echo "This works." | grep -E '[[:upper:]][[:upper:][:lower:] ]*\.'
This works.
```

##### + - Match An Element Or More Times

```bash
$ echo "This that" | grep -E '^([[:alpha:]]+ ?)+$'
This that
$ echo "a b c" | grep -E '^([[:alpha:]]+ ?)+$'
a b c
$ echo "a b 9" | grep -E '^([[:alpha:]]+ ?)+$'
$ echo "abc  d" | grep -E '^([[:alpha:]]+ ?)+$'
```

##### {} - Match An Element A Specific Number Of Times

Express minimum and maximum numbers of required matches.

| Specifier | Meaning                                                                             |
| --------- | ----------------------------------------------------------------------------------- |
| {n}       | Match the preceding element if it occurs exactly n times                            |
| {n,m}     | Match the preceding element if it occurs at least n times, but no more than m times |
| {n,}      | Match the preceding element if it occurs n or more times                            |
| {,m}      | Match the preceding element if it occurs no more that m times                       |

```bash
$ echo "(555) 123-4567" | grep -E '^\(?[[:digit:]]{3}\)? [[:digit:]]{3}-[[:digit:]]{4}$'
```

#### Putting Regular Expressions To Work

##### Validating A Phone List With `grep`

```bash
$ for i in {1..10}; do echo "(${RANDOM:0:3}) ${RANDOM:0:3}-${RANDOM:0:4}" >> phonelist.txt; done
$ grep -Ev '^\(?[0-9]{3}\)? [0-9]{3}-[0-9]{4}$' phonelist.txt
(619) 46-2825
```

##### Find Ugly Filenames With `find`

```bash
$ find . -regex '.*[^-_./0-9a-zA-Z].*'
```

##### Searching For Files With `locate`

```bash
$ locate --regex 'bin/(bz|gz|zip)'
```

### 20 - Text Processing

| Command  | Description                                        |
| -------- | -------------------------------------------------- |
| `cat`    | Concatenate files and print on the standard output |
| `sort`   | Sort lines of text files                           |
| `uniq`   | Report or omit repeated files                      |
| `cut`    | Remove sections from each line of files            |
| `paste`  | Merge lines of files                               |
| `join`   | Join lines of two files on a common field          |
| `comm`   | Compare two sorted files line by line              |
| `diff`   | Compare files line by line                         |
| `patch`  | Apply a diff file to an original                   |
| `tr`     | Translate or delete characters                     |
| `sed`    | Stream editor for filtering and transforming text  |
| `aspell` | Interative spell checker                           |

#### Revisiting Some Old Friends

| Command | Description                                        |
| ------- | -------------------------------------------------- |
| `cat`   | Concatenate files and print on the standard output |
| `sort`  | Sort lines of text files                           |
| `uniq`  | Report or omit repeated files                      |

##### `cat`

`-A` option is used to display non printing characters in the text.
There are times when we want to know if control characters are embedded in
our otherwise visible text.

```bash
$ cat > foo.txt
    The quick brown fox jumped over the lazy dog.
$ cat -A foo.txt
^IThe quick brown fox jumped over the lazy dog.$
$ cat foo.txt
        The quick brown fox jumped over the lazy dog.
```

`^I`, this is a common notation that means "Control-I" which,
as it turns out, is the same as a tab character. `$` indicate that
our text contains trailing spaces.

| Option | Description                                                                    |
| ------ | ------------------------------------------------------------------------------ |
| `s`    | Squeeze multiple adjacent empty lines, causing the output to be single spaced. |
| `n`    | Number the output lines, starting at 1.                                        |

```bash
$ cat > foo.txt
The quick brown fox


jumped over the lazy dog.
$ cat -s foo.txt
The quick brown fox

jumped over the lazy dog.
$ cat -ns foo.txt
     1  The quick brown fox
     2
     3  jumped over the lazy dog.
```

> **MS-DOS Text Vs. Unix Text**

Unix and DOS don't define the end of a line the same way
in text files. Unix ends a line with a linefeed character '\n' (or newline) (ASCII 10)
while MS-DOS and its derivatives use the sequence carriage return `\r` (ASCII 13)
and linefeed to terminate each line of text.

There are several ways to convert files from DOS to Unix format.
On many Linux systems, there are programs called `dos2unix` and `unix2dos`.
However, if you don't have `dos2unix` on your system, don't worry. The
process of converting text from DOS to Unix format is very simple;
it simply involves the removal of the offending carriage returns.

- [with `tr`](#dostounix)

##### `sort`

```bash
$ sort > foo.txt
c
b
a
$ cat foo.txt
a
b
c
```

`sort` can accept multiple files on the command line as arguments,
**it is possible to merge multiple files into single sorted whole.**

```bash
$ sort file1.txt file2.txt file3.txt > final_sorted_list.txt
```

| Option | Long Option             | Description                                                                                                                                                                                                                             |
| ------ | ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -b     | --ignore-leading-blanks | By default, sorting is performed on entire line, starting with the first character in the line. This option cause sort to ignore leading spaces in lines and calculates sorting based on the first non-whitespace character on the line |
| -f     | --ignore-case           | Making sorting case-insensitive                                                                                                                                                                                                         |
| -n     | --numeric-sort          | Allow sorting to be performed on numeric values rather than alphabetic values                                                                                                                                                           |
| -r     | --reverse               |                                                                                                                                                                                                                                         |
| -m     | --merge                 | Treat each argument as the name of a **presorted** file. Merge multiple files into a single sorted result without performing any additional sorting.                                                                                    |
| -o     | --output=file           | Send sorted output to file rather than standard output                                                                                                                                                                                  |
| -t     | --field-separator=char  | Define the field-separator character (Default: spaces or tabs). See below example for more                                                                                                                                              |
| -k     | --key=field[,field2]    | See below example for more                                                                                                                                                                                                              |

```bash
$ ls -l /usr/bin | head
-rwxr-xr-x   1 root   wheel    204592 Sep 16 21:28 AssetCacheLocatorUtil
-rwxr-xr-x   1 root   wheel    243968 Sep 16 21:28 AssetCacheManagerUtil
-rwxr-xr-x   1 root   wheel    205696 Sep 16 21:28 AssetCacheTetheratorUtil
-rwxr-xr-x  76 root   wheel    167136 Sep 16 21:28 DeRez
-rwxr-xr-x  76 root   wheel    167136 Sep 16 21:28 GetFileInfo
-rwxr-xr-x   1 root   wheel    256544 Sep 16 21:28 IOAccelMemory
-rwxr-xr-x   1 root   wheel    194288 Sep 16 21:28 IOMFB_FDR_Loader
-rwxr-xr-x   1 root   wheel    203040 Sep 16 21:28 IOSDebug
-rwxr-xr-x  76 root   wheel    167136 Sep 16 21:28 ResMerger
$ ls -l /usr/bin | sort -nr -k 5 | head
-rwxr-xr-x   1 root   wheel  41606816 Sep 16 21:28 mpsgraphtool
-r-xr-xr-x   1 root   wheel  13113472 Sep 16 21:28 parl5.34
-r-xr-xr-x   1 root   wheel  12389760 Sep 16 21:28 parl5.30
-rwxr-xr-x   1 root   wheel   5672912 Sep 16 21:28 kmutil
-rwxr-xr-x   1 root   wheel   5304192 Sep 16 21:28 vim
-rwxr-xr-x   1 root   wheel   4079680 Sep 16 21:28 sqlite3
-rwxr-xr-x   1 root   wheel   4012032 Sep 16 21:28 dig
-rwxr-xr-x   1 root   wheel   4010592 Sep 16 21:28 delv
-rwxr-xr-x   1 root   wheel   3979120 Sep 16 21:28 nslookup
-rwxr-xr-x   1 root   wheel   3979072 Sep 16 21:28 host
```

Playing with option k to sort multiple key

```bash
$ cat distros.txt
SUSE 10.2   12/07/2006
Fedora 10   11/25/2008
SUSE 11.0   06/19/2008
Ubuntu 8.04 04/24/2008
Fedora 8    11/08/2007
SUSE 10.3   10/04/2007
Ubuntu 6.10 10/26/2006
Fedora 7 05/31/2007
Ubuntu 7.10 10/18/2007
Ubuntu 7.04 04/19/2007
SUSE   10.1 05/11/2006
Fedora 6 10/24/2006
Fedora 9 05/13/2008
Ubuntu 6.06 06/01/2006
Ubuntu 8.10 10/30/2008
Fedora 5 03/20/2006
$ sort --key=1,1 --key=2n distros.txt
Fedora 5    03/20/2006
Fedora 6    10/24/2006
Fedora 7    05/31/2007
Fedora 8    11/08/2007
Fedora 9    05/13/2008
Fedora 10   11/25/2008
SUSE   10.1 05/11/2006
SUSE   10.2 12/07/2006
SUSE   10.3 10/04/2007
SUSE   11.0 06/19/2008
Ubuntu 6.06 06/01/2006
Ubuntu 6.10 10/26/2006
Ubuntu 7.04 04/19/2007
Ubuntu 7.10 10/18/2007
Ubuntu 8.04 04/24/2008
Ubuntu 8.10 10/30/2008
$ sort -k 3.7nbr -k 3.1nbr -k 3.4nbr distros.txt
Fedora 10   11/25/2008
Ubuntu 8.10 10/30/2008
SUSE 11.0   06/19/2008
Fedora 9 05/13/2008
Ubuntu 8.04 04/24/2008
Fedora 8    11/08/2007
Ubuntu 7.10 10/18/2007
SUSE 10.3   10/04/2007
Fedora 7 05/31/2007
Ubuntu 7.04 04/19/2007
SUSE 10.2   12/07/2006
Ubuntu 6.10 10/26/2006
Fedora 6 10/24/2006
Ubuntu 6.06 06/01/2006
SUSE   10.1 05/11/2006
Fedora 5 03/20/2006
```

`k 1,1` specified a range of fields to include in the first key, the digit after comma
is the place of character at that field.

`k 2n` field 2 is the sort key and that sort should be numeric

If no range is specified, `sort` uses a key that begins with the specified
field and extends to the **end of the line**.

Some files don't use tab and spaces as field delimiters:

```bash
$ head /etc/passwd
root:x:0:0:root:/root:/bin/bash
daemon:x:1:1:daemon:/usr/sbin:/bin/sh
bin:x:2:2:bin:/bin:/bin/sh
sys:x:3:3:sys:/dev:/bin/sh
sync:x:4:65534:sync:/bin:/bin/sync
games:x:5:60:games:/usr/games:/bin/sh
man:x:6:12:man:/var/cache/man:/bin/sh
lp:x:7:7:lp:/var/spool/lpd:/bin/sh
mail:x:8:8:mail:/var/mail:/bin/sh
news:x:9:9:news:/var/spool/news:/bin/sh
$ sort -t ':' -k 7 /etc/passwd
me:x:1001:1001:Myself,,,:/home/me:/bin/bash
root:x:0:0:root:/root:/bin/bash
dhcp:x:101:102::/nonexistent:/bin/false
gdm:x:106:114:Gnome Display Manager:/var/lib/gdm:/bin/false
hplip:x:104:7:HPLIP system user,,,:/var/run/hplip:/bin/false
klog:x:103:104::/home/klog:/bin/false
messagebus:x:108:119::/var/run/dbus:/bin/false
polkituser:x:110:122:PolicyKit,,,:/var/run/PolicyKit:/bin/fals
```

##### `uniq`

**When given a sorted file (include standard input)**, it removes
any duplicate lines sends the results to standard output.

**Tip**: While `uniq` is a traditional Unix tool often used with `sort`,
the GNU version of `sort` supports **a `-u` option**, which removes duplicates
from the sorted output.

```bash
$ uniq foo.txt
```

**Important**: `uniq` only removes duplicate lines which are adjacent
to each other.

| Option | Description                                                                                                                                                                            |
| ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -c     | Output a list of duplicate lines preceded by the number of times the line occurs                                                                                                       |
| -d     | Only output repeated lines                                                                                                                                                             |
| -u     | Only output unique lines                                                                                                                                                               |
| -f n   | Ignore n leading fields in each line. Fields are separated by white space as they are in `sort`; however, unlike `sort`, `uniq` has no option for setting an alternate field separator |
| -i     | Ignore case during the line comparisons                                                                                                                                                |
| -s n   | Skip(ignore) the leading n characters of each line                                                                                                                                     |

#### Slicing And Dicing

| Command | Description                               |
| ------- | ----------------------------------------- |
| `cut`   | Remove sections from each line of files   |
| `paste` | Merge lines of files                      |
| `join`  | Join lines of two files on a common field |

##### `cut`

Extract a section of text from a line and output the extracted section
to standard output. It can accept multiple file arguments or input
from standard input.

| Option        | Description                                                                                                                                                     |
| ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -c char_list  | Extract the portion of the line defined by char_list. The list may consist of one or more comma-separated numerical ranges. (e.g `7-10`, `23-`)                     |
| -f field_list | Extract one or more fields from the line as defined by field_list. The list may contain one or more fields or field ranges separated by commas. (e.g `1` `1,2`) |
| -d delim_char | When -f is specified, use _delim_char_ as the field delimiting character. By default, fields must be separated by a single tab character                        |
| --complement  | Extract the entire line of text, except for those portions specified by -c and/or -f.                                                                           |

```bash
$ cat -A distros.txt
SUSE^I10.2^I12/07/2006$
Fedora^I10^I11/25/2008$
SUSE^I11.0^I06/19/2008$
Ubuntu^I8.04^I04/24/2008$
Fedora^I8^I11/08/2007$
SUSE^I10.3^I10/04/2007$
Ubuntu^I6.10^I10/26/2006$
Fedora^I7^I05/31/2007$
Ubuntu^I7.10^I10/18/2007$
Ubuntu^I7.04^I04/19/2007$
SUSE^I10.1^I05/11/2006$
Fedora^I6^I10/24/2006$
Fedora^I9^I05/13/2008$
Ubuntu^I6.06^I06/01/2006$
Ubuntu^I8.10^I10/30/2008$
Fedora^I5^I03/20/2006$
$ cut -f 3 distros.txt
12/07/2006
11/25/2008
06/19/2008
04/24/2008
11/08/2007
10/04/2007
10/26/2006
05/31/2007
10/18/2007
04/19/2007
05/11/2006
10/24/2006
05/13/2008
06/01/2006
10/30/2008
03/20/2006
$ cut -f 3 distros.txt | cut -c 7-10
2006
2008
2008
2008
2007
2007
2006
2007
2007
2007
2006
2006
2008
2006
2008
2006
```

```bash
$ cut -d ':' -f 1 /etc/passwd |head
root
daemon
bin
sys
sync
games
man
lp
mail
news
```

> **Expanding Tabs**

If we wanted a file that could be fully manipulated with cut by characters rather than fields, this would require us to replace the tab characters within the
file with the corresponding number of spaces: a tool `expand` accept either
one or more file arguments or standard input, and output the modified
text to standard output.

```bash
# expand replaced tabs to spaces
$ expand distros.txt | cut -c 23-
```

##### `paste`

Does the opposite of `cut`. Rather than extracting a column of text
from file, it adds one or more columns of text to a file. It does this
by reading multiple files and combining the fields found in each file
into a single stream on standard output.

```bash
$ cat distros-by-date.txt
Fedora  10      11/25/2008
Ubuntu  8.10    10/30/2008
SUSE    11.0    06/19/2008
Fedora  9       05/13/2008
Ubuntu  8.04    04/24/2008
Fedora  8       11/08/2007
Ubuntu  7.10    10/18/2007
SUSE    10.3    10/04/2007
Fedora  7       05/31/2007
Ubuntu  7.04    04/19/2007
SUSE    10.2    12/07/2006
Ubuntu  6.10    10/26/2006
Fedora  6       10/24/2006
Ubuntu  6.06    06/01/2006
SUSE    10.1    05/11/2006
Fedora  5       03/20/2006
$ cut -f 1,2 distros-by-date.txt > distros-versions.txt
$ cut -f 3 distros-by-date.txt > distros-dates.txt
$ paste distros-dates.txt > distros-versions.txt
12/07/2006      SUSE    10.2
11/25/2008      Fedora  10
06/19/2008      SUSE    11.0
04/24/2008      Ubuntu  8.04
11/08/2007      Fedora  8
10/04/2007      SUSE    10.3
10/26/2006      Ubuntu  6.10
05/31/2007      Fedora  7
10/18/2007      Ubuntu  7.10
04/19/2007      Ubuntu  7.04
05/11/2006      SUSE    10.1
10/24/2006      Fedora  6
05/13/2008      Fedora  9
06/01/2006      Ubuntu  6.06
10/30/2008      Ubuntu  8.10
03/20/2006      Fedora  5
```

##### `join`

`join` joins data from multiple files based on shared key field.

```bash
$ cat distros-by-date.txt
Fedora  10      11/25/2008
Ubuntu  8.10    10/30/2008
SUSE    11.0    06/19/2008
Fedora  9       05/13/2008
Ubuntu  8.04    04/24/2008
Fedora  8       11/08/2007
Ubuntu  7.10    10/18/2007
SUSE    10.3    10/04/2007
Fedora  7       05/31/2007
Ubuntu  7.04    04/19/2007
SUSE    10.2    12/07/2006
Ubuntu  6.10    10/26/2006
Fedora  6       10/24/2006
Ubuntu  6.06    06/01/2006
SUSE    10.1    05/11/2006
Fedora  5       03/20/2006
$ cut -f 3 distros-by-date.txt > distros-dates.txt
$ cut -f 2,2 distros-by-date.txt > distros-versions.txt
$ cut -f 1,1 distros-by-date.txt > distros-names.txt
$ paste distros-dates.txt distros-names.txt > distros-key-names.txt
$ paste distros-dates.txt distros-versions.txt > distros-key-versions.txt
$ join distros-key-names.txt distros-key-versions.txt
12/07/2006 SUSE 10.2
11/25/2008 Fedora 10
06/19/2008 SUSE 11.0
04/24/2008 Ubuntu 8.04
11/08/2007 Fedora 8
10/04/2007 SUSE 10.3
10/26/2006 Ubuntu 6.10
05/31/2007 Fedora 7
10/18/2007 Ubuntu 7.10
04/19/2007 Ubuntu 7.04
05/11/2006 SUSE 10.1
10/24/2006 Fedora 6
05/13/2008 Fedora 9
06/01/2006 Ubuntu 6.06
10/30/2008 Ubuntu 8.10
03/20/2006 Fedora 5
```

**Note**: By default, `join` uses whitespace as the input field
delimiter and a single space as the output field delimiter. This
behavior can be modified by specifying options.

#### Comparing Text

| Command | Description                           |
| ------- | ------------------------------------- |
| `comm`  | Compare two sorted files line by line |
| `diff`  | Compare files line by line            |
| `patch` | Apply a diff file to an original      |

##### `comm`

Compares two text files and displays the lines that are unique to each one
and the lines they have in common.

```bash
$ cat > file1.txt
a
b
c
d
$ cat > file2.txt
b
c
d
e
$ comm file1.txt file2.txt
a
                b
                c
                d
        e
```

`comm` supports options in the form `-n`, where n is either 1, 2, or 3, these
options specify which column(s) to suppress:

```bash
$ comm -12 file1.txt file2.txt
b
c
d
```

##### `diff`

`diff` is also used to to detect the differences between files -- in more
complex ways and support many output formats and abilities.

One common use for `diff` is the creation of `diff` files or `patches`
that are used by programs such as `patch`

```bash
$ diff file1.txt file2.txt
1d0
< a
4a4
> e
```

**Where old_file and new_file are either single files or directories containing files**

| Change | Description                                                                                                 |
| ------ | ----------------------------------------------------------------------------------------------------------- |
| r1ar2  | Append the lines at the position 2 in the second file to the position r1 in the first file                  |
| r1cr2  | Change (replace) the lines at position r1 with the lines at the position r2 in the second file.             |
| r1dr2  | Delete the lines in the first file at position r1, which would have appeared at range r2 in the second file |

While this format is the default, it is not as widely used as other, optional formats: `context format` and `unified format`.

> **Context Format**

```bash
$ diff -c file1.txt file2.txt
*** file1.txt   2024-01-02 08:25:18.902221584 +0000
--- file2.txt   2024-01-02 08:25:27.518259157 +0000
***************
*** 1,4 ****
- a
  b
  c
  d
--- 1,4 ----
  b
  c
  d
+ e
```

`*** 1,4 ***` indicates lines 1 through 4 in the first file.

`--- 1,4 ---` indicates lines 1 through 4 in the second file.

| Indicator | Meaning                                                             |
| --------- | ------------------------------------------------------------------- |
| blank     | No difference between two files                                     |
| -         | A line deleted -- appear in first file but not in the second file   |
| +         | A line added -- appear in the second file but not in the first file |
| !         | A line changed                                                      |

> **Unified Format**

```bash
$ diff -u file1.txt file2.txt
--- file1.txt   2024-01-02 08:25:18.902221584 +0000
+++ file2.txt   2024-01-02 08:25:27.518259157 +0000
@@ -1,4 +1,4 @@
-a
 b
 c
 d
+e
```

| Character | Meaning                                                             |
| --------- | ------------------------------------------------------------------- |
| blank     | No difference between two files                                     |
| -         | A line deleted -- appear in first file but not in the second file   |
| +         | A line added -- appear in the second file but not in the first file |

##### `patch`

Apply changes to text files. The _diff_ file contains the change from
the previous version to the new versions. Using `diff/patch` offers
two significant advantages:

1. The diff file is very small, compared to the full size of the source tree
2. The diff file concisely shows the change being made, allowing reviewers
   of patch to quickly evaluate it.

To prepare diff file:

`diff -Naur old_file new_file > diff_file`

Patch the old file into new file:

`patch < diff_file`

```bash
$ diff -Naur file1.txt file2.txt > patchfile.txt
$ patch < patchfile.txt
$ cat file1.txt
b
c
d
e
```

`-r` - supports recursion of a directory tree.

#### Editing On The Fly

| Command  | Description                                       |
| -------- | ------------------------------------------------- |
| `tr`     | Translate or delete characters                    |
| `sed`    | Stream editor for filtering and transforming text |
| `aspell` | Interative spell checker                          |

##### `tr`

Transliterate characters -- character-based search-and-replace operation.

```bash
$ echo "lowercase letters" | tr a-z A-Z
```

Character sets may be expressed in one of three ways:

1. An enumerated list: e.g ABCDWFGHIJKLMNOPQRSTUVWXYZ

2. A character range: e.g. A-Z

   - Note that this method is sometimes subject to the same issues as other commands due to the locale collation order, and thus should be used with caution.

3. POSIX character classes: e.g [:upper:]

In most cases, both character sets should be of equal length; however, it is possible
for the first set to be larger than the second: if we wish to convert multiple
characters to a single character:

```bash
$ echo "lowercase letters" | tr [:lower:] A
AAAAAAAAA AAAAAAA
```

<a id="dostounix"></a>

In addition to transliteration, `tr` allows characters to simply be deleted from
the input stream. For example, to remove MS.DOS text files carriage return characters at
the end of each line:

```bash
$ tr -d '\r' < dos_file > unix_file
```

`tr` can "squeeze" (delete) repeated adjoining instances of a character.

```bash
$ echo "aaabbbccc" | tr -s ab
abccc
$ echo "abcabcabc" | tr -s ab
abcabcabc
```

> **ROT13: The Not-So-Secret Decoder Ring**

One amusing use of `tr` is to perform `ROT13` encoding of text.
It is a trivial type of encryption based on a simple substitution
cipher. It is used sometimes on text to obscure potentially offensive content.
The method simply moves each character 13 places up the alphabet. Since
this is half way up the possible 26 characters, performing the algorithm a
second time on the text restores it to its original form.

```bash
$ echo "secret text" | tr a-zA-Z n-za-mN-ZA-M
frperg grkg
$ echo "frperg grkg" | tr a-zA-Z n-za-mN-ZA-M
secret text
```

##### `sed`

_Stream editor_. `sed` is a powerful and somewhat complex program
(there are entire books about it), so we will not cover it completely here.

**`-i`: to rewrite the file with the changes applied rather than
sending the edited output to standard output**

**`-E`: Use extended regular expressions. e.g `sed -E 's/[0-9]+//g' input-file`**

```bash
$ echo "front" | sed 's/front/back/'
back
```

By convention, the slash character as a delimiter. The choice of the
delimiter character is arbitrary:

```bash
$ echo "front" | sed 's_front_back_'
back
```

Most commands in `sed` may be preceded by an address, which specifies which line(s)
of the input stream will be edited. If the address is omitted, then the
editing command is carried out on every line in the input stream. The simplest form
of address is a line number.

```bash
$ echo "front\nfront" | sed '2s/front/back/'
front
back
```

| Address      | Description                                                                                                    |
| ------------ | -------------------------------------------------------------------------------------------------------------- |
| n            | A line number where n is a positive integer                                                                    |
| $            | The last line                                                                                                  |
| /regexp/     | Lines matching a POSIX basic regular expression. Note that the regular expression                              |
|              | is delimited by slash characters. Optionally, the regular expression may be delimited                          |
|              | by an alternate character, by specifying the expression with `\cregexpc`, where `c` is the alternate character |
| addr1, addr2 | A range, addr2 is inclusive. Addresses may be any of the single address forms above                            |
| first~step   | Matched the line represented by the number first, then each subsequent                                         |
|              | line at step intervals. e.g `1~2` refers to each odd numbered line.                                            |
| addr1,+n     | Match addr1 and the following n lines                                                                          |
| addr!        | Match all lines except `addr`, which may be any of the forms above                                             |

Examples:

```bash
$ sed -n '1,5p' distros.txt
SUSE    10.2    12/07/2006
Fedora  10      11/25/2008
SUSE    11.0    06/19/2008
Ubuntu  8.04    04/24/2008
Fedora  8       11/08/2007
```

`p` command causes a matched line to be printed. -- print only line 1 to line 5

`n` no auto-print option, to cause `sed` not to print every line by default

```bash
$ sed -n '/SUSE/p' distros.txt
SUSE    10.2    12/07/2006
SUSE    11.0    06/19/2008
SUSE    10.3    10/04/2007
SUSE    10.1    05/11/2006
```

Above is the used of regex

```bash
$ sed -n '/SUSE/!p' distros.txt
Fedora  10      11/25/2008
Ubuntu  8.04    04/24/2008
Fedora  8       11/08/2007
Ubuntu  6.10    10/26/2006
Fedora  7       05/31/2007
Ubuntu  7.10    10/18/2007
Ubuntu  7.04    04/19/2007
Fedora  6       10/24/2006
Fedora  9       05/13/2008
Ubuntu  6.06    06/01/2006
Ubuntu  8.10    10/30/2008
Fedora  5       03/20/2006
```

Here is more commands:

| Command                | Description                                                                                                                                   |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `=`                    | Output current line number                                                                                                                    |
| `a`                    | Append text after the current line.                                                                                                           |
| `d`                    | Delete the current line                                                                                                                       |
| `i`                    | insert text in front of the current line.                                                                                                     |
| `p`                    | Print the current line. By default, `sed` prints every line and only edits                                                                    |
|                        | lines that matched a specified address within the file. The default behavior                                                                  |
|                        | can be overridden by specifying the `-n` option.                                                                                              |
| `q`                    | Exit `sed` without processng any more lines. If the `-n` option is not specified, output the current line                                     |
| `Q`                    | Exit `sed` without processing any more lines                                                                                                  |
| `s/regexp/replacement` | Substitute the contents of replacement wherever regexp is found                                                                               |
|                        | `replacement` may include the special character &, which is equivalent to the text matched by regexp.                                         |
|                        | In addition, `replacement` may include the sequences \1 through \9, which are the contents of the corresponding subexpressions in regexp      |
| `y/set1/set2`          | Perform transliteration by converting characters from set1 to                                                                                 |
|                        | corresponding characters in set2. **Unlike `tr`, `sed` requires that both sets be of same length**. And it does not support character ranges. |

Examples:

```bash
$ cat distros.txt
SUSE    10.2    12/07/2006
Fedora  10      11/25/2008
SUSE    11.0    06/19/2008
Ubuntu  8.04    04/24/2008
Fedora  8       11/08/2007
SUSE    10.3    10/04/2007
Ubuntu  6.10    10/26/2006
Fedora  7       05/31/2007
Ubuntu  7.10    10/18/2007
Ubuntu  7.04    04/19/2007
SUSE    10.1    05/11/2006
Fedora  6       10/24/2006
Fedora  9       05/13/2008
Ubuntu  6.06    06/01/2006
Ubuntu  8.10    10/30/2008
Fedora  5       03/20/2006
$ sed 's/\([0-9]\{2\}\)\/\([0-9]\{2\}\)\/\([0-9]\{4\}\)$/\3-\1-\2/' distros.txt
SUSE    10.2    2006-12-07
Fedora  10      2008-11-25
SUSE    11.0    2008-06-19
Ubuntu  8.04    2008-04-24
Fedora  8       2007-11-08
SUSE    10.3    2007-10-04
Ubuntu  6.10    2006-10-26
Fedora  7       2007-05-31
Ubuntu  7.10    2007-10-18
Ubuntu  7.04    2007-04-19
SUSE    10.1    2006-05-11
Fedora  6       2006-10-24
Fedora  9       2008-05-13
Ubuntu  6.06    2006-06-01
Ubuntu  8.10    2008-10-30
Fedora  5       2006-03-20
```

It is also possible to construct more complex commands in a script file using the `-f` option.

```js
$ cat -n distros.sed
     1  # sed script to produce Linux distributions report
     2
     3  1 i\
     4  \
     5  Linux Distributions Report\
     6
     7  s/\([0-9]\{2\}\)\/\([0-9]\{2\}\)\/\([0-9]\{4\}\)$/\3-\1-\2/
     8  y/abcdefghijklmnopqrstuvwxyz/ABCDEFGHIJKLMNOPQRSTUVWXYZ/
$ sed -f distros.sed distros.txt

Linux Distributions Report

SUSE    10.2    2006-12-07
FEDORA  10      2008-11-25
SUSE    11.0    2008-06-19
UBUNTU  8.04    2008-04-24
FEDORA  8       2007-11-08
SUSE    10.3    2007-10-04
UBUNTU  6.10    2006-10-26
FEDORA  7       2007-05-31
UBUNTU  7.10    2007-10-18
UBUNTU  7.04    2007-04-19
SUSE    10.1    2006-05-11
FEDORA  6       2006-10-24
FEDORA  9       2008-05-13
UBUNTU  6.06    2006-06-01
UBUNTU  8.10    2008-10-30
FEDORA  5       2006-03-20
```

**Note**: A line-continuation character is formed by a backslash followed immediately
by a carriage return. No intermediary spaces are permitted.

Line 6 is actually the end of our inserted text and ends with a plain
carriage return rather than a line-continuation character, signaling the end of the
`i` command.

> **People Who Like `sed` Also Like...**

`sed` is most often used for simple, one-line tasks rather than long
scripts. Many users prefer other tools for larger tasks. The most popular
of these are `awk` and `perl`. While both `awk` and `perl` are outside the scope
of this book, there are very good skills for the Linux command line user.

**Tip**: Also try googling "sed one liners", "sed cheat sheets"

##### `aspell`

An interactive spelling checker.

```bash
$ aspell check textfile
```

```bash
$ cat foo.txt
The quick brown fox jimped over the laxy dog.
$ aspell check foo.txt
$ cat foo.txt
The quick brown fox jumped over the lazy dog.
```

**Unless told otherwise via the command line option `--dont-backup`, `aspell`
creates a backup file containing the original text by appending the extention `.bak`
to the file name.**

```bash
$ cat foo.txt.bak
The quick brown fox jimped over the laxy dog.
```

`aspell` will see the contents of the HTML tags as misspelled. This problem can be
overcome by include `-H` (HTML) checking-mode option: `aspell -H check foo.txt`

**Note**: By default, `aspell` will ignore URLs and email addresses in text.
This behavior can be overridden with command line options. It is also possible to specify
which markup tags are checked and skipped. See the `aspell` man page for details.

### 21 - Formatting Output

| Command  | Description                          |
| -------- | ------------------------------------ |
| `nl`     | Number lines                         |
| `fold`   | Wrap each line to a specified length |
| `fmt`    | A simple text formatter              |
| `pr`     | Prepare text for printing            |
| `printf` | Format and print data                |
| `groff`  | A document formatting system         |

#### Simple Formatting Tools

##### `nl` - Number Lines

It numbers lines. In its simplest use, it resembles `cat -n`:

```bash
$ nl distros.txt | head
     1  SUSE    10.2    12/07/2006
     2  Fedora  10      11/25/2008
     3  SUSE    11.0    06/19/2008
     4  Ubuntu  8.04    04/24/2008
     5  Fedora  8       11/08/2007
     6  SUSE    10.3    10/04/2007
     7  Ubuntu  6.10    10/26/2006
     8  Fedora  7       05/31/2007
     9  Ubuntu  7.10    10/18/2007
    10  Ubuntu  7.04    04/19/2007
```

Like `cat`, `nl` can accept either multiple files as command line arguments,
or standard input.

`nl` supports a concept called "logical pages" when numbering:

| Markup    | Meaning                      |
| --------- | ---------------------------- |
| \\:\\:\\: | Start of logical page header |
| \\:\\:    | Start of logical page body   |
| \\:       | Start of logical page footer |

| Option    | Meaning                                                                                            |
| --------- | -------------------------------------------------------------------------------------------------- |
| -b style  | Set body numbering to style, where style is one of the following                                   |
|           | `a` = number all lines                                                                             |
|           | `t` = number only non-blank lines. This is the default                                             |
|           | `n` = none                                                                                         |
|           | `pregexp` = number only lines matching basci regular expression `regexp`                           |
| -f style  | Set footer numbering to style. Defaults is n (none)                                                |
| -h style  | Set header numbering to style. Default is n (none)                                                 |
| -i number | Set page numbering increment to number. Default is one                                             |
| -n format | Sets numbering format to format, where the format is                                               |
|           | `ln` = left justified, without leading zeros                                                       |
|           | `rn` = right justified, without leading zeros. This is the default                                 |
|           | `rz` = right justified, with leading zeros                                                         |
| -p        | Do not reset page numbering at the beginning of each logical page                                  |
| -s string | Add string to the end of each line number to create a separator. Default is a single tab character |
| -v number | Set first line number of each logical page to _number_. Default is one.                            |
| -w width  | Set width of the line number field to width. Default is 6                                          |

```bash
$ cat distros-nl.sed
# sed script to produce Linux distributions report

1 i\
\\:\\:\\:\
\
Linux Distributions Report\
\
Name       Ver.   Released\
----       ----   --------\
\\:\\:
s/\([0-9]\{2\}\)\/\([0-9]\{2\}\)\/\([0-9]\{4\}\)$/\3-\1-\2/
$ a\
\\:\
\
End Of Report
```

```bash
$ sort -k 1,1 -k 2n distros.txt | sed -f distros-nl.sed | nl


       Linux Distributions Report

       Name       Ver.   Released
       ----       ----   --------

     1  Fedora  5       2006-03-20
     2  Fedora  6       2006-10-24
     3  Fedora  7       2007-05-31
     4  Fedora  8       2007-11-08
     5  Fedora  9       2008-05-13
     6  Fedora  10      2008-11-25
     7  SUSE    10.1    2006-05-11
     8  SUSE    10.2    2006-12-07
     9  SUSE    10.3    2007-10-04
    10  SUSE    11.0    2008-06-19
    11  Ubuntu  6.06    2006-06-01
    12  Ubuntu  6.10    2006-10-26
    13  Ubuntu  7.04    2007-04-19
    14  Ubuntu  7.10    2007-10-18
    15  Ubuntu  8.04    2008-04-24
    16  Ubuntu  8.10    2008-10-30


       End Of Report
```

We can repeat the command and experiment with different options for nl. Some interesting ones are:

```bash
$ sort -k 1,1 -k 2n distros.txt | sed -f distros-nl.sed | nl -n rz


       Linux Distributions Report

       Name       Ver.   Released
       ----       ----   --------

000001  Fedora  5       2006-03-20
000002  Fedora  6       2006-10-24
000003  Fedora  7       2007-05-31
000004  Fedora  8       2007-11-08
000005  Fedora  9       2008-05-13
000006  Fedora  10      2008-11-25
000007  SUSE    10.1    2006-05-11
000008  SUSE    10.2    2006-12-07
000009  SUSE    10.3    2007-10-04
000010  SUSE    11.0    2008-06-19
000011  Ubuntu  6.06    2006-06-01
000012  Ubuntu  6.10    2006-10-26
000013  Ubuntu  7.04    2007-04-19
000014  Ubuntu  7.10    2007-10-18
000015  Ubuntu  8.04    2008-04-24
000016  Ubuntu  8.10    2008-10-30


       End Of Report
$ sort -k 1,1 -k 2n distros.txt | sed -f distros-nl.sed | nl -w 3 -s ' '


    Linux Distributions Report

    Name       Ver.   Released
    ----       ----   --------

  1 Fedora      5       2006-03-20
  2 Fedora      6       2006-10-24
  3 Fedora      7       2007-05-31
  4 Fedora      8       2007-11-08
  5 Fedora      9       2008-05-13
  6 Fedora      10      2008-11-25
  7 SUSE        10.1    2006-05-11
  8 SUSE        10.2    2006-12-07
  9 SUSE        10.3    2007-10-04
 10 SUSE        11.0    2008-06-19
 11 Ubuntu      6.06    2006-06-01
 12 Ubuntu      6.10    2006-10-26
 13 Ubuntu      7.04    2007-04-19
 14 Ubuntu      7.10    2007-10-18
 15 Ubuntu      8.04    2008-04-24
 16 Ubuntu      8.10    2008-10-30


    End Of Report
```

##### `fold` - Wrap Each Line To A Specified Length

Folding is the process of breaking lines of text at a specified width.

```bash
$ echo "The quick brown fox jumped over the lazy dog." | fold -w 12
The quick br
own fox jump
ed over the
lazy dog.
```

`-s` - cause fold to break the line at the last available space before the
line width is reached:

```bash
$ echo "The quick brown fox jumped over the lazy dog." | fold -w 12 -s
The quick
brown fox
jumped over
the lazy
dog.
```

##### `fmt` - A Simple Text Formatter

The `fmt` program also fold text, plus a lot more. Basically, it fills and
joins lines in text while preserving blank lines and indentation

Reformat to fit a fifty character wide column

```bash
$ info fmt | fmt -w 50 | head -n 20
File: coreutils.info,  Node: fmt invocation,
Next: pr invocation,  Up: Formatting file
contents

4.1 ‘fmt’: Reformat paragraph text
==================================

‘fmt’ fills and joins lines to produce output
lines of (at most) a given number of characters
(75 by default).  Synopsis:

     fmt [OPTION]... [FILE]...

   ‘fmt’ reads from the specified FILE
   arguments (or standard input if
none are given), and writes to standard output.

   By default, blank lines, spaces between words,
   and indentation are
preserved in the output; successive input lines
```

`fmt` is preserving the indentation of the first line. It provides an
option to correct this: `c`

```bash
$ info fmt | fmt -cw 50 | head -n 20
File: coreutils.info,  Node: fmt invocation,
Next: pr invocation,  Up: Formatting file
contents

4.1 ‘fmt’: Reformat paragraph text
==================================

‘fmt’ fills and joins lines to produce output
lines of (at most) a given number of characters
(75 by default).  Synopsis:

     fmt [OPTION]... [FILE]...

   ‘fmt’ reads from the specified FILE
arguments (or standard input if none are given),
and writes to standard output.

   By default, blank lines, spaces between words,
and indentation are preserved in the output;
successive input lines with different indentation
```

Other interesting options:

| Option    | Description                                                                                                                                                                                                                                                                                                                         |
| --------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -c        | Operate in crown margin mode. This preseves the indentation of the first two lines                                                                                                                                                                                                                                                  |
|           | of a paragraph. Subsequent lines are aligned with the indentation of the second line                                                                                                                                                                                                                                                |
| -p string | Only format those line beginning with the prefix string.                                                                                                                                                                                                                                                                            |
|           | This option can be used to format text in source code comments. For example , `-p #`                                                                                                                                                                                                                                                |
| -s        | Split-only mode. Lines will only be split to fit the specified column width.                                                                                                                                                                                                                                                        |
|           | Short lines will not be joined to fill lines. This mode is useful when                                                                                                                                                                                                                                                              |
|           | formatting text such as code where joining is not desired.                                                                                                                                                                                                                                                                          |
| -u        | Performing uniform spacing. This will apply traditional "typewriter-style" ormatting to the text. This means a single space between words and two spaces between sentences. This mode is useful for removing “justification,” that is, text that has been padded with spaces to force alignment on both the left and right margins. |
| -w width  | Format text to fit within a column width characters wide. The default is 75 characters. Note: fmt actually formats lines slightly shorter than the specified width to allow for line balancing.                                                                                                                                     |

Examples:

```bash
$ cat fmt-code.txt
# This file contains code with comments.

# This line is a comment.
# Followed by another comment line.
# And another.

This, on the other hand, is a line of code.
And another line of code.
And another.
$ fmt -w 50 -p '# ' fmt-code.txt
# This file contains code with comments.

# This line is a comment.  Followed by another
# comment line.  And another.

This, on the other hand, is a line of code.
And another line of code.
And another.
```

##### `pr` - Format Text For Printing

It is used to _paginate text_. It separates each page with several lines of
white space and creates a default header containing the file modification time,
filename, and page number.

```bash
$ pr -l 15 -w 65 distros.txt


2024-01-08 02:35                distros.txt                Page 1


SUSE    10.2    12/07/2006
Fedora  10      11/25/2008
SUSE    11.0    06/19/2008
Ubuntu  8.04    04/24/2008
Fedora  8       11/08/2007







2024-01-08 02:35                distros.txt                Page 2


SUSE    10.3    10/04/2007
Ubuntu  6.10    10/26/2006
Fedora  7       05/31/2007
Ubuntu  7.10    10/18/2007
Ubuntu  7.04    04/19/2007

...
```

`-l` for page length

`-w` for page width

View [~~here~~, I skipped this chapter..](#printing) for more options to control page layout.

##### `printf` - Format And Print Data

"print formatted", was originally developed for the C programming language
and has been implemented in many programming languages including the shell.
In fact, in bash, `printf` is a builtin.

Unlike the other commands in this chapter, the printf command is not used for
pipelines (it does not accept standard input) nor does it find frequent application
directly on the command line (it’s mostly used in scripts). So
why is it important? Because it is so widely used.

```bash
$ printf "format" arguments
$ printf "I formatted the string: %s\n" foo
I formatted the string: foo
```

| Specifier | Descripion                                          |
| --------- | --------------------------------------------------- |
| d         | Signed decimal interger                             |
| f         | Floating point number                               |
| o         | Octal number                                        |
| s         | String                                              |
| x         | Hexadecimal number using lowercase a-f where needed |
| X         | Same as x but use uppercase letters                 |
| %         | Print literal % symbol (i.e. "%%")                  |

```bash
$ printf "%d, %f, %o, %s, %x, %X\n" 380 380 380 380 380 380
380, 380.000000, 574, 380, 17c, 17C
```

If more arguments provided, it will be multiple output:

```bash
$ printf "%d, %f, %o, %s, %x, %X\n" 380 380 380 380 380 380 380
380, 380.000000, 574, 380, 17c, 17C
380, 0.000000, 0, , 0, 0
```

Several optional components may be added to the conversion specifier to
adjust its output:

```bash
%[flags][width].[precision]conversion_specification
```

| Component  | Description                                                                                                                                                                                                   |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| flags      | There are five different flags:                                                                                                                                                                               |
|            | # - Use the "alternate format" for output. This varies by datatype. For `o` (octal number) conversion, the output is prefixed with 0. For `x` and `X`, the output is prefixed with `0x` and `0X` respectively |
|            | `0` - Pad the output with zeros. This means that the field will be filled with leading zeors, as in "000380"                                                                                                  |
|            | `-` - Left-align the output, by default, `printf` right-aligns output                                                                                                                                         |
|            | ` ` (space) - produce a leading space for positive number                                                                                                                                                     |
|            | `+` - Sign positive number, by default, `printf` signs negative numbers                                                                                                                                       |
| width      | A number specifying the minimum field width                                                                                                                                                                   |
| .precision | For floating point numbers it is decimal point. For string, it is the number of characters to output                                                                                                          |

Here are some examples :

| Argument    | Format    | Result     |
| ----------- | --------- | ---------- |
| 380         | "%#x"     | 0x17c      |
| 380         | "%05d"    | 00380      |
| 380         | "%05.5f"  | 380.00000  |
| 380         | "%010.5f" | 0380.00000 |
| 380         | "%+d"     | +380       |
| 380         | "%-d"     | 380        |
| abcdefghijk | "%5s"     | abcdefjhik |
| abcdefghijk | "%.5s"    | abcde      |

Some numbers with neat formatting:

```bash
$ printf "Line: %05d %15.3f Result: %+15d\n" 1071 3.14156295 32589
Line: 01071           3.142 Result:          +32589
```

#### Document Formatting Systems

##### `groff`

`Skip for now`

<a id="printing"></a>

### 22 - Printing

`Skip for now`

### 23 - Compiling Programs

`Skip for now`

## 4 - Writing Shell Scripts

### 24 - Writing Your First Script

#### What Are Shell Scripts

In the simplest terms, a shell script is a file containing a series
of commands. The shell reads this file and carries out the commands as though
they have been entered directly on the command line.

#### How To Write A Shell Script

To successfully create and run a shell script, we need to do three things:

1. Write a script

2. Make the script executable

3. Put the script somewhere the shell can find it

#### Script File Format

The first line of our script is a little mysterious. It looks as if it should be
a comment, since it starts with #, but it looks too purposeful to be just that.
The `#!` character sequence is, in fact, a special construct called a _shebang_.
It is used to tell the system the name of the interpreter that should be used to execute
the script that follows. **Every shell script should include this as its first line**.

> Name: `hello_world`

```bash
#!/bin/bash

# This is our first script.

echo 'Hello World!'
```

#### Executable Permissions

There are two common permission settings for scripts; 755 for scripts that everyone can execute, and 700 for scripts that only the owner can execute. Note that scripts must be readable in order to be executed.

```bash
$ chmod 755 hello_world
$ chmod u+x hello_world
```

#### Script File Location

```bash
$ ./hello_world
Hello World!
$ hello_world
Command 'hello_world' not found
```

We can create `bin` directory and place our script within it,
it should start to work like other program

```bash
$ vi .bashrc
...
export PATH=~/bin:"$PATH"
...
$ . .bashrc
```

The dot(.) command is a synonym for the `source` command, a shell builtin
which reads a specified file of shell commands and treats it like input
from the keyboard.

##### Good Locations For Scripts

**The `~/bin` directory is a good place to put scripts intended for personal use**.
If we write a script that everyone on a system is allowed to use, the traditional
location is `/usr/local/bin`. Scripts intended for use by the system administrator
are often located in `/usr/local/sbin`. In most cases,
locally supplied software, whether scripts or compiled programs, should be placed in the
`/usr/local` hierarchy and not in `/bin` or `/usr/bin`. These directories are specified
by the Linux Filesystem Hierarchy Standard to contain only files supplied and maintained by
the Linux distributor.

#### More Formatting Tricks

One of the key goals of serious script writing is ease of maintenance

##### Long Option Names

In the interests of reduced typing, short options are preferred when entering options on the command line, but when writing scripts, long options can provide improved readability.

```bash
$ ls --all --directory
# is better than below in script
$ ls -ad
```

##### Indentation And Line-Continuation

When employing long commands, readability can be enhanced by spreading the com- mand over several lines.

```bash
find playground \
  \( \
    -type f \
    -not -perm 0600 \
    -exec chmod 0600 '{}' ';' \
  \) \
  -or \
  \( \
    -type d \
    -not -perm 0700 \
    -exec chmod 0700 '{}' ';' \
  \)

# is better than
find playground \( -type f -not -perm 0600 -exec chmod 0600 '{}' ';' \) -or \( -type d -not -perm 0700 -exec chmod 0700 '{}' ';' \)
```

> **Configuring `vim` For Script Writing**

`:syntax on`

Turns on syntax highlighting

`:set hlsearch`

Turns on the option to highlight search results.

`:set tabstop=4`

Sets the number of columns occupied by a tab character. The default is 8 columns.

`:set autoindent`

Causes `vim` to indent a new line the same amount as the line just typed. To
stop indentation, type `Ctrl-d`

These changes can be made permanent by adding these commands (without the leading colon characters) to your ~/.vimrc file.

### 25 - Starting A Project

**Tip**: you can use `file:///home/username/foo.html` to view the file in browser. `firefox <file>`

```bash
#!/bin/bash

# Program to output a system information page

echo "<HTML>"
echo "  <HEAD>"
echo "    <TITLE>Page Title</TITLE>"
echo "  </HEAD>"
echo "  <BODY>"
echo "    Page body."
echo "  </BODY>"
echo "</HTML>"
```

We can actually combine all the echo commands:

```bash
#!/bin/bash

# Program to output a system information page

echo "<HTML>
  <HEAD>
    <TITLE>Page Title</TITLE>
  </HEAD>
  <BODY>
    Page body.
  </BODY>
</HTML>"
```

#### Variables And Constants

```bash
#!/bin/bash

# Program to output a system information page

title="System Information Report"

echo "<HTML>
  <HEAD>
    <TITLE>$title</TITLE>
  </HEAD>
  <BODY>
    <h1>$title</h1>
  </BODY>
</HTML>"
```

There are some rules about variable names:

- Variable names may consist of alphanumeric characters (letters and numbers)
  and underscore characters

- The first character of a variable name must be either a letter or an underscore.

- Spaces and punctuation symbols are not allowed.

A common convention is to use **uppercase letters to designate constants** and
**lower case letters for true variables**. Although the shell makes no distinction
between variables and constants.

```bash
#!/bin/bash

# Program to output a system information page

TITLE="System Information Report"

echo "<HTML>
  <HEAD>
    <TITLE>$TITLE</TITLE>
  </HEAD>
  <BODY>
    <h1>$TITLE</h1>
  </BODY>
</HTML>"
```

**Note**: The shell actually does provide a way to enforce the immutability
of constants, through the use of the `declare` builtin command with the `-r` (read-only)
option:

```bash
declare -r TITLE="Page Title"
```

The shell would prevent any subsequent assignment to `TITLE`. **This feature is
rarely used, but it exists for very formal scripts.**

##### Assigning Values To Variables And Constants

```bash
variable=value
```

**Note that in an assignment, there must be no spaces between the variable name,
the equals sign, and the value.**

So what can the value consist of? Anything that we can expand into a string:

```bash
a=z                     # Assign the string "z"
b="a string"            # Embeded spaces must be within quotes.
c="a string and $b"     # Other expansions such as variables can be expanded into the assignments
c="a string and ${b}andc"     # Other expansions such as variables can be expanded into the assignments
d=$(ls -l foo.txt)      # Results of a command.
e=$((5*7))              # Arithmetic expansion.
f="\t\ta string\n"      # Escape sequences such as tabs and newlines.
```

Multiple variable assignments may be done on a single line:

```bash
a=5 b="a string"
```

**During expansion, variable names may be surrounded by optional curly braces "{}".**
This is useful in cases where a variable name becomes ambiguous due to its
surrounding context.

```bash
$ hello=hello
$ echo ${hello}world
helloworld
```

```bash
#!/bin/bash

# Program to output a system information page

TITLE="System Information Report For $HOSTNAME"
CURRENT_TIME=$(date +"%x %r %Z")
TIMESTAMP="Generated $CURRENT_TIME, by $USER"

echo "<HTML>
  <HEAD>
    <TITLE>$TITLE</TITLE>
  </HEAD>
  <BODY>
    <h1>$TITLE</h1>
    <p>$TIMESTAMP</p>
  </BODY>
</HTML>"
```

#### Here Documents

We've looked at two different methods of outputting our text, both
using the `echo` command. There is third way called a _here document or here script_.

```bash
command << token
text
token
```

Where command is the name of command that accepts standard input and token
is a string used to indicate the end of the embedded text:

```bash
#!/bin/bash

# Program to output a system information page

TITLE="System Information Report For $HOSTNAME"
CURRENT_TIME=$(date +"%x %r %Z")
TIMESTAMP="Generated $CURRENT_TIME, by $USER"

cat << _EOF_
<HTML>
  <HEAD>
    <TITLE>$TITLE</TITLE>
  </HEAD>
  <BODY>
    <h1>$TITLE</h1>
    <p>$TIMESTAMP</p>
  </BODY>
</HTML>
_EOF_
```

**Important**: The token must appear alone and that there must not be
trailing spaces on the line.

So what's the advantage of using a here document? It's mostly the
same as `echo`, except that, by default, single and double quotes within
here documents **lose their special meaning to the shell**.:

```bash
$ foo="some text"
$ cat << _EOF_
> $foo
> "$foo"
> '$foo'
> \$foo
> ${foo}
some text
"some text"
'some text'
$foo
some text
```

**Here documents can be used with any command that accepts standard input.**
Below uses here document to pass a series of commands to the `ftp` program
in order to retrieve a file from a remote FTP server:

```bash
#!/bin/bash

# Script to retrieve a file via FTP

FTP_SERVER=ftp.nl.debian.org
FTP_PATH=/debian/dists/bookworm/main/installer-i386/current/images/cdrom
REMOTE_FILE=debian-cd_info.tar.gz

ftp -n << _EOF_
open $FTP_SERVER
user anonymous me@linuxbox
cd $FTP_PATH
hash
get $REMOTE_FILE
bye
_EOF_
ls -l $REMOTE_FILE
```

If we change the direction redirection operator from "<<" to "<<-", the
shell will ignore leading tab characters in the here document.
This allows a here document to be indented (tabs! Not space!), which can improve readability:

```bash
#!/bin/bash

# Script to retrieve a file via FTP

FTP_SERVER=ftp.nl.debian.org
FTP_PATH=/debian/dists/bookworm/main/installer-i386/current/images/cdrom
REMOTE_FILE=debian-cd_info.tar.gz

ftp -n << _EOF_
    open $FTP_SERVER
    user anonymous me@linuxbox
    cd $FTP_PATH
    hash
    get $REMOTE_FILE
    bye
_EOF_
ls -l $REMOTE_FILE
```
### 26 - Top-Down Design

The process of identifying the top-level steps and developing
increasingly detailed views of those steps is called _top-down design_.
This technique allows us to break large complex tasks into many small,
simple tasks. Top-down design is a common method of designing programs
and one that is well suited to shell programming in particular.

#### Shell Functions

**Tip**: `$FUNCTION` variable is has the function name

Our script currently performs the following steps to generate the HTML document:

1. Open page.
2. Open page header.
3. Set page title.
4. Close page header.
5. Open page body.
6. Output page heading.
7. Output timestamp.
8. Close page body.
9. Close page.

For our next stage of development, we will add some tasks between steps 7 and 8. These will include:

- System uptime and load. This is the amount of time since the last shutdown or re- boot and the average number of tasks currently running on the processor over sev- eral time intervals.

- Disk space. The overall use of space on the system’s storage devices.

- Home space. The amount of storage space being used by each user.

If we had a command for each of these tasks, we could add them to our script
simply through command substitution:

```bash
#!/bin/bash

# Program to output a system information page

TITLE="System Information Report For $HOSTNAME"
CURRENT_TIME=$(date +"%x %r %Z")
TIMESTAMP="Generated $CURRENT_TIME, by $USER"

cat << _EOF_
<HTML>
  <HEAD>
    <TITLE>$TITLE</TITLE>
  </HEAD>
  <BODY>
    <H1>$TITLE</H1>
    <P>$TIMESTAMP</P>
    $(report_uptime)
    $(report_disk_space)
    $(report_home_space)
  </BODY>
</HTML>
_EOF_
```

We could create these additional commands two ways:

- Write three separate scripts and place them in a directory listed in our PATH.

- Embed the scripts within our program as _shell functions_.

  - Shell functions are "mini-scripts" that are located inside other scripts and can
    act as autonomous programs.

  - Two syntactic forms:

  - ```bash
    # 1
    function name {
      commands return
    }

    # 2
    name() {
      commands
      return
    }
    ```

```bash
#!/bin/bash

# Shell function demo

function funct {
  echo "Step 2"
  return
}

funct3 () {
  echo "Step 3"
  return
}

echo "Step 1" # => Step 1
funct # => Step 2
funct3 # => Step 3
```

**Important**: shell function definitions must appear in the script
before they are called.

To call function inside a text:

```bash
#!/bin/bash

# Program to output a system information page

TITLE="System Information Report For $HOSTNAME"
CURRENT_TIME=$(date +"%x %r %Z")
TIMESTAMP="Generated $CURRENT_TIME, by $USER"

report_uptime() {
	echo "Function report_uptime executed."
	return
}
report_disk_space() {
	echo "Function report_disk_space executed."
	return
}
report_home_space() {
	echo "Function report_home_space executed."
	return
}

cat <<_EOF_
<HTML>
  <HEAD>
    <TITLE>$TITLE</TITLE>
  </HEAD>
  <BODY>
    <H1>$TITLE</H1>
    <P>$TIMESTAMP</P>
    $(report_uptime)
    $(report_disk_space)
    $(report_home_space)
  </BODY>
</HTML>
_EOF_
```

#### Local Variables

Inside shell functions, it is often desirable to have local variables.
Local variables are only accessible within the shell function in which
they are defined and cease to exist once the shell function
terminates.

```bash
#!/bin/bash

# local-vars: script to demonstrate local variables

foo=0 # global var

funct_1() {
  local foo
  foo=1
  echo "funct_1: foo = $foo"
}

funct_2() {
  local foo
  foo=2
  echo "funct_2: foo = $foo"
}

echo "global: foo = $foo" # => global: foo = 0
funct_1 # => funct_1: foo = 1
echo "global: foo = $foo" # => global: foo = 0
funct_2 # => funct_2: foo = 2
echo "global: foo = $foo" # => global: foo = 0
```

#### Keep Scripts Running

```bash
#!/bin/bash

# Program to output a system information page

TITLE="System Information Report For $HOSTNAME"
CURRENT_TIME=$(date +"%x %r %Z")
TIMESTAMP="Generated $CURRENT_TIME, by $USER"

report_uptime() {
    cat <<- _EOF_
        <H2>System Uptime</H2>
        <PRE>$(uptime)</PRE>
        _EOF_
	return
}
report_disk_space() {
    cat <<- _EOF_
        <H2>Disk Space Utilization</H2>
        <PRE>$(df -h)</PRE>
        _EOF_
	return
}
report_home_space() {
    cat <<- _EOF_
        <H2>Home Space Utilization</H2>
        <PRE>$(du -sh /home/*)</PRE>
        _EOF_
	return
}

cat << _EOF_
<HTML>
  <HEAD>
    <TITLE>$TITLE</TITLE>
  </HEAD>
  <BODY>
    <H1>$TITLE</H1>
    <P>$TIMESTAMP</P>
    $(report_uptime)
    $(report_disk_space)
    $(report_home_space)
  </BODY>
</HTML>
_EOF_
```

**Tip**: We use the du command with the -sh options to perform
this task. This, however, is not a complete solution to the problem.
While it will work on some systems (Ubuntu, for example), it
will not work on others. The reason is that many systems set the permissions
of home directories to prevent them from being world-readable, which is a
reasonable security measure. Below to solve it:

```bash
report_home_space() {
  if [[ $(id -u) -eq 0 ]]; then
    cat <<- _EOF_
      <H2>Home Space Utilization (All Users)</H2>
      <PRE>$(du -sh /home/*)</PRE>
      _EOF_
  else
    cat <<- _EOF_
      <H2>Home Space Utilization ($USER)</H2>
      <PRE>$(du -sh $HOME)</PRE>
      _EOF_
  fi
  return
}
```

> **Shell Functions In Your `.bashrc` File**

**Shell functions make excellent replacements for aliases, and are
actually the preferred method of creating small commands for personal use**.

### 27 - Flow Control: Branching With if

#### If

```bash
x=5

if [ $x -eq 5 ]; then
  echo "x equals 5."
else
  echo "x does not equal 5."
fi
```

```bash
if commands; then
  commands
[elif commands; then
  commands...]
[else
  commands]
fi
```

#### Exit Status

`$?` to view the previous command exit status

Commands (including the scripts and shell functions we write)
issue a value to the system when they terminate, called an _exit status_.
This value, which is an integer in the range of 0 to 255, indicates
the success or failure of the command's execution. By convention,
**a value of zero indicates success and any other value indicates failure**.

```bash
$ ls -d /usr/bin
/usr/bin
$ echo $?
0
$ ls -d /bin/usr
ls: cannot access '/bin/usr': No such file or directory
2
```

Man pages often include a section entitled "Exit Status", describing what codes
are used.

**Tip**: `true` always executes successfully while `false` always executes unsuccessfully.

```bash
$ true
$ echo $?
0
$ false
$ echo $?
1
$ if true; then echo "It's true."; fi
It's true.
$ if false; then echo "It's true."; fi
$ if ls /usr/bin > /dev/null; then echo "hello"; fi
hello
$ if ls /bin/usr 2> /dev/null; then echo "asd"; fi
```

If a list of commands follows if, the last command in the list is evaluated:

```bash
$ if false; true; then echo "It's true."; fi
It's true.
$ if true; false; then echo "It's true."; fi
```

#### `test`

By far, the command used most frequently with `if` is `test`.

It has two equivalent forms:

```bash
# 1
test expression
# 2 **More popular**
[ expression ]
```

The `test` command returns an exit status of zero when the expression is true
and a status of one when the expression is false.

##### File Expressions

| Expression      | Is true If                                                                                                                                   |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| file1 -ef file2 | same inode numbers                                                                                                                           |
| file1 -nt file2 | newer than                                                                                                                                   |
| file1 -ot file2 | older than                                                                                                                                   |
| -b file         | file exists and is a block-special (device) file                                                                                             |
| -c file         | file exists and is a character-special (device) file                                                                                         |
| -d file         | file exists and is a directory                                                                                                               |
| -f file         | file exists and is a regular file                                                                                                            |
| -e file         | file exists                                                                                                                                  |
| -g file         | file exists and is set-group-ID.                                                                                                             |
| -G file         | file exists and is owned by effective group ID                                                                                               |
| -k file         | file exists and has its "sticky bit" set                                                                                                     |
| -L file         | file exists and is a symbolic link                                                                                                           |
| -O file         | file exists and is owned by the effective user ID                                                                                            |
| -p file         | file exists and is a named pipe                                                                                                              |
| -r file         | file exists and is readable (has readable permission for the effective user)                                                                 |
| -s file         | file exists and has a length greater than zero                                                                                               |
| -S file         | file exists and is a network socket                                                                                                          |
| -t fd           | fd is a file descriptor directed to/from the terminal. This can be used to determine whether standard input/output/error is being redirected |
| -u file         | file exists and is setuid                                                                                                                    |
| -w file         | file exists and is writable (has write permission for the effective user)                                                                    |
| -x file         | file exists and is executable (had execute/search permission for the effective user)                                                         |

Here we have a script that demonstrates some of the file expressions:

```bash
#!/bin/bash

# test-file: Evaluate the status of a file

FILE=~/.bashrc

if [ -e "$FILE" ]; then
  if [ -f "$FILE" ]; then
    echo "$FILE is a regular file."
  fi
  if [ -d "$FILE" ]; then
    echo "$FILE is a directory."
  fi
  if [ -r "$FILE" ]; then
    echo "$FILE is readable."
  fi
  if [ -w "$FILE" ]; then
    echo "$FILE is writable."
  fi
  if [ -x "$FILE" ]; then
    echo "$FILE is executale/searchable."
  fi
else
  echo "$FILE does not exist"
  exit 1
fi

exit
```

**Important**: \$FILE is quoted within expression. It is not required,
but is a defense against the parameter being empty. If the parameter expansion
\$FILE were to result in an empty value, **it would cause an error (the
operators would be interpreted as non-null strings rather than operator)**.
Using the quotes around the parameter ensures that the operator is always
followed by a string, even if the string is empty.

```bash
$ if [ -e $INVALIDVAR ]; then echo "entered!"; fi
entered!
# because it sees if [ -e ], and -e is a non null string
```

**Tip**: If no argument pass to `exit` statement, it defaults to the **exit status
of the last command executed**. The exit command appearing on the last line of
the script is there as a formality. When a script “runs off the end”
(reaches end of file), it terminates with an exit status of the last command executed by default, anyway.

Similarly, shell functions can return an exit status by including an integer
argument to the `return` command:

```bash
test_file() {
    # test-file: Evaluate the status of a file

    FILE=~/.bashrc

    if [ -e "$FILE" ]; then
      if [ -f "$FILE" ]; then
        echo "$FILE is a regular file."
      fi
      if [ -d "$FILE" ]; then
        echo "$FILE is a directory."
      fi
      if [ -r "$FILE" ]; then
        echo "$FILE is readable."
      fi
      if [ -w "$FILE" ]; then
        echo "$FILE is writable."
      fi
      if [ -x "$FILE" ]; then
        echo "$FILE is executale/searchable."
      fi
    else
      echo "$FILE does not exist"
      return 1
    fi
}
```

##### String Expression

| Expression                               | Is True If...                                                                                                                   |
| ---------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| string                                   | string is not null                                                                                                              |
| -n string                                | The length of string is greater than zero                                                                                       |
| -z string                                | The length of string is zero                                                                                                    |
| string1 = string2<br/>string1 == string2 | string1 and string2 are equal. Single or double equal signs may be used, but the use of double equal signs is greatly preferred |
| string1 != string2                       | not equal                                                                                                                       |
| string1 > string2                        | string1 sorts after string2                                                                                                     |
| string1 < string2                        | string1 sorts before string2                                                                                                    |

**Warning**: The `>` and `<` expression operators must be quoted (or escaped
with a backslash) when used with `test` (both methods). If they are not, they will be interpreted
by the shell as redirection operators, with potentially destructive results. Also
note that while the `bash` documentation states that the sorting order conforms
to the collation order of the current locale, **it does not**. ASCII (POSIX) order
is used in versions of `bash` up to and including 4.0.

```bash
#!/bin/bash

# test-string: evaluate the value of a string

ANSWER=maybe

if [ -z "$ANSWER" ]; then
  echo "There is no answer." >&2
  exit 1
fi

if [ "$ANSWER" = "yes" ]; then
  echo "The answer is YES."
elif [ "$ANSWER" = "no" ]; then
  echo "The answer is NO."
elif [ "$ANSWER" = "maybe" ]; then
  echo "The answer is MAYBE."
else
  echo "The answer is UNKNOWN."
fi
```

**Important**: `>&2` redirects the error message to standard error,
which is the proper thing to do with error messages.

##### Integer Expressions

| Expression            | Is True If...                                 |
| --------------------- | --------------------------------------------- |
| integer1 -eq integer2 | integer1 is equal to integer2                 |
| integer1 -ne integer2 | integer1 is not equal to integer2             |
| integer1 -le integer2 | integer1 is less than or equal to integer2    |
| integer1 -lt integer2 | integer1 is less than to integer2             |
| integer1 -ge integer2 | integer1 is greater than or equal to integer2 |
| integer1 -gt integer2 | integer1 is greater than to integer2          |

```bash
#!/bin/bash

# test-integer: evaluate the value of an integer.

INT=-5

if [ -z "$INT" ]; then
  echo "INT is empty." >&2
  exit 1
fi

if [ $INT -eq 0 ]; then
  echo "INT is zero."
else
  if [ $INT -lt 0 ]; then
    echo "INT is negative"
  else
    echo "INT is positive"
  fi
  if [ $((INT %2)) -eq 0 ]; then
    echo "INT is even"
  else
    echo "INT is odd."
  fi
fi
```

**Tip**: The script would fail if the constant INT contained anything except an
integer (float fails also). The script needs a way to verify that the constant
contains an integer. See below for `[[ ]]` expression.

#### A More Modern Version Of `test`

Recent version of `bash` include a compound command that acts as an enhanced replacement
for `test`. It uses the following syntax:

```bash
[[ expression ]]
```

The `[[ ]]` command is very similar to `test` (it supports all of its expressions),
but adds an important new string expression:

`string1 =~ regex`

**Tip**: It is okay to expand variable without $ in `[[ ]]`: eg. `[[ INT -lt 0 ]]`

```bash
#!/bin/bash

# test-integer2: evaluate the value of an integer.

INT=-5

if [[ "$INT" =~ ^-?[0-9]+$ ]]; then
  if [ $INT -eq 0 ]; then
    echo "INT is zero."
  else
    if [ $INT -lt 0 ]; then
      echo "INT is negative"
    else
      echo "INT is positive"
    fi
    if [ $((INT %2)) -eq 0 ]; then
      echo "INT is even"
    else
      echo "INT is odd."
    fi
  fi
else
  echo "INT is not an integer." >&2
  exit 1
fi
```

Another added feature of `[[ ]]` is that the `==` operator supports pattern matching
the same way pathname expansion does. For example:

```bash
$ FILE=foo.bar
$ if [[ $FILE == foo.* ]]; then echo "$FILE matches pattern 'foo.*'"; fi
foo.bar matches pattern 'foo.*'
```

This makes `[[ ]]` useful for evaluating file and pathnames.

#### (( )) - Designed For Integers

It supports a full set of arithmetic evaluations. [See Chapter 34](#ch34)

`(( ))` is used to perform arithmetic truth tests. An arithmetic truth test results
in true if result of the arithmetic evaluation **is non-zero**.

```bash
$ if ((1)); then echo "It is true."; fi
It is true.
$ if ((0)); then echo "It is true."; fi
```

```bash
#!/bin/bash

# test-integer2: evaluate the value of an integer.

INT=-5

if [[ "$INT" =~ ^-?[0-9]+$ ]]; then
  if ((INT == 0)); then
    echo "INT is zero."
  else
    if ((INT <= 0)); then
      echo "INT is negative"
    else
      echo "INT is positive"
    fi
    if ((INT%2 == 0)); then
      echo "INT is even"
    else
      echo "INT is odd."
    fi
  fi
else
  echo "INT is not an integer." >&2
  exit 1
fi
```

**Tip**: because the compound command (( )) is part of the shell syntax
rather than an ordinary command, and it deals only with integers, it is
able to recognize variables by name and does not require expansion to be performed.

#### Combining Expressions

| Operation | test | `[[ ]]` and `(( ))` |
| --------- | ---- | ------------------- |
| AND       | -a   | &&                  |
| OR        | -o   | \|\|                |
| NOT       | !    | !                   |

```bash
#!/bin/bash

# test-integer3: determine if an integer is within a
# specified range of values.

MIN_VAL=1
MAX_VAL=100

INT=50

if [[ "$INT" =~ ^-?[0-9]+$ ]]; then
  if [[ INT -ge MIN_VAL && INT -le MAX_VAL ]]; then
    echo "$INT is within $MIN_VAL to $MAX_VAL."
  else
    echo "$INT is out of range."
  fi
else
  echo "INT is not an integer." >&2
  exit 1
fi
```

We could have also coded this using `test`:

```bash
if [ INT -ge MIN_VAL -a INT -le MAX_VAL ]; then
    echo "$INT is within $MIN_VAL to $MAX_VAL."
else
    echo "$INT is out of range."
fi
```

The `!` negation operator reverses the outcome of an expression.

```bash
#!/bin/bash

# test-integer4: determine if an integer is outside a
# specified range of values.

MIN_VAL=1
MAX_VAL=100

INT=50

if [[ "$INT" =~ ^-?[0-9]+$ ]]; then
  if [[ ! (INT -ge MIN_VAL && INT -le MAX_VAL) ]]; then
    echo "$INT is outside $MIN_VAL to $MAX_VAL."
  else
    echo "$INT is in range."
  fi
else
  echo "INT is not an integer." >&2
  exit 1
fi
```

**Coding this with `test` would be done in this way**:

```bash
if [ ! \( $INT -ge $MIN_VAL -a $INT -le $MAX_VAL \) ]; then
  echo "$INT is outside $MIN_VAL to $MAX_VAL."
else
  echo "$INT is in range"
fi
```

**Since all expressions and operators used by `test` are treated as command
arguments by the shell (unlike `[[ ]]` and `(( ))`), characters which have
special meaning to `bash`, such as `>`, `<`, `(`, and `)`, must be quoted or escaped.**

> **`test` or `[[ ]]`?**

Seeing that `test` and `[[ ]]` do roughly the same thing, which is preferable?
`test` is traditional (and part of POSIX), whereas `[[ ]]` is specific to
bash. It’s important to know how to use `test`, since it
is very widely used, but `[[ ]]` is clearly more useful and is easier to code.

> **Portability Is the Hobgoblin Of Little Minds**

If you talk to “real” Unix people, you quickly discover that many of them don’t like Linux very much. They regard it as impure and unclean. One tenet of Unix followers is that everything should be “portable.” This means that any script you write should be able to run, unchanged, on any Unix-like system.

Unix people have good reason to believe this. Having seen what proprietary extensions to commands and shells did to the Unix world before POSIX, they are naturally wary of the effect of Linux on their beloved OS.

But portability has a serious downside. It prevents progress. It requires that things are always done using “lowest common denominator” techniques. In the case of shell programming, it means making everything compatible with sh, the original Bourne shell.

This downside is the excuse that proprietary vendors use to justify their proprietary extensions, only they call them “innovations.” But they are really just lock-in devices for their customers.

**The GNU tools, such as bash, have no such restrictions. They encourage portability by supporting standards and by being universally available. You can install bash and the other GNU tools on almost any kind of system, even Windows, without cost. So feel free to use all the features of bash. It’s really portable.**

#### Control Operators: Another Way To Branch

`bash` provides two control operators that can perform branching.

```bash
command1 && command2
# and
command1 || command2
```

```bash
$ mkdir temp && cd temp
$ [ -d temp ] || mkdir temp
```

The first command create a directory named `temp`, and if it succeeds, the
current working directory will be changed to `temp`.

The second command is attempted only if the `mkdir` command is successful. This
type of construct is very handy for handling errors:

```bash
[ -d temp ] || exit 1
```

### 28 - Reading Keyboard Input

#### `read` - Read Values From Standard Input

```bash
read [-options] [variables...]
```

_variable_ is the name of one or more variables used to hold the input value.
If no variable name is supplied, **the shell variable REPLY contains the line of data.**

```bash
#!/bin/bash

# read-integer: evaluate the value of an integer.

echo -n "Please enter an integer -> "
read int

if [[ "$int" =~ ^-?[0-9]+$ ]]; then
  if [ $int -eq 0 ]; then
    echo "int is zero."
  else
    if [ $int -lt 0 ]; then
      echo "$int is negative"
    else
      echo "$int is positive"
    fi
    if [ $((int %2)) -eq 0 ]; then
      echo "$int is even"
    else
      echo "$int is odd."
    fi
  fi
else
  echo "input value is not an integer." >&2
  exit 1
fi
```

`read` can assign input to multiple variables:

```bash
#!/bin/bash

# read-multiple: read multiple values from keyboard

echo -n "Enter one or more values > "
read var1 var2 var3 var4 var5

echo "var1 = '$var1'"
echo "var2 = '$var2'"
echo "var3 = '$var3'"
echo "var4 = '$var4'"
echo "var5 = '$var5'"

# === Result ===
$ read-multiple
Enter one or more values > 1 2 3
var1 = '1'
var2 = '2'
var3 = '3'
var4 = ''
var5 = ''
$ read-multiple
Enter one or more values > 1 2 3 4 5 a b c
var1 = '1'
var2 = '2'
var3 = '3'
var4 = '4'
var5 = '5 a b c'
```

If no variables are listed after the `read` command, a shell variable,
`REPLY`, will be assigned all the input:

```bash
#!/bin/bash

# read-single: read multiple values into default variable

echo -n "Enter one or more values > "
read

echo "REPLY = '$REPLY'"

# === Result ===
Enter one or more values > a b c d
REPLY = 'a b c d'
```

##### Options

| Option       | Description                                                                                                                                      |
| ------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| -a array     | Assign the input to _array_, starting with index zero                                                                                            |
| -d delimiter | The first character in the string delimiter is used to indicate end of input, rather than a newline character                                    |
| -e           | Use Readline to handle input. This permits input editing in the same manner as the command line                                                  |
| -i string    | Use string as a default reply if the user simply presses Enter. Requires the `-e` option                                                         |
| -n num       | Read num characters of input, rather than an entire line                                                                                         |
| -p prompt    | Display a prompt for input using the string prompt.                                                                                              |
| -r           | Raw mode. Do not interpret backslash characters as escapes                                                                                       |
| -s           | Silent mode. Do no echo characters to the display as they are typed. This is useful when inputting passwords and other confidential information. |
| -t seconds   | Timeout. Terminate input after seconds. `read` returns a non-zero exit status if an input times out.                                             |
| -u fd        | Use input from file descriptor `fd`, rather than standard input                                                                                  |

```bash
#!/bin/bash

# read-single: read multiple values into default variable

read -p "Enter on or more value >"

echo "REPLY = '$REPLY'"
```

With `-t` and `-s`, we can write a script that reads "secret" input and times out if
the input is not completed in a specified time:

```bash
#!/bin/bash

# read-secret: input a secret passphrase

if read -t 10 -sp "Enter secret passphrase > " secret_pass; then
  echo -e "\nSecret passphrase = '$secret_pass'"
else
  echo -e "\nInput timed out" >&2
  exit 1
fi
```

It's possible to supply the user with default response using `-e` and `-i` options:

```bash
#!/bin/bash

# read-default: supply a default value if user presses Enter key.

read -e -p "What is your user name? " -i $USER
echo "You answered: '$REPLY'"
```

##### IFS

The behavior of words splitting is configured by a shell variable named `IFS` (
for Internal Field Separator). The default value of `IFS` contains a
space, a tab, and a newline character, each of which will separate items
from one another.

We can adjust the value of `IFS` to control the separation of fields input to `read`.

```bash
#!/bin/bash

# read-ifs: read fields from a file

FILE=/etc/passwd

read -p "Enter a username > " user_name

file_info=$(grep "^$user_name:" $FILE)

if [ -n "$file_info" ]; then
  IFS=":" read user pw uid gid name home shell <<< "$file_info"
  echo "User =          '$user'"
  echo "UID =           '$uid'"
  echo "GID =           '$gid'"
  echo "Full Name =     '$name'"
  echo "Home Dir =      '$home'"
  echo "Shell =         '$shell'"
else
  echo "No such user '$user_name'" >&2
  exit 1
fi
```

**The IFS line consists of three parts: a variable assignment,
a `read` command with a list of variable names as arguments, and a strange
new redirection operator. The shell allows one or more variable
assignments to take place immediately before a command**. The effect of the
assignment is **temporary**; only changing the environment for the duration
of the command.

Alternative, we could have coded it this way:

```bash
#!/bin/bash

# read-ifs: read fields from a file

FILE=/etc/passwd

read -p "Enter a username > " user_name

file_info=$(grep "^$user_name:" $FILE)

if [ -n "$file_info" ]; then
  OLD_IFS="$IFS"
  IFS=":"
  read user pw uid gid name home shell <<< "$file_info"
  IFS="$OLD_IFS"
  echo "User =          '$user'"
  echo "UID =           '$uid'"
  echo "GID =           '$gid'"
  echo "Full Name =     '$name'"
  echo "Home Dir =      '$home'"
  echo "Shell =         '$shell'"
else
  echo "No such user '$user_name'" >&2
  exit 1
fi
```

The `<<<` operator indicates a _here string_. It is like a here document,
only shorter, consisting of a single string.

However, the below method does not work:

```bash
echo "$file_info" | IFS=":" read user pw uid gid name home shell
```

Because...

> **You Can't Pipe read**

In `bash`, pipelines create subshells. These are copies of the shell
and its environment which are used to execute the command in the pipeline.
In our example `echo "foo" | read`, `read` is executed in a subshell.

Subshells in the Unix-like systems create copies of the environment for the
processes to use while they execute. When the processes finishes the copy
of the environment is destroyed. This means that a subshell **can never
alter the environment of its parent process**. In the example,
`read` assigns the value "foo" to the variable `REPLY` in its subshell's environment,
but when the command exits, the subshell and its environment are destroyed,
and the effect of the assignment is **lost**.
Using _here string_ is one way to work around this behavior. The other
method is discussed in [Chapter 36 -- Process substitution](#ch36-process-substitution)

#### Validating Input

Here we have an example program that validates various kinds of input:

```bash
#!/bin/bash

# read-validate: validate input

invalid_input() {
  echo "Invalid input '$REPLY'" >&2
  exit 1
}

read -p "Enter a single item > "

[[ -z $REPLY ]] && invalid_input

(( $(echo $REPLY | wc -w) > 1 )) && invalid_input

if [[ $REPLY =~ ^[-[:alnum:]\._]+$ ]]; then
  echo "'$REPLY' is a valid filename."
  if [[ -e $REPLY ]]; then
    echo "And file '$REPLY' exists."
  else
    echo "However, file '$REPLY' does not exists."
  fi

  # is input a floating point number?
  if [[ $REPLY =~ ^-?[[:digit:]]*\.[[:digit:]]+$ ]]; then
    echo "'$REPLY' is a floating point number."
  else
    echo "'$REPLY' is not a floating point number."
  fi

  # is input an integer?
  if [[ $REPLY =~ ^-?[[:digit:]]+$ ]]; then
    echo "'$REPLY' is an integer."
  else
    echo "'$REPLY' is not an integer."
  fi
else
  echo "The string '$REPLY' is not a valid filename."
fi
```

#### Menus

A common type of interactivity is called _menu-driven_. In menu-driven programs, the user is presented with a list of choices and is asked to choose one.

We can construct our `sys_info_page` program to below menu:

```bash
Please Select:

1. Display System Information
2. Display Disk Space
3. Display Home Space Utilization
4. Quit

Enter selection [0-3] >
```

```bash
#!/bin/bash

# read-menu: a menu driven system information program

clear
echo "
Please Select:

1. Display System Information
2. Display Disk Space
3. Display Home Space Utilization
4. Quit
"
read -p "Enter selection [0-3] > "

if [[ $REPLY =~ ^[0-3]$ ]]; then
  if [[ $REPLY == 0 ]]; then
    echo "Program terminated."
    exit
  fi
  if [[ $REPLY == 1 ]]; then
    echo "Hostname: $HOSTNAME"
    uptime
    exit
  fi
  if [[ $REPLY == 2 ]]; then
    df -h
    exit
  fi
  if [[ $REPLY == 3 ]]; then
    if [[ $(id -u) -eq 0 ]]; then
      echo "Home Space Utilization (All Users)"
      du -sh /home/*
    else
      echo "Home Space Utilization ($USER)"
      du -sh $HOME
    fi
    exit
  fi
else
  echo "Invalid entry." >&2
  exit 1
fi
```

**Important**: Notice the use of the `exit` command in this script. **The
presence of multiple exit points in a program is generally a bad idea (it
makes program logic harder to understand), but it works in this script.**

### 29 - Flow Control: Looping With `while` / `until`

#### Looping

##### `while`

Syntax:

```bash
while commands; do
  commands;
done
```

```bash
#!/bin/bash

# while-count: display a series of numbers

count=1

while [[ $count -le 5 ]]; do
  echo $count
  count=$((count + 1))
done
echo "Finished."
```

Like `if`, `while` evaluates the exit status of a list of command.
As long as the exit status is zero, it performs the commands inside the loop.

```bash
#!/bin/bash

# while-menu: a menu driven system information program

DELAY=3 # Number of seconds to display results

while [[ $REPLY != 0 ]]; do
  clear
  cat <<- _EOF_
    Please Select:

    1. Display System Information
    2. Display Disk Space
    3. Display Home Space Utilization
    0. Quit
_EOF_
  read -p "Enter selection [0-3] > "

  if [[ $REPLY =~ ^[0-3]$ ]]; then
    if [[ $REPLY == 1 ]]; then
      echo "Hostname: $HOSTNAME"
      uptime
      sleep $DELAY
    fi
    if [[ $REPLY == 2 ]]; then
      df -h
      sleep $DELAY
    fi
    if [[ $REPLY == 3 ]]; then
      if [[ $(id -u) -eq 0 ]]; then
        echo "Home Space Utilization (All Users)"
        du -sh /home/*
      else
        echo "Home Space Utilization ($USER)"
        du -sh $HOME
      fi
      sleep $DELAY
    fi
  else
    echo "Invalid entry."
    sleep $DELAY
  fi
done
echo "Program terminated."
```

#### Breaking Out Of A Loop

`break` and `continue`

```bash
#!/bin/bash

# while-menu2: a menu driven system information program

DELAY=3 # Number of seconds to display results

while true; do
  clear
  cat <<- _EOF_
    Please Select:

    1. Display System Information
    2. Display Disk Space
    3. Display Home Space Utilization
    0. Quit
_EOF_
  read -p "Enter selection [0-3] > "

  if [[ $REPLY =~ ^[0-3]$ ]]; then
    if [[ $REPLY == 1 ]]; then
      echo "Hostname: $HOSTNAME"
      uptime
      sleep $DELAY
      continue
    fi
    if [[ $REPLY == 2 ]]; then
      df -h
      sleep $DELAY
      continue
    fi
    if [[ $REPLY == 3 ]]; then
      if [[ $(id -u) -eq 0 ]]; then
        echo "Home Space Utilization (All Users)"
        du -sh /home/*
      else
        echo "Home Space Utilization ($USER)"
        du -sh $HOME
      fi
      sleep $DELAY
      continue
    fi
    if [[ $REPLY == 0 ]]; then
      break
    fi
  else
    echo "Invalid entry."
    sleep $DELAY
  fi
done
echo "Program terminated."
```

##### `until`

Does the opposite of `while`

```bash
#!/bin/bash

# until-count: display a series of numbers

count=1

until [[ $count -gt 5 ]]; do
  echo $count
  count=$((count + 1))
done
echo "Finished"
```

#### Reading Files With Loops

Place the redirection operator after `done` statement.

```bash
#!/bin/bash

# while-read: read lines from a file

while read distro version release; do
  printf "Distro: %s\tVersion: %s\tReleased: %s\n" \
    $distro \
    $version \
    $release
done < distros.txt
```

It will exit after each line is read, with a zero exit status until
the end-of-file is reached. At that point, it will exit with a non-zero
exit status, thereby terminating the loop.

It is also possible to pipe standard input into a loop:

```bash
#!/bin/bash

# while-read2: read lines from a file

sort -k 1,1 -k 2n distros.txt | while read distro version release; do

  printf "Distro: %s\tVersion: %s\tReleased: %s\n" \
    $distro \
    $version \
    $release
done
```

However, it is important to remember **that since a pipe will execute the loop
in a subshell, any variables created or assigned within the loop will be lost
when the loop terminates**.

### 30 - Troubleshooting

#### Syntactic Errors

Syntactic errors involves **mistyping some element of shell syntax**, lead
to the shell refusing to execute the script.

##### Missing Quotes

```bash
$ cat -n ./trouble && ./trouble
     1  #!/bin/bash
     2
     3  # trouble: script to demonstrate common errors
     4
     5  number=1
     6
     7  if [ $number = 1 ]; then
     8    # focus on below line
     9    echo "Number is equal to 1
    10  else
    11    echo "Number is not equal to 1.
    12  fi
./trouble: line 11: unexpected EOF while looking for matching `"'
./trouble: line 13: syntax error: unexpected end of file
```

Interestingly, the line numbers reported are not where the missing
quote was removed, but rather much later in the program.
We can see why: if we follow the program after the missing quote.
`bash` will continue looking for the closing quote until it finds one,
which is the second `echo` command. `bash` becomes very confused
after that, and the syntax of the `if` command is broken because
the `fi` statement is now inside a quoted (but open) string.

In long scripts, this kind of error can be quite hard to find.
Using an editor with syntax highlighting will help.

##### Missing Or Unexpected Tokens

```bash
$ cat -n ./trouble && ./trouble
     1  #!/bin/bash
     2
     3  # trouble: script to demonstrate common errors
     4
     5  number=1
     6
     7  if [ $number = 1 ] then
     8    echo "Number is equal to 1."
     9  else
    10    echo "Number is not equal to 1."
    11  fi
./trouble: line 9: syntax error near unexpected token `else'
./trouble: line 9: `else'
```

With the semicolon removed, the word `then` is added to the list of arguments,
which is syntactically legal. The following `echo` command is legal, too, it's
interpreted as another command in the list of commands that `if` will
evaluate for an exit code. The `else` is encountered next, but it's out of place,
since the shell recognize it as a _reserved_ word and not the name of a command,
hence the error message.

##### Unanticipated Expansions

```bash
$ cat -n ./trouble && ./trouble
     1  #!/bin/bash
     2
     3  # trouble: script to demonstrate common errors
     4
     5  number=
     6
     7  if [ $number = 1 ]; then
     8    echo "Number is equal to 1."
     9  else
    10    echo "Number is not equal to 1."
    11  fi
./trouble: line 7: [: =: unary operator expected
Number is not equal to 1.
```

The `test` undergoes expansion with `number` being empty:

```bash
[ = 1]
```

The `=` operator is a binary operator, but the first is missing.
So the `test` failed (because of the error), the `if`
command receives a non-zero exit code and acts accordingly, and the second
`echo` command is executed.

The problem can be corrected by adding quotes around the first argument in the `test`
command:

```bash
[ "$number" = 1]
```

Then when expansion occurs, the result will be this:

```bash
[ "" = 1 ]
```

In addition to empty strings, quotes should be used in cases
where a value could expand into multi-word strings, as with
filenames containing embedded spaces.

#### Logical Errors

Unlike syntactic errors, logical errors do not prevent a script from running.
The script will run, but it will not produce the desired result, due to a problem
with its logic. Here are a few of the most common kinds found in scripts:

1. **Incorrect conditional expression**

2. **"Off by one" errors**. It is possible to overlook that the loop may
   require that the counting start with zero.

3. **Unanticipated situations**. Such as a filename that contains embedded spaces
   that expands into multiple command arguments rather than a single filename.

##### Defensive Programming

It is important to verify assumptions when programming. This means
a careful evaluation of the `exit status` of programs and commands that
are used by a script.

```bash
$ cd $dir_name
$ rm *
```

If the `cd` fails, the script deletes the files in the current working directory.
If the variable dir_name is empty, then the script deletes all the files in the home directory.

It **might** be wise to make the execution of `rm` contingent on the success of `cd`

```bash
cd $dir_name && rm *
```

However, it did not solve the empty variable case. **Tip**: If it was quoted, it executes `cd ""`, which will still delete the files in current directory

This could also be avoided by checking to see if the `dir_name` actually exists

```bash
[[ -d $dir_name ]] && cd $dir_name && rm *
```

Often, it is best to terminate the script with an error when an situation such
as the one above occurs:

```bash
if [[ ! -d "$dir_name" ]]; then
  echo "No such directory: '$dir_name'" >&2
  exit 1
fi
if ! cd $dir_name; then
  echo "Cannot cd to '$dir_name'" >&2
  exit 1
fi
if ! rm *; then
  echo "File deletion failed. Check results" >&2
  exit 1
fi
```

##### Verifying Input

A general rule of good programming is that if a program accepts input,
it must be able to deal with anything it receives. This usually means
that input must be carefully screened, to ensure that only valid
input is accepted for further processing.

```bash
[[ $REPLY =~ ^[0-3]$]]
```

Sometimes these sorts of tests can be very challenging to write, but
the effort is necessary to produce a quality script.

> **Design Is A Function Of Time**

Sometimes, a "quick-and-dirty" script will do if it's only going to be used once
and only used by the programmer. Such scripts don't need a lot of comments
and defensive checks.

On the other hand, if a script is intended for production use, that is, a script
that will be used over and over for an important task or by multiple users, it needs
much more careful development.

#### Testing

Testing the original fragment of code would be dangerous, since its purpose is to
delete files, but could modify the code to make the test safe:

```bash
if [[ -d $dir_name ]]; then
  if cd $dir_name; then
    echo rm * # TESTING
  else
    echo "cannot cd to '$dir_name'" >&2
    exit 1
  fi
else
  echo "no such directory: '$dir_name'" >&2
  exit 1
fi
exit # TESTING

# ... other part of the script
```

The most important change is placing an `echo` command just before the `rm` command
to allow the command and its expanded argument list to be displayed.

We place an `exit` command to conclude the test and prevent any other part of the
script from being carried out.

We also include some comments that act as "markers" for our test-related changes.
These can be used to help find and remove the changes when testing is complete.

#### Debugging

##### Finding The Problem Area

Commenting out sections of a scripts to see if the removal
of the code has any impact on the behavior of the bug.

##### Tracing

One tracing method involves placing informative messages in a script
that display the location of execution. We can add messages to our code fragment:

```bash
echo "preparing to delete files" >&2
if [[ -d $dir_name ]]; then
  if cd $dir_name; then
echo "deleting files" >&2
    rm *
  else
    echo "cannot cd to '$dir_name'" >&2
    exit 1
  fi
else
  echo "no such directory: '$dir_name'" >&2
  exit 1
fi
echo "file deletion complete" >&2
```

We send the messages to standard error to separate them from normal output.
We also do not indent the lines containing the messages, so it is
easier to find when it's time to remove them.

`bash` also provides a method of tracing, implemented by the `-x` option
and the `set` command with the `-x` option.

```bash
#!/bin/bash -x

# trouble: script to demonstrate common errors

number=1

if [ $number = 1 ]; then
  echo "Number is equal to 1."
else
  echo "Number is not equal to 1."
fi
```

When executed, the results look like this:

```bash
$ ./testing
+ number=1
+ '[' 1 = 1 ']'
+ echo 'Number is equal to 1.'
Number is equal to 1.
```

The leading plus signs indicate the display of the trace to distinguish
them from lines of regular output. The plus sign is the default character for
trace output. It is contained in the `PS4` shell variable. The
contents of this variable can be adjusted to make the prompt more useful.

**Note that the single quotes are required to prevent expansion until the
prompt is actually used**

```bash
$ export='$LINENO + '
$ ./testing
3 + number=1
5 + '[' 1 = 1 ']'
6 + echo 'Number is equal to 1.'
Number is equal to 1.
```

To perform a trace on a selected portion of a script, rather than the entire script,
we can use the `set` command with the `-x` option:

```bash
#!/bin/bash

# trouble: script to demonstrate common errors

number=1

set -x # Turn on tracing
if [ $number = 1 ]; then
  echo "Number is equal to 1."
else
  echo "Number is not equal to 1."
fi
set +x # Turn off tracing
```

This technique can be used to examine multiple portions of a
troublesome script.

##### Examining Values During Execution

It is often useful, along with tracing, to display the content of variables to see the
internal workings of a script while it is being executed.

```bash
#!/bin/bash

# trouble: script to demonstrate common errors

number=1

echo "number=$number" # DEBUG
#...
```

This technique is particularly useful when watching the behavior of loops and arithmetic within scripts.

### 31 - Flow Control: Branching With `case`

#### `case`

```bash
case word in
  [pattern [| pattern]...) commands ;;]...
esac
```

```bash
#!/bin/bash

# case-menu: a menu driven system information program

clear
echo "
Please Select:

1. Display System Information
2. Display Disk Space
3. Display Home Space Utilization
0. Quit
"
read -p "Enter selection [0-3] > "

case $REPLY in
  0)  echo "Program terminated."
      exit
      ;;
  1)  echo "Hostname: $HOSTNAME"
      uptime
      ;;
  2)  df -h
      ;;
  3)  if [[ $(id -u) -eq 0 ]]; then
        echo "Home Space Utilization (All Users)"
        du -sh /home/*
      else
        echo "Home Space Utilization ($USER)"
        du -sh $HOME
      fi
      ;;
  *)  echo "Invalid entry." >&2
      exit 1
      ;;
esac
```

**Tip**: The cases does not fallback to the next case. There is no need
to `break` it. **After a match is found, no further matches are attempted**.

See ";;&" also to perform multiple actions

##### Patterns

There are some valid patterns:

| Pattern        | Description                                                                                                                                                                                                                    |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| a)             | Matches if word equals "a"                                                                                                                                                                                                     |
| \[\[:alpha:]]) | Matches if word is a single alphabetic character.                                                                                                                                                                              |
| ???)           | Matches if word is exactly three characters long.                                                                                                                                                                              |
| \*.txt)        | Matches if word ends with the characters ".txt"                                                                                                                                                                                |
| \*)            | Matches any value of word. \*\*It is a good practice to include this as the last pattern in a `case` command, to catch any values of word that did not match a previous pattern; that is to catch any possible invalid values. |

```bash
#!/bin/bash

read -p "enter word > "

case $REPLY in
  [[:alpha:]]) echo "is a single alphabetic character." ;;
  [ABC][0-9]) echo "is A, B, or C followed by a digit." ;;
  ???) echo "is three characters long.";;
  *.txt) echo "is a word ending in '.txt'" ;;
  *) echo "is something else.";;
esac
```

It is also possible to combine multiple patterns using the vertical bar
character as a separator. This creates an "or" conditional pattern.

```bash
#!/bin/bash

# case-menu: a menu driven system information program

clear
echo "
Please Select:

A. Display System Information
B. Display Disk Space
C. Display Home Space Utilization
Q. Quit
"
read -p "Enter selection [A, B, C or Q] > "

case $REPLY in
  q|Q)  echo "Program terminated."
      exit
      ;;
  a|A)  echo "Hostname: $HOSTNAME"
      uptime
      ;;
  b|B)  df -h
      ;;
  c|C)  if [[ $(id -u) -eq 0 ]]; then
        echo "Home Space Utilization (All Users)"
        du -sh /home/*
      else
        echo "Home Space Utilization ($USER)"
        du -sh $HOME
      fi
      ;;
  *)  echo "Invalid entry." >&2
      exit 1
      ;;
esac
```

##### Performing Multiple Actions

How if we want to match "a" which is both lower case and alphabetic, as well as
a hexadecimal digit?

In `bash` prior to version 4.0 there was **no way for `case` to match more than one test**.
Modern versions of `bash`, add the ";;&" notation to terminate each action:

```bash
#!/bin/bash

# case4-2: test a character

read -n 1 -p "Type a character > "
echo
case $REPLY in
  [[:upper:]]) echo "'$REPLY' is upper case.            " ;;&
  [[:lower:]]) echo "'$REPLY' is lower case.            " ;;&
  [[:alpha:]]) echo "'$REPLY' is alphabetic.            " ;;&
  [[:digit:]]) echo "'$REPLY' is a digit.               " ;;&
  [[:graph:]]) echo "'$REPLY' is a visible character.   " ;;&
  [[:punct:]]) echo "'$REPLY' is a puntuation symbol.   " ;;&
  [[:space:]]) echo "'$REPLY' is a whitespace character." ;;&
  [[:xdigit:]]) echo "'$REPLY' is a hexadecimal digit.   " ;;&
esac
```

```bash
$ case4-2
Type a character > a
'a' is lower case.
'a' is alphabetic.
'a' is a visible character.
'a' is a hexadecimal digit.
```

**The addition of the ";;&" syntax allows case to continue on to the next
test rather than simply terminating.**

### 32 - Positional Parameters

#### Accessing The Command Line

The shell provides a set of variables called _positional parameters_ that contain
the individual words on the command line. The variables are named `0` through `9`.

```bash
$ cat posit-param
#!/bin/bash

# posit-param: script to view command line parameters

echo "
\$0 = $0
\$1 = $1
\$2 = $2
\$3 = $3
\$4 = $4
\$5 = $5
\$6 = $6
\$7 = $7
\$8 = $8
\$9 = $9
"


$ posit-param 1 2 3 4 5 6 7 8 9 10 11
$0 = ./posit-param
$1 = 1
$2 = 2
$3 = 3
$4 = 4
$5 = 5
$6 = 6
$7 = 7
$8 = 8
$9 = 9
```

**Tip**: You can actually access more than nine parameters using parameter
expansion: e.g `${10}`, `${11}`

##### Determining The Number of Arguments

The shell also provides a variable , `$#`, that yields the **number of arguments**
on the command line:

```bash
$ cat posit-param
#!/bin/bash

# posit-param: script to view command line parameters

Number of arguments: $#
echo "
\$0 = $0
\$1 = $1
\$2 = $2
\$3 = $3
\$4 = $4
\$5 = $5
\$6 = $6
\$7 = $7
\$8 = $8
\$9 = $9
"
$ ./posit-param 1 2 3 4 5 6 7 8 9 10 11
Number of arguments: 11

$0 = ./posit-param
$1 = 1
$2 = 2
$3 = 3
$4 = 4
$5 = 5
$6 = 6
$7 = 7
$8 = 8
$9 = 9

```

##### `shift` -- Getting Access To Many Arguments

How can we process many arguments like below?

```bash
$ posit-param *
```

The shell provides a method, albeit a clumsy one, to do this. The
`shift` command causes all the parameters to "move down one" each time it
is executed. In fact, by using `shift`, it is possible to get by with
only `$1` parameter (in addition to `$0`, which never changes). After one `shift`,
`$#` reduced by one.

```bash
#!/bin/bash

# posit-param2: script to display all arguments

count=1

while [[ $# -gt 0 ]]; do
  echo "Argument $count = $1"
  count=$((count + 1))
  shift
done
```

##### `getopts`

`bash` includes a builtin command called getopts.

```bash
$ getopts optstring name [args]
```

`optstring` - If a character is followed by a colon,
it indicates that this specific option requires an argument.
For example, an optstring of `abc` would mean the script is looking
for options `-a`, `-b`, or `-c` without any arguments. If the optstring
is `a:bc`, it means the script is looking for options `-a`, `-b`, and `-c`,
but only `-a` requires an argument.

`name` - a variable that `getopts` uses to store the current option
that it is processing.

`[args]` - Optional. `getopts` will parse these arguments instead of the
positional parameters.

```bash
#!/bin/bash

OPTSTRING=":x:y:"

while getopts ${OPTSTRING} opt; do
  case ${opt} in
    x)
      echo "Option -x was triggered, Argument: ${OPTARG}"
      ;;
    y)
      echo "Option -y was triggered, Argument: ${OPTARG}"
      ;;
    :)
      echo "Option -${OPTARG} requires an argument."
      exit 1
      ;;
    ?)
      echo "Invalid option: -${OPTARG}."
      exit 1
      ;;
  esac
done

$ getopts-while -x -y
Option -x was triggered, Argument: -y
$ getopts-while -x hello -y world
Option -x was triggered, Argument: hello
Option -y was triggered, Argument: world
```

##### Simple Applications

```bash
#!/bin/bash

# file_info: simple file information program

PROGRAME=$(basename $0)

if [[ -e "$1" ]]; then
  echo -e "\nFile Type:"
  file $1
  echo -e "\nFile Status:"
  stat $1
else
  echo "$PROGNAME: usage: $PROGNAME file" >&2
  exit 1
fi

$ file-info
file-info: usage: file-info file
```

##### Using Positional Parameters With Shell Functions

```bash
#!/bin/bash

# file_info: simple file information program

PROGNAME=$(basename $0)

file_info() {
	# file_info: function to display file information
	if [[ -e "$1" ]]; then
		echo -e "\nFile Type:"
		file $1
		echo -e "\nFile Status:"
		stat $1
	else
		echo "$FUNCNAME: usage: $FUNCNAME file" >&2
		exit 1
	fi
}
```

Notice that the PROGNAME variable was changed to the shell variable FUNCNAME. The shell automatically updates this variable to keep track of the currently executed shell function.

#### Handling Positional Parameters _En Masse_

The shell provides two special parameters, which expand into
the complete list of positional parameters.

| Parameter | Description                                                                                                                                                                                                                                                                              |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `$*`      | Expands into the list of positional parameters, starting with 1. **When surrounded by double quotes, it expands into a double quoted string containing all the positional parameters**, each separated by the first character of the IFS shell variable (by default a space character ). |
| `$@`      | Expands into the list of positional parameters, starting with 1. **When surrounded by double quotes, it expands each positional parameter into a separate word surrounded by double quotes**                                                                                             |

```bash
$ cat posit-params3
#!/bin/bash

# posit-params3: script to demonstrate $* and $@

print_params() {
  echo "\$1 = $1"
  echo "\$2 = $2"
  echo "\$3 = $3"
  echo "\$4 = $4"
}

pass_params() {
  echo -e "\n" '$* :'; print_params $*
  echo -e "\n" '"$*" :'; print_params "$*"
  echo -e "\n" '$@ :'; print_params $@
  echo -e "\n" '"$@" :'; print_params "$@"
}

pass_params "word" "word with spaces"
$ posit-params3
 $* :
$1 = word
$2 = word
$3 = with
$4 = spaces

 "$*" :
$1 = word word with spaces
$2 =
$3 =
$4 =

 $@ :
$1 = word
$2 = word
$3 = with
$4 = spaces

 "$@" :
$1 = word
$2 = word with spaces
$3 =
$4 =

```

Hence, "$@" is by far the most useful for most situations, because it preserves
the integrity of each positional parameter.

#### A More Complete Application

We are going to resume work on our `sys_info_page` program.
We are going to include:

- **Output file**. `-f file` or `--file file` to specify a name for a file to contain
  the program's output

- **Interactive mode**. `-i` or `--interactive` to prompt user for an output
  filename and will determine if the specified file already exists. If it does,
  the user will be prompted before existing file is overwritten.

- **Help**. `-h` or `--help` may be specified to cause the program to output
  an informative usage message.

```bash
#!/bin/bash

usage() {
  echo "$PROGNAME: usage: $PROGNAME [-f file | -i]"
  return
}

report_uptime() {
  cat <<- _EOF_
      <H2>System Uptime</H2>
      <PRE>$(uptime)</PRE>
_EOF_
	return
}

report_disk_space() {
  cat <<- _EOF_
      <H2>Disk Space Utilization</H2>
      <PRE>$(df -h)</PRE>
_EOF_
	return
}

report_home_space() {
  if [[ $(id -u) -eq 0 ]]; then
    cat <<- _EOF_
      <H2>Home Space Utilization (All Users)</H2>
      <PRE>$(du -sh /Users/*)</PRE>
_EOF_
  else
    cat <<- _EOF_
      <H2>Home Space Utilization ($USER)</H2>
      <PRE>$(du -sh $HOME)</PRE>
_EOF_
  fi
  return
}

# process command line options

interactive=
filename=

while [[ -n "$1" ]]; do
  case $1 in
    -f | --file)        shift
                        filename=$1
                        ;;
    -i | --interactive) interactive=1
                        ;;
    -h | --help)        usage
                        exit
                        ;;
    *)                  usage >&2
                        exit 1
                        ;;
  esac
  shift
done
```

Interactive mode:

```bash
if [[ -n $interactive ]]; then
  while true; do
    read -p "Enter name of output file: " filename
    if [[ -e $filename ]]; then
      read -p "'$filename' exists. Overwrite? [y/n/q] > "
      case $REPLY in
        Y|y)  break
              ;;
        Q|q)  echo "Program terminated"
              exit
              ;;
        *)    continue
              ;;
      esac
    elif [[ -z $filename ]]; then
      continue
    else
      break
    fi
  done
fi
```

Output:

```bash
write_html_page() {
  cat <<- _EOF_
  <HTML>
    <HEAD>
      <TITLE>$TITLE</TITLE>
    </HEAD>
    <BODY>
      <H1>$TITLE</H1>
      <P>$TIMESTAMP</P>
      $(report_uptime)
      $(report_disk_space)
      $(report_home_space)
    </BODY>
  </HTML>
_EOF_
  return
}

if [[ -n $filename ]]; then
  if touch $filename && [[ -f $filename ]]; then
    write_html_page > $filename
  else
    echo "PROGNAME: Cannot write file '$filename'" >&2
    exit 1
  fi
else
  write_html_page
fi
```

### 33 - Flow Control: Looping With `for`

In modern versions of `bash`, `for` is available in two forms.

#### `for`: Traditional Shell Form

```bash
for variable [in words]; do
  commands
done
```

The `for` command is useful on the command line:

```bash
$ for i in A B C D; do echo $i; done
```

The really powerful feature of `for` is the number of interesting ways we can create
the list of words. For example, through brace expansion:

```bash
$ for i in {001..100}; do echo $i; done
```

or command substitution:

```bash
#!/bin/bash

# longest-word: find longest string in a file

while [[ -n $1 ]]; do
  if [[ -r $1 ]]; then
    max_word=
    max_len=0
    for i in $(strings $1); do
      len=$(echo -n $i | wc -c)
      if (( len > max_len )); then
        max_len=$len
        max_word=$i
      fi
    done
    echo "$1: '$max_word' ($max_len characters)"
  fi
  shift
done
```

**Tip**: strings - find the **printable strings** in a object, or other binary, file

`for` defaults to processing the positional parameters:

```bash
#!/bin/bash

# longest-word2: find longest string in a file

for i; do
  if [[ -r $i ]]; then
    max_word=
    max_len=0
    for j in $(strings $i); do
      len=$(echo -n $j | wc -c)
      if (( len > max_len )); then
        max_len=$len
        max_word=$j
      fi
    done
    echo "$i: '$max_word' ($max_len characters)"
  fi
done
```

#### `for:` C Language Form

Recent versions of `bash` have added a second form of `for` command
syntax, one that resembles the form found in the C programming language

```bash
for (( expression1; expression2; expression3 )); do
  commands
done
```

Here is a typical application:

```bash
#!/bin/bash

# simple_counter: demo of C style for command

for (( i=0; i<5; i=i+1 )); do
  echo $i
done
```

#### Summary

```bash
report_home_space() {
  if [[ $(id -u) -eq 0 ]]; then
    cat <<- _EOF_
      <H2>Home Space Utilization (All Users)</H2>
      <PRE>$(du -sh /Users/*)</PRE>
    _EOF_
  else
    cat <<- _EOF_
      <H2>Home Space Utilization ($USER)</H2>
      <PRE>$(du -sh $HOME)</PRE>
    _EOF_
  fi
  return
}
```

We will rewrite `report_home_space()` function to provide more detail
for each user's home directory, and include the total number of files
and subdirectories in each:

```bash
report_home_space() {
  local format="%8s%10s%10s\n"
  local i dir_list total_files total_dirs total_size user_name

  if [[ $(id -u) -eq 0 ]]; then
    dir_list=/home/*
    user_name="All Users"
  else
    dir_list=$HOME
    user_name=$USER
  fi

  echo "<H2>Home Space Utilization ($user_name)</H2>"

  for i in $dir_list; do
    total_files=$(find $i -type f | wc -l)
    total_dirs=$(find $i -type d | wc -l)
    total_size=$(du -sh $i | cut -f 1)

    echo "<H3>$i</H3>"
    echo "<PRE>"
    printf "$format" "Dirs" "Files" "Size"
    printf "$format" "----" "-----" "----"
    printf "$format" $total_dirs $total_files $total_size
    echo "</PRE>"
  done
  return
}
```

<a id="ch34"></a>

### 34 - Strings And Numbers

#### Parameter Expansion

Though parameter expansion came up in Chapter 7, we did not cover it in detail
because most parameter expansions are used in scripts rather than on the command
line.

##### Basic Parameters

```bash
$h
${h}hello
$1
${11}
```

##### Expansions To Manage Empty Variables

These expansions are handy for handling missing positional parameters and assigning
default values to parameters.

> **`${parameter:-word}`**

If parameter is unset (i.e., does not exist), this expansion results in the value
of word.

```bash
$ foo=
$ echo ${foo:-"substitute value if unset"}
substitute value if unset
$ echo $foo

$ foo=bar
$ echo ${foo:-"substitute value if unset"}
bar
$ echo $foo
bar
```

> **`${parameter:=word}`**

Same as above expansion. In addition, **the value of word is assigned to parameter**.

```bash
$ foo=
$ echo ${foo:="default value if unset"}
default value if unset
$ echo $foo
default value if unset
$ foo=bar
$ echo ${foo:="defalt value if unset"}
bar
$ echo $foo
bar
```

> **${parameter:?word}**

If parameter is unset or empty, this expansion causes the script
to exit with an error, **the contents of word are sent to standard error**.

```bash
$ foo=
$ echo ${foo:?"parameter is empty"}
parameter is empty
$ echo $foo

$ foo=bar
$ echo ${foo:?"parameter is empty"}
bar
$ echo $foo
bar
```

> **${parameter:+word}**

If parameter is unset or empty, the expansion results in nothing. If parameter
is not empty, the value of word is substituted for parameter; however, the
**value of parameter is not changed**.

```bash
$ foo=
$ echo ${foo:+"substitute value if set"}

$ echo $foo

$ foo=bar
$ echo ${foo:+"substitute value if set"}
substitute value if set
$ echo $foo
bar
```

<a name="expansion-return-var-name"></a>

##### Expansions That Return Variable Names

The shell has the ability to return the names of variables.
This is used in some rather exotic situations.

```bash
${!prefix*}
${!prefix@}
```

```bash
$ echo ${!BASH*}
BASH BASHOPTS BASHPID BASH_ALIASES BASH_ARGC BASH_ARGV BASH_ARGV0 BASH_CMDS BASH
_COMMAND BASH_COMPLETION_VERSINFO BASH_LINENO BASH_REMATCH BASH_SOURCE BASH_SUBS
HELL BASH_VERSINFO BASH_VERSION
```

##### String Operations

There is a large set of expansions that can be used to operate on strings.

> **${#parameter}**

Expands into **length of the string contained by parameter**. Normally, parameter
is a string; however, if parameter is either '@' or '\*', then
the expansion results in the number of position parameters.

```bash
$ foo="This string is long."
$ echo "'foo' is ${#foo} characters long."
'foo' is 20 characters long.
```

```bash
$ cat ascript
#!/bin/bash
echo ${#*} # or ${#@}
$ ascript a b c
3
```

> **${parameter:offset}**
>
> **${parameter:offset:length}**

Extract a portion of the string contained in parameter. Begins
at offset to the end of string unless the length is specified.
If the value of offset is negative, it is taken to means it starts from
the end of the string rather than the beginning. \*\*Note that negative values
must be preceded by a space to prevent confusion with ${parameter:-word} expression.

```bash
$ foo="This string is long"
$ echo ${foo:5}
string is long.
$ echo ${foo:5:6}
string
$ echo ${foo: -5:4}
long
```

If parameter is '@', the result of the expansion is length positional parameters,
starting at offset.

```bash
$ cat ascript
#!/bin/bash
echo ${@:2}
$ ascript a b c d
b c d
```

> **${parameter#pattern}**
>
> **${parameter##pattern}**

Removes a leading portion of the string contained in parameter defined
by pattern. pattern is a wildcard pattern like those used in pathname expansion.
The difference in the two forms is that the '#' form removes the shortest match,
while the '##' form removes the longest match.

```bash
$ foo=file.txt.zip
$ echo ${foo#*.}
txt.zip
$ echo ${foo##*.}
zip
```

> **${parameter%pattern}**
>
> **${parameter%%pattern}**

They are same as '#' and '##' expansions above, except they remove
text from the end of the string rather than from the beginning.

```bash
$ foo=file.txt.zip
$ echo ${foo%.*}
file.txt
$ echo ${foo%%.*}
file
```

> **${parameter/pattern/string}**
>
> **${parameter//pattern/string}**
>
> **${parameter/#pattern/string}**
>
> **${parameter/%pattern/string}**

This expansion performs a search-and-replace upon the contents of parameter.
In the normal form, only **the first occurrence of pattern** is replaced.
In the '//' form, all occurrences are replaced. The '/#' form requires that the
match occur at the beginning of the string, and the '.%' form requires the match
to occur at the end of the string. `/string` may be omitted to cause the matched text to be deleted.

```bash
$ foo=JPG.JPG
$ echo ${foo/JPG/jpg}
jpg.JPG
$ echo ${foo//JPG/jpg}
jpg.jpg
$ echo ${foo/#JPG/jpg}
jpg.JPG
$ echo ${foo/%JPG/jpg}
JPG.jpg
```

The string manipulation expansions can be used as substitutes for other common commands
such as `sed` and `cut`.

Example here modifies `longest-word` program, which use ${#j} to replace $(echo $j | wc -c)

```bash
#!/bin/bash

# longest-word3: find longest string in a file

for i; do
  if [[ -r $i ]]; then
    max_word=
    max_len=0
    for j in $(strings $i); do
      len=${#j}
      if (( len > max_len )); then
        max_len=$len
        max_word=$j
      fi
    done
    echo "$i: '$max_word' ($max_len characters)"
  fi
done
```

Let's compare it:

```bash
$ time longest-word2 longest-word
longest-word: 'longest-word:' (      13 characters)
longest-word2 longest-word  0.08s user 0.17s system 99% cpu 0.249 total
$ time longest-word3 longest-word
longest-word: 'longest-word:' (13 characters)
longest-word3 longest-word  0.01s user 0.01s system 83% cpu 0.027 total
```

Wow, a very significant improvement.

##### Case Conversion

Recent versions of `bash` have support for upper/lowercase conversion
of strings. `bash` has four parameter expansions and two options to
the `declare` command to support it.

```bash
$ cat ul-declare
#!/bin/bash

# ul-declare: demonstrate case conversion via declare

declare -u upper
declare -l lower

if [ "$1" ]; then
  upper="$1"
  lower="$1"
  echo $upper
  echo $lower
fi
$ ul-declare heLLo
HELLO
hello
```

There are four parameter expansions that perform upper/lowercase conversion:

| Format         | Result                                                                        |
| -------------- | ----------------------------------------------------------------------------- |
| ${paremeter,,} | Expand the value of parameter into all lowercase.                             |
| ${parameter,}  | Expand the value of parameter changing only the first character to lower case |
| ${parameter^^} | Expand the value of parameter into all uppercase letters                      |
| ${paramter^}   | Expand the value of parameter changing only the first character to uppercase  |

Example:

```bash
#!/bin/bash

# ul-param - demonstrate case conversion via parameter expansion

if [[ $1 ]]; then
  echo ${1,,}
  echo ${1,}
  echo ${1^^}
  echo ${1^}
fi

$ ul-param aBc
abc
aBc
ABC
ABc
```

#### Arithmetic Evaluation And Expansion

We looked at arithmetic expansion in Chapter 7.

**IMPORTANT**: The shell's arithmetic ONLY OPERATES ON INTEGERS

```bash
$((expression))
```

##### Number Bases

| Notation    | Description                |
| ----------- | -------------------------- |
| number      | Decimal (base 10) integers |
| 0number     | Octal                      |
| 0xnumber    | hexadecimal notation       |
| base#number | number is in base          |

Some examples:

```bash
$ echo $((0xff))
255
$ echo $((2#11111111))
255
$ echo $((017))
15
```

##### Unary Operators

`+`, `-`

##### Simple Arithmetic

| Operator | Description        |
| -------- | ------------------ |
| +        | Addition           |
| -        | Subtraction        |
| \*       | Multiplication     |
| /        | Integer division   |
| \*\*     | Exponentiation     |
| %        | Modulo (remainder) |

**Tip**: The results of division are always whole numbers: `$(( 5 / 2 ))` is 2

##### Assignment

We can do the assignment within arithmetic expression:

```bash
$ if (( foo = 5 )); then echo "It is true."; fi
$ echo $foo
5
```

**Important**: in `(())` expression, `==` evaluates equivalence. This
can be very confusion because `test` command accepts a single `=` for
string equivalent. This is yet another reason to use the more modern
`[[]]` and `(())` compound commands in place of `test`

In addition to the `=`, the shell also provides notations that
perform some very useful assignments:

| Notation             | Description                                   |
| -------------------- | --------------------------------------------- |
| `parameter = value`  | Simple assignment. Assigns value to parameter |
| `parameter += value` | Addition                                      |
| `parameter -= value` | Subtraction                                   |
| `parameter *= value` | Multiplication                                |
| `parameter /= value` | Division                                      |
| `parameter %= value` | Modulo                                        |
| `parameter++`        | Variable post-increment                       |
| `parameter--`        | Variable post-decrement                       |
| `++parameter`        | Variable pre-increment                        |
| `--parameter`        | Variable pre-decrement                        |

```bash
$ foo=1
$ echo $((foo++))
1
$ echo $foo
2
```

##### Bit Operations

| Operator | Description                                                       |
| -------- | ----------------------------------------------------------------- |
| ~        | Bitwise negation. Negate all the bits in a number.                |
| <<       | Left bitwise shift. Shift all the bits in a number to the left.   |
| >>       | Right bitwise shift. Shift all the bits in a number to the right. |
| &        | Bitwise AND                                                       |
| \|       | Bitwise OR                                                        |
| ^        | Bitwise XOR                                                       |

```bash
$ for ((i=0; i<8; ++i)); do echo $((1<<i)); done
```

##### Logic

| Operator          | Description              |
| ----------------- | ------------------------ |
| <=                | Less than or equal to    |
| >=                | Greater than or equal to |
| <                 | Less than                |
| >                 | Greater than             |
| ==                | Equal to                 |
| !=                | Not equal to             |
| &&                | Logical AND              |
| \|\|              | Logical OR               |
| expr1?expr2:expr3 | Ternary operator         |

When used for logical operations, expressions follow the rules of arithmetic
logic; that is, expressions that evaluate as zero are considered false,
while non-zero expressions are considered true.

```bash
$ if ((1)); then echo "true"; else echo "false"; fi
true
$ if ((0)); then echo "true"; else echo "false"; fi
false
```

**Important**: Three expressions must be arithmetic expression (strings won't work) in ternary operator.

**Important**: Performing assignment within the expression is not straightforward. `bash` will declare an error if so. This problem can be mitigated by surrounding the assignment expression with parentheses

```bash
$ a=0
$ ((a<1?++a:--a))
$ echo $a
1
$ ((a<1?++a:--a))
$ echo $a
0

$ a=0
$ ((a<1?a+=1:a-=1))
-bash: ((: a<1?a+=1:a-=1: attempted assignment to non-variable (error token is "
-=1")
$ ((a<1?(a+=1):(a-=1)))
```

```bash
#!/bin/bash

# arith-loop: script to demonstrate arithmetic operators

finished=0
a=0
printf "a\ta**2\ta**3\n"
printf "==\t====\t====\n"

until ((finished)); do
	b=$((a ** 2))
	c=$((a ** 3))
	printf "%d\t%d\t%d\n" $a $b $c
	((a < 10 ? ++a : (finished = 1)))
done
$ arith-loop
a       a**2    a**3
==      ====    ====
0       0       0
1       1       1
2       4       8
3       9       27
4       16      64
5       25      125
6       36      216
7       49      343
8       64      512
9       81      729
10      100     1000
```

#### `bc` - An Arbitrary Precision Calculator Language

How to perform higher math or even just use floating point numbers?
The answer is .. **we can't**.

To do this, we need to use an external program. There are several approaches
we can take. Embedding Perl or AWK programs is one possible solution.

Another approach is to use a specialized calculator program. One such
program found on most Linux systems is called `bc`.

The `bc` program reads a file written in its own C-like language and executes
it. We won't cover `bc` entirely here, just enough to get a taste.

```bash
/* A very simple bc script */

2 + 2
```

##### Using `bc`

```bash
$ bc foo.bc
bc 1.07.1
Copyright 1991-1994, 1997, 1998, 2000, 2004, 2006, 2008, 2012-2017 Free Software
 Foundation, Inc.
This is free software with ABSOLUTELY NO WARRANTY.
For details type `warranty'.
4
```

`-q` to suppress the copyright message.

It is also possible to pass a script to `bc` via standard input

```bash
$ bc -q < foo.bc
4
```

The ability to take standard input means that we can use here documents,
here strings, and pipes to pass scripts. This is a here string example:

```bash
$ bc <<< "2+2"
4
```

##### An Example Script

```bash
#!/bin/bash

# loan-calc: script to calculate monthly loan payments

PROGNAME=$(basename $0)

usage () {
  cat <<- EOF
  Usage: $PROGNAME PRINCIPAL INTEREST MONTHS

  Where:

  PRINCIPAL is the amount of the loan.
  INTEREST is the APR as a number (7% = 0.07).
  MONTHS is the length of the loan's term.

EOF
}

if (($# != 3)); then
  usage
  exit 1
fi

principal=$1
interest=$2
months=$3

bc <<- EOF
scale = 10
i = $interest / 12
p = $principal
n = $months
a = p * ((i * ((1 + i) ^ n)) / (((1 + i) ^ n) - 1))
print a, "\n"
EOF
```

A full description of the `bc` scripting language is provided
by the `bc` man page. While its mathematical notation
is slightly different from that of the shell (`bc` more closely
resembles C), most of it will be quite familiar, based on
what we have learned so far.

### 35 - Arrays

Arrays in `bash` are limited to a single dimension.
Array support first appeared in `bash` version 2. **The original Unix shell
program, `sh` did not support arrays at all**

#### Creating An Array

```bash
$ a[1]=foo
$ echo ${a[1]}
foo
```

An array can be created with the `declare` command:

```bash
$ declare -a a
```

Using `-a` option, this example of `declare` creates the array `a`.

#### Assigning Values To An Array

```bash
name[subscript]=value
```

**Tip**: subscript can be arithmetic expression

Multiple values may be assigned using the following syntax:

```bash
name=(value1 value2 ...)
$ days=(Sun Mon Tue Wed Thu Fri Sat)
```

It is also possible to assign values to a specific element by specifying
a subscript for each value:

```bash
$ days=([0]=Sun [1]=Mon [2]=Tue [3]=Wed [4]=Thu [5]=Fri [6]=Sat)
```

#### Accessing Array Elements

```bash
#!/bin/bash

# hours: script to count files by modification time

usage () {
  echo "usage: $(basename $0) directory >&2"
}

# Check that argument is a directory
if [[ ! -d $1 ]]; then
  usage
  exit 1
fi

# Initialize array
for i in {0..23}; do hours[i]=0; done

# stat below return: "2024-01-22 07:40:43.069178053 +0000"
for i in $(stat -c %y "$1"/* | cut -c 12-13); do
  j=${i/#0}
  ((++hour[j]))
  ((++count))
done

# Display data
echo -e "----\t-----\t----\t-----"
for i in {0..11}; do
  j=$((i + 12))
  printf "%02d\t%d\t%02d\t%d\%02d\t%d\n" $i ${hours[i]} $j ${hours[j]}
done
printf "\nTotal files = %d\n" $count

$ hours .
Hour    Files   Hour    Files
----    -----   ----    -----
00      1       12      0
01      0       13      0
02      1       14      0
03      0       15      0
04      2       16      0
05      0       17      0
06      2       18      0
07      3       19      0
08      3       20      0
09      1       21      0
10      0       22      0
11      0       23      0

Total files = 13

```

**Tip**: `j=${i/#0}` remove all the leading zeros from the hour field

#### Array Operations

Many common array operations. Such things as deleting arrays, determining
their size, sorting, eetc.

##### Outputting The Entire Contents Of An Array

The subscript `*` and `@` can be used to access every element in an array.
As with positional parameters, the `@` notation is the more useful of two.

```bash
$ animals=("a dog" "a cat" "a fish")
$ for i in ${animals[*]}; do echo $i; done
a
dog
a
cat
a
fish
$ for i in ${animals[@]}; do echo $i; done
a
dog
a
cat
a
fish
$ for i in "${animals[*]}"; do echo $i; done
a dog a cat a fish
$ for i in "${animals[@]}"; do echo $i; done
a dog
a cat
a fish
```

**The behavior of notations is identical until they are quoted**.

##### Determining The Number Of Array Elements

```bash
$ a[100]=foo
$ echo ${#a[@]} # number of array elements
1
$ echo ${#a[100]} # length of element 100
3
$ for i in "${a[@]}"; do echo $i; done
foo
```

**Important**: Note that a[0] to a[99] would not be initialized with empty values and counted.

##### Finding The Subscripts Used By An Array

[Expansion -- Return variable name](#expansion-return-var-name)

```bash
${!array[*]}
${!array[@]}
```

```bash
$ foo=([2]=a [4]=b [6]=c)
$ for i in "${foo[@]}"; do echo $i; done
a
b
c
$ for i in "${!foo[@]}"; do echo $i; done
2
4
6
```

##### Adding Elements To The End Of An Array

By using the `+=` assignment operator, we can automatically append values
to the end of an array.

```bash
$ foo=(a b c)
$ echo ${foo[@]}
a b c
$ foo+=(d e f)
$ echo ${foo[@]}
a b c d e f
```

##### Sorting An Array

The shell has no direct way of doing this, but it's not hard to do with a little
coding:

```bash
#!/bin/bash

# array-sort: Sort an array

a=(f e d c b a)

echo "Original array: ${a[@]}"
a_sorted=($(for i in "${a[@]}"; do echo $i; done | sort))
echo "Sorted array:   ${a_sorted[@]}"

$ array-sort
Original array: f e d c b a
Sorted array:   a b c d e f
```

**Tip**: What is on a_sorted line? Look at [group command](#group-commands)

##### Deleting An Array

Use the `unset` command

```bash
$ foo=(a b c d e f)
$ echo ${foo[@]}
a b c d e f
$ unset foo
$ echo ${foo[@]}

$
```

Unset can also delete single array element

```bash
$ foo=(a b c d e f)
$ echo ${foo[@]}
a b c d e f
$ unset 'foo[2]'
$ echo ${foo[@]}
a b d e f
```

**Important**: The array element MUST be quoted to prevent the shell from
performing pathname expansion.

```bash
$ foo=(a b c d e f)
$ echo ${#foo[@]}
6
$ foo=
$ echo ${#foo[@]}
6
$ echo ${foo[@]}
b c d e f
```

**Reference to an array variable without a subscript refers to element zero
of the array**:

```bash
$ foo=(a b c d e f)
$ foo=A
$ echo ${foo[@]}
A b c d e f
```

#### Associative Arrays

Recent version of `bash` now support _associative arrays_.
Which **use strings rather than integers as array indexes**.
This capability allow interesting new approaches to managin data.

```bash
#!/bin/bash

declare -A colors
colors["red"]="#ff0000"
colors["green"]="#00ff00"
colors["blue"]="#0000ff"
echo ${colors[@]}
echo ${!colors[@]}
echo ${color[0]}

$ associate-array
#0000ff #ff0000 #00ff00
blue red green

$
```

### 36 Exotica

There are a few that, while not in common use, are helpful for certain programming problems. We will cover them here.

<a name="group-commands"></a>

#### Group Commands And Subshells

`bash` allows commands to be grouped together. This can be done in one of two
ways:

_Group command_:

```bash
{ command1; command2; [command3; ...] }
```

_Subshell_:

```bash
( command1; command2; [command3; ...] )
```

**Important**: The braces must be separated from the commands by a space and the
last command must be terminated with either semicolon or a newline prior to the
closing brace.

They are both used to manage redirection.

```bash
$ ls -l > output.txt
$ echo "Listing of foo.txt" >> output.txt
$ cat foo.txt >> output.txt
```

Using a group command, we could code this as follows:

```bash
{ ls -l; echo "Listing of foo.txt"; cat foo.txt; } > output.txt
```

Using a subshell is similar:

```bash
(ls -l; echo "Listing of foo.txt"; cat foo.txt;) > output.txt
```

But where a group command or subshell really shines is with pipelines.
It is often useful to combine the results of several commands into a single stream

```bash
{ ls -l; echo "Listing of foo.txt"; cat foo.txt; } | lpr
```

```bash
#!/bin/bash

# array-2: Use arrays to tally file owners

declare -A files file_group file_owner groups owners

if [[ ! -d "$1" ]]; then
  echo "Usage: array-2 dir" >&2
  exit 1
fi

for i in "$1"/*; do
  owner=$(stat -c %U "$i")
  group=$(stat -c %G "$i")
  files["$i"]="$i"
  file_owner["$i"]=$owner
  file_group["$i"]=$group
  ((++owners[$owner]))
  ((++groups[$group]))
done

# List the collected files
{ for i in "${files[@]}"; do
    printf "%-40s %-10s %-10s\n" \
      "$i" ${file_owner["$i"]} ${file_group["$i"]}
  done } | sort

echo

# List owners
echo "File owners:"
{ for i in "${!owners[@]}"; do
    printf "%-10s %5d file(s)\n" "$i" ${owners["$i"]}
  done } | sort

echo

# List groups
echo "File group owners:"
{ for i in "${!groups[@]}"; do
    printf "%-10s %5d file(s)\n" "$i" ${groups["$i"]}
  done } | sort
```

```bash
$ array_2
./array_2                                ubuntu     ubuntu
./associate-array                        ubuntu     ubuntu
./case4-2                                ubuntu     ubuntu
./debian-cd_info.tar.gz                  ubuntu     ubuntu
./find_playground                        ubuntu     ubuntu
./foo.bc                                 ubuntu     ubuntu
./hello_world                            ubuntu     ubuntu
./hours                                  ubuntu     ubuntu
./read-default                           ubuntu     ubuntu
./retrieve_via_ftp                       ubuntu     ubuntu
./sys_info_page                          ubuntu     ubuntu
./test.html                              ubuntu     ubuntu
./ul-declare                             ubuntu     ubuntu
./ul-param                               ubuntu     ubuntu

File owners:
ubuntu        14 file(s)

File group owners:
ubuntu        14 file(s)
```

##### Process Substitution

While they look similar and can both be used to combine streams
for redirection, there is an important difference between group commands
and subshells. Whereas a group command executes all of its commands in
the current shell, a subshell executes its commands in a child copy of the
current shell. This means that the environment is copied and given
to a new instance of the shell. **When the subshell exits, the
copy of the environment is lost, so any changes made to the subshell's
environment is lost as well**. Therefore in most cases, unless a script
requires a subshell, group commands are preferable to subshells. Group
commands are both faster and require less memory.

```bash
$ echo "foo" | read
$ echo $REPLY
```

Above examples shows a `read` command in a pipeline does not work as we might
intuitively expect. It is because the `read` command is executed in a subshell,
and its copy of `REPLY` is destroyed when the subshell terminates.

<a id="ch36-process-substitution"></a>

Fortunately, the shell provides an exotic form of expansion called _process substitution_
that can be used to work around this problem.

Process substitution is expressed in two ways:

For processes that produce standard output:

```bash
<(list)
```

or, for process that intake standard input:

```bash
>(list)
```

We can employ process substitution like this to solve our problem with `read`:

```bash
read < <(echo "foo")
echo $REPLY
```

Process substitution allows us to treat the output of a subshell as an
ordinary file for purposes of redirection. In fact, since it is
a form of expansion, we can examine its real value:

```bash
$ echo <(echo "foo")
/dev/fd/13
```

We see that the output of the subshell is being provided by a file named `/dev/fd/63`

Process substitution is often with loops containing `read`. Here
is an example of a `read` loop that processes the contents of a directory listing
created by a subshell:

```bash
$ cat pro_sub
#!/bin/bash

# pro-sub: demo of process substitution

while read attr links owner group size date time filename; do
  cat <<- EOF
      Filename:   $filename
      Size:       $size
      Owner:      $owner
      Group:      $group
      Modified:   $date $time
      Links:      $links
      Attributes: $attr
EOF
done < <(ls -l | tail -n +2)

$ pro_sub | head -n 20
      Filename:   12:11 arith-loop
      Size:       263
      Owner:      alexsoon
      Group:      staff
      Modified:   Jan 18
      Links:      1
      Attributes: -rwxr--r--@
      Filename:   12:30 array-2
      Size:       826
      Owner:      alexsoon
      Group:      staff
      Modified:   Jan 23
      Links:      1
      Attributes: -rwxr--r--@
      Filename:   16:24 array-sort
      Size:       186
      Owner:      alexsoon
      Group:      staff
      Modified:   Jan 22
      Links:      1
```

#### Traps

We can add responding signals capability to our scripts, too. Larger
and more complicated scripts may benefit from having a signal handling
routine.

When we design a large, complicated script, it is important to consider
what happens if the user logs off or shuts down the computer while the script
is running. When such an event occurs, a signal will be sent to all affected
processes. In turn, the programs representing those processes can perform actions
to ensure a proper and orderly termination of the program.

Let's say, for example, that we wrote a script that created a temporary file
during its execution. In the course of good design, we would have the script delete
the file when the script finishes it work. It would also be smart to have the script
delete the file if a signal is received indicating that the program was going to be
terminated prematurely.

```bash
trap argument signal [signal...]
```

where argument is a string which will be read and treated as a command
and _signal_ is the specification of a signal that will trigger the execution
of the interpreted command.

```bash
#!/bin/bash

# trap-demo: simple signal handling demo

trap "echo 'I am ignoring you.'" SIGINT SIGTERM

for i in {1..5}; do
  echo "Iteration $i of 5"
  sleep 5
done
```

When user attempts to stop the script by pressing `Ctrl-C`, the `echo 'I am ignoring you.'` will be executed.

Constructing a string to form a useful sequence of commands can be awkward,
so it is common practice to specify a shell function as the command.

```bash
#!/bin/bash

# trap-demo2: simple signal handling demo

exit_on_signal_SIGINT () {
  echo "Script interrupted." 2>&1
  exit 0
}

exit_on_signal_SIGTERM () {
  echo "Script terminated." 2>&1
  exit 0
}

trap exit_on_signal_SIGINT SIGINT
trap exit_on_signal_SIGTERM SIGTERM

for i in {1..5}; do
  echo "Iteration $i of 5"
  sleep 5
done
```

---

> **Temporary Files**

One reason signal handlers are included in scripts is to remove temporary files
that the script may create to hold intermediate results during execution.

Traditionally, programs on Unix-like systems create their temporary files in the
`/tmp` directory, a shared directory intended for such files. However, since
the directory is shared, this poses certain security concerns,
particularly for programs running with superuser privileges.
Aside from the obvious step of setting proper permissions for files exposed to all
users of the system, **it is important to give temporary files non-predictable filenames**. This avoids an exploit known as a _temp race attack_.

One way to create a non-predictable (but still descriptive)
name is to do something like this:

```bash
tempfile=/tmp/$(basename $0).$$.$RANDOM
```

`$$` contains the process ID. `$RANDOM` is a random number between `1-32767`,
which is not a very large range in computer terms, so a single instance
of the variable is not sufficient to overcome a determined attacker.

A better way is to use the `mktemp` program (not to be confused
with the `mktemp` standard library function) to both name and create the
temporary file. The `mktemp` program accepts a template as an argument
that is used to build the filename. The template should include a series of "X"
characters, which are replaced by a corresponding number of random letters and
numbers. The longer the series of "X" **characters**, the longer the series of random
characters.

```bash
tempfile=$(mktemp /tmp/foo.bar.$$.XXXXXXXXX)
```

For scripts that are executed by regular users, it may be wise to avoid the
use of the `/tmp` directory and create a directory for temporary files within
the user's home directory, with a line of code such as this:

```bash
[[ -d $HOME/tmp ]] || mkdir $HOME/tmp
```

---

#### Asynchronous Execution

`bash` has a builtin command to help manage _asynchronous execution_ such as this.
The `wait` command causes a parent script to pause until a specified process
(i.e., the child script) finishes

##### `wait`

We have two scripts, a parent script:

```bash
#!/bin/bash

# async-parent: Asynchronous execution demo (parent)

echo "Parent: starting..."

echo "Parent: launching child script..."
async-child &
pid=$!
echo "Parent: child (PID= $pid) launched"

echo "Parent: continueing"
sleep 2

echo "Parent: pausing to wait for child to finish..."
wait $pid

echo "Parent: child is finished. Continuing..."
echo "Parent: parent is done. Exiting."
```

and a child script:

```bash
#!/bin/bash

# async-child: Asynchronous execution demo (child)

echo "Child: child is running..."
sleep 5
echo "Child: child is done. Exiting."
```

Output:

```bash
$ async_parent
Parent: starting...
Parent: launching child script...
Parent: child (PID= 56371) launched
Parent: continueing
Child: child is running...
Parent: pausing to wait for child to finish...
Child: child is done. Exiting.
Parent: child is finished. Continuing...
Parent: parent is done. Exiting.
```

#### Named Pipes

In most Unix-like system, it is possible to create a special type of
file called a _named pipe_. Named pipes are used to create a connection
between two processes and can be used just like other types of files.
They are not that popular, but they're good to know about.

There is a common programming architecture called _client-server_, which
can make use of a communication method such as named pipes, as well as other
kinds of _interprocess communication_ such as network connections.

Named pipes behave like files, but actually form first-in first-out (FIFO)
buffers. With named pipes, it is possible to set up something like this:

```bash
process1 > named_pipe
```

and

```bash
process2 < named_pipe
```

and it will behave as if:

```bash
process1 | process2
```

##### Setting Up A Named Pipe

Creating a named pipe with `mkfifo` command:

```bash
$ mkfifo pipe1
$ ls -l pipe1
prw-r--r--@ 1 alexsoon  staff  0 Jan 23 15:25 pipe1
```

**Tip**: `p` indicating that it is a named pipe.

##### Using Named Pipes

To demonstrate how the named pipe works, we will need two terminal
windows. In the first terminal, we enter a simple command and redirect
its output to the named pipe:

```bash
$ ls -l > pipe1
```

After we press the `Enter` key, the command will appear to hang. This is because
there is nothing receiving data from the other end of the pipe yet. We say it is
`blocked`.

Using the second terminal window, we enter this command:

```bash
$ cat < pipe1
```

and the directory listing produced from the first terminal window
appears in the second terminal as the output from the `cat` command.
The first terminal successfully completes once it is no longer blocked.

## Other

| Command    | Description                                      |
| ---------- | ------------------------------------------------ |
| `df`       | Display free disk space                          |
| `free`     | Display free memory                              |
| `du`       | display disk usage statistics                    |
| `basename` | Return filename or directory portion of pathname |
| `time`     | Executes and times the specified utility         |

### Variable

| Variable  | Description                                         |
| --------- | --------------------------------------------------- |
| `$$`      | process ID                                          |
| `$RANDOM` | random number `1-32767`                             |
| `$!`      | The PID of the last job **put into the background** |
| `$?`      | Status of previous command exit status              |
