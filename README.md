# sys-info-cli


### Steps to Make Your CLI App Globally Accessible

To run your CLI command globally (like `sysinfo cpu`), you need to ensure that the binary is placed in a directory that is included in your system's PATH. Here’s how to do that:

1. **Move the Binary to a Directory in PATH:**
   You can move the compiled binary to a directory that is already in your PATH, such as `/usr/local/bin` or `~/bin`.

   ```bash
   sudo mv ./sysinfo /usr/local/bin/
   ```

   If you want to use `~/bin`, make sure it's added to your PATH. You can add the following line to your `~/.bashrc` or `~/.bash_profile`:

   ```bash
   export PATH="$HOME/bin:$PATH"
   ```

   After editing the file, run:

   ```bash
   source ~/.bashrc
   ```

2. **Test the Command:**
   After moving the binary, you should be able to run the command globally:

   ```bash
   sysinfo cpu
   ```

### Alternative: Using an Alias

If you prefer not to move the binary, you can create an alias in your shell configuration (like `.bashrc` or `.zshrc`):

```bash
alias sysinfo='/path/to/your/sysinfo'
```

Replace `/path/to/your/sysinfo` with the actual path to your binary. Then, run:

```bash
source ~/.bashrc
```

