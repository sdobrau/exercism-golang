// @ts-check

/**
 * Given a certain command, help the chatbot recognize whether the command is valid or not.
 *
 * @param {string} command
 * @returns {boolean} whether or not is the command valid
 */

export function isValidCommand(command) {
    return /^Chatbot/i.test(command)
}

/**
 * Given a certain message, help the chatbot get rid of all the emoji's encryption through the message.
 *
 * @param {string} message
 * @returns {string} The message without the emojis encryption
 */
export function removeEmoji(message) {
    let removeEmojis = new RegExp("emoji[0-9]*", 'g')

    let newMessage = message.replace(removeEmojis, "")
    return newMessage
}

/**
 * Given a certain phone number, help the chatbot recognize whether it is in the correct format.
 *
 * @param {string} number
 * @returns {string} the Chatbot response to the phone Validation
 */
export function checkPhoneNumber(number) {
    if (/\(\+[0-9]{2}\) [0-9]{3}-[0-9]{3}-[0-9]{3}/.test(number)) {
	return "Thanks! You can now download me to your phone."
    } else {
	return `Oops, it seems like I can't reach out to ${number}`
    }
}

/**
 * Given a certain response from the user, help the chatbot get only the URL.
 *
 * @param {string} userInput
 * @returns {string[] | null} all the possible URL's that the user may have answered
 */
export function getURL(userInput) {
    let urlPattern = new RegExp("(https?\\://www\\.[a-z]*\\.[a-z]*)|(www\\.[a-z]+\\.[a-z]+)|([a-z]+\\.[a-z]{1,3})|(https?\\://[a-z]*\\.[a-z]*)", "g")
    
    let results = []
    for (let match of userInput.matchAll(urlPattern)) {
	results.push(match[0])
    }
    return results
}

/**
 * Greet the user using the full name data from the profile.
 *
 * @param {string} fullName
 * @returns {string} Greeting from the chatbot
 */
export function niceToMeetYou(fullName) {
    // first lowercase for regexp consistency
    let fullNameLower = fullName.toLowerCase()
    // first and second capture group
    let nameCaptureGroup = new RegExp("([a-z]*), ([a-z]*)",  "")
    let match = fullNameLower.match(nameCaptureGroup)
    let firstName = match[1] // first capture group
    // Capital case
    firstName = firstName.charAt(0).toUpperCase() + firstName.slice(1).toLowerCase()
    
    let secondName = match[2] // second capture group
	// Capital case
	secondName = secondName.charAt(0).toUpperCase() + secondName.slice(1).toLowerCase()

    return `Nice to meet you, ${secondName} ${firstName}`
}
