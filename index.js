const WASM_PATH = "main.wasm";

async function init() {
  const importObj = { module: {} };

  const go = new Go();

  const response = await fetch(WASM_PATH);
  const buffer = await response.arrayBuffer();
  const obj = await WebAssembly.instantiate(buffer, go.importObject);

  go.run(obj.instance);
}

init();
