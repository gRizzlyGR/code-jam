import readline from 'readline';

let cases = 0;

const reader = readline.createInterface({
    input: process.stdin,
});

reader.on('line', (line: string) => {
    const parts = line.split(' ')
    const mural = parts[2];

    let cost = 0;
    // First line is the number of cases, the array will be empty, therefore
    // `mural` is undefined
    if (mural) {
        const x = Number(parts[0]);
        const y = Number(parts[1]);
        let prev = '';
        [...mural].forEach(letter => {
            switch (letter) {
                case 'C':
                    if (prev === 'J') {
                        cost += y;
                    }
                    prev = letter
                    break;
                case 'J':
                    if (prev === 'C') {
                        cost += x;
                    }
                    prev = letter
                    break;
                default: {
                    // Ignore '?'
                }
            }
        })
        cases++
        console.log(`Case #${cases}: ${cost}`)
    }
})



