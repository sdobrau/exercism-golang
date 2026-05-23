"""Functions to automate Conda airlines ticketing system."""

import random

def generate_seat_letters(number):
    """Generate a series of letters for airline seats.

    Parameters:
        number (int): Total number of seat letters to be generated.

    Returns:
        generator: A generator that yields seat letters.

    Note:
        Seat letters are generated from A to D.
        After D the sequence starts again with A.
        For example: A, B, C, D, A, B

    """

    seats = "ABCD"
    for i in range(number):
        yield seats[i % 4]
    

def generate_seats(number):
    """Generate a series of identifiers for airline seats.

    Parameters:
            number (int): The total number of seats to be generated.

    Returns:
            generator: A generator that yields seat numbers.

    Note:
            A seat number consists of the row number and the seat letter.
            There is no row 13, and each row has 4 seats.

        Seats should be sorted from low to high.
        For example: 3C, 3D, 4A, 4B

    """
    row = 1
    seat_letters = generate_seat_letters(number)
    for i in range (number):
        row = int((i+4) / 4)
        print(f"Row is {row}")
        if row >= 13:
            row = row + 1
        yield str(row) + next(seat_letters)
            

def assign_seats(passengers):
    """Assign seats to passengers.

    Parameters:
                passengers (list[str]): A list of strings containing names of passengers.

    Returns:
                dict: With passenger names as keys and seat numbers as values.
                Example output: {"Adele": "1A", "Björk": "1B"}

    """
    passenger_seat_dict = {}

    gen = generate_seats(len(passengers))
    for passenger in passengers:
        passenger_seat_dict[passenger] = next(gen)
    return passenger_seat_dict


def generate_codes(seat_numbers, flight_id):
    """Generate codes for a ticket.

    Parameters:
            seat_numbers (list[str]): A list of seat numbers.
            flight_id (str): A string containing the flight identifier.

    Returns:
            generator: A generator that yields 12 character long ticket codes.

    """
    for seat_number in seat_numbers:
        number_of_zeroes = 12 - len(seat_number + flight_id)
        zero_string = ""
        for i in range(number_of_zeroes):
            zero_string += "0"
        yield seat_number + flight_id + zero_string
