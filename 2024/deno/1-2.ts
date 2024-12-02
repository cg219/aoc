export function populateLists(data: string, list1: number[], list2: number[]) {
    const res = data.trim().match(/(\d+)\s+(\d+)/)

    if (res) {
        list1.push(Number(res[1]))
        list2.push(Number(res[2]))
    }
}

async function *iterate(src: string) {
    let data = []; 
    
    for (const s of src) {
        if (s == "\n") {
            yield data.join("")
            data = []
            continue;
        }

        if (s == "\r") continue;

        data.push(s)
    }
}

if (import.meta.main) {
    const data = await Deno.readTextFile("1.txt") 
    const list1: number[] = []
    const list2: number[] = []
    const cache: Map<number, number> = new Map()
    const res: number[] = []

    for await (const line of iterate(data)) {
        populateLists(line, list1, list2)
    }

    list2.forEach((v) => cache.set(v, (cache.get(v) ?? 0) + 1))

    list1.forEach((v) => {
        if (cache.has(v)) {
            res.push(cache.get(v) * v)
        }
    })

    console.log(res.reduce((total, val) => val + total, 0))
}

