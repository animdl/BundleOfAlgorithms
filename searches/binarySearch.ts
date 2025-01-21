function binarySearch(arr: number[], target: number): number {
    let lowIndex: number = 0, 
        highIndex: number = arr.length - 1;

    if (lowIndex > highIndex || target < arr[lowIndex] || target > arr[highIndex]) {
        return -1; // sentinel value returned if target is out of bounds
    }

    while (lowIndex <= highIndex) {
        let midIndex = Math.floor((lowIndex + highIndex) / 2);

        if (target === arr[midIndex]) {
            return midIndex;
        } else if (target > arr[midIndex]) {
            lowIndex = midIndex + 1;
        } else {
            highIndex = midIndex;
        }
    }

    return -1;
}

console.log(binarySearch([1, 3, 4, 7, 12, 32], 7));