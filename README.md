# PON-GOLANG-TYPES

## Introduction

This repository contains a collection of custom data types and structures designed to enhance the functionality and organization of other repositories. These data types are intended to be used as building blocks for various components of other projects, making the code more efficient, readable, and maintainable.
You can either clone this repo or install it as a library:

### Clone the Repository
Begin by cloning this repository to your local machine:

```bash
cd <project-directory>
git clone https://github.com/bsn-eng/pon-golang-types.git
```

### Install as library

To install the library run:

```bash
cd <project-directory>
go get github.com/bsn-eng/pon-golang-types
```

## Usage

To use simply import and start using. Example:
```
import beaconTypes "github.com/bsn-eng/pon-golang-types/beaconclient"
.
.
.
make(chan beaconTypes.HeadEventData)
```

## License

This data types library is licensed under the **MIT License**. See the `LICENSE` file in root directory for the full text of the license. You are free to use, modify, and distribute these data types in your projects, subject to the terms and conditions of the MIT License.
