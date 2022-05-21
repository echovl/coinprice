# coinprice

Utility to obtain cryptocurrency prices, made for Polybar.

## Installation

```
go install github.com/echovl/coinprice@latest
```

## Usage

```ini
[module/btc]
type = custom/script
exec = coinprice BTCUSDT
tail = true

interval = 5
format = <label>
label = BTC $%output%
```
