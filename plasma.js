window.plasmaRender = (w, h, ctx, buffer64) => {
    let buffer = atob(buffer64);
    for (let i = 0; i < h; i += 1) {
        for (let j = 0; j < w; j += 1) {
            let base = (i * w + j) * 3;
            ctx.fillStyle = `rgb(${buffer.charCodeAt(base)},${buffer.charCodeAt(base+1)}, ${buffer.charCodeAt(base+2)})`;
            ctx.fillRect(j, i, j + 1, i + 1);
        }
    }
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