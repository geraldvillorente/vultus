/*
Copyright © 2021 Gerald Villorente <geraldvillorente@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package profile

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"

	"github.com/geraldvillorente/vultus/inc"
	"github.com/geraldvillorente/vultus/layout"
)

// Info structure.
type Info struct {
	ID          string `header:"id"`
	Name        string `header:"name"`
	Company     string `header:"company"`
	Joined      string `header:"joined"`
	Followers   string `header:"followers"`
	Following   string `header:"following"`
	PublicRepos string `header:"public_repos"`
	Hireable    string `header:"hireable"`
}

// GetProfile data.
func GetProfile(handle string) {
	var (
		data layout.Profile
		info Info

		printer = tableprinter.New(os.Stdout)
	)

	endpoint := inc.Profile(handle)
	response := inc.Query(endpoint)

	byt, _ := json.Marshal(response)
	json.Unmarshal(byt, &data)

	/**
	 * Convert timestamp into date.
	 */
	format := "Jan-02-06"
	// Date created.
	created := data.CreatedAt

	info = Info{
		strconv.Itoa(data.ID),
		data.Name,
		data.Company,
		created.Format(format),
		strconv.Itoa(data.Followers),
		strconv.Itoa(data.Following),
		strconv.Itoa(data.PublicRepos),
		strconv.FormatBool(data.Hireable),
	}

	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	printer.Print(info)
}
