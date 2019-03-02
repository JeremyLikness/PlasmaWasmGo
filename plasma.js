window.plasmaRender = (w, h, ctx, buffer64) => {
    let bytes = atob(buffer64);
    let buffer = new Uint8ClampedArray(bytes.length);
    for (let i = 0; i < bytes.length; i+=1) {
        buffer[i] = bytes.charCodeAt(i);
    }
    let imageData = new ImageData(buffer, w, h);
    ctx.putImageData(imageData, 0, 0);
};

if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("plasma.wasm"), go.importObject).then((result) => {
    mod = result.module;
    inst = result.instance;
    console.clear();
    go.run(inst);
    inst = WebAssembly.instantiate(mod, go.importObject);
}).catch((err) => {
    console.error(err);
});