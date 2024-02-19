import { toFileUrl, join } from "https://deno.land/std@0.208.0/path/mod.ts";

async function loadGames() {
    const res = await fetch(toFileUrl(join(Deno.cwd(), "data.txt")));

    if (!res.ok || !res.body) return;

    const decoder = new TextDecoder();
    const lines: Uint8Array[] = [];
    let remainder = false;

    for await (const chunk of res.body) {
        let i = 0;
        const len = chunk.length;
        let start = 0;

        for (i; i < len; i++) {
            if (chunk[i] == 10) {
                if (remainder) {
                    let old = lines[lines.length - 1];
                    let merged = new Uint8Array(old.length + i);
                    
                    merged.set(old);
                    merged.set(chunk.slice(start, i), old.length);

                    lines[lines.length - 1] = merged;
                    remainder = false;
                } else {
                    lines.push(chunk.slice(start, i));
                }

                start = i + 1;
            }
        }

        if (start <= chunk.length) {
            lines.push(chunk.slice(start, i + 1));
            remainder = true;
        }
   }

   lines.forEach((l) => console.log(decoder.decode(l)));

}

await loadGames();
