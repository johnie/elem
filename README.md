# Elem CLI

## Overview

A personal cli which allows me to turn an "element" on or off by making a GET request to a specified [ifttt](https://ifttt.com) URL.

## Installation

### Using `go install`

If you have Go installed, you can install the Elem CLI directly using:

```bash
go install github.com/johnie/elem@latest
```

## Configuration

The Elem CLI uses an environment variable `ELEMEMNT_URL` to determine the base URL for the GET request. You can set this environment variable in a `.elem` file located in your home directory.

Create a `.elem` file in your home directory (`~/.elem`) with the following content:

## Usage

After installation, you can use the Elem CLI as follows:

To turn the element on:

```bash
elem on
```

To turn the element off:

```bash
elem off
```

The CLI will make a GET request to the URL specified by ELEMEMNT_URL with the action (on or off) appended to it.
