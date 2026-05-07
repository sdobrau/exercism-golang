// @ts-check

/**
 * Create an appointment
 *
 * @param {number} days
 * @param {number} [now] (ms since the epoch, or undefined)
 *
 * @returns {Date} the appointment
 */
// What's wrong with this function?
export function createAppointment(days, now = Date.now()) {
  let nowTime = new Date(now)
  let nowDateDays = nowTime.getDate()
  let nowTimePlusNDays = nowDateDays + days
  nowTime.setDate(nowTimePlusNDays)
  return nowTime
}

/**
 * Generate the appointment timestamp
 *
 * @param {Date} appointmentDate
 *
 * @returns {string} timestamp
 */
export function getAppointmentTimestamp(appointmentDate) {
    return appointmentDate.toISOString()
}

/**
 * Get details of an appointment
 *
 * @param {string} timestamp (ISO 8601)
 *
 * @returns {Record<'year' | 'month' | 'date' | 'hour' | 'minute', number>} the appointment details
 */
export function getAppointmentDetails(timestamp) {
    let date = new Date(timestamp)
    return {
	'year': date.getFullYear(),
	'month': date.getMonth(),
	'date': date.getDate(),
	'hour': date.getHours(),
	'minute': date.getMinutes(),
    }
}

/**
 * Update an appointment with given options
 *
 * @param {string} timestamp (ISO 8601)
 * @param {Partial<Record<'year' | 'month' | 'date' | 'hour' | 'minute', number>>} options
 *
 * @returns {Record<'year' | 'month' | 'date' | 'hour' | 'minute', number>} the appointment details
 */
// more idiomatic way ?
export function updateAppointment(timestamp, options) {
    let date = new Date(timestamp)

    if (options["year"] != null) {
	date.setFullYear(options["year"])
    }
    if (options["month"] != null) {
	date.setMonth(options["month"]) // start from 0 up to 11
    }
    if (options["date"] != null) {
	date.setDate(options["date"])
    }
    if (options["hour"] != null) {
	date.setHours(options["hour"])
    }
    if (options["minute"] != null) {
	date.setMinutes(options["minute"])
    }

    return {
	'year': date.getFullYear(),
	'month': date.getMonth(),
	'date': date.getDate(),
	'hour': date.getHours(),
	'minute': date.getMinutes(),
    }
}

/**
 * Get available time in seconds (rounded) between two appointments
 *
 * @param {string} timestampA (ISO 8601)
 * @param {string} timestampB (ISO 8601)
 *
 * @returns {number} amount of seconds (rounded)
 */
export function timeBetween(timestampA, timestampB) {
    let dateOne = new Date(timestampA)
    let dateTwo = new Date(timestampB)
    return Math.round((dateTwo - dateOne) / 1000)
}

/**
 * Get available times between two appointment
 *
 * @param {string} appointmentTimestamp (ISO 8601)
 * @param {string} currentTimestamp (ISO 8601)
 */
export function isValid(appointmentTimestamp, currentTimestamp) {
    let appointmentDate = new Date(appointmentTimestamp)
    let currentDate = new Date(currentTimestamp)
    if ((appointmentDate.getTime() - currentDate.getTime()) > 0) {
	return true
    } else {
	return false
    }
}

