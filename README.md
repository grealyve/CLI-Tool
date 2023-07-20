# CLI App

This is a Command Line Interface (CLI) application that provides weather forecast information for a given city. Additionally, it offers functionalities for pinging domains, checking disk usage, and creating a to-do list.

## Table of Contents

- [Features](#features)
- [Build](#build)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contributing](#contributing)

## Features

1. **Weather Forecast:** Obtain the weather forecast for a city of your choice by providing its name as an argument to the app.

2. **Ping a Domain:** Use the `cli-tool net ping -d <domain>` command to ping a specific domain and check its connectivity.

3. **Disk Usage Info:** Get disk usage statistics by using the `cli-tool info diskUsage -p <path>` command. It will show you the usage percentage of the specified disk path.

4. **To-Do List:** Create and manage your own to-do list with the `todo add/delete/complete` command. Keep track of your tasks and mark them as completed.

## Build

1. Navigate to the project directory.
```Bash
For Linux:

$ go build .

$ sudo mv ./cli-tool /usr/local/bin/

For Windows:

$ go build .
```


## Usage

1. Open a terminal or command prompt.

2. Navigate to the project directory.

3. Run the CLI application using the following command:

```Bash
$ cli-tool -h
```

## Commands

1. **Weather Forecast**
```Bash
$ cli-tool weather <Name of City>
```

2. **To-Do List**
```Bash
$ cli-tool todo add <anything you want>

$ cli-tool todo delete <Index of todo>

$ cli-tool todo complete <Index of todo>

To list the todo table:
$ cli-tool todo -l
```

3. **Ping a Domain**
```Bash
$ cli-tool net ping -d <domain>
```

4. **Disk Usage Info**

```Bash
$ cli-tool info diskUsage -p <path>
```

## Examples

1. Get the weather forecast for London:
```Bash 
$ ./cli-tool weather London
London, United Kingdom: 22C, Sunny, Humidity: %38
11:00 - 23C, Chance of rain: 0%, Partly cloudy
12:00 - 22C, Chance of rain: 0%, Partly cloudy
13:00 - 21C, Chance of rain: 0%, Cloudy
14:00 - 20C, Chance of rain: 68%, Patchy rain possible
15:00 - 18C, Chance of rain: 75%, Patchy rain possible                                                                                 
16:00 - 18C, Chance of rain: 71%, Light rain shower                                                                                    
17:00 - 17C, Chance of rain: 63%, Patchy rain possible                                                                                 
18:00 - 17C, Chance of rain: 81%, Patchy rain possible  
```

2. Ping a domain (e.g., google.com):
* PS: Ping only works on Windows for now... That will be updated.
```Powershell
.\cli-tool.exe net ping -d www.google.com

Pinging www.google.com [172.217.17.100] with 32 bytes of data:
Reply from 172.217.17.100: bytes=32 time=48ms TTL=114
Reply from 172.217.17.100: bytes=32 time=48ms TTL=114
Reply from 172.217.17.100: bytes=32 time=48ms TTL=114
Reply from 172.217.17.100: bytes=32 time=48ms TTL=114

Ping statistics for 172.217.17.100:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 48ms, Maximum = 48ms, Average = 48ms
```

3. Check disk usage for a path:
```Powershell
.\cli-tool.exe info diskUsage -p .
Free: 185238
Available: 185238
Size: 476922
Used: 291684
Usage: 61.15964 %
```
4. Create and manage your to-do list:
```Bash
To add a task to todo list:
$ ./cli-tool todo add <anything you want>

To delete a task:
$ ./cli-tool todo delete <Index of todo>

To complete a task:
$ ./cli-tool todo complete <Index of todo>

To list the todo table:
$ ./cli-tool todo -l

```


## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please feel free to create an issue or submit a pull request.
