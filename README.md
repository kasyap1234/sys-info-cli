

# sys-info-cli

sys-info-cli is a powerful command-line interface tool for retrieving and displaying system information on your computer. It provides detailed statistics about CPU, RAM, disk usage, processes, and more.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install sys-info-cli, follow these steps:

1. Ensure you have Go 1.23.1 or later installed on your system.
2. Clone the repository:


README.md
git clone https://github.com/yourusername/sys-info-cli.git

3. Navigate to the project directory:



cd sys-info-cli

4. Build the project:


go build -o sysinfo

5. Move the binary to a directory in your PATH:



sudo mv ./sysinfo /usr/local/bin/


## Usage

After installation, you can use sys-info-cli by running the `sysinfo` command followed by a subcommand:




sysinfo [command]


## Commands

sys-info-cli supports the following commands:

- `cpu`: Display CPU statistics and information
- `ram`: Show RAM usage statistics
- `disk`: List disk partitions and their details
- `processList`: Display information about running processes
- `List Pids`: Show all process IDs currently running on the system
- `ExistsPid`: Check if a process with a given PID exists

### CPU Information

To get CPU statistics:



sysinfo cpu


This command displays CPU usage percentage and detailed information about each CPU core.

### RAM Information

To get RAM statistics:




sysinfo ram


This command shows total RAM, used RAM, free RAM, and usage percentage.

### Disk Information

To get disk partition information:




sysinfo disk


This command lists all disk partitions with their device names, mount points, file system types, and mount options.

### Process List

To get a list of running processes:




sysinfo processList


This command displays information about all running processes, including PID, name, status, CPU usage, and memory usage.

### List PIDs

To list all running process IDs:


sysinfo List Pids


### Check if PID Exists

To check if a specific PID exists:




sysinfo ExistsPid [PID]


Replace `[PID]` with the process ID you want to check.

## Configuration

sys-info-cli uses a configuration file located at `$HOME/.sysinfo.yaml` by default. You can specify a different configuration file using the `--config` flag:



sysinfo --config /path/to/config.yaml [command]


## Contributing

Contributions to sys-info-cli are welcome! Please follow these steps to contribute:

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Make your changes and commit them with descriptive commit messages
4. Push your changes to your fork
5. Submit a pull request to the main repository
