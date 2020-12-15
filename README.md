# Taskit

Taskit is a task manager that tells you what's need to be worked, being worked on and where something is in a process. Imagine a
white board, filled with lists of sticky notes, with each note as a task for you. And everything build in a CLI program! The 
frontend was built using Go, which I'm still new to (actually this is my first project using Go). For the backend, I've used the
old good Node.js.

# Install

First, clone this repository. Inside the backend folder (taskit-webserver), you can find some scripts that will help you install and uninstall
the local server. Pretty easy, just run it with 

```
./runWebserver
```

which will be built as a docker container when you install. So, the only thing you need
to have in your machine to run the taskit webserver is docker! For the frontend, I also made available some scripts to help with the 
installation. Just run the installFrontend script using

```
./installFrontend
```

inside the taskit-frontend folder. After running the command, you can restart the pc
to apply the changes, or, if you prefer, run the command 

```
source ~/.profile
```

to be able to use the taskit in that terminal session (only). After installed, you can use taskit by just typing 

```
taskit [command]
```

Remember that you must have the webserver installed in your machine. And that's it. Everything ready to go! Down below I'll list
and explain the commands you can give to taskit to start building your CLI task board. Feel free to use the -h flag if you need
help.

# Commands

## board

```
taskit board [flags]
```

The board command allows you to see all of your tasks, including all of their properties (like the taskID, which will be needed in other commands).
The board organizes tasks according to their current status. Possible status are:

- ToDo
- Working
- Done
- Closed

The board command does not need any arguments and accepts the following flags, which will filter the tasks according to their properties:

- -c, --createdAt string   Filter tasks by created day (added time)
- -d, --deadline string    Filter tasks by deadline
- -h, --help               help for board
- -p, --priority int       Filter tasks by priority (default -1)

## create

```
taskit create
```

The create command allows you to create a new task. This command will go through the steps to create a task. He receives no
argument and has no flag. Only the title of the task and its status are mandatory at the time of creation. Feel free to
skip optional steps!

## set

```
taskit set [taskID] [property] [newValue]
```
The set command will change the property passed as an argument to a new value. That simple! The list of properties available for change are:

- title
- description
- status
- priority (task priority. Highest prioritized task appears first in the task board)
- deadline (the day that marks the deadline for the task). Must be of the format YYYY-MM-DD or "no-deadline", to remove the deadline
- timeEstimate (estimated time to complete the task, in hours)

Usage example:

```
taskit set 1 title "A new title"
```
```
taskit set 1 deadline 2020-12-15
```

## remove

```
taskit remove [taskID]
```
The remove command will transfer the task to Closed status. If the task has been completed, it will be stored as completed,
otherwise it will be stored as non-completed.

Usage example:

```
taskit remove 1
```

## block

```
taskit block [taskID]
```

The block command will block the task. A blocked task cannot be passed to Done status. A task can be blocked in two ways
(when running the command, you will be asked which way to go):
- Locked until unblocked with the unblock command
- Blocked by another task. In this case, if the blocking task is passed to Done, then the blocked task will be automatically unlocked.
Note: Tasks blocked by another task can also be unlocked by the unblock command

Usage example:

```
taskit block 1
> Is blocked by other task? (y/n): y
> Task that is blocking this task: 2
```

## unblock

```
taskit unblock [taskID]
```

The unblock command will unlock the task, regardless of the type of block that was used.
Usage example:

```
taskit unblock 1
```

## time

```
taskit time [taskID]
```

The time command will show the time that the task spent in Working status

```
taskit time 1
```