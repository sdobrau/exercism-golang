// @ts-check

export class ArgumentError extends Error {}
export class OverheatingError extends Error {
    constructor(temperature) {
	super(`The temperature is ${temperature} ! Overheating !`);
	this.temperature = temperature;
    }
}

/**
 * Check if the humidity level is not too high.
 *
 * @param {number} humidityPercentage
 * @throws {Error}
 */
export function checkHumidityLevel(humidityPercentage) {
    if (humidityPercentage > 70) {
	throw new Error("Humidity too high")
    }
}

/**
 * Check if the temperature is not too high.
 *
 * @param {number|null} temperature
 * @throws {ArgumentError|OverheatingError}
 */
export function reportOverheating(temperature) {
    if (temperature === null) {
	throw new ArgumentError("Argument is null")
    } else if (temperature > 500) {
	throw new OverheatingError(temperature)
    }
}

export function monitorTheMachine(actions) {
    try {
	actions.check()
    } catch (error) {
	if (error instanceof ArgumentError) {
	    actions.alertDeadSensor()
	} else if (error instanceof OverheatingError) {
	    // turn on the warning light
	    if (error.temperature < 600) {
		actions.alertOverheating()
		// situation critical
	    } else if (error.temperature > 600) {
		actions.shutdown()
	    }
	} else {
	    // rethrow otherwise
	    throw error
	}
    }
}
