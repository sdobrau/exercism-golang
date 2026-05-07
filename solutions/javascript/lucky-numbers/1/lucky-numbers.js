// @ts-check

/**
 * Calculates the sum of the two input arrays.
 *
 * @param {number[]} array1
 * @param {number[]} array2
 * @returns {number} sum of the two arrays
 */
export function twoSum(array1, array2) {
    const pattern = new RegExp(',', 'g');
    
    return Number(String(array1).replace(pattern,"")) +	Number(String(array2).replace(pattern,""))
}

/**
 * Checks whether a number is a palindrome.
 *
 * @param {number} value
 * @returns {boolean} whether the number is a palindrome or not
 */
export function luckyNumber(value) {
    let reversedNumber = ""
    for (let i = String(value).length - 1; i >= 0; i--) {
	reversedNumber += String(value)[i]
    }
    if (String(reversedNumber) === String(value)) {
	return true
    } else {
	return false
    }
}

/**
 * Determines the error message that should be shown to the user
 * for the given input value.
 *
 * @param {string|null|undefined} input
 * @returns {string} error message
 */
// taken from SleeplessByte
export function errorMessage(input) {
    if (!input) {
	return 'Required field'
    }
    return Number(input) ? '' : 'Must be a number besides 0'
}
