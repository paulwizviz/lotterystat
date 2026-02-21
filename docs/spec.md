# Application Specification

## Frontend Specification

- The proposed dashboard is as follows:

![img dashboard](../assets/img/dashboard.png)

- Apply this [theme](https://mui.com/material-ui/getting-started/templates/dashboard/) to the dashboard.

## Backend Specification

The backend RESTFul APIs are:

- `/` - Root endpoint delivers the web.
- `/tball/draw/frequency` - Thunderball balls frequencies.
- `/tball/tball/frequency` - Thunderball frequencies.

## App CLI Specification

- `ebz` - root command to trigger help
- `ebz --start` or `ebz -s` - root command to start frontend.
- `ebz tball` - sub command related to Thunderball draws.
- `ebz tball persists -f <filename>` - sub command to persists Thunderball csv file.
- `ebz euro` - sub command related to EuroMillions draws.
- `ebz euro persists -f <filename>` - sub command to persists EuroMillions csv file.
- `ebz lotto` - sub command related to Lotto draws.
- `ebz lotto persists -f <filename>` - sub command to persists Lotto csv file.
- `ebz sflife` - sub command related to Set For Life draws.
- `ebz sflife persists -f <filename>` - sub command to persists Set For Life csv file.
