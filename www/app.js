// Global state
let wasmLoaded = false;
let pdateParse = null;

// DOM elements
const input = document.getElementById('datetime-input');
const parseBtn = document.getElementById('parse-btn');
const clearBtn = document.getElementById('clear-btn');
const nowBtn = document.getElementById('now-btn');
const loading = document.getElementById('loading');
const error = document.getElementById('error');
const results = document.getElementById('results');

// Show loading spinner initially
loading.classList.remove('hidden');

// Initialize WASM
async function initWasm() {
    try {
        const go = new Go();
        const result = await WebAssembly.instantiateStreaming(
            fetch('pdate.wasm'),
            go.importObject
        );

        // Run the Go program
        go.run(result.instance);

        // Wait a bit for the Go program to register the function
        await new Promise(resolve => setTimeout(resolve, 100));

        // Check if the function is available
        if (typeof window.pdateParse === 'function') {
            pdateParse = window.pdateParse;
            wasmLoaded = true;
            loading.classList.add('hidden');
            console.log('pdate WASM module loaded successfully');
            // Focus the input field once WASM is loaded
            input.focus();
        } else {
            throw new Error('pdateParse function not found');
        }
    } catch (err) {
        console.error('Failed to load WASM module:', err);
        loading.classList.add('hidden');
        showError(`Failed to load WASM module: ${err.message}`);
    }
}

// Show error message
function showError(message) {
    error.textContent = message;
    error.classList.remove('hidden');
    results.classList.add('hidden');
}

// Hide error message
function hideError() {
    error.classList.add('hidden');
}

// Display results
function displayResults(data) {
    hideError();

    // Input section
    document.getElementById('result-input').textContent = data.input;
    document.getElementById('result-parsed').textContent = data.parsed;

    // UTC section
    document.getElementById('result-utc-formatted').textContent = data.utc.formatted;
    document.getElementById('result-utc-rfc3339').textContent = data.utc.rfc3339;
    document.getElementById('result-utc-rfc1123z').textContent = data.utc.rfc1123z;
    document.getElementById('result-utc-unix').textContent = data.utc.unix;
    document.getElementById('result-utc-unix-milli').textContent = data.utc.unixMilli;
    document.getElementById('result-utc-unix-nano').textContent = data.utc.unixNano;

    // Local section
    document.getElementById('result-local-formatted').textContent = data.local.formatted;
    document.getElementById('result-local-timezone').textContent = data.local.timezone;

    // Relative section
    document.getElementById('result-relative-formatted').textContent = data.relative.formatted;

    results.classList.remove('hidden');
}

// Parse datetime
function parseDateTime() {
    if (!wasmLoaded) {
        showError('WASM module is still loading. Please wait...');
        return;
    }

    const value = input.value.trim();
    if (!value) {
        showError('Please enter a datetime string or ULID');
        return;
    }

    try {
        const result = pdateParse(value);

        if (result.error) {
            showError(`Parse error: ${result.error}`);
        } else {
            displayResults(result);
        }
    } catch (err) {
        showError(`Unexpected error: ${err.message}`);
    }
}

// Clear input and results
function clearAll() {
    input.value = '';
    hideError();
    results.classList.add('hidden');
    input.focus();
}

// Use current time
function useNow() {
    // Get current Unix timestamp
    const now = Math.floor(Date.now() / 1000);
    input.value = now.toString();
    parseDateTime();
}

// Event listeners
parseBtn.addEventListener('click', parseDateTime);
clearBtn.addEventListener('click', clearAll);
nowBtn.addEventListener('click', useNow);

input.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
        parseDateTime();
    }
});

// Copy to clipboard functionality
function setupCopyButtons() {
    document.querySelectorAll('.copy-btn').forEach(button => {
        button.addEventListener('click', async (e) => {
            const targetId = button.getAttribute('data-copy-target');
            const targetElement = document.getElementById(targetId);

            if (!targetElement) return;

            const textToCopy = targetElement.textContent;

            try {
                await navigator.clipboard.writeText(textToCopy);

                // Visual feedback
                button.classList.add('copied');
                const originalTitle = button.getAttribute('title');
                button.setAttribute('title', 'Copied!');

                // Reset after 2 seconds
                setTimeout(() => {
                    button.classList.remove('copied');
                    button.setAttribute('title', originalTitle);
                }, 2000);
            } catch (err) {
                console.error('Failed to copy:', err);
                // Fallback for browsers that don't support clipboard API
                const textArea = document.createElement('textarea');
                textArea.value = textToCopy;
                textArea.style.position = 'fixed';
                textArea.style.left = '-999999px';
                document.body.appendChild(textArea);
                textArea.select();
                try {
                    document.execCommand('copy');
                    button.classList.add('copied');
                    setTimeout(() => button.classList.remove('copied'), 2000);
                } catch (fallbackErr) {
                    console.error('Fallback copy failed:', fallbackErr);
                }
                document.body.removeChild(textArea);
            }
        });
    });
}

// Initialize WASM on page load
initWasm();

// Setup copy buttons after results are displayed
const originalDisplayResults = displayResults;
displayResults = function(data) {
    originalDisplayResults(data);
    setupCopyButtons();
};
