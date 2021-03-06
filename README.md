## Galeb Client

[![Build Status](https://travis-ci.org/galeb/galeb-client.svg?branch=master)](https://travis-ci.org/galeb/galeb-client)
[![Coverage Status](https://coveralls.io/repos/github/galeb/galeb-client/badge.svg?branch=master)](https://coveralls.io/github/galeb/galeb-client?branch=master)

A easy wrapper CLI to use the [Galeb Manager](https://github.com/galeb/galeb-manager)

### Test

```bash
go test -v
```

### Build

```bash
go build -o galeb
```

### Usage

```bash
export GALEB_HOST="http://galeb.yourhost.com"
export GALEB_TOKEN="123456789"

galeb pool
```

### LICENSE

```Copyright
Copyright (c) 2014-2015 Globo.com - All rights reserved.

This source is subject to the Apache License, Version 2.0.
Please see the LICENSE file for more information.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
