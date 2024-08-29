### INFO

Hello, world!
It's my first application: cli-task-tracker
The idea for this project was taken from [roadmap site](https://roadmap.sh/projects/task-tracker)

### Requirements

go v1.23.0

### Installation
To begin you need to download this:
```sh
git clone https://github.com/HoneyVakkaria/cli-task-tracker.git
```

Next, you can use makefile to install:
```sh
sudo make install # to install

sudo make uninstall # to remove
```

Or do it manually:
```sh
# To install
go build -o task-tracker
sudo cp task-tracker /usr/local/bin/
rm -f task-tracker

# To remove
sudo rm -f /usr/local/bin/task-tracker
```
> **Note:**
>
> The installation is only for Linux and macOS users.


### How to use

Here are some examples that cover the full functionality of the project:
```sh
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```
