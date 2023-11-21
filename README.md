# Oculatus – Ticker Data Fetcher

This program is a simple terminal tool for monitoring ticker data. It fetches the latest prices for a set of ticker symbols every 5 minutes and presents the data in a visually appealing figlet-table format.

## Why?

I have a large wall-mounted display primarily dedicated to monitoring diverse software logs. Given the surplus terminal space available, I decided to use it for more practical data – specifically, ticker symbols of companies I'm interested in. The design ensures that I can easily see data for 6-7 tickers from any corner of the room.

Because I track only a few tickers, I employ only the initial letter of the symbol, a shorthand I find convenient due to my familiarity with the corresponding companies.

## How to Use

**Clone the Repository:**
   ```bash
   git clone git@github.com:nvrmndmnm/oculatus.git
   cd oculatus
```

**Build and Run:**
   ```bash
   go build ./cmd/oculatus
   ./oculatus META,AAPL,NFLX,GOOG,TSLA
```

**Example Output:**

![Oculatus example output](/assets/images/scr.png "Example Output")

## Acknowledgments

This program utilizes the Yahoo Finance API for fetching real-time stock prices with the [goyhfin](https://github.com/svarlamov/goyhfin) library. It is not intended for production use.


