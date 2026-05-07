// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */

// `'Pure Strawberry Joy'` takes 0.5 minutes, `'Energizer'` and
// `'Green Garden'` take 1.5 minutes each, `'Tropical Island'` takes 3
// minutes and `'All or Nothing'` takes 5 minutes.
export function timeToMixJuice(name) {
    switch (name) {
    case "Pure Strawberry Joy":
	return 0.5
    case "Energizer":
	return 1.5
    case "Green Garden":
	return 1.5
    case "Tropical Island":
	return 3
    case "All or Nothing":
	return 5
    default:
	return 2.5
    }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */


// Implement the function `limesToCut` which takes the number
// of lime wedges Li Mei needs to cut and an array representing
// the supply of whole limes she has at hand.
// She can get 6 wedges from a `'small'` lime, 8 wedges from a
// `'medium'` lime and 10 from a `'large'` lime.
// She always cuts the limes in the order in which they appear
// in the list, starting with the first item.
// She keeps going until she reached the number of wedges that
// she needs or until she runs out of limes.

// Li Mei would like to know in advance how many limes she
// needs to cut. The `limesToCut` function should return the
// number of limes to cut.
// limesToCut(25, ['small', 'small', 'large', 'medium', 'small']);
// => 4

// from MichaelBrig
export function limesToCut(wedgesNeeded, limes) {
    let totalWedges = 0;
    let limesCut = 0;
    let wedges = 0;  
    
    while ((limesCut < limes.length) && (totalWedges <= wedgesNeeded)) {
	switch (limes[limesCut]) {
	case "small":
            wedges = 6;
            break;
	case "medium":
            wedges = 8;
            break;
	case "large":
            wedges = 10;
            break;
	}    
	totalWedges += wedges;

	if (wedgesNeeded > 0) {
	    limesCut++;
	}
    }
    return limesCut;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
// partly inspired from MichaelBrig's solution
export function remainingOrders(timeLeft, orders) {
    let i = 0
    while (timeLeft > 0) {
	timeLeft -= timeToMixJuice(orders[0])
	orders.shift()
    }
    return orders
}
