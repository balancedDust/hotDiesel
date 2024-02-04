package dto

// Base Dto
// Could be used as Base for High Level Dtos like IndigoFlightRoute , AirAsiaFlightRoute
// By Base being Composition for High Level
// Implementing Same Interfaces as well for Getters Setters
type FlightRoute struct {
	departureTime    string
	arrivalTime      string
	duration         string
	flightStops      string
	flightNumbers    []string
	departureAirport string
	arrivalAirport   string
	minimumFare      string
	availableFares   []string
}

func FlightRouteBuilder() FlightRoute {
	return FlightRoute{}
}

func (f FlightRoute) SetDepartureTime(d string) FlightRoute {
	f.departureTime = d
	return f
}

func (f FlightRoute) SetArrivalTime(a string) FlightRoute {
	f.arrivalTime = a
	return f
}

func (f FlightRoute) SetDurationTime(d string) FlightRoute {
	f.duration = d
	return f
}

func (f FlightRoute) SetFlightStops(fs string) FlightRoute {
	f.flightStops = fs
	return f
}

func (f FlightRoute) SetFlightNumbers(fns []string) FlightRoute {
	f.flightNumbers = fns
	return f
}

func (f FlightRoute) SetDepartureAirport(d string) FlightRoute {
	f.departureAirport = d
	return f
}

func (f FlightRoute) SetArrivalAirport(a string) FlightRoute {
	f.arrivalAirport = a
	return f
}

func (f FlightRoute) SetMinimumFare(m string) FlightRoute {
	f.minimumFare = m
	return f
}

func (f FlightRoute) SetAvailableFares(afs []string) FlightRoute {
	f.availableFares = afs
	return f
}
