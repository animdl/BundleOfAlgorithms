function jumpSearch(arr: number[], target: number): number {
    // check bounds
    if (arr[arr.length - 1] < target || arr[0] > target) {
        return -1;
    }

    const jumpAmount: number = Math.floor(Math.sqrt(arr.length));
    let startIndex: number = 0;

    // find section
    while (arr[startIndex + jumpAmount - 1] < target) {

        if (startIndex + jumpAmount >= arr.length) {
            break;
        } else {
            startIndex += jumpAmount;
        }
    }

    // linear search through section
    for (let i = startIndex; i < Math.min(startIndex + jumpAmount, arr.length); i++) {
        if (arr[i] === target) {
            return i;
        }
    }

    return -1;
}

console.log(jumpSearch([1, 3, 4, 7, 12, 32], 7));