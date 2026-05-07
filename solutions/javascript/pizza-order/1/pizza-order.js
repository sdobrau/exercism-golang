/// <reference path="./global.d.ts" />
//
// @ts-check

// Taken from SleeplessByte

/**
 * @type {Record<Pizza, number>}
 */
const PIZZA_PRICES = {
    Margherita: 7,
    Caprese: 9,
    Formaggio: 10,
};

// with one extra picked off using rest
export function pizzaPrice(pizza, ...[extra, ...otherExtras]) {
    switch (extra) {
    case 'ExtraSauce': {
	return 1 + pizzaPrice(pizza, ...otherExtras); // with one extra picked off
    }
    case 'ExtraToppings': {
	return 2 + pizzaPrice(pizza, ...otherExtras);
    }
    default: {
	return PIZZA_PRICES[pizza];
    }
    }
}
/**
 * Calculate the prize of the total order, given individual orders
 *
 * @param {PizzaOrder[]} pizzaOrders a list of pizza orders
 * @returns {number} the price of the total order
 */
export function orderPrice(pizzaOrders) {
    // result is the accumulator (price)
    // pizzaPrice called with extras
    // each order is passed to pizzaPrice with its price properties
    
    return pizzaOrders.reduce(
	(result, order) => result + pizzaPrice(order.pizza, ...order.extras),
	0
    );
}
