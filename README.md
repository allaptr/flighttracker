# Flight Tracker Server!
There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

 The flight tracker API accepts a POST request that includes a list of flights defined by the source and destination airport codes. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports. The successful response will show the the starting and the ending destinations of the itinerary. 
## Examples:
```
[["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
[["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
```

## Behavior
```
  $ ./flighttracker &
  [1] 21507
```
  #### Request from a linux terminal 
```
  $ curl -v -H "Content-Type:application/json"  -d '[["ATL", "EWR"], ["SFO", "ATL"]]' http://localhost:8080/calculate
  ["SFO", "EWR"]
  $
```
  ### Request from Windows Powershell
```
  > curl.exe -v -H 'Content-Type:application/json'  -d '[[\"ATL\", \"EWR\"], [\"SFO\", \"ATL\"]]' http://localhost:8080/calculate
   ["SFO", "EWR"]
```


