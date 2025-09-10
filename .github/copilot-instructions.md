# Pixop - OpenGL Graphics Visualization Tool

Pixop is a Go application that creates mathematical visualizations using the Pixel graphics library. It includes Lissajous curves, Koch snowflakes, epicycles, and gem patterns.

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

### System Dependencies - CRITICAL FIRST STEP
Before building, you MUST install the required X11 and OpenGL development libraries:

```bash
sudo apt-get update
sudo apt-get install -y libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libgl1-mesa-dev libxi-dev pkg-config libxxf86vm-dev
```

**NEVER skip this step** - the build will fail with missing X11/OpenGL headers without these libraries.

### Build and Test Process
Bootstrap, build, and test the repository:

```bash
# Navigate to project root
cd /path/to/pixop

# Clean and download dependencies (first time or after go.mod changes)
go mod tidy

# Build the application - NEVER CANCEL: Set timeout to 5+ minutes minimum
go build -o pixop
# Build time: ~0.8 seconds normally, but up to 30 seconds on first build due to C compilation

# Run tests - NEVER CANCEL: Set timeout to 5+ minutes minimum  
go test ./...
# Test time: ~10 seconds, includes math integration tests

# Verify binary was created
ls -la pixop
./pixop --help
```

### Running the Application
**IMPORTANT**: This is a GUI application that requires a display/X11 server:

```bash
# Check available commands
./pixop --help

# Run visualization commands (requires display)
./pixop lissajous
./pixop koch
./pixop epicycles  
./pixop gem

# Custom window size
./pixop --width 1920 --height 1080 lissajous
```

**Note**: In headless environments, the application will fail with "X11: The DISPLAY environment variable is missing" - this is expected behavior.

## Validation

### Build Validation
- Always run `go build -o pixop` after making changes - it must succeed without errors
- Run `go vet ./...` to check for Go programming errors - it must pass cleanly
- Run `go test ./...` to ensure all tests pass - currently only math package has tests

### Code Quality
- Use `gofmt -d .` to check code formatting - output should be empty
- Code is already well-formatted and passes `go vet`
- No additional linters like golint or staticcheck are currently configured

### Functional Testing
**CANNOT be done in headless environments** - the application requires an X11 display to run visualizations.

In environments with a display:
1. Run `./pixop lissajous` and verify a window opens with animated sinusoidal patterns
2. Test other visualization modes: `koch`, `epicycles`, `gem`
3. Test custom window sizes with `--width` and `--height` flags

## Common Tasks

### Repository Structure
```
pixop/
├── main.go              # CLI application entry point
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
├── cmd/                 # Scene implementations
│   ├── cmd.go          # Scene interface definition
│   ├── lissajous/      # Lissajous curve visualization
│   ├── koch/           # Koch snowflake visualization  
│   ├── epicycles/      # Epicycle animation
│   └── gem/            # Gem pattern visualization
├── global/             # Global state and timing
│   └── global.go      # StartTime, Time, Bounds variables
└── math/               # Mathematical utilities
    ├── constants.go    # Mathematical constants (TwoPi)
    ├── oscillators.go  # Oscillator functions
    ├── integration.go  # Numerical integration
    └── integration_test.go # Math tests
```

### Key Dependencies
- `github.com/faiface/pixel` - Main graphics library
- `github.com/urfave/cli` - Command line interface
- `github.com/stretchr/testify` - Testing assertions
- System: X11, OpenGL, GLFW libraries

### Troubleshooting Build Issues

**"X11/Xlib.h: No such file or directory"**:
```bash
sudo apt-get install -y libx11-dev
```

**"cannot find -lXxf86vm"**:
```bash
sudo apt-get install -y libxxf86vm-dev
```

**"Package 'gl', required by 'virtual:world', not found"**:
```bash
sudo apt-get install -y libgl1-mesa-dev pkg-config
```

### Development Workflow
1. Make code changes
2. Run `go build -o pixop` to verify compilation
3. Run `go test ./...` to verify tests still pass
4. Run `go vet ./...` to check for programming errors
5. For GUI changes, test with `./pixop [command]` in an environment with display

### Performance Expectations
- **Build time**: 0.8 seconds (normal), up to 30 seconds (first build with C compilation)
- **Test time**: ~10 seconds
- **Application startup**: Instant in GUI environments
- **NEVER CANCEL** builds or tests - use timeouts of 5+ minutes minimum

### Common Gotchas
- The application REQUIRES a display - will not run in headless environments
- System libraries must be installed before building - there are no pure Go alternatives
- The CLI has a typo: "hidth" instead of "width" in the help text (in main.go line 43)
- Graphics performance depends on GPU drivers and OpenGL support