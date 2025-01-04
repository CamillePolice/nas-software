# File Organizer

A cross-platform command-line tool to organize files in a directory by automatically sorting them into appropriate folders.

## Usage

Run the program with the `-path` flag followed by the directory path you want to organize. The command varies based on your operating system:

### Linux

```bash
./file-organizer-linux-amd64 -path "/path/to/your/directory"
```

### macOS

```bash
./file-organizer-darwin-amd64 -path "/path/to/your/directory"
```

### Windows

```bash
.\file-organizer-windows-amd64.exe -path "C:\path\to\your\directory"
```

### Flags

- `-path`: (Required) Specifies the directory path to organize files

### Examples

Linux:

```bash
./file-organizer-linux-amd64 -path "/home/user/Downloads"
```

macOS:

```bash
./file-organizer-darwin-amd64 -path "/Users/username/Downloads"
```

Windows:

```bash
.\file-organizer-windows-amd64.exe -path "C:\Users\username\Downloads"
```

## Scheduling Automated Organization

### Linux (Cron)

1. Open terminal and edit crontab:

```bash
crontab -e
```

2. Add this line to run daily at 2 AM (adjust the path accordingly):

```bash
0 2 * * * /absolute/path/to/file-organizer-linux-amd64 -path "/path/to/organize"
```

### macOS (Cron)

1. Open terminal and edit crontab:

```bash
crontab -e
```

2. Add this line to run daily at 2 AM (adjust the path accordingly):

```bash
0 2 * * * /absolute/path/to/file-organizer-darwin-amd64 -path "/path/to/organize"
```

### Windows (Task Scheduler)

1. Open Task Scheduler
2. Click "Create Basic Task"
3. Set Name: "File Organizer" and Description
4. Set Trigger: Daily
5. Choose start time (e.g., 2:00 AM)
6. Action: Start a program
7. Program/script: Path to `file-organizer-windows-amd64.exe`
8. Add arguments: `-path "C:\path\to\organize"`
9. Finish the wizard

Common Cron Schedule Examples:

- Every hour: `0 * * * *`
- Every day at 2 AM: `0 2 * * *`
- Every Sunday at 3 AM: `0 3 * * 0`
- Every first day of the month at 4 AM: `0 4 1 * *`

The program will automatically create necessary folders and organize your files based on their types.
