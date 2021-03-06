# licensor

> Inject custom license header in your source code, quickly.

### Features

- You can add custom license header to infinite files automatically.
- If you want to change license, you can. Just minor modification is required in configuration.
- You can also remove all the license header that was added by licensor.

### Installation

[Download latest binary.](https://github.com/Marvin9/licensor/releases/latest)

![](https://img.shields.io/github/v/release/Marvin9/licensor?label=latest)

Unzip tar 

**Linux/MacOS**

make binary access global.

```
sudo mv /path/to/zip/licensor /usr/local/bin
```

**Windows**

create folder to store licensor.exe at any stable place and set path in environment.

Suppose licensor.exe is at ```C:\licensor\licensor.exe```

```
SET "PATH=C:\licensor;%PATH%"
```

### Usage


**Using command line**

```
licensor -project ./ -ext go js -license ./LICENSE.md
```

<table>
    <tr>
        <th>Flag</th>
        <th>Usage</th>
        <th>Example</th>
        <th>Required</th>
        <th>Default</th>
    </tr>
    <tr>
        <td>-help</td>
        <td>.</td>
        <td>
            <code>
                licensor -help
            </code>
        </td>
        <td>false</td>
        <td></td>
    </tr>
    <tr>
        <td>-project</td>
        <td>Project directory path.</td>
        <td>
            <code>
                licensor -project ./
                <br/>
                licensor -project ./app
            </code>
        </td>
        <td>false</td>
        <td>./</td>
    </tr>
    <tr>
        <td>-ext</td>
        <td>Extensions of files which you want to add license header</td>
        <td>
            <code>
                licensor -ext go py c cpp
            </code>
        </td>
        <td>true</td>
        <td></td>
    </tr>
    <tr>
        <td>-license</td>
        <td>
        Custom license template which you want to add in source code. You can provide path or url.
        </td>
        <td>
            <code>
                licensor -license Your-License-Template.txt 
            </code>
            <br/>
            <code>
                licensor -license url-that-returns-license-text
            </code>
        </td>
        <td>true</td>
        <td></td>
    </tr>
    <tr>
        <td>-template</td>
        <td>Variables value for your license template</td>
        <td>
            <code>
                licensor -template "{\"foo\":\"bar\"}"
            </code>
        </td>
        <td>Only if license template required</td>
    </tr>
    <tr>
        <td>-ignore</td>
        <td>File(s)/Dir(s) to ignore</td>
        <td>
            <code>
                licensor -ignore ./foo ./bar/a.go
            </code>
        </td>
        <td>false</td>
        <td></td>
    </tr>
</table>

**Using YML**

```
licensor
```

- licensor.yml
```yml
project: #[project directory path] [default: "./"]
extensions: #[required]
 - go
 - py
 - c
license: #[license file path or url] [required]
# Example license template:
# Copyright {{year}} {{owner}}
# template should be
template:
 year: 2020
 owner: Mayursinh Sarvaiya
ignore:
 - ./foo #[directory]
 - ./bar/baz.go #[file]
```

### Demo

> Experiment on kubernetes 8000+ go files.

![](./assets/licensor_demo.gif)

***Commands used***:
```
# to inject license header
licensor -ext go -license ./LICENSE -ignore ./vendor

# to remove license header
licensor -ext go -remove -ignore ./vendor
```
[Read article on medium](https://medium.com/@mayursinhsarvaiya/add-license-to-your-source-code-with-licensor-34590e8b18bd)
