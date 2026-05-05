// @ts-check

import { argv0 } from "node:process"

/**
 * Generates a random starship registry number.
 *
 * @returns {string} the generated registry number.
 */
export function randomShipRegistryNumber() {
    return "NCC-" + String(Math.floor(Math.random() * 9999) + 1000)
}

/**
 * Generates a random stardate.
 *
 * @returns {number} a stardate between 41000 (inclusive) and 42000 (exclusive).
 */
export function randomStardate() {
    return Math.random() * 999 + 41000
}

/**
 * Generates a random planet class.
 *
 * @returns {string} a one-letter planet class.
 */
// D, H, J, K, L, M, N, R, T, and Y.
export function randomPlanetClass() {
    let numberToLetter0to25 = { '0': "D", '1': "H", '2': "J", '3': "K", '4': "L", '5': "M", '6': "N", '7': "R",  '8': "T", '9': "Y" }
    let randNr = Math.floor(Math.random() * 10) + 0
    return numberToLetter0to25[String(randNr)]
}
