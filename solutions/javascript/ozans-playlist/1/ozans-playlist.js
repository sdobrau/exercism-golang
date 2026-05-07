// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Removes duplicate tracks from a playlist.
 *
 * @param {string[]} playlist
 * @returns {string[]} new playlist with unique entries
 */
export function removeDuplicates(playlist) {
    return Array.from(new Set(playlist))
}

/**
 * Checks whether a playlist includes a track.
 *
 * @param {string[]} playlist
 * @param {string} track
 * @returns {boolean} whether the track is in the playlist
 */
export function hasTrack(playlist, track) {
    return playlist.includes(track)
}

/**
 * Adds a track to a playlist.
 *
 * @param {string[]} playlist
 * @param {string} track
 * @returns {string[]} new playlist
 */
export function addTrack(playlist, track) {
    if (playlist.find((value) => value === track)) {
	// track found, don't do anything
	return playlist
    } else {
	// track not found, add it
	playlist.push(track)
	return playlist
    }
}

/**
 * Deletes a track from a playlist.
 *
 * @param {string[]} playlist
 * @param {string} track
 * @returns {string[]} new playlist
 */
export function deleteTrack(playlist, track) {
    let foundTrackIndex = playlist.findIndex((value, index) => value === track)
    // remove
    if (foundTrackIndex != -1) {
	playlist.splice(foundTrackIndex, 1)
    } 
    return playlist
}

/**
 * Lists the unique artists in a playlist.
 *
 * @param {string[]} playlist
 * @returns {string[]} list of artists
 */
export function listArtists(playlist) {
    // populate new array with artists using capture group
    // first lowercase for regexp consistency
    // first and second capture group
    let artistArray = []
    for (let song of playlist) {
	
	let nameCaptureGroup = new RegExp("(.*) - (.*)",  "")
	let match = song.match(nameCaptureGroup)
	let trackName = match[1] // first capture group
	let artistName = match[2]
	artistArray.push(artistName)
    }
    let artistSet = new Set(artistArray)
    return Array.from(artistSet)
}
