# Taskit

Taskit is a task manager that tells you what's need to be worked, being worked on and where something is in a process. Imagine a
white board, filled with lists of sticky notes, with each note as a task for you. And everything build in a CLI program! The 
frontend was built using Go, which I'm still new to (actually this is my first project using Go). For the backend, I've used the
old good Node.js. Inside the backend folder (taskit-webserver), you can find some scripts that will help you install and uninstall
the local server (pretty easy, just run it), which will be built as a docker container when you install. So, the only thing you need
to have in your machine to run the taskit webserver is docker! For the frontend, the only thing you need is the taskit binary. Run
it with

```
./taskit [command]
```

Remember that you must have the webserver installed in your machine. And that's it. Everything ready to go! Down below I'll list
and explain the commands you can give to taskit to start building your CLI task board.