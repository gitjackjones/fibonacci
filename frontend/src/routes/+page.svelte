<script>
    let number = 1;
    let result = '';
    let isError = false;
  
    async function getFibonacciSequence() {
        isError = false;
    try {
        const response = await fetch(`http://localhost:8080/api/fibonacci/${number}`);
        if (response.ok) {
            const data = await response.json();
            result = data.join(', ');
        } else {
            const errorMessage = await response.text();
            console.error(`Failed to process the fibonacci sequence:`, errorMessage);
            result = `Failed to process the fibonacci sequence: ` + errorMessage;
            isError = true;
        }
    } catch (error) {
        console.error(`Failed to fetch the fibonacci sequence:`, error);
        result = `Failed to fetch the fibonacci sequence: ` + error.message;
        isError = true;
    }
}
  </script>
  
  <div class="min-h-screen flex flex-col justify-center bg-gray-800 text-white">
    <div class="max-w-md mx-auto space-y-6">
        <h1 class="text-xl font-semibold text-center">Jack's Fibonacci Sequence Generator</h1>
        <div class="flex items-center space-x-2">
            <input type="number" min="1" bind:value={number} class="border border-gray-600 bg-gray-700 text-white p-2 rounded w-full" placeholder="Enter a number"/>
            <button on:click={getFibonacciSequence} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                Generate
            </button>
        </div>
        <div>
            <p class="text-lg">
                Result: <span class={`font-bold ${isError ? 'text-red-500' : 'text-white'}`}>{result}</span>
            </p>
        </div>
    </div>
</div>
  