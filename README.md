# groupie-tracker-search-bar

## Description

    This is an optional subroutine of the groupie-tracker. It involves adding a new element - the search bar. The search will be performed according to the following criteria: by the name of the artists or by the names of the groups, by the name of the participants, by locations, by the date of creation or by the date of the first album.

## Autors

    The project was developed by yzhumyroб akalbiat and rneverov, yzhumyro was responsible for the front-end and assistant in the back-end, akalbiat and rneverov in turn helped with the front-end and was responsible for the back-end

## Usage

    Go to the project directory. Next, enter the command to go to the directory with the main project file:

    go run .

    After launching the project, follow this link --> http://localhost:4000/

    To see more information about an artist or a band, CLICK on their profile

## Implementation details:

    Consider this function

    What happens in this function:
    1. If the variable containing the text from the search bar is empty, then all artists are returned
    2. A variable containing the search text is converted to lowercase for search regardless of case
    3. The searched variables are created to store the artists that fall under the search and the heck variable to check for a match in the Artist Locations field
    4. A loop is set to iterate through all the artists
    5. Next are the checks: the Name field is checked first, after the First Album field is checked, and 3 the CreationDate field is checked, if there was a match, the cycle ends
    6. Next, a loop is set to iterate through the Members field
    7. If one of the elements matches, the loop ends and skips the next check
    8. If there was no match up to this point, the check for the Locations field begins and iterates through its elements using a loop
    9. Return a variable that stores matched artists
```` golang
    func Searching(searchingText string) []Artists {
    	if searchingText == "" {
		return Artist
	}
	
	searchingText = strings.ToLower(searchingText)
	var searched []Artists
	var check bool

	for i := 0; i < len(Artist); i++ {
		if strings.ToLower(Artist[i].Name) != strings.ReplaceAll(strings.ToLower(Artist[i].Name), searchingText, "") {
			searched = append(searched, Artist[i])

			continue
		}

		if strings.ToLower(Artist[i].FirstAlbum) != strings.ReplaceAll(strings.ToLower(Artist[i].FirstAlbum), searchingText, "") {
			searched = append(searched, Artist[i])

			continue
		}

		if strings.ToLower(strconv.Itoa(Artist[i].CreationDate)) != strings.ReplaceAll(strings.ToLower(strconv.Itoa(Artist[i].CreationDate)), searchingText, "") {
			searched = append(searched, Artist[i])

			continue
		}

		for k := 0; k < len(Artist[i].Members); k++ {
			check = true

			if strings.ToLower(Artist[i].Members[k]) != strings.ReplaceAll(strings.ToLower(Artist[i].Members[k]), searchingText, "") {
				searched = append(searched, Artist[i])

				check = false

				break
			}
		}

		if check {
			for j := 0; j < len(Artist[i].Locations.Locations); j++ {
				if strings.ToLower(Artist[i].Locations.Locations[j]) != strings.ReplaceAll(strings.ToLower(Artist[i].Locations.Locations[j]), searchingText, "") {
					searched = append(searched, Artist[i])

					break
				}
			}
		}
	}

	return searched

    //Using this function, a search is implemented in the subroutine groupie-tracker-search-bar
}
````
