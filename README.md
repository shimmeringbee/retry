# Shimmering Bee: Retry

[![license](https://img.shields.io/github/license/shimmeringbee/retry.svg)](https://github.com/shimmeringbee/retry/blob/master/LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg)](https://github.com/RichardLitt/standard-readme)
[![Actions Status](https://github.com/shimmeringbee/retry/workflows/test/badge.svg)](https://github.com/shimmeringbee/retry/actions)

> Simple utility to retry a task with a regenerated context each time.

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Background

Many aspects of the Shimmering Bee suite require communication with devices over a network, as such these communications
may need to be retried. Each attempt needs to have a time limit, and there must be an overall limit. 

This library uses contexts with timeouts to handle this scenario.

## Install

Add an import and most IDEs will `go get` automatically, if it doesn't `go build` will fetch.

```go
import "github.com/shimmeringbee/retry"
```

## Usage

Wrapping retry around a network function which may fail but can be safely retried. Retry assumes success if the
nested function returns nil, if err it will be immediately retried. This assumes that the network function correctly
handles a context that expires.

```go
		err := retry.Retry(context.Background(), 250*time.Millisecond, 3, func(ctx context.Context) error {
			return riskyNetworkFunction(ctx)
		})
```

## Maintainers

[@pwood](https://github.com/pwood)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/shimmeringbee/retry/issues/new) or submit PRs.

All Shimmering Bee projects follow the [Contributor Covenant](https://shimmeringbee.io/docs/code_of_conduct/) Code of Conduct.

## License

   Copyright 2019-2020 Shimmering Bee Contributors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.