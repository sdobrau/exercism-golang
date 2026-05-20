"""Functions which helps the locomotive engineer to keep track of the train."""


def get_list_of_wagons(*args):
    """Return a list of wagons, given an arbitrary amount of wagon numbers.

    Parameters:
        An arbitrary number of wagon numbers, unpacked.

    Returns:
        list: A list of wagon numbers.
    """
    *wagon_list, = args
    return wagon_list

def fix_list_of_wagons(each_wagons_id, missing_wagons):
    """Fix the list of wagons.

    Parameters:
        each_wagons_id (list[int]): The list of wagons.
        missing_wagons (list[int]): The list of missing wagons.

    Returns:
        list[int]: The corrected list of wagons.
    """

    wagons_after_1 = []
    for index, item in enumerate(each_wagons_id):
        if item == 1:
            wagons_after_1 = each_wagons_id[index+1:]
            wagons_before_1 = each_wagons_id[:index]
    [*restmis] = missing_wagons
    [*restwagafter1] = wagons_after_1
    [*restwagbefore1] = wagons_before_1

    return [1, *restmis, *restwagafter1, *restwagbefore1]


def add_missing_stops(route, **kwargs):
    """Add missing stops to route dict.

    Parameters:
        route (dict): The dict of routing information.
        (dict): An arbitrary number of stops.

    Returns:
        dict: The updated route dictionary.
    """

    stops_list = []
    route["stops"] = []
    for value in kwargs.values():
        stops_list.append(value)
        route["stops"] = stops_list
    return route

def extend_route_information(route, more_route_information):
    """Extend route information with more_route_information.

    Parameters:
            route (dict): The route information.
            more_route_information (dict): The extra route information.

    Returns:
            dict: The extended route information.
    """

    return  {**route, **more_route_information}


def fix_wagon_depot(wagons_rows):
    """Fix the list of rows of wagons.

        Parameters:
        wagons_rows (list[list[tuple]]): The list of rows of wagons.

    Returns:
        list[list[tuple]]: the list of rows of wagons.
    """
    [[a, b, c],
     [d, e, f],
     [g, h, i]] = wagons_rows
    return [[a, d, g],
            [b, e, h],
            [c, f, i]]
