// @ts-check

/**
 * Double every card in the deck.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} deck with every card doubled
 */
export function seeingDouble(deck) {
    return deck.map((card) => card * 2)
}

/**
 *  Creates triplicates of every 3 found in the deck.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} deck with triplicate 3s
 */
export function threeOfEachThree(deck) {
    for (let i = 0; i < deck.length; i++)
	if (deck[i] == 3) {
	    // triple by inserting at i+1,i+2 [ 1, 2, 3] > 1, 2, 3, 3, 3
	    deck.splice(i, 0, 3)
	    deck.splice(i, 0, 3)
	    i += 2
	}
    return deck
}

/**
 * Extracts the middle two cards from a deck.
 * Assumes a deck is always 10 cards.
 *
 * @param {number[]} deck of 10 cards
 *
 * @returns {number[]} deck with only two middle cards
 */
export function middleTwo(deck) {
    return [deck[4], deck[5]]
}

/**
 * Moves the outside two cards to the middle.
 *
 * @param {number[]} deck with even number of cards
 *
 * @returns {number[]} transformed deck
 */

export function sandwichTrick(deck) {
    let middleIndex = (deck.length / 2 - 1)
    let firstNum = deck[0]
    let lastNum = deck[deck.length - 1]
    // remove first and last num
    deck.splice(0, 1)
    deck.splice(deck.length - 1, 1)

    // insert the numbers from top to bottom in the middle
    deck.splice(deck.length / 2, 0, firstNum)
    deck.splice(deck.length / 2, 0, lastNum)
    return deck
}

/**
 * Removes every card from the deck except 2s.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} deck with only 2s
 */
export function twoIsSpecial(deck) {
    let twoArray = []
    for (let i = 0; i < deck.length; i++) {
	if (deck[i] == 2) {
	    twoArray.push(2)
	}
    }
    return twoArray
}

/**
 * Returns a perfectly order deck from lowest to highest.
 *
 * @param {number[]} deck shuffled deck
 *
 * @returns {number[]} ordered deck
 */
export function perfectlyOrdered(deck) {
    deck.sort((item1, item2) => {
	if (item1 < item2) {
	    return -1; // first item smaller
	}
	if (item1 > item2) {
	    return 1; // first item larger
	}
	return 0; // equal
    });
    return deck
}

/**
 * Reorders the deck so that the top card ends up at the bottom.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} reordered deck
 */
export function reorder(deck) {
    deck.reverse()
    return deck
}
