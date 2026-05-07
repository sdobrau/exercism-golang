/// <reference path="./global.d.ts" />

// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 * 
 * export function yourFunction(...) {
 *   ...
 * }
 
 * @param {number} remainingTime
 */
export function cookingStatus(remainingTime) {
    if (remainingTime === 0) {
	return "Lasagna is done."
    } else if (remainingTime > 0) {
	return "Not done, please wait."
    } else if (!remainingTime) { // not provided
	return "You forgot to set the timer."
    }
}

/**
 * @param {string | any[]} layers
 * @param {number} averageLayerPreparationTime
 */
export function preparationTime(layers, averageLayerPreparationTime) {
    if (!averageLayerPreparationTime) {
	averageLayerPreparationTime = 2
    }
    return layers.length * averageLayerPreparationTime
}

/**
 * @param {any} layers
 */
export function quantities (layers) {
    let noodlesQuantity = 0
    let sauceLitres = 0
    for (let i = 0; i < layers.length; i++) {
	if (layers[i] === "noodles") {
	    noodlesQuantity += 50
	} else if (layers[i] === "sauce") {
	    sauceLitres += 0.2
	}
    }
    return {'noodles': noodlesQuantity,
	    'sauce': sauceLitres,
	   }
}

export function addSecretIngredient (ingredients1, ingredients2) {
    let secretIngredient = ingredients1[ingredients1.length - 1]
    ingredients2.push(secretIngredient)
}

function scaleRecipe (recipe, portions) {
    return {'noodles': recipe['noodles'] * (portions / 2),
	    'sauce': recipe['sauce'] * (portions / 2),
	    'mozzarella': recipe['mozzarella'] * (portions / 2),
	    'meat': recipe['meat'] * (portions / 2)
	   }
}

