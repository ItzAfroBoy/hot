<div align="center">
<pre>
 __              __      
/\ \            /\ \__   
\ \ \___     ___\ \ ,_\  
 \ \  _ `\  / __`\ \ \/  
  \ \ \ \ \/\ \L\ \ \ \_ 
   \ \_\ \_\ \____/\ \__\
    \/_/\/_/\/___/  \/__/
<br>
A shitty hot reload function for your compiled projects written in Go
<br>
<img alt="GitHub License" src="https://img.shields.io/github/license/ItzAfroBoy/hot"> <img alt="GitHub tag (with filter)" src="https://img.shields.io/github/v/tag/ItzAfroBoy/hot?label=version"> <a href="https://www.codefactor.io/repository/github/itzafroboy/hot"><img src="https://www.codefactor.io/repository/github/itzafroboy/hot/badge" alt="CodeFactor" /></a> <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/ItzAfroBoy/hot">

</pre>
</div>

## Installation

### Install with go

```shell
go install github.com/ItzAfroBoy/hot@latest
hot ...
```

### Build from source

```shell
git clone https://github.com/ItzAfroBoy/hot
cd hot
go install
hot ...
```

## Usage

`Usage: hot [--cmd COMMAND] <files>...`

- `--cmd`: Command to run on file change. Arguments to be separated with `;`
