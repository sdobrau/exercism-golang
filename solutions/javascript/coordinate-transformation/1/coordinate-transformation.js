// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * @typedef {function (number, number): [number, number]} CallbackType
 */

/**
 * Create a function that returns a function making use of a closure to
 * perform a repeatable 2d translation of a coordinate pair.
 *
 * @param {number} dx the translate x component
 * @param {number} dy the translate y component
 *
 * @returns {CallbackType} a function which takes an x, y parameter, returns the
 *  translated coordinate pair in the form [x, y]
 */
export function translate2d(dx, dy) {
    return function (x, y) {
	return [dx + x, dy + y]
    }
}

/**
 * Create a function that returns a function making use of a closure to
 * perform a repeatable 2d scale of a coordinate pair.
 *
 * @param {number} sx the amount to scale the x component
 * @param {number} sy the amount to scale the y component
 *
 * @returns {CallbackType} a function which takes an x, y parameter, returns the
 *  scaled coordinate pair in the form [x, y]
 */
export function scale2d(sx, sy) {
    return function (x, y) {
	return [sx * x, sy * y];
    }
}

/**
 * Create a composition function that returns a function that combines two
 * functions to perform a repeatable transformation
 *
 * @param {CallbackType} f the first function to apply
 * @param {CallbackType} g the second function to apply
 *
 * @returns {CallbackType} a function which takes an x, y parameter, returns the
 *  transformed coordinate pair in the form [x, y]
 */
export function composeTransform(f, g) {
    return function composed (x , y) {
	// taken from SleeplessByte.
	// g takes as arg the output of f(x, y) spread
	// ...
	return g(...f(x, y))
	
    }
}

/**
 * Return a function that memoizes the last result.  If the arguments are the same as the last call,
 * then memoized result returned.
 *
 * @param {CallbackType} f the transformation function to memoize, assumes takes two arguments 'x' and 'y'
 *
 * @returns {CallbackType} a function which takes x and y arguments, and will either return the saved result
 *  if the arguments are the same on subsequent calls, or compute a new result if they are different.
 */

export function memoizeTransform(f) {
    // taken from SleeplessByte
    let previousX, previousY, previousResult // closure that encloses state
    // on subsequent runs this is compared with subsequent
    // given args
    
    return function memoised(x, y) {
	// if available then 
	if (previousX === x && previousY === y) {
	    return previousResult // just return prev res
	}
	previousX = x // set state in enclosure
	previousY = y // set state in enclosure
	return previousResult = f(x, y) //compute and store in state
    }
}
