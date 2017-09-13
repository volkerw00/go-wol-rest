# go-wol-rest

[![Build Status](https://travis-ci.org/voowoo/go-wol-rest.svg?branch=master)](https://travis-ci.org/voowoo/go-wol) [![Coverage Status](https://coveralls.io/repos/voowoo/go-wol-rest/badge.svg)](https://coveralls.io/r/voowoo/go-wol-rest)

Wake on LAN over REST

## Starting the server

    go-wol-rest -b ${ip}

Starts the server with an default IPv4 to broadcast to. The IP is mandatory.

## Endpoint

    /wake?mac=${mac}&broadcastIP=${ip}

### Parameters

| Name        | Description                         | Optional  |
| ------------|-------------------------------------|-----------|
| mac         | The MAC in format AB:CD:EF:12:34:56 | false     |
| broadcastIP | The IPv4 adress to broadcast to     | true      |

### Responses

* Sending the magic packet was succesful

        HTTP/1.1 204 No Content
        Date: Mon, 11 Sep 2017 18:35:15 GMT

* Failed to parse MAC

        HTTP/1.1 400 Bad Request
        Content-Type: application/json; charset=utf-8
        Date: Mon, 11 Sep 2017 18:23:30 GMT
        Content-Length: 52

        {"message":"failed to parse ${mac} as a MAC adress"}

* Failed to parse broadcast IP

        HTTP/1.1 400 Bad Request
        Content-Type: application/json; charset=utf-8
        Date: Mon, 11 Sep 2017 18:32:38 GMT
        Content-Length: 50

        {"message":"failed to parse ${ip} as a IP adress"}

* Failed to send the magic packet

        HTTP/1.1 500 Internal Server Error
        Content-Type: application/json; charset=utf-8
        Date: Mon, 11 Sep 2017 18:32:38 GMT
        Content-Length: 63

        {"message":"failed to send a magic packet for ${mac} to ${ip}"}

