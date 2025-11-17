# pdate Web Interface

A WebAssembly-powered web interface for parsing datetime strings and ULIDs.

## Building

From the project root directory:

```bash
make wasm
```

This will:
1. Compile the Go code to WebAssembly (`pdate.wasm`)
2. Copy the required `wasm_exec.js` file from your Go installation

## Running Locally

From the project root directory:

```bash
make serve
```

Then open your browser to http://localhost:8080

## Usage

1. Enter a datetime string, Unix timestamp, or ULID in the input field
2. Click "Parse" or press Enter
3. View the parsed results in various formats

### Supported Input Formats

- Unix timestamps (seconds, milliseconds, microseconds, nanoseconds)
- Unix timestamps with fractional seconds (e.g., `1705318200.5`)
- ISO 8601 / RFC3339 (e.g., `2024-01-15T10:30:00Z`)
- Various other datetime formats
- ULIDs (e.g., `01HQVXYZ1234567890ABCDEFGH`)

## Files

- `index.html` - Main HTML page
- `styles.css` - Styling
- `app.js` - JavaScript application logic
- `pdate.wasm` - Compiled WebAssembly binary (generated)
- `wasm_exec.js` - Go WASM runtime support (copied from Go installation)
