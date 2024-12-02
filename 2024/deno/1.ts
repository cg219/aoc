export function populateLists(data: string, list1: number[], list2: number[]) {
    const res = data.trim().match(/(\d+)\s+(\d+)/)

    if (res && res.length > 2) {
        list1.push(Number(res[1]))
        list2.push(Number(res[2]))
    }
}

function sortLists(...lists: number[][]) {
    lists.forEach((list) => {
        list.sort((a, b) => {
            if (a < b) return -1;
            if (a > b) return 1;
            return 0;
        })
    })
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

function process(a: number, b: number): Number {
    return Math.abs(a - b)
}

if (import.meta.main) {
    const data = await Deno.readTextFile("1.txt") 
    const list1: number[] = []
    const list2: number[] = []

    for await (const line of iterate(data)) {
        populateLists(line, list1, list2)
    }

    sortLists(list1, list2)

    const results = list1.map((v, i) => process(v, list2[i]))
    console.log(results.reduce((total, val) => val + total, 0))
}
