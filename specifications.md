# Specifications

At the below all specifications of Gate are listed. specifications seperated by each version released. every future
specification add to the [todo](#todo) section and will be implemented as soon as posible. Open
an [issue](https://github.com/mrNobody95/Gate/issues/new) if your desired features.

## Todo

- v0.0.1:
    - ~~Check ADX indicator~~
    - ~~Check parabolic sar indicator~~


- v0.0.3:
    - Switch to microservice

## Doing

- v0.0.2:
    - Accounting process
    - Implement fake dealing

## Done

- v0.0.1:
    - Implement [nobitex](https://apidocs.nobitex.ir) API endpoints
    - Reading system setting from YAML file
    - Starting system from command line
    - Save markets and resolutions read from setting file to database
    - Collecting candle data from brokerage
    - Calculating Bollinger band indicator
    - Calculating Stochastic indicator
    - Calculating RSI indicator
    - Calculating ADX indicator
    - Calculating EMA and SMA indicators
    - Check MACD indicator
    - Implement [coinex](https://github.com/coinexcom/coinex_exchange_api) API endpoints


- v0.0.2:
    - Check for signal of indicators
    - Create a buy or sell order
    - Check for placed orders
    - Cancel orders