import { toFileUrl } from "https://deno.land/std@0.208.0/path/to_file_url.ts";
import { resolve } from "https://deno.land/std@0.208.0/path/resolve.ts";
import { CHAR_0 } from "https://deno.land/std@0.208.0/path/_common/constants.ts";
class AdventDayOne {
    private dataFile;
    private words: Set<string> = new Set();
    private wordMap: Map<string, number> = new Map();

    constructor(filepath: string) { 
        this.dataFile = filepath;    
        
        const list = [ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" ];
        
        list.forEach((word, index) => {
            this.wordMap.set(word, index + 1);
            this.words.add(word);
        })

    }

    async run2() {
        const res = await fetch(toFileUrl(resolve(Deno.cwd(), this.dataFile)));
        const decoder = new TextDecoder();

        let holder = '';
        let total: [number]= [0]
        let found: [number?] = [];

        if (!res.body) return;
        
        let s = '';

        for await (const chunk of res.body) {
            const decodedChunk = decoder.decode(chunk);
            const chunkedarray = decodedChunk.split('');

            for (const letter of chunkedarray) {
                if (letter == '\n') {
                    s.split('').forEach((letter) => {
                        if (parseInt(letter) >= 0 ) {
                            found.push(Number(letter));
                            holder = holder.slice(-1);
                        } else {
                            holder = `${holder}${letter}`;

                            for (const [word] of this.words.entries()) {
                                if (holder.includes(word)) {
                                    found.push(this.wordMap.get(word));
                                    holder = holder.slice(-1);
                                    break;
                                }
                            }
                        }

                    });

                    if (found.length) {
                        total.push(Number(`${found.at(0)}${found.at(-1)}`))
                    }

                    found = [];
                    holder = '';
                    s = '';
                } else {
                    s = `${s}${letter}`
                }
            }
        }

        console.log('Run2:', total.reduce((t, v) => t + v, 0));
        return total;
    }

    async run1() {
        let total = 0;
        const res = await fetch(toFileUrl(resolve(Deno.cwd(), this.dataFile)));
        const decoder = new TextDecoder();

        if (!res.body) return;
            
        let s = '';

        for await (const chunk  of res.body) {
            const decodedChunk = decoder.decode(chunk);
            const chunkedarray = decodedChunk.split('');

            for (const letter of chunkedarray) {
                if (letter == '\n') {
                    const numString: string = s.replace(/[A-Za-z]+/g, '').trim();

                    if (numString) {
                        total += Number(`${numString.at(0)}${Number(numString.at(-1))}`);
                    }
                    s = '';
                }

                s = `${s}${letter}`;
            }
            
        }

        console.log('Run1:', total)
        return total;
    }

}

const filepath = "data.txt";
const ad = new AdventDayOne(filepath);

// await ad.run1();
 await ad.run2();

