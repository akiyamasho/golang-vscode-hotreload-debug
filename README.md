# Golang + Hot Reloading + VSCode On-demand Debugging

This is a simple repo that allows both hot-reloading and debugging **on-demand**.

Note that debugging is only available on-demand, which means that you have to attach the debugger to the app when you need to debug or on app reload.
Automatically re-attaching the debugger on app reload is a work in progress.

### Requirements

- `go` (tested on v1.21)
- `air`
    ```bash
    go install github.com/cosmtrek/air@latest
    ```
- `delve`
    ```bash
    go install github.com/go-delve/delve/cmd/dlv@latest
    ```

### Deployment
![golang-debug-hotreload](https://github.com/akiyamasho/golang-vscode-hotreload-debug/assets/35907066/8fe87e12-1c8f-40dd-aba1-2ee4e99f378a)

1. Run `air` in the root directory of the project to run the app
    ```bash
    air
    ```
    This will run both `air` and attach a `dlv` debugger to the app for later attachment.

2. Update the code as needed and it will hot-reload.
    
3. (debugging) If you want to debug and use breakpoints, run the `Attach to dlv` launch configuration on VSCode and set breakpoints for debugging.

4. (debugging) Delve will now stop at breakpoints.

    NOTE: You can try it by setting a breakpoint in the `main()` function in `main.go` and running a `curl` command to the `/health` endpoint of the app.
    ```bash
    curl localhost:8000/health
    ```
